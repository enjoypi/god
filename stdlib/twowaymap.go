package stdlib

import "sync"

type TwoWayMap struct {
	concurrent bool

	kv    map[interface{}]interface{}
	vk    map[interface{}]interface{}
	mutex *sync.Mutex
}

func NewTwoWayMap(concurrent bool) *TwoWayMap {
	m := &TwoWayMap{concurrent: concurrent}
	if concurrent {
		m.mutex = &sync.Mutex{}
	}
	m.kv = make(map[interface{}]interface{})
	m.vk = make(map[interface{}]interface{})
	return m
}

// Delete deletes the value for a key.
func (m *TwoWayMap) Delete(key interface{}) {
	if !m.concurrent {
		goto do
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()

do:
	if v, ok := m.kv[key]; ok {
		delete(m.kv, key)
		delete(m.kv, v)
	}

}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *TwoWayMap) Load(key interface{}) (value interface{}, ok bool) {
	return nil, false
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *TwoWayMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return nil, false
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *TwoWayMap) LoadAndDelete(key interface{}) (value interface{}, loaded bool) {
	return nil, false
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *TwoWayMap) Range(f func(key, value interface{}) bool) {
}

// Store sets the value for a key.
func (m *TwoWayMap) Store(key, value interface{}) {
	if !m.concurrent {
		goto do
	}
	m.mutex.Lock()
	defer m.mutex.Lock()

do:
	m.kv[key] = value
	m.vk[value] = key
}
