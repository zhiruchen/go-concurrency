package main

import (
	"fmt"
)

func main() {
	for x := range randomBits() {
		fmt.Println(x)
	}
}

func randomBits() <-chan int {
	ch := make(chan int)

	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()

	return ch
}
