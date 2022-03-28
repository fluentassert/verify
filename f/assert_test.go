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

func TestAssertGt(t *testing.T) {
	f.OrderedAssert(t, 3).Gt(1, "should be greater than 1")
}

func TestRequireGt(t *testing.T) {
	f.OrderedRequire(t, 3.1).Gt(3, "should be greater than 1")
}

func TestAssertNil(t *testing.T) {
	var b []byte
	f.Assert(t, b).Eq(nil, "should be nil")
}

func TestAssertErr(t *testing.T) {
	got := errors.New("critical")
	f.ErrorAssert(t, got).Returned("should be an error")
}

func TestRequireNoErr(t *testing.T) {
	var got error
	f.ErrorRequire(t, got).Nil("should be no error")
}

func TestAssertPanic(t *testing.T) {
	act := func() { panic("boom!") }
	f.Assert(t, act).Panic("should panic")
}

func TestAssertNoPanic(t *testing.T) {
	act := func() {}
	f.Assert(t, act).NoPanic("should return normally")
}
