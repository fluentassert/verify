// Package f contains convenience assertion functions.
package f

import (
	"github.com/pellared/fluentassert/pred"
	"golang.org/x/exp/constraints"
)

// Assertion wraps a value to assert.
type Assertion[T any] struct {
	helper func()
	failf  func(format string, args ...interface{})
	got    T
}

// OrderedAssertion wraps an ordered value to assert.
type OrderedAssertion[T constraints.Ordered] struct {
	Assertion[T]
}

// Errorer reports error.
type Errorer interface {
	Errorf(format string, args ...interface{})
	Helper()
}

// Fataler reports fatal error.
type Fataler interface {
	Fatalf(format string, args ...interface{})
	Helper()
}

// Assert prepares an assertion which will t.Errorf if the predicate does not match.
func Assert[T any](t Errorer, got T) Assertion[T] {
	return Assertion[T]{t.Helper, t.Errorf, got}
}

// OrderedRequire prepares an assertion which will t.Fatalf if the predicate does not match.
func OrderedRequire[T constraints.Ordered](t Fataler, got T) OrderedAssertion[T] {
	return OrderedAssertion[T]{Assertion: Assertion[T]{t.Helper, t.Fatalf, got}}
}

// OrderedAssert prepares an assertion which will t.Errorf if the predicate does not match.
func OrderedAssert[T constraints.Ordered](t Errorer, got T) OrderedAssertion[T] {
	return OrderedAssertion[T]{Assertion: Assertion[T]{t.Helper, t.Errorf, got}}
}

// Require prepares an assertion which will t.Fatalf if the predicate does not match.
func Require[T any](t Fataler, got T) Assertion[T] {
	return Assertion[T]{t.Helper, t.Fatalf, got}
}

// Should checks the given predicate.
func (a Assertion[T]) Should(predicate func(got T) string, msg string, args ...interface{}) bool {
	a.helper()
	failMsg := predicate(a.got)
	if failMsg == "" {
		return true
	}
	a.failf(msg+"\n"+failMsg, args...)
	return false
}

// Eq checks if got is equal to want.
func (a Assertion[T]) Eq(want T, msg string, args ...interface{}) bool {
	a.helper()
	return a.Should(pred.Eq(want), msg, args...)
}

// Gt checks if got is greater to want.
func (a OrderedAssertion[T]) Gt(want T, msg string, args ...interface{}) bool {
	a.helper()
	return a.Should(pred.Gt(want), msg, args...)
}

// Err checks if got is an error.
func (a Assertion[T]) Err(msg string, args ...interface{}) bool {
	a.helper()
	return a.Should(pred.Err[T], msg, args...)
}

// Panic checks if got is a function that panics when executed.
func (a Assertion[T]) Panic(msg string, args ...interface{}) bool {
	a.helper()
	return a.Should(pred.Panic[T], msg, args...)
}

// NoPanic checks if got is a function that returns when executed.
func (a Assertion[T]) NoPanic(msg string, args ...interface{}) bool {
	a.helper()
	return a.Should(pred.NoPanic[T], msg, args...)
}
