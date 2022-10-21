package verify_test

import (
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/fluentassert/verify"
)

func assertEqual[T any](t *testing.T, got, want T, opts ...cmp.Option) {
	t.Helper()
	if diff := cmp.Diff(want, got, opts...); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("want = true; got = false")
	}
}

func assertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("want = true; got = false")
	}
}

func assertPassed(t *testing.T, got verify.FailureMessage) {
	t.Helper()
	if got != "" {
		t.Errorf("\nSHOULD PASS; GOT:\n%s", string(got))
	}
}

func assertFailed(t *testing.T, got verify.FailureMessage, substr string) {
	t.Helper()
	if !strings.Contains(string(got), substr) {
		t.Errorf("\nSHOULD FAIL AND CONTAIN:\n%s\nGOT:\n%s", substr, string(got))
	}
}

func assertNoLeak(t *testing.T) {
	t.Helper()
	goNum := runtime.NumGoroutine()
	t.Cleanup(func() {
		t.Helper()
		assertEqual(t, goNum, runtime.NumGoroutine())
	})
}
