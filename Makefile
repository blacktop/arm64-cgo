REPO=blacktop
NAME=arm64-cgo
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)


build:
	CGO_ENABLED=1 go build -o disass

all:
	CGO_ENABLED=1 c-for-go -debug disassembler.yml

.PHONY: test
test: ## Run disass on hello-mte
	@echo " > disassembling hello-mte\n"
	@dist/arm64-cgo_darwin_amd64/disass  ../../Proteas/hello-mte/hello-mte _test

.PHONY: dry_release
dry_release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo " > Creating Pre-release Build ${NEXT_VERSION}"
	@goreleaser build --rm-dist --skip-validate --snapshot

.PHONY: release
release: ## Create a new release from the NEXT_VERSION
	@echo " > Creating Release ${NEXT_VERSION}"
	@hack/make/release ${NEXT_VERSION}

clean: ## Clean up artifacts
	@echo " > Cleaning"
	rm *.tar || true
	rm *.ipsw || true
	rm kernelcache.release.* || true
	rm -rf dist

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help