package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/actors"
	"github.com/enjoypi/god/events"
	"github.com/enjoypi/god/logger"
	"github.com/enjoypi/god/types"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Supervisor struct {
	actors.SimpleActor
	*actors.ActorManager
}

func NewSupervisor() (*Supervisor, error) {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		return nil, fmt.Errorf("fail to initialize supervisor")
	}
	sup.ActorManager = actors.NewActorManager()

	return sup, nil
}

func (sup *Supervisor) Initialize() error {
	_ = sup.SimpleActor.Initialize()
	return nil
}

func (sup *Supervisor) Handle(message types.Message) types.Message {
	return nil
}

func (sup *Supervisor) HandleActor(actor types.ActorID, message types.Message) {

}
func (sup *Supervisor) Start(v *viper.Viper, actorType types.ActorType) (actors.Actor, error) {
	actor := actors.NewActor(actorType)
	if actor == nil {
		return nil, fmt.Errorf("invalid actor type")
	}

	// actor must be initial before using, or maybe lock
	if err := actor.Initialize(v); err != nil {
		return nil, err
	}

	actors.Go(func(exitChan actors.ExitChan) (types.Message, error) {
		defer actor.Terminate()

		mq := actor.MessageQueue()

		for {

			select {
			case msg := <-mq:
				if ce := logger.L.Check(zapcore.DebugLevel, "RECV"); ce != nil {
					ce.Write(
						zap.String("type", actorType),
						zap.String("actor", fmt.Sprintf("%p", actor)),
						zap.Any("message", msg))
				}
				if err := actor.Handle(msg); err != nil {
					logger.L.Warn("handle wrong", zap.Error(err))
				}
			case <-exitChan:
				return &events.EvStopped{}, nil
			}
		}
	}, func(message types.Message, err error) {
		sup.HandleActor(actor.ID(), message)
	})

	return actor, nil
}
