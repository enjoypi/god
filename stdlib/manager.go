package stdlib

import (
	"context"
	"sync"

	"github.com/enjoypi/god/def"
	"github.com/enjoypi/god/logger"
	"go.uber.org/atomic"
	"go.uber.org/zap"
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
	defaultActorID atomic.Uint32
)

func init() {
	defaultManager = NewManager()
	defaultActorID.Store(uint32(def.AIDUser))
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

func (m *Manager) NewActor(actorType def.ActorType, id def.ActorID) Actor {
	actor := defaultFactory.NewActor(actorType, id)
	if actor == nil {
		return nil
	}

	if _, loaded := m.actors.LoadOrStore(id, actor); loaded {
		logger.L.Warn("exists actor",
			zap.String("type", string(actorType)),
			zap.Uint32("id", id),
		)
	}
	return actor
}

func (m *Manager) Post(ctx context.Context, receiver def.ActorID, message def.Message) {
	actor := m.Get(receiver)
	if actor != nil {
		actor.Post(ctx, message)
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

func NewActor(actorType def.ActorType, id def.ActorID) Actor {
	if id == 0 {
		id = defaultActorID.Inc()
	}
	return defaultManager.NewActor(actorType, id)
}

func Post(ctx context.Context, receiver def.ActorID, message def.Message) {
	defaultManager.Post(ctx, receiver, message)
}

func Wait() {
	defaultManager.Wait()
}
