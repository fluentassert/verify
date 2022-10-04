package verify_test

import (
	"math"
	"testing"

	"github.com/fluentassert/verify"
)

func TestNumber(t *testing.T) {
	t.Run("InDelta", func(t *testing.T) {
		t.Run("Near", func(t *testing.T) {
			msg := verify.Number(0.0).InDelta(1, 10)
			assertPassed(t, msg)
		})
		t.Run("Far", func(t *testing.T) {
			msg := verify.Number(0.0).InDelta(-100, 10)
			assertFailed(t, msg, "absolute error (distance) between numbers is greater than delta")
		})
	})

	t.Run("NotInDelta", func(t *testing.T) {
		t.Run("Near", func(t *testing.T) {
			msg := verify.Number(0.0).NotInDelta(-1.0, 10)
			assertFailed(t, msg, "absolute error (distance) between numbers is lesser or equal than delta")
		})
		t.Run("Far", func(t *testing.T) {
			msg := verify.Number(0.0).NotInDelta(100, 10)
			assertPassed(t, msg)
		})
	})

	t.Run("InEpsilon", func(t *testing.T) {
		t.Run("Near", func(t *testing.T) {
			msg := verify.Number(1.0).InEpsilon(1, 2)
			assertPassed(t, msg)
		})
		t.Run("Far", func(t *testing.T) {
			msg := verify.Number(100.0).InEpsilon(1, 2)
			assertFailed(t, msg, "relative error between numbers is greater than epsilon")
		})
	})

	t.Run("NotInEpsilon", func(t *testing.T) {
		t.Run("Near", func(t *testing.T) {
			msg := verify.Number(0.0).NotInEpsilon(1, 2)
			assertFailed(t, msg, "relative error between numbers is lesser or equal than epsilon")
		})
		t.Run("Far", func(t *testing.T) {
			msg := verify.Number(100.0).NotInEpsilon(1, 2)
			assertPassed(t, msg)
		})
	})

	t.Run("InvalidInputs", func(t *testing.T) {
		for _, fn := range []func() verify.FailureMessage{
			func() verify.FailureMessage {
				return verify.Number(math.NaN()).InDelta(1, 2)
			},
			func() verify.FailureMessage {
				return verify.Number(1.0).NotInDelta(math.NaN(), 2)
			},
			func() verify.FailureMessage {
				return verify.Number(1.0).InDelta(1, -2)
			},
			func() verify.FailureMessage {
				return verify.Number(math.NaN()).InEpsilon(1, 2)
			},
			func() verify.FailureMessage {
				return verify.Number(1.0).NotInEpsilon(math.NaN(), 2)
			},
			func() verify.FailureMessage {
				return verify.Number(1.0).InEpsilon(1, -2)
			},
			func() verify.FailureMessage {
				return verify.Number(1.0).InEpsilon(0.0, 2)
			},
		} {
			if msg := fn(); msg == "" {
				t.Error("should fail but passed")
			}
		}
	})
}
