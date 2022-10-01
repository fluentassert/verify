package f_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
