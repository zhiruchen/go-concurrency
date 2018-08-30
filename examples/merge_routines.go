package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MergeTest() {
	rand.Seed(time.Now().Unix())
	c1 := genCh()
	c2 := genCh()
	c3 := genCh()
	c4 := genCh()
	done := make(chan struct{})

	var vs []int32
	for v := range Merge(done, c1, c2, c3, c4) {
		vs = append(vs, v)
	}
	fmt.Println("merged vals: ", vs)
}

func genCh() <-chan int32 {
	vLen := rand.Intn(10)
	var addMap = make(map[int32]bool)
	var vs []int32
	c := make(chan int32, vLen)
	for len(vs) < vLen {
		v := rand.Int31n(10)
		if _, ok := addMap[v]; ok {
			continue
		}

		addMap[v] = true
		c <- v
		vs = append(vs, v)
	}

	fmt.Println("gen vals: ", vs)
	close(c)
	return c
}

func Merge(done <-chan struct{}, cs ...<-chan int32) <-chan int32 {
	out := make(chan int32)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(sc <-chan int32) {
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
