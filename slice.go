package verify

// FluentSlice encapsulates assertions for a slice.
type FluentSlice[T any] struct {
	FluentAny[[]T]
}

// Slice is used for testing a slice.
func Slice[T any](got []T) FluentSlice[T] {
	return FluentSlice[T]{FluentAny[[]T]{got}}
}

// TODO: Empty() FailureMessage

// TODO: NotEmpty() FailureMessage

// TODO: Len(n int) FailureMessage

// TODO: Equal(elements []T) FailureMessage

// TODO: NotEqual(elements []T) FailureMessage

// TODO: Equivalent(elements []T) FailureMessage

// TODO: NotEquivalent(elements []T) FailureMessage

// TODO: Contain(elements ...T) FailureMessage

// TODO: NotContain(elements ...T) FailureMessage

// TODO: Any(func(T) bool) FailureMessage

// TODO: All(func(T) bool) FailureMessage
