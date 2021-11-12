package actors

import (
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

type factory struct {
	// 由于creator是初始化的时候顺序存入，而程序运行过程中只读，因此不需要sync.Map
	creators map[types.ActorType]ActorCreator
}

var defaultFactory = newFactory()

func newFactory() *factory {
	return &factory{creators: make(map[types.ActorType]ActorCreator)}
}

func (am *factory) RegisterCreator(actorType types.ActorType, creator ActorCreator) bool {
	_, ok := am.creators[actorType]
	if ok {
		logger.L.Error("the actor creator is already registered", zap.String("actorType", actorType))
		return false
	}
	am.creators[actorType] = creator
	return true
}

func (am *factory) NewActor(actorType types.ActorType) Actor {
	creator, ok := am.creators[actorType]
	if ok {
		actor := creator()
		actor.setType(actorType)
		return actor
	}
	logger.L.Error("no actor creator for this type", zap.String("actorType", actorType))
	return nil
}

func RegisterActorCreator(actorType types.ActorType, creator ActorCreator) bool {
	return defaultFactory.RegisterCreator(actorType, creator)
}

func NewActor(actorType types.ActorType) Actor {
	return defaultFactory.NewActor(actorType)
}
