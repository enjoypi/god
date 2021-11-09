package stdlib

import (
	"sync"

	"github.com/enjoypi/god/types"
)

type ExitChan chan int
type GoRun func(ExitChan) (types.Message, error)
type OnReturn func(types.Message, error)

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

func (a *ActorManager) Get(id types.ActorID) Actor {
	return nil
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

func (a *ActorManager) Post(sender types.ActorID, receiver types.ActorID, message types.Message) {
	actor := a.Get(receiver)
	if actor != nil {
		actor.Post(message)
		return
	}
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
