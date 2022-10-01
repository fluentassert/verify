package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func Example() {
	t := &testing.T{} // provided by test
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
	got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
	f.Obj(got).DeepEq(want).Assert(t)
}
