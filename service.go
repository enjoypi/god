package god

import (
	"math/rand"
	"sync"

	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

// service = supervisor
// 1 supervisor : 0-N actor

type ServiceType = uint16
type ActorID = uint64

type Actor struct {
	ID ActorID
	*sc.StateMachine
}

type Service struct {
	*Actor
	*zap.Logger
	ServiceType

	children sync.Map
	context  interface{}
}

func NewService(logger *zap.Logger, initialState sc.State, context interface{}) *Service {
	svc := &Service{
		Actor: &Actor{
			ID:           0,
			StateMachine: sc.NewStateMachine(initialState, context),
		},
		Logger: logger,
	}

	if err := svc.Initiate(nil); err != nil {
		return nil
	}
	return svc
}

func (svc *Service) NewActor(id ActorID, initialState sc.State, context interface{}) (*Actor, error) {
	if id == 0 {
		id = rand.Uint64()
	}

	_, ok := svc.children.Load(id)
	if ok {
		return nil, ErrDuplicateActor
	}

	machine := sc.NewStateMachine(initialState, context)
	if err := machine.Initiate(nil); err != nil {
		return nil, ErrFailedInitialization
	}

	actor := &Actor{ID: id, StateMachine: machine}
	svc.children.Store(id, actor)

	svc.Logger.Info("new actor", zap.Uint64("id", id))
	return actor, nil
}

func (svc *Service) DeleteActor(id ActorID) {
	svc.children.Delete(id)
	svc.Logger.Info("delete actor", zap.Uint64("id", id))
}
