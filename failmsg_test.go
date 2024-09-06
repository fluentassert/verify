package verify_test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestFailureMessage(t *testing.T) {
	t.Run("Assert", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			mock := &errorMock{}
			got := verify.FailureMessage("").Assert(mock)
			assertTrue(t, got)
			assertEqual(t, mock, &errorMock{})
		})
		t.Run("NoArgs", func(t *testing.T) {
			mock := &errorMock{}
			got := verify.FailureMessage("failed").Assert(mock)
			assertFalse(t, got)
			assertEqual(t, mock, &errorMock{
				HelperCalled: true,
				Called:       true,
				Args:         []any{"\nfailed"},
			})
		})
		t.Run("WithArgs", func(t *testing.T) {
			mock := &errorMock{}
			got := verify.FailureMessage("failed").Assert(mock, "arg1", 2)
			assertFalse(t, got)
			assertEqual(t, mock, &errorMock{
				HelperCalled: true,
				Called:       true,
				Args:         []any{"arg1", 2, "\nfailed"},
			})
		})
	})

	t.Run("Require", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			mock := &fatalMock{}
			got := verify.FailureMessage("").Require(mock)
			assertTrue(t, got)
			assertEqual(t, mock, &fatalMock{})
		})
		t.Run("NoArgs", func(t *testing.T) {
			mock := &fatalMock{}
			got := verify.FailureMessage("failed").Require(mock)
			assertFalse(t, got)
			assertEqual(t, mock, &fatalMock{
				HelperCalled: true,
				Called:       true,
				Args:         []any{"\nfailed"},
			})
		})
		t.Run("WithArgs", func(t *testing.T) {
			mock := &fatalMock{}
			got := verify.FailureMessage("failed").Require(mock, "arg1", 2)
			assertFalse(t, got)
			assertEqual(t, mock, &fatalMock{
				HelperCalled: true,
				Called:       true,
				Args:         []any{"arg1", 2, "\nfailed"},
			})
		})
	})

	t.Run("Assertf", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			mock := &errorfMock{}
			got := verify.FailureMessage("").Assertf(mock, "should pass")
			assertTrue(t, got)
			assertEqual(t, mock, &errorfMock{})
		})
		t.Run("NoArgs", func(t *testing.T) {
			mock := &errorfMock{}
			got := verify.FailureMessage("failed").Assertf(mock, "should pass")
			assertFalse(t, got)
			assertEqual(t, mock, &errorfMock{
				HelperCalled: true,
				Called:       true,
				Format:       "should pass%s",
				Args:         []any{"\nfailed"},
			})
		})
		t.Run("WithArgs", func(t *testing.T) {
			mock := &errorfMock{}
			got := verify.FailureMessage("failed").Assertf(mock, "should work %d", 1)
			assertFalse(t, got)
			assertEqual(t, mock, &errorfMock{
				HelperCalled: true,
				Called:       true,
				Format:       "should work %d%s",
				Args:         []any{1, "\nfailed"},
			})
		})
	})

	t.Run("Requiref", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			mock := &fatalfMock{}
			got := verify.FailureMessage("").Requiref(mock, "should pass")
			assertTrue(t, got)
			assertEqual(t, mock, &fatalfMock{})
		})
		t.Run("NoArgs", func(t *testing.T) {
			mock := &fatalfMock{}
			got := verify.FailureMessage("failed").Requiref(mock, "should pass")
			assertFalse(t, got)
			assertEqual(t, mock, &fatalfMock{
				HelperCalled: true,
				Called:       true,
				Format:       "should pass%s",
				Args:         []any{"\nfailed"},
			})
		})
		t.Run("WithArgs", func(t *testing.T) {
			mock := &fatalfMock{}
			got := verify.FailureMessage("failed").Requiref(mock, "should work %d", 1)
			assertFalse(t, got)
			assertEqual(t, mock, &fatalfMock{
				HelperCalled: true,
				Called:       true,
				Format:       "should work %d%s",
				Args:         []any{1, "\nfailed"},
			})
		})
	})

	t.Run("Prefix", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			got := verify.FailureMessage("").Prefix("[fail]")
			assertEqual(t, got, "")
		})
		t.Run("Empty", func(t *testing.T) {
			got := verify.FailureMessage("errored").Prefix("[fail] ")
			assertEqual(t, got, "[fail] errored")
		})
	})

	t.Run("AsError", func(t *testing.T) {
		t.Run("With Message", func(t *testing.T) {
			got := verify.FailureMessage("failed").AsError()
			assertEqual(t, got.Error(), "failed")
		})

		t.Run("Empty", func(t *testing.T) {
			got := verify.FailureMessage("").AsError()
			assertEqual(t, got, nil)
		})
	})
}

type errorMock struct {
	Called       bool
	Args         []any
	HelperCalled bool
}

func (mock *errorMock) Error(args ...any) {
	mock.Called = true
	mock.Args = args
}

func (mock *errorMock) Helper() {
	mock.HelperCalled = true
}

type fatalMock struct {
	Called       bool
	Args         []any
	HelperCalled bool
}

func (mock *fatalMock) Fatal(args ...any) {
	mock.Called = true
	mock.Args = args
}

func (mock *fatalMock) Helper() {
	mock.HelperCalled = true
}

type errorfMock struct {
	Called       bool
	Format       string
	Args         []any
	HelperCalled bool
}

func (mock *errorfMock) Errorf(format string, args ...any) {
	mock.Called = true
	mock.Format = format
	mock.Args = args
}

func (mock *errorfMock) Helper() {
	mock.HelperCalled = true
}

type fatalfMock struct {
	Called       bool
	Format       string
	Args         []any
	HelperCalled bool
}

func (mock *fatalfMock) Fatalf(format string, args ...any) {
	mock.Called = true
	mock.Format = format
	mock.Args = args
}

func (mock *fatalfMock) Helper() {
	mock.HelperCalled = true
}
