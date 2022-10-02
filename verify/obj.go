package verify

import "github.com/google/go-cmp/cmp"

// FluentObj encapsulates assertions for any object.
type FluentObj[T any] struct {
	Got T
}

// Obj is used for testing any object.
func Obj[T any](got T) FluentObj[T] {
	return FluentObj[T]{got}
}

// Check tests the object using the provided function.
func (x FluentObj[T]) Check(fn func(got T) FailureMessage) FailureMessage {
	return fn(x.Got)
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

// DeepEqual tests if the objects are deep equal using github.com/google/go-cmp/cmp.
func (x FluentObj[T]) DeepEqual(want T, opts ...cmp.Option) FailureMessage {
	diff := cmp.Diff(want, x.Got, opts...)
	if diff == "" {
		return ""
	}
	return FailureMessage("mismatch (-want +got):\n" + diff)
}

// NotDeepEqual tests if the objects are not deep equal using github.com/google/go-cmp/cmp.
func (x FluentObj[T]) NotDeepEqual(obj T, opts ...cmp.Option) FailureMessage {
	ok := cmp.Equal(obj, x.Got, opts...)
	if !ok {
		return ""
	}
	return FailureMessage("the objects are equal")
}

// Zero tests if the object is a zero value.
func (x FluentObj[T]) Zero() FailureMessage {
	var want T
	diff := cmp.Diff(want, x.Got)
	if diff == "" {
		return ""
	}
	return FailureMessage("not a zero value (-want +got):\n" + diff)
}

// NonZero tests if the object is a non-zero value.
func (x FluentObj[T]) NonZero() FailureMessage {
	var want T
	ok := cmp.Equal(want, x.Got)
	if !ok {
		return ""
	}
	return FailureMessage("a zero value")
}
