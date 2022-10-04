# Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).

## Developing

The latest version of Go is required.

Docker is recommended.

Run `./goyek.sh` (Bash) or `.\goyek.ps1` (PowerShell)
to execute the build pipeline.

The repository contains basic confiugration for
[Visual Studio Code](https://code.visualstudio.com/).

## Releasing

This section describes how to prepare and publish a new release.

## Pre-release

Create a pull request named `Release <version>` that does the following:

1. Update the examples in [README.md](README.md)
   and make sure the documentation is up to date.
2. Update [`CHANGELOG.md`](CHANGELOG.md).
   - Change the `Unreleased` header to represent the new release.
   - Consider adding a description for the new release.
     Especially if it adds new features or introduces breaking changes.
   - Add a new `Unreleased` header above the new release, with no details.

## Release

Create a GitHib Release named `<version>` with `v<version>` tag.

The release description should include all the release notes
from the [`CHANGELOG.md`](CHANGELOG.md) for this release.
