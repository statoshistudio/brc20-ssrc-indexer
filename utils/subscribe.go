package utils

import (
	// "errors"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Subscription struct {
	ChannelId   string    `json:"channelId"`
	ChannelName string    `json:"channelName"`
	Subscriber  string    `json:"subscriber"`
	Timestamp   int       `json:"timestamp"`
	Signature   string    `json:"signature"`
	Action      SubAction `json:"action"`
	Broadcast   bool      `json:"broadcast"`
}

func (sub *Subscription) Key() string {
	return fmt.Sprintf("/%s/%s", sub.Subscriber, sub.ChannelId)
}

func (sub *Subscription) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

type SubscriberCount struct {
	TotalSubscribers int    `json:"TotalSubscribers"`
	Channel          string `json:"channel"`
}

func ToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, i)

	fmt.Println(b)
	return b
}

func SubscriptionFromBytes(b []byte) (Subscription, error) {
	var sub Subscription
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &sub)
	return sub, err
}

func (sub *SubscriberCount) Key() string {
	return fmt.Sprintf("%d", sub.TotalSubscribers)
}

func (sub *Subscription) Hash() string {
	return hexutil.Encode(Hash(sub.ToString()))
}

func (sub *Subscription) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", sub.ChannelId))
	values = append(values, fmt.Sprintf("%s", sub.ChannelName))
	values = append(values, fmt.Sprintf("%d", sub.Timestamp))
	values = append(values, fmt.Sprintf("%s", sub.Action))
	return strings.Join(values, ",")
}
