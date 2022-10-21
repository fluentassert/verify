package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestOr(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		msg := verify.Or()
		assertPassed(t, msg)
	})
	t.Run("NoneFailed", func(t *testing.T) {
		msg := verify.Or("", "")
		assertPassed(t, msg)
	})
	t.Run("OneFailed", func(t *testing.T) {
		msg := verify.Or("", "failure")
		assertPassed(t, msg)
	})
	t.Run("TwoFailed", func(t *testing.T) {
		msg := verify.Or("one", "two")
		assertFailed(t, msg, "one\n\ntwo")
	})
}
