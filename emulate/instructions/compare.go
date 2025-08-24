package instructions

import (
	"fmt"
	"strings"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// CompareExecutor handles comparison instructions
type CompareExecutor struct {
	*BaseExecutor
}

// NewCompareExecutor creates a new compare instruction executor
func NewCompareExecutor(mnemonic, description string) *CompareExecutor {
	return &CompareExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// ValidateInstruction performs basic validation for compare instructions
func (e *CompareExecutor) ValidateInstruction(instr *disassemble.Instruction) error {
	if instr == nil {
		return core.NewEmulationError(core.ErrInvalidInstruction, 0, e.mnemonic, "nil instruction")
	}
	// For compare instructions, we don't validate against the operation string
	// since the operation format may vary
	return nil
}

// Execute executes compare instructions
func (e *CompareExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.mnemonic {
	case "CMP":
		return e.executeCMP(state, instr)
	case "CMN":
		return e.executeCMN(state, instr)
	case "TST":
		return e.executeTST(state, instr)
	case "CCMP":
		return e.executeCCMP(state, instr)
	case "CCMN":
		return e.executeCCMN(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("compare instruction %s not implemented", e.mnemonic))
	}
}

// CCMN - Conditional compare negative
func (e *CompareExecutor) executeCCMN(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CCMN requires 4 operands")
	}

	// Extract condition from encoding or string
	condition := e.extractConditionFromInstruction(fmt.Sprintf("%v", instr.Operation), instr)

	if e.evaluateCondition(state, condition) {
		// Perform CMN (add and set flags)
		src1Reg := core.MapRegister(ops[0].Registers[0])
		if src1Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid first register")
		}
		val1 := state.GetX(src1Reg)
		var val2 uint64
		if len(ops[1].Registers) > 0 {
			src2Reg := core.MapRegister(ops[1].Registers[0])
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid second register")
			}
			val2 = state.GetX(src2Reg)
		} else {
			val2 = uint64(ops[1].Immediate)
		}
		is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
		if is32Bit {
			val1 = uint64(uint32(val1))
			val2 = uint64(uint32(val2))
		}
		result, carry, overflow := e.addWithFlags(val1, val2, is32Bit)
		if is32Bit {
			state.SetN((result & 0x80000000) != 0)
		} else {
			state.SetN((result & 0x8000000000000000) != 0)
		}
		state.SetZ(result == 0)
		state.SetC(carry)
		state.SetV(overflow)
	} else {
		// Use immediate NZCV
		nzcv := uint64(ops[2].Immediate)
		state.SetN((nzcv & 8) != 0)
		state.SetZ((nzcv & 4) != 0)
		state.SetC((nzcv & 2) != 0)
		state.SetV((nzcv & 1) != 0)
	}
	return nil
}

// CMP - Compare (subtract and set flags but discard result)
func (e *CompareExecutor) executeCMP(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CMP requires at least 2 operands")
	}

	src1Reg := core.MapRegister(ops[0].Registers[0])
	if src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid first register")
	}

	val1 := state.GetX(src1Reg)
	var val2 uint64

	// Second operand can be register or immediate
	if len(ops[1].Registers) > 0 {
		src2Reg := core.MapRegister(ops[1].Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second register")
		}
		val2 = state.GetX(src2Reg)
	} else {
		// Immediate operand (including zero)
		val2 = uint64(ops[1].Immediate)
	}

	// Handle 32-bit operations (W registers)
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
	if is32Bit {
		val1 = uint64(uint32(val1))
		val2 = uint64(uint32(val2))
	}

	// Perform subtraction for comparison (val1 - val2)
	result, carry, overflow := e.subWithFlags(val1, val2, is32Bit)

	// Update flags
	if is32Bit {
		state.SetN((result & 0x80000000) != 0)
	} else {
		state.SetN((result & 0x8000000000000000) != 0)
	}
	state.SetZ(result == 0)
	state.SetC(carry) // For subtraction, C=1 means no borrow
	state.SetV(overflow)

	return nil
}

