package verify

import (
	"fmt"

	"github.com/fluentassert/verify/constraints"
)

// FluentOrdered encapsulates assertions for ordered object
// // that supports the operators < <= >= >.
type FluentOrdered[T constraints.Ordered] struct {
	FluentComparable[T]
}

// Ordered is used for testing a ordered object
// that supports the operators < <= >= >.
func Ordered[T constraints.Ordered](got T) FluentOrdered[T] {
	return FluentOrdered[T]{FluentComparable[T]{FluentObj[T]{got}}}
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
