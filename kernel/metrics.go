package kernel

import (
	"math/rand"

	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
)

type heartbeat struct {
	stdlib.DefaultActor
}

func (h heartbeat) Handle(message types.Message) types.Message {
	panic("implement me")
}

func newHeartbeat() stdlib.Actor {
	return &heartbeat{}
}

func init() {
	stdlib.RegisterActorCreator(rand.Int63(), newHeartbeat)
}
