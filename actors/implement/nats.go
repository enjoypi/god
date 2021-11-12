package implement

import (
	"fmt"

	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/settings"
	"github.com/enjoypi/god/types"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const actorTypeNats = "nats"

var conn *nats.Conn

type actorNats struct {
	actors.SimpleActor
	nats.Options

	config settings.Node
	*viper.Viper
}

func (a *actorNats) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()

	type config struct {
		Nats nats.Options
		settings.Node
	}
	var cfg config

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}
	a.config = cfg.Node

	a.Options = nats.GetDefaultOptions()
	opts := &a.Options
	opts.Url = cfg.Nats.Url
	opts.DisconnectedErrCB = a.onDisconnected
	opts.ReconnectedCB = a.onReconnected

	logger.L.Info("initialize NATS",
		zap.String("options", fmt.Sprintf("%+v", opts)))

	a.RegisterReaction(error(nil), a.onError)
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
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

	subject := fmt.Sprintf("%d.*", a.config.ID)
	_, err = conn.Subscribe(subject, a.onMsg)
	if err != nil {
		return err
	}
	return nil
}

func (a *actorNats) onDisconnected(nc *nats.Conn, err error) {
	conn = nil
	logger.L.Warn("NATS disconnected", zap.Error(err), zap.String("url", nc.Opts.Url))
	a.Post(events.EvBusDisconnected{})
}

func (a *actorNats) onReconnected(nc *nats.Conn) {
	if conn != nc {
		conn = nc
	}
	logger.L.Info("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
	a.Post(events.EvBusReconnected{})
}

func (a *actorNats) onMsg(msg *nats.Msg) {
	var nodeID types.NodeID
	var actorID types.ActorID
	if _, err := fmt.Sscanf(msg.Subject, "%d.%d", &nodeID, &actorID); err != nil {
		logger.L.Warn("invalid GOD Msg", zap.Error(err))
		return
	}
	m := natsMsg2Message(msg)

	logger.L.Debug("receive NATS Msg", zap.String("subject", msg.Subject), zap.Any("message", m))

	actors.Post(actorID, m)
}

func natsMsg2Message(msg *nats.Msg) types.Message {
	return string(msg.Data)
}

func init() {
	actors.RegisterActorCreator(actorTypeNats, func() actors.Actor {
		return &actorNats{}
	})
}
