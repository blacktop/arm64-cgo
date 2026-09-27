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
			opcode:     0x00201000,
			wantOp:     disassemble.ARM64_AMX_LDX,
			wantDisasm: "ldx\tx0",
		},
		{
			name:       "AMX LDY X0",
			opcode:     0x00201020,
			wantOp:     disassemble.ARM64_AMX_LDY,
			wantDisasm: "ldy\tx0",
		},
		{
			name:       "AMX STX X0",
			opcode:     0x00201040,
			wantOp:     disassemble.ARM64_AMX_STX,
			wantDisasm: "stx\tx0",
		},
		{
			name:       "AMX LDZ X0",
			opcode:     0x00201080,
			wantOp:     disassemble.ARM64_AMX_LDZ,
			wantDisasm: "ldz\tx0",
		},

		// AMX FMA instructions
		{
			name:       "AMX FMA64 X0",
			opcode:     0x00201140,
			wantOp:     disassemble.ARM64_AMX_FMA64,
			wantDisasm: "fma64\tx0",
		},
		{
			name:       "AMX FMA32 X0",
			opcode:     0x00201180,
			wantOp:     disassemble.ARM64_AMX_FMA32,
			wantDisasm: "fma32\tx0",
		},

		{
			name:       "AMX GENLUT X5",
			opcode:     0x002012c5,
			wantOp:     disassemble.ARM64_AMX_GENLUT,
			wantDisasm: "genlut\tx5",
		},
		{
			name:       "AMX SET",
			opcode:     0x00201220,
			wantOp:     disassemble.ARM64_AMX_SET,
			wantDisasm: "set",
		},
		{
			name:       "AMX CLR",
			opcode:     0x00201221,
			wantOp:     disassemble.ARM64_AMX_CLR,
			wantDisasm: "clr",
		},
		{
			name:    "AMX op 17 with operand > 1",
			opcode:  0x00201222,
			wantErr: true,
		},
		{
			name:    "AMX op 23 is unallocated",
			opcode:  0x002012e0,
			wantErr: true,
		},

		// Guarded execution
		{
			name:       "GEXIT",
			opcode:     0x00201400,
			wantOp:     disassemble.ARM64_GEXIT,
			wantDisasm: "gexit",
		},
		{
			name:       "GENTER #2",
			opcode:     0x00201422,
			wantOp:     disassemble.ARM64_GENTER,
			wantDisasm: "genter	#0x2",
		},

		// Memory compression
		{
			name:       "WKDMC",
			opcode:     0x00200822,
			wantOp:     disassemble.ARM64_WKDMC,
			wantDisasm: "wkdmc\tx2, x1",
		},
		{
			name:       "WKDMC x2, x3",
			opcode:     0x00200862,
			wantOp:     disassemble.ARM64_WKDMC,
			wantDisasm: "wkdmc\tx2, x3",
		},
		{
			name:       "WKDMD",
			opcode:     0x00200c22,
			wantOp:     disassemble.ARM64_WKDMD,
			wantDisasm: "wkdmd\tx2, x1",
		},

		// Words outside Apple's 0x0020xxxx space are not Apple instructions
		{
			name:    "reserved 0x00102000",
			opcode:  0x00102000,
			wantErr: true,
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

				if instr.Disassembly != tt.wantDisasm {
					t.Errorf("Disassembly = %q, want %q", instr.Disassembly, tt.wantDisasm)
				}
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
