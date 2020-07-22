package god

import (
	"sync"

	"go.uber.org/zap"
)

type Node struct {
	*zap.Logger

	services map[ServiceType]*Service
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID <= 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{
		Logger:   logger,
		services: make(map[ServiceType]*Service),
	}, nil
}

func (n *Node) AddService(srvType ServiceType, srv *Service) error {
	_, ok := n.services[srvType]
	if ok {
		return ErrDuplicateService
	}

	n.services[srvType] = srv
	return nil
}

func (n *Node) Serve() error {
	var wg sync.WaitGroup
	for _, svc := range n.services {
		wg.Add(1)
		go func(svc *Service) {
			svc.Run()
			wg.Done()
		}(svc)
	}
	wg.Wait()

	return nil
}
