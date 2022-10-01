package f

import "github.com/google/go-cmp/cmp"

// FluentObj encapsulates assertions for any objet.
type FluentObj[T any] struct {
	Got T
}

// Obj is used for testing any object.
func Obj[T any](got T) FluentObj[T] {
	return FluentObj[T]{got}
}

// Check tests the object using the provided function
func (x FluentObj[T]) Check(fn func(got T) string) FailureMessage {
	return FailureMessage(fn(x.Got))
}

// Should tests if the object meets the predicate criteria.
func (x FluentObj[T]) Should(pred func(got T) bool) FailureMessage {
	if pred(x.Got) {
		return ""
	}
	return "object does not meet the predicate criteria"
}

// ShouldNot tests if the object does not the predicate criteria.
func (x FluentObj[T]) ShouldNot(fn func(got T) bool) FailureMessage {
	if !fn(x.Got) {
		return ""
	}
	return "object meets the predicate criteria"
}

// DeepEqual tests of the objects are deep equal using github.com/google/go-cmp/cmp.
func (x FluentObj[T]) DeepEqual(want T, opts ...cmp.Option) FailureMessage {
	diff := cmp.Diff(want, x.Got, opts...)
	if diff == "" {
		return ""
	}
	return FailureMessage("mismatch (-want +got):\n" + diff)
}

// NotDeepEqual tests of the objects are not deep equal using github.com/google/go-cmp/cmp.
func (x FluentObj[T]) NotDeepEqual(want T, opts ...cmp.Option) FailureMessage {
	ok := cmp.Equal(want, x.Got, opts...)
	if !ok {
		return ""
	}
	return FailureMessage("the objects are equal")
}
