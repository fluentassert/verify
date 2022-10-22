package verify

import "fmt"

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

// TODO: Contain(elements map[K]V) FailureMessage

// TODO: NotContain(elements map[K]V) FailureMessage

// TODO: ContainKey(keys ...K) FailureMessage

// TODO: NotContainKey(keys ...K) FailureMessage

// TODO: ContainPair(K,V) FailureMessage

// TODO: NotContainPair(K,V) FailureMessage

// TODO: Any(func(K,V) bool) FailureMessage

// TODO: All(func(K,V) bool) FailureMessage

// TODO: None(func(K,V) bool) FailureMessage
