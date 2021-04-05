SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

all: fmt
	echo 'make all'
fmt:
	gofmt -w .
mod:
	go mod tidy
lint:
	golangci-lint run
.PHONY: test
test:
	sh scripts/test.sh
help:
	@echo "fmt - go format"
	@echo "mod - go mod tidy"
	@echo "lint - run golangci-lint"
	@echo "test - unit test"