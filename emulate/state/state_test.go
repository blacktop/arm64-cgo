package state

import (
	"bytes"
	"encoding/binary"
	"strings"
	"testing"

	"github.com/blacktop/arm64-cgo/emulate/core"
)

func TestNewState(t *testing.T) {
	state := NewState()

	// Check default values
	if state.GetPC() != 0 {
		t.Errorf("Expected PC to be 0, got %x", state.GetPC())
	}

	if state.GetSP() != 0x7ffff0000000 {
		t.Errorf("Expected SP to be 0x7ffff0000000, got %x", state.GetSP())
	}

	// Check all registers are zero
	for i := range 31 {
		if state.GetX(i) != 0 {
			t.Errorf("Expected X%d to be 0, got %x", i, state.GetX(i))
		}
		if state.GetW(i) != 0 {
			t.Errorf("Expected W%d to be 0, got %x", i, state.GetW(i))
		}
	}

	// Check flags are false
	if state.GetN() || state.GetZ() || state.GetC() || state.GetV() {
		t.Error("Expected all flags to be false initially")
	}
}

func TestNewStateWithConfig(t *testing.T) {
	config := core.StateConfig{
		StackSize:   0x200000, // 2MB
		StackBase:   0x8000000000000000,
		InitialPC:   0x1000,
		InitialSP:   0x8000000000000000,
		EnableDebug: true,
	}

	state := NewStateWithConfig(config)

	if state.GetPC() != 0x1000 {
		t.Errorf("Expected PC to be 0x1000, got %x", state.GetPC())
	}

	if state.GetSP() != 0x8000000000000000 {
		t.Errorf("Expected SP to be 0x8000000000000000, got %x", state.GetSP())
	}

	if len(state.stack) != 0x200000 {
		t.Errorf("Expected stack size to be 0x200000, got %x", len(state.stack))
	}
}

func TestRegisterAccess(t *testing.T) {
	state := NewState()

	// Test X register access
	state.SetX(5, 0x123456789abcdef0)
	if state.GetX(5) != 0x123456789abcdef0 {
		t.Errorf("Expected X5 to be 0x123456789abcdef0, got %x", state.GetX(5))
	}

	// Test W register access (should clear upper 32 bits)
	state.SetW(10, 0x12345678)
	if state.GetW(10) != 0x12345678 {
		t.Errorf("Expected W10 to be 0x12345678, got %x", state.GetW(10))
	}
	if state.GetX(10) != 0x12345678 {
		t.Errorf("Expected X10 to be 0x12345678 (upper bits cleared), got %x", state.GetX(10))
	}

	// Test register validity tracking
	if !state.IsValid(5) {
		t.Error("Expected X5 to be marked as valid")
	}
	if !state.IsWide(5) {
		t.Error("Expected X5 to be marked as wide")
	}
	if !state.IsValid(10) {
		t.Error("Expected W10 to be marked as valid")
	}
	if state.IsWide(10) {
		t.Error("Expected W10 to be marked as not wide")
	}

	// Test invalid register indices
	if state.GetX(-1) != 0 || state.GetX(31) != 0 {
		t.Error("Invalid register indices should return 0")
	}
}

func TestSpecialRegisters(t *testing.T) {
	state := NewState()

	// Test PC
	state.SetPC(0x100000000)
	if state.GetPC() != 0x100000000 {
		t.Errorf("Expected PC to be 0x100000000, got %x", state.GetPC())
	}

	// Test SP
	state.SetSP(0x7fff00000000)
	if state.GetSP() != 0x7fff00000000 {
		t.Errorf("Expected SP to be 0x7fff00000000, got %x", state.GetSP())
	}
}

