package f_test

import (
	"errors"
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestAssertEq(t *testing.T) {
	f.Assert(t, 1).Eq(1, "should work with int")
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
