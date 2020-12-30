
GOBIN ?= go

help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
.PHONY: help

largetype: ## build and run largetype example
	$(GOBIN) run ./examples/largetype -font="Monaco" "Hello world"
.PHONY: largetype

topframe: ## build and run topframe example
	$(GOBIN) run ./examples/topframe
.PHONY: topframe