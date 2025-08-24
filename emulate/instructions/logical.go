package instructions

import (
	"fmt"
	"math/bits"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// rotateRight rotates value right by r within the given width (32 or 64)
func rotateRight(value uint64, r uint32, width int) uint64 {
	if width == 32 {
		r %= 32
		v := uint32(value)
		return uint64((v >> r) | (v << ((32 - r) & 31)))
	}
	r %= 64
	return (value >> r) | (value << ((64 - r) & 63))
}

// maskOnes returns a mask with the lowest 'width' bits set
func maskOnes(width int) uint64 {
	if width >= 64 {
		return ^uint64(0)
	}
	if width <= 0 {
		return 0
	}
	return (uint64(1) << width) - 1
}

// insertField inserts 'field' (width bits) at bit position lsb, wrapping around modulo size bits
func insertField(field uint64, lsb int, width int, size int) uint64 {
	field &= maskOnes(width)
	if width == size {
		return field // entire register
	}
	if lsb+width <= size {
		return field << lsb
	}
	// Wrapping case
	topWidth := size - lsb
	part1 := (field & maskOnes(topWidth)) << lsb
	part2 := field >> uint(topWidth) // remaining low bits go at bottom
	return part1 | (part2 & maskOnes((lsb+width)%size))
}

// Replicate repeats a bit pattern across 64 bits.
func replicate(pattern uint64, esize int) uint64 {
	if esize >= 64 {
		return pattern
	}
	result := uint64(0)
	for i := 0; i < 64; i += esize {
		result |= pattern << i
	}
	return result
}

// decodeBitMasks decodes logical-immediate masks (ARM DecodeBitMasks)
func decodeBitMasks(n, imms, immr, sf uint32) (wmask uint64, tmask uint64, ok bool) {
	// Determine element width
	width := 64
	if sf == 0 {
		width = 32
		if n != 0 { // immN must be 0 for 32-bit encodings
			return 0, 0, false
		}
	}

	// len = HighestSetBit(immN:NOT(imms))
	notImms := (^imms) & 0x3f
	temp := (n << 6) | notImms
	if temp == 0 {
		return 0, 0, false
	}
	len := 31 - bits.LeadingZeros32(uint32(temp))
	if len < 1 {
		return 0, 0, false
	}

	levels := uint32((1 << uint(len)) - 1)
	S := imms & levels
	R := immr & levels
	esize := 1 << uint(len)

	// Build base element of S+1 ones
	elemMask := (uint64(1) << (uint(S) + 1)) - 1
	// Rotate element right by R within esize
	r := R % uint32(esize)
	if r != 0 {
		low := elemMask & ((1 << r) - 1)
		elemMask = (elemMask >> r) | (low << (uint(esize) - uint(r)))
	}
	// Replicate rotated element across full width
	wmask = 0
	for i := 0; i < width; i += int(esize) {
		wmask |= elemMask << i
	}
	// Test mask is unrotated element replicated
	baseElem := (uint64(1) << (uint(S) + 1)) - 1
	tmask = 0
	for i := 0; i < width; i += int(esize) {
		tmask |= baseElem << i
	}

	if width == 32 {
		wmask &= 0xffffffff
		tmask &= 0xffffffff
	}
	return wmask, tmask, true
}

// LogicalExecutor handles logical operations (AND, ORR, EOR, etc.)
type LogicalExecutor struct {
	*BaseExecutor
}

// NewLogicalExecutor creates a new logical instruction executor
func NewLogicalExecutor(mnemonic, description string) *LogicalExecutor {
	return &LogicalExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// Execute executes logical instructions
func (e *LogicalExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.mnemonic {
	case "AND":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a & b }, false)
	case "ANDS":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a & b }, true)
	case "ORR":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a | b }, false)
	case "EOR":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a ^ b }, false)
	case "BIC":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a &^ b }, false)
	case "BICS":
		return e.executeLogicalOp(state, instr, func(a, b uint64) uint64 { return a &^ b }, true)
	case "TST":
		return e.executeTST(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("logical instruction %s not implemented", e.mnemonic))
	}
}

