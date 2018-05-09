package examples

import (
	"reflect"
	"testing"
)

func TestConcurrentStack(t *testing.T) {
	cases := []struct {
		es     []interface{}
		size   int32
		result []interface{}
	}{
		{
			es:     []interface{}{1, 2, 3},
			size:   3,
			result: []interface{}{3, 2, 1},
		},
		{
			es:     []interface{}{"a", "b", "c"},
			size:   3,
			result: []interface{}{"c", "b", "a"},
		},
		{
			es: []interface{}{
				[]int{1},
				[]int{2},
				[]int{3},
			},
			size: 3,
			result: []interface{}{
				[]int{3},
				[]int{2},
				[]int{1},
			},
		},
	}

	for _, cc := range cases {
		s := NewConcurrentStack()
		for _, v := range cc.es {
			s.Push(v)
		}

		size := s.Size()
		if cc.size != size {
			t.Errorf("exepct: %d, get: %d\n", cc.size, size)
		}

		var outes []interface{}
		var i int32
		for i = 0; i < size; i++ {
			v, _ := s.Pop()
			outes = append(outes, v)
		}

		t.Logf("result: %v, outes: %v\n", cc.result, outes)

		if !reflect.DeepEqual(cc.result, outes) {
			t.Errorf("expect: %v, get: %v\n", cc.result, outes)
		}
	}
}
