.PHONY: all test build

GO = $(shell which go)

all: vet format build test benchmark

vet:
	$(GO) vet

format:
	$(GO) fmt

benchmark:
	$(GO) test -bench .

version:
	sed -i 's =.* =\ "$(shell git describe)" ' version.go

build: version
	$(GO) build

test:
	$(GO) test
