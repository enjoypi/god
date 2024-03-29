package stdlib

import (
	"fmt"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/event"
	"github.com/enjoypi/god/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Supervisor struct {
	//actor.SimpleActor
}

func NewSupervisor() (*Supervisor, error) {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		return nil, fmt.Errorf("fail to initialize supervisor")
	}

	return sup, nil
}

func (sup *Supervisor) Initialize() error {
	return nil
}

func (sup *Supervisor) Handle(message def.Message) def.Message {
	return nil
}

func (sup *Supervisor) HandleActor(actor def.ActorID, message def.Message) {

}
func (sup *Supervisor) Start(v *viper.Viper, actorType def.ActorType, actorID def.ActorID) (Actor, error) {
	actor := NewActor(actorType, actorID)
	if actor == nil {
		return nil, fmt.Errorf("invalid actor type")
	}

	// actor must be initial before using, or maybe lock
	if err := actor.Initialize(v); err != nil {
		return nil, err
	}

	Go(func(exitChan ExitChan) (def.Message, error) {
		defer actor.Terminate()

		mq := actor.MessageQueue()

		for {

			select {
			case m := <-mq:
				//if ce := logger.L.Check(zapcore.DebugLevel, "RECV"); ce != nil {
				//	ce.Write(
				//		zap.String("type", string(actorType)),
				//		zap.Uint32("actor", actor.ID()),
				//		zap.Any("message", sc.TypeOf(m.Message)))
				//}
				if err := actor.Handle(m.Context, m.Message, sup); err != nil {
					logger.L.Warn("handle wrong", zap.Error(err))
				}
			case <-exitChan:
				return &event.EvStopped{}, nil
			}
		}
	}, func(message def.Message, err error) {
		sup.HandleActor(actor.ID(), message)
	})

	return actor, nil
}
