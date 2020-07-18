package god

import (
	"context"
	"net"

	"github.com/enjoypi/god/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNodeServer
	*zap.Logger
}

func Start(cfg *Config, logger *zap.Logger) error {
	lis, err := net.Listen("tcp", cfg.ListenAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server{Logger: logger})

	logger.Info(lis.Addr().String())
	return s.Serve(lis)
}

func (s *server) Auth(ctx context.Context, in *pb.AuthReq) (*pb.AuthAck, error) {
	s.Logger.Debug("Received:", zap.String("cookie", in.GetCookie()))
	return &pb.AuthAck{Code: pb.ErrorCode_OK}, nil
}
