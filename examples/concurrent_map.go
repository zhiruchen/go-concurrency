package examples

import (
	"sync"
)

type ConcurrentMap struct {
	mu sync.RWMutex
	m  map[string]interface{}
}

func NewConcurrentMap(size int) *ConcurrentMap {
	s := 0
	if size > 0 {
		s = size
	}
	return &ConcurrentMap{
		m: make(map[string]interface{}, s),
	}
}

func (m *ConcurrentMap) Put(key string, val interface{}) {
	m.mu.Lock()
	m.m[key] = val
	m.mu.Unlock()
}

func (m *ConcurrentMap) Get(key string) (interface{}, bool) {
	m.mu.RLock()
	v, ok := m.m[key]
	m.mu.RUnlock()
	return v, ok
}

func (m *ConcurrentMap) Size() int {
	m.mu.RLock()
	size := len(m.m)
	m.mu.RUnlock()

	return size
}
