package core

import (
	"sync"

	"github.com/enjoypi/god"

	"go.uber.org/zap"
)

type Factory struct {
	sync.Map
}

var (
	DefaultFactory Factory
)

func (factory *Factory) RegisterCreator(actorType ActorType, creator ActorCreator) bool {
	_, ok := factory.Load(actorType)
	if ok {
		god.Logger.Error("the actor creator is already registered", zap.Int64("actorType", actorType))
		return false
	}
	factory.Store(actorType, creator)
	return true
}

func (factory *Factory) NewActor(actorType ActorType) Actor {
	creator, ok := factory.Load(actorType)
	if ok {
		return creator.(ActorCreator)()
	}
	god.Logger.Error("no actor creator for this type", zap.Int64("actorType", actorType))
	return nil
}

func RegisterActorCreator(actorType ActorType, creator ActorCreator) bool {
	return DefaultFactory.RegisterCreator(actorType, creator)
}

func NewActor(actorType ActorType) Actor {
	return DefaultFactory.NewActor(actorType)
}
