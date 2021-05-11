package f_test

import (
	"errors"
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestAssertEq(t *testing.T) {
	testCases := []struct {
		desc string
		x    interface{}
		y    interface{}
	}{
		{
			desc: "int",
			x:    1,
			y:    1,
		},
		{
			desc: "map",
			x:    map[string]int{"a": 1, "b": 2},
			y:    map[string]int{"b": 2, "a": 1},
		},
		{
			desc: "error",
			x:    errors.New("abc"),
			y:    errors.New("abc"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			f.Assert(t, tc.x).Eq(tc.y, tc.desc)
		})
	}
}

func TestRequireEq(t *testing.T) {
	got := []int{1, 2}
	f.Require(t, got).Eq([]int{1, 2}, "should work with slices")
}

func TestAssertNil(t *testing.T) {
	f.Assert(t, nil).Nil("should be nil")
}

func TestAssertErr(t *testing.T) {
	got := errors.New("critical")
	f.Assert(t, got).Err("should be an error")
}

func TestAssertPanic(t *testing.T) {
	act := func() { panic("boom!") }
	f.Assert(t, act).Panic("should panic")
}

func TestAssertNoPanic(t *testing.T) {
	act := func() {}
	f.Assert(t, act).NoPanic("should return normally")
}
