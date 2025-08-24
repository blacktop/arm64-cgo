package instructions

import (
	"fmt"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// ArithmeticExecutor handles arithmetic instructions
type ArithmeticExecutor struct {
	*BaseExecutor
}

// NewArithmeticExecutor creates a new arithmetic instruction executor
func NewArithmeticExecutor(mnemonic, description string) *ArithmeticExecutor {
	return &ArithmeticExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// Execute executes arithmetic instructions
func (e *ArithmeticExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.mnemonic {
	case "ADD":
		return e.executeADD(state, instr)
	case "ADDG":
		return e.executeADDG(state, instr)
	case "ADDS":
		return e.executeADDS(state, instr)
	case "SUB":
		return e.executeSUB(state, instr)
	case "SUBS":
		return e.executeSUBS(state, instr)
	case "MUL":
		return e.executeMUL(state, instr)
	case "UDIV":
		return e.executeUDIV(state, instr)
	case "SDIV":
		return e.executeSDIV(state, instr)
	case "MADD":
		return e.executeMADD(state, instr)
	case "MSUB":
		return e.executeMSUB(state, instr)
	case "SMULL":
		return e.executeSMULL(state, instr)
	case "UMULL":
		return e.executeUMULL(state, instr)
	case "NEG":
		return e.executeNEG(state, instr)
	case "NEGS":
		return e.executeNEGS(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("arithmetic instruction %s not implemented", e.mnemonic))
	}
}

func (e *ArithmeticExecutor) executeADD(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADD requires at least 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get destination register
	var setDstRegister func(uint64)
	if len(dstOp.Registers) > 0 {
		dstRegID := dstOp.Registers[0]
		reg := uint32(dstRegID)
		if reg == 66 { // explicit SP id
			setDstRegister = func(val uint64) { state.SetSP(val) }
		} else if reg == 31 || reg == 65 { // WZR/XZR
			// XZR/WZR discard writes
			setDstRegister = func(val uint64) {}
		} else {
			dstReg := core.MapRegister(dstRegID)
			if dstReg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid destination register: %d", reg))
			}
			setDstRegister = func(val uint64) { state.SetX(dstReg, val) }
		}
	} else {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "no destination register")
	}

	// Get first operand value
	var val1 uint64
	if len(src1Op.Registers) > 0 {
		src1RegID := src1Op.Registers[0]
		reg := uint32(src1RegID)
		if reg == 66 { // SP
			val1 = state.GetSP()
		} else if reg == 31 || reg == 65 { // WZR/XZR
			val1 = 0
		} else {
			src1Reg := core.MapRegister(src1RegID)
			if src1Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid source register: %d", reg))
			}
			val1 = state.GetX(src1Reg)
		}
	}

	// Get second operand (register or immediate)
	var val2 uint64
	if len(src2Op.Registers) > 0 {
		src2RegID := src2Op.Registers[0]
		reg := uint32(src2RegID)
		if reg == 66 { // SP
			val2 = state.GetSP()
		} else if reg == 31 || reg == 65 { // WZR/XZR
			val2 = 0
		} else {
			src2Reg := core.MapRegister(src2RegID)
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid second source register: %d", reg))
			}
			val2 = state.GetX(src2Reg)
		}
	} else {
		// Immediate operand - check for optional LSL #12 shift
		val2 = uint64(src2Op.Immediate)
		if (instr.Raw>>22)&1 == 1 { // Check the 'sh' bit
			val2 <<= 12
		}
	}

	// Handle shift/extend on second operand
	if src2Op.ShiftValueUsed {
		var err error
		val2, err = applyShift(val2, src2Op)
		if err != nil {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
		}
	}

	// Handle optional shift/extend in 4th operand (for extended registers)
	if len(instr.Operands) > 3 {
		shiftOp := instr.Operands[3]
		if shiftOp.ShiftValueUsed {
			var err error
			val2, err = applyShift(val2, shiftOp)
			if err != nil {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
			}
		}
	}

	// Perform addition
	result := val1 + val2

	// Handle 32-bit vs 64-bit operations
	if len(dstOp.Registers) > 0 && uint32(dstOp.Registers[0]) <= 31 {
		// W register (32-bit operation)
		result = uint64(uint32(result))
	}

	// Set result
	setDstRegister(result)

	return nil
}

