package verify

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// FluentSlice encapsulates assertions for a slice.
type FluentSlice[T any] struct {
	FluentAny[[]T]
}

// Slice is used for testing a slice.
func Slice[T any](got []T) FluentSlice[T] {
	return FluentSlice[T]{FluentAny[[]T]{got}}
}

// Empty tests if the slice is empty.
func (x FluentSlice[T]) Empty() FailureMessage {
	if len(x.Got) == 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not an empty slice\ngot: %+v", x.Got))
}

// NotEmpty tests if the slice is not empty.
func (x FluentSlice[T]) NotEmpty() FailureMessage {
	if len(x.Got) > 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("an empty slice\ngot: %+v", x.Got))
}

// Equivalent tests if the slice has the same items as want in any order.
func (x FluentSlice[T]) Equivalent(want []T, opts ...cmp.Option) FailureMessage {
	extraGot, extraWant := x.diff(want, opts)
	if len(extraGot) == 0 && len(extraWant) == 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not equivalent\ngot: %+v\nwant: %+v\nextra got: %+v\nextra want: %+v", x.Got, want, extraGot, extraWant))
}

// NotEquivalent tests if the slice does not have the same items as want in any order.
func (x FluentSlice[T]) NotEquivalent(want []T, opts ...cmp.Option) FailureMessage {
	extraGot, extraWant := x.diff(want, opts)
	if len(extraGot) != 0 || len(extraWant) != 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("equivalent\ngot: %+v", extraGot))
}

func (x FluentSlice[T]) diff(want []T, opts []cmp.Option) (extraGot, extraWant []T) {
	aLen := len(x.Got)
	bLen := len(want)

	// Mark indexes in list that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if cmp.Equal(want[j], x.Got[i], opts...) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraGot = append(extraGot, x.Got[i])
		}
	}

	for j := 0; j < bLen; j++ {
		if visited[j] {
			continue
		}
		extraWant = append(extraWant, want[j])
	}

	return
}

// Contain tests if the slice contains the item.
func (x FluentSlice[T]) Contain(item T, opts ...cmp.Option) FailureMessage {
	for _, v := range x.Got {
		if cmp.Equal(item, v, opts...) {
			return ""
		}
	}
	return FailureMessage(fmt.Sprintf("slice does not contain the item\ngot: %+v\nitem: %+v", x.Got, item))
}

// NotContain tests if the slice does not contain the item.
func (x FluentSlice[T]) NotContain(item T, opts ...cmp.Option) FailureMessage {
	for _, v := range x.Got {
		if cmp.Equal(item, v, opts...) {
			return FailureMessage(fmt.Sprintf("slice contains the item\ngot: %+v\nitem: %+v", x.Got, item))
		}
	}
	return ""
}

// Any tests if any of the slice's items meets the predicate criteria.
func (x FluentSlice[T]) Any(predicate func(T) bool) FailureMessage {
	for _, v := range x.Got {
		if predicate(v) {
			return ""
		}
	}
	return FailureMessage(fmt.Sprintf("none item does meet the predicate criteria\ngot: %+v", x.Got))
}

// All tests if all of the slice's items meet the predicate criteria.
func (x FluentSlice[T]) All(predicate func(T) bool) FailureMessage {
	for _, v := range x.Got {
		if !predicate(v) {
			return FailureMessage(fmt.Sprintf("an item does not meet the predicate criteria\ngot: %+v\nitem: %+v", x.Got, v))
		}
	}
	return ""
}

// None tests if none of the slice's items meets the predicate criteria.
func (x FluentSlice[T]) None(predicate func(T) bool) FailureMessage {
	for _, v := range x.Got {
		if predicate(v) {
			return FailureMessage(fmt.Sprintf("an item meets the predicate criteria\ngot: %+v\nitem: %+v", x.Got, v))
		}
	}
	return ""
}

// Len tests the length of the slice.
func (x FluentSlice[T]) Len(want int) FailureMessage {
	gotLen := len(x.Got)
	if gotLen != want {
		return FailureMessage(fmt.Sprintf("has different length\ngot: %+v\nlen: %v\nwant: %v", x.Got, gotLen, want))
	}
	return ""
}
