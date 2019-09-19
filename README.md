bcbc
======

[![GoDoc](https://godoc.org/github.com/hexindai/bcbc/bank?status.svg)](https://godoc.org/github.com/hexindai/bcbc/bank)
[![GoVersion](https://img.shields.io/github/v/release/hexindai/bcbc)](https://github.com/hexindai/bcbc/releases)

/**bcbc**/ : China UnionPay **B**ank **C**ard **B**IN **C**hecker

A tool used for checking bank card BIN in both CLI and HTTP server mode.

## Install

1. Download directly from [HERE](https://github.com/hexindai/bcbc/releases)

2. If you are a developer and Go installed, you can build from source code.

```bash
go get -u -v github.com/hexindai/bcbc
```

## Usage

Show this command help

```
➜ bcbc -h
```

#### CLI mode

```bash
$ bcbc search -c 6222021234567890123 -o json

> {"bin":"622202","bank":"ICBC","name":"中国工商银行","type":"DC","length":19}
```

#### Server mode

```bash
$ bcbc serve -p :3232

$ curl http://127.0.0.1:3232/cardInfo.json\?cardNo\=6222021234567890123

> {"bin":"622202","bank":"ICBC","name":"中国工商银行","type":"DC","length":19}
```

## Contribution

* Add new BIN: `make add len=16 bin=621245`
* Build for generating source files: `make build`
* Change version in file `cmd/bcbc.go`
* Commit! (I will review and release it.)

## License

[MIT License](LICENSE) / Copyright (c) 2018