// CMN - Compare negative (add and set flags but discard result)
func (e *CompareExecutor) executeCMN(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CMN requires at least 2 operands")
	}

	src1Reg := core.MapRegister(ops[0].Registers[0])
	if src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid first register")
	}

	val1 := state.GetX(src1Reg)
	var val2 uint64

	// Second operand can be register or immediate
	if len(ops[1].Registers) > 0 {
		src2Reg := core.MapRegister(ops[1].Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second register")
		}
		val2 = state.GetX(src2Reg)
	} else {
		// Immediate operand (including zero)
		val2 = uint64(ops[1].Immediate)
	}

	// Handle 32-bit operations (W registers)
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
	if is32Bit {
		val1 = uint64(uint32(val1))
		val2 = uint64(uint32(val2))
	}

	// Perform addition for comparison (val1 + val2)
	result, carry, overflow := e.addWithFlags(val1, val2, is32Bit)

	// Update flags
	if is32Bit {
		state.SetN((result & 0x80000000) != 0)
	} else {
		state.SetN((result & 0x8000000000000000) != 0)
	}
	state.SetZ(result == 0)
	state.SetC(carry)
	state.SetV(overflow)

	return nil
}

// TST - Test bits (AND and set flags but discard result)
func (e *CompareExecutor) executeTST(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TST requires at least 2 operands")
	}

	src1Reg := core.MapRegister(ops[0].Registers[0])
	if src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid first register")
	}

	val1 := state.GetX(src1Reg)
	var val2 uint64

	// Second operand can be register or immediate
	if len(ops[1].Registers) > 0 {
		src2Reg := core.MapRegister(ops[1].Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second register")
		}
		val2 = state.GetX(src2Reg)
	} else {
		// Immediate operand (including zero)
		val2 = uint64(ops[1].Immediate)
	}

	// Handle 32-bit operations (W registers)
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
	if is32Bit {
		val1 = uint64(uint32(val1))
		val2 = uint64(uint32(val2))
	}

	// Perform AND operation (result not stored, only used for flags)
	result := val1 & val2

	// Set condition flags based on result
	state.SetZ(result == 0) // Zero flag
	if is32Bit {
		state.SetN((result & 0x80000000) != 0) // Negative flag (check bit 31 for 32-bit)
	} else {
		state.SetN((result & 0x8000000000000000) != 0) // Negative flag (check bit 63 for 64-bit)
	}
	// C and V flags are not affected by TST

	return nil
}

// CCMP - Conditional compare
func (e *CompareExecutor) executeCCMP(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 4 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CCMP requires 4 operands")
	}

	// Extract condition from encoding or string
	condition := e.extractConditionFromInstruction(fmt.Sprintf("%v", instr.Operation), instr)

	// Check if condition is met
	if e.evaluateCondition(state, condition) {
		// Condition met: perform normal CMN
		src1Reg := core.MapRegister(ops[0].Registers[0])
		if src1Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid first register")
		}

		val1 := state.GetX(src1Reg)
		var val2 uint64

		// Second operand can be register or immediate
		if len(ops[1].Registers) > 0 {
			src2Reg := core.MapRegister(ops[1].Registers[0])
			if src2Reg == -1 {
				return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), "invalid second register")
			}
			val2 = state.GetX(src2Reg)
		} else {
			// Immediate operand (including zero)
			val2 = uint64(ops[1].Immediate)
		}

		// Handle 32-bit operations
		// W registers are 1-31, X registers are 34-64 in the disassemble package
		is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
		if is32Bit {
			val1 = uint64(uint32(val1))
			val2 = uint64(uint32(val2))
		}

		// Perform comparison (addition)
		result, carry, overflow := e.addWithFlags(val1, val2, is32Bit)

		// Update flags
		if is32Bit {
			state.SetN((result & 0x80000000) != 0)
		} else {
			state.SetN((result & 0x8000000000000000) != 0)
		}
		state.SetZ(result == 0)
		state.SetC(carry)
		state.SetV(overflow)
	} else {
		// Condition not met: set flags to immediate value
		nzcv := uint64(ops[2].Immediate) // NZCV immediate
		state.SetN((nzcv & 8) != 0)      // Bit 3
		state.SetZ((nzcv & 4) != 0)      // Bit 2
		state.SetC((nzcv & 2) != 0)      // Bit 1
		state.SetV((nzcv & 1) != 0)      // Bit 0
	}

	return nil
}

