package actors

import (
	"sync"
)

type ExitChan chan int
type GoRun func(ExitChan, interface{}) (interface{}, error)
type OnGoReturn func(interface{}, error)

type Actors struct {
	wg sync.WaitGroup
	ExitChan
}

var (
	DefaultActors Actors
)

func init() {
	DefaultActors.ExitChan = make(ExitChan)
}

func (a *Actors) Close() {
	close(a.ExitChan)
}

func (a *Actors) Go(run GoRun, parameter interface{}, onRet OnGoReturn) {
	a.wg.Add(1)
	go func() {
		defer a.wg.Done()
		ret, err := run(a.ExitChan, parameter)
		if onRet != nil {
			onRet(ret, err)
		}
	}()
}

func (a *Actors) Wait() {
	a.wg.Wait()
}

func Close() {
	DefaultActors.Close()
}

func Go(run GoRun, parameter interface{}, onRet OnGoReturn) {
	DefaultActors.Go(run, parameter, onRet)
}

func Wait() {
	DefaultActors.Wait()
}
