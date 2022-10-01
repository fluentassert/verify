package f

import "github.com/google/go-cmp/cmp"

type FluentObj[T any] struct {
	Got T
}

func Obj[T any](got T) FluentObj[T] {
	return FluentObj[T]{got}
}

func (x FluentObj[T]) Should(predicate func(got T) string) FailureMessage {
	return FailureMessage(predicate(x.Got))
}

func (x FluentObj[T]) DeepEq(want T) FailureMessage {
	diff := cmp.Diff(want, x.Got)
	if diff == "" {
		return ""
	}
	return FailureMessage("mismatch (-want +got):\n" + diff)
}
