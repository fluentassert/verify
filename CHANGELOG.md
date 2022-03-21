# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased](https://github.com/pellared/fluentassert/compare/v0.1.0...HEAD)

Redesign of the API by using generics.

### Added

- Add `f.OrderedAssert` and `f.OrderedRequire` which
  operates on `constraints.Ordered` instead of `any`.
- Add `Gt` ordered assertion that checks if `got` is greater than `want`.

### Changed

- Require Go 1.18.
- Existing parameters are `any` instead of `interface{}`.

### Fixed

- Fix error reporting line (usage of `t.Helper()`).

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
