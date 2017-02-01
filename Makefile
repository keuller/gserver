.PHONY: build fmt lint run test vendor_clean vendor_get vendor_update vet list

# Prepend our _vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
GOPATH := ${PWD}/_vendor:${GOPATH}
export GOPATH

BINARY=./bin/gserver

VERSION=1.3.2
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: build

build:
	go build ${LDFLAGS} -o ${BINARY} ./src/*.go

run: build
	${BINARY}

osx:
	GOOS="darwin" GOARCH="amd64" go build ${LDFLAGS} -o ${BINARY}-osx ./src/*.go

linux:
	GOOS="linux" GOARCH="amd64" go build ${LDFLAGS} -o ${BINARY}-linux ./src/*.go

windows:
	GOOS="windows" GOARCH="amd64" go build ${LDFLAGS} -o ${BINARY}.exe ./src/*.go

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

test:
	go test ./src/...

install:
	go install ${LDFLAGS} -o ${BINARY} ./src/*.go

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

list:
	    @$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs

vendor_clean:
	rm -dRf ./_vendor/src

# We have to set GOPATH to just the _vendor
# directory to ensure that `go get` doesn't
# update packages in our primary GOPATH instead.
# This will happen if you already have the package
# installed in GOPATH since `go get` will use
# that existing location as the destination.
vendor_get: vendor_clean
	GOPATH=${PWD}/_vendor go get -d -u -v \
	github.com/gorilla/context \
	github.com/gorilla/websocket \
	github.com/gorilla/handlers \
	github.com/gorilla/mux

vendor_update: vendor_get
	rm -rf `find ./_vendor/src -type d -name .git` \
	&& rm -rf `find ./_vendor/src -type d -name .hg` \
	&& rm -rf `find ./_vendor/src -type d -name .bzr` \
	&& rm -rf `find ./_vendor/src -type d -name .svn`

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/...
