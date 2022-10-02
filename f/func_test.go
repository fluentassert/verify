package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestFunc(t *testing.T) {
	t.Run("Panics", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := f.Func(func() { panic("exception") }).Panics()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := f.Func(func() {}).Panics()
			assertFailed(t, msg, "the function returned instead of panicking")
		})
	})

	t.Run("NotPanics", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := f.Func(func() {}).NotPanics()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := f.Func(func() { panic("exception") }).NotPanics()
			assertFailed(t, msg, "the function panicked")
		})
	})
}
