#!/bin/bash

# delete old swagger
rm -f .docs/swagger/swagger.json

# generate new swagger
go generate

# validate the swagger
swagger validate .docs/swagger/swagger.json