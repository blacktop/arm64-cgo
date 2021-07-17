REPO=blacktop
NAME=arm64-cgo
CUR_VERSION=$(shell svu current)
NEXT_VERSION=$(shell svu patch)

build:
	CGO_ENABLED=1 go build -o disass

all:
	CGO_ENABLED=1 c-for-go -debug disassembler.yml

clean:
	rm -f disassembler/cgo_helpers.go disassembler/cgo_helpers.h disassembler/cgo_helpers.c
	rm -f disassembler/const.go disassembler/doc.go disassembler/types.go
	rm -f disassembler/disassembler.go

test:
	cd disassembler && go build

.PHONY: dry_release
dry_release: ## Run goreleaser without releasing/pushing artifacts to github
	@echo " > Creating Pre-release Build ${NEXT_VERSION}"
	@goreleaser build --rm-dist --skip-validate --snapshot

cross:
	docker run --rm --privileged \
		-v $PWD:/go/src/github.com/blacktop/arm64-cgo \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /go/src/github.com/blacktop/arm64-cgo \
		mailchain/goreleaser-xcgo --snapshot --rm-dist
