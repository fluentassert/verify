// Package pred contains predicates used by f.Assertion.
package pred

import (
	"fmt"
	"reflect"
)

// Eq checks if got is equal to want.
func Eq(want interface{}) func(got interface{}) string {
	return func(got interface{}) string {
		if reflect.DeepEqual(got, want) {
			return ""
		}
		return fmt.Sprintf("got: %+v\nwant: %+v", got, want)
	}
}
