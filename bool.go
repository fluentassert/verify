package verify

// True tests if the object is a true value.
func True[T ~bool](got T) FailureMessage {
	if got {
		return ""
	}
	return "the value is false"
}

// False tests if the object is a false value.
func False[T ~bool](got T) FailureMessage {
	if !got {
		return ""
	}
	return "the value is true"
}
