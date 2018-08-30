package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	var n = 1000
	ch := make(chan int, n)
	go pump(n, ch)
Loop:
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				break Loop
			}
			fmt.Printf("received: %d ||", v)
		}
	}
}

func pump(n int, ch chan int) {
	for i := 0; i < n; i++ {
		fmt.Printf("sending %d ||", i)
		ch <- i
	}

	close(ch)
}
