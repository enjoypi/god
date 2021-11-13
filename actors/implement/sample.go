package implement

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/spf13/viper"
)

type actorSample struct {
	actors.SimpleActor
}

func (a *actorSample) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorSample) onStart(message def.Message) def.Message {
	return nil
}

func newSample() actors.Actor {
	return &actorSample{}
}

func init() {
	actors.RegisterActorCreator(def.ATSample, newSample)
}
