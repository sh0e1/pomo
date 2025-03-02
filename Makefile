SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := help

BIN_DIR := $(shell pwd)/bin

GOLANGCI_LINT         := $(abspath $(BIN_DIR)/golangci-lint)
GOLANGCI_LINT_VERSION := v1.63.4

.PHONY: lint
lint: ## Run golangci-lint after installation if golangci-lint is not installed
	@if ! type $(GOLANGCI_LINT) > /dev/null 2>&1; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_LINT_VERSION); \
	fi
	$(GOLANGCI_LINT) run ./...

.PHONY: test
test: ## Run tests with coverage
	go test -v -race -count=1 -cover ./...

.PHONY: help
help: ## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-30s %s\n", $$1, $$2}'
