#!/bin/bash

export GOPATH=$(pwd)
go clean
go build -o bin/crypthelper-open src/crypthelper-open.go
go build -o bin/crypthelper-close src/crypthelper-close.go
