package examples

import (
	"fmt"
	"runtime"
	"sync"
)

func RunWaitGroup() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			for c := 'a'; c < 'a'+26; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 3; i++ {
			for c := 'A'; c < 'A'+26; c++ {
				fmt.Printf("%c ", c)
			}
		}
	}()

	fmt.Println("waiting to finish!")
	wg.Wait()
	fmt.Println("finished!")
}
