#!/bin/bash
export GOPATH=`pwd`
go vet ./... && env GOMAXPROCS=4 go test -v ./...
