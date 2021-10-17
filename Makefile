REPO=blacktop
NAME=arm64-cgo
CLI=github.com/blacktop/arm64-cgo/cmd/disass
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)


.PHONY: build-deps
build-deps: ## Install the build dependencies
	@echo " > Installing build deps"
	brew install go goreleaser
	go get -u github.com/crazy-max/xgo
	go install golang.org/x/tools/cmd/stringer

.PHONY: build
build: ## Build disass locally
	@echo " > Building locally"
	CGO_ENABLED=1 go build -o disass.${CUR_VERSION} ./cmd/disass

.PHONY: test
test: build ## Test disass on hello-mte
	@echo " > disassembling hello-mte\n"
	@./disass.${CUR_VERSION} ../../Proteas/hello-mte/hello-mte --all | bat -l s --tabs 0 -p --theme Nord --wrap=never

.PHONY: dry_release
dry_release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo " > Creating Pre-release Build ${NEXT_VERSION}"
	# @goreleaser build --rm-dist --skip-validate --snapshot --id darwin #-f .goreleaser.mac.yml
	@goreleaser release --skip-validate --skip-publish

.PHONY: release
release: ## Create a new release from the NEXT_VERSION
	@echo " > Creating Release ${NEXT_VERSION}"
	@hack/make/release ${NEXT_VERSION}
	@goreleaser #--rm-dist

.PHONY: cross
cross: ## Create xgo releases
	@echo " > Creating xgo releases"
	@mkdir -p dist/xgo
	@cd dist/xgo; xgo --targets=*/amd64 -go 1.16.5 -ldflags='-s -w' -out disass-${NEXT_VERSION} ${CLI}

clean: ## Clean up artifacts
	@echo " > Cleaning"
	rm -rf dist
	rm disass.v* || true

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help