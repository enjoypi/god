package sample

import "github.com/enjoypi/god/def"

type option struct {
	Actors []struct {
		Type string
		def.ActorID
	}
}
