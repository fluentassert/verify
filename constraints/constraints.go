// Package constraints defines constraints to be used with type parameters.
package constraints

// Number is a constraint that permits any numeric type
// that supports the operators < <= >= > + - * /.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

// Ordered is a constraint that permits any ordered type:
// any type that supports the operators < <= >= >.
type Ordered interface {
	Number | ~string
}
