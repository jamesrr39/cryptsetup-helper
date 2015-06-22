GOPATH := /home/james/src/dm-crypt-helper/go-app

build:
	GOPATH=$(GOPATH) go clean
	GOPATH=$(GOPATH) go build -o bin/crypthelper -a src/crypthelper.go

