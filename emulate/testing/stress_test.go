package testing

import (
	"encoding/binary"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/blacktop/arm64-cgo/emulate"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// StressTest represents a stress testing scenario
type StressTest struct {
	name           string
	instructionGen func(int) [][]byte
	iterations     int
	initialState   map[int]uint64
	description    string
}

// TestLargeInstructionSequences tests emulator with very large instruction sequences
func TestLargeInstructionSequences(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large instruction sequence test in short mode")
	}

	tests := []StressTest{
		{
			name: "arithmetic_heavy_sequence",
			instructionGen: func(count int) [][]byte {
				// Use shared instruction patterns
				pattern := [][]byte{
					CommonInstructions.AddX0Imm1,
					CommonInstructions.AddX1X0,
					CommonInstructions.SubX0Imm1,
					CommonInstructions.AddX2X1X0,
				}

				instrs := make([][]byte, count)
				for i := 0; i < count; i++ {
					instrs[i] = pattern[i%4]
				}
				return instrs
			},
			iterations:   10000,
			initialState: map[int]uint64{0: 0, 1: 0, 2: 0},
			description:  "Large sequence of arithmetic operations",
		},
		{
			name: "register_intensive_sequence",
			instructionGen: func(count int) [][]byte {
				instrs := make([][]byte, count)
				for i := 0; i < count; i++ {
					// Cycle through different register combinations
					src := i % 30
					dst := (i + 1) % 30
					// ADD Xdst, Xsrc, #1
					instrBytes := make([]byte, 4)
					binary.LittleEndian.PutUint32(instrBytes, 0x91000400|uint32(dst)|uint32(src<<5))
					instrs[i] = instrBytes
				}
				return instrs
			},
			iterations:   5000,
			initialState: map[int]uint64{}, // Start with all zeros
			description:  "Sequence using many different registers",
		},
		{
			name: "mixed_operation_sequence",
			instructionGen: func(count int) [][]byte {
				instrs := make([][]byte, count)
				rand.Seed(42) // Deterministic for reproducible tests

				// Use shared instruction patterns
				validInstrs := [][]byte{
					CommonInstructions.AddX0Imm1,
					CommonInstructions.AddX1X0,
					CommonInstructions.AddX2X1X0,
					CommonInstructions.SubX0Imm1,
					// Add some variations with immediate values
					{0x21, 0x04, 0x00, 0x91}, // ADD X1, X1, #1
					{0x42, 0x04, 0x00, 0x91}, // ADD X2, X2, #1
					{0x21, 0x04, 0x00, 0xd1}, // SUB X1, X1, #1
				}

				for i := range count {
					instrs[i] = validInstrs[rand.Intn(len(validInstrs))]
				}
				return instrs
			},
			iterations:   8000,
			initialState: map[int]uint64{0: 100, 1: 200, 2: 300},
			description:  "Mixed operations with random instruction selection",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Logf("Starting stress test: %s", test.description)
			startTime := time.Now()

			// Generate instruction sequence
			instructions := test.instructionGen(test.iterations)

			// Setup engine
			engine := emulate.NewEngine()
			testState := state.NewState()
			testState.SetSP(0)

			// Set initial state
			for reg, val := range test.initialState {
				testState.SetX(reg, val)
			}
			engine.SetState(testState)

			// Execute all instructions
			pc := uint64(0x10000000)
			for i, instrBytes := range instructions {
				instrValue := binary.LittleEndian.Uint32(instrBytes)
				err := engine.ExecuteInstruction(pc, instrValue)
				if err != nil {
					t.Fatalf("Instruction %d failed: %v", i, err)
				}
				pc += 4
			}

			executionTime := time.Since(startTime)
			finalState := engine.GetState()

			t.Logf("✅ Successfully executed %d instructions in %v", test.iterations, executionTime)
			t.Logf("   Performance: %.2f instructions/second", float64(test.iterations)/executionTime.Seconds())
			t.Logf("   Final state: X0=%d, X1=%d, X2=%d",
				finalState.GetX(0), finalState.GetX(1), finalState.GetX(2))
		})
	}
}

