package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestAny(t *testing.T) {
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	t.Run("DeepEqual", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := verify.Any(got).DeepEqual(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := verify.Any(got).DeepEqual(want)
			assertFailed(t, msg, "mismatch (-want +got):\n")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Any(got).DeepEqual(nil)
			assertPassed(t, msg)
		})
	})

	t.Run("NotDeepEqual", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := verify.Any(got).NotDeepEqual(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := verify.Any(got).NotDeepEqual(want)
			assertFailed(t, msg, "the objects are equal")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Any(got).NotDeepEqual(nil)
			assertFailed(t, msg, "the objects are equal")
		})
	})

	t.Run("Check", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			fn := func(x A) verify.FailureMessage {
				return ""
			}
			msg := verify.Any(A{}).Check(fn)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			fn := func(x A) verify.FailureMessage {
				return "failure"
			}
			msg := verify.Any(A{}).Check(fn)
			assertFailed(t, msg, "failure")
		})
	})

	t.Run("Should", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := verify.Any(A{}).Should(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := verify.Any(A{}).Should(pred)
			assertFailed(t, msg, "object does not meet the predicate criteria")
		})
	})

	t.Run("ShouldNot", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := verify.Any(A{}).ShouldNot(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := verify.Any(A{}).ShouldNot(pred)
			assertFailed(t, msg, "object meets the predicate criteria")
		})
	})
}
