package sample

import (
	"github.com/enjoypi/god/actors"
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

	discovery actors.Actor
	monitor   actors.Actor
	messaging actors.Actor
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
	var c conf
	if err := v.Unmarshal(&c); err != nil {
		return err
	}

	for _, a := range c.Actors {
		actor, err := k.sup.Start(v, def.String2Type(a.Type), a.ActorID)
		if err != nil {
			return err
		}
		actor.Post(&events.EvStart{})
	}
	return nil
}

func (k *sample) Stop() {

}
