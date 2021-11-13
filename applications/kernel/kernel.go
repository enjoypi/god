package kernel

import (
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

const name = "kernel"

func init() {
	stdlib.RegisterApplication(name, newKernel)
}

type kernel struct {
	sup *stdlib.Supervisor
}

func newKernel(v *viper.Viper) (def.Application, error) {
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
	for _, a := range def.KernelActors {
		actor, err := k.sup.Start(v, a.ActorType, a.ActorID)
		if err != nil {
			return err
		}
		actor.Post(&events.EvStart{})
	}
	return nil
}

func (k *kernel) Stop() {

}
