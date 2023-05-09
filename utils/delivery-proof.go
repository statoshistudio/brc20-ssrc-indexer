package utils

import (
	"encoding/json"
	"fmt"
	"strconv"

	// "math"
	"strings"
)

// DeliveryProof
type DeliveryProof struct {
	MessageHash   string `json:"mHash"`
	MessageSender string `json:"sender"`
	NodeAddress   string `json:"addr"`
	Timestamp     int    `json:"ts"`
	Signature     string `json:"sign"`
	Block         string `json:"block"`
	Index         int    `json:"id"`
}

func (msg *DeliveryProof) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func (msg *DeliveryProof) Key() string {
	return fmt.Sprintf("/%s/%s", msg.MessageHash, msg.MessageSender)
}
func (msg *DeliveryProof) BlockKey() string {
	return fmt.Sprintf("/%s", msg.Block)
}

func (msg *DeliveryProof) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", string(msg.MessageHash)))
	values = append(values, fmt.Sprintf("%s", msg.NodeAddress))
	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.Timestamp)))
	return strings.Join(values, ",")
}

// func NewSignedDeliveryProof(data []byte, privateKey string) DeliveryProof {
// 	message, _ := DeliveryProofFromBytes(data)
// 	_, sig := Sign(message.ToString(), privateKey)
// 	message.Signature = sig
// 	return message
// }

func DeliveryProofFromBytes(b []byte) (DeliveryProof, error) {
	var message DeliveryProof
	err := json.Unmarshal(b, &message)
	return message, err
}

// DeliveryClaim
type DeliveryClaim struct {
	NodeHeight int      `json:"nodeHeight"`
	Signature  string   `json:"signature"`
	Amount     string   `json:"amount"`
	Proofs     []string `json:"proofs"`
}

func (msg *DeliveryClaim) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func DeliveryClaimFromBytes(b []byte) (DeliveryClaim, error) {
	var message DeliveryClaim
	err := json.Unmarshal(b, &message)
	return message, err
}
