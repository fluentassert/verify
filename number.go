package verify

import (
	"fmt"
	"math"

	"github.com/fluentassert/verify/constraints"
)

// NumberFloat encapsulates assertions for numbers.
type NumberFloat[T constraints.Number] struct {
	FluentOrdered[T]
}

// Number is used for testing numbers.
func Number[T constraints.Number](got T) NumberFloat[T] {
	return NumberFloat[T]{FluentOrdered[T]{FluentComparable[T]{FluentAny[T]{got}}}}
}

// InDelta tests that the numbers have an absolute error (distance) less or equal than delta.
func (x NumberFloat[T]) InDelta(want T, delta float64) FailureMessage {
	distance, msg := x.calcDistance(want, delta)
	if msg != "" {
		return msg
	}
	if distance < -delta || distance > delta {
		return FailureMessage(fmt.Sprintf("absolute error (distance) between numbers is greater than delta\nrelative error: %v\ndelta: %g\ngot: %v\nwant: %v", distance, delta, x.Got, want))
	}
	return ""
}

// NotInDelta tests that the numbers have an absolute error (distance) greater than delta.
func (x NumberFloat[T]) NotInDelta(want T, delta float64) FailureMessage {
	distance, msg := x.calcDistance(want, delta)
	if msg != "" {
		return msg
	}
	if distance < -delta || distance > delta {
		return ""
	}
	return FailureMessage(fmt.Sprintf("absolute error (distance) between numbers is lesser or equal than delta\nrelative error: %v\ndelta: %g\ngot: %v\nwant: %v", distance, delta, x.Got, want))
}

func (x NumberFloat[T]) calcDistance(want T, delta float64) (float64, FailureMessage) {
	if math.IsNaN(delta) || delta < 0 {
		return 0, "delta must be a non-negative number"
	}
	wantF := float64(want)
	gotF := float64(x.Got)
	if math.IsNaN(wantF) {
		return 0, "want is NaN"
	}
	if math.IsNaN(gotF) {
		return 0, "got is NaN"
	}
	return wantF - gotF, ""
}

// InEpsilon tests that the numbers have a relative error less or equal than epsilon.
func (x NumberFloat[T]) InEpsilon(want T, epsilon float64) FailureMessage {
	relativeError, msg := x.calcRelativeError(want, epsilon)
	if msg != "" {
		return msg
	}
	if relativeError > epsilon {
		return FailureMessage(fmt.Sprintf("relative error between numbers is greater than epsilon\nrelative error: %g\nepsilon: %g\ngot: %v\nwant: %v", relativeError, epsilon, x.Got, want))
	}
	return ""
}

// NotInEpsilon tests that the numbers have a relative error greater than epsilon.
func (x NumberFloat[T]) NotInEpsilon(want T, epsilon float64) FailureMessage {
	relativeError, msg := x.calcRelativeError(want, epsilon)
	if msg != "" {
		return msg
	}
	if relativeError > epsilon {
		return ""
	}
	return FailureMessage(fmt.Sprintf("relative error between numbers is lesser or equal than epsilon\nrelative error: %g\nepsilon: %g\ngot: %v\nto: %v", relativeError, epsilon, x.Got, want))
}

func (x NumberFloat[T]) calcRelativeError(want T, epsilon float64) (float64, FailureMessage) {
	if math.IsNaN(epsilon) || epsilon < 0 {
		return 0, "epsilon must be a non-negative number"
	}
	wantF := float64(want)
	gotF := float64(x.Got)
	if math.IsNaN(wantF) {
		return 0, "want is NaN"
	}
	if math.IsNaN(gotF) {
		return 0, "got is NaN"
	}
	if wantF == 0 {
		return 0, "want must have a value other than zero to calculate the relative error"
	}
	relativeError := math.Abs(wantF-gotF) / math.Abs(wantF)
	return relativeError, ""
}
