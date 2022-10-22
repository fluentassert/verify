package verify_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/fluentassert/verify"
)

func TestNoError(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		got := verify.NoError(nil)
		assertPassed(t, got)
	})
	t.Run("Failed", func(t *testing.T) {
		got := verify.NoError(errors.New("some error"))
		assertFailed(t, got, "non-nil error:\nsome error")
	})
}

func TestIsError(t *testing.T) {
	t.Run("Passed", func(t *testing.T) {
		got := verify.IsError(errors.New(""))
		assertPassed(t, got)
	})
	t.Run("Failed", func(t *testing.T) {
		var err error
		got := verify.IsError(err)
		assertFailed(t, got, "the error is <nil>")
	})
}

func TestError(t *testing.T) {
	err := errors.New("expected error")
	t.Run("Is", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Error(fmt.Errorf("wrap: %w", err)).Is(err)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Error(fmt.Errorf("wrap: %v", err)).Is(err)
			assertFailed(t, got, "no error in err's chain matches")
		})
	})
	t.Run("IsNot", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			got := verify.Error(fmt.Errorf("wrap: %v", err)).IsNot(err)
			assertPassed(t, got)
		})
		t.Run("Failed", func(t *testing.T) {
			got := verify.Error(fmt.Errorf("wrap: %w", err)).IsNot(err)
			assertFailed(t, got, "some error in err's chain matches")
		})
	})

	t.Run("As", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			var wantErr *stubError
			got := verify.Error(fmt.Errorf("wrap: %w", &stubError{})).As(&wantErr)
			assertPassed(t, got)
			assertEqual(t, wantErr, &stubError{})
		})
		t.Run("Failed", func(t *testing.T) {
			var wantErr *stubError
			got := verify.Error(fmt.Errorf("wrap: %v", &stubError{})).As(&wantErr)
			assertFailed(t, got, "no error in err's chain matches")
			assertEqual(t, wantErr, nil)
		})
	})
	t.Run("AsNot", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			var wantErr *stubError
			got := verify.Error(fmt.Errorf("wrap: %v", &stubError{})).AsNot(&wantErr)
			assertPassed(t, got)
			assertEqual(t, wantErr, nil)
		})
		t.Run("Failed", func(t *testing.T) {
			var wantErr *stubError
			got := verify.Error(fmt.Errorf("wrap: %w", &stubError{})).AsNot(&wantErr)
			assertFailed(t, got, "some error in err's chain matches")
			assertEqual(t, wantErr, &stubError{})
		})
	})

	t.Run("has assertions from Obj", func(t *testing.T) {
		want := errors.New("an error")
		got := verify.Error(want).FluentAny.Got // type embedding done properly
		assertEqual(t, got, want, cmpopts.EquateErrors())
	})
	t.Run("has assertions from String, Ordered, Comparable for string", func(t *testing.T) {
		want := "an error"
		got := verify.Error(errors.New(want)).FluentString.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}

type stubError struct{}

func (*stubError) Error() string { return "stubError" }
