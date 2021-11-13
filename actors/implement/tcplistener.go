package implement

import (
	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/events"
	"github.com/spf13/viper"
)

type actorTCPListener struct {
	actors.SimpleActor
}

func (a *actorTCPListener) Initialize(v *viper.Viper) error {
	_ = a.SimpleActor.Initialize()
	a.RegisterReaction((*events.EvStart)(nil), a.onStart)
	return nil
}

func (a *actorTCPListener) onStart(message def.Message) def.Message {
	return nil
}

func newTCPListener() actors.Actor {
	return &actorTCPListener{}
}

func init() {
	actors.RegisterActorCreator(def.ATTCPListener, newTCPListener)
}
