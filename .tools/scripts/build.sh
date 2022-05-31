#!/bin/bash

# clean out existing build artifacts
rm -rf bin

# gather dependencies
go mod download
go mod vendor

# build binary
go build -o bin/skatepark-api-go