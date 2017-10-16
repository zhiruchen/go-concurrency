package lock

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	mu sync.Mutex
	v  int
}

func (a *AtomicInt) Add(n int) {
	a.mu.Lock()
	a.v += n
	a.mu.Unlock()
}

func (a *AtomicInt) Read() int {
	a.mu.Lock()
	n := a.v
	a.mu.Unlock()
	return n
}

func Main() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1)
		close(wait)
	}()

	n.Add(2)
	<-wait
	fmt.Println(n.Read())
}
