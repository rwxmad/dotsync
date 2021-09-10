BUILDPATH=$(CURDIR)
GO=$(shell which go)
GOINSTALL=$(GO) install
GOCLEAN=$(GO) clean

EXENAME=main

export PATH := $(abspath bin/):${PATH}

build:
	go build -o ./bin/

.PHONY: clean
clean:
	rm -rf bin/

fmt:
	gofmt -l -s -w .
