package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// MockState implements core.State for testing
type MockSystemState struct {
	registers  [31]uint64
	pc         uint64
	sp         uint64
	n, z, c, v bool
}

func NewMockSystemState() *MockSystemState {
	return &MockSystemState{
		pc: 0x1000,
		sp: 0x7fff0000,
	}
}

func (s *MockSystemState) GetX(reg int) uint64 {
	if reg < 0 || reg >= 31 {
		return 0
	}
	return s.registers[reg]
}

func (s *MockSystemState) SetX(reg int, val uint64) {
	if reg >= 0 && reg < 31 {
		s.registers[reg] = val
	}
}

func (s *MockSystemState) GetW(reg int) uint32 {
	return uint32(s.GetX(reg))
}

func (s *MockSystemState) SetW(reg int, val uint32) {
	s.SetX(reg, uint64(val))
}

func (s *MockSystemState) GetPC() uint64     { return s.pc }
func (s *MockSystemState) SetPC(addr uint64) { s.pc = addr }
func (s *MockSystemState) GetSP() uint64     { return s.sp }
func (s *MockSystemState) SetSP(addr uint64) { s.sp = addr }

func (s *MockSystemState) GetN() bool    { return s.n }
func (s *MockSystemState) SetN(val bool) { s.n = val }
func (s *MockSystemState) GetZ() bool    { return s.z }
func (s *MockSystemState) SetZ(val bool) { s.z = val }
func (s *MockSystemState) GetC() bool    { return s.c }
func (s *MockSystemState) SetC(val bool) { s.c = val }
func (s *MockSystemState) GetV() bool    { return s.v }
func (s *MockSystemState) SetV(val bool) { s.v = val }

func (s *MockSystemState) UpdateFlags(result uint64, isNegative bool) {
	s.n = isNegative
	s.z = result == 0
}

// Stub implementations for other required methods
func (s *MockSystemState) ReadMemory(addr uint64, size int) ([]byte, error) { return nil, nil }
func (s *MockSystemState) WriteMemory(addr uint64, data []byte)             {}
func (s *MockSystemState) ReadUint64(addr uint64) (uint64, error)           { return 0, nil }
func (s *MockSystemState) ReadUint32(addr uint64) (uint32, error)           { return 0, nil }
func (s *MockSystemState) WriteUint64(addr uint64, value uint64)            {}
func (s *MockSystemState) WriteUint32(addr uint64, value uint32)            {}
func (s *MockSystemState) ReadString(addr uint64) (string, error)           { return "", nil }
func (s *MockSystemState) ReadStringWithLength(addr uint64, length int) (string, error) {
	return "", nil
}
func (s *MockSystemState) ResolvePointer(addr uint64) (uint64, error) { return 0, nil }
func (s *MockSystemState) FollowPointerChain(addr uint64, maxDepth int) (uint64, error) {
	return 0, nil
}
func (s *MockSystemState) ReadStringAtPointer(ptrAddr uint64) (string, error)  { return "", nil }
func (s *MockSystemState) SetMemoryReadHandler(handler core.MemoryReadHandler) {}
func (s *MockSystemState) SetStringPoolHandler(handler core.StringPoolHandler) {}
func (s *MockSystemState) SetPointerResolver(resolver core.PointerResolver)    {}
func (s *MockSystemState) Clone() core.State                                   { return &MockSystemState{} }
func (s *MockSystemState) Reset()                                              {}

