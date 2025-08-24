package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

func TestCompareExecutor_CMP(t *testing.T) {
	tests := []struct {
		name    string
		val1    uint64
		val2    uint64
		is32Bit bool
		expectN bool
		expectZ bool
		expectC bool
		expectV bool
	}{
		{
			name:    "CMP equal values",
			val1:    10,
			val2:    10,
			expectN: false,
			expectZ: true,
			expectC: true, // No borrow
			expectV: false,
		},
		{
			name:    "CMP val1 > val2",
			val1:    20,
			val2:    10,
			expectN: false,
			expectZ: false,
			expectC: true, // No borrow
			expectV: false,
		},
		{
			name:    "CMP val1 < val2",
			val1:    10,
			val2:    20,
			expectN: true, // Result is negative
			expectZ: false,
			expectC: false, // Borrow occurred
			expectV: false,
		},
		{
			name:    "CMP signed overflow (positive - negative = negative)",
			val1:    0x7FFFFFFFFFFFFFFF, // Max positive int64
			val2:    0x8000000000000000, // Min negative int64
			expectN: true,               // Result is negative
			expectZ: false,
			expectC: false, // Borrow occurred
			expectV: true,  // Signed overflow
		},
		{
			name:    "CMP 32-bit equal values",
			val1:    10,
			val2:    10,
			is32Bit: true,
			expectN: false,
			expectZ: true,
			expectC: true, // No borrow
			expectV: false,
		},
		{
			name:    "CMP 32-bit signed overflow",
			val1:    0x7FFFFFFF, // Max positive int32
			val2:    0x80000000, // Min negative int32
			is32Bit: true,
			expectN: true, // Result is negative
			expectZ: false,
			expectC: false, // Borrow occurred
			expectV: true,  // Signed overflow
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewCompareExecutor("CMP", "Compare")
			state := state.NewState()

			// Set up registers
			state.SetX(0, tt.val1)
			state.SetX(1, tt.val2)

			// Create instruction
			var reg0, reg1 disassemble.Register
			if tt.is32Bit {
				reg0 = disassemble.Register(1) // W0
				reg1 = disassemble.Register(2) // W1
			} else {
				reg0 = disassemble.Register(34) // X0
				reg1 = disassemble.Register(35) // X1
			}

			instr := &disassemble.Instruction{
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{reg0}},
					{Registers: []disassemble.Register{reg1}},
				},
			}

			// Execute instruction
			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			// Check flags
			if state.GetN() != tt.expectN {
				t.Errorf("Expected N=%v, got N=%v", tt.expectN, state.GetN())
			}
			if state.GetZ() != tt.expectZ {
				t.Errorf("Expected Z=%v, got Z=%v", tt.expectZ, state.GetZ())
			}
			if state.GetC() != tt.expectC {
				t.Errorf("Expected C=%v, got C=%v", tt.expectC, state.GetC())
			}
			if state.GetV() != tt.expectV {
				t.Errorf("Expected V=%v, got V=%v", tt.expectV, state.GetV())
			}
		})
	}
}

func TestCompareExecutor_CMP_Immediate(t *testing.T) {
	executor := NewCompareExecutor("CMP", "Compare")
	state := state.NewState()

	// Set up register
	state.SetX(0, 15)

	// Create instruction: CMP X0, #10
	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Immediate: 10},
		},
	}

	// Execute instruction
	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Check flags (15 - 10 = 5, positive result)
	if state.GetN() {
		t.Error("Expected N=false for positive result")
	}
	if state.GetZ() {
		t.Error("Expected Z=false for non-zero result")
	}
	if !state.GetC() {
		t.Error("Expected C=true for no borrow")
	}
	if state.GetV() {
		t.Error("Expected V=false for no overflow")
	}
}

