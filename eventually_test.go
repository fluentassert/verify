package verify_test

import (
	"testing"
	"time"

	"github.com/fluentassert/verify"
)

func TestEventually(t *testing.T) {
	assertNoLeak(t)
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("InitialPassed", func(t *testing.T) {
		msg := verify.Eventually(timeout, interval, func() verify.FailureMessage {
			return ""
		})
		assertPassed(t, msg)
	})
	t.Run("SecondPassed", func(t *testing.T) {
		shouldPass := false
		msg := verify.Eventually(timeout, interval, func() verify.FailureMessage {
			if !shouldPass {
				shouldPass = true // next exeucution will pass
				return "fail"
			}
			return ""
		})
		assertPassed(t, msg)
	})
	t.Run("ReturnedTooLate", func(t *testing.T) {
		msg := verify.Eventually(timeout, interval, func() verify.FailureMessage {
			time.Sleep(2 * timeout)
			return ""
		})
		assertFailed(t, msg, "function never passed, last failure message:\n")
	})
	t.Run("Failed", func(t *testing.T) {
		msg := verify.Eventually(timeout, interval, func() verify.FailureMessage {
			return "constant failure"
		})
		assertFailed(t, msg, "function never passed, last failure message:\nconstant failure")
	})
}

func TestEventuallyChan(t *testing.T) {
	assertNoLeak(t)
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("Passed", func(t *testing.T) {
		timer := time.NewTimer(timeout)
		defer timer.Stop()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		msg := verify.EventuallyChan(timer.C, ticker.C, func() verify.FailureMessage {
			return ""
		})
		assertPassed(t, msg)
	})
	t.Run("TimeoutBeforeStart", func(t *testing.T) {
		ch := make(chan struct{})
		close(ch)
		msg := verify.EventuallyChan(ch, ch, func() verify.FailureMessage {
			return ""
		})
		assertFailed(t, msg, "function never passed, last failure message:\n")
	})
	t.Run("Failed", func(t *testing.T) {
		ch := make(chan struct{})
		msg := verify.EventuallyChan(time.After(timeout), ch, func() verify.FailureMessage {
			return "constant failure"
		})
		assertFailed(t, msg, "function never passed, last failure message:\nconstant failure")
	})
}
