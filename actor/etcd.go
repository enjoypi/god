package actor

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorEtcd struct {
	stdlib.SimpleActor
}

func (a *actorEtcd) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorEtcd) onStart(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	return nil
}

func newEtcd() stdlib.Actor {
	return &actorEtcd{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATEtcd, newEtcd)
}
