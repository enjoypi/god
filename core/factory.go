package core

import "go.uber.org/zap"

type actorFactory map[ActorType]NewActor

var (
	defaultActorFactory = make(actorFactory)
)

func (factory actorFactory) RegisterActorCreator(actorType ActorType, creator NewActor) bool {
	_, ok := factory[actorType]
	if ok {
		logger.Error("the actor creator is already registered", zap.Int64("actorType", actorType))
		return false
	}
	factory[actorType] = creator
	return true
}

func (factory actorFactory) new(actorType ActorType) Actor {
	creator, ok := factory[actorType]
	if ok {
		return creator()
	}
	logger.Error("no actor creator for this type", zap.Int64("actorType", actorType))
	return nil
}