func TestFlags(t *testing.T) {
	state := NewState()

	// Test individual flag setting
	state.SetN(true)
	if !state.GetN() {
		t.Error("Expected N flag to be true")
	}

	state.SetZ(true)
	if !state.GetZ() {
		t.Error("Expected Z flag to be true")
	}

	state.SetC(true)
	if !state.GetC() {
		t.Error("Expected C flag to be true")
	}

	state.SetV(true)
	if !state.GetV() {
		t.Error("Expected V flag to be true")
	}

	// Test flag clearing
	state.SetN(false)
	if state.GetN() {
		t.Error("Expected N flag to be false")
	}

	// Test UpdateFlags
	state.UpdateFlags(0, false)
	if !state.GetZ() || state.GetN() {
		t.Error("UpdateFlags(0, false) should set Z=true, N=false")
	}

	state.UpdateFlags(0x8000000000000000, true)
	if state.GetZ() || !state.GetN() {
		t.Error("UpdateFlags(0x8000000000000000, true) should set Z=false, N=true")
	}
}

func TestMemoryOperations(t *testing.T) {
	state := NewState()

	// Test basic memory write/read
	testData := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	addr := uint64(0x1000)

	state.WriteMemory(addr, testData)
	readData, err := state.ReadMemory(addr, len(testData))
	if err != nil {
		t.Fatalf("Failed to read memory: %v", err)
	}

	if !bytes.Equal(testData, readData) {
		t.Errorf("Memory read/write mismatch. Expected %v, got %v", testData, readData)
	}

	// Test uint64 operations
	value64 := uint64(0x123456789abcdef0)
	state.WriteUint64(addr+16, value64)
	readValue64, err := state.ReadUint64(addr + 16)
	if err != nil {
		t.Fatalf("Failed to read uint64: %v", err)
	}
	if readValue64 != value64 {
		t.Errorf("Expected uint64 %x, got %x", value64, readValue64)
	}

	// Test uint32 operations
	value32 := uint32(0x12345678)
	state.WriteUint32(addr+24, value32)
	readValue32, err := state.ReadUint32(addr + 24)
	if err != nil {
		t.Fatalf("Failed to read uint32: %v", err)
	}
	if readValue32 != value32 {
		t.Errorf("Expected uint32 %x, got %x", value32, readValue32)
	}
}

func TestStackMemory(t *testing.T) {
	state := NewState()

	// Test stack access
	stackAddr := state.GetSP()
	testData := []byte{0xaa, 0xbb, 0xcc, 0xdd}

	state.WriteMemory(stackAddr, testData)
	readData, err := state.ReadMemory(stackAddr, len(testData))
	if err != nil {
		t.Fatalf("Failed to read from stack: %v", err)
	}

	if !bytes.Equal(testData, readData) {
		t.Errorf("Stack read/write mismatch. Expected %v, got %v", testData, readData)
	}

	// Test stack bounds checking
	_, err = state.ReadMemory(stackAddr+uint64(len(state.stack)), 1)
	if err == nil {
		t.Error("Expected error when reading beyond stack bounds")
	}
}

func TestStringOperations(t *testing.T) {
	state := NewState()

	// Test null-terminated string
	testStr := "Hello, World!"
	strData := append([]byte(testStr), 0) // Add null terminator
	addr := uint64(0x2000)

	state.WriteMemory(addr, strData)
	readStr, err := state.ReadString(addr)
	if err != nil {
		t.Fatalf("Failed to read string: %v", err)
	}

	if readStr != testStr {
		t.Errorf("Expected string '%s', got '%s'", testStr, readStr)
	}

	// Test string with length
	readStrLen, err := state.ReadStringWithLength(addr, len(testStr)+5)
	if err != nil {
		t.Fatalf("Failed to read string with length: %v", err)
	}

	if readStrLen != testStr {
		t.Errorf("Expected string '%s', got '%s'", testStr, readStrLen)
	}
}

