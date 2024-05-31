#!/bin/bash

swag init -g cmd/main.go -d ./,./cmd -o api -ot yaml

openapi-generator-cli generate -i api/swagger.yaml -o api -g openapi-yaml --additional-properties=x-extension.openapi=3.0.0,outputFile=openapi.yaml
