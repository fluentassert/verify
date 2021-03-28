// Package f contains convenience assertion functions.
package f

import (
	"testing"

	"github.com/pellared/fluentassert/pred"
)

// Assertion wraps a value to assert.
type Assertion struct {
	T     testing.TB
	Failf func(format string, args ...interface{})
	Got   interface{}
}

// Assert prepares an assertion which will t.Errorf if the predicate does not match.
func Assert(t testing.TB, got interface{}) Assertion {
	return Assertion{t, t.Errorf, got}
}

// Require prepares an assertion which will t.Fatalf if the predicate does not match.
func Require(t testing.TB, got interface{}) Assertion {
	return Assertion{t, t.Fatalf, got}
}

// Should checks the given predicate.
func (a Assertion) Should(predicate func(got interface{}) string, msg string, args ...interface{}) bool {
	a.T.Helper()
	failMsg := predicate(a.Got)
	if failMsg == "" {
		return true
	}
	a.Failf(msg+"\n"+failMsg, args...)
	return false
}

// Eq checks if got is equal to want.
func (a Assertion) Eq(want interface{}, msg string, args ...interface{}) bool {
	a.T.Helper()
	return a.Should(pred.Eq(want), msg, args...)
}

// Nil checks if got is nil.
func (a Assertion) Nil(msg string, args ...interface{}) bool {
	a.T.Helper()
	return a.Should(pred.Eq(nil), msg, args...)
}

// Err checks if got is an error.
func (a Assertion) Err(msg string, args ...interface{}) bool {
	a.T.Helper()
	return a.Should(pred.Err, msg, args...)
}