// Helper methods

// addWithFlags performs addition and returns carry and overflow flags
func (e *CompareExecutor) addWithFlags(a, b uint64, is32Bit bool) (result uint64, carry bool, overflow bool) {
	if is32Bit {
		// 32-bit addition
		a32 := uint32(a)
		b32 := uint32(b)
		result32 := a32 + b32
		result = uint64(result32)
		carry = result32 < a32 // Unsigned overflow

		// Signed overflow: result and operands have different signs
		signA := a32 & 0x80000000
		signB := b32 & 0x80000000
		signResult := result32 & 0x80000000
		overflow = (signA == signB) && (signA != signResult)
	} else {
		// 64-bit addition
		result = a + b
		carry = result < a // Unsigned overflow

		// Signed overflow: result and operands have different signs
		signA := a & 0x8000000000000000
		signB := b & 0x8000000000000000
		signResult := result & 0x8000000000000000
		overflow = (signA == signB) && (signA != signResult)
	}

	return
}

// subWithFlags performs subtraction and returns carry and overflow flags
func (e *CompareExecutor) subWithFlags(a, b uint64, is32Bit bool) (result uint64, carry bool, overflow bool) {
	if is32Bit {
		// 32-bit subtraction
		a32 := uint32(a)
		b32 := uint32(b)
		result32 := a32 - b32
		result = uint64(result32)
		carry = a32 >= b32 // No borrow (C=1 means no borrow in ARM)

		// Signed overflow for subtraction
		signA := a32 & 0x80000000
		signB := b32 & 0x80000000
		signResult := result32 & 0x80000000
		overflow = (signA != signB) && (signA != signResult)
	} else {
		// 64-bit subtraction
		result = a - b
		carry = a >= b // No borrow (C=1 means no borrow in ARM)

		// Signed overflow for subtraction
		signA := a & 0x8000000000000000
		signB := b & 0x8000000000000000
		signResult := result & 0x8000000000000000
		overflow = (signA != signB) && (signA != signResult)
	}

	return
}

