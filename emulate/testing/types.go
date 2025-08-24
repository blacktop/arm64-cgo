package testing

import "github.com/blacktop/arm64-cgo/disassemble"

// EmulatorState represents the state of an ARM64 emulator at a point in time
type EmulatorState struct {
	// Our emulator state
	Registers map[int]uint64  // X0-X30 register values
	PC        uint64          // Program counter
	SP        uint64          // Stack pointer
	OurFlags  map[string]bool // NZCV flags from our emulator

	// Hypervisor state for comparison
	HypervisorRegisters map[int]uint64  // X0-X30 register values from Hypervisor
	HypervisorPC        uint64          // Program counter from Hypervisor
	HypervisorSP        uint64          // Stack pointer from Hypervisor
	Flags               map[string]bool // NZCV flags
}

// StateDifference represents a difference between our emulator and Hypervisor
type StateDifference struct {
	Type            string // "register", "flag", "memory"
	Name            string // Register name (e.g., "X0", "PC", "N")
	OurValue        any    // Value from our emulator
	HypervisorValue any    // Value from Hypervisor
}

// ComparisonResult represents the result of comparing a single instruction execution
type ComparisonResult struct {
	PC              uint64                   // Program counter where instruction was executed
	Instruction     *disassemble.Instruction // Decoded instruction
	InstrBytes      []byte                   // Raw instruction bytes
	BeforeState     *EmulatorState           // State before execution
	AfterState      *EmulatorState           // State after execution
	HypervisorError error                    // Error from Hypervisor execution (if any)
	OurError        error                    // Error from our emulator execution (if any)
	Differences     []StateDifference        // Differences found between emulators
}

// SequenceComparisonResult represents the result of comparing instruction sequence execution
type SequenceComparisonResult struct {
	StartPC         uint64            // Starting program counter
	Instructions    [][]byte          // Sequence of instruction bytes
	InitialState    *EmulatorState    // State before sequence execution
	FinalState      *EmulatorState    // State after sequence execution
	HypervisorError error             // Error from Hypervisor execution (if any)
	OurError        error             // Error from our emulator execution (if any)
	Differences     []StateDifference // Differences found between emulators
}

// TestCase represents a single test case for instruction validation
type TestCase struct {
	Name         string            // Test case name
	InstrBytes   []byte            // Instruction bytes to test
	PC           uint64            // Program counter to execute at
	InitialRegs  map[int]uint64    // Initial register state
	InitialMem   map[uint64][]byte // Initial memory state
	ExpectedRegs map[int]uint64    // Expected register state (optional)
	ExpectedMem  map[uint64][]byte // Expected memory state (optional)
	Description  string            // Test case description
}

// TestSuite represents a collection of related test cases
type TestSuite struct {
	Name        string     // Suite name
	Description string     // Suite description
	TestCases   []TestCase // Individual test cases
}

// ValidationResult represents the overall result of running a test suite
type ValidationResult struct {
	Suite       *TestSuite         // The test suite that was run
	TotalTests  int                // Total number of tests
	PassedTests int                // Number of tests that passed
	FailedTests int                // Number of tests that failed
	Results     []ComparisonResult // Individual test results
	Summary     string             // Summary of results
}

// MemoryComparison represents a comparison of memory states
type MemoryComparison struct {
	Address        uint64 // Memory address
	OurData        []byte // Data from our emulator
	HypervisorData []byte // Data from Hypervisor
	Differences    []int  // Byte indices where data differs
}

// InstructionStats represents statistics about instruction execution
type InstructionStats struct {
	Mnemonic       string   // Instruction mnemonic
	TotalTests     int      // Total tests for this instruction
	PassedTests    int      // Passed tests
	FailedTests    int      // Failed tests
	SuccessRate    float64  // Success rate percentage
	CommonFailures []string // Common failure patterns
}

// BenchmarkResult represents performance comparison results
type BenchmarkResult struct {
	InstructionCount      int     // Number of instructions executed
	OurTime               int64   // Execution time in nanoseconds (our emulator)
	HypervisorTime        int64   // Execution time in nanoseconds (Hypervisor)
	OurMemoryUsage        int64   // Memory usage in bytes (our emulator)
	HypervisorMemoryUsage int64   // Memory usage in bytes (Hypervisor)
	Speedup               float64 // Performance ratio (positive = we're faster)
}

// ErrorAnalysis represents analysis of execution errors
type ErrorAnalysis struct {
	InstructionBytes []byte // The instruction that caused the error
	PC               uint64 // Program counter
	OurError         error  // Error from our emulator
	HypervisorError  error  // Error from Hypervisor
	ErrorType        string // Classification of error type
	Severity         string // Error severity (low, medium, high, critical)
	Recommendation   string // Suggested fix or investigation
}
