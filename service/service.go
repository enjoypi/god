package service

import (
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
	"math/rand"
)

// 1 service 1 supervisor 0-N actor

type ServiceType = uint16
type Actor = sc.StateMachine

type Service struct {
	*zap.Logger
	Type ServiceType
	*Actor
	children map[uint32]*Actor
}

func NewService(srvType ServiceType, initialState sc.State, initialEvent sc.Event, logger *zap.Logger) *Service {
	srv := &Service{
		Type:     srvType,
		children: make(map[uint32]*Actor),
		Logger:   logger,
	}

	srv.Actor = sc.NewStateMachine(initialState, srv)
	if err := srv.Initiate(initialEvent); err != nil {
		return nil
	}

	return srv
}

func (srv *Service) newActor() *Actor {
	actorID := rand.Uint32()
	actor := sc.NewStateMachine(nil, nil)
	if err := actor.Initiate(nil); err != nil {
		return nil
	}
	srv.children[actorID] = actor
	go actor.Run()
	return actor
}
