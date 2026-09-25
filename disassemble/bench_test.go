package disassemble

import (
	"encoding/binary"
	"errors"
	"math"
	"reflect"
	"strconv"
	"testing"
)

func newBenchAddrs() []uint64 {
	addrs := make([]uint64, len(benchWords))
	for i := range addrs {
		addrs[i] = 0x1000 + uint64(i)*4
	}
	return addrs
}

func assertBatchDecoded(
	t *testing.T, decoded int,
	addrs []uint64, words []uint32, out []Inst,
) {
	t.Helper()

	want := len(addrs)
	if decoded != want {
		t.Fatalf("decoded %d, want %d", decoded, want)
	}

	for i, inst := range out[:decoded] {
		if inst.Address != addrs[i] {
			t.Errorf("[%d] address mismatch: %#x vs %#x",
				i, inst.Address, addrs[i])
		}
		if inst.Operation == ARM64_ERROR {
			t.Errorf("[%d] operation is ERROR for word %#x",
				i, words[i])
		}
	}
}

// assertInstEquivalence checks that an Inst matches the legacy
// *Instruction output for the same word, including operand fields.
func assertInstEquivalence(
	t *testing.T, word uint32, inst *Inst, old *Instruction,
) {
	t.Helper()
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
		return
	}
	for i := uint8(0); i < inst.NumOps; i++ {
		if inst.Operands[i].Class != old.Operands[i].Class {
			t.Errorf("word %#x op[%d]: Class mismatch", word, i)
		}
		if inst.Operands[i].Immediate != old.Operands[i].Immediate {
			t.Errorf("word %#x op[%d]: Immediate mismatch", word, i)
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
		assertInstEquivalence(t, word, &inst, old)
	}
}

// TestDecomposeIntoResetsReusedInst decodes into an Inst that holds a
// previous decode with more registers per operand.
func TestDecomposeIntoResetsReusedInst(t *testing.T) {
	var decoder Decoder
	for _, pair := range [][2]uint32{
		{benchWords[1], benchWords[4]}, // ADD X0, X0, #16 then BL
		{benchWords[7], 0xd503201f},    // STP X29, X30, [SP, #-16]! then NOP
	} {
		var reused, fresh Inst
		if err := decoder.DecomposeInto(0x1000, pair[0], &reused); err != nil {
			t.Fatal(err)
		}
		if err := decoder.DecomposeInto(0x1004, pair[1], &reused); err != nil {
			t.Fatal(err)
		}
		if err := decoder.DecomposeInto(0x1004, pair[1], &fresh); err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(reused, fresh) {
			t.Errorf("decoding %#x into an Inst that held %#x left stale fields", pair[1], pair[0])
		}
	}
}

func TestDecomposeBatch(t *testing.T) {
	addrs := newBenchAddrs()
	out := make([]Inst, len(benchWords))

	decoded, err := DecomposeBatch(addrs, benchWords, out)
	if err != nil {
		t.Fatal(err)
	}
	assertBatchDecoded(t, decoded, addrs, benchWords, out)
}

func TestDecoderDecomposeIntoEquivalence(t *testing.T) {
	var results [1024]byte
	decoder := NewDecoder(0)
	for _, word := range benchWords {
		old, err := Decompose(0x1000, word, &results)
		if err != nil {
			t.Fatalf("Decompose(%#x) failed: %v", word, err)
		}
		var inst Inst
		if err := decoder.DecomposeInto(0x1000, word, &inst); err != nil {
			t.Fatalf("decoder.DecomposeInto(%#x) failed: %v", word, err)
		}
		assertInstEquivalence(t, word, &inst, old)
	}
}

func TestDecoderDecomposeBatch(t *testing.T) {
	addrs := newBenchAddrs()
	out := make([]Inst, len(benchWords))
	decoder := NewDecoder(len(benchWords))

	decoded, err := decoder.DecomposeBatch(addrs, benchWords, out)
	if err != nil {
		t.Fatal(err)
	}
	assertBatchDecoded(t, decoded, addrs, benchWords, out)
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
	addrs := newBenchAddrs()
	out := make([]Inst, len(benchWords))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = DecomposeBatch(addrs, benchWords, out)
	}
}

func BenchmarkDecoderDecomposeInto(b *testing.B) {
	var inst Inst
	decoder := NewDecoder(0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, word := range benchWords {
			_ = decoder.DecomposeInto(0x1000, word, &inst)
		}
	}
}

func BenchmarkDecoderDecomposeBatch(b *testing.B) {
	addrs := newBenchAddrs()
	out := make([]Inst, len(benchWords))
	decoder := NewDecoder(len(benchWords))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = decoder.DecomposeBatch(addrs, benchWords, out)
	}
}

// Words whose single decode fails, one per failure category.
const (
	wordUndefined   uint32 = 0x0c858ca9 // DECODE_STATUS_UNDEFINED; batch decoding once reported it as st2
	wordUnallocated uint32 = 0xffffffff // DECODE_STATUS_UNALLOCATED
	wordBadOperands uint32 = 0x1918f46c // DECODE_STATUS_ERROR_OPERANDS
	wordNop         uint32 = 0xd503201f
)

func batchAddrs(n int) []uint64 {
	addrs := make([]uint64, n)
	for i := range addrs {
		addrs[i] = 0x2000 + uint64(i)*4
	}
	return addrs
}

