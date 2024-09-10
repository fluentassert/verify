# Changelog

All notable changes to this library are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this library adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/fluentassert/verify/compare/v1.1.0...HEAD)

### Added

- Add `FailureMessage.Err` method together with `AssertionError` type
  to represent assertion results as `error` type.

## [1.1.0](https://github.com/fluentassert/verify/releases/tag/v1.1.0) - 2024-02-06

This release adds length assertions.

### Added

- Add `Len` assertion for `string`, `[]T`, `map[K]V` types.

## [1.0.0](https://github.com/fluentassert/verify/releases/tag/v1.0.0) - 2023-04-05

This release contains breaking changes.

The repository is moved to `github.com/fluentassert/verify`
and the `f` package is renamed to `verify`.

The main additions are the new assertions for
`bool`, `constraints.Ordered`, `constraints.Float`, `constraints.Integer`,
`string`, `error`, `[]T`, `map[K]V`, `func()` types.

### Added

- Add `True`, `False`, assertion functions.
- Add `Nil`, `NotNil`, assertion functions.
- Add `NoError`, `IsError` assertion functions.
- Add `Panics`, `NoPanic` assertion functions.
- Add `Eventually`, `EventuallyChan` assertion functions.
- Add `Ordered` function which provides following assertions,
  in addition to `Comparable`, via `FluentOrdered` type:
  - `Lesser`
  - `LesserOrEqual`
  - `GreaterOrEqual`
  - `Greater`
- Add `String` function which provides following assertions,
  in addition to `Ordered`, via `FluentString` type:
  - `Empty`
  - `NotEmpty`
  - `Contain`
  - `NotContain`
  - `Prefix`
  - `NoPrefix`
  - `Sufix`
  - `NoSufix`
  - `EqualFold`
  - `NotEqualFold`
  - `MatchRegex`
  - `NotMatchRegex`
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
- Add `Slice` function which provides following assertions,
  in addition to `Any`, via `FluentSlice` type:
  - `Empty`
  - `NotEmpty`
  - `Equivalent`
  - `NotEquivalent`
  - `Contain`
  - `NotContain`
  - `Any`
  - `All`
  - `None`
- Add `Map` function which provides following assertions,
  in addition to `Any`, via `FlientMap` type:
  - `Empty`
  - `NotEmpty`
  - `Contain`
  - `NotContain`
  - `ContainPair`
  - `NotContainPair`
  - `Any`
  - `All`
  - `None`
- Add `FailureMessage.Prefix` method together with `And` and `Or` functions
  to facilitate creating complex assertions.

### Changed

- The `f` package is renamed to `verify`.
- Rename `Obj` and `FluentObj` to `Any` and `FluentAny`.
- Rename `Comparable` and `FluentComparable` to `Obj` and `FluentObj`.
- Change the `Check` assertion for `any` object so that the
  provided function has to return `FailureMessage`
  instead of a `string`.
- `Zero` and `NonZero` methods are moved to `FluentComparable`.
- Better failure messages.

## [0.2.0](https://github.com/fluentassert/verify/releases/tag/v0.2.0) - 2022-10-01

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

## [0.1.0](https://github.com/fluentassert/verify/releases/tag/v0.1.0) - 2021-05-11

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