func TestPointerOperations(t *testing.T) {
	state := NewState()

	// Test basic pointer resolution
	targetAddr := uint64(0x3000)
	ptrAddr := uint64(0x2000)

	// Write pointer value
	state.WriteUint64(ptrAddr, targetAddr)

	resolvedAddr, err := state.ResolvePointer(ptrAddr)
	if err != nil {
		t.Fatalf("Failed to resolve pointer: %v", err)
	}

	if resolvedAddr != targetAddr {
		t.Errorf("Expected resolved address %x, got %x", targetAddr, resolvedAddr)
	}

	// Test pointer chain following
	chain1 := uint64(0x4000)
	chain2 := uint64(0x5000)
	finalAddr := uint64(0x6000)

	state.WriteUint64(ptrAddr, chain1)
	state.WriteUint64(chain1, chain2)
	state.WriteUint64(chain2, finalAddr)

	chainResult, err := state.FollowPointerChain(ptrAddr, 5)
	if err != nil {
		t.Fatalf("Failed to follow pointer chain: %v", err)
	}

	if chainResult != finalAddr {
		t.Errorf("Expected final address %x, got %x", finalAddr, chainResult)
	}

	// Test string at pointer
	testStr := "Pointer String"
	strData := append([]byte(testStr), 0)
	strAddr := uint64(0x7000)

	state.WriteMemory(strAddr, strData)
	state.WriteUint64(ptrAddr, strAddr)

	ptrStr, err := state.ReadStringAtPointer(ptrAddr)
	if err != nil {
		t.Fatalf("Failed to read string at pointer: %v", err)
	}

	if ptrStr != testStr {
		t.Errorf("Expected string '%s', got '%s'", testStr, ptrStr)
	}
}

func TestCustomHandlers(t *testing.T) {
	state := NewState()

	// Test custom memory handler
	customData := []byte{0xde, 0xad, 0xbe, 0xef}
	customAddr := uint64(0x10000000)

	state.SetMemoryReadHandler(func(addr uint64, size int) ([]byte, error) {
		if addr == customAddr && size == len(customData) {
			return customData, nil
		}
		return nil, core.ErrMemoryAccess
	})

	readData, err := state.ReadMemory(customAddr, len(customData))
	if err != nil {
		t.Fatalf("Failed to read via custom handler: %v", err)
	}

	if !bytes.Equal(customData, readData) {
		t.Errorf("Custom handler data mismatch. Expected %v, got %v", customData, readData)
	}

	// Test custom string handler
	customStr := "Custom String"
	state.SetStringPoolHandler(func(addr uint64) (string, error) {
		if addr == customAddr {
			return customStr, nil
		}
		return "", core.ErrMemoryAccess
	})

	readStr, err := state.ReadString(customAddr)
	if err != nil {
		t.Fatalf("Failed to read via custom string handler: %v", err)
	}

	if readStr != customStr {
		t.Errorf("Expected custom string '%s', got '%s'", customStr, readStr)
	}

	// Test custom pointer resolver
	customTarget := uint64(0x20000000)
	state.SetPointerResolver(func(addr uint64) (uint64, error) {
		if addr == customAddr {
			return customTarget, nil
		}
		return 0, core.ErrMemoryAccess
	})

	resolvedAddr, err := state.ResolvePointer(customAddr)
	if err != nil {
		t.Fatalf("Failed to resolve via custom resolver: %v", err)
	}

	if resolvedAddr != customTarget {
		t.Errorf("Expected resolved address %x, got %x", customTarget, resolvedAddr)
	}
}

func TestStateManagement(t *testing.T) {
	state := NewState()

	// Set up some state
	state.SetX(5, 0x123456789abcdef0)
	state.SetPC(0x100000000)
	state.SetSP(0x7fff00000000)
	state.SetN(true)
	state.SetZ(true)

	testData := []byte{0x01, 0x02, 0x03, 0x04}
	state.WriteMemory(0x1000, testData)

	// Test cloning
	cloned := state.Clone().(*ARM64State)

	// Verify cloned state matches original
	if cloned.GetX(5) != state.GetX(5) {
		t.Error("Cloned state X5 doesn't match original")
	}
	if cloned.GetPC() != state.GetPC() {
		t.Error("Cloned state PC doesn't match original")
	}
	if cloned.GetSP() != state.GetSP() {
		t.Error("Cloned state SP doesn't match original")
	}
	if cloned.GetN() != state.GetN() || cloned.GetZ() != state.GetZ() {
		t.Error("Cloned state flags don't match original")
	}

	// Verify memory is cloned
	clonedData, err := cloned.ReadMemory(0x1000, len(testData))
	if err != nil {
		t.Fatalf("Failed to read memory from cloned state: %v", err)
	}
	if !bytes.Equal(testData, clonedData) {
		t.Error("Cloned state memory doesn't match original")
	}

	// Verify independence (modifying clone shouldn't affect original)
	cloned.SetX(5, 0xdeadbeef)
	if state.GetX(5) == 0xdeadbeef {
		t.Error("Modifying cloned state affected original")
	}

	// Test reset
	state.Reset()
	if state.GetX(5) != 0 || state.GetPC() != 0 || state.GetN() || state.GetZ() {
		t.Error("Reset didn't clear state properly")
	}

	// Verify memory is cleared
	_, err = state.ReadMemory(0x1000, len(testData))
	if err == nil {
		t.Error("Expected error reading memory after reset")
	}
}

