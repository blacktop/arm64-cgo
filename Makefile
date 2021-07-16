all:
	c-for-go -ccdefs disassembler.yml

clean:
	rm -f disassembler/cgo_helpers.go disassembler/cgo_helpers.h disassembler/cgo_helpers.c
	rm -f disassembler/const.go disassembler/doc.go disassembler/types.go
	rm -f disassembler/disassembler.go

test:
	cd disassembler && go build
