/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"net"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"

	"github.com/ByteGum/go-ssrc/pkg/core/db"
	"github.com/ByteGum/go-ssrc/pkg/core/evm/abis/stake"
	processor "github.com/ByteGum/go-ssrc/pkg/core/processor"
	rpcServer "github.com/ByteGum/go-ssrc/pkg/core/rpc"
	ws "github.com/ByteGum/go-ssrc/pkg/core/ws"
	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var logger = utils.Logger

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
)

type Flag string

const (
	NODE_PRIVATE_KEY Flag = "node-private-key"
	NETWORK               = "network"
	RPC_PORT         Flag = "rpc-port"
	WS_ADDRESS       Flag = "ws-address"
)
const MaxDeliveryProofBlockSize = 1000

var deliveryProofBlockMutex sync.RWMutex

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		daemonFunc(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	daemonCmd.Flags().StringP(string(NODE_PRIVATE_KEY), "k", "", "The node private key. This is the nodes identity")
	daemonCmd.Flags().StringP(string(NETWORK), "m", MAINNET, "Network mode")
	daemonCmd.Flags().StringP(string(RPC_PORT), "p", utils.DefaultRPCPort, "RPC server port")
	daemonCmd.Flags().StringP(string(WS_ADDRESS), "w", utils.DefaultWebSocketAddress, "http service address")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
} // use default options

