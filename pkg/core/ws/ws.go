package ws

import (
	// "errors"
	"context"
	"log"
	"net/http"
	"time"

	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/gorilla/websocket"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type WsService struct {
	Ctx                    *context.Context
	Cfg                    *utils.Configuration
	ClientHandshakeChannel *chan *utils.ClientHandshake
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewWsService(mainCtx *context.Context) *WsService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	verificationc, _ := (*mainCtx).Value(utils.ClientHandShackChId).(*chan *utils.ClientHandshake)
	return &WsService{
		Ctx:                    mainCtx,
		Cfg:                    cfg,
		ClientHandshakeChannel: verificationc,
	}
}

func newResponse(status string, data interface{}) *RpcResponse {
	d := RpcResponse{
		Status: status,
		Data:   data,
	}
	return &d
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (p *WsService) ServeWebSocket(w http.ResponseWriter, r *http.Request) {

	c, err := upgrader.Upgrade(w, r, nil)
	log.Print("New ServeWebSocket c : ", c.RemoteAddr())

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	hasVerifed := false
	time.AfterFunc(5000*time.Millisecond, func() {
		if !hasVerifed {
			c.Close()
		}
	})
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break

		} else {
			err = c.WriteMessage(mt, (append(message, []byte("recieved Signature")...)))
			if err != nil {
				log.Println("Error:", err)
			} else {
				// signature := string(message)
				verifiedRequest, _ := utils.ClientHandshakeFromBytes(message)
				verifiedRequest.Socket = c
				log.Println("verifiedRequest.Message: ", verifiedRequest.Message)

				if utils.VerifySignature(verifiedRequest.Signer, verifiedRequest.Message, verifiedRequest.Signature) {
					// verifiedConn = append(verifiedConn, c)
					hasVerifed = true
					log.Println("Verification was successful: ", verifiedRequest)
					*p.ClientHandshakeChannel <- &verifiedRequest
				}
				log.Println("message:", string(message))
				log.Printf("recv: %s - %d - %s\n", message, mt, c.RemoteAddr())
			}

		}
	}

}
