package instructions

import (
	"fmt"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// MoveExecutor handles move and data transfer instructions
type MoveExecutor struct {
	*BaseExecutor
}

// NewMoveExecutor creates a new move instruction executor
func NewMoveExecutor(mnemonic, description string) *MoveExecutor {
	return &MoveExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// Execute executes move instructions
func (e *MoveExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.mnemonic {
	case "MOV":
		return e.executeMOV(state, instr)
	case "MOVZ":
		return e.executeMOVZ(state, instr)
	case "MOVN":
		return e.executeMOVN(state, instr)
	case "MOVK":
		return e.executeMOVK(state, instr)
	case "ADR":
		return e.executeADR(state, instr)
	case "ADRP":
		return e.executeADRP(state, instr)
	case "SXTB":
		return e.executeSXTB(state, instr)
	case "SXTH":
		return e.executeSXTH(state, instr)
	case "SXTW":
		return e.executeSXTW(state, instr)
	case "UXTB":
		return e.executeUXTB(state, instr)
	case "UXTH":
		return e.executeUXTH(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("move instruction %s not implemented", e.mnemonic))
	}
}

// MOV - Move (register or immediate)
func (e *MoveExecutor) executeMOV(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOV requires at least 2 operands")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	// Get destination register
	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOV requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	var value uint64

	// Handle MOV alias encodings
	// - ORR (immediate) with XZR source -> MOV immediate
	// - ORR (register) with XZR source -> MOV register
	// Prefer decoding ORR immediate mask when instr.Raw indicates logical immediate form.

	// Get source value (register or immediate)
	if len(srcOp.Registers) > 0 {
		srcRegID := srcOp.Registers[0]
		// Handle SP register specifically, as MOV Xd, SP is a common pattern.
		if uint32(srcRegID) == 66 { // Special ID for SP register from disassembler
			value = state.GetSP()
		} else {
			// Register source
			srcReg := core.MapRegister(srcRegID)
			if srcReg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid source register")
			}
			value = state.GetX(srcReg)
		}
	} else {
		// Immediate source
		// MOV has multiple alias forms:
		// 1) MOV (wide immediate) -> alias of MOVZ/MOVN encodings
		// 2) MOV (logical immediate) -> alias of ORR (immediate) with ZR
		// Prefer decoding the specific encoding from raw bits to choose path.

		raw := instr.Raw
		// Detect Move Wide class (matches MOVZ/MOVN encodings)
		isMovWide := (raw&0x7F800000) == 0x52800000 || (raw&0x7F800000) == 0x12800000 // MOVZ or MOVN pattern
		if isMovWide {
			imm16 := uint64((raw >> 5) & 0xFFFF)
			hw := (raw >> 21) & 0x3
			shiftAmount := uint64(hw * 16)
			opc := (raw >> 29) & 0x3 // 00: MOVN, 10: MOVZ

			if shiftAmount%16 != 0 || shiftAmount > 48 {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid shift amount for MOV (wide immediate)")
			}

			field := (imm16 & 0xFFFF) << shiftAmount
			// Determine width by destination register kind
			is32 := uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31
			if opc == 0 { // MOVN alias
				if is32 {
					value = (^(field)) & 0xFFFFFFFF
				} else {
					value = (^(field)) & 0xFFFFFFFFFFFFFFFF
				}
			} else { // treat as MOVZ alias (zero others)
				value = field
			}
		} else {
			// Logical immediate (ORR alias). Use operand immediate if present; fallback to decoding.
			imm := uint64(srcOp.Immediate)
			if imm == 0 && raw != 0 {
				// Decode ORR-immediate logical mask
				// Width from sf bit
				sf := uint32((raw >> 31) & 1)
				n := (raw >> 22) & 1
				immr := (raw >> 16) & 0x3f
				imms := (raw >> 10) & 0x3f
				if wmask, _, ok := decodeBitMasks(n, imms, immr, sf); ok {
					imm = wmask
				}
			}
			value = imm
		}
	}

	// Handle shift if present
	if len(instr.Operands) > 2 && instr.Operands[2].ShiftValueUsed {
		var err error
		value, err = e.applyShift(value, instr.Operands[2])
		if err != nil {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
		}
	}

	// Store value, handling 32-bit vs 64-bit
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	if uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(value))
	} else {
		state.SetX(dstReg, value)
	}

	return nil
}

