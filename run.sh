#! /bin/sh

set -e

glide install

export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED=0

go build -v -o dist/glifo

## TODO add docker
# docker build -t glifo . 