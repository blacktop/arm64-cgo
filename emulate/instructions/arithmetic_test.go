package instructions

import (
	"encoding/binary"
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// MockState implements the core.State interface for testing
type MockState struct {
	registers [31]uint64
	sp        uint64
	pc        uint64
	flags     struct {
		n, z, c, v bool
	}
	memory map[uint64][]byte
}

func NewMockState() *MockState {
	return &MockState{
		registers: [31]uint64{},
		sp:        0x7ffff0000000,
		pc:        0x100000000,
		memory:    make(map[uint64][]byte),
	}
}

func (s *MockState) GetX(reg int) uint64 {
	if reg < 0 || reg >= 31 {
		return 0
	}
	return s.registers[reg]
}

func (s *MockState) SetX(reg int, val uint64) {
	if reg >= 0 && reg < 31 {
		s.registers[reg] = val
	}
}

func (s *MockState) GetW(reg int) uint32 {
	if reg < 0 || reg >= 31 {
		return 0
	}
	return uint32(s.registers[reg])
}

func (s *MockState) SetW(reg int, val uint32) {
	if reg >= 0 && reg < 31 {
		s.registers[reg] = uint64(val)
	}
}

func (s *MockState) GetPC() uint64     { return s.pc }
func (s *MockState) SetPC(addr uint64) { s.pc = addr }
func (s *MockState) GetSP() uint64     { return s.sp }
func (s *MockState) SetSP(addr uint64) { s.sp = addr }

func (s *MockState) GetN() bool    { return s.flags.n }
func (s *MockState) SetN(val bool) { s.flags.n = val }
func (s *MockState) GetZ() bool    { return s.flags.z }
func (s *MockState) SetZ(val bool) { s.flags.z = val }
func (s *MockState) GetC() bool    { return s.flags.c }
func (s *MockState) SetC(val bool) { s.flags.c = val }
func (s *MockState) GetV() bool    { return s.flags.v }
func (s *MockState) SetV(val bool) { s.flags.v = val }

func (s *MockState) UpdateFlags(result uint64, isNegative bool) {
	s.flags.z = result == 0
	s.flags.n = isNegative
}

// Memory operations
func (s *MockState) ReadMemory(addr uint64, size int) ([]byte, error) {
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		if mem, exists := s.memory[addr+uint64(i)]; exists && len(mem) > 0 {
			data[i] = mem[0]
		}
	}
	return data, nil
}

func (s *MockState) WriteMemory(addr uint64, data []byte) {
	for i, b := range data {
		s.memory[addr+uint64(i)] = []byte{b}
	}
}

