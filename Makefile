.PHONY: all test build

GO = $(shell which go)

all: build test

build:
	$(GO) build

test:
	$(GO) test
