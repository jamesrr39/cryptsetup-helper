GOPATH := $(shell pwd)

build:
	GOPATH=$(GOPATH) go clean
	GOPATH=$(GOPATH) go build -o bin/crypthelper -a src/crypthelper.go

test:
	GOPATH=$(GOPATH) go test ./...
