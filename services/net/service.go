package net

import (
	"net"

	"github.com/enjoypi/god"
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	Config
	*zap.Logger
	*god.Node
	pb.UnimplementedSessionServer

	childState sc.State
	godSvc     *god.Service
}

func NewService(cfg Config, logger *zap.Logger, node *god.Node, initialState sc.State, childState sc.State) *god.Service {
	svc := &Service{
		Config:     cfg,
		Logger:     logger,
		Node:       node,
		childState: childState,
	}
	svc.godSvc = god.NewService(logger, svc, initialState, svc)
	actors.Go(func(actors.ExitChan, interface{}) (interface{}, error) {
		err := svc.Serve()
		return nil, err
	}, nil, nil)
	return svc.godSvc
}

func (svc *Service) Serve() error {
	lis, err := net.Listen("tcp", svc.Config.Net.ListenAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterSessionServer(s, svc)

	svc.Info(lis.Addr().String())
	return s.Serve(lis)
}

func (svc *Service) Flow(stream pb.Session_FlowServer) error {
	actor, err := svc.godSvc.NewAgent(0, svc.childState,
		&Session{Logger: svc.Logger, Node: svc.Node, Session_FlowServer: stream})
	if err != nil {
		return err
	}

	actor.Run(actors.DefaultActors.ExitChan)
	svc.godSvc.RemoveAgent(actor.ID)
	return nil
}
