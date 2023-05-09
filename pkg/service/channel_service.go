package services

import (
	// "errors"
	"context"
	"errors"

	"github.com/ByteGum/go-ssrc/pkg/core/db"
	utils "github.com/ByteGum/go-ssrc/utils"
)

type ChannelService struct {
	Ctx context.Context
	Cfg utils.Configuration
}

func NewChannelService(mainCtx *context.Context) *ChannelService {
	ctx := *mainCtx
	cfg, _ := ctx.Value(utils.ConfigKey).(*utils.Configuration)
	return &ChannelService{
		Ctx: ctx,
		Cfg: *cfg,
	}
}

func (p *ChannelService) NewChannelSubscription(sub *utils.Subscription) error {
	// subscribersc, ok := p.Ctx.Value(utils.SubscribeChId).(*chan *utils.Subscription)

	// validate before storing
	if utils.IsValidSubscription(*sub, true) {
		channelSubscriberStore, ok := p.Ctx.Value(utils.NewChannelSubscriptionStore).(*db.Datastore)
		if !ok {
			return errors.New("Could not connect to subscription datastore")
		}
		error := channelSubscriberStore.Set(p.Ctx, db.Key(sub.Key()), sub.ToJSON(), false)
		if error != nil {
			return error
		}
	}
	return nil
}
