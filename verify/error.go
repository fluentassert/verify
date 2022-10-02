package verify

// NoError tests if the error is nil.
func NoError(err error) FailureMessage {
	if err == nil {
		return ""
	}
	return "non-nil error:\n" + FailureMessage(err.Error())
}

// FluentError encapsulates assertions for error object.
type FluentError struct {
	FluentObj[error]
	FluentString[string]
}

// Error is used for testing error object.
func Error(got error) FluentError {
	res := FluentError{FluentObj: FluentObj[error]{got}}
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
