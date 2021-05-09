// Package f contains convenience assertion functions.
package f

import (
	"github.com/pellared/fluentassert/pred"
)

// Assertion wraps a value to assert.
type Assertion struct {
	h     helperer
	failf func(format string, args ...interface{})
	got   interface{}
}

// Errorer reports error.
type Errorer interface {
	Errorf(format string, args ...interface{})
}

// Fataler reports fatal error.
type Fataler interface {
	Fatalf(format string, args ...interface{})
}

// Assert prepares an assertion which will t.Errorf if the predicate does not match.
func Assert(t Errorer, got interface{}) Assertion {
	h, _ := t.(helperer)
	return Assertion{h, t.Errorf, got}
}

// Require prepares an assertion which will t.Fatalf if the predicate does not match.
func Require(t Fataler, got interface{}) Assertion {
	h, _ := t.(helperer)
	return Assertion{h, t.Fatalf, got}
}

// Helper marks the calling function as a test helper function.
// When printing file and line information, that function will be skipped.
// Helper may be called simultaneously from multiple goroutines.
func (a Assertion) Helper() {
	if a.h != nil {
		a.h.Helper()
	}
}

// Should checks the given predicate.
func (a Assertion) Should(predicate func(got interface{}) string, msg string, args ...interface{}) bool {
	a.Helper()
	failMsg := predicate(a.got)
	if failMsg == "" {
		return true
	}
	a.failf(msg+"\n"+failMsg, args...)
	return false
}

// Eq checks if got is equal to want.
func (a Assertion) Eq(want interface{}, msg string, args ...interface{}) bool {
	a.Helper()
	return a.Should(pred.Eq(want), msg, args...)
}

// Nil checks if got is nil.
func (a Assertion) Nil(msg string, args ...interface{}) bool {
	a.Helper()
	return a.Should(pred.Eq(nil), msg, args...)
}

// Err checks if got is an error.
func (a Assertion) Err(msg string, args ...interface{}) bool {
	a.Helper()
	return a.Should(pred.Err, msg, args...)
}

// Panic checks if got is a function that panics when executed.
func (a Assertion) Panic(msg string, args ...interface{}) bool {
	a.Helper()
	return a.Should(pred.Panic, msg, args...)
}

// NoPanic checks if got is a function that returns when executed.
func (a Assertion) NoPanic(msg string, args ...interface{}) bool {
	a.Helper()
	return a.Should(pred.NoPanic, msg, args...)
}

type helperer interface {
	Helper()
}