func TestCompareExecutor_CMN(t *testing.T) {
	tests := []struct {
		name    string
		val1    uint64
		val2    uint64
		is32Bit bool
		expectN bool
		expectZ bool
		expectC bool
		expectV bool
	}{
		{
			name:    "CMN normal addition",
			val1:    10,
			val2:    5,
			expectN: false,
			expectZ: false,
			expectC: false, // No carry
			expectV: false,
		},
		{
			name:    "CMN result zero",
			val1:    0xFFFFFFFFFFFFFFF6, // -10 in two's complement
			val2:    10,
			expectN: false,
			expectZ: true, // Result is zero
			expectC: true, // Carry occurred
			expectV: false,
		},
		{
			name:    "CMN unsigned overflow",
			val1:    0xFFFFFFFFFFFFFFFF, // Max uint64
			val2:    1,
			expectN: false,
			expectZ: true, // Wraps to zero
			expectC: true, // Carry occurred
			expectV: false,
		},
		{
			name:    "CMN signed overflow",
			val1:    0x7FFFFFFFFFFFFFFF, // Max positive int64
			val2:    1,
			expectN: true, // Result is negative due to overflow
			expectZ: false,
			expectC: false, // No carry in 64-bit addition
			expectV: true,  // Signed overflow
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewCompareExecutor("CMN", "Compare negative")
			state := state.NewState()

			// Set up registers
			state.SetX(0, tt.val1)
			state.SetX(1, tt.val2)

			// Create instruction
			var reg0, reg1 disassemble.Register
			if tt.is32Bit {
				reg0 = disassemble.Register(1) // W0
				reg1 = disassemble.Register(2) // W1
			} else {
				reg0 = disassemble.Register(34) // X0
				reg1 = disassemble.Register(35) // X1
			}

			instr := &disassemble.Instruction{
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{reg0}},
					{Registers: []disassemble.Register{reg1}},
				},
			}

			// Execute instruction
			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			// Check flags
			if state.GetN() != tt.expectN {
				t.Errorf("Expected N=%v, got N=%v", tt.expectN, state.GetN())
			}
			if state.GetZ() != tt.expectZ {
				t.Errorf("Expected Z=%v, got Z=%v", tt.expectZ, state.GetZ())
			}
			if state.GetC() != tt.expectC {
				t.Errorf("Expected C=%v, got C=%v", tt.expectC, state.GetC())
			}
			if state.GetV() != tt.expectV {
				t.Errorf("Expected V=%v, got V=%v", tt.expectV, state.GetV())
			}
		})
	}
}

func TestCompareExecutor_TST(t *testing.T) {
	tests := []struct {
		name    string
		val1    uint64
		val2    uint64
		is32Bit bool
		expectN bool
		expectZ bool
	}{
		{
			name:    "TST with result zero",
			val1:    0x0F0F0F0F0F0F0F0F,
			val2:    0xF0F0F0F0F0F0F0F0,
			expectN: false,
			expectZ: true, // AND result is zero
		},
		{
			name:    "TST with positive result",
			val1:    0x0F0F0F0F0F0F0F0F,
			val2:    0x0F0F0F0F0F0F0F0F,
			expectN: false,
			expectZ: false, // AND result is non-zero and positive
		},
		{
			name:    "TST with negative result (bit 63 set)",
			val1:    0x8000000000000000,
			val2:    0xFFFFFFFFFFFFFFFF,
			expectN: true, // Bit 63 is set in result
			expectZ: false,
		},
		{
			name:    "TST 32-bit with negative result (bit 31 set)",
			val1:    0x80000000,
			val2:    0xFFFFFFFF,
			is32Bit: true,
			expectN: true, // Bit 31 is set in 32-bit result
			expectZ: false,
		},
		{
			name:    "TST 32-bit with zero result",
			val1:    0x0F0F0F0F,
			val2:    0xF0F0F0F0,
			is32Bit: true,
			expectN: false,
			expectZ: true, // AND result is zero
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewCompareExecutor("TST", "Test bits")
			state := state.NewState()

			// Set up registers
			state.SetX(0, tt.val1)
			state.SetX(1, tt.val2)

			// Create instruction
			var reg0, reg1 disassemble.Register
			if tt.is32Bit {
				reg0 = disassemble.Register(1) // W0
				reg1 = disassemble.Register(2) // W1
			} else {
				reg0 = disassemble.Register(34) // X0
				reg1 = disassemble.Register(35) // X1
			}

			instr := &disassemble.Instruction{
				Operands: []disassemble.Operand{
					{Registers: []disassemble.Register{reg0}},
					{Registers: []disassemble.Register{reg1}},
				},
			}

			// Execute instruction
			err := executor.Execute(state, instr)
			if err != nil {
				t.Fatalf("Execute failed: %v", err)
			}

			// Check flags (TST only affects N and Z flags)
			if state.GetN() != tt.expectN {
				t.Errorf("Expected N=%v, got N=%v", tt.expectN, state.GetN())
			}
			if state.GetZ() != tt.expectZ {
				t.Errorf("Expected Z=%v, got Z=%v", tt.expectZ, state.GetZ())
			}
		})
	}
}

func TestCompareExecutor_TST_Immediate(t *testing.T) {
	executor := NewCompareExecutor("TST", "Test bits")
	state := state.NewState()

	// Set up register
	state.SetX(0, 0xFF00FF00FF00FF00)

	// Create instruction: TST X0, #0x00FF00FF00FF00FF
	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Immediate: 0x00FF00FF00FF00FF},
		},
	}

	// Execute instruction
	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Check flags (AND result should be zero)
	if state.GetN() {
		t.Error("Expected N=false")
	}
	if !state.GetZ() {
		t.Error("Expected Z=true for zero result")
	}
}

