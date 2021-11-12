package implement

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
)

type actorEtcd struct {
	actors.SimpleActor
}

func (a *actorEtcd) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorEtcd) onStart(message types.Message) types.Message {
	return nil
}

func newEtcd() actors.Actor {
	return &actorEtcd{}
}

func init() {
	actors.RegisterActorCreator(types.ATEtcd, newEtcd)
}
