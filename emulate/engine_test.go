package emulate

import (
	"encoding/binary"
	"testing"

	"github.com/blacktop/arm64-cgo/emulate/state"
)

func TestNewEngine(t *testing.T) {
	engine := NewEngine()
	if engine == nil {
		t.Fatal("NewEngine() returned nil")
	}

	// Check default configuration
	if engine.maxInstructions != 10000 {
		t.Errorf("Expected maxInstructions to be 10000, got %d", engine.maxInstructions)
	}

	if engine.enableTrace {
		t.Error("Expected tracing to be disabled by default")
	}

	if !engine.StopOnError {
		t.Error("Expected StopOnError to be true by default")
	}

	if engine.GetInstructionCount() != 0 {
		t.Errorf("Expected instruction count to be 0, got %d", engine.GetInstructionCount())
	}
}

func TestNewEngineWithConfig(t *testing.T) {
	config := &EngineConfig{
		MaxInstructions: 5000,
		EnableTrace:     true,
		StopOnError:     false,
		InitialPC:       0x1000,
		InitialSP:       0x2000,
	}

	engine := NewEngineWithConfig(config)
	if engine == nil {
		t.Fatal("NewEngineWithConfig() returned nil")
	}

	if engine.maxInstructions != 5000 {
		t.Errorf("Expected maxInstructions to be 5000, got %d", engine.maxInstructions)
	}

	if !engine.enableTrace {
		t.Error("Expected tracing to be enabled")
	}

	if engine.StopOnError {
		t.Error("Expected StopOnError to be false")
	}

	if engine.GetPC() != 0x1000 {
		t.Errorf("Expected PC to be 0x1000, got 0x%x", engine.GetPC())
	}

	if engine.GetSP() != 0x2000 {
		t.Errorf("Expected SP to be 0x2000, got 0x%x", engine.GetSP())
	}
}

func TestNewEngineWithState(t *testing.T) {
	customState := state.NewState()
	customState.SetX(0, 0x12345678)
	customState.SetPC(0x4000)

	engine := NewEngineWithState(customState)
	if engine == nil {
		t.Fatal("NewEngineWithState() returned nil")
	}

	if engine.GetRegister(0) != 0x12345678 {
		t.Errorf("Expected X0 to be 0x12345678, got 0x%x", engine.GetRegister(0))
	}

	if engine.GetPC() != 0x4000 {
		t.Errorf("Expected PC to be 0x4000, got 0x%x", engine.GetPC())
	}
}

func TestDefaultEngineConfig(t *testing.T) {
	config := DefaultEngineConfig()
	if config == nil {
		t.Fatal("DefaultEngineConfig() returned nil")
	}

	if config.MaxInstructions != 10000 {
		t.Errorf("Expected MaxInstructions to be 10000, got %d", config.MaxInstructions)
	}

	if config.EnableTrace {
		t.Error("Expected EnableTrace to be false")
	}

	if !config.StopOnError {
		t.Error("Expected StopOnError to be true")
	}

	if config.InitialSP != 0x7fff00000000 {
		t.Errorf("Expected InitialSP to be 0x7fff00000000, got 0x%x", config.InitialSP)
	}

	if config.StackSize != 1024*1024 {
		t.Errorf("Expected StackSize to be 1MB, got %d", config.StackSize)
	}
}

func TestLoadInstructions(t *testing.T) {
	engine := NewEngine()

	// Test valid aligned instructions (NOP instructions: 0xd503201f)
	instructions := []byte{
		0x1f, 0x20, 0x03, 0xd5, // NOP
		0x1f, 0x20, 0x03, 0xd5, // NOP
		0x1f, 0x20, 0x03, 0xd5, // NOP
	}

	err := engine.LoadInstructions(0x1000, instructions)
	if err != nil {
		t.Fatalf("LoadInstructions failed: %v", err)
	}

	// Verify instructions were loaded
	data, err := engine.GetMemory(0x1000, len(instructions))
	if err != nil {
		t.Fatalf("Failed to read loaded instructions: %v", err)
	}

	for i, b := range instructions {
		if data[i] != b {
			t.Errorf("Instruction byte %d mismatch: expected 0x%02x, got 0x%02x", i, b, data[i])
		}
	}
}