// TestMemoryIntensiveOperations tests memory-heavy instruction sequences
func TestMemoryIntensiveOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping memory intensive test in short mode")
	}

	const iterations = 1000
	const memorySize = 0x10000 // 64KB

	t.Logf("Testing memory-intensive operations with %d iterations", iterations)

	// Setup engine with large memory region
	engine := emulate.NewEngine()
	testState := state.NewState()
	testState.SetSP(0)

	// Allocate memory region
	baseAddr := uint64(0x10000000)
	testState.WriteMemory(baseAddr, make([]byte, memorySize))

	// Set initial registers
	testState.SetX(0, baseAddr)   // Base address
	testState.SetX(1, 0xDEADBEEF) // Value to store
	engine.SetState(testState)

	startTime := time.Now()
	pc := uint64(0x20000000)

	// Stress test with many memory operations
	for i := range iterations {
		offset := uint64(i * 8) // 8-byte aligned offsets

		// Update address register: ADD X0, XZR, #(baseAddr + offset)
		// This is complex to encode properly, so use simpler approach
		currentState := engine.GetState()
		currentState.SetX(0, baseAddr+offset)
		engine.SetState(currentState)

		// STR X1, [X0] - Store value at current address
		storeInstr := uint32(0xf9000001)
		err := engine.ExecuteInstruction(pc, storeInstr)
		if err != nil {
			t.Fatalf("Store instruction %d failed: %v", i, err)
		}
		pc += 4

		// LDR X2, [X0] - Load value back
		loadInstr := uint32(0xf9400002)
		err = engine.ExecuteInstruction(pc, loadInstr)
		if err != nil {
			t.Fatalf("Load instruction %d failed: %v", i, err)
		}
		pc += 4

		// Verify loaded value matches stored value
		finalState := engine.GetState()
		if finalState.GetX(2) != 0xDEADBEEF {
			t.Errorf("Memory integrity check failed at iteration %d: expected 0xDEADBEEF, got 0x%x",
				i, finalState.GetX(2))
		}
	}

	executionTime := time.Since(startTime)
	totalOps := iterations * 2 // store + load per iteration

	t.Logf("✅ Successfully executed %d memory operations in %v", totalOps, executionTime)
	t.Logf("   Performance: %.2f memory ops/second", float64(totalOps)/executionTime.Seconds())
	t.Logf("   Memory region: %d bytes at 0x%x", memorySize, baseAddr)
}

// TestResourceExhaustionBoundaries tests behavior at resource limits
func TestResourceExhaustionBoundaries(t *testing.T) {
	tests := []struct {
		name        string
		maxInstrs   int
		description string
	}{
		{"small_limit", 10, "Very small instruction limit"},
		{"medium_limit", 1000, "Medium instruction limit"},
		{"large_limit", 100000, "Large instruction limit"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := emulate.DefaultEngineConfig()
			config.MaxInstructions = test.maxInstrs
			engine := emulate.NewEngineWithConfig(config)

			testState := state.NewState()
			testState.SetSP(0)
			testState.SetX(0, 0)
			engine.SetState(testState)

			instrValue := uint32(0x91000400) // ADD X0, X0, #1

			t.Logf("Testing %s with limit of %d instructions", test.description, test.maxInstrs)

			// Execute instructions until limit is hit
			var i int
			var lastErr error
			for i = 0; i < test.maxInstrs+10; i++ { // Try to exceed limit
				pc := uint64(0x10000000 + i*4)
				err := engine.ExecuteInstruction(pc, instrValue)
				if err != nil {
					lastErr = err
					break
				}
			}

			if i >= test.maxInstrs {
				if lastErr == nil {
					t.Errorf("Expected instruction limit error but none occurred")
				} else {
					t.Logf("✅ Correctly hit instruction limit at %d instructions: %v", i, lastErr)
				}
			} else {
				t.Logf("✅ Executed %d instructions within limit of %d", i, test.maxInstrs)
			}

			// Check final state
			finalState := engine.GetState()
			expectedValue := uint64(i) // Each instruction adds 1
			if finalState.GetX(0) != expectedValue {
				t.Errorf("Final state incorrect: expected X0=%d, got X0=%d", expectedValue, finalState.GetX(0))
			}
		})
	}
}

