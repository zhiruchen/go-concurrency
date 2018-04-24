package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan string, 3)
	c2 := make(chan string, 3)
	c3 := make(chan string, 3)
	done := make(chan struct{})
	defer close(done)

	vs := []string{"hello", "my", "friend"}

	for i := 1; i <= 3; i++ {
		c1 <- vs[i-1]
	}

	for i := 1; i <= 3; i++ {
		c2 <- vs[i-1]
	}

	for i := 1; i <= 3; i++ {
		c3 <- vs[i-1]
	}
	close(c1)
	close(c2)
	close(c3)

	fmt.Printf("start merge...\n")
	for v := range merge(done, c1, c2, c3) {
		fmt.Println(v)
	}
}

func merge(done <-chan struct{}, cs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	output := func(c <-chan string) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
