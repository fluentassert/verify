.DEFAULT_GOAL := all

.PHONY: all
all: ## build pipeline
all: install build fmt lint test

.PHONY: install
install: ## go install tools
	$(call print-target)
	cd tools && go install $(shell cd tools && go list -f '{{ join .Imports " " }}' -tags=tools)

.PHONY: build
build: ## go build
	$(call print-target)
	go build ./...

.PHONY: fmt
fmt: ## go fmt
	$(call print-target)
	go fmt ./...

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	golangci-lint run

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
