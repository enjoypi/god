package core

import (
	"sync"
)

type ExitChan chan int
type GoRun func(ExitChan) (Message, error)
type OnReturn func(Message, error)

type ActorManager struct {
	wg sync.WaitGroup
	ExitChan
}

var (
	DefaultActorManager ActorManager
)

func init() {
	DefaultActorManager.ExitChan = make(ExitChan, 1)
}

func (a *ActorManager) Close() {
	close(a.ExitChan)
}

func (a *ActorManager) Go(run GoRun, onRet OnReturn) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		ret, err := run(a.ExitChan)
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

func Go(run GoRun, onRet OnReturn) {
	DefaultActorManager.Go(run, onRet)
}

func Wait() {
	DefaultActorManager.Wait()
}
