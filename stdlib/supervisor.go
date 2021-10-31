package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/types"
	"go.uber.org/zap"
)

type Supervisor struct {
	DefaultActor
}

func NewSupervisor() (*Supervisor, error) {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		return nil, fmt.Errorf("fail to initialize supervisor")
	}

	return sup, nil
}

func (sup *Supervisor) Initialize() error {
	_ = sup.DefaultActor.Initialize()
	return nil
}

func (sup *Supervisor) Handle(message types.Message) types.Message {
	return nil
}

func (sup *Supervisor) HandleActor(actor types.ActorID, message types.Message) {

}
func (sup *Supervisor) Start(actorType types.ActorType) (Actor, error) {
	actor := NewActor(actorType)
	if actor == nil {
		return nil, fmt.Errorf("invalid actor type")
	}

	// actor must be initial before using, or maybe lock
	if err := actor.Initialize(); err != nil {
		return nil, err
	}

	Go(func(exitChan ExitChan) (types.Message, error) {
		defer actor.Terminate()

		mq := actor.messageQueue()

		for {

			select {
			case msg := <-mq:
				L.Debug("receive message",
					zap.String("actor", fmt.Sprintf("%p", actor)),
					zap.String("mq", fmt.Sprintf("%p", mq)),
					zap.Any("message", msg))
				actor.Handle(msg)
			case <-exitChan:
				return types.EvStopped, nil
			}
		}
	}, func(message types.Message, err error) {
		sup.HandleActor(actor.ID(), message)
	})

	return actor, nil
}
