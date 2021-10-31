package actors

import (
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const actorTypeEtcd = "etcd"

type actorEtcd struct {
	stdlib.DefaultActor
}

func (a *actorEtcd) Initialize(v *viper.Viper) error {
	_ = a.DefaultActor.Initialize()
	a.Register(types.EvStart, a.onStart)
	return nil
}

func (a *actorEtcd) onStart(message types.Message) types.Message {
	logger.L.Debug("onStart",
		zap.String("actor", a.Type()),
	)

	return nil
}

func newEtcd() stdlib.Actor {
	return &actorEtcd{}
}

func init() {
	stdlib.RegisterActorCreator(actorTypeEtcd, newEtcd)
}
