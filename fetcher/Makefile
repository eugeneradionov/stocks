BIN_NAME := $(or $(PROJECT_NAME),'fetcher')

.PHONY: all dep lint build

all: dep lint build

dep: ## Download required dependencies
	GO111MODULE=on go mod vendor

lint: ## Lint all files in the project
	@golangci-lint run -c .golangci.yml


build: dep ## Build the binary file
	GO111MODULE=on go build -o ../bin/${BIN_NAME} -a -tags netgo -ldflags '-w -extldflags "-static"' .
