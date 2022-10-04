package verify_test

import (
	"context"
	"testing"
	"time"

	"github.com/fluentassert/verify"
)

func TestPeriodic(t *testing.T) {
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("Eventually", func(t *testing.T) {
		t.Run("InitialPassed", func(t *testing.T) {
			msg := verify.Periodic(timeout, interval, func() verify.FailureMessage {
				return ""
			}).Eventually()
			assertPassed(t, msg)
		})
		t.Run("SecondPassed", func(t *testing.T) {
			shouldPass := false
			msg := verify.Periodic(timeout, interval, func() verify.FailureMessage {
				if !shouldPass {
					shouldPass = true // next exeucution will pass
					return "fail"
				}
				return ""
			}).Eventually()
			assertPassed(t, msg)
		})
		t.Run("NeverReturned", func(t *testing.T) {
			ch := make(chan struct{}, 1)
			defer close(ch)
			msg := verify.Periodic(timeout, interval, func() verify.FailureMessage {
				<-ch
				return ""
			}).Eventually()
			assertFailed(t, msg, "timeout\nfunction has never returned")
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.Periodic(timeout, interval, func() verify.FailureMessage {
				return "constant failure"
			}).Eventually()
			assertFailed(t, msg, "timeout\nfunction always failed\nlast failure message:\nconstant failure")
		})
	})

	t.Run("EventuallyContext", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.Periodic(timeout, interval, func() verify.FailureMessage {
				return ""
			}).EventuallyContext(context.Background())
			assertPassed(t, msg)
		})
	})
}

func TestAsync(t *testing.T) {
	timeout := 100 * time.Millisecond
	interval := 10 * time.Millisecond

	t.Run("Eventually", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			timer := time.NewTimer(timeout)
			defer timer.Stop()
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			msg := verify.Async(timer.C, ticker.C, func() verify.FailureMessage {
				return ""
			}).Eventually()
			assertPassed(t, msg)
		})
		t.Run("TimeoutBeforeStart", func(t *testing.T) {
			ch := make(chan struct{})
			close(ch)
			msg := verify.Async(ch, ch, func() verify.FailureMessage {
				return ""
			}).Eventually()
			assertFailed(t, msg, "timeout\nfunction has never returned")
		})
	})

	t.Run("EventuallyContext", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			timer := time.NewTimer(timeout)
			defer timer.Stop()
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			msg := verify.Async(timer.C, ticker.C, func() verify.FailureMessage {
				return ""
			}).EventuallyContext(context.Background())
			assertPassed(t, msg)
		})
		t.Run("DoneBeforeStart", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			ch := make(chan struct{})
			msg := verify.Async(ch, ch, func() verify.FailureMessage {
				return ""
			}).EventuallyContext(ctx)
			assertFailed(t, msg, "context canceled\nfunction has never returned")
		})
		t.Run("Failed", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			ch := make(chan struct{})
			msg := verify.Async(ch, ch, func() verify.FailureMessage {
				return "constant failure"
			}).EventuallyContext(ctx)
			assertFailed(t, msg, "context deadline exceeded\nfunction always failed\nlast failure message:\nconstant failure")
		})
	})
}
