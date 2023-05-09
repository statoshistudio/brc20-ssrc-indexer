package services

import (
	// "errors"
	"context"
	"errors"
	"fmt"
	"strings"

	utils "github.com/ByteGum/go-ssrc/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Flag string

// !sign web3 m
// type msgError struct {
// 	code int
// 	message string
// }

type MessageService struct {
	Ctx context.Context
	Cfg utils.Configuration
}

type Subscribe struct {
	channel   string
	timestamp string
}

func NewMessageService(mainCtx *context.Context) *MessageService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	return &MessageService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *MessageService) Send(chatMsg utils.ChatMessage, senderSignature string) (*utils.ClientMessage, error) {
	if strings.ToLower(chatMsg.Origin) != strings.ToLower(utils.GetPublicKey(p.Cfg.EvmPrivateKey)) {
		return nil, errors.New("Invalid Origin node address: " + chatMsg.Origin + " is not")
	}
	if utils.IsValidMessage(chatMsg, senderSignature) {
		channel := strings.Split(chatMsg.Header.Receiver, ":")

		if utils.Contains(chatMsg.Header.Channels, "*") || utils.Contains(chatMsg.Header.Channels, strings.ToLower(channel[0])) {

			privateKey := p.Cfg.EvmPrivateKey

			// TODO:
			// if its an array check the channels .. if its * allow
			// message server::: store messages, require receiver to request message through an endpoint

			signature, _ := utils.Sign(senderSignature, privateKey)
			message := utils.ClientMessage{}
			message.Message = chatMsg
			message.SenderSignature = senderSignature
			message.NodeSignature = hexutil.Encode(signature)
			outgoingMessageC, ok := p.Ctx.Value(utils.OutgoingMessageChId).(*chan *utils.ClientMessage)
			if !ok {
				utils.Logger.Error("Could not connect to outgoing channel")
				panic("outgoing channel fail")
			}
			*outgoingMessageC <- &message
			fmt.Printf("Testing my function%s, %s", chatMsg.ToString(), chatMsg.Body.SubjectHash)
			return &message, nil
		}
	}
	return nil, fmt.Errorf("Invalid message signer")
}
