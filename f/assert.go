// Package f contains convenience assertion functions.
package f

import (
	"fmt"
	"reflect"
	"testing"
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

// Eq checks if the value is equal to want value.
func (a Assertion) Eq(want interface{}, msg string, args ...interface{}) bool {
	a.T.Helper()
	return a.Should(eq(want), msg, args...)
}

func eq(want interface{}) func(got interface{}) string {
	return func(got interface{}) string {
		if reflect.DeepEqual(got, want) {
			return ""
		}
		return fmt.Sprintf("got: %+v\nwant: %+v", got, want)
	}
}
