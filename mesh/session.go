package mesh

import (
	"github.com/enjoypi/god/pb"
	"go.uber.org/zap"
)

type Session struct {
	pb.Session_FlowServer
}

func NewSession(server pb.Session_FlowServer) *Session {

}

