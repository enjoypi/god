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
	return actorTypeName[at]
}

func GetActorType(name string) ActorType {
	if t, ok := nameTypeMap[name]; ok {
		return t
	}
	return ATUser
}

const (
	ATSample      ActorType = 0
	ATNats        ActorType = 1
	ATEtcd        ActorType = 2
	ATPrometheus  ActorType = 3
	ATQuic        ActorType = 4
	ATNetListener ActorType = 5

	ATUser ActorType = 1000
)

var actorTypeName = [...]string{"sample", "NATS", "etcd", "Prometheus", "QUIC", "NetListener"}

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
	{ATEtcd, ActorID(ATEtcd)},
	{ATNats, ActorID(ATNats)},
	{ATPrometheus, ActorID(ATPrometheus)},
}