func (e *ArithmeticExecutor) executeADDS(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ADDS requires at least 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get destination register
	var setDstRegister func(uint64)
	if len(dstOp.Registers) > 0 {
		dstRegID := dstOp.Registers[0]
		if uint32(dstRegID) == 31 { // Register 31 can be SP or XZR depending on context
			// In ADDS, if destination is 31, it's XZR (discard result)
			setDstRegister = func(val uint64) {} // No-op for XZR
		} else {
			dstReg := core.MapRegister(dstRegID)
			if dstReg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid destination register")
			}
			setDstRegister = func(val uint64) { state.SetX(dstReg, val) }
		}
	} else {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "no destination register")
	}

	// Get first operand value
	var val1 uint64
	if len(src1Op.Registers) > 0 {
		src1RegID := src1Op.Registers[0]
		if uint32(src1RegID) == 31 { // Register 31 can be SP or XZR depending on context
			val1 = 0 // XZR reads as zero
		} else {
			src1Reg := core.MapRegister(src1RegID)
			if src1Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid source register")
			}
			val1 = state.GetX(src1Reg)
		}
	}

	// Get second operand (register or immediate)
	var val2 uint64
	if len(src2Op.Registers) > 0 {
		src2RegID := src2Op.Registers[0]
		if uint32(src2RegID) == 31 { // Register 31 can be SP or XZR depending on context
			val2 = 0 // XZR reads as zero
		} else {
			src2Reg := core.MapRegister(src2RegID)
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid second source register")
			}
			val2 = state.GetX(src2Reg)
		}
	} else {
		// Immediate operand - check for optional LSL #12 shift
		val2 = uint64(src2Op.Immediate)
		if (instr.Raw>>22)&1 == 1 { // Check the 'sh' bit
			val2 <<= 12
		}
	}

	// Handle shift/extend on second operand
	if src2Op.ShiftValueUsed {
		var err error
		val2, err = applyShift(val2, src2Op)
		if err != nil {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
		}
	}

	// Handle optional shift/extend in 4th operand (for extended registers)
	if len(instr.Operands) > 3 {
		shiftOp := instr.Operands[3]
		if shiftOp.ShiftValueUsed {
			var err error
			val2, err = applyShift(val2, shiftOp)
			if err != nil {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
			}
		}
	}

	// Perform addition with overflow detection
	result, carry, overflow := addWithFlags(val1, val2)

	// Handle 32-bit vs 64-bit operations
	if len(dstOp.Registers) > 0 && uint32(dstOp.Registers[0]) <= 31 {
		// W register (32-bit operation)
		result = uint64(uint32(result))
		// Recalculate flags for 32-bit result
		carry = (uint32(val1) + uint32(val2)) < uint32(val1)
		overflow = ((val1 ^ result) & (val2 ^ result) & 0x80000000) != 0
	}

	// Set result
	setDstRegister(result)

	// Update flags
	state.SetZ(result == 0)
	state.SetN((result & 0x8000000000000000) != 0)
	state.SetC(carry)
	state.SetV(overflow)

	return nil
}

