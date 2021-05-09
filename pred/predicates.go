// Package pred contains predicates used by f.Assertion.
package pred

import (
	"fmt"
	"reflect"
)

// Eq checks if got is equal to want.
func Eq(want interface{}) func(got interface{}) string {
	return func(got interface{}) string {
		if reflect.DeepEqual(got, want) {
			return ""
		}
		return fmt.Sprintf("got: %+v\nwant: %+v", got, want)
	}
}

// Err checks if got is an error.
func Err(got interface{}) string {
	if _, ok := got.(error); ok {
		return ""
	}
	return fmt.Sprintf("got: %+v\nwant an error", got)
}

// Panic checks if got is a function that panics when executed.
func Panic(got interface{}) (msg string) {
	fn, ok := got.(func())
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
func NoPanic(got interface{}) (msg string) {
	fn, ok := got.(func())
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
