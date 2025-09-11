package addressing

import (
	"fmt"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// AddressingMode represents the type of addressing mode
type AddressingMode int

const (
	// ModeOffset represents [Xn, #imm] - Base register + immediate offset
	ModeOffset AddressingMode = iota
	// ModePreIndex represents [Xn, #imm]! - Pre-index with writeback
	ModePreIndex
	// ModePostIndex represents [Xn], #imm - Post-index with writeback
	ModePostIndex
	// ModeRegister represents [Xn] or [Xn, Xm] - Base register only or base + index register
	ModeRegister
	// ModeExtended represents [Xn, Xm, extend] - Base + extended index register
	ModeExtended
)

// AddressCalculation holds the result of an address calculation
type AddressCalculation struct {
	Address           uint64 // The calculated effective address
	WritebackValue    uint64 // Value to write back to base register (if applicable)
	RequiresWriteback bool   // Whether writeback is required
	BaseRegister      int    // Index of base register for writeback
}

// ShiftType constants matching ARM64 specification
const (
	ShiftTypeNone = 0
	ShiftTypeLSL  = 1
	ShiftTypeLSR  = 2
	ShiftTypeASR  = 3
	ShiftTypeROR  = 4
	ShiftTypeUXTW = 5
	ShiftTypeSXTW = 6
	ShiftTypeSXTX = 7
	ShiftTypeUXTX = 8
	ShiftTypeSXTB = 9
	ShiftTypeSXTH = 10
	ShiftTypeUXTH = 11
	ShiftTypeUXTB = 12
	ShiftTypeMSL  = 13
)

// ExtendInfo represents extend operation information
type ExtendInfo struct {
	Type   int    // Extend type (UXTW, SXTW, etc.)
	Amount uint32 // Shift amount after extend
}

// ShiftInfo represents shift operation information
type ShiftInfo struct {
	Type   int    // Shift type (LSL, LSR, ASR, ROR)
	Amount uint32 // Shift amount
}

// CalculateAddress calculates the effective memory address for load/store operations
func CalculateAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	switch operand.Class {
	case disassemble.MEM_OFFSET:
		return calculateOffsetAddress(state, operand)
	case disassemble.MEM_PRE_IDX:
		return calculatePreIndexAddress(state, operand)
	case disassemble.MEM_POST_IDX:
		return calculatePostIndexAddress(state, operand)
	case disassemble.MEM_REG:
		return calculateRegisterAddress(state, operand)
	case disassemble.MEM_EXTENDED:
		return calculateExtendedAddress(state, operand)
	default:
		// Default to register addressing mode for operands without explicit class
		// This handles cases like [Xn] where Class is not set
		return calculateRegisterAddress(state, operand)
	}
}

// calculateOffsetAddress handles [Xn, #imm] - Base register + immediate offset
func calculateOffsetAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	if len(operand.Registers) < 1 {
		return nil, fmt.Errorf("offset addressing: missing base register")
	}

	baseAddr, err := getRegisterValue(state, operand.Registers[0])
	if err != nil {
		return nil, err
	}

	// Calculate address with immediate offset
	addr := baseAddr
	if operand.SignedImm {
		// Many signed memory offsets in A64 are 9-bit (e.g., LDUR/STUR)
		const signedMemImmWidth = 9
		off := signExtend(operand.Immediate, signedMemImmWidth)
		addr = uint64(int64(addr) + off)
	} else {
		addr += operand.Immediate
	}

	return &AddressCalculation{
		Address:           addr,
		WritebackValue:    0,
		RequiresWriteback: false,
		BaseRegister:      -1,
	}, nil
}

// calculatePreIndexAddress handles [Xn, #imm]! - Pre-index with writeback
func calculatePreIndexAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	if len(operand.Registers) < 1 {
		return nil, fmt.Errorf("pre-index addressing requires base register")
	}

	regID := operand.Registers[0]
	baseReg := mapRegister(regID)
	isSP := false
	var baseAddr uint64

	if baseReg >= 0 {
		baseAddr = getXRegister(state, baseReg)
	} else {
		// Handle SP as base register
		reg := uint32(regID)
		if reg == 31 || reg > 64 {
			isSP = true
			baseAddr = state.GetSP()
		} else {
			return nil, fmt.Errorf("invalid base register")
		}
	}

	// Calculate new address with sign-extended offset (pre-index typically uses 9-bit signed imm)
	const prePostImmWidth = 9
	var off int64
	if operand.SignedImm {
		off = signExtend(operand.Immediate, prePostImmWidth)
	} else {
		off = int64(operand.Immediate)
	}
	newAddr := uint64(int64(baseAddr) + off)

	baseIndex := baseReg
	if isSP {
		baseIndex = 31 // sentinel to indicate SP for ApplyWriteback
	}

	return &AddressCalculation{
		Address:           newAddr,
		WritebackValue:    newAddr,
		RequiresWriteback: true,
		BaseRegister:      baseIndex,
	}, nil
}

