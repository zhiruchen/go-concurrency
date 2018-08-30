package examples

import (
	"testing"
)

func TestGoroutineStart(t *testing.T) {
	calculateWhenStartGoroutine()

	var x []string
	if x == nil {
		t.Logf("x is nil: %v, len: %d\n", x, len(x))
	}
}
