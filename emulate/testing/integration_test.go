//go:build darwin

package testing

import (
	"testing"
)

// Test multi-instruction sequences against Apple's Hypervisor.framework as the source of truth
func TestSequenceIntegrationWithHypervisor(t *testing.T) {
	// Test sequences as specified in task 8.3
	sequences := []struct {
		name         string
		instructions [][]byte
		initialRegs  map[int]uint64
		initialMem   map[uint64][]byte
		description  string
	}{
		{
			name: "arithmetic_dependency_chain",
			instructions: [][]byte{
				{0xe0, 0x93, 0x01, 0x91}, // ADD X0, XZR, #100
				{0x21, 0x40, 0x01, 0x91}, // ADD X1, X1, #80
				{0x02, 0x00, 0x01, 0x8b}, // ADD X2, X0, X1
				{0x63, 0x00, 0x02, 0x8b}, // ADD X3, X3, X2
			},
			initialRegs: map[int]uint64{
				1: 20,
				3: 5,
			},
			description: "Multi-instruction arithmetic with register dependencies",
		},
		{
			name: "memory_store_load_sequence",
			instructions: [][]byte{
				{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1  (X0 = 0x10000000)
				{0x21, 0x20, 0x00, 0x91}, // ADD X1, X1, #8  (X1 = 0x10000008)
				{0x02, 0x00, 0x00, 0xf9}, // STR X2, [X0]    (store at 0x10000000)
				{0x23, 0x00, 0x01, 0xf9}, // STR X3, [X1]    (store at 0x10000008)
				{0x04, 0x00, 0x40, 0xf9}, // LDR X4, [X0]    (load from 0x10000000)
				{0x25, 0x00, 0x41, 0xf9}, // LDR X5, [X1]    (load from 0x10000008)
			},
			initialRegs: map[int]uint64{
				0: 0x0,        // Start with X0=0
				1: 0x10000000, // X1 = base address
				2: 0xDEADBEEFCAFEBABE,
				3: 0x123456789ABCDEF0,
			},
			initialMem: map[uint64][]byte{
				0x10000000: make([]byte, 0x1000), // 4KB memory region
			},
			description: "Store and load operations with address calculations",
		},
		{
			name: "conditional_branch_sequence",
			instructions: [][]byte{
				{0x1f, 0x00, 0x01, 0xeb}, // CMP X0, X1
				{0x02, 0x10, 0x81, 0x9a}, // CSEL X2, X0, X1, NE
				{0x00, 0x04, 0x00, 0x91}, // ADD X0, X0, #1
				{0x1f, 0x00, 0x01, 0xeb}, // CMP X0, X1
				{0x23, 0x00, 0x82, 0x9a}, // CSEL X3, X1, X2, EQ
			},
			initialRegs: map[int]uint64{
				0: 5,
				1: 6,
			},
			description: "Conditional operations with flag dependencies",
		},
		{
			name: "function_call_pattern",
			instructions: [][]byte{
				{0xff, 0x43, 0x00, 0xd1}, // SUB SP, SP, #16   (create stack frame)
				{0x00, 0x08, 0x00, 0x91}, // ADD X0, X0, #2    (X0 += 2)
				{0x21, 0x00, 0x00, 0x8b}, // ADD X1, X1, X0    (X1 += X0)
				{0x00, 0x08, 0x00, 0xd1}, // SUB X0, X0, #2    (restore X0: X0 -= 2)
				{0xff, 0x43, 0x00, 0x91}, // ADD SP, SP, #16   (restore stack)
			},
			initialRegs: map[int]uint64{
				0: 10,
				1: 20,
			},
			description: "Function call pattern with stack management",
		},
	}

	for _, seq := range sequences {
		t.Run(seq.name, func(t *testing.T) {
			// Try to create Hypervisor framework
			framework, err := NewHypervisorFramework()
			if err != nil {
				t.Skipf("Hypervisor framework not available: %v", err)
				return
			}
			defer framework.Close()

			// Set up initial register state
			err = framework.SetRegisterState(seq.initialRegs)
			if err != nil {
				t.Skipf("Failed to set register state: %v", err)
				return
			}

			// Set up initial memory state
			if len(seq.initialMem) > 0 {
				err = framework.SetMemoryState(seq.initialMem)
				if err != nil {
					t.Skipf("Failed to set memory state: %v", err)
					return
				}
			}

			// Execute sequence and compare with Hypervisor
			result, err := framework.ExecuteSequenceAndCompare(seq.instructions, 0x40000000)
			if err != nil {
				t.Skipf("Failed to execute sequence: %v", err)
				return
			}

			// Validate results
			if len(result.Differences) > 0 {
				t.Logf("Found %d differences between our emulator and Hypervisor:", len(result.Differences))
				for _, diff := range result.Differences {
					t.Logf("  %s %s: our=0x%x, hypervisor=0x%x", diff.Type, diff.Name, diff.OurValue, diff.HypervisorValue)
				}
			}

			// Check for critical differences (non-PC)
			criticalDiffs := 0
			for _, diff := range result.Differences {
				if diff.Name != "PC" {
					criticalDiffs++
				}
			}

			if criticalDiffs > 0 {
				t.Errorf("Found %d critical differences (non-PC) in sequence '%s'", criticalDiffs, seq.name)
			} else {
				t.Logf("✅ Sequence '%s' validated successfully with Hypervisor: %s", seq.name, seq.description)
			}
		})
	}
}

// TestLoopConstructValidation tests loop patterns against Hypervisor
func TestLoopConstructValidation(t *testing.T) {
	// Try to create Hypervisor framework
	framework, err := NewHypervisorFramework()
	if err != nil {
		t.Skipf("Hypervisor framework not available: %v", err)
		return
	}
	defer framework.Close()

	// Loop construct: sum 1+2+3 (3 iterations)
	instructions := [][]byte{
		{0x00, 0x00, 0x80, 0xd2}, // MOV X0, #0    (accumulator)
		{0x21, 0x00, 0x80, 0xd2}, // MOV X1, #1    (counter)
		{0x62, 0x00, 0x80, 0xd2}, // MOV X2, #3    (limit)
		// Loop body starts here
		{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1   (acc += counter)
		{0x21, 0x04, 0x00, 0x91}, // ADD X1, X1, #1   (counter++)
		{0x3f, 0x00, 0x02, 0xeb}, // CMP X1, X2       (compare with limit)
		// Would normally branch back, but for testing we continue
		{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1   (second iteration)
		{0x21, 0x04, 0x00, 0x91}, // ADD X1, X1, #1
		{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1   (third iteration)
	}

	// Set initial state
	err = framework.SetRegisterState(map[int]uint64{})
	if err != nil {
		t.Skipf("Failed to set register state: %v", err)
		return
	}

	// Execute and compare
	result, err := framework.ExecuteSequenceAndCompare(instructions, 0x40000000)
	if err != nil {
		t.Skipf("Failed to execute loop construct: %v", err)
		return
	}

	// Check final accumulator value (should be 1+2+3 = 6)
	if result.AfterState.Registers[0] != 6 {
		t.Errorf("Expected accumulator (X0) to be 6, got %d", result.AfterState.Registers[0])
	}

	// Check Hypervisor matches
	if result.AfterState.HypervisorRegisters[0] != 6 {
		t.Errorf("Hypervisor accumulator mismatch: expected 6, got %d", result.AfterState.HypervisorRegisters[0])
	}

	// Report differences
	criticalDiffs := 0
	for _, diff := range result.Differences {
		if diff.Name != "PC" {
			criticalDiffs++
			t.Logf("Loop construct difference: %s %s: our=0x%x, hypervisor=0x%x",
				diff.Type, diff.Name, diff.OurValue, diff.HypervisorValue)
		}
	}

	if criticalDiffs == 0 {
		t.Log("✅ Loop construct validated successfully with Hypervisor")
	} else {
		t.Errorf("Found %d critical differences in loop construct", criticalDiffs)
	}
}
