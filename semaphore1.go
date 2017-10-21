package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Empty interface{}

var empty Empty

func main() {
	N := 10
	data := make([]int64, N)
	sem := make(chan Empty, N)
	res := make([]int64, N)

	rand.Seed(42)
	for i := 0; i < N; i++ {
		data[i] = rand.Int63n(100)
	}

	for i, n := range data {
		go func(i int, n int64) {
			fmt.Printf("calculating %d, %d\n", i, n)
			res[i] = doSomeCalculate(i, n)
			sem <- empty
		}(i, n)
	}

	for i := 1; i <= N; i++ {
		<-sem
	}

	fmt.Printf("after all calculation done!")
}

func doSomeCalculate(i int, n int64) int64 {
	v := int64(i) * n
	time.Sleep(100 * time.Millisecond)
	return v
}
