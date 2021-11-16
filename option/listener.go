package option

import "github.com/enjoypi/god/def"

type Listen struct {
	def.ActorType
	def.ActorID

	Network string
	Address string
	Handler def.ActorType
}
