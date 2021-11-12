package actors

import (
	"fmt"
	"reflect"

	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/types"
	sc "github.com/enjoypi/gostatechart"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Handler interface {
	Handle(message types.Message) error // will be called in actor's goroutine
}

type Receiver interface {
	Post(message types.Message) // post message to actor's message queue, must thread safe
}

type handle func(message types.Message) types.Message // will be called in actor's goroutine
type DefaultImplement interface {
	ID() types.ActorID
	Type() types.ActorType

	MessageQueue() types.MessageQueue
	setID(id types.ActorID)
	setType(actorType types.ActorType)

	Handler
	Receiver
}

type Actor interface {
	Initialize(v *viper.Viper) error // will be called by supervisor
	Terminate()                      // will be called in actor's goroutine

	DefaultImplement
}
type ActorCreator func() Actor

type SimpleActor struct {
	actorType types.ActorType
	id        types.ActorID
	mq        types.MessageQueue
	sc.SimpleState
}

// must be called by outer Initialize and ignore error
func (a *SimpleActor) Initialize() error {
	a.mq = make(types.MessageQueue, 1)
	return fmt.Errorf("no Initialize implment")
}

func (a *SimpleActor) Handle(message types.Message) error {
	//if !ok {
	//	logger.L.Warn("invalid reactor in actor",
	//		zap.String("type", a.actorType),
	//		zap.Any("ID", a.id),
	//		zap.Any("message", message),
	//	)
	//	return nil
	//}
	ret := a.SimpleState.React(message)
	if ret != nil {
		a.Post(ret)
	}

	return nil
}

func (a *SimpleActor) ID() types.ActorID {
	return a.id
}

func (a *SimpleActor) Type() types.ActorType {
	return a.actorType
}

// no any check for performance
// Post will lock if the mq has not been initial
func (a *SimpleActor) Post(message types.Message) {
	if ce := logger.L.Check(zapcore.DebugLevel, "POST"); ce != nil {
		ce.Write(
			zap.String("type", a.actorType.String()),
			zap.Uint32("actor", a.id),
			zap.Any("message", reflect.TypeOf(message)))
	}
	a.mq <- message
}

func (a *SimpleActor) Terminate() {

}

func (a *SimpleActor) MessageQueue() types.MessageQueue {
	return a.mq
}

func (a *SimpleActor) setID(id types.ActorID) {
	a.id = id
}

func (a *SimpleActor) setType(actorType types.ActorType) {
	a.actorType = actorType
}
