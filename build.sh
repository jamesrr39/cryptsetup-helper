#!/bin/bash

export GOPATH=$(pwd)
go build src/crypthelper-open.go
go build src/crypthelper-close.go
