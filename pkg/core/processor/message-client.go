package processor

import (
	"context"
	"sync"

	db "github.com/ByteGum/go-ssrc/pkg/core/db"
	"github.com/ByteGum/go-ssrc/utils"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/websocket"
	"github.com/ipfs/go-datastore/query"
)

var logger = utils.Logger

func ValidateMessageClient(
	ctx context.Context,
	connectedSubscribers *map[string]map[string][]*websocket.Conn,
	clientHandshake *utils.ClientHandshake,
	channelSubscriberStore *db.Datastore) {
	// VALIDATE AND DISTRIBUTE
	utils.Logger.Infof("Signer:  %s\n", clientHandshake.Signer)
	results, err := channelSubscriberStore.Query(ctx, query.Query{
		Prefix: "/" + clientHandshake.Signer,
	})
	if err != nil {
		utils.Logger.Errorf("Channel Subscriber Store Query Error %w", err)
		return
	}
	entries, _err := results.Rest()
	for i := 0; i < len(entries); i++ {
		_sub, _ := utils.SubscriptionFromBytes(entries[i].Value)
		if (*connectedSubscribers)[_sub.ChannelId] == nil {
			(*connectedSubscribers)[_sub.ChannelId] = map[string][]*websocket.Conn{}
		}
		(*connectedSubscribers)[_sub.ChannelId][_sub.Subscriber] = append((*connectedSubscribers)[_sub.ChannelId][_sub.Subscriber], clientHandshake.Socket)
	}
	utils.Logger.Infof("results:  %s  -  %w\n", entries[0].Value, _err)
}

func ValidateAndAddToDeliveryProofToBlock(ctx context.Context,
	proof *utils.DeliveryProof,
	deliveryProofStore *db.Datastore,
	channelSubscriberStore *db.Datastore,
	stateStore *db.Datastore,
	localBlockStore *db.Datastore,
	MaxBlockSize int,
	mutex *sync.RWMutex,
) {
	err := deliveryProofStore.Set(ctx, db.Key(proof.Key()), proof.ToJSON(), true)
	if err == nil {
		// msg, err := validMessagesStore.Get(ctx, db.Key(fmt.Sprintf("/%s/%s", proof.MessageSender, proof.MessageHash)))
		// if err != nil {
		// 	// invalid proof or proof has been tampered with
		// 	return
		// }
		// get signer of proof
		susbscriber, err := utils.GetSigner(proof.ToString(), proof.Signature)
		if err != nil {
			// invalid proof or proof has been tampered with
			return
		}
		// check if the signer of the proof is a member of the channel
		isSubscriber, err := channelSubscriberStore.Has(ctx, db.Key("/"+susbscriber+"/"+proof.MessageHash))
		if isSubscriber {
			// proof is valid, so we should add to a new or existing batch
			var block *utils.Block
			var err error
			txn, err := stateStore.NewTransaction(ctx, false)
			if err != nil {
				utils.Logger.Errorf("State query errror %w", err)
				// invalid proof or proof has been tampered with
				return
			}
			blockData, err := txn.Get(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey))
			if err != nil {
				logger.Errorf("State query errror %w", err)
				// invalid proof or proof has been tampered with
				txn.Discard(ctx)
				return
			}
			if len(blockData) > 0 && block.Size < MaxBlockSize {
				block, err = utils.BlockFromBytes(blockData)
				if err != nil {
					logger.Errorf("Invalid batch %w", err)
					// invalid proof or proof has been tampered with
					txn.Discard(ctx)
					return
				}
			} else {
				// generate a new batch
				block = utils.NewBlock()

			}
			block.Size += 1
			if block.Size >= MaxBlockSize {
				block.Closed = true
				block.NodeHeight = utils.GetNodeHeight()
			}
			// save the proof and the batch
			block.Hash = hexutil.Encode(utils.Hash(proof.Signature + block.Hash))
			err = txn.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey), block.ToJSON())
			if err != nil {
				logger.Errorf("Unable to update State store errror %w", err)
				txn.Discard(ctx)
				return
			}
			proof.Block = block.BlockId
			proof.Index = block.Size
			err = deliveryProofStore.Put(ctx, db.Key(proof.Key()), proof.ToJSON())
			if err != nil {
				txn.Discard(ctx)
				logger.Errorf("Unable to save proof to store error %w", err)
				return
			}
			err = localBlockStore.Put(ctx, db.Key(utils.CurrentDeliveryProofBlockStateKey), block.ToJSON())
			if err != nil {
				logger.Errorf("Unable to save batch error %w", err)
				txn.Discard(ctx)
				return
			}
			err = txn.Commit(ctx)
			if err != nil {
				logger.Errorf("Unable to commit state update transaction errror %w", err)
				txn.Discard(ctx)
				return
			}
			// dispatch the proof and the batch
			if block.Closed {
				utils.OutgoingDeliveryProof_BlockC <- block
			}
			utils.OutgoingDeliveryProofC <- proof

		}

	}

}
