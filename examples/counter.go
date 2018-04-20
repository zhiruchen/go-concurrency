package examples

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var (
	counter int32
)

func IncrCounter() {
	wg.Add(2)
	go incrCounter()
	go incrCounter()

	wg.Wait()
	fmt.Println(counter)
}

func incrCounter() {
	defer wg.Done()

	for count := 1; count <= 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

func AtomicIncrCounter() {
	wg.Add(2)
	go atomicIncrCounter()
	go atomicIncrCounter()

	wg.Wait()
	fmt.Println(counter)
}

func atomicIncrCounter() {
	defer wg.Done()

	for count := 1; count <= 2; count++ {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}
