default: help


help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


run: ## Quick build and run
	go run -ldflags="-X main.mode=devonly" .

test: ## Run golang tests
	@go test ./internal/utils

build: ## Build a release version
	@$(MAKE) --no-print-directory clean
	$(MAKE) --no-print-directory ui-build
	go build -o .dist/tutor


.PHONY: ui-watch
ui-watch:
	@$(MAKE) --no-print-directory clean
	@cd ui && npm run watch


.PHONY: ui-build
ui-build:
	@$(MAKE) --no-print-directory clean
	@cd ui && npm run prod


.PHONY: clean
clean: ## Clean generated files, logs, caches
	@rm -rf ui/public/*
	@rm -rf .dist

.PHONY: init
init: ## Clean generated files, logs, caches
	@mkdir ./credentials
	@cp env.example ./credentials/env
	@touch ./credentials/gcp.json
