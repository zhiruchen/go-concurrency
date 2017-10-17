package main

import (
	"log"
	"sync"
)

func main() {
	people := []string{"zhangsna", "Lisi", "Wangmazi", "Xiaoming"}
	match := make(chan string, 1)

	wg := new(sync.WaitGroup)
	for _, name := range people {
		wg.Add(1)
		go seek(name, match, wg)
	}
	wg.Wait()

	select {
	case name := <-match:
		log.Printf("No one received %s's message.\n", name)
	default:
	}
}

func seek(name string, match chan string, wg *sync.WaitGroup) {
	select {
	case peer := <-match:
		log.Printf("%s received a message from %s.\n", name, peer)
	case match <- name:
	}

	wg.Done()
}
