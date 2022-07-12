include .env

SHELL=bash
GOPATH=$(shell go env | grep GOPATH | cut -d= -f2 | xargs echo)
PACKAGES=$(shell go list ./... | grep -v /vendor/ | grep -v mocks)
IMAGE_NAME=go-rest-api
CONFIRM=0
VERSION=1.0.0

.PHONY: build vendor test

help:
	@grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

__check-tools:
	@for TOOL in $(TOOLS); do \
		type $$TOOL &> /dev/null && ([ $(CONFIRM) -eq 1 ] && echo "$$TOOL: OK" || true) || (echo "$$TOOL: MISSING"; exit 1); \
	done

check-tools: ## Check if required tools are installed
	@$(MAKE) __check-tools CONFIRM=1

vendor: ## Install dependencies
	@go mod download
	@go mod vendor

build: vendor ## Build the application
	@rm -rf ./build
	@mkdir -p ./build
	@go build -mod=vendor -o ./build/go-rest-api -ldflags "-X go-rest-api/internal.version=$(VERSION)" ./cmd/...

run: vendor ## Run the application in development mode
	@go run -mod=vendor ./cmd

lint: ## Run linters
	@$(GOPATH)/bin/revive -formatter friendly $(PACKAGES)
	@go vet $(PACKAGES)

test: ## Run unit tests
	@go test -timeout 10s ./...

clean: ## Cleanup
	@rm -rm ./build
	@rm -rf ./vendor

docker-build: ## Build docker image
	@docker build -t go-rest-api --build-arg VERSION=$(VERSION)

docker-run: ## Runs the image locally
	@docker run --name go-rest-api -d do-rest-api:$(VERSION)
