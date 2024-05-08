help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ## Quick build and run
	go run -ldflags="-X main.mode=watch" .

build: ## Build a release version
	@$(MAKE) --no-print-directory clean
	npm run prod
	go build .

clean: ## Clean generated files, logs, caches
	@rm -rf public/*
	@rm -rf tutor
