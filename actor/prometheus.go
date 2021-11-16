package actor

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

const actorTypePrometheus = "prometheus"

type actorPrometheus struct {
	stdlib.SimpleActor
}

func (a *actorPrometheus) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorPrometheus) onStart(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	return nil
}
func newPrometheus() stdlib.Actor {
	return &actorPrometheus{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATPrometheus, newPrometheus)
}
