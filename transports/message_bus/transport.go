package message_bus

import (
	"fmt"

	sc "github.com/enjoypi/gostatechart"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Transport struct {
	Config
	*zap.Logger

	childState sc.State
	*nats.Conn

	*test
}

func NewTransport(cfg Config, logger *zap.Logger, nodeID uint32) *Transport {
	trans := &Transport{
		Config: cfg,
		Logger: logger,
	}

	if err := trans.Initialize(fmt.Sprintf("%d", nodeID)); err != nil {
		logger.Error("initialize NATS failed", zap.Error(err))
		return nil
	}

	return trans
}

func mergeOpts(config Config, options *nats.Options) {
	cfg := config.Nats
	options.Url = cfg.Url
}

func (trans *Transport) Initialize(name string) error {
	opts := nats.GetDefaultOptions()
	mergeOpts(trans.Config, &opts)

	opts.Name = name
	opts.DisconnectedErrCB = trans.onDisconnected
	opts.ReconnectedCB = trans.onReconnected

	nc, err := opts.Connect()
	if err != nil {
		return err
	}
	trans.Conn = nc
	trans.Info("NATS connected", zap.String("url", nc.ConnectedUrl()))

	trans.test = newTest(trans.Logger)
	return nil
}

func (trans *Transport) onDisconnected(nc *nats.Conn, err error) {
	trans.Logger.Warn("NATS disconnected", zap.Error(err))
}

func (trans *Transport) onReconnected(nc *nats.Conn) {
	trans.Logger.Info("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
	if trans.Conn != nc {
		trans.Conn = nc
	}
}

//func (trans *Transport) Run() error {
//	sub, err := trans.Conn.SubscribeSync(">")
//	if err != nil {
//		return err
//	}
//
//	for {
//		msg, err := sub.NextMsg(time.Hour)
//		if err != nil {
//			break
//		}
//		trans.PostEvent(msg)
//	}
//	return sub.Unsubscribe()
//}
