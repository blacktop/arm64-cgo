package disassemble

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand/v2"
	"os"
	"runtime"
	"sync"
	"testing"
)

// corpusTemplate fixes an encoding's opcode bits (base) and lets the
// generator fill its operand fields (mask).
type corpusTemplate struct {
	base, mask uint32
	weight     int      // share of the corpus, in basis points
	enc        encoding // what every filled-in word must decode as
}

// corpusTemplates are the 40 most frequent encodings in the macOS 27 (26A428)
// Mac18,5 kernelcache __TEXT_EXEC, which together cover 88% of its 15.2M
// instructions, weighted by their frequency there.
var corpusTemplates = []corpusTemplate{
	{0xaa0003e0, 0x001f001f, 995, ENC_MOV_ORR_64_LOG_SHIFT},     // mov xd, xm
	{0xf9400000, 0x003fffff, 873, ENC_LDR_64_LDST_POS},          // ldr xt, [xn, #imm]
	{0x91000000, 0x007fffff, 842, ENC_ADD_64_ADDSUB_IMM},        // add xd, xn, #imm{, lsl #12}
	{0x90000000, 0x60ffffff, 585, ENC_ADRP_ONLY_PCRELADDR},      // adrp xd, label
	{0x94000000, 0x03ffffff, 543, ENC_BL_ONLY_BRANCH_IMM},       // bl label
	{0x52800000, 0x003fffff, 519, ENC_MOV_MOVZ_32_MOVEWIDE},     // mov wd, #imm (movz)
	{0xf2800000, 0x007fffff, 447, ENC_MOVK_64_MOVEWIDE},         // movk xd, #imm{, lsl #n}
	{0xa9000000, 0x003fffff, 430, ENC_STP_64_LDSTPAIR_OFF},      // stp xt, xt2, [xn, #imm]
	{0xd2800000, 0x007fffff, 262, ENC_MOV_MOVZ_64_MOVEWIDE},     // mov xd, #imm (movz)
	{0x14000000, 0x03ffffff, 241, ENC_B_ONLY_BRANCH_IMM},        // b label
	{0x54000000, 0x00ffffef, 234, ENC_B_ONLY_CONDBRANCH},        // b.cond label
	{0xf9000000, 0x003fffff, 233, ENC_STR_64_LDST_POS},          // str xt, [xn, #imm]
	{0xa9400000, 0x003fffff, 229, ENC_LDP_64_LDSTPAIR_OFF},      // ldp xt, xt2, [xn, #imm]
	{0xb9400000, 0x003fffff, 215, ENC_LDR_32_LDST_POS},          // ldr wt, [xn, #imm]
	{0x8b000000, 0x001fffff, 182, ENC_ADD_64_ADDSUB_SHIFT},      // add xd, xn, xm{, lsl #n}
	{0xd73f0800, 0x000003ff, 167, ENC_BLRAA_64P_BRANCH_REG},     // blraa xn, xm
	{0xdac11800, 0x000003ff, 164, ENC_AUTDA_64P_DP_1SRC},        // autda xd, xn
	{0xd503241f, 0x000000c0, 153, ENC_BTI_HB_HINTS},             // bti {c|j|jc}
	{0xb4000000, 0x00ffffff, 138, ENC_CBZ_64_COMPBRANCH},        // cbz xt, label
	{0x39400000, 0x003fffff, 98, ENC_LDRB_32_LDST_POS},          // ldrb wt, [xn, #imm]
	{0x7100001f, 0x007fffe0, 94, ENC_CMP_SUBS_32S_ADDSUB_IMM},   // cmp wn, #imm
	{0xb9000000, 0x003fffff, 91, ENC_STR_32_LDST_POS},           // str wt, [xn, #imm]
	{0xd1000000, 0x007fffff, 86, ENC_SUB_64_ADDSUB_IMM},         // sub xd, xn, #imm
	{0xdac1a7fe, 0x00000000, 84, ENC_PACIBSPPC_64LR_DP_1SRC},    // pacibsppc
	{0xeb00001f, 0x001fffe0, 80, ENC_CMP_SUBS_64_ADDSUB_SHIFT},  // cmp xn, xm{, lsl #n}
	{0x34000000, 0x00ffffff, 79, ENC_CBZ_32_COMPBRANCH},         // cbz wt, label
	{0x8b200000, 0x001fefff, 69, ENC_ADD_64_ADDSUB_EXT},         // add xd, xn, wm, <extend>
	{0x36000000, 0x80ffffff, 68, ENC_TBZ_ONLY_TESTBRANCH},       // tbz rt, #bit, label
	{0x9a800000, 0x001ff3ff, 67, ENC_CSEL_64_CONDSEL},           // csel xd, xn, xm, cond
	{0x39000000, 0x003fffff, 66, ENC_STRB_32_LDST_POS},          // strb wt, [xn, #imm]
	{0x5520001f, 0x001fffe0, 66, ENC_RETABSPPC_ONLY_MISCBRANCH}, // retabsppc label
	{0x72800000, 0x003fffff, 64, ENC_MOVK_32_MOVEWIDE},          // movk wd, #imm{, lsl #16}
	{0xa8c00000, 0x003fffff, 56, ENC_LDP_64_LDSTPAIR_POST},      // ldp xt, xt2, [xn], #imm
	{0x37000000, 0x80ffffff, 52, ENC_TBNZ_ONLY_TESTBRANCH},      // tbnz rt, #bit, label
	{0xa9800000, 0x003fffff, 52, ENC_STP_64_LDSTPAIR_PRE},       // stp xt, xt2, [xn, #imm]!
	{0xeb20001f, 0x001fefe0, 51, ENC_CMP_SUBS_64S_ADDSUB_EXT},   // cmp xn, wm, <extend>
	{0xf8400c00, 0x001ff3ff, 45, ENC_LDR_64_LDST_IMMPRE},        // ldr xt, [xn, #imm]!
	{0xd65f03c0, 0x00000000, 42, ENC_RET_64R_BRANCH_REG},        // ret
	{0x35000000, 0x00ffffff, 41, ENC_CBNZ_32_COMPBRANCH},        // cbnz wt, label
	{0xad000000, 0x003fffff, 39, ENC_STP_Q_LDSTPAIR_OFF},        // stp qt, qt2, [xn, #imm]
}

