# FluentAssert

> Fluent API for assertions.

[![GitHub Release](https://img.shields.io/github/v/release/pellared/fluentassert)](https://github.com/pellared/fluentassert/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/pellared/fluentassert.svg)](https://pkg.go.dev/github.com/pellared/fluentassert)
[![go.mod](https://img.shields.io/github/go-mod/go-version/pellared/fluentassert)](go.mod)
[![LICENSE](https://img.shields.io/github/license/pellared/fluentassert)](LICENSE)
[![Build Status](https://img.shields.io/github/workflow/status/pellared/fluentassert/build)](https://github.com/pellared/fluentassert/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/pellared/fluentassert)](https://goreportcard.com/report/github.com/pellared/fluentassert)
[![Codecov](https://codecov.io/gh/pellared/fluentassert/branch/main/graph/badge.svg)](https://codecov.io/gh/pellared/fluentassert)

:construction: This library is currently in **experimental phase**.

`Star` this repository if you find it valuable and worth maintaining.

`Watch` this repository to get notified about new releases, issues, etc.

## Example

```go
func TestFoo(t *testing.T) {
	got, err := Foo()

	f.Require(t, err).Eq(nil, "should return any error") // works like t.Fatalf, stops execution if fails
	f.Assert(t, got).Eq("bar", "should return proper value") // works like t.Errorf, continues execution if fails
}
```

## Why

1. I always had trouble what parameter should go first and which once second. Having a Fluent API would make it obvious and easier to use ([more](https://dave.cheney.net/2019/09/24/be-wary-of-functions-which-take-several-parameters-of-the-same-type)). It also reduces the possibility to make a bug in the library. E.g. in [testify](https://github.com/stretchr/testify) the function [Contains](https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/assert#Contains) has different order of arguments than [other functions](https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/assert#Equal).

2. Encourages to add an additional [assertion message](http://xunitpatterns.com/Assertion%20Message.html) as suggested in [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures).

3. Customization via `Should` method. Example:

    ```go
    func TestError(t *testing.T) {
        got := errors.New("some error")

        f.Assert(t, got).Should(BeError(), "should return an error")
    }

    // BeError checks if got is an error.
    func BeError() func(got interface{}) string {
        return func(got interface{}) string {
            if _, ok := got.(error); ok {
                return ""
            }
            return fmt.Sprintf("got: %+v\nshould be an error", got)
        }
    }
    ```

4. Customization via type embedding. Example:

    ```go
    func Foo() (string, error) {
        return "", errors.New("not implemented")
    }

    func Test(t *testing.T) {
        got, err := Foo()

        Assert(t, got).Eq("", "should return nothing")
        Assert(t, err).Err("", "should return an error")
    }

    type Assertion struct {
        f.Assertion
    }

    func Assert(t testing.TB, got interface{}) Assertion {
        return Assertion{f.Assert(t, got)}
    }

    // Err checks if got is equal to want.
    func (a Assertion) Err(msg string, args ...interface{}) bool {
        a.T.Helper()
        return a.Should(isError(), msg, args...)
    }

    func isError() func(got interface{}) string {
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

Use [Discussions](https://github.com/pellared/fluentassert/discussions) or write to me: *Robert Pajak* @ [Gophers Slack](https://invite.slack.golangbridge.org/).

You can also create an issue or a pull request.
