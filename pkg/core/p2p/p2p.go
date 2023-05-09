package p2p

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	// "github.com/gin-gonic/gin"

	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/libp2p/go-libp2p"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/multiformats/go-multiaddr"

	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/core/routing"
	noise "github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	"github.com/sirupsen/logrus"
	// rest "messagingprotocol/pkg/core/rest"
	// dhtConfig "github.com/libp2p/go-libp2p-kad-dht/internal/config"
)

var logger = utils.Logger
var config utils.Configuration

var protocolId string
var privKey crypto.PrivKey

const DiscoveryServiceTag = "icm-network"
const (
	MessageChannel       string = "icm-message-channel"
	SubscriptionChannel         = "icm-subscription-channel"
	BatchChannel                = "icm-batch-channel"
	DeliveryProofChannel        = "icm-delivery-proof"
)

var peerStreams = make(map[string]peer.ID)
var peerPubKeys = make(map[peer.ID][]byte)
var node *host.Host
var idht *dht.IpfsDHT

type connectionNotifee struct {
}
type discoveryNotifee struct {
	h host.Host
}

// defaultNick generates a nickname based on the $USER environment variable and
// the last 8 chars of a peer ID.
func defaultNick(p peer.ID) string {
	// TODO load name from flag/config
	return fmt.Sprintf("%s-%s", os.Getenv("USER"), shortID(p))
}

// shortID returns the last 8 chars of a base58-encoded peer id.
func shortID(p peer.ID) string {
	pretty := p.Pretty()
	return pretty[len(pretty)-12:]
}

func Discover(ctx context.Context, h host.Host, kdht *dht.IpfsDHT, rendezvous string) {

	routingDiscovery := discovery.NewRoutingDiscovery(kdht)
	discovery.Advertise(ctx, routingDiscovery, rendezvous)

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

			peers, err := discovery.FindPeers(ctx, routingDiscovery, rendezvous)
			if err != nil {
				log.Fatal(err)
			}
			logger.Debugf("Found peers: %d", len(peers)-1)
			for _, p := range peers {

				if p.ID == h.ID() {
					continue
				}

				if h.Network().Connectedness(p.ID) != network.Connected {
					_, err = h.Network().DialPeer(ctx, p.ID)
					if err != nil {
						logger.Debugf("Failed to connect to peer: %s \n%s", p.ID.Pretty(), err.Error())
						h.Peerstore().RemovePeer(p.ID)
						kdht.ForceRefresh()
						continue
					}
					logger.Debugf("Connected to discovered peer: %s \n", p.ID.Pretty())
					handleConnect(&h, &p)
				}
			}
		}
	}
}

