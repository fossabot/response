-include .env

VERSION := $(shell git describe --tags --always)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME = response

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := "${GOBASE}/.bin/${PROJECTNAME}"
GOFILES := $(wildcard *.go)

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

# Redirect error output to a file, so we can show it in development mode.
STDERR := /tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the running server
PID := /tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

.DEFAULT_GOAL := all

## setup: Setup the project for local development. Thanks for contributing!
setup:
	@echo "  >  Initialiing for local development, thanks for your contribution!"
	@$(MAKE) go-clean go-get

## exec: Run given command, wrapped with custom GOBIN. e.g; make exec run="go test ./..."
exec:
	@GOBIN=$(GOBIN) $(run)

## run: Builds and runs the Response binary with any additional args. e.g; make run serve
run: go-build
	@$(GOBIN)/response $(RUN_ARGS)

## go-clean: Clean build files. Runs `go clean` internally.
go-clean:
	@echo "  >  Cleaning project .bin/"
	@rm -rf $(GOBASE)/.bin
	@GOBIN=$(GOBIN) go clean

## go-install: installs the binary in your GOPATH for testing outside of this directory.
go-install:
	@echo "  >  Installing the Response binary..."
	@go install $(GOFILES)

## go-get: Perform a `go get` in the Response project.
go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOBIN=$(GOBIN) go get $(get)

## go-tidy: Cleans up unused dependencies. This internally calls `go mod tidy`.
go-tidy:
	@echo "  >  Tidying up unused packages"
	@GOBIN=$(GOBIN) go mod tidy

## go-build: Perform a build for the current architecture.
go-build:
	@echo "  >  Building binary..."
	@GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/response main.go

## go-test: Run application unit and integration tests.
go-test:
	@echo "  >  Testing Response..."
	@GOBIN=$(GOBIN) go test -v ./...

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " The following targets are available for the "$(PROJECTNAME)" project:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo