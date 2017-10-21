package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int64)
	go sum([]int64{1, 3, 4, 5, 6, 10}, ch)

	fmt.Printf("sum = %d\n", <-ch)
}

func sum(numbers []int64, sum chan int64) {
	total := int64(0)
	for _, n := range numbers {
		total += n
		time.Sleep(time.Millisecond)
	}

	sum <- total
}
