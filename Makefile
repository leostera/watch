.PHONY: all ci setup lint benchmark build test install uninstall

PREFIX ?= /usr/local
GO = $(shell which go)

all: lint build test benchmark

ci: setup all

setup:
	go get -u github.com/alecthomas/gometalinter
	$(shell echo $$GOPATH/bin/gometalinter) --install

lint:
	$(shell echo $$GOPATH/bin/gometalinter) @.gometalinter
	$(GO) vet

benchmark:
	$(GO) test -bench .

build:
	$(GO) build -o ./watch

test:
	$(GO) test

install: build
	install watch $(PREFIX)/bin/watch

uninstall:
	rm -rf $(PREFIX)/bin/watch
