package verify

import (
	"fmt"
)

// FluentOrdered encapsulates assertions for ordered object
// that supports the operators < <= >= >.
type FluentOrdered[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string] struct {
	FluentObj[T]
}

// Ordered is used for testing a ordered object
// that supports the operators < <= >= >.
func Ordered[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string](got T) FluentOrdered[T] {
	return FluentOrdered[T]{FluentObj[T]{FluentAny[T]{got}}}
}

// Lesser tests the objects using < operator.
func (x FluentOrdered[T]) Lesser(than T) FailureMessage {
	if x.Got < than {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object is not lesser\ngot: %v\nthan: %v", x.Got, than))
}

// LesserOrEqual tests the objects using <= operator.
func (x FluentOrdered[T]) LesserOrEqual(than T) FailureMessage {
	if x.Got <= than {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object is not lesser or equal\ngot: %v\nthan: %v", x.Got, than))
}

// GreaterOrEqual tests the objects using >= operator.
func (x FluentOrdered[T]) GreaterOrEqual(than T) FailureMessage {
	if x.Got >= than {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object is not greater or equal\ngot: %v\nthan: %v", x.Got, than))
}

// Greater tests the objects using > operator.
func (x FluentOrdered[T]) Greater(than T) FailureMessage {
	if x.Got > than {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object is not greater\ngot: %v\nthan: %v", x.Got, than))
}
