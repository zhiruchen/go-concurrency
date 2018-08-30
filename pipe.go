package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	N := 100
	sendChan := make(chan int, N)
	receiveChan := make(chan int, N)

	go func() {
		for i := 1; i <= N; i++ {
			fmt.Printf("sending %d\n", i)
			sendChan <- i
		}
		close(sendChan)
	}()

	go processChannel(sendChan, receiveChan)

	for v := range receiveChan {
		fmt.Printf("received %d\n", v)
	}
}

func processChannel(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 100
	}

	close(out)
}