// Helper function to create a mock system instruction
func createMockSystemInstruction(operation string, operands []disassemble.Operand) *disassemble.Instruction {
	switch operation {
	case "NOP":
		return &disassemble.Instruction{Operation: disassemble.ARM64_NOP, Operands: operands}
	case "MRS":
		return &disassemble.Instruction{Operation: disassemble.ARM64_MRS, Operands: operands}
	case "MSR":
		return &disassemble.Instruction{Operation: disassemble.ARM64_MSR, Operands: operands}
	case "HLT":
		return &disassemble.Instruction{Operation: disassemble.ARM64_HLT, Operands: operands}
	case "HINT":
		return &disassemble.Instruction{Operation: disassemble.ARM64_HINT, Operands: operands}
	case "ISB":
		return &disassemble.Instruction{Operation: disassemble.ARM64_ISB, Operands: operands}
	case "DSB":
		return &disassemble.Instruction{Operation: disassemble.ARM64_DSB, Operands: operands}
	case "DMB":
		return &disassemble.Instruction{Operation: disassemble.ARM64_DMB, Operands: operands}
	case "YIELD":
		return &disassemble.Instruction{Operation: disassemble.ARM64_YIELD, Operands: operands}
	case "WFE":
		return &disassemble.Instruction{Operation: disassemble.ARM64_WFE, Operands: operands}
	case "WFI":
		return &disassemble.Instruction{Operation: disassemble.ARM64_WFI, Operands: operands}
	case "SEV":
		return &disassemble.Instruction{Operation: disassemble.ARM64_SEV, Operands: operands}
	case "SEVL":
		return &disassemble.Instruction{Operation: disassemble.ARM64_SEVL, Operands: operands}
	case "SYS":
		return &disassemble.Instruction{Operation: disassemble.ARM64_SYS, Operands: operands}
	case "SYSL":
		return &disassemble.Instruction{Operation: disassemble.ARM64_SYSL, Operands: operands}
	case "XPACLRI":
		return &disassemble.Instruction{Operation: disassemble.ARM64_XPACLRI, Operands: operands}
	case "PACIBSP":
		return &disassemble.Instruction{Operation: disassemble.ARM64_PACIBSP, Operands: operands}
	case "PACIASP":
		return &disassemble.Instruction{Operation: disassemble.ARM64_PACIASP, Operands: operands}
	case "PACIAZ":
		return &disassemble.Instruction{Operation: disassemble.ARM64_PACIAZ, Operands: operands}
	case "PACIBZ":
		return &disassemble.Instruction{Operation: disassemble.ARM64_PACIBZ, Operands: operands}
	case "AUTIBSP":
		return &disassemble.Instruction{Operation: disassemble.ARM64_AUTIBSP, Operands: operands}
	case "SB":
		return &disassemble.Instruction{Operation: disassemble.ARM64_SB, Operands: operands}
	case "ADD":
		return &disassemble.Instruction{Operation: disassemble.ARM64_ADD, Operands: operands}
	default:
		return &disassemble.Instruction{Operation: disassemble.ARM64_ERROR, Operands: operands}
	}
}

func TestSystemExecutor_NOP(t *testing.T) {
	executor := NewSystemExecutor("NOP", "No operation")
	state := NewMockSystemState()

	// Create a mock instruction for NOP
	instr := createMockSystemInstruction("NOP", nil)

	// Execute NOP
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("NOP execution failed: %v", err)
	}

	// NOP should not change any state
	if state.GetX(0) != 0 {
		t.Errorf("NOP should not modify registers")
	}
}

func TestSystemExecutor_MRS_NZCV(t *testing.T) {
	executor := NewSystemExecutor("MRS", "Move from system register")
	state := NewMockSystemState()

	// Set some flags
	state.SetN(true)
	state.SetZ(false)
	state.SetC(true)
	state.SetV(false)

	// Create a mock instruction for MRS X0, NZCV
	instr := createMockSystemInstruction("MRS", []disassemble.Operand{
		{
			Registers: []disassemble.Register{disassemble.Register(34)}, // X0
		},
		{
			// System register operand - we'll rely on string parsing
		},
	})

	// Execute MRS
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("MRS execution failed: %v", err)
	}

	// Check that X0 contains the NZCV flags
	expected := uint64((1 << 31) | (1 << 29)) // N=1, Z=0, C=1, V=0
	if state.GetX(0) != expected {
		t.Errorf("MRS NZCV failed: expected 0x%x, got 0x%x", expected, state.GetX(0))
	}
}

