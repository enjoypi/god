package stdlib

import (
	"github.com/enjoypi/god/types"
)

type Supervisor struct {
	Actor
}

func NewSupervisor() *Supervisor {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		L.Panic(err.Error())
		return nil
	}

	return sup
}

func (sup *Supervisor) Initialize() error {
	return nil
}

func (sup *Supervisor) Handle(message types.Message) types.Message {
	return nil
}

func (sup *Supervisor) HandleActor(actor types.ActorID, message types.Message) {

}
func (sup *Supervisor) Start(actorType types.ActorType) Actor {
	actor := NewActor(actorType)
	if actor == nil {
		return nil
	}

	// actor must be initial before using, or maybe lock
	if err := actor.Initialize(); err != nil {
		L.Error(err.Error())
		return nil
	}

	Go(func(exitChan ExitChan) (types.Message, error) {
		defer actor.Terminate()

		mq := actor.messageQueue()
		for {
			select {
			case msg := <-mq:
				actor.Handle(msg)
			case <-exitChan:
				return types.EvStopped, nil
			}
		}
	}, func(message types.Message, err error) {
		sup.HandleActor(actor.ID(), message)
	})

	return actor
}
