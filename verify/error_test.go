package verify_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/pellared/fluentassert/verify"
)

func TestError(t *testing.T) {
	t.Run("has assertions from Obj", func(t *testing.T) {
		want := errors.New("an error")
		got := verify.Error(want).FluentObj.Got // type embedding done properly
		assertEqual(t, got, want, cmpopts.EquateErrors())
	})

	t.Run("has assertions from String, Ordered, Comparable for string", func(t *testing.T) {
		want := "an error"
		got := verify.Error(errors.New(want)).FluentString.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}
