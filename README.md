# clifeed

Simple CLI feed lister.

`clifeed` reads a list of RSS and Atom feed URLs stored in `~/.clifeed` and
prints all the items published in the last month.

## Installation

```shell
go get github.com/berfr/clifeed
```

## Usage

With list of feeds stored in `~/.clifeed` and `~/go/bin` in `$PATH`:

```shell
clifeed
```
