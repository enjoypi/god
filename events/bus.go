package events

import "github.com/enjoypi/god/def"

type EvBusConnected struct {
}

type EvBusDisconnected struct {
}

type EvBusReconnected struct {
}

type Subscription struct {
	def.ActorID
	Subject string
}
