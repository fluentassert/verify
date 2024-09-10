package verify

// AssertionError is an error type used to represent failure messages from assertions.
// It is compatible with the error interface and can be used in instances where an error shall be returned instead of failing a test.
type AssertionError struct {
	Message FailureMessage
}

// Error returns the failure message as a string. It makes AssertionError compatible with the error interface.
func (err *AssertionError) Error() string {
	return string(err.Message)
}
