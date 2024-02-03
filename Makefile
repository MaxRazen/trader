SHELL=/bin/bash

help: ## Print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Test applications
	go test -v ./...

build: ## Build applications
	go build -o ./.bin/ ./backtest/backtest.go

run-app: ## Build and run App
	go run ./trader.go

drop-builds: ## Remove all builds
	rm -f ./.builds/trader
