package actor

import (
	"context"
	"net"

	"go.uber.org/zap"

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
}

func (a *actorSocketListener) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)

	a.Viper = v
	return nil
}

func (a *actorSocketListener) onStart(ctx context.Context, message def.Message, args ...interface{}) def.Message {

	opt := ctx.Value("option").(option.Listen)
	listener, err := net.Listen(opt.Network, opt.Address)
	if err != nil {
		return err
	}
	a.listener = listener

	go stdlib.Catch(func() {
		for {
			conn, err := a.listener.Accept()
			logger.PanicOnError("net accept", err)
			logger.L.Debug("socket connected",
				zap.String("network", conn.LocalAddr().Network()),
				zap.String("local", conn.LocalAddr().String()),
				zap.String("remote", conn.RemoteAddr().String()),
			)

			sup := args[0].(*stdlib.Supervisor)

			var actor stdlib.Actor
			switch conn.(type) {
			case *net.TCPConn:
				actor, err = sup.Start(a.Viper, def.ATTcp, 0)
			case *net.UDPConn:
				actor, err = sup.Start(a.Viper, def.ATUdp, 0)
			}

			if err != nil {
				logger.L.Error("start actor failed", zap.Error(err))
				continue
			}

			actor.Post(ctx, &event.EvSocketConnected{Conn: conn})
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
