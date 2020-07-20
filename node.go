package god

import (
	"context"
	"net"
	"reflect"
	"time"

	"github.com/enjoypi/god/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedNodeServer
	*zap.Logger
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Start(cfg *Config, logger *zap.Logger) error {
	if cfg.Node.ID <= 0 {
		return ErrInvalidNodeID
	}
	lis, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &Node{Logger: logger})

	logger.Info(lis.Addr().String())
	return s.Serve(lis)
}

func (n *Node) Auth(ctx context.Context, in *pb.AuthReq) (*pb.AuthAck, error) {
	return &pb.AuthAck{Code: pb.ErrorCode_OK}, nil
}

func (n *Node) Ping(ctx context.Context, in *pb.Heartbeat) (*pb.Heartbeat, error) {
	in.ToTimestamp = time.Now().UnixNano()
	return in, nil
}

func (n *Node) Flow(stream pb.Node_FlowServer) error {
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
