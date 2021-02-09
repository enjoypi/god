package core

import "fmt"

type Receiver interface {
	Post(message Message) // post message to actor's message queue, must thread safe
}

type DefaultImplement interface {
	ID() ActorID
	messageQueue() MessageQueue
	setID(id ActorID)

	Receiver
}

type Actor interface {
	Handle(message Message) Message // will be called in actor's goroutine
	Initialize() error              // will be called by supervisor
	Terminate()                     // will be called in actor's goroutine

	DefaultImplement
}
type ActorCreator func() Actor

type DefaultActor struct {
	id ActorID
	mq MessageQueue
}

// must be called by outer Initialize and ignore error
func (a *DefaultActor) Initialize() error {
	a.mq = make(MessageQueue, 1)
	return fmt.Errorf("no Initialize implment")
}

func (a *DefaultActor) ID() ActorID {
	return a.id
}

// no any check for performance
// Post will lock if the mq has not been initial
func (a *DefaultActor) Post(message Message) {
	a.mq <- message
}

func (a *DefaultActor) Terminate() {

}

func (a *DefaultActor) messageQueue() MessageQueue {
	return a.mq
}

func (a *DefaultActor) setID(id ActorID) {
	a.id = id
}
