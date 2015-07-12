GOPATH := $(shell pwd)

all: build test

build:
	GOPATH=$(GOPATH) go clean
	GOPATH=$(GOPATH) go build -o bin/crypthelper -a src/crypthelper.go

test:
	GOPATH=$(GOPATH) go test ./...

