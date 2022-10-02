package verify

// FailureMessage encapsulates failure message
// that can by emitted using objects compatible
// with the testing.TB interface.
type FailureMessage string

// Assert calls t.Error if the failure message is not empty.
func (msg FailureMessage) Assert(t interface{ Error(args ...any) }, args ...any) bool {
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
func (msg FailureMessage) Require(t interface{ Fatal(args ...any) }, args ...any) bool {
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

// Merge accumalates a non-empty failure message.
func (msg FailureMessage) Merge(header string, failureMessage FailureMessage) FailureMessage {
	if failureMessage == "" {
		return msg
	}
	if msg == "" {
		return FailureMessage(header) + "\n" + failureMessage
	}
	return msg + "\n\n" + FailureMessage(header) + "\n" + failureMessage
}
