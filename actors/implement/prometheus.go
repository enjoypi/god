package implement

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/spf13/viper"
)

const actorTypePrometheus = "prometheus"

type actorPrometheus struct {
	actors.SimpleActor
}

func (a *actorPrometheus) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorPrometheus) onStart(message def.Message) def.Message {
	return nil
}
func newPrometheus() actors.Actor {
	return &actorPrometheus{}
}

func init() {
	actors.RegisterActorCreator(def.ATPrometheus, newPrometheus)
}
