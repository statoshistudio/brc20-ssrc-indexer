/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"net"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"

	server "github.com/ByteGum/go-ssrc/pkg/core/api"
	"github.com/ByteGum/go-ssrc/pkg/core/db"
	indexer "github.com/ByteGum/go-ssrc/pkg/core/indexer"
	p2p "github.com/ByteGum/go-ssrc/pkg/core/p2p"
	processor "github.com/ByteGum/go-ssrc/pkg/core/processor"
	rpcServer "github.com/ByteGum/go-ssrc/pkg/core/rpc"
	"github.com/ByteGum/go-ssrc/pkg/core/sql"
	ws "github.com/ByteGum/go-ssrc/pkg/core/ws"
	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
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
	errc := make(chan error)

	channelSubscriptionStore := db.New(&ctx, utils.ChannelSubscriptionStore)
	newChannelSubscriptionStore := db.New(&ctx, utils.NewChannelSubscriptionStore)

	unconfurmedBlockStore := db.New(&ctx, utils.UnconfirmedDeliveryProofStore)

	ctx = context.WithValue(ctx, utils.NewChannelSubscriptionStore, newChannelSubscriptionStore)
	// sqliteDB, sqliteDBErr := sqliteDB.InitializeDb(fmt.Sprintf("%s/indexer.db", cfg.DataDir), sqliteDB.Migrations)
	// var mu sync.Mutex
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

		if err := recover(); err != nil {
			wg.Done()
			errc <- fmt.Errorf("P2P error: %g", err)
		}
		p2p.Run(&ctx)
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
		server.HandleRequest()
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

	wg.Add(1)
	go func() {
		logger.Infof("HTTP CAll : %+s", cfg.OrdinalApi)
		page := 0

		inscriptionIdCh := make(chan string)
		pendingTransferInscriptionCh := make(chan sql.PendingTransferInscriptionModel)

		// resp, err := indexer.GetUnitDataByIdFromServer(&ctx, "e9ec244139fdd654e25085d88db97f740463bbc3757169a7a24c4bb62f10c8aci0")
		// rr, _ := json.Marshal(resp)
		// logger.Infof("TEST RUN  : %+s ", string(rr))

		wg.Add(1)
		go func() {
			for {
				resp, err := indexer.GetDataFromServer(&ctx, &page)
				if err != nil {
					logger.Fatal("indexer error: ", err)
				}

				logger.Infof("page index : %d", page)

				_list := resp.Data

				for i := len(_list) - 1; i >= 0; i-- {
					inscriptionIdCh <- _list[i]
				}
				if resp.Meta.Pagination.Next != nil {
					page = int(*resp.Meta.Pagination.Next)
				} else {
					time.Sleep(20 * time.Second)
				}

			}
		}()

		wg.Add(1)
		go func() {
			for {
				_list, err := sql.GetPendingInscriptions(sql.SqlDB)
				utils.Logger.Infof("@@@@GetPendingInscriptions err = %s  ", len(_list))
				if err != nil {
					logger.Fatal("GetPendingInscriptions error: ", err)
				}
				for i := 0; i < len(_list); i++ {

					pendingTransferInscriptionCh <- _list[i]
				}

				time.Sleep(20 * time.Second)

			}
		}()

		wg.Add(1)
		go func() {
			for {
				select {

				case id := <-inscriptionIdCh:
					// utils.Logger.Infof("+_+_+___+_++_ Number:  %d  ", id)
					func() {

						inscription, err := indexer.GetUnitDataByIdFromServer(&ctx, id)
						if err != nil {
							fmt.Println("--------")
							fmt.Println(err)
							return
						}
						acceptedTypes := [2]string{
							"text/plain;charset=utf-8",
							"application/json"}

						if acceptedTypes[0] != inscription.Inscription.ContentType && acceptedTypes[1] != inscription.Inscription.ContentType {
							return
						}
						if sql.SqlDBErr != nil {
							logger.Infof("Errr %s", sql.SqlDBErr)
						}

						var staticInscriptionStructure indexer.StaticInscriptionStructure

						err = json.Unmarshal(inscription.Inscription.Body, &staticInscriptionStructure)
						if err != nil {
							logger.Infof("Errr %s", err)
						}
						if staticInscriptionStructure.P == "brc-20" {
							content := inscription.Inscription.GetContent()
							sql.SqlDB.Transaction(func(tx *gorm.DB) error {

								if err := indexer.SaveUnitInscription(tx, inscription); err != nil {
									// return any error will rollback
									return err
								}

								if content.Op == "deploy" {
									if err := indexer.SaveNewToken(tx, &content, inscription.GenesisAddress); err != nil {
										// return any error will rollback
										return err
									}
								}
								if content.Op == "mint" {
									if err := indexer.PerformMintOperation(tx, &content, *inscription); err != nil {
										// return any error will rollback
										return err
									}
								}

								if content.Op == "transfer" {
									// genesisTransaction, err := indexer.GetUnitTransactionByIdFromServer(&ctx, inscription.GenesisAddress)
									// if err != nil {
									// 	logger.Infof("Transfer Errr : %s - %s", inscription.GenesisAddress, err)
									// 	return err
									// }
									if err := indexer.PerformTransferOperation(tx, &content, *inscription, inscription.GenesisAddress); err != nil {
										// return any error will rollback
										return err
									}
								}

								// return nil will commit the whole transaction
								return nil
							})

						}
						sql.SaveNewAccount(sql.SqlDB, inscription.GenesisAddress)
					}()
				case pendingTransferInscription := <-pendingTransferInscriptionCh:
					func() {
						utils.Logger.Infof("@@@@pendingTransferInscription err = %s  ", pendingTransferInscription.InscriptionId)
						inscription, err := indexer.GetUnitDataByIdFromServer(&ctx, pendingTransferInscription.InscriptionId)
						if err != nil {
							fmt.Println("--------")
							fmt.Println(err)
							return
						}

						if pendingTransferInscription.GenesisAddress == inscription.Address {
							logger.Infof("@@@ Thesame pendingTransferInscription.Address == inscription.Address %s == %s", pendingTransferInscription.GenesisAddress, inscription.Address)
							return
						}
						logger.Infof("@@@ @@@@@@NOT Thesame pendingTransferInscription.Address == inscription.Address %s == %s", pendingTransferInscription.GenesisAddress, inscription.Address)
						if sql.SqlDBErr != nil {
							logger.Infof("Errr %s", sql.SqlDBErr)
						}
						currentOwner := ""
						previousOwner := ""
						txAddress := strings.Split(inscription.Satpoint, ":")[0]

						for i := 0; i < 1000; i++ {
							if currentOwner == pendingTransferInscription.GenesisAddress {
								break
							}
							nextTransaction, err := indexer.GetUnitTransactionByIdFromServer(&ctx, txAddress)
							if err != nil {
								logger.Infof("genesisTransaction Errr %s", err)
							}
							previousOwner = currentOwner
							currentOwner = nextTransaction.Data.Transaction.Output[0].Address
							txAddress = strings.Split(nextTransaction.Data.Transaction.Input[0].PreviousOutput, ":")[0]

						}

						content := inscription.Inscription.GetContent()
						if err := indexer.CreditPendingOperation(sql.SqlDB, &content, *inscription, previousOwner); err != nil {
							// return any error will rollback
							logger.Infof("@@@ @@@@@@NOT Thesame AFTER Errrr %s", err)
							return
						}

						logger.Infof("@@@ @@@@@@NOT Thesame AFTER pendingTransferInscription.Address == inscription.Address %s == %s == %s", pendingTransferInscription.GenesisAddress, inscription.Address, previousOwner)

						//Perform Overations

						//Perform Overations

						//Perform Overations

						// sql.SqlDB.Delete(&pendingTransferInscription)

					}()

				}

			}
		}()

	}()

}
