.PHONY: all test build

GO = $(shell which go)

all: build test benchmark

benchmark:
	$(GO) test -bench .

build:
	$(GO) build

test:
	$(GO) test
