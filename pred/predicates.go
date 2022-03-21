// Package pred contains predicates used by f.Assertion.
package pred

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Eq checks if got is equal to want.
func Eq[T any](want T) func(got T) string {
	return func(got T) string {
		if reflect.DeepEqual(got, want) {
			return ""
		}
		return fmt.Sprintf("got: %+v\nwant: %+v", got, want)
	}
}

// Gt checks if got is equal to want.
func Gt[T constraints.Ordered](want T) func(got T) string {
	return func(got T) string {
		if got > want {
			return ""
		}
		return fmt.Sprintf("got: %+v\nwant greater than: %+v", got, want)
	}
}

// Err checks if got is an error.
func Err[T any](got T) string {
	var v interface{} = got
	if _, ok := v.(error); ok {
		return ""
	}
	return fmt.Sprintf("got: %+v\nwant an error", got)
}

// Panic checks if got is a function that panics when executed.
func Panic[T any](got T) (msg string) {
	var v interface{} = got
	fn, ok := v.(func())
	if !ok {
		return "got: should be a func()"
	}
	defer func() {
		if r := recover(); r == nil {
			msg = "got: returned\nwant a panic"
		}
	}()
	fn()
	return
}

// NoPanic checks if got is a function that returns when executed.
func NoPanic[T any](got T) (msg string) {
	var v interface{} = got
	fn, ok := v.(func())
	if !ok {
		return "got: should be a func()"
	}
	defer func() {
		if r := recover(); r != nil {
			msg = "got: panicked\nwant to return"
		}
	}()
	fn()
	return
}
