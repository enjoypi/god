package mesh

import (
	"github.com/enjoypi/god"
	sc "github.com/enjoypi/gostatechart"
	etcdclient "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

type Params struct {
	Config
	*zap.Logger
}

func NewService(cfg Config, logger *zap.Logger) *god.Service {
	return god.NewService(
		logger,
		(*main)(nil),
		&Params{cfg, logger},
	)
}

type main struct {
	// implement sc.State
	sc.SimpleState

	*etcdclient.Client
	*Params
}

func (m *main) Begin(ctx interface{}, event sc.Event) sc.Event {
	m.Params = ctx.(*Params)
	if c, err := dialEtcd(m.Config, m.Logger); err != nil {
		return err
	} else {
		m.Client = c
	}
	return nil
}

func (s *main) GetTransitions() sc.Transitions {
	return nil
}
