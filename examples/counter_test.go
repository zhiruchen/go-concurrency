package examples

import (
	"testing"
)

func TestIncrCounter(t *testing.T) {
	IncrCounter()
}

func TestAtomicIncrCounter(t *testing.T) {
	AtomicIncrCounter()
}
