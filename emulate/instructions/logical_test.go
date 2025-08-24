package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

func TestLogicalExecutor_AND(t *testing.T) {
	tests := []struct {
		name     string
		val1     uint64
		val2     uint64
		expected uint64
		is32bit  bool
	}{
		{"64-bit AND all ones", 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, false},
		{"64-bit AND with zero", 0xFFFFFFFFFFFFFFFF, 0x0000000000000000, 0x0000000000000000, false},
		{"64-bit AND pattern", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, 0x0000000000000000, false},
		{"64-bit AND mixed", 0x123456789ABCDEF0, 0xF0E1D2C3B4A59687, 0x1020524090A49680, false},
		{"32-bit AND", 0x12345678, 0x87654321, 0x02244220, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("AND", "Bitwise AND")

			// Create mock instruction
			var instr *disassemble.Instruction
			if tt.is32bit {
				instr = createMockInstruction("AND", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_W0),
					createRegisterOperand(disassemble.REG_W1),
					createRegisterOperand(disassemble.REG_W2),
				})
			} else {
				instr = createMockInstruction("AND", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_X0),
					createRegisterOperand(disassemble.REG_X1),
					createRegisterOperand(disassemble.REG_X2),
				})
			}

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			var result uint64
			if tt.is32bit {
				result = uint64(s.GetW(0))
			} else {
				result = s.GetX(0)
			}

			if result != tt.expected {
				t.Errorf("Expected 0x%X, got 0x%X", tt.expected, result)
			}
		})
	}
}

func TestLogicalExecutor_ANDS(t *testing.T) {
	tests := []struct {
		name     string
		val1     uint64
		val2     uint64
		expected uint64
		expectN  bool
		expectZ  bool
		is32bit  bool
	}{
		{"64-bit ANDS zero result", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, 0x0000000000000000, false, true, false},
		{"64-bit ANDS negative result", 0xFFFFFFFFFFFFFFFF, 0x8000000000000000, 0x8000000000000000, true, false, false},
		{"64-bit ANDS positive result", 0x7FFFFFFFFFFFFFFF, 0x1234567890ABCDEF, 0x1234567890ABCDEF, false, false, false},
		{"32-bit ANDS negative", 0xFFFFFFFF, 0x80000000, 0x80000000, true, false, true},
		{"32-bit ANDS zero", 0x12345678, 0x87654321, 0x02244220, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("ANDS", "Bitwise AND and set flags")

			// Create mock instruction
			var instr *disassemble.Instruction
			if tt.is32bit {
				instr = createMockInstruction("ANDS", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_W0),
					createRegisterOperand(disassemble.REG_W1),
					createRegisterOperand(disassemble.REG_W2),
				})
			} else {
				instr = createMockInstruction("ANDS", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_X0),
					createRegisterOperand(disassemble.REG_X1),
					createRegisterOperand(disassemble.REG_X2),
				})
			}

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			var result uint64
			if tt.is32bit {
				result = uint64(s.GetW(0))
			} else {
				result = s.GetX(0)
			}

			if result != tt.expected {
				t.Errorf("Expected result 0x%X, got 0x%X", tt.expected, result)
			}

			if s.GetN() != tt.expectN {
				t.Errorf("Expected N flag %v, got %v", tt.expectN, s.GetN())
			}

			if s.GetZ() != tt.expectZ {
				t.Errorf("Expected Z flag %v, got %v", tt.expectZ, s.GetZ())
			}

			// C and V should always be cleared for logical operations
			if s.GetC() != false {
				t.Errorf("Expected C flag false, got %v", s.GetC())
			}

			if s.GetV() != false {
				t.Errorf("Expected V flag false, got %v", s.GetV())
			}
		})
	}
}

func TestLogicalExecutor_ORR(t *testing.T) {
	tests := []struct {
		name     string
		val1     uint64
		val2     uint64
		expected uint64
	}{
		{"ORR all zeros", 0x0000000000000000, 0x0000000000000000, 0x0000000000000000},
		{"ORR with all ones", 0x0000000000000000, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF},
		{"ORR pattern", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, 0xFFFFFFFFFFFFFFFF},
		{"ORR mixed", 0x123456789ABCDEF0, 0xF0E1D2C3B4A59687, 0xF2F5D6FBBEBDDEF7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("ORR", "Bitwise OR")

			instr := createMockInstruction("ORR", []disassemble.Operand{
				createRegisterOperand(disassemble.REG_X0),
				createRegisterOperand(disassemble.REG_X1),
				createRegisterOperand(disassemble.REG_X2),
			})

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected 0x%X, got 0x%X", tt.expected, result)
			}
		})
	}
}

