package core

type DefaultImplement interface {
	ID() ActorID
	MessageQueue() MessageQueue

	Post(message Message) // post message to actor's message queue, must thread safe
}

type Actor interface {
	Handle(message Message) Message // be called by actor's goroutine
	Initialize() error              // must be called by supervisor
	Terminate()                     // must thread safe

	DefaultImplement
}
type NewActor func() Actor

type ActorImpl struct {
	id ActorID
	mq MessageQueue
}

func (a *ActorImpl) MessageQueue() MessageQueue {
	return a.mq
}

func (a *ActorImpl) ID() ActorID {
	return a.id
}

func (a *ActorImpl) Post(message Message) {
	if a.mq == nil {
		a.mq = make(MessageQueue, 1)
	}
	a.mq <- message
}
