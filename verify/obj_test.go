package verify_test

import (
	"testing"

	"github.com/pellared/fluentassert/verify"
)

func TestObj(t *testing.T) {
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	t.Run("DeepEqual", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := verify.Obj(got).DeepEqual(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := verify.Obj(got).DeepEqual(want)
			assertFailed(t, msg, "mismatch (-want +got):\n")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Obj(got).DeepEqual(nil)
			assertPassed(t, msg)
		})
	})

	t.Run("NotDeepEqual", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := verify.Obj(got).NotDeepEqual(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := verify.Obj(got).NotDeepEqual(want)
			assertFailed(t, msg, "the objects are equal")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Obj(got).NotDeepEqual(nil)
			assertFailed(t, msg, "the objects are equal")
		})
	})

	t.Run("Zero", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := A{}
			msg := verify.Obj(got).Zero()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			got := A{Str: "wrong"}
			msg := verify.Obj(got).Zero()
			assertFailed(t, msg, "not a zero value (-want +got):\n")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Obj(got).Zero()
			assertPassed(t, msg)
		})
	})

	t.Run("NonZero", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := A{Str: "string"}
			msg := verify.Obj(got).NonZero()
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			got := A{}
			msg := verify.Obj(got).NonZero()
			assertFailed(t, msg, "a zero value")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := verify.Obj(got).NonZero()
			assertFailed(t, msg, "a zero value")
		})
	})

	t.Run("Check", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			fn := func(x A) verify.FailureMessage {
				return ""
			}
			msg := verify.Obj(A{}).Check(fn)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			fn := func(x A) verify.FailureMessage {
				return "failure"
			}
			msg := verify.Obj(A{}).Check(fn)
			assertFailed(t, msg, "failure")
		})
	})

	t.Run("Should", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := verify.Obj(A{}).Should(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := verify.Obj(A{}).Should(pred)
			assertFailed(t, msg, "object does not meet the predicate criteria")
		})
	})

	t.Run("ShouldNot", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := verify.Obj(A{}).ShouldNot(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := verify.Obj(A{}).ShouldNot(pred)
			assertFailed(t, msg, "object meets the predicate criteria")
		})
	})
}
