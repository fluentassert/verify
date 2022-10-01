package f_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/pellared/fluentassert/f"
)

func assertEqual[T any](t *testing.T, got, want T) {
	t.Helper()
	if diff := cmp.Diff(want, got); diff != "" {
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

func assertPassed(t *testing.T, got f.FailureMessage) {
	t.Helper()
	if got != "" {
		t.Errorf("should pass; got = %s", string(got))
	}
}

func assertFailed(t *testing.T, got f.FailureMessage, substr string) {
	t.Helper()
	if !strings.Contains(string(got), substr) {
		t.Errorf("should cointain = %s; got = %s", substr, string(got))
	}
}
