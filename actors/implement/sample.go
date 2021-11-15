package implement

import (
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorSample struct {
	stdlib.SimpleActor
}

func (a *actorSample) Initialize(v *viper.Viper, sup *stdlib.Supervisor) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorSample) onStart(message def.Message) def.Message {
	return nil
}

func newSample() stdlib.Actor {
	return &actorSample{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATSample, newSample)
}