func TestLogicalExecutor_EOR(t *testing.T) {
	tests := []struct {
		name     string
		val1     uint64
		val2     uint64
		expected uint64
	}{
		{"EOR same values", 0x123456789ABCDEF0, 0x123456789ABCDEF0, 0x0000000000000000},
		{"EOR with zero", 0x123456789ABCDEF0, 0x0000000000000000, 0x123456789ABCDEF0},
		{"EOR with all ones", 0x123456789ABCDEF0, 0xFFFFFFFFFFFFFFFF, 0xEDCBA9876543210F},
		{"EOR pattern", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, 0xFFFFFFFFFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("EOR", "Bitwise exclusive OR")

			instr := createMockInstruction("EOR", []disassemble.Operand{
				createRegisterOperand(disassemble.REG_X0),
				createRegisterOperand(disassemble.REG_X1),
				createRegisterOperand(disassemble.REG_X2),
			})

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected 0x%X, got 0x%X", tt.expected, result)
			}
		})
	}
}

func TestLogicalExecutor_BIC(t *testing.T) {
	tests := []struct {
		name     string
		val1     uint64
		val2     uint64
		expected uint64
	}{
		{"BIC clear all", 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0x0000000000000000},
		{"BIC clear none", 0xFFFFFFFFFFFFFFFF, 0x0000000000000000, 0xFFFFFFFFFFFFFFFF},
		{"BIC pattern", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, 0xAAAAAAAAAAAAAAAA},
		{"BIC mixed", 0x123456789ABCDEF0, 0xF0E1D2C3B4A59687, 0x021404380A184870},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("BIC", "Bit clear")

			instr := createMockInstruction("BIC", []disassemble.Operand{
				createRegisterOperand(disassemble.REG_X0),
				createRegisterOperand(disassemble.REG_X1),
				createRegisterOperand(disassemble.REG_X2),
			})

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			result := s.GetX(0)
			if result != tt.expected {
				t.Errorf("Expected 0x%X, got 0x%X", tt.expected, result)
			}
		})
	}
}

func TestLogicalExecutor_TST(t *testing.T) {
	tests := []struct {
		name    string
		val1    uint64
		val2    uint64
		expectN bool
		expectZ bool
		is32bit bool
	}{
		{"TST zero result", 0xAAAAAAAAAAAAAAAA, 0x5555555555555555, false, true, false},
		{"TST negative result", 0xFFFFFFFFFFFFFFFF, 0x8000000000000000, true, false, false},
		{"TST positive result", 0x7FFFFFFFFFFFFFFF, 0x1234567890ABCDEF, false, false, false},
		{"TST 32-bit negative", 0xFFFFFFFF, 0x80000000, true, false, true},
		{"TST 32-bit zero", 0x12345678, 0x87654321, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state.NewState()
			s.SetX(1, tt.val1)
			s.SetX(2, tt.val2)

			executor := NewLogicalExecutor("TST", "Test bits")

			var instr *disassemble.Instruction
			if tt.is32bit {
				instr = createMockInstruction("TST", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_W1),
					createRegisterOperand(disassemble.REG_W2),
				})
			} else {
				instr = createMockInstruction("TST", []disassemble.Operand{
					createRegisterOperand(disassemble.REG_X1),
					createRegisterOperand(disassemble.REG_X2),
				})
			}

			err := executor.Execute(s, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			if s.GetN() != tt.expectN {
				t.Errorf("Expected N flag %v, got %v", tt.expectN, s.GetN())
			}

			if s.GetZ() != tt.expectZ {
				t.Errorf("Expected Z flag %v, got %v", tt.expectZ, s.GetZ())
			}

			// C and V should always be cleared for TST
			if s.GetC() != false {
				t.Errorf("Expected C flag false, got %v", s.GetC())
			}

			if s.GetV() != false {
				t.Errorf("Expected V flag false, got %v", s.GetV())
			}
		})
	}
}

func TestLogicalExecutor_WithImmediate(t *testing.T) {
	s := state.NewState()
	s.SetX(1, 0x123456789ABCDEF0)

	executor := NewLogicalExecutor("AND", "Bitwise AND")

	instr := createMockInstruction("AND", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createImmediateOperand(0xFF00FF00FF00FF00),
	})

	err := executor.Execute(s, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	result := s.GetX(0)
	expected := uint64(0x120056009A00DE00)
	if result != expected {
		t.Errorf("Expected 0x%X, got 0x%X", expected, result)
	}
}


