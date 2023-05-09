package processor

import (
	"context"

	db "github.com/ByteGum/go-ssrc/pkg/core/db"
	"github.com/ByteGum/go-ssrc/utils"
)

func ProcessSentMessage(ctx context.Context, unsentMessageP2pStore *db.Datastore, outMessage *utils.ClientMessage) {
	// VALIDATE AND DISTRIBUTE
	utils.Logger.Infof("\nSending out message %s\n", outMessage.Message.Body.MessageHash)
	unsentMessageP2pStore.Set(ctx, db.Key(outMessage.Key()), outMessage.ToJSON(), false)
	utils.OutgoingMessagesD_P2P_c <- outMessage
	utils.IncomingMessagesP2P2_D_c <- outMessage
	utils.Logger.Infof("\nSending out complete\n")
}
