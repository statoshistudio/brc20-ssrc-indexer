package rpc

import (
	// "errors"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	services "github.com/ByteGum/go-ssrc/pkg/service"
	utils "github.com/ByteGum/go-ssrc/utils"
	shell "github.com/ipfs/go-ipfs-api"
)

type Flag string

type IpfsData struct {
	Message string
	Subject string
}

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type RpcService struct {
	Ctx            *context.Context
	Cfg            *utils.Configuration
	MessageService *services.MessageService
	ChannelService *services.ChannelService
}

type RpcResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRpcService(mainCtx *context.Context) *RpcService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &RpcService{
		Ctx:            mainCtx,
		Cfg:            cfg,
		MessageService: services.NewMessageService(mainCtx),
		ChannelService: services.NewChannelService(mainCtx),
	}
}

func newResponse(status string, data interface{}) *RpcResponse {
	d := RpcResponse{
		Status: status,
		Data:   data,
	}
	return &d
}

// NewClient creates an http.Client that automatically perform basic auth on each request.
func NewClient(projectId, projectSecret string) *http.Client {
	return &http.Client{
		Transport: authTransport{
			RoundTripper:  http.DefaultTransport,
			ProjectId:     projectId,
			ProjectSecret: projectSecret,
		},
	}
}

// authTransport decorates each request with a basic auth header.
type authTransport struct {
	http.RoundTripper
	ProjectId     string
	ProjectSecret string
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.ProjectId, t.ProjectSecret)
	return t.RoundTripper.RoundTrip(r)
}

func (p *RpcService) SendMessage(param []byte, reply *RpcResponse) error {
	request, err := utils.JsonMessageFromBytes(param)
	if err != nil {
		return err
	}

	utils.Logger.Infof("ipfs address %s", p.Cfg.Ipfs.Host)
	client := NewClient(p.Cfg.Ipfs.ProjectId, p.Cfg.Ipfs.ProjectSecret)
	sh := shell.NewShellWithClient(p.Cfg.Ipfs.Host, client)
	chatMsg, err := utils.CreateMessageFromJson(request)
	if err != nil {
		return err
	}
	if len(request.Subject) > 0 && len(request.Message) > 0 {
		ipfsData := &IpfsData{
			Message: request.Message,
			Subject: request.Subject,
		}
		tsdBin, _ := json.Marshal(ipfsData)
		reader := bytes.NewReader(tsdBin)

		cid, err := sh.Add(reader)
		if err != nil {
			utils.Logger.Errorf("ipfs error:: %w", err)
		}
		chatMsg.Body.CID = cid
		utils.Logger.Infof("IPFS messageCID::: %s", cid)
	}

	c, err := (*p.MessageService).Send(chatMsg, request.Signature)
	if err != nil {
		return err
	}
	reply = newResponse("success", c)

	return nil
}

func (p *RpcService) Subscription(params []byte, reply *RpcResponse) error {
	request, err1 := utils.SubscriptionFromBytes(params)
	if err1 != nil {
		return err1
	}
	utils.Logger.Debug("Subscription request:::", request)
	err := (*p.ChannelService).NewChannelSubscription(&request)
	if err != nil {
		return err
	}
	reply = newResponse("success", request)
	return nil
}

// when orgs create channels...
// they sign the names to use the channel
// the signer of the message must be the signer of the name of the channel
// means they will have to put the private key of the channel...
// in case of security breach.. someone can get a hold of the private key.. which is catastrophic...
// cold storage - creating a channel should be a superior key generated by a hardware wallet
// the main private key that signs channel names should be off the server
// so they should be able to approve another key pair
// so when there is a breach.. they can approve

// methods:::
// approval process -  owners will sign and save key pair of - expiry/timestamp, public address of the sender and channels(one channel or channels(wild card))
// add sender approval signature...
// check that the signer of the approval is the signer of the message

// another issue - proof of delivery
// when someone logs in .. we generate an ephemeral receiver key value pair signature...
// can use the ephemeral key/pair to sign the delivery proof
// check if the person that signs a disposable key pair is the one that signs the delivery proof

// discuss message servers...
// we want to give ppl the ability to create their own storage
// sender will put info in the message and signs it
// we need a standard to retrieve server info... comm between final client and server will be