func TestUtilityMethods(t *testing.T) {
	state := NewState()

	// Test SetupFunctionCall
	returnAddr := uint64(0x200000000)
	state.SetupFunctionCall(returnAddr)
	if state.GetX(30) != returnAddr {
		t.Errorf("Expected LR (X30) to be %x, got %x", returnAddr, state.GetX(30))
	}

	// Test SetupFunctionCallWithArgs
	args := []uint64{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9}
	state.SetupFunctionCallWithArgs(returnAddr, args...)

	// Check first 8 args are in X0-X7
	for i := 0; i < 8; i++ {
		if state.GetX(i) != args[i] {
			t.Errorf("Expected X%d to be %x, got %x", i, args[i], state.GetX(i))
		}
	}

	// Test SetStackPointer
	newSP := uint64(0x8000000000000000)
	state.SetStackPointer(newSP)
	if state.GetSP() != newSP {
		t.Errorf("Expected SP to be %x, got %x", newSP, state.GetSP())
	}

	// Test LoadRegisters
	values := []uint64{0x10, 0x20, 0x30, 0x40, 0x50}
	state.LoadRegisters(values)
	for i, val := range values {
		if state.GetX(i) != val {
			t.Errorf("Expected X%d to be %x, got %x", i, val, state.GetX(i))
		}
	}

	// Test ClearRegisters
	state.ClearRegisters()
	for i := 0; i < 31; i++ {
		if state.GetX(i) != 0 {
			t.Errorf("Expected X%d to be 0 after clear, got %x", i, state.GetX(i))
		}
		if state.IsValid(i) {
			t.Errorf("Expected X%d to be invalid after clear", i)
		}
	}

	// Test SetupTestCase
	startAddr := uint64(0x100000000)
	state.SetupTestCase(startAddr, returnAddr)
	if state.GetPC() != startAddr {
		t.Errorf("Expected PC to be %x, got %x", startAddr, state.GetPC())
	}
	if state.GetX(30) != returnAddr {
		t.Errorf("Expected LR to be %x, got %x", returnAddr, state.GetX(30))
	}
}

func TestContextSetup(t *testing.T) {
	state := NewState()

	// Test kernel context setup
	state.SetupKernelContext()
	if state.GetSP() != 0xfffffff000000000 {
		t.Errorf("Expected kernel SP to be 0xfffffff000000000, got %x", state.GetSP())
	}

	// Test user context setup
	state.SetupUserContext()
	if state.GetSP() != 0x7ffff0000000 {
		t.Errorf("Expected user SP to be 0x7ffff0000000, got %x", state.GetSP())
	}
}

func TestMultiBlockMemoryOperations(t *testing.T) {
	state := NewState()

	// Test writing across multiple aligned blocks
	addr := uint64(0x1005)   // Unaligned address
	data := make([]byte, 20) // Spans multiple 8-byte blocks
	for i := range data {
		data[i] = byte(i + 1)
	}

	state.WriteMemory(addr, data)

	// Read back the data
	readData, err := state.ReadMemory(addr, len(data))
	if err != nil {
		t.Fatalf("Failed to read multi-block memory: %v", err)
	}

	if !bytes.Equal(data, readData) {
		t.Errorf("Multi-block memory read/write mismatch. Expected %v, got %v", data, readData)
	}

	// Test reading partial data from blocks
	partialData, err := state.ReadMemory(addr+5, 10)
	if err != nil {
		t.Fatalf("Failed to read partial memory: %v", err)
	}

	expected := data[5:15]
	if !bytes.Equal(expected, partialData) {
		t.Errorf("Partial memory read mismatch. Expected %v, got %v", expected, partialData)
	}
}