func TestSystemExecutor_MRS_StringBased(t *testing.T) {
	executor := NewSystemExecutor("MRS", "Move from system register")
	state := NewMockSystemState()

	// For unit testing, we test the default behavior when instruction string is empty
	// In real usage, the disassembler would provide proper instruction strings
	// This test verifies that MRS works with the fallback NZCV behavior

	// Set some flags for testing
	state.SetN(true)
	state.SetZ(false)
	state.SetC(true)
	state.SetV(false)

	// Create mock instruction for MRS (will fallback to NZCV behavior)
	instr := createMockSystemInstruction("MRS", []disassemble.Operand{
		{
			Registers: []disassemble.Register{disassemble.Register(35)}, // X1
		},
		{
			// System register operand - empty for testing
		},
	})

	// Execute MRS
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("MRS execution failed: %v", err)
	}

	// Check result - should get NZCV flags as fallback
	expected := uint64((1 << 31) | (1 << 29)) // N=1, Z=0, C=1, V=0
	if state.GetX(1) != expected {
		t.Errorf("MRS fallback failed: expected 0x%x, got 0x%x", expected, state.GetX(1))
	}
}

func TestSystemExecutor_MSR_NZCV(t *testing.T) {
	executor := NewSystemExecutor("MSR", "Move to system register")
	state := NewMockSystemState()

	// Set X0 to a value with specific flag bits
	flagValue := uint64((1 << 31) | (1 << 30)) // N=1, Z=1, C=0, V=0
	state.SetX(0, flagValue)

	// Create a mock instruction for MSR NZCV, X0
	instr := createMockSystemInstruction("MSR", []disassemble.Operand{
		{
			// System register operand - we'll rely on string parsing
		},
		{
			Registers: []disassemble.Register{disassemble.Register(34)}, // X0
		},
	})

	// Execute MSR
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("MSR execution failed: %v", err)
	}

	// Check that flags were set correctly
	if !state.GetN() {
		t.Errorf("MSR NZCV failed: N flag should be set")
	}
	if !state.GetZ() {
		t.Errorf("MSR NZCV failed: Z flag should be set")
	}
	if state.GetC() {
		t.Errorf("MSR NZCV failed: C flag should be clear")
	}
	if state.GetV() {
		t.Errorf("MSR NZCV failed: V flag should be clear")
	}
}

func TestSystemExecutor_MSR_Immediate(t *testing.T) {
	executor := NewSystemExecutor("MSR", "Move to system register")
	state := NewMockSystemState()

	// Create a mock instruction for MSR NZCV, #0xC0000000 (N=1, Z=1)
	instr := createMockSystemInstruction("MSR", []disassemble.Operand{
		{
			// System register operand - we'll rely on string parsing
		},
		{
			Immediate: 0xC0000000,
		},
	})

	// Execute MSR
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("MSR execution failed: %v", err)
	}

	// Check that flags were set correctly
	if !state.GetN() {
		t.Errorf("MSR NZCV immediate failed: N flag should be set")
	}
	if !state.GetZ() {
		t.Errorf("MSR NZCV immediate failed: Z flag should be set")
	}
	if state.GetC() {
		t.Errorf("MSR NZCV immediate failed: C flag should be clear")
	}
	if state.GetV() {
		t.Errorf("MSR NZCV immediate failed: V flag should be clear")
	}
}

func TestSystemExecutor_HLT(t *testing.T) {
	executor := NewSystemExecutor("HLT", "Halt processor")
	state := NewMockSystemState()

	// Create a mock instruction for HLT
	instr := createMockSystemInstruction("HLT", nil)

	// Execute HLT - should return an error indicating halt
	err := executor.Execute(state, instr)
	if err == nil {
		t.Errorf("HLT should return an error to indicate halt condition")
	}

	// Check that it's the expected halt error
	if emulErr, ok := err.(*core.EmulationError); ok {
		if emulErr.Err != core.ErrExecutionLimit {
			t.Errorf("HLT should return ErrExecutionLimit, got %v", emulErr.Err)
		}
	} else {
		t.Errorf("HLT should return EmulationError, got %T", err)
	}
}