// calculatePostIndexAddress handles [Xn], #imm - Post-index with writeback
func calculatePostIndexAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	if len(operand.Registers) < 1 {
		return nil, fmt.Errorf("post-index addressing requires base register")
	}

	regID := operand.Registers[0]
	baseReg := mapRegister(regID)
	isSP := false
	var baseAddr uint64

	if baseReg >= 0 {
		baseAddr = getXRegister(state, baseReg)
	} else {
		// Handle SP as base register
		reg := uint32(regID)
		if reg == 31 || reg > 64 {
			isSP = true
			baseAddr = state.GetSP()
		} else {
			return nil, fmt.Errorf("invalid base register")
		}
	}

	// Calculate new address for writeback using sign-extended offset
	const prePostImmWidth = 9
	var off int64
	if operand.SignedImm {
		off = signExtend(operand.Immediate, prePostImmWidth)
	} else {
		off = int64(operand.Immediate)
	}
	writebackAddr := uint64(int64(baseAddr) + off)

	baseIndex := baseReg
	if isSP {
		baseIndex = 31 // sentinel to indicate SP for ApplyWriteback
	}

	return &AddressCalculation{
		Address:           baseAddr, // Use original base address for memory operation
		WritebackValue:    writebackAddr,
		RequiresWriteback: true,
		BaseRegister:      baseIndex,
	}, nil
}

// calculateRegisterAddress handles [Xn] or [Xn, Xm] - Base register only or base + index register
func calculateRegisterAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	if len(operand.Registers) < 1 {
		return nil, fmt.Errorf("register addressing: missing base register")
	}

	baseAddr, err := getRegisterValue(state, operand.Registers[0])
	if err != nil {
		return nil, err
	}

	// Check if there's an index register
	if len(operand.Registers) >= 2 {
		indexVal, err := getRegisterValue(state, operand.Registers[1])
		if err != nil {
			return nil, err
		}
		// Apply extend/shift if provided
		if operand.ShiftValueUsed {
			indexVal = applyExtendType(indexVal, operand.ShiftType, operand.ShiftValue)
		}
		baseAddr += indexVal
	}

	return &AddressCalculation{
		Address:           baseAddr,
		WritebackValue:    0,
		RequiresWriteback: false,
		BaseRegister:      -1,
	}, nil
}

// calculateExtendedAddress handles [Xn, Xm, extend] - Base + extended index register
func calculateExtendedAddress(state core.State, operand disassemble.Operand) (*AddressCalculation, error) {
	if len(operand.Registers) < 2 {
		return nil, fmt.Errorf("extended addressing: insufficient registers")
	}

	baseAddr, err := getRegisterValue(state, operand.Registers[0])
	if err != nil {
		return nil, err
	}

	indexVal, err := getRegisterValue(state, operand.Registers[1])
	if err != nil {
		return nil, err
	}

	// Apply extension based on the operand's ShiftType and ShiftValue
	if operand.ShiftValueUsed {
		extendedVal := applyExtendType(indexVal, operand.ShiftType, operand.ShiftValue)
		baseAddr += extendedVal
	} else {
		baseAddr += indexVal
	}

	return &AddressCalculation{
		Address:           baseAddr,
		WritebackValue:    0,
		RequiresWriteback: false,
		BaseRegister:      -1,
	}, nil
}

// ApplyWriteback applies the writeback value to the base register if required
func ApplyWriteback(state core.State, calc *AddressCalculation) error {
	if !calc.RequiresWriteback || calc.BaseRegister < 0 {
		return nil
	}

	if calc.BaseRegister < 31 {
		state.SetX(calc.BaseRegister, calc.WritebackValue)
	} else {
		// Handle SP register
		state.SetSP(calc.WritebackValue)
	}

	return nil
}

// ApplyExtend applies ARM64 extend operations to a value
func ApplyExtend(value uint64, extendType int, amount uint32) uint64 {
	switch extendType {
	case ShiftTypeUXTW:
		// Unsigned extend word to 64 bits (zero-extend 32-bit value)
		value = uint64(uint32(value))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeSXTW:
		// Signed extend word to 64 bits (sign-extend 32-bit value)
		value = uint64(int64(int32(value)))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeUXTX:
		// Unsigned extend doubleword (no-op for 64-bit, optional shift)
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeSXTX:
		// Signed extend doubleword (no-op for 64-bit, optional shift)
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeUXTB:
		// Unsigned extend byte (zero-extend 8-bit value)
		value = uint64(uint8(value))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeSXTB:
		// Signed extend byte (sign-extend 8-bit value)
		value = uint64(int64(int8(value)))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeUXTH:
		// Unsigned extend halfword (zero-extend 16-bit value)
		value = uint64(uint16(value))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	case ShiftTypeSXTH:
		// Signed extend halfword (sign-extend 16-bit value)
		value = uint64(int64(int16(value)))
		if amount > 0 && amount < 64 {
			value = value << amount
		}
		return value
	default:
		return value
	}
}

