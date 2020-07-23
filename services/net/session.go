package net

import (
	"github.com/enjoypi/god/pb"
	"go.uber.org/zap"
)

type Session struct {
	*zap.Logger
	pb.Session_FlowServer
}
