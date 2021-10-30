package stdlib

import "sync"

type Manager struct {
	kv  sync.Map
	v   sync.Map
	idV sync.Map
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Store(key, value interface{}) {
	m.kv.Store(key, value)
}
