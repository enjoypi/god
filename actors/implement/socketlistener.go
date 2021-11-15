package implement

import (
	"net"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/stdlib"
	"github.com/spf13/viper"
)

type conf struct {
	Network   string
	Address   string
	ConnActor string
}

type actorSocketListener struct {
	stdlib.SimpleActor
	conf
	listener net.Listener
	*viper.Viper
	sup *stdlib.Supervisor
}

func (a *actorSocketListener) Initialize(v *viper.Viper, sup *stdlib.Supervisor) error {
	if err := v.Unmarshal(&a.conf); err != nil {
		return err
	}

	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)

	a.Viper = v
	a.sup = sup
	return nil
}

func (a *actorSocketListener) onStart(message def.Message) def.Message {

	listener, err := net.Listen(a.Network, a.Address)
	if err != nil {
		return err
	}
	a.listener = listener

	go stdlib.Catch(func() {
		for {
			conn, err := a.listener.Accept()
			logger.CheckError("net accept", err)

			actor, err := a.sup.Start(a.Viper, def.GetActorType(a.conf.ConnActor), 0)
			logger.CheckError("start net actor", err)

			actor.Post(&events.EvNetConnected{Conn: conn})
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
