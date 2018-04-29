package examples

import (
	"testing"
)

func TestTimeoutWithctx(t *testing.T) {
	timeoutWithCtx()
}

func TestPassValueWithCtx(t *testing.T) {
	passValueWithCtx("context value")
}

func TestCancelCtx(t *testing.T) {
	cancelCtx()
}

func TestDeadlineCtx(t *testing.T) {
	deadlineCtx()
}
