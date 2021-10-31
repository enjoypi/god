package actors

import (
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

const actorTypePrometheus = "prometheus"

type actorPrometheus struct {
	stdlib.DefaultActor
}

func (a *actorPrometheus) Initialize() error {
	_ = a.DefaultActor.Initialize()
	a.Register(types.EvStart, a.onStart)
	return nil
}

func (a *actorPrometheus) onStart(message types.Message) types.Message {
	stdlib.L.Debug("onStart",
		zap.String("actor", a.Type()),
	)

	return nil
}
func newPrometheus() stdlib.Actor {
	return &actorPrometheus{}
}

func init() {
	stdlib.RegisterActorCreator(actorTypePrometheus, newPrometheus)
}
