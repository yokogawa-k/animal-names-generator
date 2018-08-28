SHELL := /bin/bash

CURRENT_REVISION = $(shell git rev-parse --short HEAD)

SRC := $(shell find . -not -path "*/vendor/*" -name "*.go")
BIN := animal-names-generator

all: fmt lint test

build: $(SRC)
	go build -ldflags="-s -w -X main.COMMIT=$(CURRENT_REVISION)" -o $(BIN) ./cmd/...

fmt:
	gofmt -s -w $(SRC)

test: testdeps
	go test $(VERBOSE_FLAG) ./...

lint: testdeps
	go vet ./...
	golint -set_exit_status ./...

testdeps:
	go get -d -v -t ./...
	go get github.com/golang/lint/golint

release_deps:
	go get github.com/mitchellh/gox
	go get github.com/axw/gocov/gocov
	go get github.com/mattn/goveralls
	go get golang.org/x/tools/cover

cross_build:
	gox --osarch "darwin/amd64 linux/386 linux/amd64 windows/amd64" \
		-output "dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-ldflags="-s -w -X main.COMMIT=$(CURRENT_REVISION)" ./cmd/...
.PHONY: all build fmt test lint testdeps release_deps cross_build
