package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/pellared/fluentassert/verify"
)

func TestAsync(t *testing.T) {
	verify.Periodic(10*time.Second, time.Second, func() verify.FailureMessage {
		client := http.Client{Timeout: time.Second}
		_, err := client.Get("http://not-existing:1234")
		if err != nil {
			return verify.Error(err).Zero()
		}
		return ""
	}).Eventually().Assert(t)
}
