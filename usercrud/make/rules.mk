# This file defines GNU Make targets

GOPATH=$(shell go env GOPATH)
export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

all: build

.PHONY: modules
modules:
	go mod tidy

# Builds Go binaries
.PHONY: build
build: modules
	bin/run-go-build .

