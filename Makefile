build:
	CGO_ENABLED=1 go build -o disass

all:
	c-for-go -debug disassembler.yml

clean:
	rm -f disassembler/cgo_helpers.go disassembler/cgo_helpers.h disassembler/cgo_helpers.c
	rm -f disassembler/const.go disassembler/doc.go disassembler/types.go
	rm -f disassembler/disassembler.go

test:
	cd disassembler && go build
