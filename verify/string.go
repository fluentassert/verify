package verify

import (
	"fmt"
	"strings"
)

// FluentString encapsulates assertions for string object.
type FluentString[T ~string] struct {
	FluentOrdered[T]
}

// String is used for testing a string object.
func String[T ~string](got T) FluentString[T] {
	return FluentString[T]{FluentOrdered[T]{FluentComparable[T]{FluentObj[T]{got}}}}
}

// TODO: Add Empty

// TODO: Add NotEmpty

// Contain tests if the string contains the substring.
func (x FluentString[T]) Contain(substr string) FailureMessage {
	if strings.Contains(string(x.Got), substr) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object does not contain the substring\ngot: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}

// NotContain tests if the string does not contain the substring.
func (x FluentString[T]) NotContain(substr string) FailureMessage {
	if !strings.Contains(string(x.Got), substr) {
		return ""
	}

	return FailureMessage(fmt.Sprintf("the object contains the substring\ngot: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}

// TODO: Prefix

// TODO: Add NoPrefix

// TODO: Sufix

// TODO: Add NoSufix

// TODO: EqualFold

// TODO: NotEqualFold

// TODO: MatchRegex

// TODO: NotMatchRegex
