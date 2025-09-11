package emulate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
	"github.com/blacktop/arm64-cgo/emulate/instructions"
	"github.com/blacktop/arm64-cgo/emulate/state"
)

// Engine provides the primary ARM64 emulation interface using the modular instruction system
type Engine struct {
	state                   core.State
	registry                core.InstructionRegistry
	maxInstructions         int
	instrCount              int
	enableTrace             bool
	trace                   []core.InstructionInfo
	StopOnError             bool
	LastError               error
	StopAddress             uint64
	ShouldHaltPreHandler    func(state core.State, info core.InstructionInfo) bool
	ShouldHaltPostHandler   func(state core.State, info core.InstructionInfo) bool
	ShouldTakeBranchHandler func(state core.State, info core.InstructionInfo) bool
}

// EngineConfig provides configuration options for the ARM64 emulator engine
type EngineConfig struct {
	MaxInstructions         int
	EnableTrace             bool
	StopOnError             bool
	InitialPC               uint64
	InitialSP               uint64
	StackSize               int
	StackBase               uint64
	MemoryHandler           core.MemoryReadHandler
	StringHandler           core.StringPoolHandler
	PointerResolver         core.PointerResolver
	StopAddress             uint64
	ShouldHaltPreHandler    func(state core.State, info core.InstructionInfo) bool
	ShouldHaltPostHandler   func(state core.State, info core.InstructionInfo) bool
	ShouldTakeBranchHandler func(state core.State, info core.InstructionInfo) bool
}

// DefaultEngineConfig returns a configuration with sensible defaults
func DefaultEngineConfig() *EngineConfig {
	return &EngineConfig{
		MaxInstructions: 10000,
		EnableTrace:     false,
		StopOnError:     true,
		InitialPC:       0,
		InitialSP:       0x7fff00000000, // Top of stack
		StackSize:       1024 * 1024,    // 1MB stack
		StackBase:       0x7fff00000000 - 1024*1024,
	}
}

// NewEngine creates a new ARM64 emulation engine with default configuration
func NewEngine() *Engine {
	config := DefaultEngineConfig()
	return NewEngineWithConfig(config)
}

// NewEngineWithConfig creates a new ARM64 emulation engine with the specified configuration
func NewEngineWithConfig(config *EngineConfig) *Engine {
	st := state.NewState()

	// Apply configuration to state
	if config.InitialPC != 0 {
		st.SetPC(config.InitialPC)
	}
	if config.InitialSP != 0 {
		st.SetSP(config.InitialSP)
	}

	// Set up memory handlers if provided
	if config.MemoryHandler != nil {
		st.SetMemoryReadHandler(config.MemoryHandler)
	}
	if config.StringHandler != nil {
		st.SetStringPoolHandler(config.StringHandler)
	}
	if config.PointerResolver != nil {
		st.SetPointerResolver(config.PointerResolver)
	}

	return &Engine{
		state:                   st,
		registry:                instructions.DefaultRegistry(),
		maxInstructions:         config.MaxInstructions,
		instrCount:              0,
		enableTrace:             config.EnableTrace,
		trace:                   make([]core.InstructionInfo, 0),
		StopOnError:             config.StopOnError,
		StopAddress:             config.StopAddress,
		ShouldHaltPreHandler:    config.ShouldHaltPreHandler,
		ShouldHaltPostHandler:   config.ShouldHaltPostHandler,
		ShouldTakeBranchHandler: config.ShouldTakeBranchHandler,
	}
}

// NewEngineWithState creates a new engine with a custom initial state
func NewEngineWithState(st core.State) *Engine {
	return &Engine{
		state:           st,
		registry:        instructions.DefaultRegistry(),
		maxInstructions: 10000,
		instrCount:      0,
		enableTrace:     false,
		trace:           make([]core.InstructionInfo, 0),
		StopOnError:     true,
	}
}

