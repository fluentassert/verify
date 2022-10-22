package verify

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

// FluentMap encapsulates assertions for a map.
type FluentMap[K comparable, V any] struct {
	FluentAny[map[K]V]
}

// Map is used for testing a map.
func Map[K comparable, V any](got map[K]V) FluentMap[K, V] {
	return FluentMap[K, V]{FluentAny[map[K]V]{got}}
}

// Empty tests if the slice is empty.
func (x FluentMap[K, V]) Empty() FailureMessage {
	if len(x.Got) == 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not an empty map\ngot: %+v", x.Got))
}

// NotEmpty tests if the slice is not empty.
func (x FluentMap[K, V]) NotEmpty() FailureMessage {
	if len(x.Got) > 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("an empty map\ngot: %+v", x.Got))
}

// Contain tests if the map contains all pairs from want.
func (x FluentMap[K, V]) Contain(want map[K]V, opts ...cmp.Option) FailureMessage {
	missing := x.miss(want, opts)
	if len(missing) == 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("not contains all pairs\ngot: %+v\nwant: %+v\nmissing: %+v", x.Got, want, missing))
}

// NotContain tests if the map does not contains all pairs from want.
func (x FluentMap[K, V]) NotContain(want map[K]V, opts ...cmp.Option) FailureMessage {
	missing := x.miss(want, opts)
	if len(missing) > 0 {
		return ""
	}
	return FailureMessage(fmt.Sprintf("contains all pairs\ngot: %+v\nwant: %+v", x.Got, want))
}

func (x FluentMap[K, V]) miss(want map[K]V, opts []cmp.Option) map[K]V {
	missing := map[K]V{}
	for k, v := range want {
		got, ok := x.Got[k]
		if !ok {
			missing[k] = v
			continue
		}
		if !cmp.Equal(v, got, opts...) {
			missing[k] = v
			continue
		}
	}
	return missing
}

// ContainPair tests if the map contains the pair.
func (x FluentMap[K, V]) ContainPair(k K, v V, opts ...cmp.Option) FailureMessage {
	got, ok := x.Got[k]
	if !ok {
		return FailureMessage(fmt.Sprintf("has no value under key\ngot: %+v\nkey: %+v\nvalue: %+v", x.Got, k, v))
	}
	if !cmp.Equal(v, got, opts...) {
		return FailureMessage(fmt.Sprintf("has different value under key\nkey: %+v\ngot: %+v\nwant: %+v", k, got, v))
	}
	return ""
}

// NotContainPair tests if the map does not contain the pair.
func (x FluentMap[K, V]) NotContainPair(k K, v V, opts ...cmp.Option) FailureMessage {
	got, ok := x.Got[k]
	if !ok {
		return ""
	}
	if !cmp.Equal(v, got, opts...) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("contains the pair\ngot: %+v\nkey: %+v\nvalue: %+v", x.Got, k, v))
}

// Any tests if any of the slice's item meets the predicate criteria.
func (x FluentMap[K, V]) Any(predicate func(K, V) bool) FailureMessage {
	for k, v := range x.Got {
		if predicate(k, v) {
			return ""
		}
	}
	return FailureMessage(fmt.Sprintf("none pair does meet the predicate criteria\ngot: %+v", x.Got))
}

// All tests if all of the slice's items meets the predicate criteria.
func (x FluentMap[K, V]) All(predicate func(K, V) bool) FailureMessage {
	for k, v := range x.Got {
		if !predicate(k, v) {
			return FailureMessage(fmt.Sprintf("a pair does not meet the predicate criteria\ngot: %+v\npair: %+v", x.Got, v))
		}
	}
	return ""
}

// None tests if all of the slice's item does not meet the predicate criteria.
func (x FluentMap[K, V]) None(predicate func(K, V) bool) FailureMessage {
	for k, v := range x.Got {
		if predicate(k, v) {
			return FailureMessage(fmt.Sprintf("a pair meets the predicate criteria\ngot: %+v\npair: %+v", x.Got, v))
		}
	}
	return ""
}
