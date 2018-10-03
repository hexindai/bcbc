bcbc
======

bcbc is a command about China's bankcard info

* Generate a random card No
* Start a http server for checking bankcard info
* Check card BIN via this command

## Install

1. Download

:point_right: [HERE](https://github.com/hexindai/bcbc/releases)

2. If you are a developer and Go installed

```bash
go get -u -v github.com/hexindai/bcbc
```

## Usage

```
➜ bcbc -h

bcbc is a command for searching China's bankcard info

Usage:
  bcbc [command]

Available Commands:
  help        Help about any command
  list        List all bank card BINs
  random      Return a random bankcard
  search      Search bankcard info
  serve       Serve as a http server
  version     Print version and exit

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
