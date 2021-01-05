package core

type Actor interface {
	Handle(event Event) // be called by goroutine
	ID() ActorID
	Impl() *ActorImpl
	Initialize() error // must be called by supervisor
	Terminate()        // must thread safe
}

type ActorImpl struct {
	id ActorID
}

type NewActor func() Actor

func (a *ActorImpl) ID() ActorID {
	return a.id
}

func (a *ActorImpl) Run(exitChan ExitChan, event Event) (Event, error) {
	select {
	case <-exitChan:
		return nil, nil
	}
}

const (
	EvStopped = iota
	EvPanic
)
