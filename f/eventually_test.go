package f_test

import (
	"context"
	"testing"
	"time"

	"github.com/pellared/fluentassert/f"
)

func TestEventually(t *testing.T) {
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("InitialPassed", func(t *testing.T) {
		msg := f.Eventually(timeout, interval, func() f.FailureMessage {
			return ""
		})
		assertPassed(t, msg)
	})

	t.Run("SecondPassed", func(t *testing.T) {
		shouldPass := false
		msg := f.Eventually(timeout, interval, func() f.FailureMessage {
			if !shouldPass {
				shouldPass = true // next exeucution will pass
				return "fail"
			}
			return ""
		})
		assertPassed(t, msg)
	})

	t.Run("NeverReturned", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		defer close(ch)
		msg := f.Eventually(timeout, interval, func() f.FailureMessage {
			<-ch
			return ""
		})
		assertFailed(t, msg, "function has never returned")
	})

	t.Run("Failed", func(t *testing.T) {
		msg := f.Eventually(timeout, interval, func() f.FailureMessage {
			return "constant failure"
		})
		assertFailed(t, msg, "function always failed, last failure message:\nconstant failure")
	})
}

func TestEventuallyContext(t *testing.T) {
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("InitialPassed", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		msg := f.EventuallyContext(ctx, interval, func() f.FailureMessage {
			return ""
		})
		assertPassed(t, msg)
	})

	t.Run("SecondPassed", func(t *testing.T) {
		shouldPass := false
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		msg := f.EventuallyContext(ctx, interval, func() f.FailureMessage {
			if !shouldPass {
				shouldPass = true // next exeucution will pass
				return "fail"
			}
			return ""
		})
		assertPassed(t, msg)
	})

	t.Run("NeverReturned", func(t *testing.T) {
		ch := make(chan struct{}, 1)
		defer close(ch)
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		msg := f.EventuallyContext(ctx, interval, func() f.FailureMessage {
			<-ch
			return ""
		})
		assertFailed(t, msg, "function has never returned")
	})

	t.Run("Failed", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		msg := f.EventuallyContext(ctx, interval, func() f.FailureMessage {
			return "constant failure"
		})
		assertFailed(t, msg, "function always failed, last failure message:\nconstant failure")
	})
}