func TestEndianness(t *testing.T) {
	state := NewState()
	addr := uint64(0x1000)

	// Test little-endian byte order for uint64
	value64 := uint64(0x123456789abcdef0)
	state.WriteUint64(addr, value64)

	// Read as bytes and verify little-endian order
	data, err := state.ReadMemory(addr, 8)
	if err != nil {
		t.Fatalf("Failed to read memory: %v", err)
	}

	expected := make([]byte, 8)
	binary.LittleEndian.PutUint64(expected, value64)

	if !bytes.Equal(data, expected) {
		t.Errorf("Endianness mismatch for uint64. Expected %v, got %v", expected, data)
	}

	// Test little-endian byte order for uint32
	value32 := uint32(0x12345678)
	state.WriteUint32(addr+8, value32)

	data32, err := state.ReadMemory(addr+8, 4)
	if err != nil {
		t.Fatalf("Failed to read memory: %v", err)
	}

	expected32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(expected32, value32)

	if !bytes.Equal(data32, expected32) {
		t.Errorf("Endianness mismatch for uint32. Expected %v, got %v", expected32, data32)
	}
}

// Tests for enhanced ARM64 state functionality

func TestRegisterValidityTracking(t *testing.T) {
	state := NewState()

	// Test initial state - all registers should be invalid
	for i := range uint8(31) {
		if state.IsRegisterValid(i) {
			t.Errorf("Register %d should be invalid initially", i)
		}
		if state.IsRegisterWide(i) {
			t.Errorf("Register %d should not be wide initially", i)
		}
	}

	// Test setting registers valid
	state.SetRegisterValid(5, true)
	state.SetRegisterWide(5, true)

	if !state.IsRegisterValid(5) {
		t.Error("Register 5 should be valid after setting")
	}
	if !state.IsRegisterWide(5) {
		t.Error("Register 5 should be wide after setting")
	}

	// Test setting register as 32-bit (not wide)
	state.SetRegisterValid(10, true)
	state.SetRegisterWide(10, false)

	if !state.IsRegisterValid(10) {
		t.Error("Register 10 should be valid after setting")
	}
	if state.IsRegisterWide(10) {
		t.Error("Register 10 should not be wide after setting to false")
	}

	// Test validity mask retrieval
	valid, wide, qvalid := state.GetValidityMask()
	expectedValid := uint32((1 << 5) | (1 << 10))
	expectedWide := uint32(1 << 5)

	if valid != expectedValid {
		t.Errorf("Expected valid mask 0x%x, got 0x%x", expectedValid, valid)
	}
	if wide != expectedWide {
		t.Errorf("Expected wide mask 0x%x, got 0x%x", expectedWide, wide)
	}
	if qvalid != 0 {
		t.Errorf("Expected qvalid mask 0, got 0x%x", qvalid)
	}

	// Test clearing validity
	state.ClearRegisterValidity()
	valid, wide, qvalid = state.GetValidityMask()

	if valid != 0 || wide != 0 || qvalid != 0 {
		t.Errorf("Expected all validity masks to be 0 after clearing, got valid=0x%x wide=0x%x qvalid=0x%x",
			valid, wide, qvalid)
	}

	// Test boundary conditions
	state.SetRegisterValid(31, true) // Should be ignored (invalid register)
	if state.IsRegisterValid(31) {
		t.Error("Register 31 should remain invalid (out of bounds)")
	}
}

