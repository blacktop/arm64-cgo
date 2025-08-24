//go:build darwin

package testing

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"time"

	"github.com/blacktop/arm64-cgo/emulate"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// CPUState represents the complete ARM64 CPU state
type CPUState struct {
	X0   uint64 `json:"x0"`
	X1   uint64 `json:"x1"`
	X2   uint64 `json:"x2"`
	X3   uint64 `json:"x3"`
	X4   uint64 `json:"x4"`
	X5   uint64 `json:"x5"`
	X6   uint64 `json:"x6"`
	X7   uint64 `json:"x7"`
	X8   uint64 `json:"x8"`
	X9   uint64 `json:"x9"`
	X10  uint64 `json:"x10"`
	X11  uint64 `json:"x11"`
	X12  uint64 `json:"x12"`
	X13  uint64 `json:"x13"`
	X14  uint64 `json:"x14"`
	X15  uint64 `json:"x15"`
	X16  uint64 `json:"x16"`
	X17  uint64 `json:"x17"`
	X18  uint64 `json:"x18"`
	X19  uint64 `json:"x19"`
	X20  uint64 `json:"x20"`
	X21  uint64 `json:"x21"`
	X22  uint64 `json:"x22"`
	X23  uint64 `json:"x23"`
	X24  uint64 `json:"x24"`
	X25  uint64 `json:"x25"`
	X26  uint64 `json:"x26"`
	X27  uint64 `json:"x27"`
	X28  uint64 `json:"x28"`
	FP   uint64 `json:"fp"`
	LR   uint64 `json:"lr"`
	SP   uint64 `json:"sp"`
	PC   uint64 `json:"pc"`
	CPSR uint32 `json:"cpsr"`
}

// HypervisorResponse represents the JSON response from hv execute
type HypervisorResponse struct {
	State    CPUState               `json:"state"`
	ExitInfo map[string]interface{} `json:"exit_info"`
	Memory   map[string]string      `json:"memory"`
}

// ExecuteResult represents the result of instruction execution
type ExecuteResult struct {
	FinalState     CPUState          `json:"final_state"`
	MemoryRegions  map[uint64][]byte `json:"memory_regions"`
	ExitInfo       string            `json:"exit_info"`
	ExecutionError string            `json:"execution_error,omitempty"`
}

// Helper functions for register mapping using reflection

// setRegistersFromMap sets CPUState registers from a map using reflection
func setRegistersFromMap(cpuState *CPUState, initialRegs map[int]uint64) {
	v := reflect.ValueOf(cpuState).Elem()
	t := v.Type()

	// Create a map of register names to their indices for fast lookup
	regFieldMap := make(map[int]int)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// Map X0-X28 fields to register numbers
		switch field.Name {
		case "X0":
			regFieldMap[0] = i
		case "X1":
			regFieldMap[1] = i
		case "X2":
			regFieldMap[2] = i
		case "X3":
			regFieldMap[3] = i
		case "X4":
			regFieldMap[4] = i
		case "X5":
			regFieldMap[5] = i
		case "X6":
			regFieldMap[6] = i
		case "X7":
			regFieldMap[7] = i
		case "X8":
			regFieldMap[8] = i
		case "X9":
			regFieldMap[9] = i
		case "X10":
			regFieldMap[10] = i
		case "X11":
			regFieldMap[11] = i
		case "X12":
			regFieldMap[12] = i
		case "X13":
			regFieldMap[13] = i
		case "X14":
			regFieldMap[14] = i
		case "X15":
			regFieldMap[15] = i
		case "X16":
			regFieldMap[16] = i
		case "X17":
			regFieldMap[17] = i
		case "X18":
			regFieldMap[18] = i
		case "X19":
			regFieldMap[19] = i
		case "X20":
			regFieldMap[20] = i
		case "X21":
			regFieldMap[21] = i
		case "X22":
			regFieldMap[22] = i
		case "X23":
			regFieldMap[23] = i
		case "X24":
			regFieldMap[24] = i
		case "X25":
			regFieldMap[25] = i
		case "X26":
			regFieldMap[26] = i
		case "X27":
			regFieldMap[27] = i
		case "X28":
			regFieldMap[28] = i
		case "FP":
			regFieldMap[29] = i // X29 is Frame Pointer
		case "LR":
			regFieldMap[30] = i // X30 is Link Register
		}
	}

	// Set register values from the map
	for regNum, value := range initialRegs {
		if fieldIndex, exists := regFieldMap[regNum]; exists {
			field := v.Field(fieldIndex)
			if field.CanSet() && field.Kind() == reflect.Uint64 {
				field.SetUint(value)
			}
		}
		// Handle SP specially (register 31)
		if regNum == 31 {
			if spField := v.FieldByName("SP"); spField.IsValid() && spField.CanSet() {
				spField.SetUint(value)
			}
		}
	}
}

