package types

import "github.com/enjoypi/gostatechart"

type ActorType = string
type ActorID = uint64
type Message = gostatechart.Event
type MessageQueue chan Message

type EvStart struct {
}

type EvStopped struct {
}
