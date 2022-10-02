# fluentassert

> Extensible, type-safe, fluent assertion Go library.

[![Go Reference](https://pkg.go.dev/badge/github.com/pellared/fluentassert.svg)](https://pkg.go.dev/github.com/pellared/fluentassert)
[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/pellared/fluentassert)](https://github.com/pellared/fluentassert/releases)
[![go.mod](https://img.shields.io/github/go-mod/go-version/pellared/fluentassert)](go.mod)
[![LICENSE](https://img.shields.io/github/license/pellared/fluentassert)](LICENSE)

[![Build Status](https://img.shields.io/github/workflow/status/pellared/fluentassert/build)](https://github.com/pellared/fluentassert/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/pellared/fluentassert)](https://goreportcard.com/report/github.com/pellared/fluentassert)
[![Codecov](https://codecov.io/gh/pellared/fluentassert/branch/main/graph/badge.svg)](https://codecov.io/gh/pellared/fluentassert)

Please ⭐ `Star` this repository if you find it valuable and worth maintaining.

## Description

The fluent API makes the assertion code easier
to read and write ([more](https://dave.cheney.net/2019/09/24/be-wary-of-functions-which-take-several-parameters-of-the-same-type)).

The generics (type parameters) make the usage type-safe.

The library is [extensible](#custom-assertions).

## Examples

### Quick start

```go
package test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func Foo() (string, error) {
	return "wrong", nil
}

func TestFoo(t *testing.T) {
	got, err := Foo()
	f.Obj(err).Zero().Require(t, "should be no error") // Require(f) uses t.Fatal(f), stops execution if fails
	f.String(got).Contains("ok").Assert(t)             // Assert(f) uses t.Error(f), continues execution if fails
}
```

```sh
$ go test
--- FAIL: TestFoo (0.00s)
    basic_test.go:16:
        the object does not contain the substring
        got: "wrong"
        substr: "ok"
```

### Deep equality

```go
package test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

type A struct {
	Str   string
	Bool  bool
	Slice []int
}

func TestDeepEqual(t *testing.T) {
	got := A{Str: "wrong", Slice: []int{1, 4}}

	f.Obj(got).DeepEqual(
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

### Asynchronous assertion

```go
package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/pellared/fluentassert/f"
)

func TestAsync(t *testing.T) {
	f.Periodic(10*time.Second, time.Second, func() f.FailureMessage {
		client := http.Client{Timeout: time.Second}
		resp, err := client.Get("http://not-existing:1234")
		if err != nil {
			return f.FailureMessage(err.Error())
		}
		return f.Obj(resp.StatusCode).DeepEqual(http.StatusOK)
	}).Eventually().Assert(t)
}
```

```sh
$ go test
--- FAIL: TestAsync (10.00s)
    async_test.go:19:
        timeout
        function always failed
        last failure message:
        Get "http://not-existing:1234": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
```

## Custom assertions

You can take advantage of the `f.FailureMessage` and `f.Fluent*` types
to create your own fluent assertions.

For reference, take a look at the implementation
of existing fluent assertions in this repository
(for example [comparable.go](f/comparable.go)).

### Supported Go versions

Minimal supported Go version is 1.18.

## Contributing

Feel free to create an issue or propose a pull request.

### Developing

Run `./goyek.sh` (Bash) or `.\goyek.ps1` (PowerShell)
to execute the build pipeline.

The repository contains confiugration for
[Visual Studio Code](https://code.visualstudio.com/).

## License

**fluentassert** is licensed under the terms of [the MIT license](LICENSE).

[`github.com/google/go-cmp`](https://github.com/google/go-cmp)
(license: [BSD-3-Clause](https://pkg.go.dev/github.com/google/go-cmp/cmp?tab=licenses))
is the only [third-party dependency](go.mod).
