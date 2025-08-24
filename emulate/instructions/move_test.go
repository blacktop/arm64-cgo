package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

func TestMoveExecutor_MOV(t *testing.T) {
	tests := []struct {
		name        string
		instruction *disassemble.Instruction
		setupState  func(s core.State)
		checkState  func(t *testing.T, s core.State)
		expectError bool
	}{
		{
			name: "MOV X1, X0 - register to register",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOV,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{X1_REG}},
					{Registers: []disassemble.Register{X0_REG}},
				},
			},
			setupState: func(s core.State) {
				s.SetX(0, 0x1234567890ABCDEF)
			},
			checkState: func(t *testing.T, s core.State) {
				if got := s.GetX(1); got != 0x1234567890ABCDEF {
					t.Errorf("Expected X1=0x1234567890ABCDEF, got 0x%x", got)
				}
			},
		},
		{
			name: "MOV W1, W0 - 32-bit register to register",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOV,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{W1_REG}},
					{Registers: []disassemble.Register{W0_REG}},
				},
			},
			setupState: func(s core.State) {
				s.SetX(0, 0x1234567890ABCDEF)
			},
			checkState: func(t *testing.T, s core.State) {
				if got := s.GetW(1); got != 0x90ABCDEF {
					t.Errorf("Expected W1=0x90ABCDEF, got 0x%x", got)
				}
				// Upper 32 bits should be zeroed
				if got := s.GetX(1); got != 0x90ABCDEF {
					t.Errorf("Expected X1=0x90ABCDEF (upper bits zeroed), got 0x%x", got)
				}
			},
		},
		{
			name: "MOV X1, #0x1234 - immediate to register",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOV,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{X1_REG}},
					{Immediate: 0x1234},
				},
			},
			setupState: func(s core.State) {},
			checkState: func(t *testing.T, s core.State) {
				if got := s.GetX(1); got != 0x1234 {
					t.Errorf("Expected X1=0x1234, got 0x%x", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewMoveExecutor("MOV", "Move register or immediate")
			s := state.NewState()
			tt.setupState(s)

			err := executor.Execute(s, tt.instruction)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !tt.expectError {
				tt.checkState(t, s)
			}
		})
	}
}

func TestMoveExecutor_MOVZ(t *testing.T) {
	tests := []struct {
		name        string
		instruction *disassemble.Instruction
		checkState  func(t *testing.T, s core.State)
		expectError bool
	}{
		{
			name: "MOVZ X1, #0x1234 - basic immediate",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOVZ,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{X1_REG}},
					{Immediate: 0x1234},
				},
			},
			checkState: func(t *testing.T, s core.State) {
				if got := s.GetX(1); got != 0x1234 {
					t.Errorf("Expected X1=0x1234, got 0x%x", got)
				}
			},
		},
		{
			name: "MOVZ X1, #0x1234, LSL #16 - shifted immediate",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOVZ,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{X1_REG}},
					{Immediate: 0x1234},
					{ShiftValueUsed: true, ShiftType: disassemble.SHIFT_TYPE_LSL, ShiftValue: 16},
				},
			},
			checkState: func(t *testing.T, s core.State) {
				expected := uint64(0x1234) << 16
				if got := s.GetX(1); got != expected {
					t.Errorf("Expected X1=0x%x, got 0x%x", expected, got)
				}
			},
		},
		{
			name: "MOVZ W1, #0x1234 - 32-bit operation",
			instruction: &disassemble.Instruction{
				Operation: disassemble.ARM64_MOVZ,
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{W1_REG}},
					{Immediate: 0x1234},
				},
			},
			checkState: func(t *testing.T, s core.State) {
				if got := s.GetW(1); got != 0x1234 {
					t.Errorf("Expected W1=0x1234, got 0x%x", got)
				}
				// Upper 32 bits should be zeroed
				if got := s.GetX(1); got != 0x1234 {
					t.Errorf("Expected X1=0x1234 (upper bits zeroed), got 0x%x", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewMoveExecutor("MOVZ", "Move wide with zero")
			s := state.NewState()

			err := executor.Execute(s, tt.instruction)
			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if !tt.expectError {
				tt.checkState(t, s)
			}
		})
	}
}

func TestMoveExecutor_HelperFunctions(t *testing.T) {
	t.Run("IsMoveInstruction", func(t *testing.T) {
		tests := []struct {
			mnemonic string
			expected bool
		}{
			{"MOV", true},
			{"MOVZ", true},
			{"MOVN", true},
			{"MOVK", true},
			{"ADR", true},
			{"ADRP", true},
			{"SXTB", true},
			{"SXTH", true},
			{"SXTW", true},
			{"UXTB", true},
			{"UXTH", true},
			{"ADD", false},
			{"SUB", false},
		}

		for _, tt := range tests {
			if got := IsMoveInstruction(tt.mnemonic); got != tt.expected {
				t.Errorf("IsMoveInstruction(%s) = %v, want %v", tt.mnemonic, got, tt.expected)
			}
		}
	})

	t.Run("IsWideMove", func(t *testing.T) {
		tests := []struct {
			mnemonic string
			expected bool
		}{
			{"MOVZ", true},
			{"MOVN", true},
			{"MOVK", true},
			{"MOV", false},
			{"ADR", false},
		}

		for _, tt := range tests {
			if got := IsWideMove(tt.mnemonic); got != tt.expected {
				t.Errorf("IsWideMove(%s) = %v, want %v", tt.mnemonic, got, tt.expected)
			}
		}
	})

	t.Run("ValidateWideMove", func(t *testing.T) {
		tests := []struct {
			mnemonic    string
			shiftAmount uint64
			expectError bool
		}{
			{"MOVZ", 0, false},
			{"MOVZ", 16, false},
			{"MOVZ", 32, false},
			{"MOVZ", 48, false},
			{"MOVZ", 8, true},  // Invalid: not multiple of 16
			{"MOVZ", 64, true}, // Invalid: too large
			{"MOV", 0, true},   // Invalid: not a wide move
		}

		for _, tt := range tests {
			err := ValidateWideMove(tt.mnemonic, tt.shiftAmount)
			if tt.expectError && err == nil {
				t.Errorf("ValidateWideMove(%s, %d) expected error but got none", tt.mnemonic, tt.shiftAmount)
			} else if !tt.expectError && err != nil {
				t.Errorf("ValidateWideMove(%s, %d) unexpected error: %v", tt.mnemonic, tt.shiftAmount, err)
			}
		}
	})
}

func TestRegisterMoveInstructions(t *testing.T) {
	registry := NewRegistry()
	RegisterMoveInstructions(registry)

	expectedInstructions := []string{
		"MOV", "MOVZ", "MOVN", "MOVK",
		"ADR", "ADRP",
		"SXTB", "SXTH", "SXTW", "UXTB", "UXTH",
	}

	for _, instr := range expectedInstructions {
		if _, exists := registry.Get(instr); !exists {
			t.Errorf("Instruction %s not registered", instr)
		}
	}
}
