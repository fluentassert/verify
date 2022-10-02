package f

import (
	"context"
	"time"
)

// FluentAsync encapsulates asynchronous assertions using channels.
type FluentAsync[TFn ~func() FailureMessage, TTimerPayload, TTickPayload any] struct {
	Fn      TFn
	Timeout <-chan (TTimerPayload)
	Ticker  <-chan (TTickPayload)
}

// Async is used for asynchronous testing using channels for timeout and ticking.
func Async[TFn ~func() FailureMessage, TTimerPayload, TTickPayload any](timeout <-chan (TTimerPayload), ticker <-chan (TTickPayload), fn TFn) FluentAsync[TFn, TTimerPayload, TTickPayload] {
	return FluentAsync[TFn, TTimerPayload, TTickPayload]{fn, timeout, ticker}
}

// Eventually executes the test function until it returns an empty FailureMessage
// or timeout elapses.
func (x FluentAsync[TFn, TTimerPayload, TTickPayload]) Eventually() FailureMessage {
	ctx := context.Background()
	return eventually(ctx, x.Timeout, x.Ticker, x.Fn)
}

// EventuallyContext executes the test function until it returns an empty FailureMessage,
// timeout elapses or context is done.
func (x FluentAsync[TFn, TTimerPayload, TTickPayload]) EventuallyContext(ctx context.Context) FailureMessage {
	return eventually(ctx, x.Timeout, x.Ticker, x.Fn)
}

// FluentPeriodic encapsulates asynchronous assertions using constant durations.
type FluentPeriodic[T ~func() FailureMessage] struct {
	Fn       T
	Timeout  time.Duration
	Interval time.Duration
}

// Periodic is used for asynchronous testing using time.Duriation for timeout and interval.
func Periodic[T ~func() FailureMessage](timeout, interval time.Duration, fn T) FluentPeriodic[T] {
	return FluentPeriodic[T]{fn, timeout, interval}
}

// Eventually executes the test function until it returns an empty FailureMessage
// or timeout elapses.
func (x FluentPeriodic[T]) Eventually() FailureMessage {
	ctx := context.Background()
	timer := time.NewTimer(x.Timeout)
	defer timer.Stop()
	ticker := time.NewTicker(x.Interval)
	defer ticker.Stop()
	return eventually(ctx, timer.C, ticker.C, x.Fn)
}

// EventuallyContext executes the test function until it returns an empty FailureMessage,
// timeout elapses or context is done.
func (x FluentPeriodic[T]) EventuallyContext(ctx context.Context) FailureMessage {
	timer := time.NewTimer(x.Timeout)
	defer timer.Stop()
	ticker := time.NewTicker(x.Interval)
	defer ticker.Stop()
	return eventually(ctx, timer.C, ticker.C, x.Fn)
}

// eventually executes the test function until it returns an empty FailureMessage,
// timeout elapses or context is done.
func eventually[TTimerPayload, TTickPayload any](ctx context.Context, timeout <-chan (TTimerPayload), ticker <-chan (TTickPayload), fn func() FailureMessage) FailureMessage {
	var (
		err  string
		tick <-chan TTickPayload
		ch   = make(chan FailureMessage, 1)
	)
	failMsg := func(cause string) FailureMessage {
		if err == "" {
			return FailureMessage(cause + "\nfunction has never returned")
		}
		return FailureMessage(cause + "\nfunction always failed\nlast failure message:\n" + err)
	}

	select {
	case <-ctx.Done():
		return failMsg(ctx.Err().Error())
	case <-timeout:
		return failMsg("timeout")
	default:
	}

	go func() { ch <- fn() }()
	for {
		select {
		case <-ctx.Done():
			return failMsg(ctx.Err().Error())
		case <-timeout:
			return failMsg("timeout")
		case <-tick:
			tick = nil
			go func() { ch <- fn() }()
		case msg := <-ch:
			if msg == "" {
				return ""
			}
			err = string(msg)
			tick = ticker
		}
	}
}
