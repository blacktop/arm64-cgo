package disassemble_test

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

func TestAppleInstructions(t *testing.T) {
	tests := []struct {
		name       string
		opcode     uint32
		wantOp     disassemble.Operation
		wantDisasm string
		wantErr    bool
	}{
		// AMX Load/Store instructions
		{
			name:       "AMX LDX X0",
			opcode:     0x00102000,
			wantOp:     disassemble.ARM64_AMX_LDX,
			wantDisasm: "ldx x0",
		},
		{
			name:       "AMX LDY X0",
			opcode:     0x00102020,
			wantOp:     disassemble.ARM64_AMX_LDY,
			wantDisasm: "ldy x0",
		},
		{
			name:       "AMX STX X0",
			opcode:     0x00102040,
			wantOp:     disassemble.ARM64_AMX_STX,
			wantDisasm: "stx x0",
		},
		{
			name:       "AMX LDZ X0",
			opcode:     0x00102080,
			wantOp:     disassemble.ARM64_AMX_LDZ,
			wantDisasm: "ldz x0",
		},

		// AMX FMA instructions
		{
			name:       "AMX FMA64 X0",
			opcode:     0x00112040,
			wantOp:     disassemble.ARM64_AMX_FMA64,
			wantDisasm: "fma64 x0",
		},
		{
			name:       "AMX FMA32 X0",
			opcode:     0x00112080,
			wantOp:     disassemble.ARM64_AMX_FMA32,
			wantDisasm: "fma32 x0",
		},

		// Guarded execution
		{
			name:       "GEXIT",
			opcode:     0x00142000,
			wantOp:     disassemble.ARM64_GEXIT,
			wantDisasm: "gexit",
		},
		{
			name:       "GENTER #2",
			opcode:     0x00142002,
			wantOp:     disassemble.ARM64_GENTER,
			wantDisasm: "genter #2",
		},

		// Memory compression
		{
			name:       "WKDMC",
			opcode:     0x00200822,
			wantOp:     disassemble.ARM64_WKDMC,
			wantDisasm: "wkdmc",
		},
		{
			name:       "WKDMD",
			opcode:     0x00200862,
			wantOp:     disassemble.ARM64_WKDMD,
			wantDisasm: "wkdmd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var results [1024]byte
			instr, err := disassemble.Decompose(0x1000, tt.opcode, &results)

			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				if instr.Operation != tt.wantOp {
					t.Errorf("Operation = %v (%s), want %v",
						instr.Operation, instr.Operation.String(), tt.wantOp)
				}

				// Check if the operation string is correct
				opStr := instr.Operation.String()
				if opStr == "" || opStr == "error" {
					t.Errorf("Operation string is empty or error for %v", tt.wantOp)
				}

				t.Logf("Decoded: %s (op: %s)", instr.Disassembly, instr.Operation.String())
			}
		})
	}
}

func TestAppleSystemRegisters(t *testing.T) {
	// Test that Apple AMX system registers are accessible
	amxRegs := []struct {
		name string
		reg  disassemble.SystemReg
	}{
		// These would need to be added to the SystemReg constants
		// For now, just test the infrastructure works
		{"SYSREG_NONE", disassemble.SYSREG_NONE},
	}

	for _, tt := range amxRegs {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.reg.String()
			if str == "" {
				t.Errorf("SystemReg.String() returned empty for %s", tt.name)
			}
			t.Logf("%s: %s", tt.name, str)
		})
	}
}
