package god

import (
	"sync"

	"github.com/enjoypi/god/pb"

	"go.uber.org/zap"
)

type ExitChan chan int
type GoRun func(ExitChan, interface{}) (interface{}, error)
type OnGoReturn func(interface{}, error)

type Node struct {
	ID uint32
	*zap.Logger

	ExitChan

	services map[pb.ServiceType]*Service
	wg       sync.WaitGroup
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID <= 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{
		ID:       cfg.Node.ID,
		Logger:   logger,
		ExitChan: make(ExitChan),
		services: make(map[pb.ServiceType]*Service),
	}, nil
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

func (n *Node) Cast2Service(svcType pb.ServiceType, msg interface{}) error {
	mesh, ok := n.services[svcType]
	if !ok {
		return ErrNoService
	}
	mesh.PostEvent(msg)
	return nil
}

func (n *Node) Go(run GoRun, parameter interface{}, onRet OnGoReturn) {
	n.wg.Add(1)
	go func() {
		defer n.wg.Done()
		ret, err := run(n.ExitChan, parameter)
		if onRet != nil {
			onRet(ret, err)
		}
	}()
}

func (n *Node) RegisterService(svcType pb.ServiceType) error {
	return n.Cast2Service(pb.ServiceType_Mesh,
		&pb.ServiceInfo{
			NodeID:      n.ID,
			ServiceType: svcType,
		})
}

func (n *Node) Serve() error {
	for svcType, svc := range n.services {
		if svcType != pb.ServiceType_Mesh {

			if err := n.RegisterService(svcType); err != nil {
				n.Error("register service failed", zap.Error(err))
				continue
				//return nil, err
			}
		}
		n.Go(func(exitChan ExitChan, parameter interface{}) (interface{}, error) {
			svc := parameter.(*Service)
			svc.Run(exitChan)

			return nil, nil
		}, svc, nil)
	}

	n.wg.Wait()
	return nil
}

func (n *Node) Terminate() {
	close(n.ExitChan)
}
