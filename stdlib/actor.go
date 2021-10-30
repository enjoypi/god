package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/types"
)

type Receiver interface {
	Post(message types.Message) // post message to actor's message queue, must thread safe
}

type DefaultImplement interface {
	ID() types.ActorID
	messageQueue() types.MessageQueue
	setID(id types.ActorID)

	Receiver
}

type Actor interface {
	Handle(message types.Message) types.Message // will be called in actor's goroutine
	Initialize() error                          // will be called by supervisor
	Terminate()                                 // will be called in actor's goroutine

	DefaultImplement
}
type ActorCreator func() Actor

type DefaultActor struct {
	id types.ActorID
	mq types.MessageQueue
}

// must be called by outer Initialize and ignore error
func (a *DefaultActor) Initialize() error {
	a.mq = make(types.MessageQueue, 1)
	return fmt.Errorf("no Initialize implment")
}

func (a *DefaultActor) ID() types.ActorID {
	return a.id
}

// no any check for performance
// Post will lock if the mq has not been initial
func (a *DefaultActor) Post(message types.Message) {
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
