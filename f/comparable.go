package f

// FluentComparable encapsulates assertions for comparable object.
type FluentComparable[T comparable] struct {
	FluentObj[T]
}

// Comparable is used for testing a comparable object.
func Comparable[T comparable](got T) FluentComparable[T] {
	return FluentComparable[T]{FluentObj[T]{got}}
}

// Equal tests if the objects are equal using == operator.
func (x FluentComparable[T]) Equal(want T) FailureMessage {
	if x.Got == want {
		return ""
	}
	return "the objects are not equal"
}

// NotEqual tests if the objects are equal using == operator.
func (x FluentComparable[T]) NotEqual(want T) FailureMessage {
	if x.Got != want {
		return ""
	}
	return "the objects are equal"
}
