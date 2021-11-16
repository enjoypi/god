package sample

import (
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

const name = "sample"

func init() {
	stdlib.RegisterApplication(name, newSample)
}

type sample struct {
	sup *stdlib.Supervisor

	discovery stdlib.Actor
	monitor   stdlib.Actor
	messaging stdlib.Actor
}

func newSample(v *viper.Viper) (def.Application, error) {
	sup, err := stdlib.NewSupervisor()
	if err != nil {
		return nil, err
	}

	return &sample{
		sup: sup,
	}, nil
}

func (k *sample) PrepareStop() {

}

func (k *sample) Name() string {
	return name
}

func (k *sample) Start(v *viper.Viper) error {
	var opt option
	if err := v.Unmarshal(&opt); err != nil {
		return err
	}

	for _, a := range opt.Actors {
		actor, err := k.sup.Start(v, a.ActorType, a.ActorID)
		if err != nil {
			return err
		}
		actor.Post(&events.EvStart{})
	}
	return nil
}

func (k *sample) Stop() {

}
