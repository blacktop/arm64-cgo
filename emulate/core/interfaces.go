package core

import "github.com/blacktop/arm64-cgo/disassemble"

// State represents the ARM64 processor state interface
type State interface {
	// Register access
	GetX(reg int) uint64
	SetX(reg int, val uint64)
	GetW(reg int) uint32
	SetW(reg int, val uint32)

	// Special registers
	GetPC() uint64
	SetPC(addr uint64)
	GetSP() uint64
	SetSP(addr uint64)

	// Flags access
	GetN() bool
	SetN(val bool)
	GetZ() bool
	SetZ(val bool)
	GetC() bool
	SetC(val bool)
	GetV() bool
	SetV(val bool)
	UpdateFlags(result uint64, isNegative bool)

	// Memory operations
	ReadMemory(addr uint64, size int) ([]byte, error)
	WriteMemory(addr uint64, data []byte)
	ReadUint64(addr uint64) (uint64, error)
	ReadUint32(addr uint64) (uint32, error)
	WriteUint64(addr uint64, value uint64)
	WriteUint32(addr uint64, value uint32)

	// String and pointer operations
	ReadString(addr uint64) (string, error)
	ReadStringWithLength(addr uint64, length int) (string, error)
	ResolvePointer(addr uint64) (uint64, error)
	FollowPointerChain(addr uint64, maxDepth int) (uint64, error)
	ReadStringAtPointer(ptrAddr uint64) (string, error)

	// Handler management
	SetMemoryReadHandler(handler MemoryReadHandler)
	SetStringPoolHandler(handler StringPoolHandler)
	SetPointerResolver(resolver PointerResolver)

	// State management
	Clone() State
	Reset()
}

// InstructionExecutor represents an instruction execution interface
type InstructionExecutor interface {
	Execute(state State, instr *disassemble.Instruction) error
	Supports(mnemonic string) bool
}

// AddressingMode represents different addressing modes for ARM64
type AddressingMode interface {
	CalculateAddress(state State, operand disassemble.Operand) (uint64, error)
	SupportsWriteback() bool
}

// EmulationEngine represents the main emulation engine interface
type EmulationEngine interface {
	// Core emulation
	EmulateOne(state State, instr *InstructionInfo) error
	Emulate(data []byte, startAddr uint64, state State) (State, error)
	EmulateUntil(data []byte, startAddr uint64, state State, stopAddr *uint64) (State, error)

	// Instruction management
	RegisterExecutor(mnemonic string, executor InstructionExecutor)
	GetExecutor(mnemonic string) (InstructionExecutor, bool)

	// Configuration
	SetMaxInstructions(limit int)
	GetInstructionCount() int
	Reset()
}

// MemoryManager represents memory management interface
type MemoryManager interface {
	Read(addr uint64, size int) ([]byte, error)
	Write(addr uint64, data []byte) error
	IsValidAddress(addr uint64) bool
	MapRegion(addr uint64, size int) error
	UnmapRegion(addr uint64, size int) error
}

// RegisterManager represents register management interface
type RegisterManager interface {
	GetRegister(id RegisterID) (uint64, error)
	SetRegister(id RegisterID, value uint64) error
	GetRegisterName(id RegisterID) string
	IsValidRegister(id RegisterID) bool
	ClearRegisters()
}

// FlagManager represents processor flags management interface
type FlagManager interface {
	GetFlags() ProcessorFlags
	SetFlags(flags ProcessorFlags)
	UpdateArithmeticFlags(result uint64, carry bool, overflow bool)
	EvaluateCondition(cond ConditionCode) bool
}

// InstructionRegistry manages instruction executors
type InstructionRegistry interface {
	Register(mnemonic string, executor InstructionExecutor) error
	Get(mnemonic string) (InstructionExecutor, bool)
	List() []string
	Clear()
}

// Debugger represents debugging capabilities
type Debugger interface {
	// Breakpoints
	SetBreakpoint(addr uint64) error
	RemoveBreakpoint(addr uint64) error
	IsBreakpoint(addr uint64) bool

	// Single stepping
	Step() error
	StepOver() error
	StepOut() error

	// Tracing
	EnableTrace(enable bool)
	GetTrace() []InstructionInfo
	ClearTrace()
}

// ProfilerI represents profiling capabilities
type ProfilerI interface {
	Start()
	Stop()
	GetStats() map[string]int
	Reset()
}

// Disassembler represents disassembly capabilities
type Disassembler interface {
	DisassembleAt(addr uint64, data []byte) (*disassemble.Instruction, error)
	DisassembleRange(startAddr uint64, data []byte, count int) ([]*disassemble.Instruction, error)
}

// StateSnapshot represents a snapshot of processor state
type StateSnapshot interface {
	Restore(state State) error
	GetPC() uint64
	GetRegisters() [31]uint64
	GetFlags() ProcessorFlags
}

// StateManager manages state snapshots and restoration
type StateManager interface {
	CreateSnapshot(state State) StateSnapshot
	RestoreSnapshot(state State, snapshot StateSnapshot) error
	ListSnapshots() []StateSnapshot
	ClearSnapshots()
}

// EmulationContext provides execution context
type EmulationContext interface {
	GetState() State
	GetEngine() EmulationEngine
	GetDebugger() Debugger
	GetProfiler() ProfilerI
	SetTimeout(milliseconds int)
	IsTimeout() bool
}

// Factory interfaces for creating components
type StateFactory interface {
	NewState() State
	NewStateWithConfig(config StateConfig) State
}

// Configuration structures
type StateConfig struct {
	StackSize      int
	StackBase      uint64
	InitialPC      uint64
	InitialSP      uint64
	EnableDebug    bool
	EnableProfile  bool
	MemoryHandler  MemoryReadHandler
	StringHandler  StringPoolHandler
	PointerHandler PointerResolver
}

type EngineConfig struct {
	MaxInstructions   int
	EnableTrace       bool
	EnableBreakpoints bool
	TimeoutMs         int
}
