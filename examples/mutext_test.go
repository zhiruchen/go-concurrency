package examples

import "testing"

func TestIncreaseValueCounter(t *testing.T) {
	cases := []struct {
		n      int
		result int32
	}{
		{
			10,
			10,
		},
		{
			1,
			1,
		},
		{
			100,
			100,
		},
		{
			10000,
			10000,
		},
	}

	for _, cc := range cases {
		v := increaseValueCounter(cc.n)
		if cc.result != v {
			t.Errorf("Expect: %d, get: %d\n", cc.result, v)
		}
	}
}
