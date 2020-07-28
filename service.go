package god

import (
	"math/rand"
	"sync"

	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

// service = supervisor
// 1 supervisor : 0-N actor

type ActorID = uint64

type Actor struct {
	ID ActorID
	*sc.StateMachine
}

type Service struct {
	*Actor
	*zap.Logger
	pb.ServiceType

	children    sync.Map
	realService interface{}
}

func NewService(logger *zap.Logger, realService interface{}, initialState sc.State, stateCtx interface{}) *Service {
	svc := &Service{
		Actor: &Actor{
			ID:           0,
			StateMachine: sc.NewStateMachine(initialState, stateCtx),
		},
		Logger:      logger,
		realService: realService,
	}

	if err := svc.Initiate(nil); err != nil {
		return nil
	}
	return svc
}

func (svc *Service) NewAgent(id ActorID, initialState sc.State, stateCtx interface{}) (*Actor, error) {
	if id == 0 {
		id = rand.Uint64()
	}

	_, ok := svc.children.Load(id)
	if ok {
		return nil, ErrDuplicateActor
	}

	machine := sc.NewStateMachine(initialState, stateCtx)
	if err := machine.Initiate(nil); err != nil {
		return nil, ErrFailedInitialization
	}

	actor := &Actor{ID: id, StateMachine: machine}
	svc.children.Store(id, actor)

	svc.Logger.Info("new agent", zap.Uint64("id", id))
	return actor, nil
}

func (svc *Service) RemoveAgent(id ActorID) {
	svc.children.Delete(id)
	svc.Logger.Info("remove agent", zap.Uint64("id", id))
}
