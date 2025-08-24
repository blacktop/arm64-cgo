package testing

import (
	"encoding/binary"
	"fmt"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/blacktop/arm64-cgo/emulate"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// Common test utilities to eliminate code duplication across test files

// EngineConfig provides configuration options for test engines
type EngineConfig struct {
	MaxInstructions int
	InitialSP      uint64
	InitialRegs    map[int]uint64
}

// DefaultTestEngineConfig returns a standard test engine configuration
func DefaultTestEngineConfig() *EngineConfig {
	return &EngineConfig{
		MaxInstructions: 1000000,
		InitialSP:      0,
		InitialRegs:    make(map[int]uint64),
	}
}

// CreateTestEngine creates a configured emulator engine for testing
func CreateTestEngine(config *EngineConfig) *emulate.Engine {
	if config == nil {
		config = DefaultTestEngineConfig()
	}

	emulatorConfig := emulate.DefaultEngineConfig()
	emulatorConfig.MaxInstructions = config.MaxInstructions
	engine := emulate.NewEngineWithConfig(emulatorConfig)
	
	testState := state.NewState()
	testState.SetSP(config.InitialSP)
	
	// Set initial register values
	for reg, value := range config.InitialRegs {
		if reg == 31 {
			// Special case: register 31 in ARM64 context-sensitive (SP vs XZR)
			// For tests, when initializing register 31, set both SP and allow X31 writes
			testState.SetSP(value)
		}
		testState.SetX(reg, value)
	}
	
	engine.SetState(testState)
	return engine
}

// CreateBenchmarkEngine creates an optimized engine for benchmark tests
func CreateBenchmarkEngine(maxInstructions int) *emulate.Engine {
	config := &EngineConfig{
		MaxInstructions: maxInstructions,
		InitialSP:      0,
		InitialRegs:    make(map[int]uint64),
	}
	return CreateTestEngine(config)
}

// CommonInstructions provides frequently used instruction encodings
var CommonInstructions = struct {
	AddX0Imm1    []byte // ADD X0, X0, #1
	AddX1X0      []byte // ADD X1, X1, X0
	SubX0Imm1    []byte // SUB X0, X0, #1
	AddX2X1X0    []byte // ADD X2, X1, X0
	MovX0Imm100  []byte // ADD X0, XZR, #100
	MovX1Imm200  []byte // ADD X1, XZR, #200
	LdrX0FromX0  []byte // LDR X0, [X0]
	StrX0ToX1    []byte // STR X0, [X1]
	Nop          []byte // NOP
}{
	AddX0Imm1:   []byte{0x00, 0x04, 0x00, 0x91},
	AddX1X0:     []byte{0x21, 0x00, 0x00, 0x8b},
	SubX0Imm1:   []byte{0x00, 0x04, 0x00, 0xd1},
	AddX2X1X0:   []byte{0x22, 0x00, 0x00, 0x8b},
	MovX0Imm100: []byte{0x00, 0x19, 0x80, 0x91},
	MovX1Imm200: []byte{0x01, 0x32, 0x80, 0x91},
	LdrX0FromX0: []byte{0x00, 0x00, 0x40, 0xf9},
	StrX0ToX1:   []byte{0x20, 0x00, 0x00, 0xf9},
	Nop:         []byte{0x1f, 0x20, 0x03, 0xd5},
}

// InstructionSequenceBuilder helps build instruction sequences for testing
type InstructionSequenceBuilder struct {
	instructions [][]byte
	description  string
}

// NewSequenceBuilder creates a new instruction sequence builder
func NewSequenceBuilder() *InstructionSequenceBuilder {
	return &InstructionSequenceBuilder{
		instructions: make([][]byte, 0),
	}
}

// Add adds an instruction to the sequence
func (b *InstructionSequenceBuilder) Add(instruction []byte) *InstructionSequenceBuilder {
	b.instructions = append(b.instructions, instruction)
	return b
}

// AddMultiple adds multiple instructions to the sequence
func (b *InstructionSequenceBuilder) AddMultiple(instructions ...[]byte) *InstructionSequenceBuilder {
	b.instructions = append(b.instructions, instructions...)
	return b
}

// Build returns the final instruction sequence
func (b *InstructionSequenceBuilder) Build() [][]byte {
	return b.instructions
}

// WithDescription sets a description for the sequence
func (b *InstructionSequenceBuilder) WithDescription(desc string) *InstructionSequenceBuilder {
	b.description = desc
	return b
}

// PredefinedSequences provides common instruction sequences
func PredefinedSequences() map[string][][]byte {
	return map[string][][]byte{
		"arithmetic_chain": NewSequenceBuilder().
			Add(CommonInstructions.MovX0Imm100).
			Add(CommonInstructions.MovX1Imm200).
			Add(CommonInstructions.AddX2X1X0).
			Build(),
		"counter_loop": NewSequenceBuilder().
			AddMultiple(
				CommonInstructions.AddX0Imm1,
				CommonInstructions.AddX1X0,
				CommonInstructions.SubX0Imm1,
				CommonInstructions.AddX2X1X0,
			).
			Build(),
		"memory_operations": NewSequenceBuilder().
			Add(CommonInstructions.MovX0Imm100).
			Add(CommonInstructions.StrX0ToX1).
			Add(CommonInstructions.LdrX0FromX0).
			Build(),
	}
}

// StateComparison provides utilities for comparing engine states
type StateComparison struct {
	OurState       *state.ARM64State
	ExpectedState  *state.ARM64State
	Differences    []StateDifference
}

// CompareStates compares two emulator states and returns differences
func CompareStates(our, expected *state.ARM64State) *StateComparison {
	comparison := &StateComparison{
		OurState:      our,
		ExpectedState: expected,
		Differences:   make([]StateDifference, 0),
	}

	// Compare PC (Program Counter)
	if our.GetPC() != expected.GetPC() {
		comparison.Differences = append(comparison.Differences, StateDifference{
			Type:            "register",
			Name:            "PC",
			OurValue:        fmt.Sprintf("0x%x", our.GetPC()),
			HypervisorValue: fmt.Sprintf("0x%x", expected.GetPC()),
		})
	}

	// Compare SP (Stack Pointer)
	if our.GetSP() != expected.GetSP() {
		comparison.Differences = append(comparison.Differences, StateDifference{
			Type:            "register", 
			Name:            "SP",
			OurValue:        fmt.Sprintf("0x%x", our.GetSP()),
			HypervisorValue: fmt.Sprintf("0x%x", expected.GetSP()),
		})
	}

	// Compare general purpose registers X0-X30
	for i := 0; i < 31; i++ {
		ourReg := our.GetX(i)
		expectedReg := expected.GetX(i)
		if ourReg != expectedReg {
			regName := fmt.Sprintf("X%d", i)
			// Use standard names for special registers
			switch i {
			case 29:
				regName = "FP" // Frame Pointer
			case 30:
				regName = "LR" // Link Register
			}
			
			comparison.Differences = append(comparison.Differences, StateDifference{
				Type:            "register",
				Name:            regName,
				OurValue:        fmt.Sprintf("0x%x", ourReg),
				HypervisorValue: fmt.Sprintf("0x%x", expectedReg),
			})
		}
	}

	// Compare ARM64 condition flags (NZCV)
	flags := []struct {
		name     string
		ourFlag  bool
		expFlag  bool
	}{
		{"N", our.GetN(), expected.GetN()},
		{"Z", our.GetZ(), expected.GetZ()},
		{"C", our.GetC(), expected.GetC()},
		{"V", our.GetV(), expected.GetV()},
	}

	for _, flag := range flags {
		if flag.ourFlag != flag.expFlag {
			comparison.Differences = append(comparison.Differences, StateDifference{
				Type:            "flag",
				Name:            flag.name,
				OurValue:        flag.ourFlag,
				HypervisorValue: flag.expFlag,
			})
		}
	}

	return comparison
}

// HasDifferences returns true if there are any state differences
func (sc *StateComparison) HasDifferences() bool {
	return len(sc.Differences) > 0
}

// String returns a formatted string of all differences
func (sc *StateComparison) String() string {
	if !sc.HasDifferences() {
		return "No differences found"
	}

	result := fmt.Sprintf("Found %d differences:\n", len(sc.Differences))
	for _, diff := range sc.Differences {
		result += fmt.Sprintf("  %s %s: our=%v, expected=%v\n", 
			diff.Type, diff.Name, diff.OurValue, diff.HypervisorValue)
	}
	return result
}

// PerformanceHelper provides utilities for performance testing
type PerformanceHelper struct {
	startTime time.Time
	endTime   time.Time
}

// NewPerformanceHelper creates a new performance measurement helper
func NewPerformanceHelper() *PerformanceHelper {
	return &PerformanceHelper{}
}

// Start begins performance measurement
func (ph *PerformanceHelper) Start() {
	runtime.GC() // Clean up before measurement
	ph.startTime = time.Now()
}

// Stop ends performance measurement
func (ph *PerformanceHelper) Stop() time.Duration {
	ph.endTime = time.Now()
	return ph.endTime.Sub(ph.startTime)
}

// GetRate calculates operations per second
func (ph *PerformanceHelper) GetRate(operations int) float64 {
	duration := ph.endTime.Sub(ph.startTime)
	return float64(operations) / duration.Seconds()
}

// MemoryHelper provides utilities for memory testing
type MemoryHelper struct {
	initialMem runtime.MemStats
	finalMem   runtime.MemStats
}

// NewMemoryHelper creates a new memory measurement helper
func NewMemoryHelper() *MemoryHelper {
	return &MemoryHelper{}
}

// Start begins memory measurement
func (mh *MemoryHelper) Start() {
	runtime.GC() // Clean up before measurement
	runtime.ReadMemStats(&mh.initialMem)
}

// Stop ends memory measurement
func (mh *MemoryHelper) Stop() {
	runtime.ReadMemStats(&mh.finalMem)
}

// GetAllocatedBytes returns bytes allocated during the measurement
func (mh *MemoryHelper) GetAllocatedBytes() uint64 {
	return mh.finalMem.TotalAlloc - mh.initialMem.TotalAlloc
}

// GetHeapBytes returns current heap size
func (mh *MemoryHelper) GetHeapBytes() uint64 {
	return mh.finalMem.HeapInuse
}

// ErrorHelper provides utilities for error testing
type ErrorHelper struct{}

// NewErrorHelper creates a new error testing helper
func NewErrorHelper() *ErrorHelper {
	return &ErrorHelper{}
}

// ContainsError checks if an error contains expected text (more robust than direct string matching)
func (eh *ErrorHelper) ContainsError(err error, expectedText string) bool {
	if err == nil {
		return expectedText == ""
	}
	if expectedText == "" {
		return false
	}
	
	errStr := err.Error()
	// Simple case-insensitive substring matching
	// Could be extended with regex or more sophisticated matching
	return containsIgnoreCase(errStr, expectedText)
}

// MatchesErrorPattern checks if error matches expected pattern
func (eh *ErrorHelper) MatchesErrorPattern(err error, patterns ...string) bool {
	if err == nil {
		return len(patterns) == 0
	}
	
	errStr := err.Error()
	for _, pattern := range patterns {
		if containsIgnoreCase(errStr, pattern) {
			return true
		}
	}
	return false
}

// TestHelper combines all helper utilities
type TestHelper struct {
	Performance *PerformanceHelper
	Memory      *MemoryHelper
	Error       *ErrorHelper
}

// NewTestHelper creates a comprehensive test helper
func NewTestHelper() *TestHelper {
	return &TestHelper{
		Performance: NewPerformanceHelper(),
		Memory:      NewMemoryHelper(),
		Error:       NewErrorHelper(),
	}
}

// ExecuteInstructionSequence executes a sequence of instructions and measures performance
func (th *TestHelper) ExecuteInstructionSequence(t *testing.T, engine *emulate.Engine, instructions [][]byte, startPC uint64) time.Duration {
	th.Performance.Start()
	
	pc := startPC
	for i, instrBytes := range instructions {
		instrValue := binary.LittleEndian.Uint32(instrBytes)
		err := engine.ExecuteInstruction(pc, instrValue)
		if err != nil {
			t.Fatalf("Instruction %d failed at PC 0x%x: %v", i, pc, err)
		}
		pc += 4
	}
	
	return th.Performance.Stop()
}

// Utility functions

// containsIgnoreCase performs case-insensitive substring search
func containsIgnoreCase(s, substr string) bool {
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}