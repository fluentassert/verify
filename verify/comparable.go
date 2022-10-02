package verify

import "fmt"

// FluentComparable encapsulates assertions for comparable object.
type FluentComparable[T comparable] struct {
	FluentObj[T]
}

// Comparable is used for testing a comparable object.
func Comparable[T comparable](got T) FluentComparable[T] {
	return FluentComparable[T]{FluentObj[T]{got}}
}

// Equal tests the objects using == operator.
func (x FluentComparable[T]) Equal(want T) FailureMessage {
	if x.Got == want {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the objects are not equal\ngot: %#v\nwant: %#v", x.Got, want))
}

// NotEqual tests the objects using != operator.
func (x FluentComparable[T]) NotEqual(obj T) FailureMessage {
	if x.Got != obj {
		return ""
	}
	return "the objects are equal"
}
