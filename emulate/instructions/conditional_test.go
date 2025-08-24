package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// Helper function to set up executor with test condition
func setupTestExecutor(mnemonic, description string, condition core.ConditionCode) *ConditionalExecutor {
	executor := NewConditionalExecutor(mnemonic, description)
	executor.SetTestConditionOverride(condition)
	return executor
}

func TestConditionalExecutor_CSEL(t *testing.T) {
	tests := []struct {
		name        string
		condition   core.ConditionCode
		flags       map[string]bool
		src1Val     uint64
		src2Val     uint64
		expected    uint64
		shouldMatch bool
	}{
		{
			name:        "CSEL EQ condition true",
			condition:   core.EQ,
			flags:       map[string]bool{"Z": true},
			src1Val:     0x1234,
			src2Val:     0x5678,
			expected:    0x1234,
			shouldMatch: true,
		},
		{
			name:        "CSEL EQ condition false",
			condition:   core.EQ,
			flags:       map[string]bool{"Z": false},
			src1Val:     0x1234,
			src2Val:     0x5678,
			expected:    0x5678,
			shouldMatch: false,
		},
		{
			name:        "CSEL NE condition true",
			condition:   core.NE,
			flags:       map[string]bool{"Z": false},
			src1Val:     0xAAAA,
			src2Val:     0xBBBB,
			expected:    0xAAAA,
			shouldMatch: true,
		},
		{
			name:        "CSEL GE condition true (N==V)",
			condition:   core.GE,
			flags:       map[string]bool{"N": true, "V": true},
			src1Val:     0x1111,
			src2Val:     0x2222,
			expected:    0x1111,
			shouldMatch: true,
		},
		{
			name:        "CSEL GE condition false (N!=V)",
			condition:   core.GE,
			flags:       map[string]bool{"N": true, "V": false},
			src1Val:     0x1111,
			src2Val:     0x2222,
			expected:    0x2222,
			shouldMatch: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create state and set flags
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "N":
					s.SetN(value)
				case "Z":
					s.SetZ(value)
				case "C":
					s.SetC(value)
				case "V":
					s.SetV(value)
				}
			}

			// Set source register values
			s.SetX(1, tt.src1Val) // X1
			s.SetX(2, tt.src2Val) // X2

			// Create mock instruction
			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSEL,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
					{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
					{Registers: []disassemble.Register{disassemble.Register(36)}}, // X2
				},
			}

			executor := setupTestExecutor("CSEL", "Conditional select", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CSINC(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		src1Val   uint64
		src2Val   uint64
		expected  uint64
	}{
		{
			name:      "CSINC condition true",
			condition: core.EQ,
			flags:     map[string]bool{"Z": true},
			src1Val:   0x1000,
			src2Val:   0x2000,
			expected:  0x1000, // src1 selected
		},
		{
			name:      "CSINC condition false",
			condition: core.EQ,
			flags:     map[string]bool{"Z": false},
			src1Val:   0x1000,
			src2Val:   0x2000,
			expected:  0x2001, // src2 + 1 selected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "Z":
					s.SetZ(value)
				}
			}

			s.SetX(1, tt.src1Val)
			s.SetX(2, tt.src2Val)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSINC,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
					{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
					{Registers: []disassemble.Register{disassemble.Register(36)}}, // X2
				},
			}

			executor := setupTestExecutor("CSINC", "Conditional select increment", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CSINV(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		src1Val   uint64
		src2Val   uint64
		expected  uint64
	}{
		{
			name:      "CSINV condition true",
			condition: core.NE,
			flags:     map[string]bool{"Z": false},
			src1Val:   0x1234,
			src2Val:   0x5678,
			expected:  0x1234, // src1 selected
		},
		{
			name:      "CSINV condition false",
			condition: core.NE,
			flags:     map[string]bool{"Z": true},
			src1Val:   0x1234,
			src2Val:   0x0000000000005678,
			expected:  0xFFFFFFFFFFFFA987, // ~src2 selected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "Z":
					s.SetZ(value)
				}
			}

			s.SetX(1, tt.src1Val)
			s.SetX(2, tt.src2Val)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSINV,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
					{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
					{Registers: []disassemble.Register{disassemble.Register(36)}}, // X2
				},
			}

			executor := setupTestExecutor("CSINV", "Conditional select invert", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CSNEG(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		src1Val   uint64
		src2Val   uint64
		expected  uint64
	}{
		{
			name:      "CSNEG condition true",
			condition: core.CS,
			flags:     map[string]bool{"C": true},
			src1Val:   0x1000,
			src2Val:   0x2000,
			expected:  0x1000, // src1 selected
		},
		{
			name:      "CSNEG condition false - positive number",
			condition: core.CS,
			flags:     map[string]bool{"C": false},
			src1Val:   0x1000,
			src2Val:   0x2000,
			expected:  0xFFFFFFFFFFFFE000, // -src2 selected (two's complement)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "C":
					s.SetC(value)
				}
			}

			s.SetX(1, tt.src1Val)
			s.SetX(2, tt.src2Val)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSNEG,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
					{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
					{Registers: []disassemble.Register{disassemble.Register(36)}}, // X2
				},
			}

			executor := setupTestExecutor("CSNEG", "Conditional select negate", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CSET(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		expected  uint64
	}{
		{
			name:      "CSET condition true",
			condition: core.EQ,
			flags:     map[string]bool{"Z": true},
			expected:  1,
		},
		{
			name:      "CSET condition false",
			condition: core.EQ,
			flags:     map[string]bool{"Z": false},
			expected:  0,
		},
		{
			name:      "CSET GT condition true",
			condition: core.GT,
			flags:     map[string]bool{"Z": false, "N": false, "V": false},
			expected:  1,
		},
		{
			name:      "CSET GT condition false (Z=1)",
			condition: core.GT,
			flags:     map[string]bool{"Z": true, "N": false, "V": false},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "N":
					s.SetN(value)
				case "Z":
					s.SetZ(value)
				case "C":
					s.SetC(value)
				case "V":
					s.SetV(value)
				}
			}

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSET,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
				},
			}

			executor := setupTestExecutor("CSET", "Conditional set", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CSETM(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		expected  uint64
	}{
		{
			name:      "CSETM condition true",
			condition: core.NE,
			flags:     map[string]bool{"Z": false},
			expected:  0xFFFFFFFFFFFFFFFF,
		},
		{
			name:      "CSETM condition false",
			condition: core.NE,
			flags:     map[string]bool{"Z": true},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "Z":
					s.SetZ(value)
				}
			}

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CSETM,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
				},
			}

			executor := setupTestExecutor("CSETM", "Conditional set mask", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected X0 = 0x%x, got 0x%x", tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_CCMP(t *testing.T) {
	tests := []struct {
		name         string
		condition    core.ConditionCode
		flags        map[string]bool
		src1Val      uint64
		src2Val      uint64
		flagsImm     uint64
		expectedN    bool
		expectedZ    bool
		expectedC    bool
		expectedV    bool
		conditionMet bool
	}{
		{
			name:         "CCMP condition true - equal values",
			condition:    core.EQ,
			flags:        map[string]bool{"Z": true},
			src1Val:      0x1000,
			src2Val:      0x1000,
			flagsImm:     0x5, // NZCV = 0101
			expectedN:    false,
			expectedZ:    true,
			expectedC:    true,
			expectedV:    false,
			conditionMet: true,
		},
		{
			name:         "CCMP condition true - src1 > src2",
			condition:    core.EQ,
			flags:        map[string]bool{"Z": true},
			src1Val:      0x2000,
			src2Val:      0x1000,
			flagsImm:     0x5, // NZCV = 0101
			expectedN:    false,
			expectedZ:    false,
			expectedC:    true,
			expectedV:    false,
			conditionMet: true,
		},
		{
			name:         "CCMP condition false - use immediate flags",
			condition:    core.EQ,
			flags:        map[string]bool{"Z": false},
			src1Val:      0x1000,
			src2Val:      0x2000,
			flagsImm:     0xA, // NZCV = 1010
			expectedN:    true,
			expectedZ:    false,
			expectedC:    true,
			expectedV:    false,
			conditionMet: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "N":
					s.SetN(value)
				case "Z":
					s.SetZ(value)
				case "C":
					s.SetC(value)
				case "V":
					s.SetV(value)
				}
			}

			s.SetX(1, tt.src1Val)
			s.SetX(2, tt.src2Val)

			instr := &disassemble.Instruction{
				Operation: disassemble.ARM64_CCMP,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
					{Registers: []disassemble.Register{disassemble.Register(36)}}, // X2
					{Immediate: tt.flagsImm},
				},
			}

			executor := setupTestExecutor("CCMP", "Conditional compare", tt.condition)

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			if s.GetN() != tt.expectedN {
				t.Errorf("Expected N flag = %v, got %v", tt.expectedN, s.GetN())
			}
			if s.GetZ() != tt.expectedZ {
				t.Errorf("Expected Z flag = %v, got %v", tt.expectedZ, s.GetZ())
			}
			if s.GetC() != tt.expectedC {
				t.Errorf("Expected C flag = %v, got %v", tt.expectedC, s.GetC())
			}
			if s.GetV() != tt.expectedV {
				t.Errorf("Expected V flag = %v, got %v", tt.expectedV, s.GetV())
			}
		})
	}
}

