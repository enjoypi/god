package def

import (
	"context"

	"github.com/enjoypi/gostatechart"
)

type ActorType string
type ActorID = uint32
type Message = gostatechart.Event
type MessageQueue chan struct {
	context.Context
	Message
}
type NodeID = uint16
type FullID = uint64

type Reply struct {
	sender   ActorID
	receiver ActorID
	Message
}

func DecodeID(id FullID) (NodeID, ActorID) {
	return NodeID(id >> 32), ActorID(id)
}

func EncodeID(id NodeID, actorID ActorID) FullID {
	return (FullID(id) << 32) | FullID(actorID)
}

const (
	ATSample         = "sample"
	ATNats           = "NATS"
	ATEtcd           = "etcd"
	ATPrometheus     = "Prometheus"
	ATQuic           = "QUIC"
	ATSocketListener = "SocketListener"
	ATTcp            = "TCP"
	ATUdp            = "UDP"

	AIDUser = 1000
)

// KernelActors
//It use actor type as actor ID because of only one actor each type
var KernelActors = []struct {
	ActorType
	ActorID
}{
	{ATEtcd, 1},
	{ATNats, 2},
	{ATPrometheus, 3},
}