func Run(mainCtx *context.Context) {
	// fmt.Printf("publicKey %s", privateKey)
	// The context governs the lifetime of the libp2p node.
	// Cancelling it will stop the the host.

	ctx, cancel := context.WithCancel(*mainCtx)
	cfg, ok := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	if !ok {

	}
	config = *cfg
	protocolId = config.Network

	incomingMessagesC, ok := ctx.Value(utils.IncomingMessageChId).(*chan *utils.ClientMessage)
	if !ok {

	}
	outgoinMessageC, ok := ctx.Value(utils.OutgoingMessageDP2PChId).(*chan *utils.ClientMessage)
	if !ok {

	}

	subscriptionC, ok := ctx.Value(utils.SubscriptionDP2PChId).(*chan *utils.Subscription)
	if !ok {

	}

	outgoingDPBlockCh, ok := ctx.Value(utils.OutgoingDeliveryProof_BlockChId).(*chan *utils.Block)
	// outgoingProofCh, ok := ctx.Value(utils.OutgoingDeliveryProofCh).(*chan *utils.DeliveryProof)
	publishedSubscriptionC, ok := ctx.Value(utils.SubscribeChId).(*chan *utils.Subscription)
	if !ok {

	}
	defer cancel()

	if len(cfg.NodePrivateKey) == 0 {
		priv, _, err := crypto.GenerateKeyPair(
			crypto.ECDSA, // Select your key type. Ed25519 are nice short
			-1,           // Select key length when possible (i.e. RSA).
		)
		if err != nil {
			panic(err)
		}
		privKey = priv
	} else {
		priv, err := crypto.UnmarshalECDSAPrivateKey(hexutil.MustDecode(cfg.NodePrivateKey))
		if err != nil {
			panic(err)
		}
		privKey = priv
	}

	// conMgr := connmgr.NewConnManager(
	// 	100,         // Lowwater
	// 	400,         // HighWater,
	// 	time.Minute, // GracePeriod
	// )

	h, err := libp2p.New(
		// Use the keypair we generated
		libp2p.Identity(privKey),
		// Multiple listen addresses
		libp2p.ListenAddrStrings(cfg.Listeners...),
		// support TLS connections
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		// support noise connections
		libp2p.Security(noise.ID, noise.New),
		// support any other default transports (TCP)
		libp2p.DefaultTransports,
		// libp2p.Transport(ws.New),
		// Let's prevent our peer from having too many
		// connections by attaching a connection manager.

		// libp2p.ConnectionManager(connmgr.NewConnManager(
		// 	100,         // Lowwater
		// 	400,         // HighWater,
		// 	time.Minute, // GracePeriod
		// )),

		// Attempt to open ports using uPNP for NATed hosts.
		libp2p.NATPortMap(),
		// Let this host use the DHT to find other hosts

		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {

			var bootstrapPeers []peer.AddrInfo
			for _, addr := range cfg.BootstrapPeers {
				addr, _ := multiaddr.NewMultiaddr(addr)
				pi, _ := peer.AddrInfoFromP2pAddr(addr)
				bootstrapPeers = append(bootstrapPeers, *pi)
			}
			var dhtOptions []dht.Option
			dhtOptions = append(dhtOptions, dht.BootstrapPeers(bootstrapPeers...))
			if cfg.BootstrapNode {
				dhtOptions = append(dhtOptions, dht.Mode(dht.ModeServer))
			}
			kdht, err := dht.New(ctx, h, dhtOptions...)
			// validator = {
			// 	// Validate validates the given record, returning an error if it's
			// 	// invalid (e.g., expired, signed by the wrong key, etc.).
			// 	Validate(key string, value []byte) error

			// 	// Select selects the best record from the set of records (e.g., the
			// 	// newest).
			// 	//
			// 	// Decisions made by select should be stable.
			// 	Select(key string, values [][]byte) (int, error)
			// }
			// dhtOptions = append(dhtOptions, dht.NamespacedValidator("subsc"))

			idht = kdht
			if err = kdht.Bootstrap(ctx); err != nil {
				logger.Fatalf("Error starting bootstrap node %w", err)
				return nil, err
			}

			for _, addr := range cfg.BootstrapPeers {
				addr, _ := multiaddr.NewMultiaddr(addr)
				pi, err := peer.AddrInfoFromP2pAddr(addr)
				if err != nil {
					logger.Warnf("Invalid boostrap peer address (%s): %s \n", addr, err)
				} else {
					error := h.Connect(ctx, *pi)
					if error != nil {
						logger.Debugf("Unable connect to boostrap peer (%s): %s \n", addr, err)
						continue
					}
					logger.Debugf("Connected to boostrap peer (%s)", addr)
					handleConnect(&h, pi)
				}
			}
			go Discover(ctx, h, kdht, "icms")

			// routingOptions := routing.Options{
			// 	Expired: true,
			// 	Offline: true,
			// }
			// var	routingOptionsSlice []routing.Option;
			// routingOptionsSlice = append(routingOptionsSlice, routingOptions.ToOption())
			// key := "/$name/$first"
			// putErr := kdht.PutValue(ctx, key, []byte("femi"), routingOptions.ToOption())

			// if putErr != nil {
			// 	logger.Infof("Put the error %w", putErr)
			// }
			return idht, err
		}),

		// Let this host use relays and advertise itself on relays if
		// it finds it is behind NAT. Use libp2p.Relay(options...) to
		// enable active relays and more.
		// libp2p.DefaultEnableRelay(),
		//libp2p.EnableAutoRelay(),
		// If you want to help other peers to figure out if they are behind
		// NATs, you can launch the server-side of AutoNAT too (AutoRelay
		// already runs the client)
		//
		// This service is highly rate-limited and should not cause any
		// performance issues.
		libp2p.EnableNATService(),
	)

	if err != nil {
		panic(err)
	}
	h.Network().Notify(&connectionNotifee{})

	h.SetStreamHandler(protocol.ID(protocolId), handleStream)
	// create a new PubSub service using the GossipSub router
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	// setup local mDNS discovery
	err = setupDiscovery(ctx, h)
	if err != nil {
		panic(err)
	}

	node = &h

	// The last step to get fully up and running would be to connect to
	// bootstrap peers (or any other peers). We leave this commented as
	// this is an example and the peer will die as soon as it finishes, so
	// it is unnecessary to put strain on the network.

	logger.Infof("Host started with ID is %s\n", h.ID())

	messagePubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), MessageChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}
	logger.WithFields(logrus.Fields{"event": "JoinChannel", "peer": h.ID()}).Infof("Peer joined channel %s", messagePubSub.ChannelName)
	subscriptionPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), SubscriptionChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	batchPubSub, err := JoinChannel(ctx, ps, h.ID(), defaultNick(h.ID()), BatchChannel, config.ChannelMessageBufferSize)
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(5 * time.Second)
		for {
			select {
			case inMessage, ok := <-batchPubSub.Messages:
				if !ok {
					cancel()
					logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
					return
				}

				cm, err := utils.ClientMessageFromBytes(inMessage.Data)
				if err != nil {

				}
				*incomingMessagesC <- &cm
			case sub, ok := <-subscriptionPubSub.Messages:
				if !ok {
					cancel()
					logger.Fatalf("Primary Message channel closed. Please restart server to try or adjust buffer size in config")
					return
				}
				// logger.Info("Received new message %s\n", inMessage.Message.Body.Message)
				cm, err := utils.SubscriptionFromBytes(sub.Data)
				if err != nil {

				}
				logger.Info("New subscription updates:::", string(cm.ToJSON()))
				// *incomingMessagesC <- &cm
				cm.Broadcast = false
				*publishedSubscriptionC <- &cm
			}
		}
	}()

	for {
		select {
		case outMessage, ok := <-*outgoinMessageC:
			if cfg.Validator {
				if !ok {
					logger.Errorf("Outgoing channel closed. Please restart server to try or adjust buffer size in config")
					return
				}
				err := messagePubSub.Publish(utils.NewSignedPubSubMessage(outMessage.ToJSON(), cfg.EvmPrivateKey))
				if err != nil {
					logger.Errorf("Failed to publish message. Please restart server to try or adjust buffer size in config")
					return
				}
			}
		case subscription, ok := <-*subscriptionC:
			if cfg.Validator {
				if !ok {
					logger.Errorf("Subscription channel not found in the context")
					return
				}
				logger.Info("subscription channel:::", subscription.ChannelId)

				err := subscriptionPubSub.Publish(utils.NewSignedPubSubMessage(subscription.ToJSON(), cfg.EvmPrivateKey))
				if err != nil {
					logger.Errorf("Failed to publish subscription.")
					return
				}
			}
		case block, ok := <-*outgoingDPBlockCh:
			if cfg.Validator {
				if !ok {
					logger.Errorf("Subscription channel not found in the context")
					return
				}
				logger.Info("subscription channel:::", block.BlockId)
				err := batchPubSub.Publish(utils.NewSignedPubSubMessage(block.ToJSON(), cfg.EvmPrivateKey))
				if err != nil {
					logger.Errorf("Failed to publish subscription.")
					return
				}
			}
		}
	}

}

