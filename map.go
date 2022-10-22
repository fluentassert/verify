package verify

// FluentMap encapsulates assertions for a map.
type FluentMap[K comparable, V any] struct {
	FluentAny[map[K]V]
}

// Map is used for testing a map.
func Map[K comparable, V any](got map[K]V) FluentMap[K, V] {
	return FluentMap[K, V]{FluentAny[map[K]V]{got}}
}

// TODO: Empty() FailureMessage

// TODO: NotEmpty() FailureMessage

// TODO: Len(n int) FailureMessage

// TODO: Equal(elements map[K]V) FailureMessage

// TODO: NotEqual(elements map[K]V) FailureMessage

// TODO: Contain(elements map[K]V) FailureMessage

// TODO: NotContain(elements map[K]V) FailureMessage

// TODO: ContainKey(keys ...K) FailureMessage

// TODO: NotContainKey(keys ...K) FailureMessage

// TODO: ContainPair(K,V) FailureMessage

// TODO: NotContainPair(K,V) FailureMessage

// TODO: Any(func(K,V) bool) FailureMessage

// TODO: All(func(K,V) bool) FailureMessage
