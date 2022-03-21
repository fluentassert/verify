# FluentAssert

> Extensible fluent API for assertions.

[![Go Reference](https://pkg.go.dev/badge/github.com/pellared/fluentassert.svg)](https://pkg.go.dev/github.com/pellared/fluentassert)
[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/pellared/fluentassert)](https://github.com/pellared/fluentassert/releases)
[![go.mod](https://img.shields.io/github/go-mod/go-version/pellared/fluentassert)](go.mod)
[![LICENSE](https://img.shields.io/github/license/pellared/fluentassert)](LICENSE)

[![Build Status](https://img.shields.io/github/workflow/status/pellared/fluentassert/build)](https://github.com/pellared/fluentassert/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/pellared/fluentassert)](https://goreportcard.com/report/github.com/pellared/fluentassert)
[![Codecov](https://codecov.io/gh/pellared/fluentassert/branch/main/graph/badge.svg)](https://codecov.io/gh/pellared/fluentassert)

## Motivation

I always had trouble what parameter should go first and which once second.
Having a Fluent API makes it more obvious and easier to use
([more](https://dave.cheney.net/2019/09/24/be-wary-of-functions-which-take-several-parameters-of-the-same-type)).

**FluentAssert** encourages to add an additional
[assertion message](http://xunitpatterns.com/Assertion%20Message.html)
as suggested in
[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).

`Star` this repository if you find it valuable and worth maintaining.

## Quick start

```go
func TestFoo(t *testing.T) {
	got, err := Foo()

	f.Require(t, err).Eq(nil, "should be no error") // works like t.Fatalf, stops execution if fails
	f.OrderedAssert(t, got).Gt(1, "should return proper value") // works like t.Errorf, continues execution if fails
}

func Foo() (float64, error) {
	return 1.23, errors.New("not implemented")
}
```

```sh
$ go test
--- FAIL: TestFoo (0.00s)
    assert_test.go:13: should be no error
        got: not implemented
        want: <nil>
```

## Extensibility

### Using `Should` method

```go
func Test(t *testing.T) {
	got := errors.New("some error")

	f.Assert(t, got).Should(BeError(), "should return an error")
}

func BeError() func(got interface{}) string {
	return func(got interface{}) string {
		if _, ok := got.(error); ok {
			return ""
		}
		return fmt.Sprintf("got: %+v\nshould be an error", got)
	}
}
```

### Using type embedding

```go
func Test(t *testing.T) {
	got := errors.New("some error")

	Assert(t, got).Eq("", "should return nothing")
	Assert(t, err).IsError("", "should return an error")
}

type Assertion struct {
	f.Assertion
}

func Assert(t testing.TB, got interface{}) Assertion {
	return Assertion{f.Assert(t, got)}
}

func Require(t testing.TB, got interface{}) Assertion {
	return Assertion{f.Require(t, got)}
}

func (a Assertion) IsError(msg string, args ...interface{}) bool {
	a.T.Helper()
	return a.Should(beError(), msg, args...)
}

func beError() func(got interface{}) string {
	return func(got interface{}) string {
		if _, ok := got.(error); ok {
			return ""
		}
		return fmt.Sprintf("got: %+v\nshould be an error", got)
	}
}
```

## Contributing

I am open to any feedback and contribution.

Use [Discussions](https://github.com/pellared/fluentassert/discussions)
or write to me: *Robert Pajak* @ [Gophers Slack](https://invite.slack.golangbridge.org/).

You can also create an issue or a pull request.
