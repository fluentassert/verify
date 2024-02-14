package verify

import (
	"fmt"
	"regexp"
	"strings"
)

// FluentString encapsulates assertions for string object.
type FluentString[T ~string] struct {
	FluentOrdered[T]
}

// String is used for testing a string object.
func String[T ~string](got T) FluentString[T] {
	return FluentString[T]{FluentOrdered[T]{FluentObj[T]{FluentAny[T]{got}}}}
}

// Empty tests if the string is not empty.
func (x FluentString[T]) Empty() FailureMessage {
	if x.Got == "" {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value was not an empty string\ngot: %q", x.Got))
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
	return FailureMessage(fmt.Sprintf("the value does not contain the substring\ngot: %q\nsubstr: %q", x.Got, substr))
}

// NotContain tests if the string does not contain the substring.
func (x FluentString[T]) NotContain(substr string) FailureMessage {
	if !strings.Contains(string(x.Got), substr) {
		return ""
	}

	return FailureMessage(fmt.Sprintf("the value contains the substring\ngot: %q\nsubstr: %q", x.Got, substr))
}

// Prefix tests if the string starts with the prefix.
func (x FluentString[T]) Prefix(prefix string) FailureMessage {
	if strings.HasPrefix(string(x.Got), prefix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value does not have the prefix\ngot: %q\nprefix: %q", x.Got, prefix))
}

// NoPrefix tests if the string does not start with the prefix.
func (x FluentString[T]) NoPrefix(prefix string) FailureMessage {
	if !strings.HasPrefix(string(x.Got), prefix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value has the prefix\ngot: %q\nprefix: %q", x.Got, prefix))
}

// Sufix tests if the string ends with the sufix.
func (x FluentString[T]) Sufix(sufix string) FailureMessage {
	if strings.HasSuffix(string(x.Got), sufix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value does not have the sufix\ngot: %q\nsufix: %q", x.Got, sufix))
}

// NoSufix tests if the string does not end with the sufix.
func (x FluentString[T]) NoSufix(sufix string) FailureMessage {
	if !strings.HasSuffix(string(x.Got), sufix) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the value has the sufix\ngot: %q\nsufix: %q", x.Got, sufix))
}

// EqualFold tests if the values interpreted as UTF-8 strings,
// are equal under simple Unicode case-folding,
// which is a more general form of case-insensitivity.
func (x FluentString[T]) EqualFold(want string) FailureMessage {
	if strings.EqualFold(string(x.Got), want) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the string values are not equal under Unicode case folding\ngot: %q\nwant: %q", x.Got, want))
}

// NotEqualFold tests if the values interpreted as UTF-8 strings,
// are not equal under simple Unicode case-folding,
// which is a more general form of case-insensitivity.
func (x FluentString[T]) NotEqualFold(want string) FailureMessage {
	if !strings.EqualFold(string(x.Got), want) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the string values are equal under Unicode case folding\ngot: %q\nwant: %q", x.Got, want))
}

// MatchRegex tests if the string matches the regular expression.
func (x FluentString[T]) MatchRegex(regex *regexp.Regexp) FailureMessage {
	if regex.MatchString(string(x.Got)) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the string value does not match the regular expression\ngot: %q\nregex: %s", x.Got, regex.String()))
}

// NotMatchRegex tests if the string does not match the regular expression.
func (x FluentString[T]) NotMatchRegex(regex *regexp.Regexp) FailureMessage {
	if !regex.MatchString(string(x.Got)) {
		return ""
	}
	return FailureMessage(fmt.Sprintf("the string value matches the regular expression\ngot: %q\nregex: %s", x.Got, regex.String()))
}

// Len tests the length of the string.
func (x FluentString[T]) Len(want int) FailureMessage {
	gotLen := len(x.Got)
	if gotLen != want {
		return FailureMessage(fmt.Sprintf("has different length\ngot: %+v\nlen: %v\nwant: %v", x.Got, gotLen, want))
	}
	return ""
}