func (e *ArithmeticExecutor) executeSUB(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SUB requires at least 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get destination register
	var setDstRegister func(uint64)
	if len(dstOp.Registers) > 0 {
		dstRegID := dstOp.Registers[0]
		reg := uint32(dstRegID)
		if reg == 66 { // SP
			setDstRegister = func(val uint64) { state.SetSP(val) }
		} else if reg == 31 || reg == 65 { // WZR/XZR
			setDstRegister = func(val uint64) {}
		} else {
			dstReg := core.MapRegister(dstRegID)
			if dstReg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid destination register: %d", reg))
			}
			setDstRegister = func(val uint64) { state.SetX(dstReg, val) }
		}
	} else {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "no destination register")
	}

	// Get first operand value
	var val1 uint64
	if len(src1Op.Registers) > 0 {
		src1RegID := src1Op.Registers[0]
		reg := uint32(src1RegID)
		if reg == 66 { // SP
			val1 = state.GetSP()
		} else if reg == 31 || reg == 65 { // WZR/XZR
			val1 = 0
		} else {
			src1Reg := core.MapRegister(src1RegID)
			if src1Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid source register: %d", reg))
			}
			val1 = state.GetX(src1Reg)
		}
	}

	// Get second operand (register or immediate)
	var val2 uint64
	if len(src2Op.Registers) > 0 {
		src2RegID := src2Op.Registers[0]
		reg := uint32(src2RegID)
		if reg == 66 { // SP
			val2 = state.GetSP()
		} else if reg == 31 || reg == 65 { // WZR/XZR
			val2 = 0
		} else {
			src2Reg := core.MapRegister(src2RegID)
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("invalid second source register: %d", reg))
			}
			val2 = state.GetX(src2Reg)
		}
	} else {
		// Immediate operand - check for optional LSL #12 shift
		val2 = uint64(src2Op.Immediate)
		if (instr.Raw>>22)&1 == 1 { // Check the 'sh' bit
			val2 <<= 12
		}
	}

	// Handle shift/extend on second operand
	if src2Op.ShiftValueUsed {
		var err error
		val2, err = applyShift(val2, src2Op)
		if err != nil {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
		}
	}

	// Handle optional shift/extend in 4th operand (for extended registers)
	if len(instr.Operands) > 3 {
		shiftOp := instr.Operands[3]
		if shiftOp.ShiftValueUsed {
			var err error
			val2, err = applyShift(val2, shiftOp)
			if err != nil {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
			}
		}
	}

	// Perform subtraction
	result := val1 - val2

	// Handle 32-bit vs 64-bit operations
	if len(dstOp.Registers) > 0 && uint32(dstOp.Registers[0]) <= 31 {
		// W register (32-bit operation)
		result = uint64(uint32(result))
	}

	// Set result
	setDstRegister(result)

	return nil
}

func (e *ArithmeticExecutor) executeSUBS(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SUBS requires at least 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get destination register
	var setDstRegister func(uint64)
	if len(dstOp.Registers) > 0 {
		dstRegID := dstOp.Registers[0]
		if uint32(dstRegID) == 31 { // Register 31 can be SP or XZR depending on context
			// In SUBS, if destination is 31, it's XZR (discard result)
			setDstRegister = func(val uint64) {} // No-op for XZR
		} else {
			dstReg := core.MapRegister(dstRegID)
			if dstReg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid destination register")
			}
			setDstRegister = func(val uint64) { state.SetX(dstReg, val) }
		}
	} else {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "no destination register")
	}

	// Get first operand value
	var val1 uint64
	if len(src1Op.Registers) > 0 {
		src1RegID := src1Op.Registers[0]
		if uint32(src1RegID) == 31 { // Register 31 can be SP or XZR depending on context
			val1 = 0 // XZR reads as zero
		} else {
			src1Reg := core.MapRegister(src1RegID)
			if src1Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid source register")
			}
			val1 = state.GetX(src1Reg)
		}
	}

	// Get second operand (register or immediate)
	var val2 uint64
	if len(src2Op.Registers) > 0 {
		src2RegID := src2Op.Registers[0]
		if uint32(src2RegID) == 31 { // Register 31 can be SP or XZR depending on context
			val2 = 0 // XZR reads as zero
		} else {
			src2Reg := core.MapRegister(src2RegID)
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid second source register")
			}
			val2 = state.GetX(src2Reg)
		}
	} else {
		// Immediate operand - check for optional LSL #12 shift
		val2 = uint64(src2Op.Immediate)
		if (instr.Raw>>22)&1 == 1 { // Check the 'sh' bit
			val2 <<= 12
		}
	}

	// Handle shift/extend on second operand
	if src2Op.ShiftValueUsed {
		var err error
		val2, err = applyShift(val2, src2Op)
		if err != nil {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
		}
	}

	// Handle optional shift/extend in 4th operand (for extended registers)
	if len(instr.Operands) > 3 {
		shiftOp := instr.Operands[3]
		if shiftOp.ShiftValueUsed {
			var err error
			val2, err = applyShift(val2, shiftOp)
			if err != nil {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("shift error: %v", err))
			}
		}
	}

	// Perform subtraction with overflow detection
	result, carry, overflow := subWithFlags(val1, val2)

	// Handle 32-bit vs 64-bit operations
	if len(dstOp.Registers) > 0 && uint32(dstOp.Registers[0]) <= 31 {
		// W register (32-bit operation)
		result = uint64(uint32(result))
		// Recalculate flags for 32-bit result
		carry = uint32(val1) >= uint32(val2)
		overflow = ((val1 ^ val2) & (val1 ^ result) & 0x80000000) != 0
	}

	// Set result
	setDstRegister(result)

	// Update flags
	state.SetZ(result == 0)
	state.SetN((result & 0x8000000000000000) != 0)
	state.SetC(carry)
	state.SetV(overflow)

	return nil
}

