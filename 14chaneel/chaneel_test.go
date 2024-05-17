package chaneel

import (
	"testing"

	"go.uber.org/goleak"
)

func TestLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	Chaneel() // Chaneel関数に対してgoroutine leakをチェックする。
}