// ApplyShift applies ARM64 shift operations to a value
func ApplyShift(value uint64, shiftType int, amount uint32) uint64 {
	switch shiftType {
	case ShiftTypeLSL:
		// Logical shift left
		if amount < 64 {
			return value << amount
		}
		return 0
	case ShiftTypeLSR:
		// Logical shift right
		if amount < 64 {
			return value >> amount
		}
		return 0
	case ShiftTypeASR:
		// Arithmetic shift right (preserves sign)
		if amount < 64 {
			return uint64(int64(value) >> amount)
		}
		// ASR by 64 or more results in all 0s or all 1s depending on sign
		if int64(value) < 0 {
			return 0xFFFFFFFFFFFFFFFF
		}
		return 0
	case ShiftTypeROR:
		// Rotate right
		if amount > 0 {
			amount = amount % 64 // Handle amounts > 64
			return (value >> amount) | (value << (64 - amount))
		}
		return value
	default:
		return value
	}
}

// Helper functions

// mapRegister maps disassembler register IDs to our register array indices
func mapRegister(regID disassemble.Register) int {
	// Based on disassembler constants:
	// X0 = 34, X1 = 35, ..., X30 = 64
	// W0 = 1, W1 = 2, ..., W30 = 31
	reg := uint32(regID)
	if reg >= 34 && reg <= 64 {
		return int(reg - 34) // X0-X30 -> 0-30
	} else if reg >= 1 && reg <= 31 {
		return int(reg - 1) // W0-W30 -> 0-30 (same register index as X)
	}
	return -1 // Invalid register
}

// getRegisterValue gets the value from a register (handles X registers and SP)
func getRegisterValue(state core.State, regID disassemble.Register) (uint64, error) {
	regIdx := mapRegister(regID)
	if regIdx >= 0 {
		return state.GetX(regIdx), nil
	}

	// Check if this is the SP register (typically encoded differently)
	reg := uint32(regID)
	if reg == 31 || reg > 64 { // SP has special encoding
		return state.GetSP(), nil
	}

	return 0, fmt.Errorf("invalid register ID %d", regID)
}

// getXRegister gets the value from an X register by index
func getXRegister(state core.State, regIdx int) uint64 {
	if regIdx >= 0 && regIdx < 31 {
		return state.GetX(regIdx)
	}
	return 0
}

// signExtend sign-extends an immediate of given bit width and returns it as int64
func signExtend(imm uint64, width int) int64 {
	if width <= 0 || width >= 64 {
		return int64(imm)
	}
	mask := uint64((1 << width) - 1)
	u := imm & mask
	sign := uint64(1) << (width - 1)
	if (u & sign) != 0 {
		return int64(u | ^mask)
	}
	return int64(u)
}

// applyExtendType applies ARM64 extension types to a value (legacy compatibility)
func applyExtendType(value uint64, shiftType any, shiftValue uint32) uint64 {
	shiftTypeVal := 0

	// Handle different types that might be passed
	switch v := shiftType.(type) {
	case int:
		shiftTypeVal = v
	case uint32:
		shiftTypeVal = int(v)
	case uint8:
		shiftTypeVal = int(v)
	default:
		// Try to parse string representation
		str := fmt.Sprintf("%v", shiftType)
		switch str {
		case "LSL", "1":
			shiftTypeVal = ShiftTypeLSL
		case "LSR", "2":
			shiftTypeVal = ShiftTypeLSR
		case "ASR", "3":
			shiftTypeVal = ShiftTypeASR
		case "UXTW", "5":
			shiftTypeVal = ShiftTypeUXTW
		case "SXTW", "6":
			shiftTypeVal = ShiftTypeSXTW
		case "UXTX", "8":
			shiftTypeVal = ShiftTypeUXTX
		case "SXTX", "7":
			shiftTypeVal = ShiftTypeSXTX
		case "UXTB", "12":
			shiftTypeVal = ShiftTypeUXTB
		case "SXTB", "9":
			shiftTypeVal = ShiftTypeSXTB
		case "UXTH", "11":
			shiftTypeVal = ShiftTypeUXTH
		case "SXTH", "10":
			shiftTypeVal = ShiftTypeSXTH
		}
	}

	// Determine if this is an extend or shift operation
	switch shiftTypeVal {
	case ShiftTypeUXTW, ShiftTypeSXTW, ShiftTypeUXTX, ShiftTypeSXTX,
		ShiftTypeUXTB, ShiftTypeSXTB, ShiftTypeUXTH, ShiftTypeSXTH:
		return ApplyExtend(value, shiftTypeVal, shiftValue)
	case ShiftTypeLSL, ShiftTypeLSR, ShiftTypeASR, ShiftTypeROR:
		return ApplyShift(value, shiftTypeVal, shiftValue)
	default:
		return value
	}
}
