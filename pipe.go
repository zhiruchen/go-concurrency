package main

import "fmt"

func main() {
	sendChan := make(chan int, 10)
	receiveChan := make(chan int, 10)

	go func() {
		for i := 1; i <= 10; i++ {
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
