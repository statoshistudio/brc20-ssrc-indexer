package utils

import (
	// "errors"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type ClientHandshake struct {
	Signature string          `json:"signature"`
	Signer    string          `json:"signer"`
	Message   string          `json:"message"`
	Socket    *websocket.Conn `json:"socket"`
}

func (sub *ClientHandshake) ToJSON() []byte {
	m, e := json.Marshal(sub)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func ClientHandshakeFromBytes(b []byte) (ClientHandshake, error) {
	var verMsg ClientHandshake
	// if err := json.Unmarshal(b, &message); err != nil {
	// 	panic(err)
	// }
	err := json.Unmarshal(b, &verMsg)
	return verMsg, err
}
