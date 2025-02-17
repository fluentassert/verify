# fluentassert

> Extensible, type-safe, fluent assertion Go library.

[![Go Reference](https://pkg.go.dev/badge/github.com/fluentassert/verify.svg)](https://pkg.go.dev/github.com/fluentassert/verify)
[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/fluentassert/verify)](https://github.com/fluentassert/verify/releases)
[![go.mod](https://img.shields.io/github/go-mod/go-version/fluentassert/verify)](go.mod)
[![LICENSE](https://img.shields.io/github/license/fluentassert/verify)](LICENSE)

[![Build Status](https://img.shields.io/github/actions/workflow/status/fluentassert/verify/build.yml?branch=main)](https://github.com/fluentassert/verify/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluentassert/verify)](https://goreportcard.com/report/github.com/fluentassert/verify)
[![Codecov](https://codecov.io/gh/fluentassert/verify/branch/main/graph/badge.svg)](https://codecov.io/gh/fluentassert/verify)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

## Description

The fluent API makes the assertion code easier
to read and write ([more](https://dave.cheney.net/2019/09/24/be-wary-of-functions-which-take-several-parameters-of-the-same-type)).

The generics (type parameters) make the usage type-safe.

The library is [extensible](#extensibility) by design.

> [!CAUTION]
> [Avoid using assertion libraries](https://go.dev/wiki/TestComments#assert-libraries).
> Instead, use [`go-cmp`](https://github.com/google/go-cmp)
> and write custom test helpers.
> Using the popular [`testify`](https://github.com/stretchr/testify)
> may be also an acceptable choice,
> especially together with [`testifylint`](https://github.com/Antonboom/testifylint)
> to avoid common mistakes.
> Use this library if you still want to.
> Consider yourself warned.

### Quick start

```go
package test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func Foo() (string, error) {
	return "wrong", nil
}

func TestFoo(t *testing.T) {
	got, err := Foo()

	verify.NoError(err).Require(t)           // Require(f) uses t.Fatal(f), stops execution if fails
	verify.String(got).Equal("ok").Assert(t) // Assert(f) uses t.Error(f), continues execution if fails
}
```

```sh
$ go test
--- FAIL: TestFoo (0.00s)
    basic_test.go:17:
        the objects are not equal
        got: "wrong"
        want: "ok"
```

⚠ Do not forget calling
[`Assert(t)`](https://pkg.go.dev/github.com/fluentassert/verify#FailureMessage.Assert)
or [`Require(t)`](https://pkg.go.dev/github.com/fluentassert/verify#FailureMessage.Require)
which executes the actual assertion.

## Supported types

Out-of-the-box the package provides fluent assertions for the following types.
The more specific function you use, the more assertions you get.

| Go type | Assertion entry point |
| - | - |
| `interface{}` ([`any`](https://pkg.go.dev/builtin#any)) | [`verify.Any()`](https://pkg.go.dev/github.com/fluentassert/verify#Any) |
| [`comparable`](https://pkg.go.dev/builtin#comparable) | [`verify.Obj()`](https://pkg.go.dev/github.com/fluentassert/verify#Obj) |
| [`constraints.Ordered`](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered) | [`verify.Ordered()`](https://pkg.go.dev/github.com/fluentassert/verify#Ordered) |
| [`constraints.Number`](https://pkg.go.dev/golang.org/x/exp/constraints#Number) | [`verify.Number()`](https://pkg.go.dev/github.com/fluentassert/verify#Number) |
| [`string`](https://pkg.go.dev/builtin#string) | [`verify.String()`](https://pkg.go.dev/github.com/fluentassert/verify#String) |
| [`error`](https://go.dev/ref/spec#Errors) | [`verify.Error()`](https://pkg.go.dev/github.com/fluentassert/verify#Error) |
| `[]T` ([slice](https://go.dev/ref/spec#Slice_types)) | [`verify.Slice()`](https://pkg.go.dev/github.com/fluentassert/verify#Slice) |
| `map[K]V` ([map](https://go.dev/ref/spec#Map_types)) | [`verify.Map()`](https://pkg.go.dev/github.com/fluentassert/verify#Map) |

Below you can find some convenience functions.

- [`verify.NoError()`](https://pkg.go.dev/github.com/fluentassert/verify#NoError)
- [`verify.IsError()`](https://pkg.go.dev/github.com/fluentassert/verify#IsError)
- [`verify.Nil()`](https://pkg.go.dev/github.com/fluentassert/verify#Nil)
- [`verify.NotNil()`](https://pkg.go.dev/github.com/fluentassert/verify#NotNil)
- [`verify.True()`](https://pkg.go.dev/github.com/fluentassert/verify#True)
- [`verify.False()`](https://pkg.go.dev/github.com/fluentassert/verify#False)

### Deep equality

For testing deep equality use
[`DeepEqual()`](https://pkg.go.dev/github.com/fluentassert/verify#FluentAny.DeepEqual)
or [`NotDeepEqual()`](https://pkg.go.dev/github.com/fluentassert/verify#FluentAny.NotDeepEqual).

```go
package test

import (
	"testing"

	"github.com/fluentassert/verify"
)

type A struct {
	Str   string
	Bool  bool
	Slice []int
}

func TestDeepEqual(t *testing.T) {
	got := A{Str: "wrong", Slice: []int{1, 4}}

	verify.Any(got).DeepEqual(
		A{Str: "string", Bool: true, Slice: []int{1, 2}},
	).Assert(t)
}
```

```sh
$ go test
--- FAIL: TestDeepEqual (0.00s)
    deepeq_test.go:20:
        mismatch (-want +got):
          test.A{
        -       Str:  "string",
        +       Str:  "wrong",
        -       Bool: true,
        +       Bool: false,
                Slice: []int{
                        1,
        -               2,
        +               4,
                },
          }
```

### Collection assertions

The library contains many collection assertion.
Below is an example of checking unordered equality.

```go
package test

import (
	"testing"

	"github.com/fluentassert/verify"
)

func TestSlice(t *testing.T) {
	got := []int { 3, 1, 2 }

	verify.Slice(got).Equivalent([]int { 2, 3, 4 }).Assert(t)
}
```

```sh
$ go test
--- FAIL: TestSlice (0.00s)
    slice_test.go:12:
        not equivalent
        got: [3 1 2]
        want: [2 3 4]
        extra got: [1]
        extra want: [4]
```

### Periodic polling

For asynchronous testing you can use
[`verify.Eventually()`](https://pkg.go.dev/github.com/fluentassert/verify#Eventually)
or [`verify.EventuallyChan()`](https://pkg.go.dev/github.com/fluentassert/verify#EventuallyChan).

```go
package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/fluentassert/verify"
)

func TestPeriodic(t *testing.T) {
	verify.Eventually(10*time.Second, time.Second, func() verify.FailureMessage {
		client := http.Client{Timeout: time.Second}
		resp, err := client.Get("http://not-existing:1234")
		if err != nil {
			return verify.NoError(err)
		}
		return verify.Number(resp.StatusCode).Lesser(300)
	}).Assert(t)
}
```

```sh
$ go test
--- FAIL: TestPeriodic (10.00s)
    async_test.go:19:
        function never passed, last failure message:
        Get "http://not-existing:1234": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

### Custom predicates

For the most basic scenarios, you can use one of the
[`Check()`](https://pkg.go.dev/github.com/fluentassert/verify#FluentAny.Check),
[`Should()`](https://pkg.go.dev/github.com/fluentassert/verify#FluentAny.Should),
[`ShouldNot()`](https://pkg.go.dev/github.com/fluentassert/verify#FluentAny.ShouldNot)
assertions.

```go
package test

import (
	"strings"
	"testing"

	"github.com/fluentassert/verify"
)

func TestShould(t *testing.T) {
	got := "wrong"

	chars := "abc"
	verify.Any(got).Should(func(got string) bool {
		return strings.ContainsAny(got, chars)
	}).Assertf(t, "does not contain any of: %s", chars)
}
```

```sh
$ go test
--- FAIL: TestShould (0.00s)
    should_test.go:16: does not contain any of: abc
        object does not meet the predicate criteria
        got: "wrong"
```

### Panics

For testing panics use [`verify.Panics()`](https://pkg.go.dev/github.com/fluentassert/verify#Panics)
and [`verify.NotPanics()`](https://pkg.go.dev/github.com/fluentassert/verify#NotPanics).

### Custom assertion function

You can create a function that returns [`FailureMessage`](https://pkg.go.dev/github.com/fluentassert/verify#FailureMessage).
Use [`verify.And()`](https://pkg.go.dev/github.com/fluentassert/verify#And)
and [`verify.Or()`](https://pkg.go.dev/github.com/fluentassert/verify#Or)
functions together with [`Prefix()`](https://pkg.go.dev/github.com/fluentassert/verify#FailureMessage.Prefix)
method to create complex assertions.

```go
package test

import (
	"testing"

	"github.com/fluentassert/verify"
)

type A struct {
	Str string
	Ok  bool
}

func TestCustom(t *testing.T) {
	got := A{Str: "something was wrong"}

	verifyA(got).Assert(t)
}

func verifyA(got A) verify.FailureMessage {
	return verify.And(
		verify.String(got.Str).Contain("ok").Prefix("got.String: "),
		verify.True(got.Ok).Prefix("got.Ok: "),
	)
}
```

```sh
$ go test
--- FAIL: TestCustom (0.00s)
    custom_test.go:17:
        got.String: the value does not contain the substring
        got: "something was wrong"
        substr: "ok"

        got.Ok: the value is false
```

## Extensibility

You can take advantage of the [`FailureMessage`](https://pkg.go.dev/github.com/fluentassert/verify#FailureMessage)
and `Fluent*` types
to create your own fluent assertions for a given type.

For reference, take a look at the implementation
of existing fluent assertions in this repository
(for example [comparable.go](comparable.go)).

## Supported Go versions

Minimal supported Go version is 1.18.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) if you want to help.

## License

**fluentassert** is licensed under the terms of [the MIT license](LICENSE).

[`github.com/google/go-cmp`](https://github.com/google/go-cmp)
(license: [BSD-3-Clause](https://pkg.go.dev/github.com/google/go-cmp/cmp?tab=licenses))
is the only [third-party dependency](go.mod).
