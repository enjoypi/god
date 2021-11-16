package socket

import "github.com/enjoypi/god/def"

type Listener struct {
	def.ActorType
	def.ActorID

	Address string
	Handler def.ActorType
	Network string
}

type option struct {
	Listener []Listener
}