func TestSystemExecutor_MemoryBarriers(t *testing.T) {
	testCases := []struct {
		name        string
		mnemonic    string
		description string
	}{
		{"ISB", "ISB", "Instruction synchronization barrier"},
		{"DSB", "DSB", "Data synchronization barrier"},
		{"DMB", "DMB", "Data memory barrier"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor := NewSystemExecutor(tc.mnemonic, tc.description)
			state := NewMockSystemState()

			// Create a mock instruction
			instr := createMockSystemInstruction(tc.mnemonic, nil)

			// Execute barrier instruction - should be NOP
			err := executor.Execute(state, instr)
			if err != nil {
				t.Errorf("%s execution failed: %v", tc.mnemonic, err)
			}

			// Barriers should not change any state
			if state.GetX(0) != 0 {
				t.Errorf("%s should not modify registers", tc.mnemonic)
			}
		})
	}
}

func TestSystemExecutor_PowerManagement(t *testing.T) {
	testCases := []struct {
		name        string
		mnemonic    string
		description string
	}{
		{"YIELD", "YIELD", "Yield hint"},
		{"WFE", "WFE", "Wait for event"},
		{"WFI", "WFI", "Wait for interrupt"},
		{"SEV", "SEV", "Send event"},
		{"SEVL", "SEVL", "Send event local"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor := NewSystemExecutor(tc.mnemonic, tc.description)
			state := NewMockSystemState()

			// Create a mock instruction
			instr := createMockSystemInstruction(tc.mnemonic, nil)

			// Execute power management instruction - should be NOP
			err := executor.Execute(state, instr)
			if err != nil {
				t.Errorf("%s execution failed: %v", tc.mnemonic, err)
			}

			// Power management instructions should not change any state
			if state.GetX(0) != 0 {
				t.Errorf("%s should not modify registers", tc.mnemonic)
			}
		})
	}
}

func TestSystemExecutor_SYS_SYSL(t *testing.T) {
	testCases := []struct {
		name        string
		mnemonic    string
		description string
		hasResult   bool
	}{
		{"SYS", "SYS", "System instruction", false},
		{"SYSL", "SYSL", "System instruction with result", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor := NewSystemExecutor(tc.mnemonic, tc.description)
			state := NewMockSystemState()

			var operands []disassemble.Operand
			if tc.hasResult {
				// SYSL has a destination register
				operands = []disassemble.Operand{
					{
						Registers: []disassemble.Register{disassemble.Register(34)}, // X0
					},
				}
			}

			// Create a mock instruction
			instr := createMockSystemInstruction(tc.mnemonic, operands)

			// Execute system instruction
			err := executor.Execute(state, instr)
			if err != nil {
				t.Errorf("%s execution failed: %v", tc.mnemonic, err)
			}

			// SYSL should set destination register to 0
			if tc.hasResult && state.GetX(0) != 0 {
				t.Errorf("SYSL should set destination register to 0, got %d", state.GetX(0))
			}
		})
	}
}

func TestSystemExecutor_HINT(t *testing.T) {
	executor := NewSystemExecutor("HINT", "Hint instruction")
	state := NewMockSystemState()

	// Create a mock instruction for HINT
	instr := createMockSystemInstruction("HINT", nil)

	// Execute HINT - should be NOP
	err := executor.Execute(state, instr)
	if err != nil {
		t.Errorf("HINT execution failed: %v", err)
	}

	// HINT should not change any state
	if state.GetX(0) != 0 {
		t.Errorf("HINT should not modify registers")
	}
}

func TestSystemExecutor_InvalidInstruction(t *testing.T) {
	executor := NewSystemExecutor("NOP", "No operation")
	state := NewMockSystemState()

	// Test with nil instruction
	err := executor.Execute(state, nil)
	if err == nil {
		t.Errorf("Execute should fail with nil instruction")
	}
}

