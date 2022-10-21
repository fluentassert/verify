package verify

// FailureMessage encapsulates a failure message
// that can by emitted using objects compatible
// with the testing.TB interface.
type FailureMessage string

// Assert calls t.Error if the failure message is not empty.
// Calling Error on *testing.T marks the the function as having failed
// and continues its execution.
// Returns true when the failure message is empty.
func (msg FailureMessage) Assert(t interface {
	Error(args ...any)
}, args ...any,
) bool {
	if msg == "" {
		return true
	}
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	t.Error(append(args, "\n"+string(msg))...)
	return false
}

// Require calls t.Fatal if the failure message is not empty.
// Calling Fatal on *testing.T stops the test function execution.
// Returns true when the failure message is empty.
func (msg FailureMessage) Require(t interface {
	Fatal(args ...any)
}, args ...any,
) bool {
	if msg == "" {
		return true
	}
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	t.Fatal(append(args, "\n"+string(msg))...)
	return false
}

// Assertf calls t.Errorf if the failure message is not empty.
// Calling Errorf on *testing.T marks the the function as having failed
// and continues its execution.
// Returns true when the failure message is empty
func (msg FailureMessage) Assertf(t interface {
	Errorf(format string, args ...any)
}, format string, args ...any,
) bool {
	if msg == "" {
		return true
	}
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	t.Errorf(format+"%s", append(args, "\n"+string(msg))...)
	return false
}

// Requiref calls t.Fatalf if the failure message is not empty.
// Calling Fatalf on *testing.T stops the test function execution.
// Returns true when the failure message is empty.
func (msg FailureMessage) Requiref(t interface {
	Fatalf(format string, args ...any)
}, format string, args ...any,
) bool {
	if msg == "" {
		return true
	}
	if h, ok := t.(interface{ Helper() }); ok {
		h.Helper()
	}
	t.Fatalf(format+"%s", append(args, "\n"+string(msg))...)
	return false
}

// Prefix adds prefix if the failure message is not empty.
func (msg FailureMessage) Prefix(s string) FailureMessage {
	if msg == "" {
		return ""
	}
	return FailureMessage(s) + msg
}