// MOVZ - Move wide with zero
func (e *MoveExecutor) executeMOVZ(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVZ requires at least 2 operands")
	}

	dstOp := instr.Operands[0]

	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVZ requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	var immediate uint64
	var shiftAmount uint64

	// Prefer operand-provided immediate/shift when present (unit tests)
	if len(instr.Operands) > 1 && (instr.Operands[1].Immediate != 0 || len(instr.Operands) <= 2) {
		immediate = uint64(instr.Operands[1].Immediate) & 0xFFFF
		if len(instr.Operands) > 2 && instr.Operands[2].ShiftValueUsed {
			shiftOp := instr.Operands[2]
			if shiftOp.ShiftType != disassemble.SHIFT_TYPE_LSL {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "MOVZ only supports LSL shift")
			}
			shiftAmount = uint64(shiftOp.ShiftValue)
		}
	} else if instr.Raw != 0 {
		// Decode from encoding: imm16 in bits [20:5], hw in [22:21]
		instrValue := instr.Raw
		imm16 := uint64((instrValue >> 5) & 0xFFFF)
		hw := (instrValue >> 21) & 0x3
		shiftAmount = uint64(hw * 16)
		immediate = imm16
	} else {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "MOVZ requires source immediate")
	}

	// Validate shift amount
	if shiftAmount%16 != 0 || shiftAmount > 48 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid shift amount for MOVZ")
	}

	// Construct value and zero rest
	value := (immediate & 0xFFFF) << shiftAmount

	// For MOVZ the rest of the destination bits are zeroed explicitly
	if uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(value))
	} else {
		state.SetX(dstReg, value)
	}

	return nil
}

// MOVN - Move wide with NOT
func (e *MoveExecutor) executeMOVN(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVN requires at least 2 operands")
	}

	dstOp := instr.Operands[0]

	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVN requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	var immediate uint64
	var shiftAmount uint64

	// Prefer operand-provided immediate/shift when present (unit tests)
	if len(instr.Operands) > 1 && (instr.Operands[1].Immediate != 0 || len(instr.Operands) <= 2) {
		immediate = uint64(instr.Operands[1].Immediate) & 0xFFFF
		if len(instr.Operands) > 2 && instr.Operands[2].ShiftValueUsed {
			shiftOp := instr.Operands[2]
			if shiftOp.ShiftType != disassemble.SHIFT_TYPE_LSL {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "MOVN only supports LSL shift")
			}
			shiftAmount = uint64(shiftOp.ShiftValue)
		}
	} else if instr.Raw != 0 {
		instrValue := instr.Raw
		imm16 := uint64((instrValue >> 5) & 0xFFFF)
		hw := (instrValue >> 21) & 0x3
		shiftAmount = uint64(hw * 16)
		immediate = imm16
	} else {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "MOVN requires source immediate")
	}

	// Validate shift amount
	if shiftAmount%16 != 0 || shiftAmount > 48 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid shift amount for MOVN")
	}

	// Invert 16-bit field within width
	is32 := uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31
	field := (immediate & 0xFFFF) << shiftAmount
	var value uint64
	if is32 {
		value = (^field) & 0xFFFFFFFF
		state.SetW(dstReg, uint32(value))
	} else {
		value = (^field) & 0xFFFFFFFFFFFFFFFF
		state.SetX(dstReg, value)
	}

	return nil
}

// MOVK - Move wide with keep
func (e *MoveExecutor) executeMOVK(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVK requires at least 2 operands")
	}

	dstOp := instr.Operands[0]

	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MOVK requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	var immediate uint64
	var shiftAmount uint64

	// Extract immediate from instruction encoding or fallback to operand
	if instr.Raw != 0 {
		// Decode imm16 and hw per encoding
		instrValue := instr.Raw
		imm16 := uint64((instrValue >> 5) & 0xFFFF)
		hw := (instrValue >> 21) & 0x3
		shiftAmount = uint64(hw * 16)
		// If a shift operand is also present, it should match encoded hw; prefer encoded
		immediate = imm16
	} else {
		// Fallback for unit tests - use operand immediate and shift
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "MOVK requires source immediate")
		}
		srcOp := instr.Operands[1]
		immediate = uint64(srcOp.Immediate)

		// Get shift amount from operand
		if len(instr.Operands) > 2 && instr.Operands[2].ShiftValueUsed {
			shiftOp := instr.Operands[2]
			if shiftOp.ShiftType == disassemble.SHIFT_TYPE_LSL {
				shiftAmount = uint64(shiftOp.ShiftValue)
			} else {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "MOVK only supports LSL shift")
			}
		}
	}

	// Validate shift amount
	if shiftAmount%16 != 0 || shiftAmount > 48 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid shift amount for MOVK")
	}

	// Get current register value
	currentValue := state.GetX(dstReg)

	// Create mask for the 16-bit field being updated
	mask := uint64(0xFFFF) << shiftAmount

	// Clear the target field and insert new value
	newValue := (currentValue &^ mask) | ((immediate & 0xFFFF) << shiftAmount)

	// Store value, handling 32-bit vs 64-bit
	if uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31 {
		// 32-bit form writes lower 32 bits and zero-extends
		state.SetW(dstReg, uint32(newValue))
	} else {
		state.SetX(dstReg, newValue)
	}

	return nil
}

