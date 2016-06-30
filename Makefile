.PHONY: all ci setup lint benchmark build test install uninstall

PREFIX ?= /usr/local
GO = $(shell which go)
VERSION="$(shell git describe HEAD)"
NAME=$(shell basename $(shell pwd))

all: lint build test benchmark

ci: setup all

setup:
	go get -u github.com/ostera/oh-my-gosh/lib
	go get -u github.com/alecthomas/gometalinter
	$(shell echo $$GOPATH/bin/gometalinter) --install

lint:
	$(shell echo $$GOPATH/bin/gometalinter) @.gometalinter
	$(GO) vet

benchmark:
	$(GO) test -bench .

build:
	$(GO) build --ldflags "-X main.Version=$(VERSION)" -o $(NAME)

release:
	./release.sh

test:
	$(GO) test

install: build
	install $(NAME) $(PREFIX)/bin/$(NAME)

uninstall:
	rm -rf $(PREFIX)/bin/$(NAME)
