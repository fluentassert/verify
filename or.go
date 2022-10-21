package verify

// Or accumalates failure messages if all are not empty.
func Or(assertions ...FailureMessage) FailureMessage {
	var msg FailureMessage
	for _, assertion := range assertions {
		if assertion == "" {
			return ""
		}
		if msg == "" {
			msg = assertion
			continue
		}
		msg = msg + "\n\n" + assertion
	}
	return msg
}
