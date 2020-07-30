package mesh

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/enjoypi/god"
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/pb"
	mb "github.com/enjoypi/god/transports/message_bus"
	sc "github.com/enjoypi/gostatechart"
	"github.com/nats-io/nats.go"
	etcdclient "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type Service struct {
	Config
	*zap.Logger

	trans *mb.Transport
}

func NewService(cfg Config, logger *zap.Logger, trans *mb.Transport) *god.Service {
	svc := &Service{Config: normalizeConfig(cfg),
		Logger: logger,
		trans:  trans,
	}

	return god.NewService(
		logger,
		svc,
		(*main)(nil),
		svc,
	)
}

func (s *Service) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return s.trans.Conn.Subscribe(subj, cb)
}

type main struct {
	// implement sc.State
	sc.SimpleState

	*etcdclient.Client
	*Service
	buf bytes.Buffer
}

type Marshal func() (dAtA []byte, err error)

func (m *main) writeProto(marshal Marshal, size int, w io.Writer) error {
	b, err := marshal()
	if err != nil {
		return err
	}
	m.Logger.Debug("size compare", zap.Int("size", size), zap.Int("trueSize", len(b)))

	l := uint16(len(b))
	if err := binary.Write(w, binary.LittleEndian, l); err != nil {
		return err
	}

	if err := binary.Write(w, binary.LittleEndian, b); err != nil {
		return err
	}
	return nil
}

func (m *main) Begin(ctx interface{}, event sc.Event) sc.Event {
	m.Service = ctx.(*Service)
	//m.RegisterReaction((*pb.ServiceInfo)(nil), m.onServiceInfo)
	m.RegisterReaction((*pb.Echo)(nil), func(e sc.Event) sc.Event {
		buf := &m.buf
		buf.Reset()

		var header pb.Header

		header.Serial++
		header.MessageType = "pb.Echo"
		if err := m.writeProto(header.Marshal, header.Size(), buf); err != nil {
			return err
		}

		req := e.(*pb.Echo)
		if err := m.writeProto(req.Marshal, req.Size(), buf); err != nil {
			return err
		}

		m.Logger.Info("buf len", zap.Int("length", buf.Len()))
		return m.trans.Publish("echo", buf.Bytes())
	})

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
		actors.Go(m.keepAlive, leaseID, m.onDropped)
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

func (m *main) keepAlive(exitChan actors.ExitChan, i interface{}) (interface{}, error) {
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
