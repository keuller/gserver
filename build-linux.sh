#!/bin/bash
export GOOS="linux"
export GOARCH="amd64"
go build -o dist/gserver gserver.go
