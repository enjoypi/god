package types

import "github.com/enjoypi/gostatechart"

type ActorType uint32
type ActorID = uint32
type Message = gostatechart.Event
type MessageQueue chan Message
type NodeID = uint16
type FullID = uint64

func DecodeID(id FullID) (NodeID, ActorID) {
	return NodeID(id >> 16), ActorID(id)
}

func EncodeID(id NodeID, actorID ActorID) FullID {
	return (FullID(id) << 16) | FullID(actorID)
}

func (at ActorType) String() string {
	return actorTypeName[at-1]
}

const (
	ATEtcd       = 1
	ATNats       = 2
	ATPrometheus = 3

	ATUser = 1000
)

var actorTypeName = [...]string{"etcd", "NATS", "Prometheus"}

// KernelActors
//It use actor type as actor ID because of only one actor each type
var KernelActors = []struct {
	Type ActorType
	ID   ActorID
}{
	{ATEtcd, ATEtcd},
	{ATNats, ATNats},
	{ATPrometheus, ATPrometheus},
}
