package actor

import (
	"context"
	"net"

	"github.com/enjoypi/god/option"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorSocketListener struct {
	stdlib.SimpleActor
	listener net.Listener
	*viper.Viper
	sup *stdlib.Supervisor
}

func (a *actorSocketListener) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)

	a.Viper = v
	//a.sup = sup
	return nil
}

func (a *actorSocketListener) onStart(ctx context.Context, message def.Message) def.Message {

	opt := ctx.Value("option").(option.Listen)
	listener, err := net.Listen(opt.Network, opt.Address)
	if err != nil {
		return err
	}
	a.listener = listener

	go stdlib.Catch(func() {
		for {
			conn, err := a.listener.Accept()
			logger.CheckError("net accept", err)

			actor, err := a.sup.Start(a.Viper, opt.Handler, 0)
			logger.CheckError("start net actor", err)

			actor.Post(context.Background(), &event.EvSocketConnected{Conn: conn})
		}
	})
	return nil
}

func newSocketListener() stdlib.Actor {
	return &actorSocketListener{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATSocketListener, newSocketListener)
}
