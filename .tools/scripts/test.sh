#!/bin/bash

# execute golang tests
go test ./... -covermode=count -v -coverprofile cp.out

# code coverage print
go tool cover -func=cp.out