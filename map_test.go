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
}