// extractConditionFromInstruction extracts condition from instruction encoding or string
func (e *CompareExecutor) extractConditionFromInstruction(instrStr string, instr *disassemble.Instruction) core.ConditionCode {
	// Prefer structured CONDITION operand if present
	if instr != nil {
		for _, op := range instr.Operands {
			if op.Class == disassemble.CONDITION {
				// Map disassemble.condition to core.ConditionCode
				switch op.Condition {
				case disassemble.COND_EQ:
					return core.EQ
				case disassemble.COND_NE:
					return core.NE
				case disassemble.COND_CS:
					return core.CS
				case disassemble.COND_CC:
					return core.CC
				case disassemble.COND_MI:
					return core.MI
				case disassemble.COND_PL:
					return core.PL
				case disassemble.COND_VS:
					return core.VS
				case disassemble.COND_VC:
					return core.VC
				case disassemble.COND_HI:
					return core.HI
				case disassemble.COND_LS:
					return core.LS
				case disassemble.COND_GE:
					return core.GE
				case disassemble.COND_LT:
					return core.LT
				case disassemble.COND_GT:
					return core.GT
				case disassemble.COND_LE:
					return core.LE
				case disassemble.COND_AL:
					return core.AL
				case disassemble.COND_NV:
					return core.NV
				}
			}
		}
		// Fallback: decode cond from Raw bits [15:12] per ARM encoding for CCMP/CCMN
		if instr.Raw != 0 {
			condField := (instr.Raw >> 12) & 0xF
			switch condField {
			case 0:
				return core.EQ
			case 1:
				return core.NE
			case 2:
				return core.CS
			case 3:
				return core.CC
			case 4:
				return core.MI
			case 5:
				return core.PL
			case 6:
				return core.VS
			case 7:
				return core.VC
			case 8:
				return core.HI
			case 9:
				return core.LS
			case 10:
				return core.GE
			case 11:
				return core.LT
			case 12:
				return core.GT
			case 13:
				return core.LE
			case 14:
				return core.AL
			case 15:
				return core.NV
			}
		}
	}
	// String-based suffix fallback
	// Minimal string-based condition parse as fallback
	lower := strings.ToLower(instrStr)
	if strings.Contains(lower, ".eq") || strings.HasSuffix(lower, " eq") {
		return core.EQ
	} else if strings.Contains(lower, ".ne") || strings.HasSuffix(lower, " ne") {
		return core.NE
	} else if strings.Contains(lower, ".cs") || strings.Contains(lower, ".hs") || strings.HasSuffix(lower, " cs") || strings.HasSuffix(lower, " hs") {
		return core.CS
	} else if strings.Contains(lower, ".cc") || strings.Contains(lower, ".lo") || strings.HasSuffix(lower, " cc") || strings.HasSuffix(lower, " lo") {
		return core.CC
	} else if strings.Contains(lower, ".mi") || strings.HasSuffix(lower, " mi") {
		return core.MI
	} else if strings.Contains(lower, ".pl") || strings.HasSuffix(lower, " pl") {
		return core.PL
	} else if strings.Contains(lower, ".vs") || strings.HasSuffix(lower, " vs") {
		return core.VS
	} else if strings.Contains(lower, ".vc") || strings.HasSuffix(lower, " vc") {
		return core.VC
	} else if strings.Contains(lower, ".hi") || strings.HasSuffix(lower, " hi") {
		return core.HI
	} else if strings.Contains(lower, ".ls") || strings.HasSuffix(lower, " ls") {
		return core.LS
	} else if strings.Contains(lower, ".ge") || strings.HasSuffix(lower, " ge") {
		return core.GE
	} else if strings.Contains(lower, ".lt") || strings.HasSuffix(lower, " lt") {
		return core.LT
	} else if strings.Contains(lower, ".gt") || strings.HasSuffix(lower, " gt") {
		return core.GT
	} else if strings.Contains(lower, ".le") || strings.HasSuffix(lower, " le") {
		return core.LE
	}
	return core.AL
}

// evaluateCondition evaluates a condition code against current flags
func (e *CompareExecutor) evaluateCondition(state core.State, cond core.ConditionCode) bool {
	n := state.GetN()
	z := state.GetZ()
	c := state.GetC()
	v := state.GetV()

	switch cond {
	case core.EQ:
		return z
	case core.NE:
		return !z
	case core.CS:
		return c
	case core.CC:
		return !c
	case core.MI:
		return n
	case core.PL:
		return !n
	case core.VS:
		return v
	case core.VC:
		return !v
	case core.HI:
		return c && !z
	case core.LS:
		return !c || z
	case core.GE:
		return n == v
	case core.LT:
		return n != v
	case core.GT:
		return !z && (n == v)
	case core.LE:
		return z || (n != v)
	case core.AL:
		return true
	case core.NV:
		return false
	default:
		return false
	}
}

// RegisterCompareInstructions registers all compare instructions
func RegisterCompareInstructions(registry *Registry) {
	registry.Register("CMP", NewCompareExecutor("CMP", "Compare"))
	registry.Register("CMN", NewCompareExecutor("CMN", "Compare negative"))
	// TST is registered in logical instructions as it's fundamentally AND with flags
	registry.Register("CCMP", NewCompareExecutor("CCMP", "Conditional compare"))
	registry.Register("CCMN", NewCompareExecutor("CCMN", "Conditional compare negative"))
}
