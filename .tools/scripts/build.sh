#!/bin/bash

# clean out existing build artifacts
rm -rf bin

# gather dependencies
dep ensure -v

# build binary
go build -o bin/skatepark-api-go