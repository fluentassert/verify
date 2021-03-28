package f_test

import (
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
