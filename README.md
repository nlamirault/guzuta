# Guzuta

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://img.shields.io/github/release/nlamirault/guzuta.svg)](https://github.com/nlamirault/guzuta/releases)


A CLI to manage personal open source contributions.

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/guzuta/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/guzuta/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/guzuta/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/guzuta/tree/develop)


## Installation

You can download the binaries :

 * Architecture i386 [ [linux](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_linux_386.tar.gz?direct) / [windows](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_windows_386.zip?direct) / [darwin](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_darwin_386.zip?direct) / [freebsd](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_freebsd_386.zip?direct) / [openbsd](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_openbsd_386.zip?direct) ]
 * Architecture amd64 [ [linux](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_linux_amd64.tar.gz?direct) / [windows](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_windows_amd64.zip?direct) / [darwin](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_darwin_amd64.zip?direct) / [freebsd](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_freebsd_amd64.zip?direct) / [openbsd](https://dl.bintray.com//content/pacesys/utils/depcon_0.2_openbsd_amd64.zip?direct) ]


## Usage

### Github

* List user's repositories :

```bash
$ guzuta github --username=nlamirault repos
* abraracourcix - A simple URL Shortener
* aneto - A backup tool
[...]
```

* Describe user repository :

```bash
$ guzuta github --username=nlamirault --name=guzuta repo
* guzuta - A CLI to manage personal open source contributions.
```

* List user's issues :

```bash
$ guzuta github --username=nlamirault issues
- [2] open - API Statistics
- [1] open - Add Authentication
- [2] open - Invoke support
[...]
```

* List user's project issues :

```bash
$ guzuta github --username=nlamirault --name=scame issues
- [55] open - Lazy loading for modules
- [53] open - BBDB : Add feature to customize database file
- [51] open - Cask : Add support for pinned packages
- [48] open - Emacs can't start due to Org-mode error
[...]
```


### Gitlab

* List user's projects :

```bash
$ guzuta gitlab --namespace=nicolas-lamirault list
* eudyptula - The Eudyptula challenge
* dotfiles - My dotfiles
* Scame - An Emacs configuration
[...]
```

### TravisCI

* Check all projects status :

```bash
$ guzuta travisci --namespace=nlamirault
OK      nlamirault/emacs-gitlab
        nlamirault/bento
KO      nlamirault/abraracourcix
        nlamirault/nlamirault.github.io
        nlamirault/blinky
OK      nlamirault/iris
KO      nlamirault/enigma
[...]
```

* Check project status :

```bash
$ guzuta travisci --namespace=nlamirault --name=aneto
OK      nlamirault/aneto
```


### CircleCI

* Check all projects status  :

```bash
$ guzuta circleci --username=nlamirault
 OK     portefaix/portefaix-ci
 OK     nlamirault/aneto
 OK     nlamirault/gotest.el
 OK     nlamirault/phpunit.el
 OK     nlamirault/scame
[...]
```

* Check project status :

```bash
guzuta circleci --username=nlamirault --name=guzuta
$  OK     nlamirault/guzuta
```

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
[badge-release]: https://img.shields.io/github/release/nlamirault/guzuta.svg

[LICENSE]: https://github.com/nlamirault/guzuta/blob/master/LICENSE

[releases]: https://github.com/nlamirault/guzuta/releases
