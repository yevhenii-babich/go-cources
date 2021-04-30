#!/bin/sh
go mod vendor -v
docker build --tag test-app .
rm -rf vendor
docker run -p 8081:8081 --rm test-app
