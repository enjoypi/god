package actors

import (
	"sync"

	"github.com/enjoypi/god/def"
)

type ExitChan chan int
type GoRun func(ExitChan) (def.Message, error)
type OnReturn func(def.Message, error)

type Manager struct {
	actors sync.Map
	wg     sync.WaitGroup
	ExitChan
}

var (
	defaultManager *Manager
)

func init() {
	defaultManager = NewManager()
}

func NewManager() *Manager {
	return &Manager{
		ExitChan: make(ExitChan, 1),
	}
}

func (m *Manager) Close() {
	close(m.ExitChan)
}

func (m *Manager) Get(id def.ActorID) Actor {
	i, _ := m.actors.Load(id)
	return i.(Actor)
}

func (m *Manager) Go(run GoRun, onRet OnReturn) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		ret, err := run(m.ExitChan)
		if onRet != nil {
			onRet(ret, err)
		}
	}()
}

func (m *Manager) Post(receiver def.ActorID, message def.Message) {
	actor := m.Get(receiver)
	if actor != nil {
		actor.Post(message)
		return
	}
}

func (m *Manager) Wait() {
	m.wg.Wait()
}

func Close() {
	defaultManager.Close()
}

func Go(run GoRun, onRet OnReturn) {
	defaultManager.Go(run, onRet)
}

func Post(receiver def.ActorID, message def.Message) {
	defaultManager.Post(receiver, message)
}

func Wait() {
	defaultManager.Wait()
}
