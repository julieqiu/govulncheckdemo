# govulncheck demo

This repository provides a demo code for govulncheck.

## Getting Started

```
$ go install golang.org/x/vuln/cmd/govulncheck@latest
$ govulncheck ./...
```

If you run into a "command not found" error, make sure your PATH is set to include $GOPATH/bin.

The GOPATH environment variable specifies the location of your workspace. By
default, the go command sets this to $HOME/go.

```
$ export GOPATH=$HOME/go # this can be set to any location
$ export PATH=$PATH:$GOPATH/bin
```

Verify that you are running Go version 1.18 or later:

```
$ go version
```

## Running govulncheck

This repository includes several demo modules for running govulncheck.

The output for each is described in the respective directory.

Outputs were generated with
golang.org/x/vuln@v0.0.0-20230320143955-c17b6c9bd0e8.
