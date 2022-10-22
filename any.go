package verify

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// FluentAny encapsulates assertions for any object.
type FluentAny[T any] struct {
	Got T
}

// Any is used for testing any object.
func Any[T any](got T) FluentAny[T] {
	return FluentAny[T]{got}
}

// Check tests the object using the provided function.
func (x FluentAny[T]) Check(fn func(got T) FailureMessage) FailureMessage {
	return fn(x.Got)
}

// Should tests if the object meets the predicate criteria.
func (x FluentAny[T]) Should(pred func(got T) bool) FailureMessage {
	if pred(x.Got) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("object does not meet the predicate criteria\ngot: %+v", x.Got))
}

// ShouldNot tests if the object does not meet the predicate criteria.
func (x FluentAny[T]) ShouldNot(fn func(got T) bool) FailureMessage {
	if !fn(x.Got) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("object meets the predicate criteria\ngot: %+v", x.Got))
}

// DeepEqual tests if the objects are deep equal using github.com/google/go-cmp/cmp.
func (x FluentAny[T]) DeepEqual(want T, opts ...cmp.Option) FailureMessage {
	diff := cmp.Diff(want, x.Got, opts...)
	if diff == "" {
		return ""
	}
	return FailureMessage("mismatch (-want +got):\n" + diff)
}

// NotDeepEqual tests if the objects are not deep equal using github.com/google/go-cmp/cmp.
func (x FluentAny[T]) NotDeepEqual(obj T, opts ...cmp.Option) FailureMessage {
	ok := cmp.Equal(obj, x.Got, opts...)
	if !ok {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the objects are equal\ngot: %+v", x.Got))
}
