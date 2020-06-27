# git-vertag

A tool to manage version-tag with the semantic versioning specification.

[![Go Report Card](https://goreportcard.com/badge/github.com/kyoh86/git-vertag)](https://goreportcard.com/report/github.com/kyoh86/git-vertag)
[![Coverage Status](https://img.shields.io/codecov/c/github/kyoh86/git-vertag.svg)](https://codecov.io/gh/kyoh86/git-vertag)
[![Release](https://github.com/kyoh86/git-vertag/workflows/Release/badge.svg)](https://github.com/kyoh86/git-vertag/releases)

This tool supports only `update` semantic versiong tags.
So it won't delete, replase and rollback versions.

## Install

```
go get github.com/kyoh86/git-vertag
```

## Usage

`git-vertag` won't rollback versions.
There's some sub commands to manipulate versions.

| Sub command   | Description                                                   |
| ------------- | --------------------------------------------------------------|
| get (default) | Gets the current version tag.                                 |
| major         | Creates a tag for the next major version and prints it.       |
| minor         | Creates a tag for the next minor version and prints it.       |
| patch         | Creates a tag for the next patch version and prints it.       |
| pre           | Creates a tag for the next pre-release version and prints it. |
| build         | Creates a tag for the next build version and prints it.       |

See `git vertag --help-long` for detail.

## Example

### Case 1: Update major

```console
$ git vertag
v1.2.3
$ git vertag major
v2.0.0
```

### Case 2: Update minor with messages

```console
$ git vertag
v1.2.3
$ git vertag minor --message 'Supports new notation'
v1.3.0-alpha.2
```

### Case 3: Update patch with pre-release meta informations

```console
$ git vertag
v1.2.3
$ git vertag patch --pre alpha --pre 2
v1.2.4-alpha.2
```

### Case 4: Release current patch that has pre-released.

```console
$ git vertag
v1.2.4-alpha.2
$ git vertag release
v1.2.4
```

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
