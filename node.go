package god

import (
	"context"
	"net"
	"reflect"
	"time"

	"github.com/enjoypi/god/pb"
	"github.com/enjoypi/god/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedSessionServer
	*Config
	*zap.Logger
}

func NewNode(cfg *Config, logger *zap.Logger) (*Node, error) {
	if cfg.Node.ID <= 0 {
		return nil, ErrInvalidNodeID
	}
	return &Node{Config: cfg, Logger: logger}, nil
}

func (n *Node) Serve(srv *service.Service, listenAddress string) error {
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterSessionServer(s, n)

	n.Info(lis.Addr().String())
	return s.Serve(lis)
}

func (n *Node) Auth(ctx context.Context, in *pb.AuthReq) (*pb.AuthAck, error) {
	return &pb.AuthAck{Code: pb.ErrorCode_OK}, nil
}

func (n *Node) Ping(ctx context.Context, in *pb.Heartbeat) (*pb.Heartbeat, error) {
	in.ToTimestamp = time.Now().UnixNano()
	return in, nil
}

func (n *Node) Flow(stream pb.Session_FlowServer) error {
	var err error
	var header pb.Header

	for {
		if err = stream.RecvMsg(&header); err != nil {
			break
		}
		n.Info(header.String(), zap.String("type", reflect.TypeOf(header).String()))

		var req pb.AuthReq
		if err = stream.RecvMsg(&req); err != nil {
			break
		}
		n.Info(req.String(), zap.String("type", reflect.TypeOf(req).String()))

		header.Serial++
		if err = stream.SendMsg(&header); err != nil {
			break
		}

		if err := stream.SendMsg(&req); err != nil {
			break
		}
	}
	return err
}
