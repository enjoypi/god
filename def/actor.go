package def

import "github.com/enjoypi/gostatechart"

type ActorType uint32
type ActorID = uint32
type Message = gostatechart.Event
type MessageQueue chan Message
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

func (at ActorType) String() string {
	if at >= ATUser {
		return "User"
	}
	return actorTypeName[at-1]
}

const (
	ATNats       ActorType = 1
	ATEtcd       ActorType = 2
	ATPrometheus ActorType = 3

	ATUser ActorType = 1000
)

var actorTypeName = [...]string{"NATS", "etcd", "Prometheus"}

// KernelActors
//It use actor type as actor ID because of only one actor each type
var KernelActors = []struct {
	Type ActorType
	ID   ActorID
}{
	{ATEtcd, ActorID(ATEtcd)},
	{ATNats, ActorID(ATNats)},
	{ATPrometheus, ActorID(ATPrometheus)},
}
