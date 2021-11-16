package implement

import (
	"fmt"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/options"
	"github.com/enjoypi/god/stdlib"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var conn *nats.Conn

type actorNats struct {
	stdlib.SimpleActor
	nats.Options

	optNode options.Node
	*viper.Viper
}

func (a *actorNats) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()

	var opt struct {
		Nats nats.Options
		options.Node
	}
	if err := v.Unmarshal(&opt); err != nil {
		return err
	}
	a.optNode = opt.Node

	a.Options = nats.GetDefaultOptions()
	opts := &a.Options
	opts.Url = opt.Nats.Url
	opts.DisconnectedErrCB = a.onDisconnected
	opts.ReconnectedCB = a.onReconnected

	logger.L.Info("initialize NATS",
		zap.Uint32("actor", a.ID()),
		zap.String("options", fmt.Sprintf("%+v", opts)))

	a.RegisterReaction(error(nil), a.onError)
	a.RegisterReaction("", a.onString)
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorNats) onError(message def.Message) def.Message {
	logger.L.Error("error message", zap.Error(message.(error)))
	return nil
}

func (a *actorNats) onStart(message def.Message) def.Message {
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

	subject := fmt.Sprintf("%d.*", a.optNode.ID)
	_, err = conn.Subscribe(subject, a.onMsg)
	if err != nil {
		return err
	}
	return nil
}

func (a *actorNats) onString(message def.Message) def.Message {
	logger.L.Debug("on string", zap.String("message", message.(string)))
	return nil
}

func (a *actorNats) onDisconnected(nc *nats.Conn, err error) {
	stdlib.Catch(func() {
		conn = nil
		logger.L.Warn("NATS disconnected", zap.Error(err), zap.String("url", nc.Opts.Url))
		a.Post(events.EvBusDisconnected{})
	})
}

func (a *actorNats) onReconnected(nc *nats.Conn) {
	stdlib.Catch(func() {
		if conn != nc {
			conn = nc
		}
		logger.L.Info("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
		a.Post(events.EvBusReconnected{})
	})
}

func (a *actorNats) onMsg(msg *nats.Msg) {
	stdlib.Catch(func() {
		var nodeID def.NodeID
		var actorID def.ActorID
		if _, err := fmt.Sscanf(msg.Subject, "%d.%d", &nodeID, &actorID); err != nil {
			logger.L.Warn("invalid GOD Msg", zap.Error(err))
			return
		}
		m := natsMsg2Message(msg)

		logger.L.Debug("receive NATS Msg", zap.String("subject", msg.Subject), zap.Any("message", m))

		stdlib.Post(actorID, m)
	})
}

func natsMsg2Message(msg *nats.Msg) def.Message {
	return string(msg.Data)
}

func init() {
	stdlib.RegisterActorCreator(def.ATNats, func() stdlib.Actor {
		return &actorNats{}
	})
}