func TestConditionalExecutor_32BitOperations(t *testing.T) {
	s := state.NewState()
	s.SetZ(true) // Condition will be true

	// Set source values
	s.SetX(1, 0x123456789ABCDEF0)
	s.SetX(2, 0xFEDCBA9876543210)

	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_CSEL,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(1)}}, // W0
			{Registers: []disassemble.Register{disassemble.Register(2)}}, // W1
			{Registers: []disassemble.Register{disassemble.Register(3)}}, // W2
		},
	}

	executor := setupTestExecutor("CSEL", "Conditional select", core.EQ)

	err := executor.Execute(s, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Should select W1 (lower 32 bits of X1) and zero-extend to X0
	expected := uint64(0x9ABCDEF0)
	result := s.GetX(0)
	if result != expected {
		t.Errorf("Expected X0 = 0x%x, got 0x%x", expected, result)
	}
}

func TestConditionalExecutor_ConditionEvaluation(t *testing.T) {
	tests := []struct {
		name      string
		condition core.ConditionCode
		flags     map[string]bool
		expected  bool
	}{
		{"EQ with Z=1", core.EQ, map[string]bool{"Z": true}, true},
		{"EQ with Z=0", core.EQ, map[string]bool{"Z": false}, false},
		{"NE with Z=0", core.NE, map[string]bool{"Z": false}, true},
		{"NE with Z=1", core.NE, map[string]bool{"Z": true}, false},
		{"CS with C=1", core.CS, map[string]bool{"C": true}, true},
		{"CS with C=0", core.CS, map[string]bool{"C": false}, false},
		{"CC with C=0", core.CC, map[string]bool{"C": false}, true},
		{"CC with C=1", core.CC, map[string]bool{"C": true}, false},
		{"MI with N=1", core.MI, map[string]bool{"N": true}, true},
		{"MI with N=0", core.MI, map[string]bool{"N": false}, false},
		{"PL with N=0", core.PL, map[string]bool{"N": false}, true},
		{"PL with N=1", core.PL, map[string]bool{"N": true}, false},
		{"VS with V=1", core.VS, map[string]bool{"V": true}, true},
		{"VS with V=0", core.VS, map[string]bool{"V": false}, false},
		{"VC with V=0", core.VC, map[string]bool{"V": false}, true},
		{"VC with V=1", core.VC, map[string]bool{"V": true}, false},
		{"HI with C=1,Z=0", core.HI, map[string]bool{"C": true, "Z": false}, true},
		{"HI with C=0,Z=0", core.HI, map[string]bool{"C": false, "Z": false}, false},
		{"HI with C=1,Z=1", core.HI, map[string]bool{"C": true, "Z": true}, false},
		{"LS with C=0", core.LS, map[string]bool{"C": false, "Z": false}, true},
		{"LS with Z=1", core.LS, map[string]bool{"C": true, "Z": true}, true},
		{"LS with C=1,Z=0", core.LS, map[string]bool{"C": true, "Z": false}, false},
		{"GE with N=V=0", core.GE, map[string]bool{"N": false, "V": false}, true},
		{"GE with N=V=1", core.GE, map[string]bool{"N": true, "V": true}, true},
		{"GE with N!=V", core.GE, map[string]bool{"N": true, "V": false}, false},
		{"LT with N!=V", core.LT, map[string]bool{"N": true, "V": false}, true},
		{"LT with N=V", core.LT, map[string]bool{"N": false, "V": false}, false},
		{"GT with Z=0,N=V", core.GT, map[string]bool{"Z": false, "N": false, "V": false}, true},
		{"GT with Z=1", core.GT, map[string]bool{"Z": true, "N": false, "V": false}, false},
		{"GT with N!=V", core.GT, map[string]bool{"Z": false, "N": true, "V": false}, false},
		{"LE with Z=1", core.LE, map[string]bool{"Z": true, "N": false, "V": false}, true},
		{"LE with N!=V", core.LE, map[string]bool{"Z": false, "N": true, "V": false}, true},
		{"LE with Z=0,N=V", core.LE, map[string]bool{"Z": false, "N": false, "V": false}, false},
		{"AL always true", core.AL, map[string]bool{}, true},
		{"NV always false", core.NV, map[string]bool{}, false},
	}

	executor := NewConditionalExecutor("TEST", "Test executor")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			for flag, value := range tt.flags {
				switch flag {
				case "N":
					s.SetN(value)
				case "Z":
					s.SetZ(value)
				case "C":
					s.SetC(value)
				case "V":
					s.SetV(value)
				}
			}

			result := executor.evaluateCondition(s, tt.condition)
			if result != tt.expected {
				t.Errorf("Expected condition %s to be %v, got %v", tt.condition.String(), tt.expected, result)
			}
		})
	}
}