// executeLogicalOp performs a generic logical operation
func (e *LogicalExecutor) executeLogicalOp(state core.State, instr *disassemble.Instruction, op func(uint64, uint64) uint64, setFlags bool) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "logical operation requires at least 3 operands")
	}

	dstOp := instr.Operands[0]
	src1Op := instr.Operands[1]
	src2Op := instr.Operands[2]

	if len(dstOp.Registers) == 0 || len(src1Op.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "missing destination or source register")
	}

	dstReg := core.MapRegister(dstOp.Registers[0])
	src1Reg := core.MapRegister(src1Op.Registers[0])
	if dstReg == -1 || src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	val1 := state.GetX(src1Reg)
	var val2 uint64

	if len(src2Op.Registers) > 0 {
		src2Reg := core.MapRegister(src2Op.Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second source register")
		}
		val2 = state.GetX(src2Reg)
		// Apply optional shift to register operand
		if src2Op.ShiftValueUsed {
			switch src2Op.ShiftType {
			case disassemble.SHIFT_TYPE_LSL:
				amount := uint64(src2Op.ShiftValue)
				if amount >= 64 {
					val2 = 0
				} else {
					val2 = val2 << amount
				}
			case disassemble.SHIFT_TYPE_LSR:
				amount := uint64(src2Op.ShiftValue)
				if amount >= 64 {
					val2 = 0
				} else {
					val2 = val2 >> amount
				}
			case disassemble.SHIFT_TYPE_ASR:
				amount := uint64(src2Op.ShiftValue)
				if amount >= 64 {
					if (val2 & 0x8000000000000000) != 0 {
						val2 = 0xFFFFFFFFFFFFFFFF
					} else {
						val2 = 0
					}
				} else {
					val2 = uint64(int64(val2) >> amount)
				}
			case disassemble.SHIFT_TYPE_ROR:
				val2 = rotateRight(val2, uint32(src2Op.ShiftValue), 64)
			}
		}
	} else {
		// Immediate form.
		// If operand explicitly provides an immediate (unit tests), use it directly.
		if src2Op.Class == disassemble.IMM32 || src2Op.Class == disassemble.IMM64 || src2Op.SignedImm || src2Op.Immediate != 0 {
			val2 = uint64(src2Op.Immediate)
		} else if instr.Raw != 0 { // try to decode logical-immediate mask from encoding
			is32Bit := uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31
			sf := uint32(0)
			if !is32Bit {
				sf = 1
			}
			n := (instr.Raw >> 22) & 1
			immr := (instr.Raw >> 16) & 0x3f
			imms := (instr.Raw >> 10) & 0x3f
			if wmask, _, ok := decodeBitMasks(n, imms, immr, sf); ok {
				val2 = wmask
			} else {
				val2 = 0
			}
		} else {
			// No register, no explicit immediate, and no encoding to decode -> treat as zero immediate
			val2 = 0
		}
	}

	result := op(val1, val2)

	is32Bit := uint32(dstOp.Registers[0]) >= 1 && uint32(dstOp.Registers[0]) <= 31
	if is32Bit {
		result &= 0xffffffff
		val1 &= 0xffffffff
		// Ensure immediate logical masks are also 32-bit when destination is W-reg
		val2 &= 0xffffffff
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	if setFlags {
		if is32Bit {
			state.SetN((result & 0x80000000) != 0)
		} else {
			state.SetN((result & 0x8000000000000000) != 0)
		}
		state.SetZ(result == 0)
		state.SetC(false)
		state.SetV(false)
	}

	return nil
}

