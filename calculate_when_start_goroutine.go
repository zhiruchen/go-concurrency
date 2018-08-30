package examples

import (
	"fmt"
	"time"
)

func calculateWhenStartGoroutine() {
	done := make(chan struct{})
	go func(v int) {
		time.Sleep(time.Second)
		fmt.Println(v * 1000)
		done <- struct{}{}
	}(calculate(100))

	<-done
}

func calculate(i int) int {
	fmt.Println("calculate value...")
	return i * 100
}