// TestLongRunningExecution tests sustained execution over time
func TestLongRunningExecution(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long running test in short mode")
	}

	const testDurationSeconds = 5
	const checkInterval = time.Second

	t.Logf("Running sustained execution test for %d seconds", testDurationSeconds)

	config := emulate.DefaultEngineConfig()
	config.MaxInstructions = 10000000 // Large limit for long running test - enough for 5 seconds
	engine := emulate.NewEngineWithConfig(config)
	testState := state.NewState()
	testState.SetSP(0)
	testState.SetX(0, 0)
	engine.SetState(testState)

	instrValue := uint32(0x91000400) // ADD X0, X0, #1

	startTime := time.Now()
	lastCheck := startTime
	pc := uint64(0x10000000)
	instructionCount := 0

	for time.Since(startTime) < testDurationSeconds*time.Second {
		err := engine.ExecuteInstruction(pc, instrValue)
		if err != nil {
			t.Fatalf("Instruction failed after %d instructions: %v", instructionCount, err)
		}

		pc += 4
		instructionCount++

		// Periodic status check
		if time.Since(lastCheck) >= checkInterval {
			currentState := engine.GetState()
			rate := float64(instructionCount) / time.Since(startTime).Seconds()
			t.Logf("Status: %d instructions executed, %.2f instr/sec, X0=%d",
				instructionCount, rate, currentState.GetX(0))
			lastCheck = time.Now()
		}
	}

	totalTime := time.Since(startTime)
	finalState := engine.GetState()

	t.Logf("✅ Sustained execution completed:")
	t.Logf("   Total time: %v", totalTime)
	t.Logf("   Instructions executed: %d", instructionCount)
	t.Logf("   Average rate: %.2f instructions/second", float64(instructionCount)/totalTime.Seconds())
	t.Logf("   Final X0 value: %d", finalState.GetX(0))

	// Verify consistency
	if finalState.GetX(0) != uint64(instructionCount) {
		t.Errorf("State consistency error: expected X0=%d, got X0=%d", instructionCount, finalState.GetX(0))
	}
}

// TestMemoryPressure tests behavior under memory pressure
func TestMemoryPressure(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping memory pressure test in short mode")
	}

	// Force initial GC
	runtime.GC()
	var initialMem runtime.MemStats
	runtime.ReadMemStats(&initialMem)

	const engineCount = 100
	const instructionsPerEngine = 1000

	t.Logf("Creating %d engine instances with %d instructions each", engineCount, instructionsPerEngine)

	engines := make([]*emulate.Engine, engineCount)

	// Create many engine instances
	for i := 0; i < engineCount; i++ {
		engine := emulate.NewEngine()
		testState := state.NewState()
		testState.SetSP(0)
		testState.SetX(0, uint64(i))

		// Add some memory to each state
		testState.WriteMemory(0x10000000, make([]byte, 4096)) // 4KB per engine

		engine.SetState(testState)
		engines[i] = engine
	}

	// Execute instructions on all engines
	instrValue := uint32(0x91000400) // ADD X0, X0, #1

	for _, engine := range engines {
		for j := 0; j < instructionsPerEngine; j++ {
			pc := uint64(0x10000000 + j*4)
			err := engine.ExecuteInstruction(pc, instrValue)
			if err != nil {
				t.Fatalf("Execution failed under memory pressure: %v", err)
			}
		}
	}

	// Check memory usage
	runtime.GC()
	var finalMem runtime.MemStats
	runtime.ReadMemStats(&finalMem)

	memoryIncrease := finalMem.HeapInuse - initialMem.HeapInuse
	memoryPerEngine := memoryIncrease / uint64(engineCount)

	t.Logf("✅ Memory pressure test completed:")
	t.Logf("   Engines created: %d", engineCount)
	t.Logf("   Instructions per engine: %d", instructionsPerEngine)
	t.Logf("   Total instructions: %d", engineCount*instructionsPerEngine)
	t.Logf("   Memory increase: %d bytes (%.2f MB)", memoryIncrease, float64(memoryIncrease)/1024/1024)
	t.Logf("   Memory per engine: %d bytes", memoryPerEngine)

	// Verify all engines still work correctly
	for i, engine := range engines {
		state := engine.GetState()
		expected := uint64(i + instructionsPerEngine) // Initial value + increments
		if state.GetX(0) != expected {
			t.Errorf("Engine %d state corruption: expected X0=%d, got X0=%d", i, expected, state.GetX(0))
		}
	}

	t.Logf("✅ All engine states remained consistent under memory pressure")
}
