package verify

import (
	"fmt"
	"math"
)

// FluentNumber encapsulates assertions for numbers
// that supports the operators < <= >= > + - * /.
type FluentNumber[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64] struct {
	FluentOrdered[T]
}

// Number is used for testing numbers
// that supports the operators < <= >= > + - * /.
func Number[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64](got T) FluentNumber[T] {
	return FluentNumber[T]{FluentOrdered[T]{FluentObj[T]{FluentAny[T]{got}}}}
}

// InDelta tests that the numbers have an absolute error (distance) less or equal than delta.
func (x FluentNumber[T]) InDelta(want T, delta float64) FailureMessage {
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
func (x FluentNumber[T]) NotInDelta(want T, delta float64) FailureMessage {
	distance, msg := x.calcDistance(want, delta)
	if msg != "" {
		return msg
	}
	if distance < -delta || distance > delta {
		return ""
	}
	return FailureMessage(fmt.Sprintf("absolute error (distance) between numbers is lesser or equal than delta\nrelative error: %v\ndelta: %g\ngot: %v\nwant: %v", distance, delta, x.Got, want))
}

func (x FluentNumber[T]) calcDistance(want T, delta float64) (float64, FailureMessage) {
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
func (x FluentNumber[T]) InEpsilon(want T, epsilon float64) FailureMessage {
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
func (x FluentNumber[T]) NotInEpsilon(want T, epsilon float64) FailureMessage {
	relativeError, msg := x.calcRelativeError(want, epsilon)
	if msg != "" {
		return msg
	}
	if relativeError > epsilon {
		return ""
	}
	return FailureMessage(fmt.Sprintf("relative error between numbers is lesser or equal than epsilon\nrelative error: %g\nepsilon: %g\ngot: %v\nto: %v", relativeError, epsilon, x.Got, want))
}

func (x FluentNumber[T]) calcRelativeError(want T, epsilon float64) (float64, FailureMessage) {
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