// LoadInstructions loads a blob of ARM64 machine code into memory at the specified address
func (e *Engine) LoadInstructions(addr uint64, data []byte) error {
	// Validate alignment
	if addr%4 != 0 {
		return fmt.Errorf("instruction address 0x%x is not 4-byte aligned", addr)
	}
	if len(data)%4 != 0 {
		return fmt.Errorf("instruction data length %d is not 4-byte aligned", len(data))
	}

	// Validate instructions during loading
	for i := 0; i < len(data); i += 4 {
		instrAddr := addr + uint64(i)
		instrValue := uint32(data[i]) | uint32(data[i+1])<<8 | uint32(data[i+2])<<16 | uint32(data[i+3])<<24

		// Try to decode the instruction to validate it
		var results [1024]byte
		_, err := disassemble.Decompose(instrAddr, instrValue, &results)
		if err != nil {
			return fmt.Errorf("invalid instruction at offset %d (0x%x): %w", i, instrAddr, err)
		}
	}

	// Map the memory region first, then load the validated instructions
	if armState, ok := e.state.(*state.ARM64State); ok {
		armState.MapRegion(addr, len(data))
	}
	e.state.WriteMemory(addr, data)
	return nil
}

// LoadInstructionsAt is a convenience method that loads instructions and sets PC
func (e *Engine) LoadInstructionsAt(addr uint64, data []byte) error {
	if err := e.LoadInstructions(addr, data); err != nil {
		return err
	}
	e.state.SetPC(addr)
	return nil
}

// SetMaxInstructions sets the maximum number of instructions to execute
func (e *Engine) SetMaxInstructions(limit int) {
	e.maxInstructions = limit
}

// GetInstructionCount returns the current instruction count
func (e *Engine) GetInstructionCount() int {
	return e.instrCount
}

// Reset resets the engine state
func (e *Engine) Reset() {
	e.state = state.NewState()
	e.instrCount = 0
	e.trace = e.trace[:0]
	e.LastError = nil
}

// GetState returns the current emulator state
func (e *Engine) GetState() core.State {
	return e.state
}

// SetState sets a new emulator state
func (e *Engine) SetState(st core.State) {
	e.state = st
}

// SetTrace enables or disables instruction tracing
func (e *Engine) SetTrace(enable bool) {
	e.enableTrace = enable
	if !enable {
		e.trace = e.trace[:0]
	}
}

// GetTrace returns the instruction trace
func (e *Engine) GetTrace() []core.InstructionInfo {
	return e.trace
}

// ExecuteInstruction executes a single instruction using the modular instruction system
func (e *Engine) ExecuteInstruction(pc uint64, instr uint32) error {
	if e.instrCount >= e.maxInstructions {
		return fmt.Errorf("maximum instruction limit reached (%d)", e.maxInstructions)
	}

	e.state.SetPC(pc)

	// Decode the instruction
	var results [1024]byte
	decodedInstr, err := disassemble.Decompose(pc, instr, &results)
	if err != nil {
		return fmt.Errorf("failed to decode instruction at 0x%x: %w", pc, err)
	}

	// Track instruction if tracing is enabled
	if e.enableTrace {
		// Find the actual length of the decoded string
		mnemonicLen := 0
		for i, b := range results {
			if b == 0 {
				mnemonicLen = i
				break
			}
		}
		if mnemonicLen == 0 {
			mnemonicLen = len(results)
		}

		e.trace = append(e.trace, core.InstructionInfo{
			Address:     pc,
			Instruction: decodedInstr,
			Value:       instr,
			Mnemonic:    string(results[:mnemonicLen]),
		})
	}

	// Use the modular instruction registry system
	mnemonic := fmt.Sprintf("%v", decodedInstr.Operation)
	executor, found := e.registry.Get(mnemonic)
	if !found {
		err = fmt.Errorf("unsupported instruction: %s at 0x%x", mnemonic, pc)
		e.LastError = err
		if e.StopOnError {
			return err
		}
		// Advance PC and count to avoid infinite loop when skipping unsupported instructions
		e.state.SetPC(pc + 4)
		e.instrCount++
		return nil
	}

	// Execute using the modular executor
	err = executor.Execute(e.state, decodedInstr)
	if err != nil {
		e.LastError = err
		if e.StopOnError {
			return err
		}
	}

	// Update PC for sequential execution (if not modified by instruction)
	if e.state.GetPC() == pc {
		e.state.SetPC(pc + 4)
	}

	e.instrCount++
	return nil
}

