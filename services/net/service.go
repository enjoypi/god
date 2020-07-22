package net

import (
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"reflect"
)

type NetService struct {
	Config
	*zap.Logger
	*god.Service
	pb.UnimplementedSessionServer

	childState sc.State
}

func NewService(cfg Config, logger *zap.Logger, initialState sc.State, childState sc.State) *god.Service {
	svc := &NetService{
		Config:     cfg,
		Logger:     logger,
		childState: childState,
	}
	svc.Service = god.NewService(initialState, svc)
	return svc.Service
}

func (n *NetService) Serve() error {
	lis, err := net.Listen("tcp", n.Config.Net.ListenAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterSessionServer(s, n)

	n.Info(lis.Addr().String())
	go n.Run()
	return s.Serve(lis)
}

func (n *NetService) Flow(stream pb.Session_FlowServer) error {
	var err error

	actor, err := n.Service.NewActor(n.childState, 0)
	if err != nil {
		return err
	}

	for {
		var header pb.Header
		if err = stream.RecvMsg(&header); err != nil {
			break
		}

		//typ, ok := p.id2type[header.MessageType]
		//if !ok {
		//	return 0, nil, ErrMessageNotRegistered
		//}
		//// 根据类型创建一个对应的实例
		//msg0 := reflect.New(typ.Elem()).Interface().(proto.Message)

		var req pb.Heartbeat
		if err = stream.RecvMsg(&req); err != nil {
			break
		}
		n.Info(req.String(), zap.String("type", reflect.TypeOf(req).String()))
		n.PostEvent(&header)
		actor.PostEvent(&header)

		//header.Serial++
		//if err = stream.SendMsg(&header); err != nil {
		//	break
		//}
		//
		//if err := stream.SendMsg(&req); err != nil {
		//	break
		//}
	}
	return err
}
