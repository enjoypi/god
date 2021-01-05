package core

import (
	"sync"
)

type ExitChan chan int
type GoRun func(ExitChan, Event) (Event, error)
type OnGoReturn func(Event, error)

type ActorManager struct {
	wg sync.WaitGroup
	ExitChan
}

var (
	DefaultActorManager ActorManager
)

func init() {
	DefaultActorManager.ExitChan = make(ExitChan)
}

func (a *ActorManager) Close() {
	close(a.ExitChan)
}

func (a *ActorManager) Go(run GoRun, event Event, onRet OnGoReturn) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		ret, err := run(a.ExitChan, event)
		if onRet != nil {
			onRet(ret, err)
		}
	}()
}

func (a *ActorManager) Wait() {
	a.wg.Wait()
}

func Close() {
	DefaultActorManager.Close()
}

func Go(run GoRun, parameter interface{}, onRet OnGoReturn) {
	DefaultActorManager.Go(run, parameter, onRet)
}

func Wait() {
	DefaultActorManager.Wait()
}