func TestLoadInstructionsAlignment(t *testing.T) {
	engine := NewEngine()

	// Test unaligned address
	instructions := []byte{0x1f, 0x20, 0x03, 0xd5}
	err := engine.LoadInstructions(0x1001, instructions) // Unaligned address
	if err == nil {
		t.Error("Expected error for unaligned address")
	}

	// Test unaligned data length
	instructions = []byte{0x1f, 0x20, 0x03} // 3 bytes, not 4-byte aligned
	err = engine.LoadInstructions(0x1000, instructions)
	if err == nil {
		t.Error("Expected error for unaligned data length")
	}
}

func TestLoadInstructionsAt(t *testing.T) {
	engine := NewEngine()

	instructions := []byte{
		0x1f, 0x20, 0x03, 0xd5, // NOP
	}

	err := engine.LoadInstructionsAt(0x2000, instructions)
	if err != nil {
		t.Fatalf("LoadInstructionsAt failed: %v", err)
	}

	// Verify PC was set
	if engine.GetPC() != 0x2000 {
		t.Errorf("Expected PC to be 0x2000, got 0x%x", engine.GetPC())
	}

	// Verify instructions were loaded
	data, err := engine.GetMemory(0x2000, len(instructions))
	if err != nil {
		t.Fatalf("Failed to read loaded instructions: %v", err)
	}

	for i, b := range instructions {
		if data[i] != b {
			t.Errorf("Instruction byte %d mismatch: expected 0x%02x, got 0x%02x", i, b, data[i])
		}
	}
}

func TestRegisterOperations(t *testing.T) {
	engine := NewEngine()

	// Test setting and getting registers
	for i := range 31 {
		value := uint64(0x1000 + i)
		engine.SetRegister(i, value)

		if engine.GetRegister(i) != value {
			t.Errorf("Register X%d: expected 0x%x, got 0x%x", i, value, engine.GetRegister(i))
		}
	}
}

func TestPCOperations(t *testing.T) {
	engine := NewEngine()

	testValues := []uint64{0x1000, 0x2000, 0x100000000, 0xffffffff00000000}

	for _, value := range testValues {
		engine.SetPC(value)
		if engine.GetPC() != value {
			t.Errorf("PC: expected 0x%x, got 0x%x", value, engine.GetPC())
		}
	}
}

func TestSPOperations(t *testing.T) {
	engine := NewEngine()

	testValues := []uint64{0x1000, 0x2000, 0x7fff00000000, 0x8000000000000000}

	for _, value := range testValues {
		engine.SetSP(value)
		if engine.GetSP() != value {
			t.Errorf("SP: expected 0x%x, got 0x%x", value, engine.GetSP())
		}
	}
}

func TestMemoryOperations(t *testing.T) {
	engine := NewEngine()

	// Test byte array operations
	testData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	addr := uint64(0x10000)

	err := engine.SetMemory(addr, testData)
	if err != nil {
		t.Fatalf("SetMemory failed: %v", err)
	}

	readData, err := engine.GetMemory(addr, len(testData))
	if err != nil {
		t.Fatalf("GetMemory failed: %v", err)
	}

	for i, b := range testData {
		if readData[i] != b {
			t.Errorf("Memory byte %d: expected 0x%02x, got 0x%02x", i, b, readData[i])
		}
	}

	// Test 32-bit operations
	value32 := uint32(0x12345678)
	err = engine.SetMemoryUint32(addr, value32)
	if err != nil {
		t.Fatalf("SetMemoryUint32 failed: %v", err)
	}

	read32, err := engine.GetMemoryUint32(addr)
	if err != nil {
		t.Fatalf("GetMemoryUint32 failed: %v", err)
	}

	if read32 != value32 {
		t.Errorf("Memory uint32: expected 0x%x, got 0x%x", value32, read32)
	}

	// Test 64-bit operations
	value64 := uint64(0x123456789abcdef0)
	err = engine.SetMemoryUint64(addr, value64)
	if err != nil {
		t.Fatalf("SetMemoryUint64 failed: %v", err)
	}

	read64, err := engine.GetMemoryUint64(addr)
	if err != nil {
		t.Fatalf("GetMemoryUint64 failed: %v", err)
	}

	if read64 != value64 {
		t.Errorf("Memory uint64: expected 0x%x, got 0x%x", value64, read64)
	}
}

