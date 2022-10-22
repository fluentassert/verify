package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestSlice(t *testing.T) {
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	t.Run("has assertions from Any", func(t *testing.T) {
		want := []A{
			{Str: "string", Bool: true, Slice: []int{1, 2, 3}},
		}
		got := verify.Slice(want).FluentAny.Got // type embedding done properly
		assertEqual(t, got, want)
	})

	t.Run("Empty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice([]A{}).Empty()
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice([]A{{}}).Empty()
			assertFailed(t, got, "not an empty slice")
		})
	})
	t.Run("NotEmpty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice([]A{{}}).NotEmpty()
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice([]A{}).NotEmpty()
			assertFailed(t, got, "an empty slice")
		})
	})

	list := []A{
		{Str: "text", Bool: true, Slice: []int{1, 2, 3}},
		{Slice: []int{9, 8, 7}},
	}
	eq := []A{
		{Slice: []int{9, 8, 7}},
		{Str: "text", Bool: true, Slice: []int{1, 2, 3}},
	}
	notEq := []A{
		{Slice: []int{0, 0, 0}},
		{Str: "text", Bool: true, Slice: []int{1, 2, 3}},
	}
	t.Run("Equivalent", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).Equivalent(eq)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).Equivalent(notEq)
			assertFailed(t, got, "not equivalent")
		})
	})
	t.Run("NotEquivalent", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).NotEquivalent(notEq)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).NotEquivalent(eq)
			assertFailed(t, got, "equivalent")
		})
	})

	t.Run("Contain", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).Contain(A{Str: "text", Bool: true, Slice: []int{1, 2, 3}})
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).Contain(A{})
			assertFailed(t, got, "slice does not contain the item")
		})
	})
	t.Run("NotContain", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).NotContain(A{})
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).NotContain(A{Str: "text", Bool: true, Slice: []int{1, 2, 3}})
			assertFailed(t, got, "slice contains the item")
		})
	})

	t.Run("Any", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).Any(func(a A) bool { return true })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).Any(func(a A) bool { return false })
			assertFailed(t, got, "none item does meet the predicate criteria")
		})
	})
	t.Run("All", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).All(func(a A) bool { return true })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).All(func(a A) bool { return false })
			assertFailed(t, got, "an item does not meet the predicate criteria")
		})
	})
	t.Run("None", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Slice(list).None(func(a A) bool { return false })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Slice(list).None(func(a A) bool { return true })
			assertFailed(t, got, "an item meets the predicate criteria")
		})
	})
}
