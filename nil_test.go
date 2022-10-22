package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestNil(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		var err error
		msg := verify.Nil(err)
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		msg := verify.Nil(0)
		assertFailed(t, msg, "value is not nil")
	})
}

func TestNotNil(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		msg := verify.NotNil(0)
		assertPassed(t, msg)
	})
	t.Run("Failed", func(t *testing.T) {
		var err error
		msg := verify.NotNil(err)
		assertFailed(t, msg, "value is <nil>")
	})
}