func handleStream(stream network.Stream) {
	// logger.Debugf("Got a new stream! %s", stream.ID())
	// stream.SetReadDeadline()
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go readData(peerStreams[stream.ID()], rw)
	// go sendData(rw)

}

func readData(p peer.ID, rw *bufio.ReadWriter) {
	for {
		hsString, err := rw.ReadString('\n')
		if err != nil {
			logger.Errorf("Error reading from buffer %w", err)
			panic(err)
		}
		if hsString == "" {
			break
		}

		logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Info("New Handshake data from peer")
		handshake, err := utils.HandshakeFromJSON(hsString)

		if err != nil {
			logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Warnf("Failed to parse handshake: %w", err)
			break
		}
		validHandshake := isValidHandshake(handshake, p)
		if !validHandshake {
			disconnect(*node, p)
			logger.WithFields(logrus.Fields{"peer": p, "data": hsString}).Warnf("Disconnecting from peer (%s) with invalid handshake", p)
			return
		}
		validStake := isValidStake(handshake, p)
		if !validStake {
			disconnect(*node, p)
			logger.WithFields(logrus.Fields{"address": handshake.Signer, "data": hsString}).Warnf("Disconnecting from peer (%s) with inadequate stake in network", p)
			return
		}
		b, _ := hexutil.Decode(handshake.Signer)
		peerPubKeys[p] = b
		break
	}
}

