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
}
