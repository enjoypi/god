package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/def"
	sc "github.com/enjoypi/gostatechart"
	"github.com/spf13/viper"
)

type Handler interface {
	Handle(message def.Message) error // will be called in actor's goroutine
}

type Receiver interface {
	Post(message def.Message) // post message to actor's message queue, must thread safe
}

type handle func(message def.Message) def.Message // will be called in actor's goroutine
type DefaultImplement interface {
	ID() def.ActorID
	Type() def.ActorType

	MessageQueue() def.MessageQueue
	setID(id def.ActorID)
	setType(actorType def.ActorType)

	Handler
	Receiver
}

type Actor interface {
	Initialize(v *viper.Viper, sup *Supervisor) error // will be called by supervisor
	Terminate()                                       // will be called in actor's goroutine

	DefaultImplement
}
type ActorCreator func() Actor

type SimpleActor struct {
	actorType def.ActorType
	id        def.ActorID
	mq        def.MessageQueue
	sc.SimpleState
}

// must be called by outer Initialize and ignore error
func (a *SimpleActor) Initialize() error {
	a.mq = make(def.MessageQueue, 1)
	return fmt.Errorf("no Initialize implment")
}

func (a *SimpleActor) Handle(message def.Message) error {
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

func (a *SimpleActor) ID() def.ActorID {
	return a.id
}

func (a *SimpleActor) Type() def.ActorType {
	return a.actorType
}

// no any check for performance
// Post will lock if the mq has not been initial
func (a *SimpleActor) Post(message def.Message) {
	//if ce := logger.L.Check(zapcore.DebugLevel, "POST"); ce != nil {
	//	ce.Write(
	//		zap.String("type", a.actorType.String()),
	//		zap.Uint32("actor", a.id),
	//		zap.Any("message", sc.TypeOf(message)))
	//}
	a.mq <- message
}

func (a *SimpleActor) Terminate() {

}

func (a *SimpleActor) MessageQueue() def.MessageQueue {
	return a.mq
}

func (a *SimpleActor) setID(id def.ActorID) {
	a.id = id
}

func (a *SimpleActor) setType(actorType def.ActorType) {
	a.actorType = actorType
}
