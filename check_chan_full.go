package examples

import "fmt"

func CheckChanIsFull() {
	ch := make(chan string, 1)

	ch <- "v"
	select {
	case ch <- "2":
	default:
		fmt.Println("channel is full")
	}
}
