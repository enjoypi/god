package net

import (
	"github.com/enjoypi/god/pb"
)

type Session struct {
	pb.Session_FlowServer
}

func NewSession(server pb.Session_FlowServer) *Session {
	return nil
}

