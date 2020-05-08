.PHONY: dep lint

all: dep lint

dep: # Download required dependencies
	GO111MODULE=on go mod vendor

lint: ## Lint the files local env
	@golangci-lint run -c .golangci.yml

