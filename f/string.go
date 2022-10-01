package f

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

// Contains tests if the string contains the substring.
func (x FluentString[T]) Contains(substr string) FailureMessage {
	if strings.Contains(string(x.Got), substr) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object does not contain the substring\nobject: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}

// NotContains tests if the string does not contain the substring.
func (x FluentString[T]) NotContains(substr string) FailureMessage {
	if !strings.Contains(string(x.Got), substr) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the object contains the substring\nobject: \"%s\"\nsubstr: \"%s\"", x.Got, substr))
}
