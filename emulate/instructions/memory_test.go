package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
)

// Register constants for testing
const (
	W0_REG = disassemble.Register(1)  // w0
	W1_REG = disassemble.Register(2)  // w1
	W2_REG = disassemble.Register(3)  // w2
	X0_REG = disassemble.Register(34) // x0
	X1_REG = disassemble.Register(35) // x1
	X2_REG = disassemble.Register(36) // x2
)

func TestMemoryExecutor_Basic(t *testing.T) {
	// Test that we can create memory executors
	executor := NewMemoryExecutor("LDR", "Load register")
	if executor == nil {
		t.Fatal("Failed to create memory executor")
	}

	if !executor.Supports("LDR") {
		t.Error("Executor should support LDR")
	}

	if executor.Supports("ADD") {
		t.Error("Executor should not support ADD")
	}
}

func TestMemoryExecutor_LDR_Simple(t *testing.T) {
	executor := NewMemoryExecutor("LDR", "Load register")
	state := NewMockState()

	// Set up memory with test data
	testAddr := uint64(0x1000)
	testValue := uint64(0x123456789ABCDEF0)
	state.WriteUint64(testAddr, testValue)

	// Set up base register
	state.SetX(1, testAddr)

	// Create instruction: LDR X0, [X1]
	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_LDR,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{X0_REG}}, // X0
			{Registers: []disassemble.Register{X1_REG}}, // X1
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("LDR execution failed: %v", err)
	}

	result := state.GetX(0)
	if result != testValue {
		t.Errorf("Expected X0 = 0x%x, got 0x%x", testValue, result)
	}
}

func TestMemoryExecutor_STR_Simple(t *testing.T) {
	executor := NewMemoryExecutor("STR", "Store register")
	state := NewMockState()

	// Set up test data
	testAddr := uint64(0x1000)
	testValue := uint64(0x123456789ABCDEF0)

	// Set up registers
	state.SetX(0, testValue)
	state.SetX(1, testAddr)

	// Create instruction: STR X0, [X1]
	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_STR,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{X0_REG}}, // X0
			{Registers: []disassemble.Register{X1_REG}}, // X1
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("STR execution failed: %v", err)
	}

	result, err := state.ReadUint64(testAddr)
	if err != nil {
		t.Fatalf("Failed to read memory: %v", err)
	}

	if result != testValue {
		t.Errorf("Expected memory[0x%x] = 0x%x, got 0x%x", testAddr, testValue, result)
	}
}

func TestMemoryExecutor_LDRB_Simple(t *testing.T) {
	executor := NewMemoryExecutor("LDRB", "Load register byte")
	state := NewMockState()

	// Set up memory with test data
	testAddr := uint64(0x1000)
	testValue := uint8(0xAB)
	state.WriteMemory(testAddr, []byte{testValue})

	// Set up base register
	state.SetX(1, testAddr)

	// Create instruction: LDRB W0, [X1]
	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_LDRB,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{W0_REG}}, // W0
			{Registers: []disassemble.Register{X1_REG}}, // X1
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("LDRB execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(testValue)
	if result != expected {
		t.Errorf("Expected X0 = 0x%x, got 0x%x", expected, result)
	}
}

func TestMemoryExecutor_LDP_Simple(t *testing.T) {
	executor := NewMemoryExecutor("LDP", "Load pair of registers")
	state := NewMockState()

	// Set up memory with test data
	testAddr := uint64(0x1000)
	testValue1 := uint64(0x123456789ABCDEF0)
	testValue2 := uint64(0xFEDCBA0987654321)
	state.WriteUint64(testAddr, testValue1)
	state.WriteUint64(testAddr+8, testValue2)

	// Set up base register
	state.SetX(2, testAddr)

	// Create instruction: LDP X0, X1, [X2]
	instr := &disassemble.Instruction{
		Operation: disassemble.ARM64_LDP,
		Operands: []disassemble.Operand{
			{Registers: []disassemble.Register{X0_REG}}, // X0
			{Registers: []disassemble.Register{X1_REG}}, // X1
			{Registers: []disassemble.Register{X2_REG}}, // X2
		},
	}

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("LDP execution failed: %v", err)
	}

	result1 := state.GetX(0)
	result2 := state.GetX(1)

	if result1 != testValue1 {
		t.Errorf("Expected X0 = 0x%x, got 0x%x", testValue1, result1)
	}
	if result2 != testValue2 {
		t.Errorf("Expected X1 = 0x%x, got 0x%x", testValue2, result2)
	}
}

