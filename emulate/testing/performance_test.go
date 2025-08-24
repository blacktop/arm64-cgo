//go:build darwin

package testing

import (
	"encoding/binary"
	"runtime"
	"testing"
	"time"

	"github.com/blacktop/arm64-cgo/emulate"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// PerformanceTest represents a performance test scenario
type PerformanceTest struct {
	name         string
	instructions [][]byte
	initialRegs  map[int]uint64
	iterations   int
	description  string
}

// BenchmarkSingleInstructionExecution benchmarks single instruction execution
func BenchmarkSingleInstructionExecution(b *testing.B) {
	// Test basic arithmetic instruction: ADD X0, X0, #1
	instrBytes := []byte{0x00, 0x04, 0x00, 0x91}
	instrValue := binary.LittleEndian.Uint32(instrBytes)

	// Create engine with higher limits
	config := emulate.DefaultEngineConfig()
	config.MaxInstructions = 1000000 // Large limit for benchmarks
	engine := emulate.NewEngineWithConfig(config)
	testState := state.NewState()
	testState.SetSP(0)
	engine.SetState(testState)

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := engine.ExecuteInstruction(0x10000000, instrValue)
		if err != nil {
			b.Fatalf("Instruction execution failed: %v", err)
		}
	}
}

// BenchmarkSequenceExecution benchmarks multi-instruction sequence execution
func BenchmarkSequenceExecution(b *testing.B) {
	// Arithmetic dependency chain sequence
	instructions := [][]byte{
		{0xe0, 0x93, 0x01, 0x91}, // ADD X0, XZR, #100
		{0x21, 0x40, 0x01, 0x91}, // ADD X1, X1, #80
		{0x02, 0x00, 0x01, 0x8b}, // ADD X2, X0, X1
		{0x63, 0x00, 0x02, 0x8b}, // ADD X3, X3, X2
	}

	initialRegs := map[int]uint64{
		1: 20,
		3: 5,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Create fresh engine for each iteration with higher limits
		config := emulate.DefaultEngineConfig()
		config.MaxInstructions = 1000000 // Large limit for benchmarks
		engine := emulate.NewEngineWithConfig(config)
		testState := state.NewState()
		testState.SetSP(0)

		// Set initial registers
		for reg, val := range initialRegs {
			testState.SetX(reg, val)
		}
		engine.SetState(testState)

		// Execute sequence
		pc := uint64(0x10000000)
		for _, instrBytes := range instructions {
			instrValue := binary.LittleEndian.Uint32(instrBytes)
			err := engine.ExecuteInstruction(pc, instrValue)
			if err != nil {
				b.Fatalf("Instruction execution failed: %v", err)
			}
			pc += 4
		}
	}
}

// BenchmarkMemoryOperations benchmarks memory store/load operations
func BenchmarkMemoryOperations(b *testing.B) {
	// Memory store/load sequence
	instructions := [][]byte{
		{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1  (X0 = 0x10000000)
		{0x21, 0x20, 0x00, 0x91}, // ADD X1, X1, #8  (X1 = 0x10000008)
		{0x02, 0x00, 0x00, 0xf9}, // STR X2, [X0]    (store at 0x10000000)
		{0x23, 0x00, 0x01, 0xf9}, // STR X3, [X1]    (store at 0x10000008)
		{0x04, 0x00, 0x40, 0xf9}, // LDR X4, [X0]    (load from 0x10000000)
		{0x25, 0x00, 0x41, 0xf9}, // LDR X5, [X1]    (load from 0x10000008)
	}

	initialRegs := map[int]uint64{
		0: 0x0,
		1: 0x10000000,
		2: 0xDEADBEEFCAFEBABE,
		3: 0x123456789ABCDEF0,
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		// Create fresh engine for each iteration with higher limits
		config := emulate.DefaultEngineConfig()
		config.MaxInstructions = 1000000 // Large limit for benchmarks
		engine := emulate.NewEngineWithConfig(config)
		testState := state.NewState()
		testState.SetSP(0)

		// Set up memory
		testState.WriteMemory(0x10000000, make([]byte, 0x1000))

		// Set initial registers
		for reg, val := range initialRegs {
			testState.SetX(reg, val)
		}
		engine.SetState(testState)

		// Execute sequence
		pc := uint64(0x10000000)
		for _, instrBytes := range instructions {
			instrValue := binary.LittleEndian.Uint32(instrBytes)
			err := engine.ExecuteInstruction(pc, instrValue)
			if err != nil {
				b.Fatalf("Instruction execution failed: %v", err)
			}
			pc += 4
		}
	}
}

// TestPerformanceComparison compares our emulator performance against hypervisor
func TestPerformanceComparison(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance comparison in short mode")
	}

	tests := []PerformanceTest{
		{
			name: "arithmetic_sequence",
			instructions: [][]byte{
				{0xe0, 0x93, 0x01, 0x91}, // ADD X0, XZR, #100
				{0x21, 0x40, 0x01, 0x91}, // ADD X1, X1, #80
				{0x02, 0x00, 0x01, 0x8b}, // ADD X2, X0, X1
				{0x63, 0x00, 0x02, 0x8b}, // ADD X3, X3, X2
			},
			initialRegs: map[int]uint64{1: 20, 3: 5},
			iterations:  100,
			description: "Arithmetic dependency chain",
		},
		{
			name: "memory_sequence",
			instructions: [][]byte{
				{0x00, 0x00, 0x01, 0x8b}, // ADD X0, X0, X1
				{0x21, 0x20, 0x00, 0x91}, // ADD X1, X1, #8
				{0x02, 0x00, 0x00, 0xf9}, // STR X2, [X0]
				{0x04, 0x00, 0x40, 0xf9}, // LDR X4, [X0]
			},
			initialRegs: map[int]uint64{0: 0, 1: 0x10000000, 2: 0xDEADBEEF},
			iterations:  50,
			description: "Memory store and load operations",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Benchmark our emulator
			ourTime := benchmarkOurEmulator(t, test)

			// Benchmark hypervisor (using hv execute)
			hvTime := benchmarkHypervisor(t, test)

			// Calculate performance ratio
			ratio := float64(ourTime) / float64(hvTime)

			t.Logf("Performance comparison for %s:", test.description)
			t.Logf("  Our emulator: %v (%v per operation)", ourTime, ourTime/time.Duration(test.iterations))
			t.Logf("  Hypervisor:   %v (%v per operation)", hvTime, hvTime/time.Duration(test.iterations))
			t.Logf("  Ratio: %.2fx (lower is better)", ratio)

			if ratio > 100 {
				t.Logf("⚠️  Our emulator is significantly slower than hypervisor")
			} else if ratio > 10 {
				t.Logf("⚠️  Our emulator is slower than hypervisor")
			} else if ratio < 0.5 {
				t.Logf("✅ Our emulator is faster than hypervisor!")
			} else {
				t.Logf("✅ Performance is reasonable compared to hypervisor")
			}
		})
	}
}

