package mesh

import (
	"github.com/enjoypi/god"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

type Context struct {
	Config
	*zap.Logger
}

func NewService(cfg Config, logger *zap.Logger) *god.Service {
	return god.NewService(
		(*Root)(nil),
		&Context{cfg, logger},
	)
}

type Root struct {
	sc.SimpleState
	*Context
}

func (s Root) Begin(ctx interface{}, event sc.Event) sc.Event {
	s.Context = ctx.(*Context)
	return initEtcd(s.Config, s.Logger)
}

func (s Root) End(event sc.Event) sc.Event {
	panic("implement me")
}

func (s Root) GetTransitions() sc.Transitions {
	return nil
}
