package core

import "github.com/enjoypi/god"

type Supervisor struct {
	Actor
}

func NewSupervisor() *Supervisor {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		god.Logger.Panic(err.Error())
		return nil
	}

	return sup
}

func (sup *Supervisor) Initialize() error {
	return nil
}

func (sup *Supervisor) Handle(message Message) Message {
	return nil
}

func (sup *Supervisor) HandleActor(actor ActorID, message Message) {

}
func (sup *Supervisor) Start(actorType ActorType) Actor {
	actor := NewActor(actorType)
	if actor == nil {
		return nil
	}

	// actor must be initial before using, or maybe lock
	if err := actor.Initialize(); err != nil {
		god.Logger.Error(err.Error())
		return nil
	}

	Go(func(exitChan ExitChan) (Message, error) {
		defer actor.Terminate()

		mq := actor.messageQueue()
		for {
			select {
			case msg := <-mq:
				actor.Handle(msg)
			case <-exitChan:
				return EvStopped, nil
			}
		}
	}, func(message Message, err error) {
		sup.HandleActor(actor.ID(), message)
	})

	return actor
}
