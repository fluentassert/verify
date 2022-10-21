package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestPanics(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		msg := verify.Panics(func() { panic("exception") })
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		msg := verify.Panics(func() {})
		assertFailed(t, msg, "the function returned instead of panicking")
	})
}

func TestNotPanics(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		msg := verify.NotPanics(func() {})
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		msg := verify.NotPanics(func() { panic("exception") })
		assertFailed(t, msg, "the function panicked")
	})
}
