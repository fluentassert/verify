package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestComparable(t *testing.T) {
	type A struct {
		Str  string
		Bool bool
	}

	t.Run("Equal", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true}
			got := A{Str: "string", Bool: true}
			msg := verify.Comparable(got).Equal(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true}
			got := A{Str: "wrong", Bool: true}
			msg := verify.Comparable(got).Equal(want)
			assertFailed(t, msg, "the objects are not equal")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Comparable(got).Equal(nil)
			assertPassed(t, msg)
		})
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true}
			got := A{Str: "wrong", Bool: true}
			msg := verify.Comparable(got).NotEqual(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true}
			got := A{Str: "string", Bool: true}
			msg := verify.Comparable(got).NotEqual(want)
			assertFailed(t, msg, "the objects are equal")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Comparable(got).NotEqual(nil)
			assertFailed(t, msg, "the objects are equal")
		})
	})

	t.Run("Zero", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := A{}
			msg := verify.Comparable(got).Zero()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			got := A{Str: "wrong"}
			msg := verify.Comparable(got).Zero()
			assertFailed(t, msg, "not a zero value\n")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Comparable(got).Zero()
			assertPassed(t, msg)
		})
	})

	t.Run("NonZero", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := A{Str: "string"}
			msg := verify.Comparable(got).NonZero()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			got := A{}
			msg := verify.Comparable(got).NonZero()
			assertFailed(t, msg, "a zero value")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Comparable(got).NonZero()
			assertFailed(t, msg, "a zero value")
		})
	})

	t.Run("has assertions from Obj", func(t *testing.T) {
		want := A{}
		got := verify.Comparable(want).FluentObj.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}