func isValidHandshake(handshake utils.Handshake, p peer.ID) bool {
	handshakeMessage := handshake.ToJSON()
	if math.Abs(float64(handshake.Data.Timestamp-int(time.Now().Unix()))) > utils.VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"peer": p, "data": handshakeMessage}).Warnf("Hanshake Expired: %s", handshakeMessage)
		return false
	}
	message := handshake.Data.ToString()
	isValid := utils.VerifySignature(handshake.Signer, message, handshake.Signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": message, "signature": handshake.Signature}).Warnf("Invalid signer %s", handshake.Signer)
		return false
	}
	logger.Debugf("New Valid handshake from peer: %s", p)
	return true
}
func isValidStake(handshake utils.Handshake, p peer.ID) bool {
	// if handshake.Data.NodeType == utils.ValidatorNodeType && config.Validator {

	// }
	return true
}
func sendData(p peer.ID, rw *bufio.ReadWriter, data []byte) {

	// defer disconnect(*node, p)
	_, err := rw.WriteString(fmt.Sprintf("%s\n", string(data)))
	if err != nil {
		logger.Warn("Error writing to to stream")
		return
	}
	err = rw.Flush()
	if err != nil {
		// fmt.Println("Error flushing buffer")
		// panic(err)
		logger.Warn("Error flushing to to stream")
		return
	}
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	logger.Debugf("Discovered new peer %s\n", pi.ID.Pretty())
	err := n.h.Connect(context.Background(), pi)

	if err != nil {
		logger.Warningf("Unable to connect with peer: %s %w", pi.ID, err)
		return
	}
	handleConnect(&n.h, &pi)
}

func handleConnect(h *host.Host, pa *peer.AddrInfo) {
	pi := *pa
	logger.Debugf("Successfully connected to peer: %s", pi.ID)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := (*h).NewStream(ctx, pi.ID, protocol.ID(protocolId))

	if err != nil {
		logger.Warningf("Unable to establish stream with peer: %s %w", pi.ID, err.Error())
	} else {
		logger.Infof("Streaming to peer: %s", pi.ID)
		rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
		logger.Infof("New StreamID: %s", stream.ID())
		peerStreams[stream.ID()] = pi.ID
		nodeType := utils.RelayNodeType
		if config.Validator {
			nodeType = utils.ValidatorNodeType
		}
		hs := utils.CreateHandshake(defaultNick((*h).ID()), protocolId, config.EvmPrivateKey, nodeType)
		go sendData(pi.ID, rw, (&hs).ToJSON())
	}
}

func disconnect(h host.Host, id peer.ID) {
	h.Network().ClosePeer(id)
}

// DiscoveryServiceTag is used in our mDNS advertisements to discover other chat peers.

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func setupDiscovery(ctx context.Context, h host.Host) error {
	n := discoveryNotifee{h: h}
	// n.h = make(chan peer.AddrInfo)
	// setup mDNS discovery to find local peers
	disc := mdns.NewMdnsService(h, DiscoveryServiceTag, &n)
	// if err := disc.Start(); err != nil {
	// 	panic(err)
	// }
	// // disc.RegisterNotifee(&n)
	return disc.Start()
}

// Listen is called when network starts listening on an addr
func (n *connectionNotifee) Listen(netw network.Network, ma multiaddr.Multiaddr) {}

// ListenClose is called when network starts listening on an addr
func (n *connectionNotifee) ListenClose(netw network.Network, ma multiaddr.Multiaddr) {}

// Connected is called when a connection opened
func (n *connectionNotifee) Connected(netw network.Network, conn network.Conn) {
	//retain max 4 connections
	// if (len(netw.Conns()) > 4){
	// 	conn.Close()
	// 	fmt.Printf("Connection refused for peer: %v!\n", conn.RemotePeer().Pretty())
	// }a
}

// Disconnected is called when a connection closed
func (cn *connectionNotifee) Disconnected(netw network.Network, conn network.Conn) {
	id := conn.RemotePeer()
	logger.Infof("Peer disconnect: %s", id)
	idht.Host().Peerstore().RemovePeer(id)
}

// OpenedStream is called when a stream opened
func (cn *connectionNotifee) OpenedStream(netw network.Network, stream network.Stream) {}

// ClosedStream is called when a stream was closed
func (cn *connectionNotifee) ClosedStream(netw network.Network, stream network.Stream) {}
