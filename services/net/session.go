package net

import (
	"github.com/enjoypi/god"
	"github.com/enjoypi/god/pb"
	"go.uber.org/zap"
)

type Session struct {
	*zap.Logger
	*god.Node
	pb.Session_FlowServer
}
