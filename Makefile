BIN             = pid-plot
OUTPUT_DIR      = build
RELEASE_VER    := $(shell git rev-parse --short HEAD)
COVERMODE       = atomic

TEST_PACKAGES      := $(shell go list ./... | grep -v vendor | grep -v fakes | grep -v ftest)

.PHONY: help
.DEFAULT_GOAL := help

run: ## Run application (without building)
	go run *.go -d

setup: installtools ## Install and setup tools


## Run Tests ##

test: ## Perform unit tests (with verbose flag)
	go test -v -cover $(TEST_PACKAGES)

## Alternate test types not included in 'make test'

test/race: ## Perform unit tests and enable the race detector
	go test -race -cover $(TEST_PACKAGES)

test/cover: ## Run all tests + open coverage report for all packages
	echo 'mode: $(COVERMODE)' > .coverage
	for PKG in $(TEST_PACKAGES); do \
		go test -coverprofile=.coverage.tmp -tags "integration" $$PKG; \
		grep -v -E '^mode:' .coverage.tmp >> .coverage; \
	done
	go tool cover -html=.coverage
	$(RM) .coverage .coverage.tmp

installtools: ## Install development related tools
	go get github.com/kardianos/govendor
	go get -u github.com/jteeuwen/go-bindata/...

## Build ##

build: clean build/linux build/darwin ## Build for linux and darwin (save to OUTPUT_DIR/BIN)

build/linux: clean/linux ## Build for linux (save to OUTPUT_DIR/BIN)
	GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=$(RELEASE_VER)" -o $(OUTPUT_DIR)/$(BIN)-linux .

build/darwin: clean/darwin ## Build for darwin (save to OUTPUT_DIR/BIN)
	GOOS=darwin go build -a -installsuffix cgo -ldflags "-X main.version=$(RELEASE_VER)" -o $(OUTPUT_DIR)/$(BIN)-darwin .

clean: clean/darwin clean/linux ## Remove all build artifacts

clean/darwin: ## Remove darwin build artifacts
	$(RM) $(OUTPUT_DIR)/$(BIN)-darwin

clean/linux: ## Remove linux build artifacts
	$(RM) $(OUTPUT_DIR)/$(BIN)-linux

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_\/-]+:.*?## / {printf "\033[34m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | \
		sort | \
		grep -v '#'
