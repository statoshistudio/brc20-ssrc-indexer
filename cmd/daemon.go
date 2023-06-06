/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"net"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"

	ordApi "github.com/ByteGum/go-ssrc/pkg/core/api"
	"github.com/ByteGum/go-ssrc/pkg/core/db"
	indexer "github.com/ByteGum/go-ssrc/pkg/core/indexer"
	rpcServer "github.com/ByteGum/go-ssrc/pkg/core/rpc"
	"github.com/ByteGum/go-ssrc/pkg/core/sql"
	ws "github.com/ByteGum/go-ssrc/pkg/core/ws"
	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

var logger = &utils.Logger

const (
	TESTNET string = "/icm/testing"
	MAINNET        = "/icm/mainnet"
)

type Flag string

const (
	NODE_PRIVATE_KEY Flag = "node-private-key"
	NETWORK               = "network"
	RPC_PORT         Flag = "rpc-port"
	RPC_HTTP_PORT    Flag = "rpc-http-port"
	WS_ADDRESS       Flag = "ws-address"
	API_PORT         Flag = "api-port"
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
	daemonCmd.Flags().StringP(string(RPC_HTTP_PORT), "r", utils.DefaultRPCPort, "RPC http client port")
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

	rpcPort, err := cmd.Flags().GetString(string(RPC_PORT))
	rpcHttpPort, err := cmd.Flags().GetString(string(RPC_HTTP_PORT))
	rpcHttpPort = cfg.RPCHttpPort
	if len(cfg.RPCPort) == 0 {
		rpcHttpPort = utils.DefaultRPCHttpPort
	}
	cfg.RPCHttpPort = rpcHttpPort

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

	rpcPort = cfg.RPCPort
	if len(cfg.RPCPort) == 0 {
		rpcPort = utils.DefaultRPCPort
	}
	cfg.RPCPort = rpcPort

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

	newChannelSubscriptionStore := db.New(&ctx, utils.NewChannelSubscriptionStore)

	ctx = context.WithValue(ctx, utils.NewChannelSubscriptionStore, newChannelSubscriptionStore)

	defer wg.Wait()

	wg.Add(1)
	go func() {
		rpc.Register(rpcServer.NewRpcService(&ctx))
		rpc.HandleHTTP()
		listener, err := net.Listen("tcp", cfg.RPCHost+":"+rpcPort)
		if err != nil {
			logger.Fatal("ListenTCP error: ", err)
		}

		go http.Serve(listener, nil)
		time.Sleep(3000)
		logger.Infof("RPC server listening on: %+s", cfg.RPCHost+":"+rpcPort)
		sendHttp := rpcServer.NewHttpService(&ctx)
		err = sendHttp.Start()
		if err != nil {
			logger.Fatalf("Http error: %s", err.Error())
		}
		logger.Infof("New http connection")

	}()

	wg.Add(1)
	go func() {
		wss := ws.NewWsService(&ctx)
		logger.Infof("Websocket server listing on: %s\n", wsAddress)
		http.HandleFunc("/echo", wss.ServeWebSocket)

		log.Fatal(http.ListenAndServe(wsAddress, nil))
	}()

	wg.Add(1)
	go func() {
		ordApi.HandleRequest()
	}()

	wg.Add(1)
	go func() {
		logger.Infof("Ordinal data source: %+s", cfg.OrdinalApi)
		page := 0
		_config, err := sql.GetConfig(sql.SqlDB, utils.LastIndexedPageKey)
		if err == nil {
			page, err = strconv.Atoi(_config.Value)
			logger.Infof("Starting from page %d", page)
			if err != nil {
				page = 0
			}
		} else {
			logger.Infof("sql.GetConfig Errr %s", err)
		}

		inscriptionIdCh := make(chan string)
		pendingTransferInscriptionCh := make(chan sql.PendingTransferInscriptionModel)

		// resp, err := indexer.GetUnitDataByIdFromServer(&ctx, "e9ec244139fdd654e25085d88db97f740463bbc3757169a7a24c4bb62f10c8aci0")
		// rr, _ := json.Marshal(resp)
		// logger.Infof("TEST RUN  : %+s ", string(rr))

		// wg.Add(1)
		// go func() {
		// 	for {
		// 		resp, err := indexer.GetDataFromServer(&ctx, &page)
		// 		if err != nil {
		// 			logger.Fatal("indexer error: ", err)
		// 			panic(err)
		// 		}

		// 		logger.Infof("page index : %d", page)

		// 		_list := resp.Data

		// 		for i := len(_list) - 1; i >= 0; i-- {
		// 			inscriptionIdCh <- _list[i]
		// 		}
		// 		_page := strconv.Itoa(page)
		// 		_, configError := sql.SetConfig(sql.SqlDB, utils.LastIndexedPageKey, _page)
		// 		if configError != nil {
		// 			logger.Errorf("Setting last indexed page sql error: %s", err)
		// 		}
		// 		if resp.Meta.Pagination.Next != nil {
		// 			page = int(*resp.Meta.Pagination.Next)
		// 		} else {
		// 			time.Sleep(20 * time.Second)
		// 		}

		// 	}
		// }()
		wg.Add(1)
		go func() {
			for {
				time.Sleep(4 * time.Second)
				logger.Info("Checking updated inscription...")
				inscriptions, err := sql.GetUpdatedInscriptions(sql.SqlDB, 200)
				if err != nil {
					logger.Infof("Sql error %s", err.Error())
					continue
				}
				logger.Infof("Found %d updated inscriptions", len(inscriptions))
				for _, dbInscription := range inscriptions {
					inscription_id := dbInscription.InscriptionId

					inscription, err := indexer.GetUnitDataByIdFromServer(&ctx, inscription_id)
					if err != nil {
						continue
					}
					if inscription.Satpoint == dbInscription.Satpoint {
						if dbInscription.CreatedAt.Before(time.Now().Add(-24 * time.Hour)) {
							sql.DeleteUpdatedInscription(sql.SqlDB, dbInscription.ID)
						}
						continue
					}
					_, err = indexer.HandleCallback(sql.SqlDB, *inscription)
					if err != nil {
						continue
					}

				}

			}
		}()

		wg.Add(1)
		go func() {
			for {
				_list, err := sql.GetPendingInscriptions(sql.SqlDB)
				utils.Logger.Infof("@@@@GetPendingInscriptions err = %d  ", len(_list))
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
							logger.Errorf("Error %s", err.Error())
							panic(err)

						}
						if err := indexer.SaveGenericInscription(sql.SqlDB, inscription); err != nil {
							// return any error will rollback
							if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
								logger.Debugf("Errr %s", err.Error())
								return
							}
							panic(err)

						}
						if sql.SqlDBErr != nil {
							logger.Errorf("SQLDB ERROR %s", sql.SqlDBErr)
						}
						acceptedTypes := []string{
							"text/plain;charset=utf-8",
							"application/json"}

						if slices.Index(acceptedTypes, inscription.Inscription.ContentType) == -1 {
							return
						}

						var staticInscriptionStructure indexer.StaticInscriptionStructure

						err = json.Unmarshal(inscription.Inscription.Body, &staticInscriptionStructure)
						if err != nil {
							logger.WithFields(logrus.Fields{"id": inscription.InscriptionId}).Infof("Body not a json: %s", err.Error())
							return
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
						inscription, err := sql.GetUnitGenericInscription(sql.SqlDB, pendingTransferInscription.InscriptionId)
						// sql.SaveNewAccount(sql.SqlDB, inscription.GenesisAddress)
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
							logger.Infof("Error Getting pending  %s", sql.SqlDBErr)
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

						// content := inscription.Inscription.GetContent()
						var content indexer.InscriptionStructure
						err = json.Unmarshal([]byte(inscription.InscriptionBody), &content)
						if err != nil {
							log.Println("Unmershaling error:", err)
							return
						}
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
