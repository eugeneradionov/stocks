BIN_NAME := $(or $(PROJECT_NAME),'api')

MIGRATE=migrate -path sql/migrations -database postgres://postgres:12345@localhost:5432/stocks?sslmode=disable
TEST_MIGRATE=migrate -path sql/migrations -database postgres://postgres:12345@localhost:5432/stocks-test?sslmode=disable

.PHONY: all dep lint build

all: dep lint build

dep: ## Download required dependencies
	GO111MODULE=on go mod vendor

lint: ## Lint all files in the project
	@golangci-lint run -c .golangci.yml

build: dep ## Build the binary file
	GO111MODULE=on go build -o ../bin/${BIN_NAME} -a -tags netgo -ldflags '-w -extldflags "-static"' .

migrate-create: ## Create migration file with name
	migrate create -ext sql -dir sql/migrations migration

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down
