#!/bin/bash

export GOPATH=$(pwd)
go clean
go build -o bin/crypthelper-open -a src/crypthelper-open.go
#go build -o bin/crypthelper-close -a src/crypthelper-close.go
