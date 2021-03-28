package f_test

import (
	"errors"
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestFoo(t *testing.T) {
	got, err := Foo()

	f.Require(t, err).Nil("should be no error")              // works like t.Fatalf, stops execution if fails
	f.Assert(t, got).Eq("bar", "should return proper value") // works like t.Errorf, continues execution if fails
}

func Foo() (string, error) {
	return "", errors.New("not implemented")
}

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