// checkBatchAgainstSingles compares a status batch with DecomposeInto on
// each word: successes must match exactly, failures must carry the same
// error and leave only the address in out.
func checkBatchAgainstSingles(
	t *testing.T, decoded int,
	addrs []uint64, words []uint32, out []Inst, status []DecodeStatus,
) {
	t.Helper()
	var decoder Decoder
	want := 0
	for i, word := range words {
		var single Inst
		err := decoder.DecomposeInto(addrs[i], word, &single)
		if err == nil {
			want++
			if !status[i].OK() {
				t.Errorf("word %#x: batch status %s, single decode succeeded", word, status[i])
			} else if !reflect.DeepEqual(out[i], single) {
				t.Errorf("word %#x: batch instruction differs from single decode", word)
			}
			continue
		}
		if status[i].OK() {
			t.Errorf("word %#x: batch reported success, single decode failed: %v", word, err)
			continue
		}
		if got := status[i].Err(word); got == nil || got.Error() != err.Error() {
			t.Errorf("word %#x: batch error %v, single error %v", word, got, err)
		}
		if !reflect.DeepEqual(out[i], Inst{Address: addrs[i]}) {
			t.Errorf("word %#x: failed decode left %v in out", word, out[i].Operation)
		}
	}
	if decoded != want {
		t.Errorf("decoded = %d, want %d successful decodes", decoded, want)
	}
}

func TestDecomposeBatchStatusMatchesSingleDecodes(t *testing.T) {
	decoder := NewDecoder(0)
	for _, tc := range []struct {
		name  string
		words []uint32
	}{
		{"undefined probe word", []uint32{wordUndefined}},
		{"mixed", []uint32{wordNop, wordUndefined, benchWords[0], wordUnallocated, benchWords[1], wordBadOperands, benchWords[2]}},
		{"failures at both ends", []uint32{wordUnallocated, benchWords[0], benchWords[1], wordUndefined}},
		{"all failures", []uint32{wordUndefined, wordUnallocated, wordBadOperands}},
		{"empty", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			addrs := batchAddrs(len(tc.words))
			out := make([]Inst, len(tc.words))
			status := make([]DecodeStatus, len(tc.words))
			decoded, err := decoder.DecomposeBatchStatus(addrs, tc.words, out, status)
			if err != nil {
				t.Fatal(err)
			}
			checkBatchAgainstSingles(t, decoded, addrs, tc.words, out, status)
		})
	}
	status := make([]DecodeStatus, 1)
	if _, err := DecomposeBatchStatus(batchAddrs(1), []uint32{wordUndefined}, make([]Inst, 1), status); err != nil {
		t.Fatal(err)
	}
	if got := status[0].String(); got != "DECODE_STATUS_UNDEFINED" {
		t.Errorf("probe word status = %s, want DECODE_STATUS_UNDEFINED", got)
	}
}

func TestDecomposeBatchStatusResetsReusedBuffers(t *testing.T) {
	decoder := NewDecoder(0)
	out := make([]Inst, 4)
	status := make([]DecodeStatus, 4)
	for _, words := range [][]uint32{
		{benchWords[0], benchWords[1], benchWords[2], benchWords[3]},
		{wordUndefined, benchWords[4], wordUnallocated, benchWords[5]},
		{wordBadOperands},
		{benchWords[6], wordUndefined, benchWords[7], wordNop},
	} {
		addrs := batchAddrs(len(words))
		decoded, err := decoder.DecomposeBatchStatus(addrs, words, out[:len(words)], status[:len(words)])
		if err != nil {
			t.Fatal(err)
		}
		checkBatchAgainstSingles(t, decoded, addrs, words, out[:len(words)], status[:len(words)])
	}
}

func TestDecomposeBatchCountsOnlySuccessfulDecodes(t *testing.T) {
	words := []uint32{wordUndefined, benchWords[0], wordUnallocated, benchWords[1]}
	addrs := batchAddrs(len(words))
	out := make([]Inst, len(words))
	decoded, err := NewDecoder(0).DecomposeBatch(addrs, words, out)
	if err != nil {
		t.Fatal(err)
	}
	if decoded != 2 {
		t.Errorf("DecomposeBatch decoded = %d, want 2", decoded)
	}
	for _, i := range []int{0, 2} {
		if !reflect.DeepEqual(out[i], Inst{Address: addrs[i]}) {
			t.Errorf("failed word %#x left %v in out", words[i], out[i].Operation)
		}
	}
}

func TestDecomposeBatchStatusRejectsMismatchedLengths(t *testing.T) {
	words := []uint32{wordNop, wordNop}
	for _, tc := range []struct {
		name   string
		addrs  []uint64
		out    []Inst
		status []DecodeStatus
	}{
		{"addrs", batchAddrs(1), make([]Inst, 2), make([]DecodeStatus, 2)},
		{"out", batchAddrs(2), make([]Inst, 1), make([]DecodeStatus, 2)},
		{"status", batchAddrs(2), make([]Inst, 2), make([]DecodeStatus, 3)},
	} {
		decoded, err := NewDecoder(0).DecomposeBatchStatus(tc.addrs, words, tc.out, tc.status)
		if !errors.Is(err, errBatchStatusLenMismatch) || decoded != 0 {
			t.Errorf("%s length mismatch: decoded %d, err %v", tc.name, decoded, err)
		}
	}
	if _, err := NewDecoder(0).DecomposeBatch(batchAddrs(1), words, make([]Inst, 2)); !errors.Is(err, errBatchLenMismatch) {
		t.Errorf("DecomposeBatch length mismatch: err %v", err)
	}
}

func TestCheckBatchLen(t *testing.T) {
	if strconv.IntSize == 32 {
		t.Skip("int cannot exceed the C int range")
	}
	if err := checkBatchLen(math.MaxInt32); err != nil {
		t.Errorf("checkBatchLen(MaxInt32) = %v", err)
	}
	over := int64(math.MaxInt32)
	over++
	if err := checkBatchLen(int(over)); !errors.Is(err, errBatchTooLarge) {
		t.Errorf("checkBatchLen(MaxInt32+1) = %v, want errBatchTooLarge", err)
	}
}