func (e *ArithmeticExecutor) executeMUL(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MUL requires exactly 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in MUL")
	}

	// Handle XZR (zero register) - reads as 0
	var val1, val2 uint64
	if uint32(src1Op.Registers[0]) == 65 || uint32(src1Op.Registers[0]) == 32 { // XZR or WZR
		val1 = 0
	} else {
		val1 = state.GetX(src1Reg)
	}

	if uint32(src2Op.Registers[0]) == 65 || uint32(src2Op.Registers[0]) == 32 { // XZR or WZR
		val2 = 0
	} else {
		val2 = state.GetX(src2Reg)
	}

	result := val1 * val2

	// Handle XZR destination (ignore writes to zero register)
	if uint32(dstOp.Registers[0]) == 65 || uint32(dstOp.Registers[0]) == 32 { // XZR or WZR
		// Writes to XZR/WZR are ignored
		return nil
	}

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *ArithmeticExecutor) executeUDIV(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UDIV requires exactly 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in UDIV")
	}

	val1 := state.GetX(src1Reg)
	val2 := state.GetX(src2Reg)

	if val2 == 0 {
		// ARM64 UDIV returns 0 for division by zero (no exception)
		result := uint64(0)
		if uint32(dstOp.Registers[0]) <= 31 {
			state.SetW(dstReg, uint32(result))
		} else {
			state.SetX(dstReg, result)
		}
		return nil
	}

	result := val1 / val2

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *ArithmeticExecutor) executeSDIV(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SDIV requires exactly 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in SDIV")
	}

	val1 := int64(state.GetX(src1Reg))
	val2 := int64(state.GetX(src2Reg))

	if val2 == 0 {
		// ARM64 SDIV returns 0 for division by zero (no exception)
		result := int64(0)
		if uint32(dstOp.Registers[0]) <= 31 {
			state.SetW(dstReg, uint32(int32(result)))
		} else {
			state.SetX(dstReg, uint64(result))
		}
		return nil
	}

	result := val1 / val2

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(int32(result)))
	} else {
		state.SetX(dstReg, uint64(result))
	}

	return nil
}