const corpusBase = 0xfffffe0008000000

// syntheticCorpus returns 64K instruction words drawn from corpusTemplates
// with a fixed seed, so every run decodes the same instructions.
var syntheticCorpus = sync.OnceValue(func() []uint32 {
	total := 0
	for _, t := range corpusTemplates {
		total += t.weight
	}
	r := rand.New(rand.NewPCG(1, 2))
	words := make([]uint32, 1<<16)
	for i := range words {
		pick := r.IntN(total)
		for _, t := range corpusTemplates {
			if pick < t.weight {
				words[i] = t.base | r.Uint32()&t.mask
				break
			}
			pick -= t.weight
		}
	}
	return words
})

type benchCorpus struct {
	name       string
	provenance string // logged by each measured sub-benchmark
	addrs      []uint64
	words      []uint32
}

// benchCorpora returns the synthetic corpus and, when ARM64_BENCH_CORPUS
// names a file of little-endian instruction words, that file as well.
func benchCorpora(b *testing.B) []benchCorpus {
	b.Helper()
	corpora := []benchCorpus{newBenchCorpus("synthetic", "", syntheticCorpus())}

	path := os.Getenv("ARM64_BENCH_CORPUS")
	if path == "" {
		return corpora
	}
	data, err := os.ReadFile(path)
	if err != nil {
		b.Fatalf("read ARM64_BENCH_CORPUS: %v", err)
	}
	if len(data) == 0 || len(data)%4 != 0 {
		b.Fatalf("ARM64_BENCH_CORPUS %s is %d bytes; want a non-empty multiple of 4",
			path, len(data))
	}
	words := make([]uint32, len(data)/4)
	for i := range words {
		words[i] = binary.LittleEndian.Uint32(data[i*4:])
	}
	provenance := fmt.Sprintf("external corpus %s: %d words, sha256 %x, %s",
		path, len(words), sha256.Sum256(data), runtime.Version())
	return append(corpora, newBenchCorpus("external", provenance, words))
}

func newBenchCorpus(name, provenance string, words []uint32) benchCorpus {
	addrs := make([]uint64, len(words))
	for i := range addrs {
		addrs[i] = corpusBase + uint64(i)*4
	}
	return benchCorpus{name: name, provenance: provenance, addrs: addrs, words: words}
}

// runCorpora runs bench as a sub-benchmark per corpus.
func runCorpora(b *testing.B, bench func(b *testing.B, c benchCorpus)) {
	for _, c := range benchCorpora(b) {
		b.Run(c.name, func(b *testing.B) {
			if c.provenance != "" {
				b.Log(c.provenance)
			}
			b.SetBytes(int64(len(c.words) * 4))
			b.ReportAllocs()
			bench(b, c)
		})
	}
}

func BenchmarkCorpusDisassemble(b *testing.B) {
	runCorpora(b, func(b *testing.B, c benchCorpus) {
		var results [1024]byte
		for b.Loop() {
			for i, w := range c.words {
				_, _ = Disassemble(c.addrs[i], w, &results)
			}
		}
	})
}

func BenchmarkCorpusDecompose(b *testing.B) {
	runCorpora(b, func(b *testing.B, c benchCorpus) {
		var results [1024]byte
		for b.Loop() {
			for i, w := range c.words {
				_, _ = Decompose(c.addrs[i], w, &results)
			}
		}
	})
}

func BenchmarkCorpusDecomposeInto(b *testing.B) {
	runCorpora(b, func(b *testing.B, c benchCorpus) {
		var inst Inst
		for b.Loop() {
			for i, w := range c.words {
				_ = DecomposeInto(c.addrs[i], w, &inst)
			}
		}
	})
}

func BenchmarkCorpusDecoderBatch(b *testing.B) {
	runCorpora(b, func(b *testing.B, c benchCorpus) {
		out := make([]Inst, len(c.words))
		status := make([]DecodeStatus, len(c.words))
		decoder := NewDecoder(len(c.words))
		for b.Loop() {
			_, _ = decoder.DecomposeBatchStatus(c.addrs, c.words, out, status)
		}
	})
}

// Every template, with any operand fields, must decode as the encoding it
// stands for; otherwise the corpus no longer matches the kernel mix.
func TestCorpusTemplateEncodings(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	var inst Inst
	for _, tmpl := range corpusTemplates {
		for n := range 1024 {
			word := tmpl.base
			if n > 0 {
				word |= r.Uint32() & tmpl.mask
			}
			if err := DecomposeInto(0, word, &inst); err != nil {
				t.Fatalf("template %s: %#08x does not decode: %v", tmpl.enc, word, err)
			}
			if inst.Encoding != tmpl.enc {
				t.Fatalf("template %s: %#08x decodes as %s", tmpl.enc, word, inst.Encoding)
			}
		}
	}
}

// Every synthetic corpus word must decode, as real code nearly always does.
func TestSyntheticCorpusDecodes(t *testing.T) {
	var inst Inst
	for i, w := range syntheticCorpus() {
		if err := DecomposeInto(corpusBase+uint64(i)*4, w, &inst); err != nil {
			t.Fatalf("corpus word %d (%#08x) does not decode: %v", i, w, err)
		}
	}
}
