package types

import "github.com/enjoypi/gostatechart"

type ActorType = string
type ActorID = uint16
type Message = gostatechart.Event
type MessageQueue chan Message
type NodeID = uint16
type FullID = uint32

type EvStart struct {
}

type EvStopped struct {
}

func DecodeID(id FullID) (NodeID, ActorID) {
	return NodeID(id >> 16), ActorID(id)
}

func EncodeID(id NodeID, actorID ActorID) FullID {
	return (FullID(id) << 16) | FullID(actorID)
}