func (e *ArithmeticExecutor) executeMADD(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MADD requires exactly 4 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]   // Rn
	src2Op := instr.Operands[2]   // Rm
	addendOp := instr.Operands[3] // Ra

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])
	addendReg := core.MapRegister(addendOp.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 || addendReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in MADD")
	}

	// Handle XZR (zero register) - reads as 0
	var val1, val2, val3 uint64
	if uint32(src1Op.Registers[0]) == 65 || uint32(src1Op.Registers[0]) == 32 { // XZR or WZR
		val1 = 0
	} else {
		val1 = state.GetX(src1Reg)
	}

	if uint32(src2Op.Registers[0]) == 65 || uint32(src2Op.Registers[0]) == 32 { // XZR or WZR
		val2 = 0
	} else {
		val2 = state.GetX(src2Reg)
	}

	if uint32(addendOp.Registers[0]) == 65 || uint32(addendOp.Registers[0]) == 32 { // XZR or WZR
		val3 = 0
	} else {
		val3 = state.GetX(addendReg)
	}

	// MADD: Rd = Ra + (Rn * Rm)
	result := val3 + (val1 * val2)

	// Handle XZR destination (ignore writes to zero register)
	if uint32(dstOp.Registers[0]) == 65 || uint32(dstOp.Registers[0]) == 32 { // XZR or WZR
		// Writes to XZR/WZR are ignored
		return nil
	}

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *ArithmeticExecutor) executeMSUB(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "MSUB requires exactly 4 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]    // Rn
	src2Op := instr.Operands[2]    // Rm
	minuendOp := instr.Operands[3] // Ra

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])
	minuendReg := core.MapRegister(minuendOp.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 || minuendReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in MSUB")
	}

	// Handle XZR (zero register) - reads as 0
	var val1, val2, val3 uint64
	if uint32(src1Op.Registers[0]) == 65 || uint32(src1Op.Registers[0]) == 32 { // XZR or WZR
		val1 = 0
	} else {
		val1 = state.GetX(src1Reg)
	}

	if uint32(src2Op.Registers[0]) == 65 || uint32(src2Op.Registers[0]) == 32 { // XZR or WZR
		val2 = 0
	} else {
		val2 = state.GetX(src2Reg)
	}

	if uint32(minuendOp.Registers[0]) == 65 || uint32(minuendOp.Registers[0]) == 32 { // XZR or WZR
		val3 = 0
	} else {
		val3 = state.GetX(minuendReg)
	}

	// MSUB: Rd = Ra - (Rn * Rm)
	result := val3 - (val1 * val2)

	// Handle XZR destination (ignore writes to zero register)
	if uint32(dstOp.Registers[0]) == 65 || uint32(dstOp.Registers[0]) == 32 { // XZR or WZR
		// Writes to XZR/WZR are ignored
		return nil
	}

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *ArithmeticExecutor) executeSMULL(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "SMULL requires exactly 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in SMULL")
	}

	// SMULL multiplies two 32-bit signed values to produce 64-bit result
	val1 := int32(state.GetW(src1Reg))
	val2 := int32(state.GetW(src2Reg))
	result := int64(val1) * int64(val2)

	state.SetX(dstReg, uint64(result))
	return nil
}

func (e *ArithmeticExecutor) executeUMULL(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UMULL requires exactly 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	src2Reg := core.MapRegister(src2Op.Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in UMULL")
	}

	// UMULL multiplies two 32-bit unsigned values to produce 64-bit result
	val1 := uint64(state.GetW(src1Reg))
	val2 := uint64(state.GetW(src2Reg))
	result := val1 * val2

	state.SetX(dstReg, result)
	return nil
}

