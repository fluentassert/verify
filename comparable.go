package verify

import "fmt"

// FluentObj encapsulates assertions for comparable object.
type FluentObj[T comparable] struct {
	FluentAny[T]
}

// Obj is used for testing a comparable object.
func Obj[T comparable](got T) FluentObj[T] {
	return FluentObj[T]{FluentAny[T]{got}}
}

// Equal tests the objects using == operator.
func (x FluentObj[T]) Equal(want T) FailureMessage {
	if x.Got == want {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the objects are not equal\ngot: %+v\nwant: %+v", x.Got, want))
}

// NotEqual tests the objects using != operator.
func (x FluentObj[T]) NotEqual(obj T) FailureMessage {
	if x.Got != obj {
		return ""
	}
	return "the objects are equal"
}

// Zero tests if the object is a zero value.
func (x FluentObj[T]) Zero() FailureMessage {
	var want T
	if want == x.Got {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not a zero value\ngot: %+v", x.Got))
}

// NonZero tests if the object is a non-zero value.
func (x FluentObj[T]) NonZero() FailureMessage {
	var want T
	if want != x.Got {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not a zero value\ngot: %+v", x.Got))
}
