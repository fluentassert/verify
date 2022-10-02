package f

// FluentFunc encapsulates assertions for func().
type FluentFunc[T ~func()] struct {
	Got T
}

// Func is used for testing func().
func Func[T ~func()](got T) FluentFunc[T] {
	return FluentFunc[T]{got}
}

// Panics tests if the function panics when executed.
func (x FluentFunc[T]) Panics() (msg FailureMessage) {
	defer func() {
		if r := recover(); r == nil {
			msg = "the function returned instead of panicking"
		}
	}()
	x.Got()
	return
}

// NotPanics tests if the function does not panic when executed.
func (x FluentFunc[T]) NotPanics() (msg FailureMessage) {
	defer func() {
		if r := recover(); r != nil {
			msg = "the function panicked"
		}
	}()
	x.Got()
	return
}
