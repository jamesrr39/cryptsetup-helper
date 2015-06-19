GOPATH := /home/james/src/dm-crypt-helper/go-app

build:
	GOPATH=$(GOPATH) go clean
	GOPATH=$(GOPATH) go build -o bin/crypthelper-open -a src/crypthelper-open.go
	#go build -o bin/crypthelper-close -a src/crypthelper-close.go
