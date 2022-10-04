package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestTrue(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		got := verify.True(true)
		assertPassed(t, got)
	})
	t.Run("Failed", func(t *testing.T) {
		got := verify.True(false)
		assertFailed(t, got, "the value is false")
	})
}

func TestFalse(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		got := verify.False(false)
		assertPassed(t, got)
	})
	t.Run("Failed", func(t *testing.T) {
		got := verify.False(true)
		assertFailed(t, got, "the value is true")
	})
}
