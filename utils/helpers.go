package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
func ToSqlDateTime(t time.Time) string {
	t = t.UTC()
	formatted := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return formatted
}

func IsValidChannel(ch ChatMessageHeader, signature string, channelOwner string) bool {

	signer, _ := GetSigner(ch.ToApprovalString(), signature)
	if strings.ToLower(channelOwner) != strings.ToLower(signer) {
		return false
	}
	if math.Abs(float64(int(ch.ChannelExpiry)-int(time.Now().Unix()))) > VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": ch}).Warnf("Channel Expired: %s", ch.ChannelExpiry)
		return false
	}
	channel := ch.ToApprovalString()
	isValid := VerifySignature(signer, channel, signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": channel, "signature": signature}).Warnf("Invalid signer %s", signer)
		return false
	} else {

	}
	return true
}

func IsValidMessage(msg ChatMessage, signature string) bool {
	chatMessage := msg.ToJSON()
	signer, _ := GetSigner(msg.ToString(), signature)
	channel := strings.Split(msg.Header.Receiver, ":")
	channelOwner, _ := GetSigner(strings.ToLower(channel[0]), channel[1])
	if strings.ToLower(channelOwner) != strings.ToLower(signer) {
		return false
	}
	if !IsValidChannel(msg.Header, channel[1], channelOwner) {
		return false
	}
	if math.Abs(float64(int(msg.Header.Timestamp)-int(time.Now().Unix()))) > VALID_HANDSHAKE_SECONDS {
		logger.WithFields(logrus.Fields{"data": chatMessage}).Warnf("ChatMessage Expired: %s", chatMessage)
		return false
	}
	message := msg.ToString()
	isValid := VerifySignature(signer, message, signature)
	if !isValid {
		logger.WithFields(logrus.Fields{"message": message, "signature": signature}).Warnf("Invalid signer %s", signer)
		return false
	} else {

	}
	return true
}

func IsValidSubscription(
	subscription Subscription,
	verifyTimestamp bool,
) bool {
	if verifyTimestamp {
		if math.Abs(float64(int(subscription.Timestamp)-int(time.Now().Unix()))) > VALID_HANDSHAKE_SECONDS {
			logger.Info("Invalid Subscription, invalid handshake duration")
			return false
		}
	}
	return VerifySignature(subscription.Subscriber, subscription.ToString(), subscription.Signature)
}

func CreateMessageFromJson(msg MessageJsonInput) (ChatMessage, error) {

	if len(msg.Message) > 0 {
		msgHash := hexutil.Encode(Hash(msg.Message))
		if msg.MessageHash != msgHash {
			return ChatMessage{}, errors.New("Invalid Message")
		}
	}
	if len(msg.Subject) > 0 {
		subHash := hexutil.Encode(Hash(msg.Subject))
		if msg.SubjectHash != subHash {
			return ChatMessage{}, errors.New("Invalid Subject")
		}
	}
	chatMessage := ChatMessageHeader{
		Timestamp:     uint(msg.Timestamp),
		Approval:      msg.Approval,
		Receiver:      msg.Receiver,
		ChainId:       msg.ChainId,
		Platform:      msg.Platform,
		Length:        100,
		ChannelExpiry: msg.ChannelExpiry,
		Channels:      msg.Channels,
		SenderAddress: msg.SenderAddress,
		// OwnerAddress:  msg.OwnerAddress,
	}

	bodyMessage := ChatMessageBody{
		SubjectHash: msg.SubjectHash,
		MessageHash: msg.MessageHash,
	}
	_chatMessage := ChatMessage{chatMessage, bodyMessage, msg.Actions, msg.Origin}
	return _chatMessage, nil
}