func (s *MockState) ReadUint64(addr uint64) (uint64, error) {
	data, err := s.ReadMemory(addr, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(data), nil
}

func (s *MockState) ReadUint32(addr uint64) (uint32, error) {
	data, err := s.ReadMemory(addr, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(data), nil
}

func (s *MockState) WriteUint64(addr uint64, value uint64) {
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	s.WriteMemory(addr, data)
}

func (s *MockState) WriteUint32(addr uint64, value uint32) {
	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)
	s.WriteMemory(addr, data)
}
func (s *MockState) ReadString(addr uint64) (string, error)                       { return "", nil }
func (s *MockState) ReadStringWithLength(addr uint64, length int) (string, error) { return "", nil }
func (s *MockState) ResolvePointer(addr uint64) (uint64, error)                   { return 0, nil }
func (s *MockState) FollowPointerChain(addr uint64, maxDepth int) (uint64, error) { return 0, nil }
func (s *MockState) ReadStringAtPointer(ptrAddr uint64) (string, error)           { return "", nil }
func (s *MockState) SetMemoryReadHandler(handler core.MemoryReadHandler)          {}
func (s *MockState) SetStringPoolHandler(handler core.StringPoolHandler)          {}
func (s *MockState) SetPointerResolver(resolver core.PointerResolver)             {}
func (s *MockState) Clone() core.State                                            { return NewMockState() }
func (s *MockState) Reset()                                                       {}

// Helper function to create a mock instruction
func createMockInstruction(operation string, operands []disassemble.Operand) *disassemble.Instruction {
	// For testing, we'll create a minimal instruction with the operation as a string
	// The actual operation field will be set to a valid operation constant
	var op = disassemble.ARM64_ADD // Default to ADD
	switch operation {
	case "ADD":
		op = disassemble.ARM64_ADD
	case "ADDS":
		op = disassemble.ARM64_ADDS
	case "SUB":
		op = disassemble.ARM64_SUB
	case "SUBS":
		op = disassemble.ARM64_SUBS
	case "MUL":
		op = disassemble.ARM64_MUL
	case "UDIV":
		op = disassemble.ARM64_UDIV
	case "SDIV":
		op = disassemble.ARM64_SDIV
	case "MADD":
		op = disassemble.ARM64_MADD
	case "MSUB":
		op = disassemble.ARM64_MSUB
	// Logical operations
	case "AND":
		op = disassemble.ARM64_AND
	case "ANDS":
		op = disassemble.ARM64_ANDS
	case "ORR":
		op = disassemble.ARM64_ORR
	case "EOR":
		op = disassemble.ARM64_EOR
	case "BIC":
		op = disassemble.ARM64_BIC
	case "BICS":
		op = disassemble.ARM64_BICS
	case "TST":
		op = disassemble.ARM64_TST
	case "BFM":
		op = disassemble.ARM64_BFM
	case "SBFM":
		op = disassemble.ARM64_SBFM
	case "UBFM":
		op = disassemble.ARM64_UBFM
	default:
		op = disassemble.ARM64_ERROR // Use error for invalid operations
	}

	return &disassemble.Instruction{
		Operation: op,
		Operands:  operands,
	}
}

// Helper function to create register operand
func createRegisterOperand(regID disassemble.Register) disassemble.Operand {
	return disassemble.Operand{
		Registers: []disassemble.Register{regID},
	}
}

// Helper function to create immediate operand
func createImmediateOperand(value uint64) disassemble.Operand {
	return disassemble.Operand{
		Immediate: value,
	}
}

func TestArithmeticExecutor_ADD(t *testing.T) {
	executor := NewArithmeticExecutor("ADD", "Add")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 10) // X1 = 10
	state.SetX(2, 5)  // X2 = 5

	// Test ADD X0, X1, X2 (X0 = X1 + X2 = 15)
	instr := createMockInstruction("ADD", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("ADD execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(15)
	if result != expected {
		t.Errorf("ADD result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_ADD_Immediate(t *testing.T) {
	executor := NewArithmeticExecutor("ADD", "Add")
	state := NewMockState()

	// Set up initial register value
	state.SetX(1, 10) // X1 = 10

	// Test ADD X0, X1, #5 (X0 = X1 + 5 = 15)
	instr := createMockInstruction("ADD", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createImmediateOperand(5),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("ADD immediate execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(15)
	if result != expected {
		t.Errorf("ADD immediate result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_SUB(t *testing.T) {
	executor := NewArithmeticExecutor("SUB", "Subtract")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 10) // X1 = 10
	state.SetX(2, 3)  // X2 = 3

	// Test SUB X0, X1, X2 (X0 = X1 - X2 = 7)
	instr := createMockInstruction("SUB", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("SUB execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(7)
	if result != expected {
		t.Errorf("SUB result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_MUL(t *testing.T) {
	executor := NewArithmeticExecutor("MUL", "Multiply")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 6) // X1 = 6
	state.SetX(2, 7) // X2 = 7

	// Test MUL X0, X1, X2 (X0 = X1 * X2 = 42)
	instr := createMockInstruction("MUL", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("MUL execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(42)
	if result != expected {
		t.Errorf("MUL result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_UDIV(t *testing.T) {
	executor := NewArithmeticExecutor("UDIV", "Unsigned divide")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 20) // X1 = 20
	state.SetX(2, 4)  // X2 = 4

	// Test UDIV X0, X1, X2 (X0 = X1 / X2 = 5)
	instr := createMockInstruction("UDIV", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("UDIV execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(5)
	if result != expected {
		t.Errorf("UDIV result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_UDIV_DivisionByZero(t *testing.T) {
	executor := NewArithmeticExecutor("UDIV", "Unsigned divide")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 20) // X1 = 20
	state.SetX(2, 0)  // X2 = 0

	// Test UDIV X0, X1, X2 (X0 = X1 / X2 = 0, ARM64 behavior)
	instr := createMockInstruction("UDIV", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("UDIV division by zero should not error: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(0)
	if result != expected {
		t.Errorf("UDIV division by zero result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_MADD(t *testing.T) {
	executor := NewArithmeticExecutor("MADD", "Multiply-add")
	state := NewMockState()

	// Set up initial register values
	state.SetX(1, 3)  // X1 = 3 (Rn)
	state.SetX(2, 4)  // X2 = 4 (Rm)
	state.SetX(3, 10) // X3 = 10 (Ra)

	// Test MADD X0, X1, X2, X3 (X0 = X3 + (X1 * X2) = 10 + (3 * 4) = 22)
	instr := createMockInstruction("MADD", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		createRegisterOperand(disassemble.REG_X2),
		createRegisterOperand(disassemble.REG_X3),
	})

	err := executor.Execute(state, instr)
	if err != nil {
		t.Fatalf("MADD execution failed: %v", err)
	}

	result := state.GetX(0)
	expected := uint64(22)
	if result != expected {
		t.Errorf("MADD result mismatch: got %d, expected %d", result, expected)
	}
}

func TestArithmeticExecutor_InvalidInstruction(t *testing.T) {
	executor := NewArithmeticExecutor("ADD", "Add")
	state := NewMockState()

	// Test with unsupported instruction
	instr := createMockInstruction("INVALID", []disassemble.Operand{})

	err := executor.Execute(state, instr)
	if err == nil {
		t.Fatal("Expected error for invalid instruction")
	}

	// Check that it's an emulation error
	if _, ok := err.(*core.EmulationError); !ok {
		t.Errorf("Expected EmulationError, got %T", err)
	}
}

func TestArithmeticExecutor_InsufficientOperands(t *testing.T) {
	executor := NewArithmeticExecutor("ADD", "Add")
	state := NewMockState()

	// Test ADD with insufficient operands
	instr := createMockInstruction("ADD", []disassemble.Operand{
		createRegisterOperand(disassemble.REG_X0),
		createRegisterOperand(disassemble.REG_X1),
		// Missing third operand
	})

	err := executor.Execute(state, instr)
	if err == nil {
		t.Fatal("Expected error for insufficient operands")
	}

	// Check that it's an emulation error
	if _, ok := err.(*core.EmulationError); !ok {
		t.Errorf("Expected EmulationError, got %T", err)
	}
}

// Test helper functions
func TestAddWithFlags(t *testing.T) {
	// Test normal addition
	result, carry, overflow := addWithFlags(10, 5)
	if result != 15 || carry || overflow {
		t.Errorf("addWithFlags(10, 5) = (%d, %t, %t), expected (15, false, false)", result, carry, overflow)
	}

	// Test unsigned overflow
	result, carry, overflow = addWithFlags(0xFFFFFFFFFFFFFFFF, 1)
	if result != 0 || !carry {
		t.Errorf("addWithFlags overflow test failed: result=%d, carry=%t", result, carry)
	}

	// Test signed overflow (positive + positive = negative)
	result, carry, overflow = addWithFlags(0x7FFFFFFFFFFFFFFF, 1)
	if !overflow {
		t.Error("addWithFlags should detect signed overflow")
	}
}

func TestSubWithFlags(t *testing.T) {
	// Test normal subtraction
	result, carry, overflow := subWithFlags(10, 5)
	if result != 5 || !carry || overflow {
		t.Errorf("subWithFlags(10, 5) = (%d, %t, %t), expected (5, true, false)", result, carry, overflow)
	}

	// Test borrow (no carry)
	result, carry, overflow = subWithFlags(5, 10)
	if carry {
		t.Error("subWithFlags should clear carry for borrow")
	}
}

func TestApplyShift(t *testing.T) {
	// Test LSL (Logical Shift Left)
	shiftOp := disassemble.Operand{
		ShiftValueUsed: true,
		ShiftType:      disassemble.SHIFT_TYPE_LSL,
		ShiftValue:     2,
	}
	result, err := applyShift(5, shiftOp)
	if err != nil || result != 20 {
		t.Errorf("LSL shift failed: got %d, expected 20, error: %v", result, err)
	}

	// Test LSR (Logical Shift Right)
	shiftOp.ShiftType = disassemble.SHIFT_TYPE_LSR
	result, err = applyShift(20, shiftOp)
	if err != nil || result != 5 {
		t.Errorf("LSR shift failed: got %d, expected 5, error: %v", result, err)
	}

	// Test no shift
	shiftOp.ShiftValueUsed = false
	result, err = applyShift(42, shiftOp)
	if err != nil || result != 42 {
		t.Errorf("No shift failed: got %d, expected 42, error: %v", result, err)
	}
}

func TestRegisterArithmeticInstructions(t *testing.T) {
	registry := NewRegistry()
	RegisterArithmeticInstructions(registry)

	// Test that all expected instructions are registered
	expectedInstructions := []string{
		"ADD", "ADDS", "SUB", "SUBS", "MUL", "UDIV", "SDIV",
		"MADD", "MSUB", "SMULL", "UMULL", "NEG", "NEGS",
	}

	for _, instr := range expectedInstructions {
		if !registry.HasInstruction(instr) {
			t.Errorf("Instruction %s not registered", instr)
		}
	}

	// Test that we can get executors for registered instructions
	executor, exists := registry.Get("ADD")
	if !exists {
		t.Error("Could not get ADD executor from registry")
	}

	if !executor.Supports("ADD") {
		t.Error("ADD executor does not support ADD instruction")
	}
}
