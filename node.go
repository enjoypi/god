package god

import (
	"github.com/enjoypi/god/service"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type Node struct {
	*Config
	*zap.Logger

	services map[service.ServiceType]*service.NetService
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID <= 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{Config: cfg,
		Logger:   logger,
		services: make(map[uint16]*service.NetService),
	}, nil
}

func (n *Node) NewService(srvType service.ServiceType, state sc.State) (*service.NetService, error) {
	_, ok := n.services[srvType]
	if ok {
		return nil, ErrDuplicateService
	}
	srv := service.NewNetService(srvType, state, nil, n.Logger)
	n.services[srvType] = srv
	return srv, nil
}
