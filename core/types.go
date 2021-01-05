package core

import "github.com/enjoypi/gostatechart"

type ActorID = int64
type ActorType = int64
type Event = gostatechart.Event

const (
	EvStopped = iota + 1
	EvPanic
)
