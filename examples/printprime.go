package examples

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func PrintPrime() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Wait to finish")
	wg.Wait()
	fmt.Println("finished!")
}

func printPrime(prefix string) {
	defer wg.Done()

Loop:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue Loop
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer)
	}

	fmt.Println("Completed ", prefix)
}