func TestTracing(t *testing.T) {
	engine := NewEngine()
	engine.SetTrace(true)

	// Load a NOP instruction
	nopInstr := []byte{0x1f, 0x20, 0x03, 0xd5}
	err := engine.LoadInstructionsAt(0x1000, nopInstr)
	if err != nil {
		t.Fatalf("LoadInstructionsAt failed: %v", err)
	}

	// Execute the instruction
	err = engine.StepOver()
	if err != nil {
		t.Fatalf("StepOver failed: %v", err)
	}

	// Check trace
	trace := engine.GetTrace()
	if len(trace) != 1 {
		t.Errorf("Expected 1 trace entry, got %d", len(trace))
	}

	if len(trace) > 0 {
		entry := trace[0]
		if entry.Address != 0x1000 {
			t.Errorf("Expected trace address 0x1000, got 0x%x", entry.Address)
		}

		expectedValue := binary.LittleEndian.Uint32(nopInstr)
		if entry.Value != expectedValue {
			t.Errorf("Expected trace value 0x%x, got 0x%x", expectedValue, entry.Value)
		}
	}

	// Test disabling trace
	engine.SetTrace(false)

	// Load another NOP instruction for the second step
	err = engine.LoadInstructions(0x1004, nopInstr)
	if err != nil {
		t.Fatalf("Failed to load second instruction: %v", err)
	}

	err = engine.StepOver()
	if err != nil {
		t.Fatalf("Second StepOver failed: %v", err)
	}

	// Trace should still be 1 (not growing)
	trace = engine.GetTrace()
	if len(trace) != 0 { // Trace gets cleared when disabled
		t.Errorf("Expected trace to be cleared when disabled, got %d entries", len(trace))
	}
}

func TestMaxInstructions(t *testing.T) {
	engine := NewEngine()
	engine.SetMaxInstructions(2)

	// Load multiple NOP instructions
	instructions := []byte{
		0x1f, 0x20, 0x03, 0xd5, // NOP
		0x1f, 0x20, 0x03, 0xd5, // NOP
		0x1f, 0x20, 0x03, 0xd5, // NOP
		0x1f, 0x20, 0x03, 0xd5, // NOP
	}

	err := engine.LoadInstructionsAt(0x1000, instructions)
	if err != nil {
		t.Fatalf("LoadInstructionsAt failed: %v", err)
	}

	// Run should stop after max instructions
	err = engine.Run()
	// The error might be about hitting the limit OR about running out of instructions
	// Both are acceptable for this test - we just want to ensure it stops
	if err == nil && engine.GetInstructionCount() < 2 {
		t.Error("Expected engine to execute at least some instructions or return an error")
	}

	if engine.GetInstructionCount() != 2 {
		t.Errorf("Expected instruction count to be 2, got %d", engine.GetInstructionCount())
	}
}

