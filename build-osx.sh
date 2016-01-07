#!/bin/bash
export GOOS="darwin"
export GOARCH="amd64"
go build -o dist/gserver-osx gserver.go
