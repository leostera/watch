.PHONY: all test build

PREFIX ?= /usr/local
GO = $(shell which go)
SED = $(shell which sed)

all: vet format build test benchmark

vet:
	$(GO) vet

format:
	$(GO) fmt

benchmark:
	$(GO) test -bench .

version:
	$(SED) -i'' 's =.* =\ "$(shell git describe)" ' version.go

build: version
	$(GO) build

test:
	$(GO) test

install: build
	install ./watch $(PREFIX)/bin/watch

uninstall:
	rm -rf $(PREFIX)/bin/watch
