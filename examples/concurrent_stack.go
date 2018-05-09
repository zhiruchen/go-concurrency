package examples

import (
	"errors"
	"sync"
)

type concurrentStack struct {
	sync.RWMutex // guards follwing
	elements     []interface{}
	size         int32
}

func NewConcurrentStack() *concurrentStack {
	return &concurrentStack{elements: []interface{}{}, size: 0}
}

func (cs *concurrentStack) Pop() (v interface{}, err error) {
	cs.Lock()
	defer cs.Unlock()
	if cs.size <= 0 {
		return nil, errors.New("stack is empty")
	}

	v, cs.elements = cs.elements[len(cs.elements)-1], cs.elements[:len(cs.elements)-1]
	cs.size--
	return v, nil
}

func (cs *concurrentStack) Push(v interface{}) {
	cs.Lock()
	cs.elements = append(cs.elements, v)
	cs.size++
	cs.Unlock()
}

func (cs *concurrentStack) Size() int32 {
	cs.RLock()
	defer cs.RUnlock()
	return cs.size
}