// benchmarkOurEmulator measures performance of our emulator
func benchmarkOurEmulator(t *testing.T, test PerformanceTest) time.Duration {
	start := time.Now()

	for i := 0; i < test.iterations; i++ {
		// Create engine with higher instruction limit
		config := emulate.DefaultEngineConfig()
		config.MaxInstructions = test.iterations * 10 // Ensure we don't hit limits
		engine := emulate.NewEngineWithConfig(config)
		testState := state.NewState()
		testState.SetSP(0)

		// Set up memory if needed
		if test.name == "memory_sequence" {
			testState.WriteMemory(0x10000000, make([]byte, 0x1000))
		}

		// Set initial registers
		for reg, val := range test.initialRegs {
			testState.SetX(reg, val)
		}
		engine.SetState(testState)

		// Execute sequence
		pc := uint64(0x10000000)
		for _, instrBytes := range test.instructions {
			instrValue := binary.LittleEndian.Uint32(instrBytes)
			err := engine.ExecuteInstruction(pc, instrValue)
			if err != nil {
				t.Fatalf("Our emulator failed: %v", err)
			}
			pc += 4
		}
	}

	return time.Since(start)
}

// benchmarkHypervisor measures performance of Apple's Hypervisor.framework
func benchmarkHypervisor(t *testing.T, test PerformanceTest) time.Duration {
	// Try to create hypervisor framework
	framework, err := NewHypervisorFramework()
	if err != nil {
		t.Skipf("Hypervisor framework not available: %v", err)
		return 0
	}
	defer framework.Close()

	start := time.Now()

	for i := 0; i < test.iterations; i++ {
		// Set up initial state
		err = framework.SetRegisterState(test.initialRegs)
		if err != nil {
			t.Fatalf("Failed to set register state: %v", err)
		}

		// Set up memory if needed
		if test.name == "memory_sequence" {
			err = framework.SetMemoryState(map[uint64][]byte{
				0x10000000: make([]byte, 0x1000),
			})
			if err != nil {
				t.Fatalf("Failed to set memory state: %v", err)
			}
		}

		// Execute sequence
		_, err = framework.ExecuteSequenceAndCompare(test.instructions, 0x40000000)
		if err != nil {
			t.Fatalf("Hypervisor execution failed: %v", err)
		}
	}

	return time.Since(start)
}

// TestMemoryEfficiency tests memory usage and garbage collection behavior
func TestMemoryEfficiency(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping memory efficiency test in short mode")
	}

	// Force GC to get baseline
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Test memory efficiency with a single engine (more realistic usage)
	config := emulate.DefaultEngineConfig()
	config.MaxInstructions = 100000
	engine := emulate.NewEngineWithConfig(config)
	testState := state.NewState()
	testState.SetSP(0)
	testState.SetX(0, 0)
	engine.SetState(testState)

	// Execute many instructions on the same engine
	iterations := 1000
	instrBytes := []byte{0x00, 0x04, 0x00, 0x91} // ADD X0, X0, #1
	instrValue := binary.LittleEndian.Uint32(instrBytes)

	for i := 0; i < iterations; i++ {
		for j := 0; j < 10; j++ {
			err := engine.ExecuteInstruction(0x10000000, instrValue)
			if err != nil {
				t.Fatalf("Instruction execution failed: %v", err)
			}
		}
	}

	// Force GC and measure
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate memory usage
	allocatedBytes := m2.TotalAlloc - m1.TotalAlloc
	bytesPerIteration := allocatedBytes / uint64(iterations)

	t.Logf("Memory efficiency test:")
	t.Logf("  Total iterations: %d", iterations)
	t.Logf("  Total allocated: %d bytes (%.2f MB)", allocatedBytes, float64(allocatedBytes)/1024/1024)
	t.Logf("  Per iteration: %d bytes", bytesPerIteration)
	t.Logf("  Final heap size: %d bytes (%.2f MB)", m2.HeapInuse, float64(m2.HeapInuse)/1024/1024)

	// Reasonable memory usage thresholds for single engine usage
	const maxBytesPerIteration = 1000 // 1KB per iteration for single engine
	if bytesPerIteration > maxBytesPerIteration {
		t.Logf("⚠️  Memory usage: %d bytes per iteration (max recommended: %d)", bytesPerIteration, maxBytesPerIteration)
	} else {
		t.Logf("✅ Memory usage is efficient")
	}
}
