package actors

import (
	"fmt"

	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/stdlib"
	"github.com/enjoypi/god/types"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const actorTypeNats = "nats"

var conn *nats.Conn

type actorNats struct {
	stdlib.DefaultActor
	nats.Options
	*viper.Viper
}

func (a *actorNats) Initialize(v *viper.Viper) error {
	_ = a.DefaultActor.Initialize()

	type config struct {
		Nats nats.Options
	}
	var cfg config

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	a.Options = nats.GetDefaultOptions()
	opts := &a.Options
	opts.Url = cfg.Nats.Url
	opts.DisconnectedErrCB = a.onDisconnected
	opts.ReconnectedCB = a.onReconnected

	a.Register(types.EvStart, a.onStart)
	return nil
}

func (a *actorNats) onStart(message types.Message) types.Message {
	logger.L.Debug("onStart",
		zap.String("actor", a.Type()),
	)

	nc, err := a.Options.Connect()
	if err != nil {
		return err
	}
	conn = nc
	Subscribe(">")
	logger.L.Debug("NATS connected", zap.String("url", conn.ConnectedUrl()))
	return nil
}

func (a *actorNats) onDisconnected(nc *nats.Conn, err error) {
	logger.L.Debug("NATS disconnected", zap.Error(err))
}

func (a *actorNats) onReconnected(nc *nats.Conn) {
	logger.L.Debug("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
	if conn != nc {
		conn = nc
	}
}

func Subscribe(subj string) error {
	if conn == nil {
		return fmt.Errorf("no NATS connection")
	}
	_, err := conn.Subscribe(subj, handleNatsMsg)
	if err != nil {
		return err
	}
	return nil
}

func handleNatsMsg(msg *nats.Msg) {
	m := natsMsg2Message(msg)
	logger.Warn("handle NATS msg", PostActor(actorTypeNats, m))
}

func natsMsg2Message(msg *nats.Msg) types.Message {
	return types.EvNone
}

func PostActor(name string, message types.Message) error {
	return nil
}

//func (a *Transport) Run() error {
//	sub, err := a.Conn.SubscribeSync(">")
//	if err != nil {
//		return err
//	}
//
//	for {
//		msg, err := sub.NextMsg(time.Hour)
//		if err != nil {
//			break
//		}
//		a.PostEvent(msg)
//	}
//	return sub.Unsubscribe()
//}

func newActorNats() stdlib.Actor {
	return &actorNats{}
}

func init() {
	stdlib.RegisterActorCreator(actorTypeNats, newActorNats)
}
