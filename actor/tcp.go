package actor

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorTCP struct {
	stdlib.SimpleActor
}

func (a *actorTCP) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorTCP) onStart(ctx context.Context, message def.Message) def.Message {
	return nil
}

func newTCP() stdlib.Actor {
	return &actorTCP{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATSample, newTCP)
}
