package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go genNumber(ch1)
	go genNumber1(ch2)

	go printNumber(ch1, ch2)

	time.Sleep(1e6)
}

func genNumber(ch1 chan int) {
	for i := 1; ; i++ {
		ch1 <- i * 2
	}
}

func genNumber1(ch2 chan int) {
	for i := 1; ; i++ {
		ch2 <- i + 10
	}
}

func printNumber(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("received from ch1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("received from ch2: %d\n", v)
		}
	}
}