func TestQRegisterSupport(t *testing.T) {
	state := NewState()

	// Test initial state
	for i := range uint8(32) {
		if state.IsQRegisterValid(i) {
			t.Errorf("Q register %d should be invalid initially", i)
		}
		q := state.GetQ(i)
		if q[0] != 0 || q[1] != 0 {
			t.Errorf("Q register %d should be zero initially, got [0x%x, 0x%x]", i, q[0], q[1])
		}
	}

	// Test setting Q registers
	state.SetQ(0, 0x123456789abcdef0, 0xfedcba9876543210)

	if !state.IsQRegisterValid(0) {
		t.Error("Q register 0 should be valid after setting")
	}

	q := state.GetQ(0)
	if q[0] != 0x123456789abcdef0 || q[1] != 0xfedcba9876543210 {
		t.Errorf("Q register 0 value mismatch, expected [0x123456789abcdef0, 0xfedcba9876543210], got [0x%x, 0x%x]",
			q[0], q[1])
	}

	// Test Q register validity tracking
	state.SetQRegisterValid(15, true)
	if !state.IsQRegisterValid(15) {
		t.Error("Q register 15 should be valid after setting")
	}

	state.SetQRegisterValid(15, false)
	if state.IsQRegisterValid(15) {
		t.Error("Q register 15 should be invalid after clearing")
	}

	// Test boundary conditions
	state.SetQ(32, 1, 2) // Should be ignored
	q = state.GetQ(32)
	if q[0] != 0 || q[1] != 0 {
		t.Error("Out of bounds Q register access should return zero")
	}
}

func TestHostMemoryTracking(t *testing.T) {
	state := NewState()

	// Test initial state
	for i := range uint8(32) {
		if state.IsHostMemoryValid(i) != 0 {
			t.Errorf("Host memory index %d should be 0 initially", i)
		}
	}

	// Test setting memory validity levels
	state.SetHostMemoryValid(0, 1)
	if state.IsHostMemoryValid(0) != 1 {
		t.Errorf("Host memory index 0 should be 1, got %d", state.IsHostMemoryValid(0))
	}

	state.SetHostMemoryValid(5, 3)
	if state.IsHostMemoryValid(5) != 3 {
		t.Errorf("Host memory index 5 should be 3, got %d", state.IsHostMemoryValid(5))
	}

	// Verify other indices are unaffected
	if state.IsHostMemoryValid(1) != 0 {
		t.Error("Host memory index 1 should remain 0")
	}

	// Test setting multiple indices
	state.SetHostMemoryValid(10, 2)
	state.SetHostMemoryValid(31, 1)

	if state.IsHostMemoryValid(10) != 2 || state.IsHostMemoryValid(31) != 1 {
		t.Error("Multiple host memory indices not set correctly")
	}

	// Test boundary conditions
	state.SetHostMemoryValid(32, 1) // Should be ignored
	state.SetHostMemoryValid(0, 4)  // Should be ignored (level > 3)

	if state.IsHostMemoryValid(0) != 1 { // Should still be 1
		t.Error("Invalid host memory operations should be ignored")
	}
}

func TestValidatePointer(t *testing.T) {
	state := NewState()

	// Set up a register with a value
	state.SetX(5, 0x100000)
	state.SetRegisterValid(5, true)

	// Test valid pointer within bounds
	if !state.ValidatePointer(5, 0x80000, 0x200000) {
		t.Error("Pointer should be valid within bounds")
	}

	// Test pointer outside bounds
	if state.ValidatePointer(5, 0x200000, 0x300000) {
		t.Error("Pointer should be invalid outside bounds")
	}

	// Test invalid register
	if state.ValidatePointer(31, 0x80000, 0x200000) {
		t.Error("Invalid register should fail pointer validation")
	}

	// Test register without valid data
	state.SetX(10, 0x100000)
	// Don't mark as valid - the SetX call above should have marked it valid automatically
	// So we need to explicitly clear the validity
	state.SetRegisterValid(10, false)
	if state.ValidatePointer(10, 0x80000, 0x200000) {
		t.Error("Register without valid data should fail pointer validation")
	}
}

