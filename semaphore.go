package main

import (
	"fmt"
	"time"
)

func main() {
	for v := range recvBufferedChan() {
		fmt.Printf("%d, ", v)
	}
}

func sum(numbers []int64, sum chan int64) {
	total := int64(0)
	for _, n := range numbers {
		total += n
		time.Sleep(time.Millisecond)
	}

	sum <- total
}

func bufferChanSum() {
	ll := [][]int32{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11},
		{100,1000,999,8888},
	}

	results := make(chan int32, len(ll))
	for _, l := range ll {
		go func(intList []int32) {
			var sum int32
			for _, v := range intList {
				sum+= v
			}
			results <- sum
		}(l)
	}

	for i := 1; i <= len(ll); i++ {
		fmt.Println(<-results)
	}
	close(results)
}

func sendToClosedChan() {
	c := make(chan int)
	close(c)
	c <- 100
}

func receiveFromClosedChan() {
	c := make(chan int)
	close(c)
	v ,ok := <-c
	fmt.Println(v, ok)
}

func recvBufferedChan() chan int {
	vs := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		vs <- i*10
	}
	close(vs)
	return vs
}