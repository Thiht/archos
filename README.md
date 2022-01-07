# archos

`archos` prints the current OS and architecture using the Golang namings.
It can be useful to download released Go binaries with OS and architecture present in the URL.

## Use case

[golangci-lint](https://github.com/golangci/golangci-lint) [releases](https://github.com/golangci/golangci-lint/releases) have the following naming scheme:

> golangci-lint-1.43.0-darwin-amd64.tar.gz
> golangci-lint-1.43.0-darwin-arm64.tar.gz
> golangci-lint-1.43.0-freebsd-386.tar.gz
> golangci-lint-1.43.0-freebsd-amd64.tar.gz
> ...

To install the correct version depending on your OS and architecture, they [recommend](https://golangci-lint.run/usage/install/#other-ci) using their [install.sh](https://github.com/golangci/golangci-lint/blob/master/install.sh) file as follows:

```sh
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.0.0
```

You might not be comfortable executing the [`curl | sh`](https://www.seancassidy.me/dont-pipe-to-your-shell.html) pattern though. So instead thanks to `archos`, you can directly run the following command:

```sh
curl -Lo ~/.bin/golangci-lint https://github.com/golangci/golangci-lint/releases/download/v1.0.0/golangci-lint-1.0.0-$(archos).tar.gz
```

## Installation

### Using `go install`

```sh
go install github.com/Thiht/archos@latest
```

### Using the released binaries

#### With `uname`

```sh
curl -Lo ~/.bin/archos https://github.com/Thiht/archos/releases/download/v1.0.0/archos-1.0.0-$(uname -s)-$(uname -m).tar.gz
```

## Usage

```sh
$ archos
darwin-amd64

$ archos -format "{os}/{arch}"
darwin/amd64
```

## Comparison with `uname`

```sh
$ uname -s
Darwin

$ uname -m
x86_64
```
