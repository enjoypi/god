package events

import "github.com/enjoypi/god/types"

type EvBusConnected struct {
}

type EvBusDisconnected struct {
}

type EvBusReconnected struct {
}

type Subscription struct {
	types.ActorID
	Subject string
}
