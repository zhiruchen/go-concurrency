package lock

import (
	"fmt"
	"sync"
)

func Race() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Printf("%d ", i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func RightRace() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(n int) {
			fmt.Printf("%d ", n)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