func (e *LogicalExecutor) executeTST(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TST requires 2 operands")
	}

	src1Op := instr.Operands[0]
	src2Op := instr.Operands[1]

	if len(src1Op.Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "missing source register")
	}
	src1Reg := core.MapRegister(src1Op.Registers[0])
	if src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	val1 := state.GetX(src1Reg)
	var val2 uint64
	if len(src2Op.Registers) > 0 {
		src2Reg := core.MapRegister(src2Op.Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second source register")
		}
		val2 = state.GetX(src2Reg)
	} else {
		// Immediate form
		if src2Op.Class == disassemble.IMM32 || src2Op.Class == disassemble.IMM64 || src2Op.Immediate != 0 || src2Op.SignedImm {
			val2 = uint64(src2Op.Immediate)
		} else {
			is32Bit := uint32(src1Op.Registers[0]) >= 1 && uint32(src1Op.Registers[0]) <= 31
			sf := uint32(0)
			if !is32Bit {
				sf = 1
			}
			n := (instr.Raw >> 22) & 1
			immr := (instr.Raw >> 16) & 0x3f
			imms := (instr.Raw >> 10) & 0x3f
			wmask, _, ok := decodeBitMasks(n, imms, immr, sf)
			if !ok {
				return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid logical immediate encoding")
			}
			val2 = wmask
		}
	}

	result := val1 & val2
	is32Bit := uint32(src1Op.Registers[0]) >= 1 && uint32(src1Op.Registers[0]) <= 31
	if is32Bit {
		result &= 0xffffffff
		state.SetN((result & 0x80000000) != 0)
	} else {
		state.SetN((result & 0x8000000000000000) != 0)
	}
	state.SetZ(result == 0)
	state.SetC(false)
	state.SetV(false)
	return nil
}

// BitfieldExecutor handles bitfield operations (BFM, SBFM, UBFM and their aliases)
type BitfieldExecutor struct {
	*BaseExecutor
}

// NewBitfieldExecutor creates a new bitfield instruction executor
func NewBitfieldExecutor(mnemonic, description string) *BitfieldExecutor {
	return &BitfieldExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// Execute executes bitfield instructions
func (e *BitfieldExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	instrValue := instr.Raw

	switch e.mnemonic {
	case "BFM":
		return e.executeBFM(state, instr, instrValue)
	case "BFI":
		return e.executeBFI(state, instr)
	case "BFXIL":
		return e.executeBFXIL(state, instr)
	case "SBFM", "SBFIZ", "SBFX":
		return e.executeSBFM(state, instr, instrValue)
	case "UBFM":
		return e.executeUBFM(state, instr, instrValue)
	case "UBFIZ":
		return e.executeUBFIZ(state, instr)
	case "UBFX":
		return e.executeUBFX(state, instr)
	case "LSL":
		return e.executeLSL(state, instr)
	case "LSR":
		return e.executeLSR(state, instr)
	case "ASR":
		return e.executeASR(state, instr)
	case "UXTB":
		// Zero extend byte
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "UXTB requires dst, src")
		}
		dst := core.MapRegister(instr.Operands[0].Registers[0])
		src := core.MapRegister(instr.Operands[1].Registers[0])
		if dst == -1 || src == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid register")
		}
		val := state.GetX(src) & 0xFF
		if is32bitOp(instr.Operands[0]) {
			state.SetW(dst, uint32(val))
		} else {
			state.SetX(dst, val)
		}
		return nil
	case "UXTH":
		// Zero extend halfword
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "UXTH requires dst, src")
		}
		dst := core.MapRegister(instr.Operands[0].Registers[0])
		src := core.MapRegister(instr.Operands[1].Registers[0])
		if dst == -1 || src == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid register")
		}
		val := state.GetX(src) & 0xFFFF
		if is32bitOp(instr.Operands[0]) {
			state.SetW(dst, uint32(val))
		} else {
			state.SetX(dst, val)
		}
		return nil
	case "SXTB":
		// Sign extend byte
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "SXTB requires dst, src")
		}
		dst := core.MapRegister(instr.Operands[0].Registers[0])
		src := core.MapRegister(instr.Operands[1].Registers[0])
		if dst == -1 || src == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid register")
		}
		b := int8(state.GetX(src) & 0xFF)
		state.SetX(dst, uint64(int64(b)))
		return nil
	case "SXTH":
		// Sign extend halfword
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "SXTH requires dst, src")
		}
		dst := core.MapRegister(instr.Operands[0].Registers[0])
		src := core.MapRegister(instr.Operands[1].Registers[0])
		if dst == -1 || src == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid register")
		}
		h := int16(state.GetX(src) & 0xFFFF)
		state.SetX(dst, uint64(int64(h)))
		return nil
	case "SXTW":
		// Sign extend word
		if len(instr.Operands) < 2 {
			return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "SXTW requires dst, src")
		}
		dst := core.MapRegister(instr.Operands[0].Registers[0])
		src := core.MapRegister(instr.Operands[1].Registers[0])
		if dst == -1 || src == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(), fmt.Sprintf("%v", instr.Operation), "invalid register")
		}
		w := int32(state.GetX(src) & 0xFFFFFFFF)
		state.SetX(dst, uint64(int64(w)))
		return nil
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("bitfield instruction %s not implemented", e.mnemonic))
	}
}

