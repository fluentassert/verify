package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestMap(t *testing.T) {
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	t.Run("has assertions from Any", func(t *testing.T) {
		want := map[string]A{
			"id": {Str: "string", Bool: true, Slice: []int{1, 2, 3}},
		}
		got := verify.Map(want).FluentAny.Got // type embedding done properly
		assertEqual(t, got, want)
	})

	dict := map[string]A{
		"a": {Str: "text", Bool: true, Slice: []int{1, 2, 3}},
		"b": {Slice: []int{9, 8, 7}},
	}
	has := map[string]A{
		"a": {Str: "text", Bool: true, Slice: []int{1, 2, 3}},
	}
	notHas := map[string]A{
		"a": {Str: "text", Bool: true, Slice: []int{1, 2, 3}},
		"b": {Slice: []int{1, 4, 7}},
		"c": {Slice: []int{9, 8, 7}},
	}
	t.Run("Empty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(map[string]A{}).Empty()
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).Empty()
			assertFailed(t, got, "not an empty map")
		})
	})
	t.Run("NotEmpty", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).NotEmpty()
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(map[string]A{}).NotEmpty()
			assertFailed(t, got, "an empty map")
		})
	})

	t.Run("Contain", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).Contain(has)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).Contain(notHas)
			assertFailed(t, got, "not contains all pairs")
		})
	})
	t.Run("NotContain", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).NotContain(notHas)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).NotContain(has)
			assertFailed(t, got, "contains all pairs")
		})
	})

	t.Run("ContainPair", func(t *testing.T) {
		t.Run("Has", func(t *testing.T) {
			got := verify.Map(dict).ContainPair("b", A{Slice: []int{9, 8, 7}})
			assertPassed(t, got)
		})
		t.Run("DiffKey", func(t *testing.T) {
			got := verify.Map(dict).ContainPair("z", A{Slice: []int{9, 8, 7}})
			assertFailed(t, got, "has no value under key")
		})
		t.Run("DiffValue", func(t *testing.T) {
			got := verify.Map(dict).ContainPair("b", A{Slice: []int{1, 1, 1}})
			assertFailed(t, got, "has different value under key")
		})
	})
	t.Run("NotContainPair", func(t *testing.T) {
		t.Run("Has", func(t *testing.T) {
			got := verify.Map(dict).NotContainPair("b", A{Slice: []int{9, 8, 7}})
			assertFailed(t, got, "contains the pair")
		})
		t.Run("DiffKey", func(t *testing.T) {
			got := verify.Map(dict).NotContainPair("z", A{Slice: []int{9, 8, 7}})
			assertPassed(t, got)
		})
		t.Run("DiffValue", func(t *testing.T) {
			got := verify.Map(dict).NotContainPair("b", A{Slice: []int{1, 1, 1}})
			assertPassed(t, got)
		})
	})

	t.Run("Any", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).Any(func(string, A) bool { return true })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).Any(func(string, A) bool { return false })
			assertFailed(t, got, "none pair does meet the predicate criteria")
		})
	})
	t.Run("All", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).All(func(string, A) bool { return true })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).All(func(string, A) bool { return false })
			assertFailed(t, got, "a pair does not meet the predicate criteria")
		})
	})
	t.Run("None", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Map(dict).None(func(string, A) bool { return false })
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Map(dict).None(func(string, A) bool { return true })
			assertFailed(t, got, "a pair meets the predicate criteria")
		})
	})
}