func TestReset(t *testing.T) {
	engine := NewEngine()

	// Set some state
	engine.SetRegister(0, 0x12345678)
	engine.SetPC(0x2000)
	engine.SetTrace(true)

	// Execute an instruction to increment count
	nopInstr := []byte{0x1f, 0x20, 0x03, 0xd5}
	err := engine.LoadInstructionsAt(0x1000, nopInstr)
	if err != nil {
		t.Fatalf("LoadInstructionsAt failed: %v", err)
	}

	err = engine.StepOver()
	if err != nil {
		t.Fatalf("StepOver failed: %v", err)
	}

	// Reset
	engine.Reset()

	// Check that state was reset
	if engine.GetRegister(0) != 0 {
		t.Errorf("Expected X0 to be 0 after reset, got 0x%x", engine.GetRegister(0))
	}

	if engine.GetInstructionCount() != 0 {
		t.Errorf("Expected instruction count to be 0 after reset, got %d", engine.GetInstructionCount())
	}

	if len(engine.GetTrace()) != 0 {
		t.Errorf("Expected trace to be empty after reset, got %d entries", len(engine.GetTrace()))
	}

	if engine.LastError != nil {
		t.Error("Expected LastError to be nil after reset")
	}
}

func TestConfigure(t *testing.T) {
	engine := NewEngine()

	config := &EngineConfig{
		MaxInstructions: 5000,
		EnableTrace:     true,
		StopOnError:     false,
		InitialPC:       0x3000,
		InitialSP:       0x4000,
	}

	engine.Configure(config)

	if engine.maxInstructions != 5000 {
		t.Errorf("Expected maxInstructions to be 5000, got %d", engine.maxInstructions)
	}

	if !engine.enableTrace {
		t.Error("Expected tracing to be enabled")
	}

	if engine.StopOnError {
		t.Error("Expected StopOnError to be false")
	}

	if engine.GetPC() != 0x3000 {
		t.Errorf("Expected PC to be 0x3000, got 0x%x", engine.GetPC())
	}

	if engine.GetSP() != 0x4000 {
		t.Errorf("Expected SP to be 0x4000, got 0x%x", engine.GetSP())
	}
}

func TestRegistryOperations(t *testing.T) {
	engine := NewEngine()

	registry := engine.GetRegistry()
	if registry == nil {
		t.Fatal("GetRegistry() returned nil")
	}

	instructions := engine.ListSupportedInstructions()
	if len(instructions) == 0 {
		t.Error("Expected some supported instructions")
	}

	// Check that common instructions are supported
	found := false
	for _, instr := range instructions {
		if instr == "NOP" || instr == "nop" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected 'NOP' instruction to be supported. Available instructions: %v", instructions)
	}
}

func TestExecuteInstructionIntegration(t *testing.T) {
	engine := NewEngine()

	// Test executing a MOV instruction: mov x0, #42 (0xd2800540)
	movInstr := uint32(0xd2800540)

	err := engine.ExecuteInstruction(0x1000, movInstr)
	if err != nil {
		t.Fatalf("ExecuteInstruction failed: %v", err)
	}

	// Check that X0 was set to 42
	if engine.GetRegister(0) != 42 {
		t.Errorf("Expected X0 to be 42, got %d", engine.GetRegister(0))
	}

	// Check that PC was advanced
	if engine.GetPC() != 0x1004 {
		t.Errorf("Expected PC to be 0x1004, got 0x%x", engine.GetPC())
	}

	// Check instruction count
	if engine.GetInstructionCount() != 1 {
		t.Errorf("Expected instruction count to be 1, got %d", engine.GetInstructionCount())
	}
}

func TestStepOverIntegration(t *testing.T) {
	engine := NewEngine()

	// Load a MOV instruction: mov x1, #100 (0xd2800c81)
	movInstr := []byte{0x81, 0x0c, 0x80, 0xd2}
	err := engine.LoadInstructionsAt(0x2000, movInstr)
	if err != nil {
		t.Fatalf("LoadInstructionsAt failed: %v", err)
	}

	err = engine.StepOver()
	if err != nil {
		t.Fatalf("StepOver failed: %v", err)
	}

	// Check that X1 was set to 100
	if engine.GetRegister(1) != 100 {
		t.Errorf("Expected X1 to be 100, got %d", engine.GetRegister(1))
	}

	// Check that PC was advanced
	if engine.GetPC() != 0x2004 {
		t.Errorf("Expected PC to be 0x2004, got 0x%x", engine.GetPC())
	}
}

