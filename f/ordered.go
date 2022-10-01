package f

import "github.com/pellared/fluentassert/constraints"

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
	return "the object is not lesser"
}

// LesserOrEqual tests the objects using <= operator.
func (x FluentOrdered[T]) LesserOrEqual(than T) FailureMessage {
	if x.Got <= than {
		return ""
	}
	return "the object is not lesser or equal"
}

// GreaterOrEqual tests the objects using >= operator.
func (x FluentOrdered[T]) GreaterOrEqual(than T) FailureMessage {
	if x.Got >= than {
		return ""
	}
	return "the object is not greater or equal"
}

// Greater tests the objects using > operator.
func (x FluentOrdered[T]) Greater(than T) FailureMessage {
	if x.Got > than {
		return ""
	}
	return "the object is not greater"
}
