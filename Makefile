#! /usr/bin/make -f

# ---------------------------------------------------------
# Go related variables.
# ---------------------------------------------------------
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# ---------------------------------------------------------
# Go files.
# ---------------------------------------------------------
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

# ---------------------------------------------------------
# Code Checking
# ---------------------------------------------------------
check: fmt lint unit-test swag

unit-test:
	@echo "  >  Running unit tests"
	go clean -testcache
	GOBIN=$(GOBIN) go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./internal... ./pkg...

fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

lint-install:
ifeq ("$(wildcard bin/golangci-lint)","")
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
endif

lint: lint-install
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=3m
