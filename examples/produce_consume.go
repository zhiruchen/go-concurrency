package examples

import (
	"runtime"
	"sync"

	"github.com/labstack/gommon/log"
)

func pc() {
	runtime.GOMAXPROCS(4)
	c := make(chan int, 10)
	done := make(chan struct{})
	go consume(c, done)
	go produce(c)

	<-done
}

func produce(c chan<- int) {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(v int) {
			c <- v
			log.Printf("producing: %d\n", v)
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
}

func consume(c <-chan int, done chan<- struct{}) {
Loop:
	for {
		select {
		case v, ok := <-c:
			if !ok {
				done <- struct{}{}
				break Loop
			}
			log.Printf("consuming: %d\n", v)
		}
	}

	log.Printf("end consume \n")
}
