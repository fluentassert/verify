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

// Zero tests if the object is a zero value.
func (x FluentComparable[T]) Zero() FailureMessage {
	var want T
	if want == x.Got {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not a zero value\ngot: %#v", x.Got))
}

// NonZero tests if the object is a non-zero value.
func (x FluentComparable[T]) NonZero() FailureMessage {
	var want T
	if want != x.Got {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not a zero value\ngot: %#v", x.Got))
}
