.PHONY: dep lint

all: dep lint

dep: # Download required dependencies
	GO111MODULE=on go mod vendor

lint: ## Lint all files in the project
	@golangci-lint run -c .golangci.yml

