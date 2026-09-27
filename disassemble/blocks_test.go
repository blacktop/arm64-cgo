package disassemble_test

import (
	"encoding/binary"
	"slices"
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

const (
	opAdd    = 0x91000420 // add x0, x1, #1
	opB      = 0x14000002 // b 0x8
	opBEq    = 0x54000040 // b.eq 0x8
	opBL     = 0x94000004 // bl 0x10
	opCBNZ   = 0xb5000040 // cbnz x0, 0x8
	opRet    = 0xd65f03c0 // ret
	opTexit  = 0xd6ff03e0 // texit
	opTenter = 0xd4e00000 // tenter #0
	opGexit  = 0x00201400 // gexit
	opGenter = 0x00201420 // genter #0
)

func TestBlocks(t *testing.T) {
	tests := []struct {
		name       string
		words      []uint32
		wantBlocks []int // number of instructions in each block
	}{
		{"branch and return end blocks", []uint32{opAdd, opB, opAdd, opRet}, []int{2, 2}},
		{"conditional branch ends a block", []uint32{opBEq, opAdd}, []int{1, 1}},
		{"call ends a block", []uint32{opBL, opAdd}, []int{1, 1}},
		{"call stays with a CBNZ checking its result", []uint32{opBL, opCBNZ, opAdd}, []int{2, 1}},
		{"texit ends a block even before CBNZ", []uint32{opTexit, opCBNZ, opAdd}, []int{1, 1, 1}},
		{"gexit ends a block", []uint32{opGexit, opAdd}, []int{1, 1}},
		{"tenter and genter return to the next instruction", []uint32{opTenter, opGenter, opAdd}, []int{3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]byte, 0, len(tt.words)*4)
			for _, w := range tt.words {
				data = binary.LittleEndian.AppendUint32(data, w)
			}
			instrs, err := disassemble.GetInstructions(0, data)
			if err != nil {
				t.Fatalf("GetInstructions error = %v", err)
			}

			var got []int
			for _, block := range instrs.Blocks() {
				got = append(got, len(block))
			}
			if !slices.Equal(got, tt.wantBlocks) {
				t.Errorf("block sizes = %v, want %v", got, tt.wantBlocks)
			}
		})
	}
}
