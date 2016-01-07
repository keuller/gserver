#!/bin/bash
export GOOS="windows"
export GOARCH="amd64"
go build -o dist/gserver.exe gserver.go
