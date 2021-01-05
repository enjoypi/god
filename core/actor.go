package core

type Actor interface {
	Handle(event Event) // be called by goroutine
	ID() ActorID
	Initialize() error // must be called by supervisor
	Terminate()        // must thread safe
}
type NewActor func() Actor

type ActorImpl struct {
	id ActorID
}

func (a *ActorImpl) ID() ActorID {
	return a.id
}
