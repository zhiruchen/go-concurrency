package examples

import (
	"fmt"
	"sync"
)

func MergeTest() {
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

	for v := range Merge(done, c1, c2, c3) {
		fmt.Println(v)
	}
}

func Merge(done <-chan struct{}, cs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(sc <-chan string) {
			for v := range sc {
				select {
				case out <- v:
				case <-done:
					return
				}
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
