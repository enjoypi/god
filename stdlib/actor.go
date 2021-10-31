package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
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
	Register(message types.Message, h handle)
	Type() types.ActorType

	messageQueue() types.MessageQueue
	setID(id types.ActorID)
	setType(actorType types.ActorType)

	Handler
	Receiver
}

type Actor interface {
	Initialize() error // will be called by supervisor
	Terminate()        // will be called in actor's goroutine

	DefaultImplement
}
type ActorCreator func() Actor

type DefaultActor struct {
	actorType types.ActorType
	id        types.ActorID
	mq        types.MessageQueue
	reactors  map[types.Message]handle
}

// must be called by outer Initialize and ignore error
func (a *DefaultActor) Initialize() error {
	a.mq = make(types.MessageQueue, 1)
	a.reactors = make(map[types.Message]handle)
	return fmt.Errorf("no Initialize implment")
}

func (a *DefaultActor) Handle(message types.Message) error {
	h, ok := a.reactors[message]
	if !ok {
		L.Warn("invalid reactor in actor",
			zap.String("type", a.actorType),
			zap.Any("ID", a.id),
			zap.Any("message", message),
		)
		return nil
	}
	ret := h(message)
	if ret != nil {
		a.Post(ret)
	}

	return nil
}

func (a *DefaultActor) ID() types.ActorID {
	return a.id
}

func (a *DefaultActor) Register(message types.Message, h handle) {
	a.reactors[message] = h
}

func (a *DefaultActor) Type() types.ActorType {
	return a.actorType
}

// no any check for performance
// Post will lock if the mq has not been initial
func (a *DefaultActor) Post(message types.Message) {
	L.Debug("post message",
		zap.String("actor", fmt.Sprintf("%p", a)),
		zap.String("mq", fmt.Sprintf("%p", a.mq)),
		zap.Any("message", message))
	a.mq <- message
}

func (a *DefaultActor) Terminate() {

}

func (a *DefaultActor) messageQueue() types.MessageQueue {
	return a.mq
}

func (a *DefaultActor) setID(id types.ActorID) {
	a.id = id
}

func (a *DefaultActor) setType(actorType types.ActorType) {
	a.actorType = actorType
}
