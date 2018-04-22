package examples

import (
	"fmt"
	"log"
	"sync"
)

type valueCounter struct {
	sync.Mutex // guard val
	val        int32
}

func (v *valueCounter) incrValue(id string) {
	v.Lock()
	defer v.Unlock()

	// critical section
	log.Printf("incring value: %s\n", id)
	v.val++
}

func increaseValueCounter(n int) int32 {
	vc := &valueCounter{val: 0}
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			vc.incrValue(fmt.Sprintf("%d", index))
		}(i)
	}

	wg.Wait()

	return vc.val
}
