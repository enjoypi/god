package def

import "github.com/enjoypi/gostatechart"

type ActorType string
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

func GetActorType(name string) ActorType {
	if t, ok := nameTypeMap[name]; ok {
		return t
	}
	return "User"
}

const (
	ATSample         = "sample"
	ATNats           = "NATS"
	ATEtcd           = "etcd"
	ATPrometheus     = "Prometheus"
	ATQuic           = "QUIC"
	ATSocketListener = "SocketListener"

	ATUser = 1000
)

var actorTypeName = [...]string{"sample", "etcd", "NATS", "Prometheus", "QUIC", "SocketListener"}

var nameTypeMap = make(map[string]ActorType)

func init() {
	for i, name := range actorTypeName {
		nameTypeMap[name] = ActorType(i)
	}
}

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