// ADR - Address to register
func (e *MoveExecutor) executeADR(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADR requires destination register")
	}

	dstOp := instr.Operands[0]

	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADR requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	// Extract 21-bit immediate from instruction encoding
	// ADR instruction format: op(1) | immlo(2) | 10000 | immhi(19) | Rd(5)
	instrValue := instr.Raw

	// Extract immhi (bits 23-5) and immlo (bits 30-29)
	immhi := int64((instrValue >> 5) & 0x7FFFF) // Bits 23-5 (19 bits)
	immlo := int64((instrValue >> 29) & 0x3)    // Bits 30-29 (2 bits)

	// Combine into 21-bit signed offset
	offset := (immhi << 2) | immlo

	// Sign extend from 21 bits to 64 bits
	if offset&(1<<20) != 0 {
		offset |= ^int64((1 << 21) - 1) // Sign extend
	}

	// Calculate PC-relative address
	pc := state.GetPC()
	address := uint64(int64(pc) + offset)

	state.SetX(dstReg, address)
	return nil
}

// ADRP - Address to register, page
func (e *MoveExecutor) executeADRP(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADRP requires destination register")
	}

	dstOp := instr.Operands[0]

	if len(dstOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADRP requires destination register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	// Extract 21-bit immediate from instruction encoding (same as ADR)
	// ADRP instruction format: op(1) | immlo(2) | 10000 | immhi(19) | Rd(5)
	instrValue := instr.Raw

	// Extract immhi (bits 23-5) and immlo (bits 30-29)
	immhi := int64((instrValue >> 5) & 0x7FFFF) // Bits 23-5 (19 bits)
	immlo := int64((instrValue >> 29) & 0x3)    // Bits 30-29 (2 bits)

	// Combine into 21-bit signed offset
	offset := (immhi << 2) | immlo

	// Sign extend from 21 bits to 64 bits
	if offset&(1<<20) != 0 {
		offset |= ^int64((1 << 21) - 1) // Sign extend
	}

	// For ADRP: scale offset by 4KB page size and add to page-aligned PC
	pc := state.GetPC() &^ 0xFFF // Clear lower 12 bits (page align PC)
	pageOffset := offset << 12   // Scale offset to 4KB pages
	address := uint64(int64(pc) + pageOffset)

	state.SetX(dstReg, address)
	return nil
}

// SXTB - Sign extend byte
func (e *MoveExecutor) executeSXTB(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTB requires destination and source registers")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	if len(dstOp.Registers) == 0 || len(srcOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTB requires register operands")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Get low byte and sign extend
	srcValue := state.GetX(srcReg)
	byteValue := int8(srcValue & 0xFF)
	extendedValue := uint64(int64(byteValue))

	state.SetX(dstReg, extendedValue)
	return nil
}

// SXTH - Sign extend halfword
func (e *MoveExecutor) executeSXTH(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTH requires destination and source registers")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	if len(dstOp.Registers) == 0 || len(srcOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTH requires register operands")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Get low halfword and sign extend
	srcValue := state.GetX(srcReg)
	halfwordValue := int16(srcValue & 0xFFFF)
	extendedValue := uint64(int64(halfwordValue))

	state.SetX(dstReg, extendedValue)
	return nil
}

// SXTW - Sign extend word
func (e *MoveExecutor) executeSXTW(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTW requires destination and source registers")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	if len(dstOp.Registers) == 0 || len(srcOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SXTW requires register operands")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Get low word and sign extend
	srcValue := state.GetX(srcReg)
	wordValue := int32(srcValue & 0xFFFFFFFF)
	extendedValue := uint64(int64(wordValue))

	state.SetX(dstReg, extendedValue)
	return nil
}

// UXTB - Zero extend byte (alias for AND with 0xFF)
func (e *MoveExecutor) executeUXTB(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UXTB requires destination and source registers")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	if len(dstOp.Registers) == 0 || len(srcOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UXTB requires register operands")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Zero extend byte (keep only low 8 bits)
	srcValue := state.GetX(srcReg)
	extendedValue := srcValue & 0xFF

	state.SetX(dstReg, extendedValue)
	return nil
}

