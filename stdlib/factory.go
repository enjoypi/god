package stdlib

import (
	"sync"

	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

type Factory struct {
	sync.Map
}

var (
	DefaultFactory Factory
)

func (factory *Factory) RegisterCreator(actorType types.ActorType, creator ActorCreator) bool {
	_, ok := factory.Load(actorType)
	if ok {
		L.Error("the actor creator is already registered", zap.Int64("actorType", actorType))
		return false
	}
	factory.Store(actorType, creator)
	return true
}

func (factory *Factory) NewActor(actorType types.ActorType) Actor {
	creator, ok := factory.Load(actorType)
	if ok {
		return creator.(ActorCreator)()
	}
	L.Error("no actor creator for this type", zap.Int64("actorType", actorType))
	return nil
}

func RegisterActorCreator(actorType types.ActorType, creator ActorCreator) bool {
	return DefaultFactory.RegisterCreator(actorType, creator)
}

func NewActor(actorType types.ActorType) Actor {
	return DefaultFactory.NewActor(actorType)
}