// helpers for alias ops
func getDstSrcRegs(state core.State, instr *disassemble.Instruction) (int, int, bool) {
	if len(instr.Operands) < 2 {
		return -1, -1, false
	}
	dst := core.MapRegister(instr.Operands[0].Registers[0])
	src := core.MapRegister(instr.Operands[1].Registers[0])
	if dst == -1 || src == -1 {
		return -1, -1, false
	}
	return dst, src, true
}

func is32bitOp(op disassemble.Operand) bool {
	reg := uint32(op.Registers[0])
	return reg >= 1 && reg <= 31
}

func immOf(op disassemble.Operand) uint64 {
	return uint64(op.Immediate)
}

func (e *BitfieldExecutor) executeBFM(state core.State, instr *disassemble.Instruction, instrValue uint32) error {
	dstReg := core.MapRegister(instr.Operands[0].Registers[0])
	srcReg := core.MapRegister(instr.Operands[1].Registers[0])

	sf := (instrValue >> 31) & 1
	immr := (instrValue >> 16) & 0x3f
	imms := (instrValue >> 10) & 0x3f

	size := 64
	if sf == 0 {
		size = 32
	}
	R := int(immr & uint32(size-1))
	S := int(imms & uint32(size-1))
	var width int
	if S >= R {
		width = S - R + 1
	} else {
		width = size - R + S + 1
	}

	srcVal := state.GetX(srcReg)
	dstVal := state.GetX(dstReg)
	rot := rotateRight(srcVal, uint32(R), size)
	field := rot & maskOnes(width)
	mask := insertField(maskOnes(width), R, width, size)
	inserted := insertField(field, R, width, size)
	result := (dstVal & ^mask) | (inserted & mask)

	if sf == 0 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *BitfieldExecutor) executeBFI(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BFI requires dst, src, lsb, width")
	}
	lsb := uint(immOf(instr.Operands[2]))
	width := uint(immOf(instr.Operands[3]))
	if width == 0 || width > 64 || lsb >= 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid lsb/width")
	}
	mask := ((uint64(1) << width) - 1) << lsb
	srcVal := state.GetX(src)
	insert := (srcVal & ((uint64(1) << width) - 1)) << lsb
	dstVal := state.GetX(dst)
	dstVal = (dstVal & ^mask) | (insert & mask)
	if is32bitOp(instr.Operands[0]) {
		state.SetW(dst, uint32(dstVal))
	} else {
		state.SetX(dst, dstVal)
	}
	return nil
}

