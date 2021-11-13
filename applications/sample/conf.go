package sample

import "github.com/enjoypi/god/def"

type conf struct {
	Actors []struct {
		Type string
		def.ActorID
	}
}
