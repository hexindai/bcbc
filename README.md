bcbc
======

bcbc是一个根据银行卡号，通过卡bin来判断所属银行的命令行

## 安装

1. 直接下载安装

:point_right: [下载地址](https://github.com/hexindai/bcbc/releases)

2. 如果是开发者，并且安装go

```bash
go get -u -v github.com/hexindai/bcbc
```

## 用法

**以一下*nix为例，windows下用bcbc.exe**

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

#### CLI mode

```bash
bcbc search -c 6222021234567890123 -o json
```

#### Server Mode

Default port is 3232

```bash
bcbc serve -p :3232
```

Result

```
$ curl http://127.0.0.1:3232/cardInfo.json\?cardNo\=6222021234567890123

> {"bin":"622202","bank":"ICBC","name":"中国工商银行","type":"DC","length":19}
```

## License

[MIT License](LICENSE)

Copyright (c) 2018
