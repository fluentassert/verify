package verify

// NoError tests if the error is nil.
func NoError(err error) FailureMessage {
	if err == nil {
		return ""
	}
	return "non-nil error:\n" + FailureMessage(err.Error())
}

// IsError tests if the error is non-nil.
func IsError(err error) FailureMessage {
	if err != nil {
		return ""
	}
	return "the error is <nil>"
}

// FluentError encapsulates assertions for error object.
type FluentError struct {
	FluentAny[error]
	FluentString[string]
}

// Error is used for testing error object.
func Error(got error) FluentError {
	res := FluentError{FluentAny: FluentAny[error]{got}}
	if got != nil {
		res.FluentString.Got = got.Error()
	}
	return res
}

// TODO: Is

// TODO: IsNot

// TODO: As

// TODO: NotAs

// TODO: WithWrapped

// TODO: NoWrapped