// RunAt executes instructions starting from a specific address
func (e *Engine) RunAt(startPC uint64) error {
	e.state.SetPC(startPC)
	return e.Run()
}

// Run executes instructions from the current PC until a stop condition is met
func (e *Engine) Run() error {
	for e.instrCount < e.maxInstructions {
		pc := e.state.GetPC()

		// Read instruction at PC
		instr, err := e.state.ReadUint32(pc)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				// stop execution cleanly on unmapped fetch
				break
			}
			return fmt.Errorf("failed to read instruction at 0x%x: %w", pc, err)
		}

		// Optional early stop based on configured stop address or handler
		if e.StopAddress != 0 && pc == e.StopAddress {
			break
		}
		if e.ShouldHaltPreHandler != nil {
			var results [1024]byte
			if decodedInstr, derr := disassemble.Decompose(pc, instr, &results); derr == nil {
				mnemonic := fmt.Sprintf("%v", decodedInstr.Operation)
				if _, ok := e.registry.Get(mnemonic); ok {
					info := core.InstructionInfo{
						Address:     pc,
						Instruction: decodedInstr,
						Value:       instr,
						Mnemonic:    mnemonic,
					}
					if e.ShouldHaltPreHandler(e.state, info) {
						break
					}
				}
			}
		}

		// Execute the instruction
		// Optionally decide whether to take a branch before executing it
		if e.ShouldTakeBranchHandler != nil {
			var results [1024]byte
			if decodedInstr, derr := disassemble.Decompose(pc, instr, &results); derr == nil {
				mnemonic := fmt.Sprintf("%v", decodedInstr.Operation)
				if _, ok := e.registry.Get(mnemonic); ok {
					upper := strings.ToUpper(mnemonic)
					if instructions.IsBranchInstruction(upper) {
						info := core.InstructionInfo{
							Address:     pc,
							Instruction: decodedInstr,
							Value:       instr,
							Mnemonic:    mnemonic,
						}
						if !e.ShouldTakeBranchHandler(e.state, info) {
							// Skip branch: treat as NOP
							e.state.SetPC(pc + 4)
							e.instrCount++
							continue
						}
					}
				}
			}
		}
		err = e.ExecuteInstruction(pc, instr)
		if err != nil {
			return err
		}
		// Optional post-execution halt check
		if e.ShouldHaltPostHandler != nil {
			var results [1024]byte
			if decodedInstr, derr := disassemble.Decompose(pc, instr, &results); derr == nil {
				mnemonic := fmt.Sprintf("%v", decodedInstr.Operation)
				if _, ok := e.registry.Get(mnemonic); ok {
					info := core.InstructionInfo{
						Address:     pc,
						Instruction: decodedInstr,
						Value:       instr,
						Mnemonic:    mnemonic,
					}
					if e.ShouldHaltPostHandler(e.state, info) {
						break
					}
				}
			}
		}

		// Check for stop conditions (RET instruction, etc.)
		if e.state.GetPC() == 0 {
			break // Reached a RET or similar
		}
	}

	return nil
}

// StepOver executes a single instruction and advances PC
func (e *Engine) StepOver() error {
	pc := e.state.GetPC()
	instr, err := e.state.ReadUint32(pc)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			// treat as a valid stopping point
			return nil
		}
		return fmt.Errorf("failed to read instruction at 0x%x: %w", pc, err)
	}

	return e.ExecuteInstruction(pc, instr)
}