// UXTH - Zero extend halfword (alias for AND with 0xFFFF)
func (e *MoveExecutor) executeUXTH(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UXTH requires destination and source registers")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	if len(dstOp.Registers) == 0 || len(srcOp.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UXTH requires register operands")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Zero extend halfword (keep only low 16 bits)
	srcValue := state.GetX(srcReg)
	extendedValue := srcValue & 0xFFFF

	state.SetX(dstReg, extendedValue)
	return nil
}

// Helper method to apply shift operations
func (e *MoveExecutor) applyShift(value uint64, shiftOp disassemble.Operand) (uint64, error) {
	shiftType := shiftOp.ShiftType
	shiftAmount := uint64(shiftOp.ShiftValue)

	switch shiftType {
	case disassemble.SHIFT_TYPE_LSL: // Logical shift left
		if shiftAmount >= 64 {
			return 0, nil
		}
		return value << shiftAmount, nil
	case disassemble.SHIFT_TYPE_LSR: // Logical shift right
		if shiftAmount >= 64 {
			return 0, nil
		}
		return value >> shiftAmount, nil
	case disassemble.SHIFT_TYPE_ASR: // Arithmetic shift right
		if shiftAmount >= 64 {
			if (value & 0x8000000000000000) != 0 {
				return 0xFFFFFFFFFFFFFFFF, nil // Sign extend
			}
			return 0, nil
		}
		return uint64(int64(value) >> shiftAmount), nil
	case disassemble.SHIFT_TYPE_ROR: // Rotate right
		shiftAmount %= 64
		return (value >> shiftAmount) | (value << (64 - shiftAmount)), nil
	default:
		return 0, fmt.Errorf("unsupported shift type: %v", shiftType)
	}
}

// RegisterMoveInstructions registers all move instructions
func RegisterMoveInstructions(registry *Registry) {
	// Basic move instructions
	registry.Register("MOV", NewMoveExecutor("MOV", "Move register or immediate"))
	registry.Register("MOVZ", NewMoveExecutor("MOVZ", "Move wide with zero"))
	registry.Register("MOVN", NewMoveExecutor("MOVN", "Move wide with NOT"))
	registry.Register("MOVK", NewMoveExecutor("MOVK", "Move wide with keep"))

	// Address calculation
	registry.Register("ADR", NewMoveExecutor("ADR", "Address to register"))
	registry.Register("ADRP", NewMoveExecutor("ADRP", "Address to register, page"))

	// Sign/zero extension
	registry.Register("SXTB", NewMoveExecutor("SXTB", "Sign extend byte"))
	registry.Register("SXTH", NewMoveExecutor("SXTH", "Sign extend halfword"))
	registry.Register("SXTW", NewMoveExecutor("SXTW", "Sign extend word"))
	registry.Register("UXTB", NewMoveExecutor("UXTB", "Zero extend byte"))
	registry.Register("UXTH", NewMoveExecutor("UXTH", "Zero extend halfword"))
}

// Helper functions for move instruction analysis

// IsMoveInstruction checks if a mnemonic is a move instruction
func IsMoveInstruction(mnemonic string) bool {
	switch mnemonic {
	case "MOV", "MOVZ", "MOVN", "MOVK", "ADR", "ADRP",
		"SXTB", "SXTH", "SXTW", "UXTB", "UXTH":
		return true
	default:
		return false
	}
}

// IsWideMove checks if an instruction is a wide move (MOVZ, MOVN, MOVK)
func IsWideMove(mnemonic string) bool {
	switch mnemonic {
	case "MOVZ", "MOVN", "MOVK":
		return true
	default:
		return false
	}
}

// IsAddressCalculation checks if an instruction calculates addresses
func IsAddressCalculation(mnemonic string) bool {
	switch mnemonic {
	case "ADR", "ADRP":
		return true
	default:
		return false
	}
}

// IsExtensionInstruction checks if an instruction performs sign/zero extension
func IsExtensionInstruction(mnemonic string) bool {
	switch mnemonic {
	case "SXTB", "SXTH", "SXTW", "UXTB", "UXTH":
		return true
	default:
		return false
	}
}

// ValidateWideMove validates parameters for wide move instructions
func ValidateWideMove(mnemonic string, shiftAmount uint64) error {
	if !IsWideMove(mnemonic) {
		return fmt.Errorf("not a wide move instruction: %s", mnemonic)
	}

	// Wide moves only support shifts in multiples of 16, up to 48
	if shiftAmount%16 != 0 || shiftAmount > 48 {
		return fmt.Errorf("invalid shift amount for %s: %d (must be 0, 16, 32, or 48)", mnemonic, shiftAmount)
	}

	return nil
}
