#!/usr/bin/env bash

set -e

swag init -q -g ./cmd/server/main.go --parseInternal --parseDependency
echo -e "==> swag init -q -g ./cmd/server/main.go --parseInternal --parseDependency \e[42mOK\e[0m"

go build -o ./bin/server ./cmd/server
echo -e "==> go build -o ./bin/server ./cmd/server \e[42mOK\e[0m"

golint -set_exit_status=1 ./...
echo -e "==> golint -set_exit_status=1 ./... \e[42mOK\e[0m"

golangci-lint run
echo -e "==> golangci-lint run \e[42mOK\e[0m"

gosec -quiet -fmt=golint ./...
echo -e "==> gosec -quiet -fmt=golint ./... \e[42mOK\e[0m"
