package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestAnd(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		msg := verify.And()
		assertPassed(t, msg)
	})
	t.Run("NoneFailed", func(t *testing.T) {
		msg := verify.And("", "")
		assertPassed(t, msg)
	})
	t.Run("OneFailed", func(t *testing.T) {
		msg := verify.And("", "failure")
		assertFailed(t, msg, "failure")
	})
	t.Run("TwoFailed", func(t *testing.T) {
		msg := verify.And("one", "two")
		assertFailed(t, msg, "one\n\ntwo")
	})
}
