package actor

import (
	"context"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/option"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type actorTCP struct {
	stdlib.SimpleActor
	*viper.Viper
}

func (a *actorTCP) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*event.EvStart)(nil), a.onStart)
	a.RegisterReaction((*event.EvSocketConnected)(nil), a.onSocketConnected)
	return nil
}

func (a *actorTCP) onStart(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	return nil
}

func (a *actorTCP) onSocketConnected(ctx context.Context, message def.Message, args ...interface{}) def.Message {
	opt := ctx.Value("option").(option.Listen)
	//conn := message.(*event.EvSocketConnected).Conn.(*net.TCPConn)
	sup := args[0].(*stdlib.Supervisor)

	actor, err := sup.Start(a.Viper, opt.Handler, 0)
	logger.PanicOnError("new tcp handler", err)

	actor.Post(ctx, message)

	//conn.CloseWrite()
	//time.Sleep(10 * time.Second)
	//conn.CloseRead()
	return nil
}

func newTCP() stdlib.Actor {
	return &actorTCP{}
}

func init() {
	stdlib.RegisterActorCreator(def.ATTcp, newTCP)
}