func daemonFunc(cmd *cobra.Command, args []string) {
	cfg := utils.Config
	ctx := context.Background()

	connectedSubscribers := map[string]map[string][]*websocket.Conn{}

	rpcPort, err := cmd.Flags().GetString(string(RPC_PORT))
	wsAddress, err := cmd.Flags().GetString(string(WS_ADDRESS))
	network, err := cmd.Flags().GetString(string(NETWORK))
	if err != nil || len(network) == 0 {
		if len(cfg.Network) == 0 {
			panic("Network required")
		}
	}
	if len(network) > 0 {
		cfg.Network = network
	}

	if rpcPort == utils.DefaultRPCPort && len(cfg.RPCPort) > 0 {
		rpcPort = cfg.RPCPort
	}

	ctx = context.WithValue(ctx, utils.ConfigKey, &cfg)
	ctx = context.WithValue(ctx, utils.IncomingMessageChId, &utils.IncomingMessagesP2P2_D_c)
	ctx = context.WithValue(ctx, utils.OutgoingMessageChId, &utils.SentMessagesRPC_D_c)
	ctx = context.WithValue(ctx, utils.OutgoingMessageDP2PChId, &utils.OutgoingMessagesD_P2P_c)
	// incoming from client apps to daemon channel
	ctx = context.WithValue(ctx, utils.SubscribeChId, &utils.SubscribersRPC_D_c)
	// daemon to p2p channel
	ctx = context.WithValue(ctx, utils.SubscriptionDP2PChId, &utils.SubscriptionD_P2P_C)
	ctx = context.WithValue(ctx, utils.ClientHandShackChId, &utils.ClientHandshakeC)
	ctx = context.WithValue(ctx, utils.OutgoingDeliveryProof_BlockChId, &utils.OutgoingDeliveryProof_BlockC)
	ctx = context.WithValue(ctx, utils.OutgoingDeliveryProofChId, &utils.OutgoingDeliveryProofC)
	ctx = context.WithValue(ctx, utils.PubsubDeliverProofChId, &utils.PubSubInputBlockC)
	ctx = context.WithValue(ctx, utils.PubSubBlockChId, &utils.PubSubInputProofC)
	// receiving subscription from other nodes channel
	ctx = context.WithValue(ctx, utils.PublishedSubChId, &utils.PublishedSubC)

	var wg sync.WaitGroup
	// errc := make(chan error)

	channelSubscriptionStore := db.New(&ctx, utils.ChannelSubscriptionStore)
	newChannelSubscriptionStore := db.New(&ctx, utils.NewChannelSubscriptionStore)

	unconfurmedBlockStore := db.New(&ctx, utils.UnconfirmedDeliveryProofStore)

	ctx = context.WithValue(ctx, utils.NewChannelSubscriptionStore, newChannelSubscriptionStore)

	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {

			case clientHandshake, ok := <-utils.ClientHandshakeC:
				if !ok {
					logger.Errorf("Verification channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go processor.ValidateMessageClient(ctx, &connectedSubscribers, clientHandshake, channelSubscriptionStore)

			case batch, ok := <-utils.PubSubInputBlockC:
				if !ok {
					logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBlockStore.Put(ctx, db.Key(batch.Key()), batch.ToJSON())
				}()
			case proof, ok := <-utils.PubSubInputProofC:
				if !ok {
					logger.Errorf("PubsubInputBlock channel closed. Please restart server to try or adjust buffer size in config")
					wg.Done()
					return
				}
				go func() {
					unconfurmedBlockStore.Put(ctx, db.Key(proof.BlockKey()), proof.ToJSON())
				}()

			}

		}
	}()
	wg.Add(1)
	go func() {
		rpc.Register(rpcServer.NewRpcService(&ctx))
		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", cfg.RPCHost+":"+rpcPort)
		if err != nil {
			logger.Fatal("ListenTCP error: ", err)
		}
		logger.Infof("RPC server runing on: %+s", cfg.RPCHost+":"+rpcPort)
		go http.Serve(listener, nil)

	}()

	wg.Add(1)
	go func() {
		wss := ws.NewWsService(&ctx)
		logger.Infof("wsAddress: %s\n", wsAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		log.Fatal(http.ListenAndServe(wsAddress, nil))
	}()

	wg.Add(1)
	go func() {
		sendHttp := rpcServer.NewHttpService(&ctx)
		err := sendHttp.Start()
		if err != nil {
			logger.Fatalf("Http error: ", err)
		}
		logger.Infof("New http connection")
	}()

}

func parserEvent(vLog types.Log, eventName string) {
	event := stake.StakeStakeEvent{}
	contractAbi, err := abi.JSON(strings.NewReader(string(stake.StakeMetaData.ABI)))

	if err != nil {
		log.Fatal("contractAbi, err", err)
	}
	_err := contractAbi.UnpackIntoInterface(&event, eventName, vLog.Data)
	if _err != nil {
		log.Fatal("_err :  ", _err)
	}

	fmt.Println(event.Account) // foo
	fmt.Println(event.Amount)
	fmt.Println(event.Timestamp)
}

var lobbyConn = []*websocket.Conn{}
var verifiedConn = []*websocket.Conn{}

// func ServeWebSocket(w http.ResponseWriter, r *http.Request) {

// 	c, err := upgrader.Upgrade(w, r, nil)
// 	log.Print("New ServeWebSocket c : ", c.RemoteAddr())

// 	if err != nil {
// 		log.Print("upgrade:", err)
// 		return
// 	}
// 	defer c.Close()
// 	hasVerifed := false
// 	time.AfterFunc(5000*time.Millisecond, func() {

// 		if !hasVerifed {
// 			c.Close()
// 		}
// 	})
// 	_close := func(code int, t string) error {
// 		logger.Infof("code: %d, t: %s \n", code, t)
// 		return errors.New("Closed ")
// 	}
// 	c.SetCloseHandler(_close)
// 	for {
// 		mt, message, err := c.ReadMessage()
// 		if err != nil {
// 			log.Println("read:", err)
// 			break

// 		} else {
// 			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
// 			if err != nil {
// 				log.Println("Error:", err)
// 			} else {
// 				// signature := string(message)
// 				verifiedRequest, _ := utils.VerificationRequestFromBytes(message)
// 				log.Println("verifiedRequest.Message: ", verifiedRequest.Message)

// 				if utils.VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
// 					verifiedConn = append(verifiedConn, c)
// 					hasVerifed = true
// 					log.Println("Verification was successful: ", verifiedRequest)
// 				}
// 				log.Println("message:", string(message))
// 				log.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
// 			}

// 		}
// 	}

// }
