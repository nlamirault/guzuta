# Guzuta

[![License Apache 2][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]
[![coveralls][badge-coveralls]][coveralls]

A CLI to manage personal open source contributions.

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

* Launch unit tests :

        $ make test

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
[travis]: https://travis-ci.org/nlamirault/guzuta
[badge-travis]: http://img.shields.io/travis/nlamirault/guzuta.svg?style=flat
[badge-drone]: https://drone.io/github.com/nlamirault/guzuta/status.png
[drone]: https://drone.io/github.com/nlamirault/guzuta/latest
[badge-coveralls]: https://coveralls.io/repos/nlamirault/guzuta/badge.svg
[coveralls]: https://coveralls.io/r/nlamirault/guzuta

[releases]: https://github.com/nlamirault/guzuta/releases
