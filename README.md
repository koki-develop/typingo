# Typingo

[![GitHub release (latest by date)](https://img.shields.io/github/v/release/koki-develop/typingo)](https://github.com/koki-develop/typingo/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/koki-develop/typingo/ci.yml?logo=github)](https://github.com/koki-develop/typingo/actions/workflows/ci.yml)
[![Maintainability](https://img.shields.io/codeclimate/maintainability/koki-develop/typingo?style=flat&logo=codeclimate)](https://codeclimate.com/github/koki-develop/typingo/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/koki-develop/typingo)](https://goreportcard.com/report/github.com/koki-develop/typingo)
[![LICENSE](https://img.shields.io/github/license/koki-develop/typingo)](./LICENSE)

Typing game written in Go.

![demo](./docs/demo.gif)

- [Installation](#installation)
- [Usage](#usage)
- [LICENSE](#license)

## Installation

### Homebrew

```console
$ brew install koki-develop/tap/typingo
```

### `go install`

```console
$ go install github.com/koki-develop/typingo@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/typingo/releases/latest).

## Usage

```console
$ typingo --help
Typing game written in Go

Usage:
  typingo [flags]

Flags:
  -b, --beep            whether to beep when mistaken (default true)
  -h, --help            help for typingo
  -n, --num-texts int   the number of texts (default 10)
  -v, --version         version for typingo
```

## LICENSE

[MIT](./LICENSE)
