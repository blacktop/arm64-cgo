package disassemble_test

import (
	"fmt"
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

func TestSimpleAppleCheck(t *testing.T) {
	// First, verify the Apple operations exist
	ops := []disassemble.Operation{
		disassemble.ARM64_AMX_LDX,
		disassemble.ARM64_GENTER,
		disassemble.ARM64_GEXIT,
	}

	for _, op := range ops {
		name := op.String()
		if name == "" || name == "error" || name == "UNKNOWN" {
			t.Errorf("Operation %d has invalid string: %s", op, name)
		} else {
			t.Logf("Operation %d = %s", op, name)
		}
	}

	// Now test a simple known-good instruction first
	var results [1024]byte
	instr, err := disassemble.Decompose(0x1000, 0x91000000, &results) // ADD X0, X0, #0
	if err != nil {
		t.Logf("Known instruction failed: %v", err)
	} else {
		t.Logf("Known instruction succeeded: %s (op: %s)",
			instr.Disassembly, instr.Operation.String())
	}

	// Try the simplest Apple instruction
	testOpcode := uint32(0x00102000) // AMX LDX X0
	instr, err = disassemble.Decompose(0x1000, testOpcode, &results)
	if err != nil {
		t.Logf("Apple instruction 0x%08x failed: %v", testOpcode, err)
		// Let's try to understand what the decoder sees
		fmt.Printf("Opcode: 0x%08x\n", testOpcode)
		fmt.Printf("Bits [31:24]: 0x%02x\n", (testOpcode>>24)&0xFF)
		fmt.Printf("Bits [23:16]: 0x%02x\n", (testOpcode>>16)&0xFF)
		fmt.Printf("Bits [15:8]: 0x%02x\n", (testOpcode>>8)&0xFF)
		fmt.Printf("Bits [7:0]: 0x%02x\n", testOpcode&0xFF)
	} else {
		t.Logf("Apple instruction succeeded: %s (op: %s)",
			instr.Disassembly, instr.Operation.String())
	}
}
