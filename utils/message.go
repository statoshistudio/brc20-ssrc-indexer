package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	PlatformEthereum string = "ethereum"
	PlatformBitcoin         = "bitcoin"
	PlatformSolana          = "solana"
	PlatformCosmos          = "cosmos"
)

/**
CHAT MESSAGE
**/
type ChatMessageHeader struct {
	Length        int    `json:"length"`
	Sender        string `json:"from"`
	Approval      string `json:"approval"`
	Receiver      string `json:"reciever"`
	ChainId       string `json:"chainId"`
	Platform      string `json:"platform"`
	Timestamp     uint   `json:"timestamp"`
	ChannelExpiry int    `json:"channelExpiry"`
	// Wildcard      bool   `json:"wildcard"`
	Channels      []string `json:"channels"`
	SenderAddress string   `json:"senderAddress"`
	// OwnerAddress  string `json:"OwnerAddress"`
}

// TODO! platform enum channel
//! receiver field is name of channel u are sending to
// ! look for all subscribers to the channel
//! channel subscribers store
type ChatMessageBody struct {
	SubjectHash string `json:"subjectHash"`
	MessageHash string `json:"messageHash"`
	CID         string `json:"cid"`
}
type ChatMessageAction struct {
	Contract   string   `json:"contract"`
	Abi        string   `json:"abi"`
	Action     string   `json:"action"`
	Parameters []string `json:"parameters"`
}

type ChatMessage struct {
	Header  ChatMessageHeader   `json:"header"`
	Body    ChatMessageBody     `json:"body"`
	Actions []ChatMessageAction `json:"actions"`
	Origin  string              `json:"origin"`
}

func (chatMessage *ChatMessage) ToString() string {
	values := []string{}

	values = append(values, fmt.Sprintf("%s", chatMessage.Header.Receiver))
	values = append(values, fmt.Sprintf("%s", chatMessage.Header.Approval))
	values = append(values, fmt.Sprintf("%s", chatMessage.Header.ChainId))
	values = append(values, fmt.Sprintf("%s", chatMessage.Header.Platform))
	values = append(values, fmt.Sprintf("%d", chatMessage.Header.Timestamp))

	values = append(values, fmt.Sprintf("%s", chatMessage.Body.SubjectHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.MessageHash))
	values = append(values, fmt.Sprintf("%s", chatMessage.Body.CID))
	_action := []string{}
	for i := 0; i < len(chatMessage.Actions); i++ {
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Contract))
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Abi))
		_action = append(_action, fmt.Sprintf("[%d]:%s", i, chatMessage.Actions[i].Action))

		_parameter := []string{}
		for j := 0; j < len(chatMessage.Actions[i].Parameters); j++ {
			_parameter = append(_parameter, fmt.Sprintf("[%d][%d]:%s", i, j, chatMessage.Actions[i].Parameters[j]))
		}

		_action = append(_action, fmt.Sprintf("[%d]:%s", i, _parameter))
	}

	values = append(values, fmt.Sprintf("%s", _action))
	values = append(values, fmt.Sprintf("%s", chatMessage.Origin))

	return strings.Join(values, ",")
}

func (channel *ChatMessageHeader) ToApprovalString() string {
	values := []string{}

	values = append(values, fmt.Sprintf("ChannelExpiry:%d", channel.ChannelExpiry))
	// values = append(values, fmt.Sprintf("Wildcard:%s", channel.Wildcard))
	values = append(values, fmt.Sprintf("Channels:%s", channel.Channels))
	values = append(values, fmt.Sprintf("SenderAddress:%s", channel.SenderAddress))
	// values = append(values, fmt.Sprintf("OwnerAddress:%s", channel.OwnerAddress))
	return strings.Join(values, ",")
}

/**
NODE MESSAGE
**/
type ClientMessage struct {
	Message         ChatMessage `json:"message"`
	SenderSignature string      `json:"senderSignature"`
	NodeSignature   string      `json:"nodeSignature"`
}

type SuccessResponse struct {
	body ClientMessage
	meta Meta
}

type ErrorResponse struct {
	statusCode int
	meta       Meta
}

type Meta struct {
	statusCode int
	success    bool
}

func ReturnError(msg string, code int) *ErrorResponse {
	meta := Meta{statusCode: code}
	meta.success = false
	e := ErrorResponse{statusCode: code}
	e.meta = meta
	return &e
}

func (msg *ClientMessage) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}
func (msg *ClientMessage) ToString() string {
	return fmt.Sprintf("%s:%s:%s", msg.Message.ToString(), msg.SenderSignature, msg.NodeSignature)
}

func (msg *ClientMessage) Hash() []byte {
	bs := crypto.Keccak256Hash([]byte(msg.Message.ToString())).Bytes()
	return bs
}

func (msg *ClientMessage) Key() string {
	return fmt.Sprintf("/%s/%s", msg.Message.Header.Sender, hexutil.Encode(msg.Hash()))
}

func ClientMessageFromBytes(b []byte) (ClientMessage, error) {
	var message ClientMessage
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &message)
	return message, err
}

func JsonMessageFromBytes(b []byte) (MessageJsonInput, error) {
	var message MessageJsonInput
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &message)
	return message, err
}

func ClientMessageFromString(msg string) (ClientMessage, error) {
	return ClientMessageFromBytes([]byte(msg))
}

func (msg *ChatMessage) ToJSON() string {
	m, _ := json.Marshal(msg)
	return string(m)
}

func ChatMessageFromBytes(b []byte) ChatMessage {
	var message ChatMessage
	if err := json.Unmarshal(b, &message); err != nil {
		panic(err)
	}
	return message
}

func ChatMessageFromString(msg string) ChatMessage {
	return ChatMessageFromBytes([]byte(msg))
}

type MessageJsonInput struct {
	Timestamp     int      `json:"timestamp"`
	Approval      string   `json:"approval"`
	ChannelExpiry int      `json:"channelExpiry"`
	Channels      []string `json:"channels"`
	SenderAddress string   `json:"senderAddress"`
	// OwnerAddress  string              `json:"OwnerAddress"`
	Receiver    string              `json:"receiver"`
	Platform    string              `json:"platform"`
	ChainId     string              `json:"chainId"`
	Type        string              `json:"type"`
	Message     string              `json:"message"`
	Subject     string              `json:"subject"`
	Signature   string              `json:"signature"`
	Actions     []ChatMessageAction `json:"actions"`
	Origin      string              `json:"origin"`
	MessageHash string              `json:"messageHash"`
	SubjectHash string              `json:"subjectHash"`
}

// PubSubMessage
type PubSubMessage struct {
	Data      json.RawMessage `json:"data"`
	Timestamp string          `json:"timestamp"`
	Signature string          `json:"signature"`
}

func (msg *PubSubMessage) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func (msg *PubSubMessage) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("Data:%s", string(msg.Data)))
	values = append(values, fmt.Sprintf("Timestmap%s", msg.Timestamp))
	return strings.Join(values, ",")
}

func NewSignedPubSubMessage(data []byte, privateKey string) PubSubMessage {
	timestamp := int(time.Now().Unix())
	message := PubSubMessage{Data: data, Timestamp: strconv.Itoa(timestamp)}
	_, sig := Sign(message.ToString(), privateKey)
	message.Signature = sig
	return message
}

func PubSubMessageFromBytes(b []byte) (PubSubMessage, error) {
	var message PubSubMessage
	err := json.Unmarshal(b, &message)
	return message, err
}
