package net

import (
	"net"

	"github.com/enjoypi/god"
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	Config
	*zap.Logger
	pb.UnimplementedSessionServer

	childState sc.State
	godSvc     *god.Service
}

func NewService(cfg Config, logger *zap.Logger, initialState sc.State, childState sc.State) *god.Service {
	svc := &Service{
		Config:     cfg,
		Logger:     logger,
		childState: childState,
	}
	svc.godSvc = god.NewService(logger, initialState, svc)
	go func() {
		_ = svc.Serve()
	}()
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
	actor, err := svc.godSvc.NewActor(0, svc.childState, &Session{Logger: svc.Logger, Session_FlowServer: stream})
	if err != nil {
		return err
	}

	actor.Run()
	svc.godSvc.DeleteActor(actor.ID)
	return nil
}
