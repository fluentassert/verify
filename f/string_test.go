package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestString(t *testing.T) {
	t.Run("Contains", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := f.String("text").Contains("ex")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := f.String("text").Contains("asd")
			assertFailed(t, msg, "the object does not contain the substring")
		})
	})

	t.Run("NotContains", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			msg := f.String("text").NotContains("asd")
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			msg := f.String("text").NotContains("ex")
			assertFailed(t, msg, "the object contains the substring")
		})
	})

	t.Run("has assertions from Ordered, Comparable, Obj", func(t *testing.T) {
		want := "text"
		got := f.String(want).FluentOrdered.FluentComparable.FluentObj.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}
