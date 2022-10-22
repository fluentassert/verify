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

// NotContain tests if the slice does not have the same element as want in any order.
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

// TODO: Contain(elements map[K]V) FailureMessage

// TODO: NotContain(elements map[K]V) FailureMessage

// TODO: ContainKey(keys ...K) FailureMessage

// TODO: NotContainKey(keys ...K) FailureMessage

// TODO: ContainPair(K,V) FailureMessage

// TODO: NotContainPair(K,V) FailureMessage

// TODO: Any(func(K,V) bool) FailureMessage

// TODO: All(func(K,V) bool) FailureMessage

// TODO: None(func(K,V) bool) FailureMessage
