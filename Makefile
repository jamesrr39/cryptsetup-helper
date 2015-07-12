GOPATH := $(shell pwd)

all: build test

build:
	GOPATH=$(GOPATH) go clean
	GOPATH=$(GOPATH) go get github.com/stretchr/testify/assert
	GOPATH=$(GOPATH) go get github.com/stretchr/objx
	GOPATH=$(GOPATH) go build -o bin/crypthelper -a src/crypthelper.go

test:
	GOPATH=$(GOPATH) go test ./...

