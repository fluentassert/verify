package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestPanics(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		msg := f.Panics(func() { panic("exception") })
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		msg := f.Panics(func() {})
		assertFailed(t, msg, "the function returned instead of panicking")
	})
}

func TestNotPanics(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		msg := f.NotPanics(func() {})
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		msg := f.NotPanics(func() { panic("exception") })
		assertFailed(t, msg, "the function panicked")
	})
}
