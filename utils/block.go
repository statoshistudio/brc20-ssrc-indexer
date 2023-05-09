package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	// "math"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Block
type Block struct {
	BlockId    string `json:"id"`
	Size       int    `json:"s"`
	Closed     bool   `json:"cls"`
	NodeHeight int    `json:"height"`
	Hash       string `json:"hash"`
	Timestamp  int    `json:"ts"`
	sync.Mutex
}

func (msg *Block) ToJSON() []byte {
	m, _ := json.Marshal(msg)
	return m
}

func (msg *Block) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", string(msg.BlockId)))
	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.Size)))
	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.NodeHeight)))
	values = append(values, fmt.Sprintf("%s", strconv.Itoa(msg.Timestamp)))
	values = append(values, fmt.Sprintf("%s", msg.Hash))
	return strings.Join(values, ",")
}

func (msg *Block) Key() string {
	return fmt.Sprintf("/%s", msg.BlockId)
}

// func (msg *Block) Sign(privateKey string) Block {

// 	msg.Timestamp = int(time.Now().Unix())
// 	_, sig := Sign(msg.ToString(), privateKey)
// 	msg.Signature = sig
// 	return *msg
// }

func NewBlock() *Block {
	id, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_", 32)
	return &Block{BlockId: id,
		Size:   0,
		Closed: false}
}

func BlockFromBytes(b []byte) (*Block, error) {
	var message Block
	err := json.Unmarshal(b, &message)
	return &message, err
}
