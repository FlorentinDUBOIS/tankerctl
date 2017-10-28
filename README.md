# tankerctl

Export gasoline data as [sensision](http://www.warp10.io/getting-started/#data-format) metrics.

# Installation

## From GitHub release

Coming soon :).

## From source code

The installation from source is based on the `go` binary. So, you need to have a working golang installation.

```sh
$ go get github.com/FlorentinDUBOIS/tankerctl
```

# Usage

tankerctl is a command line tool built across [viper](https://github.com/spf13/viper) and [cobra](https://github.com/spf13/cobra).

To get starting, just do :

```sh
$ tankerctl
```

This command will prompt :

```
Export gasoline data as sensision metrics

Usage:
  tankerctl [command]

Available Commands:
  help        Help about any command
  version     Display tankerctl version

Flags:
  -h, --help      help for tankerctl
  -v, --verbose   enable verbose output

Use "tankerctl [command] --help" for more information about a command.
```

# License

Please, see the license file ([LICENSE](https://github.com/FlorentinDUBOIS/tankerctl/blob/master/LICENSE)).