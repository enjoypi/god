package socket

import (
	"context"

	"github.com/enjoypi/god/option"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

const name = "socket"

func init() {
	stdlib.RegisterApplication(name, newSocket)
}

type Socket struct {
	sup *stdlib.Supervisor
}

func newSocket(v *viper.Viper) (def.Application, error) {
	sup, err := stdlib.NewSupervisor()
	if err != nil {
		return nil, err
	}

	return &Socket{
		sup: sup,
	}, nil
}

func (k *Socket) PrepareStop() {

}

func (k *Socket) Name() string {
	return name
}

func (k *Socket) Start(v *viper.Viper) error {
	type optionSocket struct {
		Listener []option.Listen
	}
	var opt struct {
		Socket optionSocket
	}
	if err := v.Unmarshal(&opt); err != nil {
		return err
	}

	for _, l := range opt.Socket.Listener {
		actor, err := k.sup.Start(v, l.ActorType, l.ActorID)
		if err != nil {
			return err
		}

		ctx := context.WithValue(context.Background(), "option", l)
		actor.Post(ctx, &event.EvStart{})
	}
	return nil
}

func (k *Socket) Stop() {

}