// GetRegister returns the value of a register by index
func (e *Engine) GetRegister(index int) uint64 {
	return e.state.GetX(index)
}

// SetRegister sets the value of a register by index
func (e *Engine) SetRegister(index int, value uint64) {
	e.state.SetX(index, value)
}

// GetPC returns the current program counter
func (e *Engine) GetPC() uint64 {
	return e.state.GetPC()
}

// SetPC sets the program counter
func (e *Engine) SetPC(pc uint64) {
	e.state.SetPC(pc)
}

// GetSP returns the stack pointer
func (e *Engine) GetSP() uint64 {
	return e.state.GetSP()
}

// SetSP sets the stack pointer
func (e *Engine) SetSP(sp uint64) {
	e.state.SetSP(sp)
}

// SetMemory writes data to memory at the specified address
func (e *Engine) SetMemory(addr uint64, data []byte) error {
	// Map the memory region first
	if armState, ok := e.state.(*state.ARM64State); ok {
		armState.MapRegion(addr, len(data))
	}
	e.state.WriteMemory(addr, data)
	return nil
}

// GetMemory reads data from memory at the specified address
func (e *Engine) GetMemory(addr uint64, size int) ([]byte, error) {
	return e.state.ReadMemory(addr, size)
}

// SetMemoryUint64 writes a 64-bit value to memory
func (e *Engine) SetMemoryUint64(addr uint64, value uint64) error {
	// Map the memory region first
	if armState, ok := e.state.(*state.ARM64State); ok {
		armState.MapRegion(addr, 8) // 8 bytes for uint64
	}
	e.state.WriteUint64(addr, value)
	return nil
}

// GetMemoryUint64 reads a 64-bit value from memory
func (e *Engine) GetMemoryUint64(addr uint64) (uint64, error) {
	return e.state.ReadUint64(addr)
}

// SetMemoryUint32 writes a 32-bit value to memory
func (e *Engine) SetMemoryUint32(addr uint64, value uint32) error {
	// Map the memory region first
	if armState, ok := e.state.(*state.ARM64State); ok {
		armState.MapRegion(addr, 4) // 4 bytes for uint32
	}
	e.state.WriteUint32(addr, value)
	return nil
}

// GetMemoryUint32 reads a 32-bit value from memory
func (e *Engine) GetMemoryUint32(addr uint64) (uint32, error) {
	return e.state.ReadUint32(addr)
}

// GetRegistry returns the instruction registry (for debugging/inspection)
func (e *Engine) GetRegistry() core.InstructionRegistry {
	return e.registry
}

// SetRegistry sets a custom instruction registry
func (e *Engine) SetRegistry(registry core.InstructionRegistry) {
	e.registry = registry
}

// ListSupportedInstructions returns a list of supported instruction mnemonics
func (e *Engine) ListSupportedInstructions() []string {
	return e.registry.List()
}

// Configure applies configuration changes to the engine
func (e *Engine) Configure(config *EngineConfig) {
	if config.MaxInstructions > 0 {
		e.maxInstructions = config.MaxInstructions
	}
	e.enableTrace = config.EnableTrace
	e.StopOnError = config.StopOnError

	if config.InitialPC != 0 {
		e.state.SetPC(config.InitialPC)
	}
	if config.InitialSP != 0 {
		e.state.SetSP(config.InitialSP)
	}

	// Set up memory handlers if provided
	if config.MemoryHandler != nil {
		e.state.SetMemoryReadHandler(config.MemoryHandler)
	}
	if config.StringHandler != nil {
		e.state.SetStringPoolHandler(config.StringHandler)
	}
	if config.PointerResolver != nil {
		e.state.SetPointerResolver(config.PointerResolver)
	}
	e.StopAddress = config.StopAddress
	e.ShouldHaltPreHandler = config.ShouldHaltPreHandler
	e.ShouldHaltPostHandler = config.ShouldHaltPostHandler
	e.ShouldTakeBranchHandler = config.ShouldTakeBranchHandler
}
