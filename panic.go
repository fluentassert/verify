package verify

import "fmt"

// Panics tests if the function panics when executed.
func Panics(fn func()) (msg FailureMessage) {
	defer func() {
		if r := recover(); r == nil {
			msg = "the function returned instead of panicking"
		}
	}()
	fn()
	return
}

// NotPanics tests if the function does not panic when executed.
func NotPanics(fn func()) (msg FailureMessage) {
	defer func() {
		if r := recover(); r != nil {
			msg = FailureMessage(fmt.Sprintf("the function panicked\ngot: %v", r))
		}
	}()
	fn()
	return
}
