package service

import (
	"github.com/enjoypi/god/pb"
	sc "github.com/enjoypi/gostatechart"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"reflect"
)

type NetService struct {
	*Service
	pb.UnimplementedSessionServer
}

func NewNetService(srvType ServiceType, initialState sc.State, initialEvent sc.Event, logger *zap.Logger) *NetService {
	return &NetService{
		Service: NewService(srvType, initialState, initialEvent, logger),
	}
}

func (n *NetService) Serve(listenAddress string) error {
	lis, err := net.Listen("tcp", listenAddress)
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

	//actor := n.newActor()
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
		//actor.PostEvent(actor)

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
