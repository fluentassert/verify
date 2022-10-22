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
}