func TestMemoryExecutor_LDADDA(t *testing.T) {
	executor := NewMemoryExecutor("LDADDA", "Atomic load-add with acquire semantics")

	t.Run("word", func(t *testing.T) {
		state := NewMockState()
		addr := uint64(0x2000)
		initial := uint32(0xFFFFFFF0)
		addend := uint32(0x30)

		state.WriteUint32(addr, initial)
		state.SetX(2, addr)
		state.SetW(1, addend)

		instr := &disassemble.Instruction{
			Operation: disassemble.ARM64_LDADDA,
			Operands: []disassemble.Operand{
				{Registers: []disassemble.Register{W0_REG}},
				{Registers: []disassemble.Register{W1_REG}},
				{Registers: []disassemble.Register{X2_REG}},
			},
		}

		if err := executor.Execute(state, instr); err != nil {
			t.Fatalf("LDADDA execution failed: %v", err)
		}

		if got := state.GetW(0); got != initial {
			t.Errorf("LDADDA should return original value, want 0x%x got 0x%x", initial, got)
		}

		mem, err := state.ReadUint32(addr)
		if err != nil {
			t.Fatalf("Failed to read memory: %v", err)
		}

		expectedMem := uint32(uint64(initial) + uint64(addend))
		if mem != expectedMem {
			t.Errorf("LDADDA should store updated value, want 0x%x got 0x%x", expectedMem, mem)
		}
	})

	t.Run("doubleword", func(t *testing.T) {
		state := NewMockState()
		addr := uint64(0x3000)
		initial := uint64(0x7FFFFFFFFFFFFFF0)
		addend := uint64(0x30)

		state.WriteUint64(addr, initial)
		state.SetX(2, addr)
		state.SetX(1, addend)

		instr := &disassemble.Instruction{
			Operation: disassemble.ARM64_LDADDA,
			Operands: []disassemble.Operand{
				{Registers: []disassemble.Register{X0_REG}},
				{Registers: []disassemble.Register{X1_REG}},
				{Registers: []disassemble.Register{X2_REG}},
			},
		}

		if err := executor.Execute(state, instr); err != nil {
			t.Fatalf("LDADDA execution failed: %v", err)
		}

		if got := state.GetX(0); got != initial {
			t.Errorf("LDADDA should return original value, want 0x%x got 0x%x", initial, got)
		}

		mem, err := state.ReadUint64(addr)
		if err != nil {
			t.Fatalf("Failed to read memory: %v", err)
		}

		expectedMem := initial + addend
		if mem != expectedMem {
			t.Errorf("LDADDA should store updated value, want 0x%x got 0x%x", expectedMem, mem)
		}
	})
}

func TestMemoryHelperFunctions(t *testing.T) {
	// Test IsLoadInstruction
	loadInstructions := []string{"LDR", "LDRB", "LDRH", "LDRSB", "LDRSH", "LDRSW", "LDUR", "LDP", "LDADDA"}
	for _, instr := range loadInstructions {
		if !IsLoadInstruction(instr) {
			t.Errorf("Expected %s to be a load instruction", instr)
		}
	}

	if IsLoadInstruction("STR") {
		t.Error("STR should not be a load instruction")
	}

	// Test IsStoreInstruction
	storeInstructions := []string{"STR", "STRB", "STRH", "STUR", "STP", "LDADDA"}
	for _, instr := range storeInstructions {
		if !IsStoreInstruction(instr) {
			t.Errorf("Expected %s to be a store instruction", instr)
		}
	}

	if IsStoreInstruction("LDR") {
		t.Error("LDR should not be a store instruction")
	}

	// Test IsPairInstruction
	pairInstructions := []string{"LDP", "STP"}
	for _, instr := range pairInstructions {
		if !IsPairInstruction(instr) {
			t.Errorf("Expected %s to be a pair instruction", instr)
		}
	}

	if IsPairInstruction("LDR") {
		t.Error("LDR should not be a pair instruction")
	}

	// Test GetMemoryAccessSize
	tests := []struct {
		mnemonic string
		is32Bit  bool
		expected int
	}{
		{"LDRB", false, 1},
		{"STRB", false, 1},
		{"LDRSB", false, 1},
		{"LDRH", false, 2},
		{"STRH", false, 2},
		{"LDRSH", false, 2},
		{"LDRSW", false, 4},
		{"LDR", true, 4},
		{"STR", true, 4},
		{"LDR", false, 8},
		{"STR", false, 8},
		{"LDADDA", true, 4},
		{"LDADDA", false, 8},
		{"LDP", true, 8},
		{"STP", true, 8},
		{"LDP", false, 16},
		{"STP", false, 16},
		{"UNKNOWN", false, 0},
	}

	for _, test := range tests {
		result := GetMemoryAccessSize(test.mnemonic, test.is32Bit)
		if result != test.expected {
			t.Errorf("GetMemoryAccessSize(%s, %v) = %d, expected %d",
				test.mnemonic, test.is32Bit, result, test.expected)
		}
	}
}

func TestRegisterMemoryInstructions(t *testing.T) {
	registry := NewRegistry()
	RegisterMemoryInstructions(registry)

	// Test that all memory instructions are registered
	expectedInstructions := []string{
		"LDR", "LDRB", "LDRH", "LDRSB", "LDRSH", "LDRSW", "LDUR",
		"LDADDA",
		"STR", "STRB", "STRH", "STUR",
		"LDP", "STP",
	}

	for _, instr := range expectedInstructions {
		if !registry.HasInstruction(instr) {
			t.Errorf("Expected instruction %s to be registered", instr)
		}
	}

	// Test that we can get executors
	executor, exists := registry.Get("LDR")
	if !exists {
		t.Error("Expected to find LDR executor")
	}
	if executor == nil {
		t.Error("Expected non-nil LDR executor")
	}
}
