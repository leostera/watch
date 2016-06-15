.PHONY: all test build

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
