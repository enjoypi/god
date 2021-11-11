package actors

import (
	"sync"

	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

type ExitChan chan int
type GoRun func(ExitChan) (types.Message, error)
type OnReturn func(types.Message, error)

type ActorManager struct {
	actors sync.Map
	wg     sync.WaitGroup
	ExitChan

	// 由于creator是初始化的时候顺序存入，而程序运行过程中只读，因此不需要sync.Map
	creators map[types.ActorType]ActorCreator
}

var (
	defaultActorManager *ActorManager
)

func init() {
	defaultActorManager = NewActorManager()
}

func NewActorManager() *ActorManager {
	return &ActorManager{
		ExitChan: make(ExitChan, 1),
		creators: make(map[types.ActorType]ActorCreator),
	}
}

func (am *ActorManager) Close() {
	close(am.ExitChan)
}

func (am *ActorManager) Get(id types.ActorID) Actor {
	actor, _ := am.actors.Load(id)
	return actor.(Actor)
}

func (am *ActorManager) Go(run GoRun, onRet OnReturn) {
	am.wg.Add(1)
	go func() {
		defer am.wg.Done()
		ret, err := run(am.ExitChan)
		if onRet != nil {
			onRet(ret, err)
		}
	}()
}

func (am *ActorManager) Post(sender types.ActorID, receiver types.ActorID, message types.Message) {
	actor := am.Get(receiver)
	if actor != nil {
		actor.Post(message)
		return
	}
}

func (am *ActorManager) Wait() {
	am.wg.Wait()
}

func (am *ActorManager) RegisterCreator(actorType types.ActorType, creator ActorCreator) bool {
	_, ok := am.creators[actorType]
	if ok {
		logger.L.Error("the actor creator is already registered", zap.String("actorType", actorType))
		return false
	}
	am.creators[actorType] = creator
	return true
}

func (am *ActorManager) NewActor(actorType types.ActorType) Actor {
	creator, ok := am.creators[actorType]
	if ok {
		actor := creator()
		actor.setType(actorType)
		return actor
	}
	logger.L.Error("no actor creator for this type", zap.String("actorType", actorType))
	return nil
}

func Close() {
	defaultActorManager.Close()
}

func Go(run GoRun, onRet OnReturn) {
	defaultActorManager.Go(run, onRet)
}

func Post(sender types.ActorID, receiver types.ActorID, message types.Message) {
	defaultActorManager.Post(sender, receiver, message)
}

func Wait() {
	defaultActorManager.Wait()
}

func RegisterActorCreator(actorType types.ActorType, creator ActorCreator) bool {
	return defaultActorManager.RegisterCreator(actorType, creator)
}

func NewActor(actorType types.ActorType) Actor {
	return defaultActorManager.NewActor(actorType)
}
