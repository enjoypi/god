package actors

import (
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
)

const actorTypePrometheus = "prometheus"

type actorPrometheus struct {
	stdlib.DefaultActor
}

func (a *actorPrometheus) Initialize(v *viper.Viper) error {
	_ = a.DefaultActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorPrometheus) onStart(message types.Message) types.Message {
	return nil
}
func newPrometheus() stdlib.Actor {
	return &actorPrometheus{}
}

func init() {
	stdlib.RegisterActorCreator(actorTypePrometheus, newPrometheus)
}