func TestBitfieldExecutor_BFM(t *testing.T) {
	s := state.NewState()
	s.SetX(0, 0xAAAAAAAAAAAAAAAA) // destination
	s.SetX(1, 0x5555555555555555) // source

	executor := NewBitfieldExecutor("BFM", "Bitfield Move")

	// Mock instruction with encoding for BFM X0, X1, #0, #7
	// This should insert bits 0-7 from X1 into X0
	instr := createMockInstruction("BFM", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
	})
	instr.Raw = 0xB3400020 // BFM X0, X1, #0, #7 (64-bit, N=1, immr=0, imms=7)

	err := executor.Execute(s, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	result := s.GetX(0)
	// Should have bits 0-7 from source (0x55) and rest from destination
	expected := uint64(0xAAAAAAAAAAAAAAAB)
	if result != expected {
		t.Errorf("Expected 0x%X, got 0x%X", expected, result)
	}
}

func TestBitfieldExecutor_UBFM(t *testing.T) {
	s := state.NewState()
	s.SetX(1, 0x123456789ABCDEF0)

	executor := NewBitfieldExecutor("UBFM", "Unsigned Bitfield Move")

	// Mock instruction with encoding for UBFX X0, X1, #8, #8
	// This should extract 8 bits starting at bit 8
	instr := createMockInstruction("UBFM", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
	})
	instr.Raw = 0xD3483C20 // UBFM X0, X1, #8, #15 (extract 8 bits from bit 8)

	err := executor.Execute(s, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	result := s.GetX(0)
	// Should extract bits 8-15 which is 0xDE from the source 0x123456789ABCDEF0
	expected := uint64(0xDE)
	if result != expected {
		t.Errorf("Expected 0x%X, got 0x%X", expected, result)
	}
}

func TestBitfieldExecutor_SBFM(t *testing.T) {
	s := state.NewState()
	s.SetX(1, 0x123456789ABCDE80) // Source with sign bit set in extracted field

	executor := NewBitfieldExecutor("SBFM", "Signed Bitfield Move")

	// Mock instruction with encoding for SBFX X0, X1, #4, #4
	// This should extract 4 bits starting at bit 4 with sign extension
	instr := createMockInstruction("SBFM", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
	})
	instr.Raw = 0x93441C20 // SBFM X0, X1, #4, #7 (extract 4 bits from bit 4 with sign extension)

	err := executor.Execute(s, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	result := s.GetX(0)
	// Should extract bits 4-7 (0x8) and sign extend since MSB of extracted field is set
	expected := uint64(0xFFFFFFFFFFFFFFF8)
	if result != expected {
		t.Errorf("Expected 0x%X, got 0x%X", expected, result)
	}
}

func TestRegisterLogicalInstructions(t *testing.T) {
	registry := NewRegistry()
	RegisterLogicalInstructions(registry)

	// Test that all logical instructions are registered
	expectedInstructions := []string{
		"AND", "ANDS", "ORR", "EOR", "BIC", "BICS", "TST",
		"BFM", "BFI", "BFXIL",
		"SBFM", "SBFIZ", "SBFX", "ASR", "SXTB", "SXTH", "SXTW",
		"UBFM", "UBFIZ", "UBFX", "LSL", "LSR", "UXTB", "UXTH",
	}

	for _, instr := range expectedInstructions {
		if !registry.HasInstruction(instr) {
			t.Errorf("Instruction %s not registered", instr)
		}
	}

	// Test that we can get executors for each instruction
	for _, instr := range expectedInstructions {
		executor, exists := registry.Get(instr)
		if !exists {
			t.Errorf("Could not get executor for %s", instr)
		}
		if executor == nil {
			t.Errorf("Executor for %s is nil", instr)
		}
	}
}

func TestLogicalExecutor_ErrorHandling(t *testing.T) {
	s := state.NewState()
	executor := NewLogicalExecutor("AND", "Bitwise AND")

	// Test with insufficient operands
	instr := createMockInstruction("AND", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
	})

	err := executor.Execute(s, instr)
	if err == nil {
		t.Error("Expected error for insufficient operands")
	}

	// Test with invalid register
	instr = createMockInstruction("AND", []disassemble.Operand{
		{Registers: []disassemble.Register{disassemble.Register(999)}},
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err = executor.Execute(s, instr)
	if err == nil {
		t.Error("Expected error for invalid register")
	}
}
