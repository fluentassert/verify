package verify_test

import (
	"regexp"
	"testing"

	"github.com/fluentassert/verify"
)

func TestString(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("").Empty()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("val").Empty()
			assertFailed(t, msg, "the value was not an empty string")
		})
	})

	t.Run("NotEmpty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("val").NotEmpty()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("").NotEmpty()
			assertFailed(t, msg, "the value was \"\"")
		})
	})

	t.Run("Contains", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("text").Contain("ex")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("text").Contain("asd")
			assertFailed(t, msg, "the value does not contain the substring")
		})
	})

	t.Run("NotContains", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("text").NotContain("asd")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("text").NotContain("ex")
			assertFailed(t, msg, "the value contains the substring")
		})
	})

	t.Run("Prefix", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("[ok]a").Prefix("[ok]")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("a[ok]").Prefix("[ok]")
			assertFailed(t, msg, "the value does not have the prefix")
		})
	})

	t.Run("NoPrefix", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("a[ok]").NoPrefix("[ok]")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("[ok]a").NoPrefix("[ok]")
			assertFailed(t, msg, "the value has the prefix")
		})
	})

	t.Run("Sufix", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("a[ok]").Sufix("[ok]")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("[ok]a").Sufix("[ok]")
			assertFailed(t, msg, "the value does not have the sufix")
		})
	})

	t.Run("NoSufix", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("[ok]a").NoSufix("[ok]")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("a[ok]").NoSufix("[ok]")
			assertFailed(t, msg, "the value has the sufix")
		})
	})

	t.Run("EqualFold", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("aBc").EqualFold("ABC")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("aBc").EqualFold("aB")
			assertFailed(t, msg, "the string values are not equal under Unicode case folding")
		})
	})

	t.Run("NotEqualFold", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("aBc").NotEqualFold("aB")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("aBc").NotEqualFold("ABC")
			assertFailed(t, msg, "the string values are equal under Unicode case folding")
		})
	})

	t.Run("MatchRegex", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("3aD").MatchRegex(regexp.MustCompile("[0-9][a-z][A-Z]"))
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("123").MatchRegex(regexp.MustCompile("[0-9][a-z][A-Z]"))
			assertFailed(t, msg, "the string value does not match the regular expression")
		})
	})

	t.Run("NotMatchRegex", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := verify.String("123").NotMatchRegex(regexp.MustCompile("[0-9][a-z][A-Z]"))
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := verify.String("3aD").NotMatchRegex(regexp.MustCompile("[0-9][a-z][A-Z]"))
			assertFailed(t, msg, "the string value matches the regular expression")
		})
	})

	t.Run("has assertions from Ordered, Comparable, Obj", func(t *testing.T) {
		want := "text"
		got := verify.String(want).FluentOrdered.FluentComparable.FluentAny.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}
