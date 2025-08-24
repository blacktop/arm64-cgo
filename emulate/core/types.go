package core

import (
	"errors"
	"fmt"

	"github.com/blacktop/arm64-cgo/disassemble"
)

// Common errors for ARM64 emulation
var (
	ErrInvalidRegister    = errors.New("invalid register")
	ErrInvalidInstruction = errors.New("invalid instruction")
	ErrMemoryAccess       = errors.New("memory access error")
	ErrDivisionByZero     = errors.New("division by zero")
	ErrInvalidAddress     = errors.New("invalid address")
	ErrUnsupportedFeature = errors.New("unsupported feature")
	ErrExecutionLimit     = errors.New("execution limit exceeded")
)

// EmulationError represents a specific emulation error with context
type EmulationError struct {
	Err    error
	PC     uint64
	Instr  string
	Detail string
}

func (e *EmulationError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("emulation error at PC=0x%x, instruction=%s: %v (%s)", e.PC, e.Instr, e.Err, e.Detail)
	}
	return fmt.Sprintf("emulation error at PC=0x%x, instruction=%s: %v", e.PC, e.Instr, e.Err)
}

func (e *EmulationError) Unwrap() error {
	return e.Err
}

// NewEmulationError creates a new emulation error with context
func NewEmulationError(err error, pc uint64, instrStr string, detail string) *EmulationError {
	return &EmulationError{
		Err:    err,
		PC:     pc,
		Instr:  instrStr,
		Detail: detail,
	}
}

// HandlerFunc represents a function type for custom handlers
type HandlerFunc[T any] func(T) error

// MemoryReadHandler is a function type for custom memory reads
type MemoryReadHandler func(addr uint64, size int) ([]byte, error)

// StringPoolHandler is a function type for reading strings from external sources
type StringPoolHandler func(addr uint64) (string, error)

// PointerResolver is a function type for resolving pointers to their target addresses
type PointerResolver func(addr uint64) (uint64, error)

// RegisterID represents a register identifier
type RegisterID int

// Standard ARM64 register constants
const (
	// General purpose registers X0-X30
	X0 RegisterID = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27
	X28
	X29
	X30

	// Special purpose registers
	SP_REG // Stack Pointer
	PC_REG // Program Counter
	LR_REG // Link Register (alias for X30)
	FP_REG // Frame Pointer (alias for X29)
	ZR_REG // Zero Register

	INVALID_REG = -1
)

// ConditionCode represents ARM64 condition codes
type ConditionCode int

const (
	EQ ConditionCode = iota // Equal
	NE                      // Not Equal
	CS                      // Carry Set
	CC                      // Carry Clear
	MI                      // Minus/Negative
	PL                      // Plus/Positive
	VS                      // Overflow Set
	VC                      // Overflow Clear
	HI                      // Higher (unsigned)
	LS                      // Lower or Same (unsigned)
	GE                      // Greater or Equal (signed)
	LT                      // Less Than (signed)
	GT                      // Greater Than (signed)
	LE                      // Less or Equal (signed)
	AL                      // Always
	NV                      // Never
)

// String returns the string representation of a condition code
func (cc ConditionCode) String() string {
	switch cc {
	case EQ:
		return "EQ"
	case NE:
		return "NE"
	case CS:
		return "CS"
	case CC:
		return "CC"
	case MI:
		return "MI"
	case PL:
		return "PL"
	case VS:
		return "VS"
	case VC:
		return "VC"
	case HI:
		return "HI"
	case LS:
		return "LS"
	case GE:
		return "GE"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case LE:
		return "LE"
	case AL:
		return "AL"
	case NV:
		return "NV"
	default:
		return "UNKNOWN"
	}
}

// InstructionInfo contains metadata about an instruction
type InstructionInfo struct {
	Address     uint64
	Instruction *disassemble.Instruction
	Value       uint32
	Mnemonic    string
}

// ProcessorFlags represents the ARM64 processor state flags (NZCV)
type ProcessorFlags struct {
	N bool // Negative
	Z bool // Zero
	C bool // Carry
	V bool // Overflow
}

// PSTATE represents the processor state
type PSTATE struct {
	Flags ProcessorFlags
	// Additional processor state fields can be added here
}

// UpdateFlags updates processor flags based on a result
func (pstate *PSTATE) UpdateFlags(result uint64, isNegative bool) {
	pstate.Flags.N = isNegative
	pstate.Flags.Z = result == 0
	// C and V flags are set by specific operations
}

// MapRegister maps disassembler register IDs to our register array indices
func MapRegister(regID disassemble.Register) int {
	// Based on disassembler constants:
	// X0 = 34, X1 = 35, ..., X30 = 64, XZR = 65
	// W0 = 1, W1 = 2, ..., W30 = 31, WZR = 32
	// SP = 66
	reg := uint32(regID)
	if reg >= 34 && reg <= 64 {
		return int(reg - 34) // X0-X30 -> 0-30
	} else if reg >= 1 && reg <= 31 {
		return int(reg - 1) // W0-W30 -> 0-30 (same register index as X)
	} else if reg == 65 { // REG_XZR
		return 31 // XZR maps to register 31 (handled specially as zero register)
	} else if reg == 32 { // REG_WZR  
		return 31 // WZR maps to register 31 (handled specially as zero register)
	} else if reg == 66 { // REG_SP
		return 31 // SP maps to register 31
	}
	return -1 // Invalid register
}

// RegisterName returns the name of a register
func RegisterName(regID RegisterID) string {
	if regID >= X0 && regID <= X30 {
		return fmt.Sprintf("X%d", int(regID))
	}
	switch regID {
	case SP_REG:
		return "SP"
	case PC_REG:
		return "PC"
	case LR_REG:
		return "LR"
	case FP_REG:
		return "FP"
	case ZR_REG:
		return "XZR"
	default:
		return "UNKNOWN"
	}
}

// IsValidRegister checks if a register ID is valid
func IsValidRegister(regID RegisterID) bool {
	return regID >= X0 && regID <= ZR_REG
}

// IsGeneralPurposeRegister checks if a register is a general purpose register
func IsGeneralPurposeRegister(regID RegisterID) bool {
	return regID >= X0 && regID <= X30
}
