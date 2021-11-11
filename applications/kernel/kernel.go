package kernel

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
)

const name = "kernel"

func init() {
	stdlib.RegisterApplication(name, newKernel)
}

type kernel struct {
	sup *stdlib.Supervisor

	discovery actors.Actor
	monitor   actors.Actor
	messaging actors.Actor
}

func newKernel(v *viper.Viper) (types.Application, error) {
	sup, err := stdlib.NewSupervisor()
	if err != nil {
		return nil, err
	}

	return &kernel{
		sup: sup,
	}, nil
}

func (k *kernel) PrepareStop() {

}

func (k *kernel) Name() string {
	return name
}

func (k *kernel) Start(v *viper.Viper) error {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	cfg.ActorType = []string{"etcd", "nats", "prometheus"}
	for _, actorType := range cfg.ActorType {
		actor, err := k.sup.Start(v, actorType)
		if err != nil {
			return err
		}
		actor.Post(&events.EvStart{})
	}
	return nil
}

func (k *kernel) Stop() {

}
