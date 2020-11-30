# cws-clocking-cli

[![CI](https://github.com/yano3/cws-clocking-cli/workflows/CI/badge.svg)](https://github.com/yano3/cws-clocking-cli/actions)

Clocking in/out cli for COMPANY Web Service.

## Usage

Clocking in

```
cws-clocking-cli
```

Clocking out

```
cws-clocking-cli -out
```

## Installation

### macOS

If you use [Homebrew](https://brew.sh):

```
brew tap yano3/tap
brew install cws-clocking-cli
```

### Other platforms

Download binary from [releases page](https://github.com/yano3/cws-clocking-cli/releases) or use `go get` command.

```console
$ go get -u github.com/yano3/cws-clocking-cli
```

## Configuration

Set environment variables bellow.

```
export CWS_URL=<Put your COMPANY Web Service URL>
export CWS_USERID=<Put your user id>
export CWS_PASSWORD=<Put your password>
```
