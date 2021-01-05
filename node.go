package god

import (
	"github.com/enjoypi/god/core"
	"github.com/enjoypi/god/pb"
	mb "github.com/enjoypi/god/transports/t_nats"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type Node struct {
	ID uint32
	*zap.Logger

	services map[pb.ServiceType]*Service
	trans    *mb.Transport
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID < 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{
		ID:       cfg.Node.ID,
		Logger:   logger,
		services: make(map[pb.ServiceType]*Service),
	}, nil
}

func (n *Node) AddTransport(svcType pb.TransportType, trans *mb.Transport) error {
	n.trans = trans
	return nil
}

func (n *Node) AddService(svcType pb.ServiceType, svc *Service) error {
	if svc == nil {
		return ErrFailedInitialization
	}

	_, ok := n.services[svcType]
	if ok {
		return ErrDuplicateService
	}

	n.services[svcType] = svc
	return nil
}

func (n *Node) CastTo(svcType pb.ServiceType, msg interface{}) error {
	mesh, ok := n.services[svcType]
	if !ok {
		return ErrNoService
	}
	mesh.PostEvent(msg)
	return nil
}

func (n *Node) RealService(svcType pb.ServiceType) interface{} {
	svc, ok := n.services[svcType]
	if !ok {
		return nil
	}
	return svc.realService
}

func (n *Node) RegisterService(svcType pb.ServiceType) error {
	return n.CastTo(pb.ServiceType_Mesh,
		&pb.ServiceInfo{
			NodeID:      n.ID,
			ServiceType: svcType,
		})
}

func (n *Node) Serve() error {
	for svcType, svc := range n.services {
		if svcType != pb.ServiceType_Mesh {
			if err := n.RegisterService(svcType); err != nil {
				//n.Error("register service failed", zap.Error(err))
				continue
			}
		}
		core.Go(func(exitChan core.ExitChan, parameter interface{}) (interface{}, error) {
			svc := parameter.(*Service)
			svc.Run(exitChan)

			return nil, nil
		}, svc, nil)
	}

	core.Wait()
	return nil
}

func (n *Node) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return n.trans.Conn.Subscribe(subj, cb)
}

func (n *Node) Terminate() {
	core.Close()
}
