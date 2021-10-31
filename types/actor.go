package types

import "github.com/enjoypi/gostatechart"

type ActorType = string
type ActorID int64
type Message = gostatechart.Event
type MessageQueue chan Message

const (
	EvNone = iota
	EvStart
	EvStopped
	EvPanic
)