func (e *ArithmeticExecutor) executeNEG(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "NEG requires exactly 2 operands")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in NEG")
	}

	val := state.GetX(srcReg)
	result := uint64(-int64(val))

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *ArithmeticExecutor) executeNEGS(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) != 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "NEGS requires exactly 2 operands")
	}

	dstOp := instr.Operands[0]
	srcOp := instr.Operands[1]

	// Get register indices
	dstReg := core.MapRegister(dstOp.Registers[0])
	srcReg := core.MapRegister(srcOp.Registers[0])

	if dstReg == -1 || srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register in NEGS")
	}

	val := state.GetX(srcReg)

	// NEG is equivalent to SUB from 0
	result, carry, overflow := subWithFlags(0, val)

	// Handle 32-bit vs 64-bit operations
	if uint32(dstOp.Registers[0]) <= 31 {
		result = uint64(uint32(result))
		// Recalculate flags for 32-bit result
		carry = uint32(0) >= uint32(val)
		overflow = ((0 ^ val) & (0 ^ result) & 0x80000000) != 0
	}

	// Set result
	if uint32(dstOp.Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	// Update flags
	state.SetZ(result == 0)
	state.SetN((result & 0x8000000000000000) != 0)
	state.SetC(carry)
	state.SetV(overflow)

	return nil
}

// Helper functions for arithmetic operations

// addWithFlags performs addition and returns carry and overflow flags
func addWithFlags(a, b uint64) (result uint64, carry bool, overflow bool) {
	result = a + b
	carry = result < a // Unsigned overflow

	// Signed overflow: result and operands have different signs
	// This occurs when both operands have the same sign but result has different sign
	signA := a & 0x8000000000000000
	signB := b & 0x8000000000000000
	signResult := result & 0x8000000000000000
	overflow = (signA == signB) && (signA != signResult)

	return
}

// subWithFlags performs subtraction and returns carry and overflow flags
func subWithFlags(a, b uint64) (result uint64, carry bool, overflow bool) {
	result = a - b
	carry = a >= b // No borrow (inverse of typical subtraction carry)

	// Signed overflow for subtraction
	signA := a & 0x8000000000000000
	signB := b & 0x8000000000000000
	signResult := result & 0x8000000000000000
	overflow = (signA != signB) && (signA != signResult)

	return
}

// applyShift applies shift operations to a value
func applyShift(value uint64, shiftOp disassemble.Operand) (uint64, error) {
	if !shiftOp.ShiftValueUsed {
		return value, nil
	}

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

// executeADDG executes ADDG (Add with Tag Generation) instructions
// For our purposes, we treat it as regular ADD since we're not implementing full memory tagging
func (e *ArithmeticExecutor) executeADDG(state core.State, instr *disassemble.Instruction) error {
	return e.executeADD(state, instr)
}

// RegisterArithmeticInstructions registers all arithmetic instructions
func RegisterArithmeticInstructions(registry *Registry) {
	// Basic arithmetic
	registry.Register("ADD", NewArithmeticExecutor("ADD", "Add (immediate/register)"))
	registry.Register("ADDG", NewArithmeticExecutor("ADDG", "Add with Tag Generation"))
	registry.Register("ADDS", NewArithmeticExecutor("ADDS", "Add and set flags"))
	registry.Register("SUB", NewArithmeticExecutor("SUB", "Subtract (immediate/register)"))
	registry.Register("SUBS", NewArithmeticExecutor("SUBS", "Subtract and set flags"))

	// Multiplication and division
	registry.Register("MUL", NewArithmeticExecutor("MUL", "Multiply"))
	registry.Register("UDIV", NewArithmeticExecutor("UDIV", "Unsigned divide"))
	registry.Register("SDIV", NewArithmeticExecutor("SDIV", "Signed divide"))

	// Multiply-accumulate
	registry.Register("MADD", NewArithmeticExecutor("MADD", "Multiply-add"))
	registry.Register("MSUB", NewArithmeticExecutor("MSUB", "Multiply-subtract"))

	// Long multiply
	registry.Register("SMULL", NewArithmeticExecutor("SMULL", "Signed multiply long"))
	registry.Register("UMULL", NewArithmeticExecutor("UMULL", "Unsigned multiply long"))

	// Negation
	registry.Register("NEG", NewArithmeticExecutor("NEG", "Negate"))
	registry.Register("NEGS", NewArithmeticExecutor("NEGS", "Negate and set flags"))
}
