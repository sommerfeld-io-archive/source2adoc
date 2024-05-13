#!/bin/bash

##
# The `local-build.sh` script is a simple shell script that formats the Go code, runs the tests,
# and builds the Go app. It does not rely on the Dockerfile from this repository. It is intended
# to build the app locally on your machine without spinning up the whole Docker environment. So
# this script does not run any linters or tests other than unit tests. Make sure to run the whole
# Docker-based configuration before pushing your changes to the repository.
#
# @see docker-compose.yml
##

echo "[INFO] Setting up the Go environment ..."
go mod download
go mod tidy

echo "[INFO] Formatting the Go code ..."
go fmt ./...

echo "[INFO] Running unit tests ..."
go test ./...

echo "[INFO] Building the Go app ..."
go build .
