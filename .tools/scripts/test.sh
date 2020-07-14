#!/bin/bash

# execute golang tests
go test ./... -covermode=count -v -coverprofile cp.out

# go get cobertura lib
go get github.com/t-yuki/gocover-cobertura

# generate output files
go tool cover -html=cp.out -o cp.html && gocover-cobertura < cp.out > cp.xml

# code coverage print
go tool cover -func=cp.out