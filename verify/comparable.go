package verify

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
	return "the objects are not equal"
}

// NotEqual tests the objects using != operator.
func (x FluentComparable[T]) NotEqual(obj T) FailureMessage {
	if x.Got != obj {
		return ""
	}
	return "the objects are equal"
}
