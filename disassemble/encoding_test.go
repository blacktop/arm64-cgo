package disassemble_test

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

func TestEncodingString(t *testing.T) {
	tests := []struct {
		opcode uint32
		want   string
	}{
		{0x14000004, "B_only_branch_imm"},
		{0x94000004, "BL_only_branch_imm"},
		{0xb4000040, "CBZ_64_compbranch"},
		{0xd65f03c0, "RET_64R_branch_reg"},
		{0x91000420, "ADD_64_addsub_imm"},
		{0x00201420, "GENTER"},
		{0x00201000, "AMX_LDX"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			var results [1024]byte
			instr, err := disassemble.Decompose(0, tt.opcode, &results)
			if err != nil {
				t.Fatalf("Decompose(%#08x) error = %v", tt.opcode, err)
			}
			if got := instr.Encoding.String(); got != tt.want {
				t.Errorf("Encoding.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
