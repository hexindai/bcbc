bcbc
======

bcbc是一个根据银行卡号，通过卡bin来判断所属银行的命令行

**目前简单支持http服务启动, 未做优化切勿放到生产环境使用**

## Install

```bash
go get -u -v github.com/runrioter/bcbc
```

## Usage

```
➜ bcbc -h

bcbc is a command for searching China's bankcard info

Usage:
  bcbc [command]

Available Commands:
  help        Help about any command
  search      Search bankcard info
  serve       Start a http bankcard info server

Flags:
  -h, --help   help for bcbc

Use "bcbc [command] --help" for more information about a command.
```

## Thanks

HTTP API Provider @alipay


## License

[MIT License](LICENSE)

Copyright (c) 2018 Runrioter
