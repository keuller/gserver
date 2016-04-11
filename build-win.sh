#!/bin/bash
export GOOS="windows"
export GOARCH="amd64"
export VERSION=1.2.2
export BUILD=`git rev-parse HEAD`
go build -ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}" -o ./bin/gserver.exe ./src/gserver.go
