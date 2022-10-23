package verify

import (
	"time"
)

// Eventually executes the test function until it returns an empty FailureMessage
// or timeout elapses.
func Eventually(timeout, interval time.Duration, fn func() FailureMessage) FailureMessage {
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	return EventuallyChan(timer.C, ticker.C, fn)
}

// EventuallyChan executes the test function until it returns an empty FailureMessage or timeout elapses.
func EventuallyChan[TTimerPayload, TTickPayload any](timeout <-chan (TTimerPayload), ticker <-chan (TTickPayload), fn func() FailureMessage) FailureMessage {
	var err string
	fail := func() FailureMessage {
		return FailureMessage("function never passed, last failure message:\n" + err)
	}

	for {
		select {
		case <-timeout:
			return fail()
		default:
		}

		err = string(fn())

		select {
		case <-timeout:
			return fail()
		default:
		}

		if err == "" {
			return ""
		}

		select {
		case <-timeout:
			return fail()
		case <-ticker:
		}
	}
}