func TestEmulationSafety(t *testing.T) {
	state := NewState()

	// Test initial state
	if state.GetInstructionCount() != 0 {
		t.Error("Initial instruction count should be 0")
	}
	if state.HasEmulationError() {
		t.Error("Should not have emulation error initially")
	}

	// Test instruction counting
	err := state.IncrementInstructionCount()
	if err != nil {
		t.Errorf("Should not error on first increment: %v", err)
	}
	if state.GetInstructionCount() != 1 {
		t.Errorf("Instruction count should be 1, got %d", state.GetInstructionCount())
	}

	// Test max instruction limit
	state.SetMaxInstructions(5)
	for range 4 {
		err := state.IncrementInstructionCount()
		if err != nil {
			t.Errorf("Should not error before limit: %v", err)
		}
	}

	// Should error on exceeding limit
	err = state.IncrementInstructionCount()
	if err == nil {
		t.Error("Should error when exceeding max instructions")
	}

	// Test error tracking
	state.SetEmulationError(0x1000, "test error")
	if !state.HasEmulationError() {
		t.Error("Should have emulation error after setting")
	}

	instrCount, hasError, errorMsg := state.GetEmulationStatus()
	if instrCount != 6 || !hasError || errorMsg != "test error" {
		t.Errorf("Emulation status incorrect: count=%d hasError=%t msg=%s",
			instrCount, hasError, errorMsg)
	}

	// Test clearing error
	state.ClearEmulationError()
	if state.HasEmulationError() {
		t.Error("Should not have error after clearing")
	}

	// Test resetting instruction count
	state.ResetInstructionCount()
	if state.GetInstructionCount() != 0 {
		t.Error("Instruction count should be 0 after reset")
	}
}

func TestStateValidation(t *testing.T) {
	state := NewState()

	// Test valid state
	err := state.ValidateInternalState()
	if err != nil {
		t.Errorf("Valid state should not error: %v", err)
	}

	// Test instruction count validation
	state.instructionCount = state.maxInstructions + 1
	err = state.ValidateInternalState()
	if err == nil {
		t.Error("Should error when instruction count exceeds max")
	}

	// Reset to valid state
	state.instructionCount = 0
	err = state.ValidateInternalState()
	if err != nil {
		t.Errorf("Reset state should be valid: %v", err)
	}
}

func TestStateComparison(t *testing.T) {
	state1 := NewState()
	state2 := NewState()

	// Initially identical states should have no differences
	diffs := state1.CompareState(state2)
	if len(diffs) != 0 {
		t.Errorf("Identical states should have no differences, got: %v", diffs)
	}

	// Test differences in basic fields
	state1.SetPC(0x1000)
	diffs = state1.CompareState(state2)
	if len(diffs) == 0 {
		t.Error("States with different PC should have differences")
	}

	// Test differences in registers
	state1 = NewState()
	state1.SetX(5, 0x123)
	diffs = state1.CompareState(state2)
	if len(diffs) == 0 {
		t.Error("States with different registers should have differences")
	}

	// Test nil comparison
	diffs = state1.CompareState(nil)
	if len(diffs) == 0 || diffs[0] != "other state is nil" {
		t.Error("Comparison with nil should return appropriate error")
	}
}

func TestDumpRegisterState(t *testing.T) {
	state := NewState()

	// Set up some state
	state.SetPC(0x1000)
	state.SetX(5, 0x123456)
	state.SetRegisterValid(5, true)
	state.SetRegisterWide(5, true)

	dump := state.DumpRegisterState()

	// Should contain basic info
	if !strings.Contains(dump, "PC: 0x0000000000001000") {
		t.Error("Dump should contain PC value")
	}
	if !strings.Contains(dump, "Valid: 0x00000020") {
		t.Error("Dump should contain validity mask")
	}
	if !strings.Contains(dump, "X5 :X") {
		t.Error("Dump should show register 5 as valid and wide")
	}

	// Test with error
	state.SetEmulationError(0x2000, "test error")
	dump = state.DumpRegisterState()
	if !strings.Contains(dump, "ERROR at PC") {
		t.Error("Dump should contain error information")
	}
}

func TestBackwardCompatibility(t *testing.T) {
	state := NewState()

	// Test that legacy methods still work
	state.setValid(5, true)
	if !state.IsValid(5) {
		t.Error("Legacy setValid should work")
	}

	state.setWide(10, false)
	if state.IsWide(10) {
		t.Error("Legacy setWide should work")
	}

	// Test that existing operations still function
	state.SetX(15, 0x12345678)
	if state.GetX(15) != 0x12345678 {
		t.Error("Basic register operations should work")
	}

	// Verify that setting X register marks it as valid (existing behavior)
	if !state.IsValid(15) {
		t.Error("Setting X register should mark it as valid")
	}
	if !state.IsWide(15) {
		t.Error("Setting X register should mark it as wide")
	}
}
