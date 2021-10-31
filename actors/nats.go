package actors

import (
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

const actorTypeNats = "nats"

type actorNats struct {
	stdlib.DefaultActor
}

func (a *actorNats) Initialize() error {
	_ = a.DefaultActor.Initialize()
	a.Register(types.EvStart, a.onStart)
	return nil
}

func (a *actorNats) onStart(message types.Message) types.Message {
	stdlib.L.Debug("onStart",
		zap.String("actor", a.Type()),
	)

	return nil
}

func newActorNats() stdlib.Actor {
	return &actorNats{}
}

func init() {
	stdlib.RegisterActorCreator(actorTypeNats, newActorNats)
}
