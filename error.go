package verify

import (
	"errors"
	"fmt"
)

// NoError tests if the error is nil.
func NoError(err error) FailureMessage {
	if err == nil {
		return ""
	}
	return "non-nil error:\n" + FailureMessage(err.Error())
}

// IsError tests if the error is non-nil.
func IsError(err error) FailureMessage {
	if err != nil {
		return ""
	}
	return "the error is <nil>"
}

// FluentError encapsulates assertions for error object.
type FluentError struct {
	FluentAny[error]
	FluentString[string]
}

// Error is used for testing error object.
func Error(got error) FluentError {
	res := FluentError{FluentAny: FluentAny[error]{got}}
	if got != nil {
		res.FluentString.Got = got.Error()
	}
	return res
}

// Is tests whether any error in err's chain matches target.
func (x FluentError) Is(target error) FailureMessage {
	if errors.Is(x.Got, target) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("no error in err's chain matches\ngot: %#v\ntarget: %#v", x.Got, target))
}

// IsNot tests whether no error in err's chain matches target.
func (x FluentError) IsNot(target error) FailureMessage {
	if !errors.Is(x.Got, target) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("some error in err's chain matches\ngot: %#v\ntarget: %#v", x.Got, target))
}

// As finds the first error in err's chain that matches target, and if one is found, sets
// target to that error value. In such case it is a success..
func (x FluentError) As(target any) FailureMessage {
	if errors.As(x.Got, target) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("no error in err's chain matches\ngot: %#v\ntarget: %T", x.Got, target))
}

// AsNot finds the first error in err's chain that matches target, and if one is found, sets
// target to that error value. In such case it is a failure.
func (x FluentError) AsNot(target any) FailureMessage {
	if !errors.As(x.Got, target) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("some error in err's chain matches\ngot: %#v\ntarget: %T", x.Got, target))
}
