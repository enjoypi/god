package stdlib

import (
	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/logger"
	"go.uber.org/zap"
)

type factory struct {
	// 由于creator是初始化的时候顺序存入，而程序运行过程中只读，因此不需要sync.Map
	creators map[def.ActorType]ActorCreator
}

var defaultFactory = newFactory()

func newFactory() *factory {
	return &factory{creators: make(map[def.ActorType]ActorCreator)}
}

func (am *factory) RegisterCreator(actorType def.ActorType, creator ActorCreator) bool {
	_, ok := am.creators[actorType]
	if ok {
		logger.L.Error("the actor creator is already registered", zap.String("actorType", actorType.String()))
		return false
	}
	am.creators[actorType] = creator
	return true
}

func (am *factory) NewActor(actorType def.ActorType, id def.ActorID) Actor {
	creator, ok := am.creators[actorType]
	if ok {
		actor := creator()
		actor.setType(actorType)
		actor.setID(id)
		return actor
	}
	logger.L.Error("no actor creator for this type", zap.String("actorType", actorType.String()))
	return nil
}

func RegisterActorCreator(actorType def.ActorType, creator ActorCreator) bool {
	return defaultFactory.RegisterCreator(actorType, creator)
}
