package verify

import "fmt"

// Nil tests if provided interface value is nil.
// Use it only for interfaces.
// For structs and pointers use Obj(got).Zero().
func Nil(v any) FailureMessage {
	if v == nil {
		return ""
	}
	return FailureMessage(fmt.Sprintf("value is not nil\ngot: %+v", v))
}

// NotNil tests if provided interface is not nil.
// Use it only for interfaces.
// For structs and pointers use Obj(got).NonZero().
func NotNil(v any) FailureMessage {
	if v != nil {
		return ""
	}
	return "value is <nil>"
}
