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

func (n *Node) Serve() error {
	var wg sync.WaitGroup
	for _, svc := range n.services {
		wg.Add(1)
		go func(svc *Service) {
			defer wg.Done()
			svc.Run()
		}(svc)
	}
	wg.Wait()

	return nil
}
