.DEFAULT_GOAL := all

.PHONY: all
all: ## build pipeline
all: test lint misspell

.PHONY: test
test: ## go test with race detector and code covarage
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

.PHONY: misspell
misspell: ## misspell
	$(call print-target)
	cd tools && go install github.com/client9/misspell/cmd/misspell
	misspell -error -locale US *.md

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