func (e *BitfieldExecutor) executeBFXIL(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BFXIL requires dst, src, lsb, width")
	}
	lsb := uint(immOf(instr.Operands[2]))
	width := uint(immOf(instr.Operands[3]))
	if width == 0 || width > 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid width")
	}
	mask := uint64((uint64(1) << width) - 1)
	srcVal := state.GetX(src)
	insert := (srcVal >> lsb) & mask
	dstVal := state.GetX(dst)
	dstVal = (dstVal & ^mask) | insert
	if is32bitOp(instr.Operands[0]) {
		state.SetW(dst, uint32(dstVal))
	} else {
		state.SetX(dst, dstVal)
	}
	return nil
}

func (e *BitfieldExecutor) executeSBFM(state core.State, instr *disassemble.Instruction, instrValue uint32) error {
	dstReg := core.MapRegister(instr.Operands[0].Registers[0])
	srcReg := core.MapRegister(instr.Operands[1].Registers[0])

	sf := (instrValue >> 31) & 1
	immr := (instrValue >> 16) & 0x3f
	imms := (instrValue >> 10) & 0x3f

	size := 64
	if sf == 0 {
		size = 32
	}
	R := int(immr & uint32(size-1))
	S := int(imms & uint32(size-1))
	var width int
	if S >= R {
		width = S - R + 1
	} else {
		width = size - R + S + 1
	}

	srcVal := state.GetX(srcReg)
	rot := rotateRight(srcVal, uint32(R), size)
	result := rot & maskOnes(width)
	// Sign-extend from 'width'
	if width > 0 {
		signBit := (result >> uint(width-1)) & 1
		if signBit == 1 {
			result |= ^maskOnes(width)
		}
	}

	if sf == 0 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *BitfieldExecutor) executeUBFM(state core.State, instr *disassemble.Instruction, instrValue uint32) error {
	dstReg := core.MapRegister(instr.Operands[0].Registers[0])
	srcReg := core.MapRegister(instr.Operands[1].Registers[0])

	sf := (instrValue >> 31) & 1
	immr := (instrValue >> 16) & 0x3f
	imms := (instrValue >> 10) & 0x3f

	size := 64
	if sf == 0 {
		size = 32
	}
	R := int(immr & uint32(size-1))
	S := int(imms & uint32(size-1))
	var width int
	if S >= R {
		width = S - R + 1
	} else {
		width = size - R + S + 1
	}

	srcVal := state.GetX(srcReg)
	rot := rotateRight(srcVal, uint32(R), size)
	result := rot & maskOnes(width)

	if sf == 0 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

func (e *BitfieldExecutor) executeUBFIZ(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UBFIZ requires dst, src, lsb, width")
	}
	lsb := uint(immOf(instr.Operands[2]))
	width := uint(immOf(instr.Operands[3]))
	if width == 0 || width > 64 || lsb >= 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid lsb/width")
	}
	mask := (uint64(1) << width) - 1
	srcVal := state.GetX(src)
	result := (srcVal & mask) << lsb
	if is32bitOp(instr.Operands[0]) {
		state.SetW(dst, uint32(result))
	} else {
		state.SetX(dst, result)
	}
	return nil
}

func (e *BitfieldExecutor) executeUBFX(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "UBFX requires dst, src, lsb, width")
	}
	lsb := uint(immOf(instr.Operands[2]))
	width := uint(immOf(instr.Operands[3]))
	if width == 0 || width > 64 || lsb >= 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid lsb/width")
	}
	mask := (uint64(1) << width) - 1
	srcVal := state.GetX(src)
	result := (srcVal >> lsb) & mask
	if is32bitOp(instr.Operands[0]) {
		state.SetW(dst, uint32(result))
	} else {
		state.SetX(dst, result)
	}
	return nil
}

func (e *BitfieldExecutor) executeLSL(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LSL requires dst, src, shift")
	}
	shift := uint(immOf(instr.Operands[2]))
	srcVal := state.GetX(src)
	var result uint64
	if is32bitOp(instr.Operands[0]) {
		shift &= 31
		result = uint64(uint32(srcVal) << shift)
		state.SetW(dst, uint32(result))
	} else {
		shift &= 63
		result = srcVal << shift
		state.SetX(dst, result)
	}
	return nil
}

