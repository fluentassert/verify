package f

import (
	"context"
	"time"
)

// Eventually periodically executes the test function until it returns an empty FailureMessage
// or timeout elapses.
func Eventually(timeout, interval time.Duration, fn func() FailureMessage) FailureMessage {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var (
		msg  FailureMessage
		tick <-chan time.Time
		ch   = make(chan FailureMessage, 1)
	)
	go func() { ch <- fn() }()
	for {
		select {
		case <-timer.C:
			if msg == "" {
				return "function has never returned"
			}
			return "function always failed, last failure message:\n" + msg
		case <-tick:
			tick = nil
			go func() { ch <- fn() }()
		case msg = <-ch:
			if msg == "" {
				return ""
			}
			tick = ticker.C
		}
	}
}

// EventuallyContext periodically executes the test function until it returns an empty FailureMessage
// or the context signals done.
func EventuallyContext(ctx context.Context, interval time.Duration, fn func() FailureMessage) FailureMessage {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	var (
		msg  FailureMessage
		tick <-chan time.Time
		ch   = make(chan FailureMessage, 1)
	)
	go func() { ch <- fn() }()
	for {
		select {
		case <-ctx.Done():
			if msg == "" {
				return "function has never returned"
			}
			return "function always failed, last failure message:\n" + msg
		case <-tick:
			tick = nil
			go func() { ch <- fn() }()
		case msg = <-ch:
			if msg == "" {
				return ""
			}
			tick = ticker.C
		}
	}
}
