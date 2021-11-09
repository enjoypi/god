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

	logger.L.Info("initialize NATS",
		zap.String("options", fmt.Sprintf("%+v", opts)))

	a.RegisterReaction(error(nil), a.onError)
	a.RegisterReaction((*types.EvStart)(nil), a.onStart)
	a.RegisterReaction((*types.Subscription)(nil), a.onSubscribe)
	return nil
}

func (a *actorNats) onError(message types.Message) types.Message {
	logger.L.Error("error message", zap.Error(message.(error)))
	return nil
}

func (a *actorNats) onStart(message types.Message) types.Message {
	opts := a.Options
	nc, err := opts.Connect()
	if err != nil {
		logger.L.Warn("connect NATS", zap.Error(err), zap.String("url", opts.Url))
		opts.RetryOnFailedConnect = true
		_, _ = opts.Connect()
		return nil
	}
	conn = nc
	logger.L.Info("NATS connected", zap.String("url", conn.ConnectedUrl()))
	a.Post(&types.Subscription{Subject: actorTypeNats})
	return nil
}

func (a *actorNats) onSubscribe(message types.Message) types.Message {
	subscription := message.(*types.Subscription)
	if conn == nil {
		return fmt.Errorf("no NATS connection")
	}

	handler :=
		func(msg *nats.Msg) {
			m := natsMsg2Message(msg)
			logger.CheckError("handle NATS msg", Post2Actor(subscription.ActorID, m))
		}
	_, err := conn.Subscribe(subscription.Subject, handler)

	if err != nil {
		return err
	}
	return nil
}

func (a *actorNats) onDisconnected(nc *nats.Conn, err error) {
	conn = nil
	logger.L.Warn("NATS disconnected", zap.Error(err), zap.String("url", nc.Opts.Url))
}

func (a *actorNats) onReconnected(nc *nats.Conn) {
	if conn != nc {
		conn = nc
	}
	logger.L.Info("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
}

func natsMsg2Message(msg *nats.Msg) types.Message {
	return string(msg.Data)
}

func Post2Actor(id types.ActorID, message types.Message) error {
	logger.L.Debug("post to actor", zap.Uint16("actor", id), zap.Any("message", message))
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
