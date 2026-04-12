all: build
.PHONY: all build test clean graph

VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

build:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/buildsim ./cmd/main.go

test:
	@go test ./... -race -count=1

clean:
	@rm -rf ./bin

run: build
	@./bin/buildsim start --port 9090

run-edit: build
	@./bin/buildsim start --port 9090 --edit
