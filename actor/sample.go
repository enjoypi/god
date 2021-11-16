package actor

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorSample struct {
	stdlib.SimpleActor
}

func (a *actorSample) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorSample) onStart(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	return nil
}

func newSample() stdlib.Actor {
	return &actorSample{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATSample, newSample)
}
