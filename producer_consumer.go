package main

import (
	"fmt"
)

func produce(ch chan int, done chan bool) {
	for i := 0; i < 100; i += 10 {
		fmt.Printf("produce %d\n", i)
		ch <- i
	}
	done <- true
}

func consume(ch chan int) {
	for {
		fmt.Printf("consume %d\n", <-ch)
	}
}

func main() {
	done := make(chan bool)
	ch := make(chan int)

	go produce(ch, done)
	go consume(ch)

	<-done
}