func TestCompareExecutor_CCMP(t *testing.T) {
	executor := NewCompareExecutor("CCMP", "Conditional compare")
	state := state.NewState()

	// Set up registers
	state.SetX(0, 10)
	state.SetX(1, 5)

	// Set flags to make condition true (for this test, we'll assume AL condition)
	state.SetZ(false)

	// Create instruction: CCMP X0, X1, #0, AL
	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
			{Immediate: 0}, // NZCV immediate
			{},             // Condition placeholder
		},
	}

	// Execute instruction
	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Since condition is always true (AL), should perform normal CMP
	// 10 - 5 = 5 (positive, non-zero, no borrow, no overflow)
	if state.GetN() {
		t.Error("Expected N=false for positive result")
	}
	if state.GetZ() {
		t.Error("Expected Z=false for non-zero result")
	}
	if !state.GetC() {
		t.Error("Expected C=true for no borrow")
	}
	if state.GetV() {
		t.Error("Expected V=false for no overflow")
	}
}

func TestCompareExecutor_CCMN(t *testing.T) {
	executor := NewCompareExecutor("CCMN", "Conditional compare negative")
	state := state.NewState()

	// Set up registers
	state.SetX(0, 10)
	state.SetX(1, 5)

	// Set flags to make condition true (for this test, we'll assume AL condition)
	state.SetZ(false)

	// Create instruction: CCMN X0, X1, #0, AL
	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
			{Immediate: 0}, // NZCV immediate
			{},             // Condition placeholder
		},
	}

	// Execute instruction
	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Since condition is always true (AL), should perform normal CMN
	// 10 + 5 = 15 (positive, non-zero, no carry, no overflow)
	if state.GetN() {
		t.Error("Expected N=false for positive result")
	}
	if state.GetZ() {
		t.Error("Expected Z=false for non-zero result")
	}
	if state.GetC() {
		t.Error("Expected C=false for no carry")
	}
	if state.GetV() {
		t.Error("Expected V=false for no overflow")
	}
}

func TestCompareExecutor_InvalidOperands(t *testing.T) {
	executor := NewCompareExecutor("CMP", "Compare")
	state := state.NewState()

	// Test with insufficient operands
	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
		},
	}

	err := executor.Execute(state, instr)
	if err == nil {
		t.Error("Expected error for insufficient operands")
	}

	emulErr, ok := err.(*core.EmulationError)
	if !ok {
		t.Errorf("Expected EmulationError, got %T", err)
	}
	if emulErr.Err != core.ErrInvalidInstruction {
		t.Errorf("Expected ErrInvalidInstruction, got %v", emulErr.Err)
	}
}

func TestCompareExecutor_UnsupportedInstruction(t *testing.T) {
	executor := NewCompareExecutor("INVALID", "Invalid instruction")
	state := state.NewState()

	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
		},
	}

	err := executor.Execute(state, instr)
	if err == nil {
		t.Error("Expected error for unsupported instruction")
	}

	emulErr, ok := err.(*core.EmulationError)
	if !ok {
		t.Errorf("Expected EmulationError, got %T", err)
	}
	if emulErr.Err != core.ErrUnsupportedFeature {
		t.Errorf("Expected ErrUnsupportedFeature, got %v", emulErr.Err)
	}
}

func TestRegisterCompareInstructions(t *testing.T) {
	registry := NewRegistry()
	RegisterCompareInstructions(registry)

	// Test that all compare instructions are registered
	// Note: TST is registered in logical instructions as it's fundamentally AND with flags
	expectedInstructions := []string{"CMP", "CMN", "CCMP", "CCMN"}

	for _, instr := range expectedInstructions {
		executor, exists := registry.Get(instr)
		if !exists {
			t.Errorf("Instruction %s not registered", instr)
			continue
		}

		if !executor.Supports(instr) {
			t.Errorf("Executor for %s does not support the instruction", instr)
		}
	}
}

// Benchmark tests
func BenchmarkCompareExecutor_CMP(b *testing.B) {
	executor := NewCompareExecutor("CMP", "Compare")
	state := state.NewState()
	state.SetX(0, 100)
	state.SetX(1, 50)

	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = executor.Execute(state, instr)
	}
}

func BenchmarkCompareExecutor_TST(b *testing.B) {
	executor := NewCompareExecutor("TST", "Test bits")
	state := state.NewState()
	state.SetX(0, 0xFF00FF00FF00FF00)
	state.SetX(1, 0x00FF00FF00FF00FF)

	instr := &disassemble.Instruction{
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{disassemble.Register(34)}}, // X0
			{Registers: []disassemble.Register{disassemble.Register(35)}}, // X1
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = executor.Execute(state, instr)
	}
}
