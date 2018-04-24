package main

import (
	"fmt"
	"time"
)

func main() {
	ch := createChan(10)
	go receiveChan(ch)
	time.Sleep(1e9)
}

func createChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	return ch
}

func receiveChan(ch chan int) {
	for {
		fmt.Printf("received -> %d\n", <-ch)
	}
}
