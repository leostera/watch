.PHONY: all test build

PREFIX ?= /usr/local
GO = $(shell which go)

all: vet format build test benchmark

vet:
	$(GO) vet

format:
	$(GO) fmt

benchmark:
	$(GO) test -bench .

build:
	$(GO) build

test:
	$(GO) test

install: build
	install watch $(PREFIX)/bin/watch

uninstall:
	rm -rf $(PREFIX)/bin/watch
