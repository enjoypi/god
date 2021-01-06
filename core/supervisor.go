package core

type Supervisor struct {
	Actor
}

func NewSupervisor() *Supervisor {
	sup := &Supervisor{}
	if err := sup.Initialize(); err != nil {
		logger.Panic(err.Error())
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
	actor := defaultActorFactory.new(actorType)
	if actor == nil {
		return nil
	}

	Go(func(exitChan ExitChan) (Message, error) {
		if err := actor.Initialize(); err != nil {
			return nil, err
		}
		defer actor.Terminate()

		mq := actor.MessageQueue()
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
