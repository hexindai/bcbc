bcbc
======

[![Github Workflows](https://github.com/hexindai/bcbc/workflows/bcbc-ci-wf/badge.svg)](https://github.com/hexindai/bcbc/actions?query=workflow%3Abcbc-ci-wf)
[![GoVersion](https://img.shields.io/github/v/release/hexindai/bcbc)](https://github.com/hexindai/bcbc/releases/latest)
[![GoReportCard](https://goreportcard.com/badge/github.com/hexindai/bcbc)](https://goreportcard.com/report/github.com/hexindai/bcbc)

/**bcbc**/ : China UnionPay **B**ank **C**ard **B**IN **C**hecker

A tool used for checking bank card BIN in both CLI and HTTP server mode.

## Install

1. Download directly from [HERE](https://github.com/hexindai/bcbc/releases)

2. If you are a developer and Go installed, you can build from source code.

```bash
$ go get -u -v github.com/hexindai/bcbc
```

## Usage

Show this command help

```
$ bcbc -h
```

#### As a cli

```bash
$ bcbc search -c 6222021234567890123 -o json

> {"bin":"622202","bank":"ICBC","name":"中国工商银行","type":"DC","length":19}
```

#### As a HTTP server

```bash
$ bcbc serve -p :3232

$ curl http://127.0.0.1:3232/cardInfo.json?cardNo=6222021234567890123

> {"bin":"622202","bank":"ICBC","name":"中国工商银行","type":"DC","length":19}
```

#### As a library

See [![PkgGoDev](https://pkg.go.dev/badge/github.com/hexindai/bcbc/bank)](https://pkg.go.dev/github.com/hexindai/bcbc/bank?tab=doc)

## Contribution

* Add new BIN: `make add len=16 bin=621245`
* Build for generating source files: `make build`
* Change version in file `cmd/bcbc.go`
* Commit! (I will review and release it.)

## License

[MIT License](LICENSE) / Copyright (c) 2018 - 2020