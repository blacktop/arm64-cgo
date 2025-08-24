package state

import (
	"testing"

	"github.com/blacktop/arm64-cgo/emulate/core"
)

// TestStateInterfaceCompliance verifies that ARM64State implements core.State interface
func TestStateInterfaceCompliance(t *testing.T) {
	// This test ensures that ARM64State implements the core.State interface
	var _ core.State = (*ARM64State)(nil)

	// Create a state instance and verify it can be used as core.State
	var state core.State = NewState()

	// Test basic interface methods
	state.SetX(0, 0x123456789abcdef0)
	if state.GetX(0) != 0x123456789abcdef0 {
		t.Error("Interface method GetX/SetX not working correctly")
	}

	state.SetPC(0x100000000)
	if state.GetPC() != 0x100000000 {
		t.Error("Interface method GetPC/SetPC not working correctly")
	}

	state.SetN(true)
	if !state.GetN() {
		t.Error("Interface method GetN/SetN not working correctly")
	}

	// Test memory operations through interface
	testData := []byte{0x01, 0x02, 0x03, 0x04}
	state.WriteMemory(0x1000, testData)

	readData, err := state.ReadMemory(0x1000, len(testData))
	if err != nil {
		t.Fatalf("Interface memory read failed: %v", err)
	}

	for i, b := range testData {
		if readData[i] != b {
			t.Errorf("Interface memory operation failed at byte %d: expected %x, got %x", i, b, readData[i])
		}
	}

	// Test cloning through interface
	cloned := state.Clone()
	if cloned.GetX(0) != state.GetX(0) {
		t.Error("Interface Clone method not working correctly")
	}

	// Test reset through interface
	state.Reset()
	if state.GetX(0) != 0 || state.GetPC() != 0 {
		t.Error("Interface Reset method not working correctly")
	}
}

// TestStateFactoryPattern tests creating states with different configurations
func TestStateFactoryPattern(t *testing.T) {
	// Test default state creation
	defaultState := NewState()
	if defaultState.GetSP() != 0x7ffff0000000 {
		t.Error("Default state creation failed")
	}

	// Test configured state creation
	config := core.StateConfig{
		StackSize:   0x200000,
		StackBase:   0x8000000000000000,
		InitialPC:   0x100000000,
		InitialSP:   0x8000000000000000,
		EnableDebug: true,
	}

	configuredState := NewStateWithConfig(config)
	if configuredState.GetPC() != config.InitialPC {
		t.Error("Configured state creation failed")
	}
	if configuredState.GetSP() != config.InitialSP {
		t.Error("Configured state SP initialization failed")
	}
}

// TestStateWithCustomHandlers tests the state with custom memory handlers
func TestStateWithCustomHandlers(t *testing.T) {
	state := NewState()

	// Test with custom memory handler
	customMemoryData := map[uint64][]byte{
		0x10000000: {0xde, 0xad, 0xbe, 0xef},
		0x20000000: {0xca, 0xfe, 0xba, 0xbe},
	}

	state.SetMemoryReadHandler(func(addr uint64, size int) ([]byte, error) {
		if data, exists := customMemoryData[addr]; exists && size <= len(data) {
			return data[:size], nil
		}
		return nil, core.ErrMemoryAccess
	})

	// Test reading through custom handler
	data, err := state.ReadMemory(0x10000000, 4)
	if err != nil {
		t.Fatalf("Custom memory handler failed: %v", err)
	}

	expected := customMemoryData[0x10000000]
	for i, b := range expected {
		if data[i] != b {
			t.Errorf("Custom memory handler data mismatch at byte %d: expected %x, got %x", i, b, data[i])
		}
	}

	// Test with custom string handler
	customStrings := map[uint64]string{
		0x30000000: "Hello, World!",
		0x40000000: "Custom String Handler",
	}

	state.SetStringPoolHandler(func(addr uint64) (string, error) {
		if str, exists := customStrings[addr]; exists {
			return str, nil
		}
		return "", core.ErrMemoryAccess
	})

	str, err := state.ReadString(0x30000000)
	if err != nil {
		t.Fatalf("Custom string handler failed: %v", err)
	}

	if str != customStrings[0x30000000] {
		t.Errorf("Custom string handler returned wrong string: expected '%s', got '%s'", customStrings[0x30000000], str)
	}

	// Test with custom pointer resolver
	customPointers := map[uint64]uint64{
		0x50000000: 0x60000000,
		0x60000000: 0x70000000,
	}

	state.SetPointerResolver(func(addr uint64) (uint64, error) {
		if target, exists := customPointers[addr]; exists {
			return target, nil
		}
		return 0, core.ErrMemoryAccess
	})

	resolved, err := state.ResolvePointer(0x50000000)
	if err != nil {
		t.Fatalf("Custom pointer resolver failed: %v", err)
	}

	if resolved != customPointers[0x50000000] {
		t.Errorf("Custom pointer resolver returned wrong address: expected %x, got %x", customPointers[0x50000000], resolved)
	}
}

// TestStateErrorHandling tests error conditions and edge cases
func TestStateErrorHandling(t *testing.T) {
	state := NewState()

	// Test invalid register access (should not panic)
	if state.GetX(-1) != 0 {
		t.Error("Invalid register access should return 0")
	}
	if state.GetX(31) != 0 {
		t.Error("Invalid register access should return 0")
	}

	// Test memory access to unmapped regions
	_, err := state.ReadMemory(0xdeadbeef, 8)
	if err == nil {
		t.Error("Expected error when reading unmapped memory")
	}

	// Test string reading from invalid address
	_, err = state.ReadString(0xdeadbeef)
	if err == nil {
		t.Error("Expected error when reading string from unmapped memory")
	}

	// Test pointer resolution from invalid address
	_, err = state.ResolvePointer(0xdeadbeef)
	if err == nil {
		t.Error("Expected error when resolving pointer from unmapped memory")
	}
}

// TestStatePerformance tests basic performance characteristics
func TestStatePerformance(t *testing.T) {
	state := NewState()

	// Test register access performance (should be very fast)
	for i := range 10000 {
		state.SetX(i%31, uint64(i))
		_ = state.GetX(i % 31)
	}

	// Test memory access performance
	testData := make([]byte, 1024)
	for i := range testData {
		testData[i] = byte(i)
	}

	for i := 0; i < 1000; i++ {
		addr := uint64(0x100000 + i*1024)
		state.WriteMemory(addr, testData)
		_, err := state.ReadMemory(addr, len(testData))
		if err != nil {
			t.Fatalf("Memory operation failed at iteration %d: %v", i, err)
		}
	}
}
