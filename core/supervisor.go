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

func (sup *Supervisor) Handle(event Event) {

}
func (sup *Supervisor) HandleActor(actor ActorID, event Event) {

}
func (sup *Supervisor) Start(actorType ActorType) bool {
	actor := defaultActorFactory.new(actorType)
	if actor == nil {
		return false
	}

	Go(func(exitChan ExitChan, event Event) (Event, error) {
		if err := actor.Initialize(); err != nil {
			return nil, err
		}

		select {
		case <-exitChan:
			actor.Terminate()
			return EvStopped, nil
		}
	}, nil, func(event Event, err error) {
		sup.HandleActor(actor.ID(), event)
	})

	return true
}
