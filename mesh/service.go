package mesh

import (
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type Service struct {
	Config
	*zap.Logger
	sc.SimpleState
}

func NewService(cfg Config, logger *zap.Logger) *Service {
	return &Service{Config: cfg, Logger: logger}
}

func (s Service) Begin(context interface{}, event sc.Event) sc.Event {
	return initEtcd(s.Config, s.Logger)
}

func (s Service) End(event sc.Event) sc.Event {
	panic("implement me")
}

func (s Service) GetTransitions() sc.Transitions {
	return nil
}
