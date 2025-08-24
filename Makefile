REPO=blacktop
NAME=arm64-cgo
CLI=github.com/blacktop/arm64-cgo/cmd/disass
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)
BIN ?= disass

.PHONY: bump
bump:
	@echo "🚀 Bumping Version"
	git tag $(shell svu patch)
	git push --tags

.PHONY: build
build:
	@echo "🚀 Building Version $(shell svu current)"
	cd cmd/$(BIN) && go build -o ../../$(BIN) .

.PHONY: test
test: ## Run unit tests
	@echo " > Running unit tests"
	@go test -v ./...

.PHONY: test-mte
test-mte: build ## Test disass on hello-mte
	@echo " > disassembling hello-mte\n"
	@$(BIN) ../../Proteas/hello-mte/hello-mte --all | bat -l s --tabs 0 -p --theme Nord --wrap=never

.PHONY: release
release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo "🚀 Releasing Version $(shell svu current)"
	goreleaser build --id default --clean --snapshot --single-target --output dist/$(BIN)

.PHONY: snapshot
snapshot:
	@echo "🚀 Snapshot Version $(shell svu current)"
	goreleaser --clean --timeout 60m --snapshot

clean: ## Clean up artifacts
	@echo " > Cleaning"
	rm -rf dist/ completions/

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help