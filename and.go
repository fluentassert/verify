package verify

// And accumalates non-empty failure messages.
func And(assertions ...FailureMessage) FailureMessage {
	var msg FailureMessage
	for _, assertion := range assertions {
		if assertion == "" {
			continue
		}
		if msg == "" {
			msg = assertion
			continue
		}
		msg = msg + "\n\n" + assertion
	}
	return msg
}
