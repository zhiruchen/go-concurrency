package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)
	go pump(10, ch)
	for v := range ch {
		fmt.Printf("received: %d\n", v)
	}
}

func pump(n int, ch chan int) {
	for i := 0; i < n; i++ {
		fmt.Printf("sending %d\n", i)
		ch <- i
	}

	close(ch)
}
