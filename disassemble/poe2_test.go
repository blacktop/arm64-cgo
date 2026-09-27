package disassemble_test

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

// Vectors from binutils gas/testsuite/gas/aarch64/{poe2,tev}.d and the
// macOS 27 (26A428) Mac18,5 kernelcache GXF call stubs.
func TestPOE2Instructions(t *testing.T) {
	tests := []struct {
		opcode  uint32
		wantOp  disassemble.Operation
		wantEnc string
		want    string
	}{
		{0xd4e00000, disassemble.ARM64_TENTER, "TENTER_EX_exception", "tenter\t#0"},
		{0xd4e00020, disassemble.ARM64_TENTER, "TENTER_EX_exception", "tenter\t#0x1"},
		{0xd4e00fe0, disassemble.ARM64_TENTER, "TENTER_EX_exception", "tenter\t#0x7f"},
		{0xd4e20000, disassemble.ARM64_TENTER, "TENTER_EX_exception", "tenter\t#0, nb"},
		{0xd4e20fe0, disassemble.ARM64_TENTER, "TENTER_EX_exception", "tenter\t#0x7f, nb"},
		{0xd6ff03e0, disassemble.ARM64_TEXIT, "TEXIT_64E_branch_reg", "texit"},
		{0xd6ff07e0, disassemble.ARM64_TEXIT, "TEXIT_64E_branch_reg", "texit\tnb"},

		{0xd5800020, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\tx0, x1"},
		{0xd5800040, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\tx0, x2"},
		{0xd580003f, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\txzr, x1"},
		{0xd58003e1, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\tx1, xzr"},
		{0xd5820020, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\tx0, x1, nb"},
		{0xd58203e1, disassemble.ARM64_TCHANGEF, "TCHANGEF_reg", "tchangef\tx1, xzr, nb"},
		{0xd5900000, disassemble.ARM64_TCHANGEF, "TCHANGEF_imm", "tchangef\tx0, #0"},
		{0xd590003f, disassemble.ARM64_TCHANGEF, "TCHANGEF_imm", "tchangef\txzr, #0x1"},
		{0xd5900fea, disassemble.ARM64_TCHANGEF, "TCHANGEF_imm", "tchangef\tx10, #0x7f"},
		{0xd59207fe, disassemble.ARM64_TCHANGEF, "TCHANGEF_imm", "tchangef\tlr, #0x3f, nb"},
		{0xd5840020, disassemble.ARM64_TCHANGEB, "TCHANGEB_reg", "tchangeb\tx0, x1"},
		{0xd58403e1, disassemble.ARM64_TCHANGEB, "TCHANGEB_reg", "tchangeb\tx1, xzr"},
		{0xd586003f, disassemble.ARM64_TCHANGEB, "TCHANGEB_reg", "tchangeb\txzr, x1, nb"},
		{0xd5940000, disassemble.ARM64_TCHANGEB, "TCHANGEB_imm", "tchangeb\tx0, #0"},
		{0xd5940fea, disassemble.ARM64_TCHANGEB, "TCHANGEB_imm", "tchangeb\tx10, #0x7f"},
		{0xd5960fea, disassemble.ARM64_TCHANGEB, "TCHANGEB_imm", "tchangeb\tx10, #0x7f, nb"},

		// neighbouring encodings must keep decoding as before
		{0xd69f03e0, disassemble.ARM64_ERET, "ERET_64E_branch_reg", "eret"},
		{0xd6bf03e0, disassemble.ARM64_DRPS, "DRPS_64E_branch_reg", "drps"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			var results [1024]byte
			instr, err := disassemble.Decompose(0x1000, tt.opcode, &results)
			if err != nil {
				t.Fatalf("Decompose(%#08x) error = %v", tt.opcode, err)
			}
			if instr.Operation != tt.wantOp {
				t.Errorf("Operation = %s, want %s", instr.Operation, tt.wantOp)
			}
			if got := instr.Encoding.String(); got != tt.wantEnc {
				t.Errorf("Encoding = %q, want %q", got, tt.wantEnc)
			}
			if instr.Disassembly != tt.want {
				t.Errorf("Disassembly = %q, want %q", instr.Disassembly, tt.want)
			}
		})
	}
}

func TestPOE2Operands(t *testing.T) {
	var results [1024]byte
	instr, err := disassemble.Decompose(0, 0xd4e20fe0, &results) // tenter #127, nb
	if err != nil {
		t.Fatalf("Decompose error = %v", err)
	}
	if len(instr.Operands) != 2 {
		t.Fatalf("len(Operands) = %d, want 2", len(instr.Operands))
	}
	if got := instr.Operands[0].GetImmediate(); got != 127 {
		t.Errorf("imm = %d, want 127", got)
	}
	if got := instr.Operands[1].Name; got != "nb" {
		t.Errorf("Operands[1].Name = %q, want \"nb\"", got)
	}

	instr, err = disassemble.Decompose(0, 0xd6ff03e0, &results) // texit
	if err != nil {
		t.Fatalf("Decompose error = %v", err)
	}
	if len(instr.Operands) != 0 {
		t.Errorf("texit len(Operands) = %d, want 0", len(instr.Operands))
	}
}

// Words that differ from these instructions only in fixed bits must not decode as them.
func TestPOE2FixedBits(t *testing.T) {
	for _, opcode := range []uint32{
		0xd4e00001, // TENTER with LL != 00
		0xd4e00004, // TENTER with op2 != 000
		0xd4e01000, // TENTER with bit 12 set
		0xd4e10000, // TENTER with bit 16 set
		0xd4e40000, // TENTER with bit 18 set
		0xd6ff03c0, // TEXIT with Rn != 11111
		0xd6ff03e1, // TEXIT with op4 != 00000
		0xd6ff0be0, // TEXIT with bit 11 set
		0xd5800420, // TCHANGEF (register) with bit 10 set
		0xd5808020, // TCHANGEF (register) with bit 15 set
		0xd5810020, // TCHANGEF (register) with bit 16 set
		0xd5880020, // TCHANGEF (register) with bit 19 set
		0xd5901000, // TCHANGEF (immediate) with bit 12 set
		0xd5910000, // TCHANGEF (immediate) with bit 16 set
		0xd5980000, // TCHANGEF (immediate) with bit 19 set
	} {
		var results [1024]byte
		instr, err := disassemble.Decompose(0, opcode, &results)
		if err != nil {
			continue
		}
		switch instr.Operation {
		case disassemble.ARM64_TENTER, disassemble.ARM64_TEXIT,
			disassemble.ARM64_TCHANGEF, disassemble.ARM64_TCHANGEB:
			t.Errorf("%#08x decoded as %q, want not a POE2/TEV instruction", opcode, instr.Disassembly)
		}
	}
}
