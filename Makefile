.DEFAULT_GOAL := help

GOCMD := go
GOMOD := $(GOCMD) mod
GOINSTALL := $(GOCMD) install
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOFMT := $(GOCMD) fmt
NAME := $(shell basename $(shell pwd))
BUILDDIR := ./build
BINDIR := $(BUILDDIR)/bin

VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

## Install dependencies
.PHONY: deps
deps:
	$(GOMOD) download

## Install dependencies for develop
.PHONY: devel-deps
devel-deps: deps
	$(GOINSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2
	$(GOINSTALL) github.com/Songmu/make2help/cmd/make2help@v0.2.1
	$(GOINSTALL) golang.org/x/tools/cmd/goimports@latest

## Run tests
.PHONY: test
test: deps
	$(GOTEST) -v ./...

## Format source codes
.PHONY: fmt
fmt: deps
	$(GOFMT) ./...

## Lint
.PHONY: lint
lint: devel-deps
	golangci-lint run ./...


## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)