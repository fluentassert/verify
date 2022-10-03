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

// Empty tests if the string is not empty.
func (x FluentString[T]) Empty() FailureMessage {
	if x.Got == "" {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value was not an empty string\ngot: \"%s\"", x.Got))
}

// NotEmpty tests if the string is not empty.
func (x FluentString[T]) NotEmpty() FailureMessage {
	if x.Got != "" {
		return ""
	}
	return "the value was \"\""
}

// Contain tests if the string contains the substring.
func (x FluentString[T]) Contain(substr string) FailureMessage {
	if strings.Contains(string(x.Got), substr) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value does not contain the substring\ngot: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}

// NotContain tests if the string does not contain the substring.
func (x FluentString[T]) NotContain(substr string) FailureMessage {
	if !strings.Contains(string(x.Got), substr) {
		return ""
	}

	return FailureMessage(fmt.Sprintf("the value contains the substring\ngot: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}

// Prefix tests if the string starts with the prefix.
func (x FluentString[T]) Prefix(prefix string) FailureMessage {
	if strings.HasPrefix(string(x.Got), prefix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value does not have the prefix\ngot: \"%s\"\nprefix: \"%s\"", x.Got, prefix))
}

// NoPrefix tests if the string does not start with the prefix.
func (x FluentString[T]) NoPrefix(prefix string) FailureMessage {
	if !strings.HasPrefix(string(x.Got), prefix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value has the prefix\ngot: \"%s\"\nprefix: \"%s\"", x.Got, prefix))
}

// Sufix tests if the string ends with the sufix.
func (x FluentString[T]) Sufix(sufix string) FailureMessage {
	if strings.HasSuffix(string(x.Got), sufix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value does not have the sufix\ngot: \"%s\"\nsufix: \"%s\"", x.Got, sufix))
}

// NoSufix tests if the string does not end with the sufix.
func (x FluentString[T]) NoSufix(sufix string) FailureMessage {
	if !strings.HasSuffix(string(x.Got), sufix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value has the sufix\ngot: \"%s\"\nsufix: \"%s\"", x.Got, sufix))
}

// TODO: EqualFold

// TODO: NotEqualFold

// TODO: MatchRegex

// TODO: NotMatchRegex