// extractRegistersToMap extracts CPUState registers to a map using reflection
func extractRegistersToMap(cpuState *CPUState) map[int]uint64 {
	registers := make(map[int]uint64)
	v := reflect.ValueOf(cpuState).Elem()

	// Extract register values using field names
	fieldMap := map[string]int{
		"X0": 0, "X1": 1, "X2": 2, "X3": 3, "X4": 4, "X5": 5, "X6": 6, "X7": 7,
		"X8": 8, "X9": 9, "X10": 10, "X11": 11, "X12": 12, "X13": 13, "X14": 14, "X15": 15,
		"X16": 16, "X17": 17, "X18": 18, "X19": 19, "X20": 20, "X21": 21, "X22": 22, "X23": 23,
		"X24": 24, "X25": 25, "X26": 26, "X27": 27, "X28": 28, "FP": 29, "LR": 30,
	}

	for fieldName, regNum := range fieldMap {
		if field := v.FieldByName(fieldName); field.IsValid() && field.Kind() == reflect.Uint64 {
			registers[regNum] = field.Uint()
		}
	}

	// Handle SP (register 31)
	if spField := v.FieldByName("SP"); spField.IsValid() && spField.Kind() == reflect.Uint64 {
		registers[31] = spField.Uint()
	}

	return registers
}

// HypervisorFramework provides utilities for comparing our ARM64 emulator with Apple's Hypervisor.framework
type HypervisorFramework struct {
	ourEngine    *emulate.Engine
	hvBinaryPath string
	timeout      time.Duration
	baseAddr     uint64
	stackAddr    uint64
	stackSize    uint64
	initialMem   map[uint64][]byte // Store initial memory state for hypervisor
}

// NewHypervisorFramework creates a new test framework instance
func NewHypervisorFramework() (*HypervisorFramework, error) {
	// Find the hv binary in PATH
	hvPath, err := exec.LookPath("hv")
	if err != nil {
		return nil, fmt.Errorf("hv binary not found in PATH: %w", err)
	}

	// Create our engine
	ourEngine := emulate.NewEngine()

	// Set up default memory layout
	baseAddr := uint64(0x10000000)
	stackAddr := uint64(0x20000000)
	stackSize := uint64(0x200000) // 2MB stack

	return &HypervisorFramework{
		ourEngine:    ourEngine,
		hvBinaryPath: hvPath,
		timeout:      30 * time.Second,
		baseAddr:     baseAddr,
		stackAddr:    stackAddr,
		stackSize:    stackSize,
		initialMem:   make(map[uint64][]byte),
	}, nil
}

// SetRegisterState sets initial register values in both emulators
func (hf *HypervisorFramework) SetRegisterState(registers map[int]uint64) error {
	// Set up our engine state
	testState := state.NewState()
	for reg, val := range registers {
		if reg == 31 {
			testState.SetSP(val) // Register 31 is SP
		} else {
			testState.SetX(reg, val)
		}
	}

	// Set default stack pointer if not provided (match hypervisor's initial SP=0)
	if _, hasStack := registers[31]; !hasStack {
		testState.SetSP(0) // Hypervisor starts with SP=0
	}

	hf.ourEngine.SetState(testState)
	return nil
}

// SetMemoryState sets initial memory state
func (hf *HypervisorFramework) SetMemoryState(memory map[uint64][]byte) error {
	// Store memory state for hypervisor execution
	hf.initialMem = make(map[uint64][]byte)
	for addr, data := range memory {
		// Make a copy of the data to avoid aliasing
		dataCopy := make([]byte, len(data))
		copy(dataCopy, data)
		hf.initialMem[addr] = dataCopy
	}

	// Set memory in our emulator
	if hf.ourEngine != nil {
		for addr, data := range memory {
			// Map the region first
			if armState, ok := hf.ourEngine.GetState().(*state.ARM64State); ok {
				armState.MapRegion(addr, len(data))
			}
			// Write the data
			hf.ourEngine.GetState().WriteMemory(addr, data)
		}
	}
	return nil
}

// executeWithHypervisor executes instructions using the hv binary
func (hf *HypervisorFramework) executeWithHypervisor(instructions [][]byte, pc uint64, initialRegs map[int]uint64, initialMem map[uint64][]byte) (*ExecuteResult, error) {
	// Create temporary files
	tempDir, err := os.MkdirTemp("", "hypervisor_test_*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp dir: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Create binary file with all instructions
	codeFile := filepath.Join(tempDir, "code.bin")
	var codeData []byte
	for _, instrBytes := range instructions {
		codeData = append(codeData, instrBytes...)
	}

	if err := os.WriteFile(codeFile, codeData, 0644); err != nil {
		return nil, fmt.Errorf("failed to write code file: %w", err)
	}

	// Create initial CPU state file if we have initial registers
	var stateFile string
	if len(initialRegs) > 0 {
		stateFile = filepath.Join(tempDir, "state.json")

		// Convert our register map to the expected CPU state format
		cpuState := CPUState{
			PC: pc,
			SP: 0, // Default SP
		}

		// Set registers from initialRegs using helper function
		setRegistersFromMap(&cpuState, initialRegs)

		stateData, err := json.Marshal(cpuState)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal state: %w", err)
		}

		if err := os.WriteFile(stateFile, stateData, 0644); err != nil {
			return nil, fmt.Errorf("failed to write state file: %w", err)
		}
	}

	// Execute hv binary with correct arguments
	// Allocate sufficient memory for tests with large offsets
	memSize := len(codeData) + 64*1024
	page := os.Getpagesize()
	if memSize%page != 0 {
		memSize = ((memSize / page) + 1) * page
	}
	if memSize < 256*1024 {
		memSize = 256 * 1024
	}

	args := []string{"execute", codeFile, "--base-addr", fmt.Sprintf("%d", pc), "--mem-size", fmt.Sprintf("%d", memSize)}
	if stateFile != "" {
		args = append(args, "--state", stateFile)
	}

	cmd := exec.Command(hf.hvBinaryPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("hv execution failed: %w, output: %s", err, string(output))
	}

	// Parse the JSON output from hv execute
	var hvResponse HypervisorResponse
	if err := json.Unmarshal(output, &hvResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal hv output: %w, output: %s", err, string(output))
	}

	// Convert to our ExecuteResult format
	result := &ExecuteResult{
		FinalState: hvResponse.State,
		ExitInfo:   "success",
	}

	return result, nil
}