func (e *BitfieldExecutor) executeLSR(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LSR requires dst, src, shift")
	}
	shift := uint(immOf(instr.Operands[2]))
	srcVal := state.GetX(src)
	var result uint64
	if is32bitOp(instr.Operands[0]) {
		shift &= 31
		result = uint64(uint32(srcVal) >> shift)
		state.SetW(dst, uint32(result))
	} else {
		shift &= 63
		result = srcVal >> shift
		state.SetX(dst, result)
	}
	return nil
}

func (e *BitfieldExecutor) executeASR(state core.State, instr *disassemble.Instruction) error {
	dst, src, ok := getDstSrcRegs(state, instr)
	if !ok || len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "ASR requires dst, src, shift")
	}
	shift := uint(immOf(instr.Operands[2]))
	srcVal := state.GetX(src)
	var result uint64
	if is32bitOp(instr.Operands[0]) {
		shift &= 31
		result = uint64(int32(uint32(srcVal)) >> shift)
		state.SetW(dst, uint32(result))
	} else {
		shift &= 63
		result = uint64(int64(srcVal) >> shift)
		state.SetX(dst, result)
	}
	return nil
}

// RegisterLogicalInstructions registers all logical instructions
func RegisterLogicalInstructions(registry *Registry) {
	// Basic logical operations
	registry.Register("AND", NewLogicalExecutor("AND", "Bitwise AND"))
	registry.Register("ANDS", NewLogicalExecutor("ANDS", "Bitwise AND and set flags"))
	registry.Register("ORR", NewLogicalExecutor("ORR", "Bitwise OR"))
	registry.Register("EOR", NewLogicalExecutor("EOR", "Bitwise exclusive OR"))
	registry.Register("BIC", NewLogicalExecutor("BIC", "Bit clear (AND NOT)"))
	registry.Register("BICS", NewLogicalExecutor("BICS", "Bit clear and set flags"))
	registry.Register("TST", NewLogicalExecutor("TST", "Test bits"))

	// Bitfield operations
	registry.Register("BFM", NewBitfieldExecutor("BFM", "Bitfield Move"))
	registry.Register("BFI", NewBitfieldExecutor("BFI", "Bitfield Insert"))
	registry.Register("BFXIL", NewBitfieldExecutor("BFXIL", "Bitfield Extract and Insert Low"))

	// Signed bitfield operations
	registry.Register("SBFM", NewBitfieldExecutor("SBFM", "Signed Bitfield Move"))
	registry.Register("SBFIZ", NewBitfieldExecutor("SBFIZ", "Signed Bitfield Insert in Zero"))
	registry.Register("SBFX", NewBitfieldExecutor("SBFX", "Signed Bitfield Extract"))
	registry.Register("ASR", NewBitfieldExecutor("ASR", "Arithmetic Shift Right"))
	registry.Register("SXTB", NewBitfieldExecutor("SXTB", "Sign Extend Byte"))
	registry.Register("SXTH", NewBitfieldExecutor("SXTH", "Sign Extend Halfword"))
	registry.Register("SXTW", NewBitfieldExecutor("SXTW", "Sign Extend Word"))

	// Unsigned bitfield operations
	registry.Register("UBFM", NewBitfieldExecutor("UBFM", "Unsigned Bitfield Move"))
	registry.Register("UBFIZ", NewBitfieldExecutor("UBFIZ", "Unsigned Bitfield Insert in Zero"))
	registry.Register("UBFX", NewBitfieldExecutor("UBFX", "Unsigned Bitfield Extract"))
	registry.Register("LSL", NewBitfieldExecutor("LSL", "Logical Shift Left"))
	registry.Register("LSR", NewBitfieldExecutor("LSR", "Logical Shift Right"))
	registry.Register("UXTB", NewBitfieldExecutor("UXTB", "Unsigned Extend Byte"))
	registry.Register("UXTH", NewBitfieldExecutor("UXTH", "Unsigned Extend Halfword"))
}
