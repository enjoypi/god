package mesh

import (
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	etcdclient "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Params struct {
	Config
	*zap.Logger
	*god.Node
}

func NewService(cfg Config, logger *zap.Logger, node *god.Node) *god.Service {
	return god.NewService(
		logger,
		(*main)(nil),
		&Params{Config: normalizeConfig(cfg), Logger: logger, Node: node},
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
	m.RegisterReaction((*pb.ServiceInfo)(nil), m.onServiceInfo)

	return m.connectEtcd()
}

func (m *main) End(event sc.Event) sc.Event {
	//m.Node.Terminate()
	return m.Client.Close()
}

func (m *main) GetTransitions() sc.Transitions {
	return nil
}

func (m *main) connectEtcd() error {
	var leaseID etcdclient.LeaseID
	var err error

	m.Client, leaseID, err = dialEtcd(m.Config, m.Logger)
	for i := 0; i < m.Mesh.RetryTimes && err != nil; {
		m.Client, leaseID, err = dialEtcd(m.Config, m.Logger)
	}

	if err == nil {
		m.Node.Go(m.keepAlive, leaseID, m.onDropped)
	}
	return err
}

func (m *main) getDeps() {

}

func (m *main) onDropped(i interface{}, err error) {
	m.Client = nil

	if err == nil {
		return
	}
	m.Logger.Warn("mesh keep alive failed", zap.Error(err))

	if err := m.connectEtcd(); err != nil {
		m.Logger.Error("connect etcd failed", zap.Error(err))
	}
}

func (m *main) onEvents(events []*etcdclient.Event) {
	for _, e := range events {
		if e.IsCreate() {
			m.Logger.Debug("created",
				zap.ByteString("key", e.Kv.Key),
				zap.ByteString("value", e.Kv.Value),
			)
		}
		if e.IsModify() {
			m.Logger.Debug("modified",
				zap.ByteString("key", e.Kv.Key),
				zap.ByteString("value", e.Kv.Value),
			)
		}
	}
}

func (m *main) keepAlive(exitChan god.ExitChan, i interface{}) (interface{}, error) {
	leaseID := i.(etcdclient.LeaseID)

	keepChan, err := m.Client.KeepAlive(context.Background(), leaseID)
	if err != nil {
		return nil, err
	}

	watchChan := m.Client.Watch(context.Background(), m.Mesh.Path+"/*")

	for {
		select {
		case resp := <-watchChan:
			if resp.Err() != nil {
				return nil, resp.Err()
			}
			m.onEvents(resp.Events)
		case resp := <-keepChan:
			if resp == nil {
				return nil, ErrEtcdDropped
			}
		case <-exitChan:
			return nil, nil
		}
	}
}