// ExecuteSequenceAndCompare executes a sequence of instructions and compares results
func (hf *HypervisorFramework) ExecuteSequenceAndCompare(instructions [][]byte, pc uint64) (*ComparisonResult, error) {
	// Get initial state
	initialState := hf.ourEngine.GetState()
	initialRegs := make(map[int]uint64)
	for i := range 32 {
		if i == 31 {
			initialRegs[i] = initialState.GetSP()
		} else {
			initialRegs[i] = initialState.GetX(i)
		}
	}

	// Mirror Hypervisor's code visibility: place instruction bytes into our memory at PC
	if st, ok := hf.ourEngine.GetState().(*state.ARM64State); ok {
		var codeData []byte
		for _, instrBytes := range instructions {
			codeData = append(codeData, instrBytes...)
		}
		if len(codeData) > 0 {
			st.MapRegion(pc, len(codeData))
			st.WriteMemory(pc, codeData)
		}
	}

	// Execute on our emulator
	currentPC := pc
	for _, instrBytes := range instructions {
		if len(instrBytes) != 4 {
			return nil, fmt.Errorf("invalid instruction length: %d", len(instrBytes))
		}

		instrValue := binary.LittleEndian.Uint32(instrBytes)
		if err := hf.ourEngine.ExecuteInstruction(currentPC, instrValue); err != nil {
			return nil, fmt.Errorf("our emulator failed: %w", err)
		}
		currentPC += 4
	}

	// Get our final state
	ourFinalState := hf.ourEngine.GetState()
	ourRegisters := make(map[int]uint64)
	for i := 0; i < 32; i++ {
		if i == 31 {
			ourRegisters[i] = ourFinalState.GetSP()
		} else {
			ourRegisters[i] = ourFinalState.GetX(i)
		}
	}

	// Execute on hypervisor
	hvResult, err := hf.executeWithHypervisor(instructions, pc, initialRegs, hf.initialMem)
	if err != nil {
		return nil, fmt.Errorf("hypervisor execution failed: %w", err)
	}

	// Convert hypervisor state to our format using helper function
	hvRegisters := extractRegistersToMap(&hvResult.FinalState)

	// Compare states and find differences
	var differences []StateDifference
	for i := range 32 {
		if ourRegisters[i] != hvRegisters[i] {
			regName := fmt.Sprintf("X%d", i)
			switch i {
			case 29:
				regName = "FP"
			case 30:
				regName = "LR"
			case 31:
				regName = "SP"
			}

			differences = append(differences, StateDifference{
				Type:            "register",
				Name:            regName,
				OurValue:        ourRegisters[i],
				HypervisorValue: hvRegisters[i],
			})
		}
	}

	// Check PC difference
	ourPC := ourFinalState.GetPC()
	if ourPC != hvResult.FinalState.PC {
		differences = append(differences, StateDifference{
			Type:            "register",
			Name:            "PC",
			OurValue:        ourPC,
			HypervisorValue: hvResult.FinalState.PC,
		})
	}

	return &ComparisonResult{
		Differences: differences,
		AfterState: &EmulatorState{
			Registers:           ourRegisters,
			PC:                  ourPC,
			HypervisorRegisters: hvRegisters,
			HypervisorPC:        hvResult.FinalState.PC,
		},
	}, nil
}

// GetOurEngine returns the internal ARM64 engine for compatibility
func (hf *HypervisorFramework) GetOurEngine() *emulate.Engine {
	return hf.ourEngine
}

// Close cleans up resources
// Reset clears all state for a fresh test
func (hf *HypervisorFramework) Reset() error {
	// Create a completely new engine with fresh state
	hf.ourEngine = emulate.NewEngine()

	// Clear stored memory state
	hf.initialMem = make(map[uint64][]byte)

	return nil
}

func (hf *HypervisorFramework) Close() {
	// Nothing to clean up for now
}
