# Guzuta

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fguzuta.svg)](https://badge.fury.io/gh/nlamirault%2Fguzuta)

A CLI to manage personal open source contributions.

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/guzuta/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/guzuta/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/guzuta/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/guzuta/tree/develop)


## Installation

Download binary from [releases][] for your platform.

## Usage

```bash
$ guzuta
NAME:
   guzuta - A CLI for Open source repositories

USAGE:
   guzuta [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR(S):
   Nicolas Lamirault <nicolas.lamirault@gmail.com>

COMMANDS:
   github
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level, -l "info"       Log level (options: debug, info, warn, error, fatal, panic)
   --github-token               Github access token [$GUZUTA_GITHUB_TOKEN]
   --help, -h                   show help
   --version, -v                print the version

```

You must have an **access token** to use the CLI.


## Development

* Checkout the projet and install it into $GOPATH :

        $ mkdir -p $GOPATH/src/github.com/nlamirault
        $ git clone https://github.com/nlamirault/guzuta.git $GOPATH/src/github.com/nlamirault/guzutay
        $ cd $GOPATH/src/github.com/nlamirault/guzuta

* Install requirements :

        $ make init

* Initialize dependencies :

        $ make deps

* Make the binary:

        $ make build

* Launch all unit tests :

        $ make test

* Launch some unit tests :

        $ gb test github.com/nlamirault/guzuta/providers/gitlab/

* Check code coverage for project or specific package :

        $ make coverage

* For a new release, it will run a build which cross-compiles binaries for
  a variety of architectures and operating systems:

        $ make release


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE][] for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache_2-green.svg?style=flat
[LICENSE]: https://github.com/nlamirault/guzuta/blob/master/LICENSE

[releases]: https://github.com/nlamirault/guzuta/releases