func TestConditionalExecutor_StringConditionParsing(t *testing.T) {
	tests := []struct {
		instrStr string
		expected core.ConditionCode
	}{
		{"csel x0, x1, x2, eq", core.EQ},
		{"csel x0, x1, x2, ne", core.NE},
		{"csel x0, x1, x2, cs", core.CS},
		{"csel x0, x1, x2, hs", core.CS}, // HS is alias for CS
		{"csel x0, x1, x2, cc", core.CC},
		{"csel x0, x1, x2, lo", core.CC}, // LO is alias for CC
		{"csel x0, x1, x2, mi", core.MI},
		{"csel x0, x1, x2, pl", core.PL},
		{"csel x0, x1, x2, vs", core.VS},
		{"csel x0, x1, x2, vc", core.VC},
		{"csel x0, x1, x2, hi", core.HI},
		{"csel x0, x1, x2, ls", core.LS},
		{"csel x0, x1, x2, ge", core.GE},
		{"csel x0, x1, x2, lt", core.LT},
		{"csel x0, x1, x2, gt", core.GT},
		{"csel x0, x1, x2, le", core.LE},
		{"csel x0, x1, x2", core.AL}, // No condition defaults to AL
	}

	executor := NewConditionalExecutor("TEST", "Test executor")

	for _, tt := range tests {
		t.Run(tt.instrStr, func(t *testing.T) {
			result := executor.parseConditionFromString(tt.instrStr)
			if result != tt.expected {
				t.Errorf("Expected condition %s, got %s", tt.expected.String(), result.String())
			}
		})
	}
}

func TestConditionalExecutor_ErrorHandling(t *testing.T) {
	s := state.NewState()
	executor := NewConditionalExecutor("CSEL", "Conditional select")

	// Test insufficient operands
	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_CSEL,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
		},
	}

	err := executor.Execute(s, instr)
	if err == nil {
		t.Error("Expected error for insufficient operands")
	}

	// Test invalid register
	instr = &disassemble.Instruction{
		Operation: disassemble.ARM64_CSEL,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(999)}}, // Invalid register
			{Registers: []disassemble.Register{disassemble.Register(35)}},  // X1
			{Registers: []disassemble.Register{disassemble.Register(36)}},  // X2
		},
	}

	err = executor.Execute(s, instr)
	if err == nil {
		t.Error("Expected error for invalid register")
	}
}
