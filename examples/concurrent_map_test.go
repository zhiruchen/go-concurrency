package examples

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	connm := NewConcurrentMap(10)

	mm := map[string]interface{}{
		"1":   1,
		"2":   2,
		"3":   4,
		"4":   4,
		"5":   6,
		"100": 1000,
	}

	var wg sync.WaitGroup

	for k, v := range mm {
		wg.Add(1)
		go func(key string, val interface{}) {
			connm.Put(key, val)
			log.Println(fmt.Sprintf("put key: %s, val: %v", key, val))
			wg.Done()
		}(k, v)
	}

	for _, k := range []string{"2", "3", "4", "5", "100", "0"} {
		wg.Add(1)
		go func(k string) {
			v, ok := connm.Get(k)
			log.Println(fmt.Sprintf("key: %s, v: %v, ok: %t", k, v, ok))
			wg.Done()
		}(k)
	}

	wg.Wait()
}
