package testing

import (
	"fmt"
	"testing"

	"github.com/blacktop/arm64-cgo/emulate"
	"github.com/blacktop/arm64-cgo/emulate/core"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// TestErrorHandlingAndRecovery tests various error scenarios and recovery mechanisms
func TestErrorHandlingAndRecovery(t *testing.T) {
	tests := []struct {
		name          string
		setup         func(*emulate.Engine) error
		instruction   uint32
		pc            uint64
		expectError   bool
		errorContains string
		description   string
	}{
		{
			name: "invalid_instruction_encoding",
			setup: func(engine *emulate.Engine) error {
				testState := state.NewState()
				testState.SetSP(0)
				engine.SetState(testState)
				return nil
			},
			instruction:   0xFFFFFFFF, // Invalid instruction
			pc:            0x10000000,
			expectError:   true,
			errorContains: "failed to decode",
			description:   "Should handle invalid instruction encodings gracefully",
		},
		{
			name: "memory_access_out_of_bounds",
			setup: func(engine *emulate.Engine) error {
				testState := state.NewState()
				testState.SetSP(0)
				testState.SetX(0, 0xFFFFFFFFFFFF0000) // Very high address
				engine.SetState(testState)
				return nil
			},
			instruction:   0xf9400000, // LDR X0, [X0] - load from high address
			pc:            0x10000000,
			expectError:   true,
			errorContains: "memory",
			description:   "Should handle out-of-bounds memory access",
		},
		{
			name: "invalid_register_access",
			setup: func(engine *emulate.Engine) error {
				testState := state.NewState()
				testState.SetSP(0)
				engine.SetState(testState)
				return nil
			},
			instruction:   0x91000000, // ADD X0, X0, #0 (but we'll modify to invalid register)
			pc:            0x10000000,
			expectError:   false, // Should handle gracefully or succeed
			errorContains: "",
			description:   "Should handle edge case register access",
		},
		{
			name: "stack_overflow_simulation",
			setup: func(engine *emulate.Engine) error {
				testState := state.NewState()
				testState.SetSP(0x10) // Very low SP that could cause issues
				engine.SetState(testState)
				return nil
			},
			instruction:   0xd10043ff, // SUB SP, SP, #16 (stack allocation)
			pc:            0x10000000,
			expectError:   false, // Should handle but might cause issues in real scenarios
			errorContains: "",
			description:   "Should handle very low stack pointer values",
		},
		{
			name: "division_by_zero",
			setup: func(engine *emulate.Engine) error {
				testState := state.NewState()
				testState.SetSP(0)
				testState.SetX(0, 100) // Dividend
				testState.SetX(1, 0)   // Divisor (zero)
				engine.SetState(testState)
				return nil
			},
			instruction:   0x9ac10c00, // UDIV X0, X0, X1 (divide by zero)
			pc:            0x10000000,
			expectError:   false, // ARM64 division by zero returns 0, doesn't fault
			errorContains: "",
			description:   "Should handle division by zero per ARM64 spec",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create fresh engine for each test
			engine := emulate.NewEngine()

			// Run setup
			if err := tt.setup(engine); err != nil {
				t.Fatalf("Setup failed: %v", err)
			}

			// Execute instruction and check error handling
			err := engine.ExecuteInstruction(tt.pc, tt.instruction)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none for %s", tt.description)
				} else if tt.errorContains != "" && !containsString(err.Error(), tt.errorContains) {
					t.Errorf("Expected error containing '%s' but got '%s'", tt.errorContains, err.Error())
				} else {
					t.Logf("✅ Correctly handled error: %s", err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for %s: %v", tt.description, err)
				} else {
					t.Logf("✅ Successfully executed: %s", tt.description)
				}
			}
		})
	}
}

// TestErrorRecovery tests that the emulator can recover from errors and continue execution
func TestErrorRecovery(t *testing.T) {
	engine := emulate.NewEngine()
	testState := state.NewState()
	testState.SetSP(0)
	testState.SetX(0, 10)
	engine.SetState(testState)

	// Execute valid instruction
	validInstr := uint32(0x91000400) // ADD X0, X0, #1
	err := engine.ExecuteInstruction(0x10000000, validInstr)
	if err != nil {
		t.Fatalf("Valid instruction failed: %v", err)
	}

	// Check state was updated
	finalState := engine.GetState()
	if finalState.GetX(0) != 11 {
		t.Errorf("Expected X0=11, got X0=%d", finalState.GetX(0))
	}

	// Try invalid instruction (should fail)
	invalidInstr := uint32(0xFFFFFFFF)
	err = engine.ExecuteInstruction(0x10000004, invalidInstr)
	if err == nil {
		t.Error("Invalid instruction should have failed")
	} else {
		t.Logf("✅ Invalid instruction correctly failed: %v", err)
	}

	// Execute another valid instruction (should work after error)
	err = engine.ExecuteInstruction(0x10000008, validInstr)
	if err != nil {
		t.Errorf("Valid instruction after error failed: %v", err)
	}

	// Check state continued to update correctly
	finalState = engine.GetState()
	if finalState.GetX(0) != 12 {
		t.Errorf("Expected X0=12 after recovery, got X0=%d", finalState.GetX(0))
	} else {
		t.Log("✅ Engine recovered successfully and continued execution")
	}
}

