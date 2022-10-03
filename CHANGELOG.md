# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/pellared/fluentassert/compare/v0.2.0...HEAD)

This release contains breaking changes.
The `f` package has been renamed to `verify`.

The release provides assertions for
`constraints.Ordered`, `string`, `error`, `[]T`, `map[K]V`, `func()`.

### Added

- Add `True`, `False`, `NoError` assertion functions.
- Add `Ordered` function which provides following assertions,
  in addition to `Comparable`, via `FluentOrdered` type:
  - `Lesser`
  - `LesserOrEqual`
  - `GreaterOrEqual`
  - `Greater`
- Add `String` function which provides following assertions,
  in addition to `Ordered`, via `FluentString` type:
  - `Contain`
  - `NotContain`
- Add `Number` function which provides following assertions,
  in addition to `Ordered`, via `FluentNumber` type:
  - `InDelta`
  - `NotInDelta`
  - `InEpsilon`
  - `NotInEpsilon`
- Add `Error` function which provides following assertions,
  in addition to `Any` and `String` (for error message),
  via `FluentObj` and `FluentString` types:
  - `Is`
  - `IsNot`
  - `As`
  - `NotAs`
  - `WithWrapped`
  - `NoWrapped`
- Add `Func` function which provides following assertions
  via `FluentFunc` type:
  - `Panics`
  - `NotPanic`
- Add `Async` and `Periodic` functions which provides following assertions
  via `FluentAsync` and `FluentPeriodic` types:
  - `Eventually`
  - `EventuallyContext`
- Add `FailureMessage.Merge` method to to facilitate accumulating failure messages.

### Changed

- The `f` package is renamed to `verify`.
- Change the `Check` assertion for `any` object so that the
  provided function has to return `FailureMessage`
  instead of a `string`.
- `Zero` and `NonZero` methods are moved to `FluentComparable`.
- Better failure messages.

## [0.2.0](https://github.com/pellared/fluentassert/releases/tag/v0.2.0) - 2022-10-01

This release is a complete rewrite.
It is not compatible with the previous release.

The new API is type-safe and easier to extend.

The release provides assertions for `any`, `comparable`.

The next release is supposed to provide assertions for
`constraints.Ordered`, `string`, `error`, `[]T`, `map[K]V`, `func()`.

### Added

- Add `FailureMessage` which encapsulates the failure message
  and methods for error reporting.
- Add `Obj` function which provides following assertions
  via `FluentObject` type:
  - `Check`
  - `Should`
  - `ShouldNot`
  - `DeepEqual`
  - `NotDeepEqual`
  - `Zero`
  - `NonZero`
- Add `Comparable` function which provides following assertions,
  in addition to `Obj`, via `FluentComparable` type:
  - `Equal`
  - `NotEqual`

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
