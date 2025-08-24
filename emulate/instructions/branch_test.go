package instructions

import (
	"strings"
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

func TestBranchExecutor_B(t *testing.T) {
	tests := []struct {
		name         string
		initialPC    uint64
		targetOffset int64
		expectedPC   uint64
		signedImm    bool
	}{
		{
			name:         "forward branch",
			initialPC:    0x1000,
			targetOffset: 0x100,
			expectedPC:   0x1100,
			signedImm:    false,
		},
		{
			name:         "backward branch",
			initialPC:    0x1000,
			targetOffset: -0x100,
			expectedPC:   0xF00,
			signedImm:    true,
		},
		{
			name:         "zero offset branch",
			initialPC:    0x1000,
			targetOffset: 0,
			expectedPC:   0x1000,
			signedImm:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("B", "Branch")
			state := state.NewState()
			state.SetPC(tt.initialPC)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_B,
				Operands: []disassemble.Operand{
					{
						Immediate: uint64(tt.targetOffset),
						SignedImm: tt.signedImm,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			if state.GetPC() != tt.expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", tt.expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_BL(t *testing.T) {
	executor := NewBranchExecutor("BL", "Branch with link")
	state := state.NewState()

	initialPC := uint64(0x1000)
	targetOffset := uint64(0x100)
	expectedPC := initialPC + targetOffset
	expectedLR := initialPC + 4

	state.SetPC(initialPC)

	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_BL,
		Operands: []disassemble.Operand{
			{
				Immediate: targetOffset,
				SignedImm: false,
			},
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if state.GetPC() != expectedPC {
		t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
	}

	if state.GetX(30) != expectedLR {
		t.Errorf("Expected LR=0x%x, got LR=0x%x", expectedLR, state.GetX(30))
	}
}

func TestBranchExecutor_BR(t *testing.T) {
	executor := NewBranchExecutor("BR", "Branch to register")
	state := state.NewState()

	initialPC := uint64(0x1000)
	targetAddr := uint64(0x2000)
	targetReg := 5

	state.SetPC(initialPC)
	state.SetX(targetReg, targetAddr)

	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_BR,
		Operands: []disassemble.Operand{
			{
				Registers: []disassemble.Register{disassemble.Register(34 + targetReg)}, // X5
			},
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if state.GetPC() != targetAddr {
		t.Errorf("Expected PC=0x%x, got PC=0x%x", targetAddr, state.GetPC())
	}
}

func TestBranchExecutor_BLR(t *testing.T) {
	executor := NewBranchExecutor("BLR", "Branch with link to register")
	state := state.NewState()

	initialPC := uint64(0x1000)
	targetAddr := uint64(0x2000)
	targetReg := 5
	expectedLR := initialPC + 4

	state.SetPC(initialPC)
	state.SetX(targetReg, targetAddr)

	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_BLR,
		Operands: []disassemble.Operand{
			{
				Registers: []disassemble.Register{disassemble.Register(34 + targetReg)}, // X5
			},
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if state.GetPC() != targetAddr {
		t.Errorf("Expected PC=0x%x, got PC=0x%x", targetAddr, state.GetPC())
	}

	if state.GetX(30) != expectedLR {
		t.Errorf("Expected LR=0x%x, got LR=0x%x", expectedLR, state.GetX(30))
	}
}

func TestBranchExecutor_RET(t *testing.T) {
	tests := []struct {
		name       string
		hasOperand bool
		retReg     int
		retAddr    uint64
	}{
		{
			name:       "default RET (X30)",
			hasOperand: false,
			retReg:     30,
			retAddr:    0x2000,
		},
		{
			name:       "RET with specific register",
			hasOperand: true,
			retReg:     5,
			retAddr:    0x3000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("RET", "Return from subroutine")
			state := state.NewState()

			initialPC := uint64(0x1000)
			state.SetPC(initialPC)
			state.SetX(tt.retReg, tt.retAddr)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_RET,
			}

			if tt.hasOperand {
				instr.Operands = []disassemble.Operand{
					{
						Registers: []disassemble.Register{disassemble.Register(34 + tt.retReg)},
					},
				}
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			if state.GetPC() != tt.retAddr {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", tt.retAddr, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_CBZ(t *testing.T) {
	tests := []struct {
		name         string
		regValue     uint64
		is32Bit      bool
		shouldBranch bool
		targetOffset uint64
	}{
		{
			name:         "64-bit zero value - should branch",
			regValue:     0,
			is32Bit:      false,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "64-bit non-zero value - should not branch",
			regValue:     0x123456789ABCDEF0,
			is32Bit:      false,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "32-bit zero value - should branch",
			regValue:     0x1234567800000000, // Upper bits set, but W register is zero
			is32Bit:      true,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "32-bit non-zero value - should not branch",
			regValue:     0x12345678ABCDEF01,
			is32Bit:      true,
			shouldBranch: false,
			targetOffset: 0x100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("CBZ", "Compare and branch if zero")
			state := state.NewState()

			initialPC := uint64(0x1000)
			testReg := 5
			state.SetPC(initialPC)
			state.SetX(testReg, tt.regValue)

			var regID disassemble.Register
			if tt.is32Bit {
				regID = disassemble.Register(1 + testReg) // W5
			} else {
				regID = disassemble.Register(34 + testReg) // X5
			}

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CBZ,
				Operands: []disassemble.Operand{
					{
						Registers: []disassemble.Register{regID},
					},
					{
						Immediate: tt.targetOffset,
						SignedImm: false,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			expectedPC := initialPC
			if tt.shouldBranch {
				expectedPC += tt.targetOffset
			} else {
				expectedPC += 4
			}

			if state.GetPC() != expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_CBNZ(t *testing.T) {
	tests := []struct {
		name         string
		regValue     uint64
		is32Bit      bool
		shouldBranch bool
		targetOffset uint64
	}{
		{
			name:         "64-bit zero value - should not branch",
			regValue:     0,
			is32Bit:      false,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "64-bit non-zero value - should branch",
			regValue:     0x123456789ABCDEF0,
			is32Bit:      false,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "32-bit zero value - should not branch",
			regValue:     0x1234567800000000, // Upper bits set, but W register is zero
			is32Bit:      true,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "32-bit non-zero value - should branch",
			regValue:     0x12345678ABCDEF01,
			is32Bit:      true,
			shouldBranch: true,
			targetOffset: 0x100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("CBNZ", "Compare and branch if not zero")
			state := state.NewState()

			initialPC := uint64(0x1000)
			testReg := 5
			state.SetPC(initialPC)
			state.SetX(testReg, tt.regValue)

			var regID disassemble.Register
			if tt.is32Bit {
				regID = disassemble.Register(1 + testReg) // W5
			} else {
				regID = disassemble.Register(34 + testReg) // X5
			}

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CBNZ,
				Operands: []disassemble.Operand{
					{
						Registers: []disassemble.Register{regID},
					},
					{
						Immediate: tt.targetOffset,
						SignedImm: false,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			expectedPC := initialPC
			if tt.shouldBranch {
				expectedPC += tt.targetOffset
			} else {
				expectedPC += 4
			}

			if state.GetPC() != expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_TBZ(t *testing.T) {
	tests := []struct {
		name         string
		regValue     uint64
		bitPos       uint64
		shouldBranch bool
		targetOffset uint64
	}{
		{
			name:         "bit 0 is zero - should branch",
			regValue:     0xFFFFFFFFFFFFFFFE, // All bits set except bit 0
			bitPos:       0,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "bit 0 is one - should not branch",
			regValue:     0xFFFFFFFFFFFFFFFF, // All bits set
			bitPos:       0,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "bit 63 is zero - should branch",
			regValue:     0x7FFFFFFFFFFFFFFF, // All bits set except bit 63
			bitPos:       63,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "bit 63 is one - should not branch",
			regValue:     0x8000000000000000, // Only bit 63 set
			bitPos:       63,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "bit 32 is zero - should branch",
			regValue:     0xFFFFFFFEFFFFFFFF, // All bits set except bit 32
			bitPos:       32,
			shouldBranch: true,
			targetOffset: 0x100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("TBZ", "Test bit and branch if zero")
			state := state.NewState()

			initialPC := uint64(0x1000)
			testReg := 5
			state.SetPC(initialPC)
			state.SetX(testReg, tt.regValue)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_TBZ,
				Operands: []disassemble.Operand{
					{
						Registers: []disassemble.Register{disassemble.Register(34 + testReg)}, // X5
					},
					{
						Immediate: tt.bitPos,
					},
					{
						Immediate: tt.targetOffset,
						SignedImm: false,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			expectedPC := initialPC
			if tt.shouldBranch {
				expectedPC += tt.targetOffset
			} else {
				expectedPC += 4
			}

			if state.GetPC() != expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_TBNZ(t *testing.T) {
	tests := []struct {
		name         string
		regValue     uint64
		bitPos       uint64
		shouldBranch bool
		targetOffset uint64
	}{
		{
			name:         "bit 0 is zero - should not branch",
			regValue:     0xFFFFFFFFFFFFFFFE, // All bits set except bit 0
			bitPos:       0,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "bit 0 is one - should branch",
			regValue:     0xFFFFFFFFFFFFFFFF, // All bits set
			bitPos:       0,
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "bit 63 is zero - should not branch",
			regValue:     0x7FFFFFFFFFFFFFFF, // All bits set except bit 63
			bitPos:       63,
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "bit 63 is one - should branch",
			regValue:     0x8000000000000000, // Only bit 63 set
			bitPos:       63,
			shouldBranch: true,
			targetOffset: 0x100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor("TBNZ", "Test bit and branch if not zero")
			state := state.NewState()

			initialPC := uint64(0x1000)
			testReg := 5
			state.SetPC(initialPC)
			state.SetX(testReg, tt.regValue)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_TBNZ,
				Operands: []disassemble.Operand{
					{
						Registers: []disassemble.Register{disassemble.Register(34 + testReg)}, // X5
					},
					{
						Immediate: tt.bitPos,
					},
					{
						Immediate: tt.targetOffset,
						SignedImm: false,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			expectedPC := initialPC
			if tt.shouldBranch {
				expectedPC += tt.targetOffset
			} else {
				expectedPC += 4
			}

			if state.GetPC() != expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_ConditionalBranch(t *testing.T) {
	tests := []struct {
		name         string
		condition    string
		flags        struct{ n, z, c, v bool }
		shouldBranch bool
		targetOffset uint64
	}{
		{
			name:         "B.EQ with Z=1 - should branch",
			condition:    "b.eq",
			flags:        struct{ n, z, c, v bool }{false, true, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.EQ with Z=0 - should not branch",
			condition:    "b.eq",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "B.NE with Z=0 - should branch",
			condition:    "b.ne",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.NE with Z=1 - should not branch",
			condition:    "b.ne",
			flags:        struct{ n, z, c, v bool }{false, true, false, false},
			shouldBranch: false,
			targetOffset: 0x100,
		},
		{
			name:         "B.CS with C=1 - should branch",
			condition:    "b.cs",
			flags:        struct{ n, z, c, v bool }{false, false, true, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.CC with C=0 - should branch",
			condition:    "b.cc",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.MI with N=1 - should branch",
			condition:    "b.mi",
			flags:        struct{ n, z, c, v bool }{true, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.PL with N=0 - should branch",
			condition:    "b.pl",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.VS with V=1 - should branch",
			condition:    "b.vs",
			flags:        struct{ n, z, c, v bool }{false, false, false, true},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.VC with V=0 - should branch",
			condition:    "b.vc",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.HI with C=1,Z=0 - should branch",
			condition:    "b.hi",
			flags:        struct{ n, z, c, v bool }{false, false, true, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.LS with C=0 - should branch",
			condition:    "b.ls",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.GE with N=V=0 - should branch",
			condition:    "b.ge",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.LT with N=1,V=0 - should branch",
			condition:    "b.lt",
			flags:        struct{ n, z, c, v bool }{true, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.GT with Z=0,N=V=0 - should branch",
			condition:    "b.gt",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.LE with Z=1 - should branch",
			condition:    "b.le",
			flags:        struct{ n, z, c, v bool }{false, true, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.AL - should always branch",
			condition:    "b.al",
			flags:        struct{ n, z, c, v bool }{false, false, false, false},
			shouldBranch: true,
			targetOffset: 0x100,
		},
		{
			name:         "B.NV - should never branch",
			condition:    "b.nv",
			flags:        struct{ n, z, c, v bool }{true, true, true, true},
			shouldBranch: false,
			targetOffset: 0x100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Extract the condition from the test condition string (e.g., "b.eq" -> "B.EQ")
			conditionParts := strings.Split(tt.condition, ".")
			if len(conditionParts) != 2 {
				t.Fatalf("Invalid condition format: %s", tt.condition)
			}
			mnemonic := strings.ToUpper(conditionParts[0]) + "." + strings.ToUpper(conditionParts[1])

			executor := NewBranchExecutor(mnemonic, "Conditional branch")
			state := state.NewState()

			initialPC := uint64(0x1000)
			state.SetPC(initialPC)

			// Set flags
			state.SetN(tt.flags.n)
			state.SetZ(tt.flags.z)
			state.SetC(tt.flags.c)
			state.SetV(tt.flags.v)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_B_EQ, // Use a generic conditional branch operation
				Operands: []disassemble.Operand{
					{
						Immediate: tt.targetOffset,
						SignedImm: false,
					},
				},
			}

			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			expectedPC := initialPC
			if tt.shouldBranch {
				expectedPC += tt.targetOffset
			} else {
				expectedPC += 4
			}

			if state.GetPC() != expectedPC {
				t.Errorf("Expected PC=0x%x, got PC=0x%x", expectedPC, state.GetPC())
			}
		})
	}
}

func TestBranchExecutor_InvalidInstructions(t *testing.T) {
	tests := []struct {
		name        string
		mnemonic    string
		operands    []disassemble.Operand
		expectError bool
	}{
		{
			name:        "B with no operands",
			mnemonic:    "B",
			operands:    []disassemble.Operand{},
			expectError: true,
		},
		{
			name:        "BL with no operands",
			mnemonic:    "BL",
			operands:    []disassemble.Operand{},
			expectError: true,
		},
		{
			name:        "BR with no operands",
			mnemonic:    "BR",
			operands:    []disassemble.Operand{},
			expectError: true,
		},
		{
			name:        "CBZ with only one operand",
			mnemonic:    "CBZ",
			operands:    []disassemble.Operand{{Registers: []disassemble.Register{disassemble.Register(34)}}},
			expectError: true,
		},
		{
			name:     "TBZ with invalid bit position",
			mnemonic: "TBZ",
			operands: []disassemble.Operand{
				{Registers: []disassemble.Register{disassemble.Register(34)}},
				{Immediate: 64}, // Invalid bit position
				{Immediate: 0x100},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewBranchExecutor(tt.mnemonic, "Test instruction")
			state := state.NewState()
			state.SetPC(0x1000)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_B, // Use a generic operation for error testing
				Operands:  tt.operands,
			}

			err := executor.Execute(state, instr)
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestBranchHelperFunctions(t *testing.T) {
	tests := []struct {
		name     string
		mnemonic string
		expected bool
		testFunc func(string) bool
	}{
		{"B is branch", "B", true, IsBranchInstruction},
		{"BL is branch", "BL", true, IsBranchInstruction},
		{"CBZ is branch", "CBZ", true, IsBranchInstruction},
		{"B.EQ is branch", "B.EQ", true, IsBranchInstruction},
		{"ADD is not branch", "ADD", false, IsBranchInstruction},

		{"B is unconditional", "B", true, IsUnconditionalBranch},
		{"BL is unconditional", "BL", true, IsUnconditionalBranch},
		{"B.AL is unconditional", "B.AL", true, IsUnconditionalBranch},
		{"B.EQ is not unconditional", "B.EQ", false, IsUnconditionalBranch},
		{"CBZ is not unconditional", "CBZ", false, IsUnconditionalBranch},

		{"B.EQ is conditional", "B.EQ", true, IsConditionalBranch},
		{"B.NE is conditional", "B.NE", true, IsConditionalBranch},
		{"B is not conditional", "B", false, IsConditionalBranch},
		{"CBZ is not conditional", "CBZ", false, IsConditionalBranch},

		{"BL is link branch", "BL", true, IsLinkBranch},
		{"BLR is link branch", "BLR", true, IsLinkBranch},
		{"B is not link branch", "B", false, IsLinkBranch},
		{"BR is not link branch", "BR", false, IsLinkBranch},

		{"BR is register branch", "BR", true, IsRegisterBranch},
		{"BLR is register branch", "BLR", true, IsRegisterBranch},
		{"RET is register branch", "RET", true, IsRegisterBranch},
		{"B is not register branch", "B", false, IsRegisterBranch},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.testFunc(tt.mnemonic)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v for %s", tt.expected, result, tt.mnemonic)
			}
		})
	}
}