func TestSystemExecutor_XPACLRI_SB(t *testing.T) {
	testCases := []struct {
		name        string
		mnemonic    string
		description string
	}{
		{"XPACLRI", "XPACLRI", "Strip Pointer Authentication Code from Link Register"},
		{"SB", "SB", "Speculation barrier"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			executor := NewSystemExecutor(tc.mnemonic, tc.description)
			state := NewMockSystemState()

			// Create a mock instruction
			instr := createMockSystemInstruction(tc.mnemonic, nil)

			// Execute instruction - should be NOP
			err := executor.Execute(state, instr)
			if err != nil {
				t.Errorf("%s execution failed: %v", tc.mnemonic, err)
			}

			// Instructions should not change any state
			if state.GetX(0) != 0 {
				t.Errorf("%s should not modify registers", tc.mnemonic)
			}
		})
	}
}

func TestSystemExecutor_AUTIBSP(t *testing.T) {
	state := NewMockSystemState()
	state.SetSP(0x7fff1234000)
	baseLR := uint64(0x0000000100002222)
	state.SetX(30, baseLR)

	pacExecutor := NewSystemExecutor("PACIBSP", "Pointer authenticate LR using SP")
	if err := pacExecutor.Execute(state, createMockSystemInstruction("PACIBSP", nil)); err != nil {
		t.Fatalf("PACIBSP execution failed: %v", err)
	}

	signedLR := state.GetX(30)
	if signedLR>>56 == 0 {
		t.Fatalf("PACIBSP should set a pseudo PAC in the top byte, got 0x%x", signedLR)
	}

	autExecutor := NewSystemExecutor("AUTIBSP", "Authenticate LR using SP (B-key)")
	if err := autExecutor.Execute(state, createMockSystemInstruction("AUTIBSP", nil)); err != nil {
		t.Fatalf("AUTIBSP execution failed: %v", err)
	}

	if got := state.GetX(30); got != baseLR {
		t.Errorf("AUTIBSP should restore LR to base value, want 0x%x got 0x%x", baseLR, got)
	}

	state.SetX(30, uint64(0xAA)<<56|0x0000000000003333)
	if err := autExecutor.Execute(state, createMockSystemInstruction("AUTIBSP", nil)); err != nil {
		t.Fatalf("AUTIBSP execution failed on invalid PAC: %v", err)
	}

	if got := state.GetX(30); got != 0 {
		t.Errorf("AUTIBSP should poison LR on invalid PAC, want 0 got 0x%x", got)
	}
}

func TestSystemExecutor_UnsupportedInstruction(t *testing.T) {
	executor := NewSystemExecutor("NOP", "No operation")
	state := NewMockSystemState()

	// Create an instruction that doesn't match the executor
	instr := createMockSystemInstruction("ADD", nil)

	// Should fail validation
	err := executor.Execute(state, instr)
	if err == nil {
		t.Errorf("Execute should fail with unsupported instruction")
	}
}

func TestRegisterSystemInstructions(t *testing.T) {
	registry := NewRegistry()

	// Register system instructions
	RegisterSystemInstructions(registry)

	// Check that all expected instructions are registered
	expectedInstructions := []string{
		"NOP", "MRS", "MSR", "SYS", "SYSL",
		"ISB", "DSB", "DMB",
		"HINT", "YIELD", "WFE", "WFI", "SEV", "SEVL",
		"HLT", "XPACLRI", "PACIBSP", "PACIASP", "PACIAZ", "PACIBZ", "AUTIBSP", "SB",
	}

	for _, instr := range expectedInstructions {
		if !registry.HasInstruction(instr) {
			t.Errorf("Instruction %s not registered", instr)
		}

		executor, exists := registry.Get(instr)
		if !exists {
			t.Errorf("Could not retrieve executor for %s", instr)
		}

		if !executor.Supports(instr) {
			t.Errorf("Executor for %s does not support the instruction", instr)
		}
	}

	// Check total count
	if registry.Count() < len(expectedInstructions) {
		t.Errorf("Expected at least %d instructions, got %d", len(expectedInstructions), registry.Count())
	}
}
