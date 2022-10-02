# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/pellared/fluentassert/compare/v0.2.0...HEAD)

The release provides assertions for
`constraints.Ordered`, `string`, `error`, `[]T`, `map[K]V`, `func()`.

There is a minor breaking change in the `Check` method signature.

### Added

- Add `Eventually` and `EventuallyContext` asynchronous assertions.
- Add `Panics` and `NotPanics` function assertions.
- Add `Ordered[T constraints.Ordered](got T)` function which provides following assertion
  in addition to `Comparable(got T)`:
  - `Lesser(than T)`
  - `LesserOrEqual(than T)`
  - `GreaterOrEqual(than T)`
  - `Greater(than T)`
- Add `String[T ~string](got T)` function which provides following assertion
  in addition to `Ordered(got T)`:
  - `Contains(substr string)`
  - `NotContains(substr string)`

### Changed

- Change the `Check` assertion for `any` object so that the
  provided function has to return `FailureMessage`
  instead of a `string`.

## [0.2.0](https://github.com/pellared/fluentassert/releases/tag/v0.2.0) - 2022-10-01

This release is a complete rewrite.
It is not compatible with the previous release.

The new API is type-safe and easier to extend.

It is highly probable that future releases will have no (or minimal)
breaking changes.

The next release is supposed to provide assertions for
`constraints.Ordered`, `string`, `error`, `[]T`, `map[K]V`, `func()`.

### Added

- Add `FailureMessage` which encapsulates the failure message
  and methods for error reporting.
- Add `Obj[T any](got T)` function which provides following assertions:
  - `Check(fn func(got T) string)`
  - `Should(pred func(got T) bool)`
  - `ShouldNot(pred func(got T) bool)`
  - `DeepEqual(want T, opts ...cmp.Option)`
  - `NotDeepEqual(obj T, opts ...cmp.Option)`
  - `Zero()`
  - `NonZero()`
- Add `Comparable[T comparable](got T)` function which provides following assertions
  in addition to `Obj(got T)`:
  - `Equal(want T)`
  - `NotEqual(obj T)`

### Changed

- Require Go 1.18.

### Fixed

- Fix error reporting line (use `t.Helper()` when available).

### Removed

- Remove all previous functions and types (API rewrite).

## [0.1.0](https://github.com/pellared/fluentassert/releases/tag/v0.1.0) - 2021-05-11

First release after the experiential phase.

### Added

- Add `f.Assert` that can be used instead of `t.Error` methods.
- Add `f.Require` that can be used instead of `t.Fatal` methods.
- Add `Should` assertion that can be used with custom predicates.
- Add `Eq` assertion that checks if `got` is deeply equal to `want`.
- Add `Nil` assertion that checks if `got` is `nil`.
- Add `Err` assertion that checks if `got` is an `error`.
- Add `Panic` assertion that checks if calling `got` results in a panic.
- Add `NoPanic` assertion that checks if calling `got` returns without panicking.
