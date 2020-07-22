package god

import (
	"math/rand"

	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
)

// service = supervisor
// 1 supervisor : 0-N actor

type ServiceType = uint16
type Actor = sc.StateMachine
type ActorID = uint16

type Service struct {
	*Actor
	*zap.Logger
	ServiceType

	children map[ActorID]*Actor
	context  interface{}
}

func NewService(initialState sc.State, context interface{}) *Service {
	return &Service{
		Actor:    sc.NewStateMachine(initialState, context),
		children: make(map[ActorID]*Actor),
	}
}

func (srv *Service) NewActor(initialState sc.State, id ActorID) (*Actor, error) {
	actorID := id
	if actorID == 0 {
		actorID = ActorID(rand.Uint32())
	}

	_, ok := srv.children[actorID]
	if ok {
		return nil, ErrDuplicateActor
	}

	actor := sc.NewStateMachine(initialState, srv.context)
	if err := actor.Initiate(nil); err != nil {
		return nil, ErrFailedInitialization
	}
	srv.children[actorID] = actor
	go actor.Run()
	return actor, nil
}