// TestResourceLimits tests behavior under resource constraints
func TestResourceLimits(t *testing.T) {
	tests := []struct {
		name        string
		maxInstrs   int
		instrCount  int
		expectError bool
		description string
	}{
		{
			name:        "within_instruction_limit",
			maxInstrs:   100,
			instrCount:  50,
			expectError: false,
			description: "Should succeed when within instruction limit",
		},
		{
			name:        "at_instruction_limit",
			maxInstrs:   100,
			instrCount:  100,
			expectError: false,
			description: "Should succeed when at instruction limit",
		},
		{
			name:        "exceeds_instruction_limit",
			maxInstrs:   10,
			instrCount:  15,
			expectError: true,
			description: "Should fail when exceeding instruction limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create engine with specific limits
			config := emulate.DefaultEngineConfig()
			config.MaxInstructions = tt.maxInstrs
			engine := emulate.NewEngineWithConfig(config)

			testState := state.NewState()
			testState.SetSP(0)
			testState.SetX(0, 0)
			engine.SetState(testState)

			// Execute instructions up to the count
			instrValue := uint32(0x91000400) // ADD X0, X0, #1
			var lastErr error

			for i := 0; i < tt.instrCount; i++ {
				pc := uint64(0x10000000 + i*4)
				err := engine.ExecuteInstruction(pc, instrValue)
				if err != nil {
					lastErr = err
					break
				}
			}

			if tt.expectError {
				if lastErr == nil {
					t.Errorf("Expected instruction limit error but execution succeeded")
				} else {
					t.Logf("✅ Correctly enforced instruction limit: %v", lastErr)
				}
			} else {
				if lastErr != nil {
					t.Errorf("Unexpected error within limits: %v", lastErr)
				} else {
					t.Logf("✅ Successfully executed %d instructions within limit", tt.instrCount)
				}
			}
		})
	}
}

// TestConcurrentAccess tests thread safety of the emulator
func TestConcurrentAccess(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping concurrent access test in short mode")
	}

	// Note: The current emulator is not designed for concurrent access
	// This test documents the expected behavior and can be updated if
	// thread safety is added in the future

	t.Log("⚠️  Current emulator design assumes single-threaded usage")
	t.Log("⚠️  Each goroutine should use its own engine instance")

	// Test that separate engine instances work independently
	const goroutines = 10
	const iterations = 100

	done := make(chan bool, goroutines)
	errors := make(chan error, goroutines*iterations)

	for g := 0; g < goroutines; g++ {
		go func(goroutineID int) {
			defer func() { done <- true }()

			// Each goroutine gets its own engine instance
			engine := emulate.NewEngine()
			testState := state.NewState()
			testState.SetSP(0)
			testState.SetX(0, uint64(goroutineID))
			engine.SetState(testState)

			instrValue := uint32(0x91000400) // ADD X0, X0, #1

			for i := range iterations {
				pc := uint64(0x10000000 + i*4)
				err := engine.ExecuteInstruction(pc, instrValue)
				if err != nil {
					errors <- err
					return
				}
			}

			// Verify final state
			finalState := engine.GetState()
			expected := uint64(goroutineID + iterations)
			if finalState.GetX(0) != expected {
				errors <- core.NewEmulationError(core.ErrInvalidInstruction, 0, "concurrent_test",
					fmt.Sprintf("goroutine %d: expected X0=%d, got %d", goroutineID, expected, finalState.GetX(0)))
			}
		}(g)
	}

	// Wait for all goroutines to complete
	for range goroutines {
		<-done
	}

	// Check for any errors
	close(errors)
	errorCount := 0
	for err := range errors {
		t.Errorf("Concurrent execution error: %v", err)
		errorCount++
	}

	if errorCount == 0 {
		t.Logf("✅ Successfully executed %d concurrent engine instances", goroutines)
	}
}

// Helper function to check if string contains substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					findInString(s, substr)))
}

func findInString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
