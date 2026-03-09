package disassemble

import (
	"encoding/binary"
	"testing"
)

func TestDecomposeIntoEquivalence(t *testing.T) {
	var results [1024]byte
	for _, word := range benchWords {
		old, err := Decompose(0x1000, word, &results)
		if err != nil {
			t.Fatalf("Decompose(%#x) failed: %v", word, err)
		}

		var inst Inst
		if err := DecomposeInto(0x1000, word, &inst); err != nil {
			t.Fatalf("DecomposeInto(%#x) failed: %v", word, err)
		}

		if inst.Operation != old.Operation {
			t.Errorf("word %#x: Operation mismatch: %v vs %v",
				word, inst.Operation, old.Operation)
		}
		if inst.Raw != old.Raw {
			t.Errorf("word %#x: Raw mismatch: %#x vs %#x",
				word, inst.Raw, old.Raw)
		}
		if int(inst.NumOps) != len(old.Operands) {
			t.Errorf("word %#x: operand count mismatch: %d vs %d",
				word, inst.NumOps, len(old.Operands))
			continue
		}
		for i := uint8(0); i < inst.NumOps; i++ {
			if inst.Operands[i].Class != old.Operands[i].Class {
				t.Errorf("word %#x op[%d]: Class mismatch",
					word, i)
			}
			if inst.Operands[i].Immediate != old.Operands[i].Immediate {
				t.Errorf("word %#x op[%d]: Immediate mismatch",
					word, i)
			}
			if int(inst.Operands[i].NumRegisters) !=
				len(old.Operands[i].Registers) {
				t.Errorf("word %#x op[%d]: register count mismatch: %d vs %d",
					word, i,
					inst.Operands[i].NumRegisters,
					len(old.Operands[i].Registers))
			}
		}
	}
}

func TestDecomposeBatch(t *testing.T) {
	n := len(benchWords)
	addrs := make([]uint64, n)
	for i := range addrs {
		addrs[i] = 0x1000 + uint64(i)*4
	}
	out := make([]Inst, n)

	decoded, err := DecomposeBatch(addrs, benchWords, out)
	if err != nil {
		t.Fatal(err)
	}
	if decoded != n {
		t.Fatalf("decoded %d, want %d", decoded, n)
	}

	for i, inst := range out[:decoded] {
		if inst.Address != addrs[i] {
			t.Errorf("[%d] address mismatch: %#x vs %#x",
				i, inst.Address, addrs[i])
		}
		if inst.Operation == ARM64_ERROR {
			t.Errorf("[%d] operation is ERROR for word %#x",
				i, benchWords[i])
		}
	}
}

// Representative instruction words for benchmarks
var benchWords = []uint32{
	// ADRP X0, #0x1000
	binary.LittleEndian.Uint32([]byte{0x00, 0x08, 0x00, 0x90}),
	// ADD X0, X0, #16
	binary.LittleEndian.Uint32([]byte{0x00, 0x40, 0x00, 0x91}),
	// LDR X1, [X0, #8]
	binary.LittleEndian.Uint32([]byte{0x01, 0x04, 0x40, 0xf9}),
	// STR X1, [SP, #16]
	binary.LittleEndian.Uint32([]byte{0xe1, 0x0b, 0x00, 0xf9}),
	// BL #0x100
	binary.LittleEndian.Uint32([]byte{0x40, 0x00, 0x00, 0x94}),
	// MOV X0, X1
	binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0x01, 0xaa}),
	// RET
	binary.LittleEndian.Uint32([]byte{0xc0, 0x03, 0x5f, 0xd6}),
	// STP X29, X30, [SP, #-16]!
	binary.LittleEndian.Uint32([]byte{0xfd, 0x7b, 0xbf, 0xa9}),
	// LDP X29, X30, [SP], #16
	binary.LittleEndian.Uint32([]byte{0xfd, 0x7b, 0xc1, 0xa8}),
	// CBNZ X0, #0x10
	binary.LittleEndian.Uint32([]byte{0x80, 0x00, 0x00, 0xb5}),
}

func BenchmarkDecompose(b *testing.B) {
	var results [1024]byte
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, word := range benchWords {
			_, _ = Decompose(0x1000, word, &results)
		}
	}
}

func BenchmarkDisassemble(b *testing.B) {
	var results [1024]byte
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, word := range benchWords {
			_, _ = Disassemble(0x1000, word, &results)
		}
	}
}

func BenchmarkDecomposeInto(b *testing.B) {
	var inst Inst
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, word := range benchWords {
			_ = DecomposeInto(0x1000, word, &inst)
		}
	}
}

func BenchmarkOperationString(b *testing.B) {
	ops := []Operation{
		ARM64_ADRP, ARM64_ADD, ARM64_LDR, ARM64_STR,
		ARM64_BL, ARM64_MOV, ARM64_RET, ARM64_STP,
		ARM64_LDP, ARM64_CBNZ,
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, op := range ops {
			_ = op.String()
		}
	}
}

func BenchmarkDecomposeBatch(b *testing.B) {
	n := len(benchWords)
	addrs := make([]uint64, n)
	for i := range addrs {
		addrs[i] = 0x1000 + uint64(i)*4
	}
	out := make([]Inst, n)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = DecomposeBatch(addrs, benchWords, out)
	}
}
