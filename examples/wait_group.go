package examples

import (
	"fmt"
	"sync"
)

// Doit wait group, show how to signal all workers to end their for loop
func Doit() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doit(i, wq, done, &wg)
	}

	for i := 0; i < workerCount; i++ {
		wq <- i
	}

	// close a channel all workers are receiving from, signal all workes to end  for loop
	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()

	for {
		select {
		case m := <-wq:
			fmt.Printf("[%d] m => %d\n", workerId, m)
		case <-done:
			fmt.Printf("[%d] is done\n", workerId)
			return
		}
	}
}
