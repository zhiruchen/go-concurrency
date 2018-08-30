package examples

import (
	"fmt"
)

func ForSelect() {
	ch1 := make(chan int32, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan int, 10)

	go chan1(ch1)

	go chan2(ch2)

	go chan3(ch3)

Loop:
	for {
		if ch1 == nil && ch2 == nil && ch3 == nil {
			break
		}

		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue Loop
			}
			fmt.Println("ch1: ", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue Loop
			}
			fmt.Println("ch2: ", v)
		case v, ok := <-ch3:
			if !ok {
				ch3 = nil
				continue Loop
			}
			fmt.Println("ch3: ", v)
		}
	}
}

func chan1(ch chan int32) {
	for i := 1; i <= 10; i++ {
		ch <- int32(i)
	}
	close(ch)
}

func chan2(ch chan string) {
	for i := 1; i <= 10; i++ {
		ch <- fmt.Sprintf("%d", i)
	}
	close(ch)
}

func chan3(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
