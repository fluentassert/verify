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

// Should uses a function to test the object.
func (x FluentObj[T]) Should(fn func(got T) string) FailureMessage {
	return FailureMessage(fn(x.Got))
}

// DeepEq uses github.com/google/go-cmp/cmp for comparing objects.
func (x FluentObj[T]) DeepEq(want T) FailureMessage {
	diff := cmp.Diff(want, x.Got)
	if diff == "" {
		return ""
	}
	return FailureMessage("mismatch (-want +got):\n" + diff)
}
