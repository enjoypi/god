package god

import "golang.org/x/net/context"

type server struct{}

func newServer() *server {
	return &server{}
}

func (s *server) Auth(context.Context, *AuthReq) (*AuthAck, error) {
	return &AuthAck{Code: ErrorCode_OK}, nil
}
