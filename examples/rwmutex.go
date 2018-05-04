package examples

import (
	"sync"
)

type concurrentMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func newCMap() *concurrentMap {
	return &concurrentMap{m: make(map[string]int)}
}

func (cmap *concurrentMap) Get(key string) (int, bool) {
	cmap.mu.RLock() // RLock for reading
	defer cmap.mu.RUnlock()
	v, ok := cmap.m[key]
	return v, ok
}

func (cmap *concurrentMap) Put(k string, v int) {
	cmap.mu.Lock() // Lock for writing
	defer cmap.mu.Unlock()
	cmap.m[k] = v
}

func (cmap *concurrentMap) Delete(k string) {
	cmap.mu.Lock() // Lock for writing
	defer cmap.mu.Unlock()
	delete(cmap.m, k)
}

func (cmap *concurrentMap) Count() int {
	cmap.mu.RLock()
	defer cmap.mu.RUnlock()
	return len(cmap.m)
}

func (cmap *concurrentMap) Range() {

}