func TestRunAtIntegration(t *testing.T) {
	engine := NewEngine()
	engine.SetMaxInstructions(3)

	// Load multiple instructions
	instructions := []byte{
		0x40, 0x00, 0x80, 0xd2, // mov x0, #2
		0x61, 0x00, 0x80, 0xd2, // mov x1, #3
		0x02, 0x00, 0x01, 0x8b, // add x2, x0, x1
	}

	err := engine.LoadInstructions(0x3000, instructions)
	if err != nil {
		t.Fatalf("LoadInstructions failed: %v", err)
	}

	err = engine.RunAt(0x3000)
	if err != nil {
		t.Fatalf("RunAt failed: %v", err)
	}

	// Check results
	if engine.GetRegister(0) != 2 {
		t.Errorf("Expected X0 to be 2, got %d", engine.GetRegister(0))
	}

	if engine.GetRegister(1) != 3 {
		t.Errorf("Expected X1 to be 3, got %d", engine.GetRegister(1))
	}

	if engine.GetRegister(2) != 5 {
		t.Errorf("Expected X2 to be 5, got %d", engine.GetRegister(2))
	}

	if engine.GetInstructionCount() != 3 {
		t.Errorf("Expected instruction count to be 3, got %d", engine.GetInstructionCount())
	}
}

func TestCustomMemoryHandlers(t *testing.T) {
	// Test with custom memory handler
	readHandlerCalled := false
	customReadHandler := func(addr uint64, size int) ([]byte, error) {
		readHandlerCalled = true
		// Return some test data
		data := make([]byte, size)
		for i := range data {
			data[i] = byte(addr + uint64(i))
		}
		return data, nil
	}

	config := DefaultEngineConfig()
	config.MemoryHandler = customReadHandler

	engine := NewEngineWithConfig(config)

	// Try to read from an unmapped address - should call our handler
	data, err := engine.GetMemory(0x50000000, 4)
	if err != nil {
		t.Fatalf("GetMemory with custom handler failed: %v", err)
	}

	if !readHandlerCalled {
		t.Error("Expected custom read handler to be called")
	}

	if len(data) != 4 {
		t.Errorf("Expected 4 bytes, got %d", len(data))
	}

	// Verify the data matches what our handler should return
	expectedData := []byte{0x00, 0x01, 0x02, 0x03}
	for i, b := range expectedData {
		if data[i] != b {
			t.Errorf("Data byte %d: expected 0x%02x, got 0x%02x", i, b, data[i])
		}
	}
}

func TestEngineStateManagement(t *testing.T) {
	engine := NewEngine()

	// Get initial state
	state1 := engine.GetState()
	if state1 == nil {
		t.Fatal("GetState() returned nil")
	}

	// Modify state
	state1.SetX(5, 0x12345678)
	state1.SetPC(0x8000)

	// Create new state
	state2 := state.NewState()
	state2.SetX(10, 0x87654321)
	state2.SetPC(0x9000)

	// Set new state
	engine.SetState(state2)

	// Verify state changed
	if engine.GetRegister(5) != 0 { // Should be reset
		t.Errorf("Expected X5 to be 0 after state change, got 0x%x", engine.GetRegister(5))
	}

	if engine.GetRegister(10) != 0x87654321 {
		t.Errorf("Expected X10 to be 0x87654321, got 0x%x", engine.GetRegister(10))
	}

	if engine.GetPC() != 0x9000 {
		t.Errorf("Expected PC to be 0x9000, got 0x%x", engine.GetPC())
	}
}
