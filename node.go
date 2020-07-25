package god

import (
	"sync"

	"go.uber.org/zap"
)

type ExitChan chan int
type GoRun func(ExitChan, interface{}) (interface{}, error)
type OnGoReturn func(interface{}, error)

type Node struct {
	*zap.Logger

	ExitChan

	services map[ServiceType]*Service
	wg       sync.WaitGroup
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID <= 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{
		Logger:   logger,
		ExitChan: make(ExitChan),
		services: make(map[ServiceType]*Service),
	}, nil
}

func (n *Node) AddService(svcType ServiceType, svc *Service) error {
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

func (n *Node) Serve() error {
	for _, svc := range n.services {
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
