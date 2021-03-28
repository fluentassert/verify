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
	got := Foo()

	f.Assert(t, got).Eq("bar", "should return proper value") // works like t.Errorf, continues execution if fails
	f.Require(t, got).Eq("bar", "should return proper value") // works like t.Fatalf, stops execution if fails
}
```

## Why

1. I always had trouble what parameter should go first and which once second. Having a Fluent API would make it obvious and easier to use ([more](https://dave.cheney.net/2019/09/24/be-wary-of-functions-which-take-several-parameters-of-the-same-type)). It also reduces the possibility to make a bug in the library. E.g. in [testify](https://github.com/stretchr/testify) the function [Contains](https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/assert#Contains) has different order of arguments than [other functions](https://pkg.go.dev/github.com/stretchr/testify@v1.7.0/assert#Equal).
2. Customizable via `Should` method.

## Contributing

I am open to any feedback and contribution.

Use [Discussions](https://github.com/pellared/fluentassert/discussions) or write to me: *Robert Pajak* @ [Gophers Slack](https://invite.slack.golangbridge.org/).

You can also create an issue or a pull request.
