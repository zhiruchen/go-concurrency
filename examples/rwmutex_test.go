package examples

import (
	"testing"
)

func TestConcurrentMap_Get(t *testing.T) {
	cases := []struct {
		m   map[string]int
		key string
		val int
		ok  bool
	}{
		{
			m: map[string]int{
				"a": 1,
				"b": 2,
				"c": 9,
			},
			key: "b",
			val: 2,
			ok:  true,
		},
		{
			m: map[string]int{
				"a": 1,
				"b": 2,
				"c": 9,
			},
			key: "c",
			val: 9,
			ok:  true,
		},
		{
			m: map[string]int{
				"a": 1,
				"b": 2,
				"c": 9,
			},
			key: "x",
			val: 0,
			ok:  false,
		},
	}

	for _, cc := range cases {
		cmap := newCMap()

		for k, v := range cc.m {
			cmap.Put(k, v)
		}

		v, ok := cmap.Get(cc.key)
		if v != cc.val || ok != cc.ok {
			if v != cc.val {
				t.Errorf("expect: %v, get: %v\n", cc.val, v)
			}

			if ok != cc.ok {
				t.Errorf("expect: %v, get: %v\n", cc.ok, ok)
			}
		}
	}
}

func TestConcurrentMap_Put(t *testing.T) {
	cases := []struct {
		m     map[string]int
		count int
	}{
		{
			m: map[string]int{
				"a": 1,
				"b": 2,
				"c": 9,
			},
			count: 3,
		},
		{
			m: map[string]int{
				"a": 1,
			},
			count: 1,
		},
		{
			m:     map[string]int{},
			count: 0,
		},
	}

	for _, cc := range cases {
		cmap := newCMap()
		for k, v := range cc.m {
			cmap.Put(k, v)
		}

		count := cmap.Count()
		if count != cc.count {
			t.Errorf("expect: %v, get: %v\n", cc.count, count)
		}
	}
}

func TestConcurrentMap_Remove(t *testing.T) {
}

func TestConcurrentMap_Count(t *testing.T) {
}

func BenchmarkCMap_Get(b *testing.B) {
}

func BenchmarkCMap_Put(b *testing.B) {

}
