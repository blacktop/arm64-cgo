package instructions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// ConditionalExecutor handles conditional instructions
type ConditionalExecutor struct {
	*BaseExecutor
	// For testing: allows overriding condition extraction (unexported)
	testConditionOverride *core.ConditionCode
}

// NewConditionalExecutor creates a new conditional instruction executor
func NewConditionalExecutor(mnemonic, description string) *ConditionalExecutor {
	return &ConditionalExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// SetTestConditionOverride sets a condition override for testing
func (e *ConditionalExecutor) SetTestConditionOverride(condition core.ConditionCode) {
	e.testConditionOverride = &condition
}

// ClearTestConditionOverride clears the condition override
func (e *ConditionalExecutor) ClearTestConditionOverride() {
	e.testConditionOverride = nil
}

// Execute executes conditional instructions
func (e *ConditionalExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.GetMnemonic() {
	case "CSEL":
		return e.executeCSEL(state, instr)
    case "CSINC":
        return e.executeCSINC(state, instr)
	case "CSINV":
		return e.executeCSINV(state, instr)
	case "CSNEG":
		return e.executeCSNEG(state, instr)
	case "CSET":
		return e.executeCSET(state, instr)
	case "CSETM":
		return e.executeCSETM(state, instr)
    case "CINC":
        return e.executeCINC(state, instr)
	case "CINV":
		return e.executeCSINV(state, instr)
	case "CNEG":
		return e.executeCSNEG(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("conditional instruction %s not implemented", e.GetMnemonic()))
	}
}

// CSEL - Conditional select
func (e *ConditionalExecutor) executeCSEL(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSEL requires at least 3 operands")
	}

	// Check that each operand has at least one register
	if len(ops[0].Registers) == 0 || len(ops[1].Registers) == 0 || len(ops[2].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSEL operands missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	src1Reg := core.MapRegister(ops[1].Registers[0])
	src2Reg := core.MapRegister(ops[2].Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

	// Extract condition from instruction
	condition := e.ExtractCondition(instr)

	var result uint64
	if e.evaluateCondition(state, condition) {
		result = state.GetX(src1Reg)
	} else {
		result = state.GetX(src2Reg)
	}

	// Handle 32-bit vs 64-bit operations (W registers are 1-31, X registers are 34-64)
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CSINC - Conditional select increment
func (e *ConditionalExecutor) executeCSINC(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	// CINC is an alias for CSINC Rd, Rn, Rn, cond. The disassembler might provide fewer operands.
	if (e.GetMnemonic() == "CINC" && len(ops) < 2) || (e.GetMnemonic() != "CINC" && len(ops) < 3) {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINC/CINC requires more operands")
	}

	// Check that required operands have registers
	if len(ops[0].Registers) == 0 || len(ops[1].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINC/CINC operands missing register information")
	}
	if e.GetMnemonic() != "CINC" && len(ops[2].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINC non-alias operand missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	src1Reg := core.MapRegister(ops[1].Registers[0])
	var src2Reg int
	if e.GetMnemonic() == "CINC" {
		src2Reg = src1Reg // For CINC, Rm is the same as Rn
	} else {
		src2Reg = core.MapRegister(ops[2].Registers[0])
	}

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

    // CSINC semantics: if condition true -> Rd = Rn, else -> Rd = Rm + 1
    condition := e.ExtractCondition(instr)
    var result uint64
    if e.evaluateCondition(state, condition) {
        result = state.GetX(src1Reg)
    } else {
        result = state.GetX(src2Reg) + 1
    }

	// Handle 32-bit vs 64-bit operations (W registers are 1-31, X registers are 34-64)
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CSINV - Conditional select invert
func (e *ConditionalExecutor) executeCSINV(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if (e.GetMnemonic() == "CINV" && len(ops) < 2) || (e.GetMnemonic() != "CINV" && len(ops) < 3) {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINV/CINV requires more operands")
	}

	// Check that required operands have registers
	if len(ops[0].Registers) == 0 || len(ops[1].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINV/CINV operands missing register information")
	}
	if e.GetMnemonic() != "CINV" && len(ops[2].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSINV non-alias operand missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	src1Reg := core.MapRegister(ops[1].Registers[0])
	var src2Reg int
	if e.GetMnemonic() == "CINV" {
		src2Reg = src1Reg // For CINV, Rm is the same as Rn
	} else {
		src2Reg = core.MapRegister(ops[2].Registers[0])
	}

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

    condition := e.ExtractCondition(instr)
    // CINV alias requires inverted condition relative to CSINV semantics
    if e.GetMnemonic() == "CINV" {
        condition = e.invertCondition(condition)
    }

	// Determine width for inversion mask
	is32 := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31
	var allOnes uint64 = 0xFFFFFFFFFFFFFFFF
	if is32 {
		allOnes = 0xFFFFFFFF
	}

    var result uint64
    if e.evaluateCondition(state, condition) {
        // Condition true: select Rn
        result = state.GetX(src1Reg)
    } else {
        // Condition false: select bitwise NOT of Rm (masked to width)
        rm := state.GetX(src2Reg)
        result = (^rm) & allOnes
    }

	if is32 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CSNEG - Conditional select negate
func (e *ConditionalExecutor) executeCSNEG(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSNEG requires at least 3 operands")
	}

	// Check that each operand has at least one register
	if len(ops[0].Registers) == 0 || len(ops[1].Registers) == 0 || len(ops[2].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSNEG operands missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	src1Reg := core.MapRegister(ops[1].Registers[0])
	src2Reg := core.MapRegister(ops[2].Registers[0])

	if dstReg == -1 || src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid register")
	}

    condition := e.ExtractCondition(instr)
    // CNEG alias requires inverted condition relative to CSNEG semantics
    if e.GetMnemonic() == "CNEG" {
        condition = e.invertCondition(condition)
    }

	var result uint64
    if e.evaluateCondition(state, condition) {
        // Condition true: select Rn
        result = state.GetX(src1Reg)
    } else {
        // Condition false: select -Rm
        result = uint64(-int64(state.GetX(src2Reg)))
    }

	// Handle 32-bit vs 64-bit operations (W registers are 1-31, X registers are 34-64)
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CSET - Conditional set (alias for CSINC with XZR)
func (e *ConditionalExecutor) executeCSET(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSET requires destination register")
	}

	// Check that the operand has at least one register
	if len(ops[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSET operand missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	condition := e.ExtractCondition(instr)

	var result uint64
	if e.evaluateCondition(state, condition) {
		result = 1
	} else {
		result = 0
	}

	// Handle 32-bit vs 64-bit operations (W registers are 1-31, X registers are 34-64)
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CSETM - Conditional set mask (alias for CSINV with XZR)
func (e *ConditionalExecutor) executeCSETM(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSETM requires destination register")
	}

	// Check that the operand has at least one register
	if len(ops[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CSETM operand missing register information")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	condition := e.ExtractCondition(instr)

	var result uint64
	if e.evaluateCondition(state, condition) {
		result = 0xFFFFFFFFFFFFFFFF // All bits set
	} else {
		result = 0
	}

	// Handle 32-bit vs 64-bit operations (W registers are 1-31, X registers are 34-64)
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		state.SetW(dstReg, uint32(result))
	} else {
		state.SetX(dstReg, result)
	}

	return nil
}

// CINC - Conditional increment (alias for CSINC)
func (e *ConditionalExecutor) executeCINC(state core.State, instr *disassemble.Instruction) error {
    ops := instr.Operands
    if len(ops) < 2 {
        return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
            fmt.Sprintf("%v", instr.Operation), "CINC requires at least 2 operands")
    }

    // dst and src (Rn)
    if len(ops[0].Registers) == 0 || len(ops[1].Registers) == 0 {
        return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
            fmt.Sprintf("%v", instr.Operation), "CINC operands missing register information")
    }

    dstReg := core.MapRegister(ops[0].Registers[0])
    srcReg := core.MapRegister(ops[1].Registers[0])
    if dstReg == -1 || srcReg == -1 {
        return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
            fmt.Sprintf("%v", instr.Operation), "invalid register")
    }

    // CINC semantics: increment if condition is true
    cond := e.ExtractCondition(instr)
    var result uint64
    if e.evaluateCondition(state, cond) {
        result = state.GetX(srcReg) + 1
    } else {
        result = state.GetX(srcReg)
    }

    if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
        state.SetW(dstReg, uint32(result))
    } else {
        state.SetX(dstReg, result)
    }
    return nil
}

// CINV - Conditional invert (alias for CSINV)
func (e *ConditionalExecutor) executeCINV(state core.State, instr *disassemble.Instruction) error {
	// CINV is an alias for CSINV where src1 == src2
	return e.executeCSINV(state, instr)
}

// CNEG - Conditional negate (alias for CSNEG)
func (e *ConditionalExecutor) executeCNEG(state core.State, instr *disassemble.Instruction) error {
	// CNEG is an alias for CSNEG where src1 == src2
	return e.executeCSNEG(state, instr)
}

// CCMP - Conditional compare
func (e *ConditionalExecutor) executeCCMP(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CCMP requires at least 3 operands")
	}

	// Check that the first operand has at least one register
	if len(ops[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CCMP first operand missing register information")
	}

	src1Reg := core.MapRegister(ops[0].Registers[0])
	if src1Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	// Get first operand value
	val1 := state.GetX(src1Reg)

	// Get second operand value (register or immediate)
	var val2 uint64
	if len(ops[1].Registers) > 0 {
		src2Reg := core.MapRegister(ops[1].Registers[0])
		if src2Reg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid second source register")
		}
		val2 = state.GetX(src2Reg)
	} else {
		val2 = uint64(ops[1].Immediate)
	}

	// Third operand is immediate flags value used when condition is false
	flagsImm := uint64(ops[2].Immediate)

	// Extract condition from instruction
	condition := e.ExtractCondition(instr)

	if e.evaluateCondition(state, condition) {
		// Condition is true - perform the compare operation (val1 - val2)
		result := val1 - val2

		// Set condition flags based on subtraction result
		state.SetZ(result == 0)                                    // Zero flag
		state.SetN((result & (1 << 63)) != 0)                      // Negative flag
		state.SetC(val1 >= val2)                                   // Carry flag (no borrow)
		state.SetV(e.checkSubtractionOverflow(val1, val2, result)) // Overflow flag
	} else {
		// Condition is false - use immediate flags value
		state.SetN((flagsImm & 8) != 0) // Bit 3
		state.SetZ((flagsImm & 4) != 0) // Bit 2
		state.SetC((flagsImm & 2) != 0) // Bit 1
		state.SetV((flagsImm & 1) != 0) // Bit 0
	}

	return nil
}

// Helper methods

// invertCondition inverts a condition code
func (e *ConditionalExecutor) invertCondition(cond core.ConditionCode) core.ConditionCode {
	switch cond {
	case core.EQ:
		return core.NE
	case core.NE:
		return core.EQ
	case core.CS:
		return core.CC
	case core.CC:
		return core.CS
	case core.MI:
		return core.PL
	case core.PL:
		return core.MI
	case core.VS:
		return core.VC
	case core.VC:
		return core.VS
	case core.HI:
		return core.LS
	case core.LS:
		return core.HI
	case core.GE:
		return core.LT
	case core.LT:
		return core.GE
	case core.GT:
		return core.LE
	case core.LE:
		return core.GT
	case core.AL:
		return core.NV
	case core.NV:
		return core.AL
	default:
		return cond
	}
}

// ExtractCondition extracts condition code from instruction
func (e *ConditionalExecutor) ExtractCondition(instr *disassemble.Instruction) core.ConditionCode {
	// For testing: check if we have a condition override
	if e.testConditionOverride != nil {
		return *e.testConditionOverride
	}

	// First try to find a structured condition operand
	for _, operand := range instr.Operands {
		if operand.Class == disassemble.CONDITION {
			return e.mapConditionCode(operand.Condition)
		}
	}

	// Try to decode condition from raw encoding for conditional select class
	// CSEL/CSINC/CSINV/CSNEG and aliases encode cond in bits [15:12]
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

	// Fall back to string parsing for compatibility
	instrStr := fmt.Sprintf("%v", instr.Operation)
	return e.parseConditionFromString(instrStr)
}

// mapConditionCode maps disassembler condition codes to our condition codes
func (e *ConditionalExecutor) mapConditionCode(cond any) core.ConditionCode {
	// Handle string-based condition codes first
	if condStr, ok := cond.(string); ok {
		switch strings.ToUpper(condStr) {
		case "EQ":
			return core.EQ
		case "NE":
			return core.NE
		case "CS", "HS":
			return core.CS
		case "CC", "LO":
			return core.CC
		case "MI":
			return core.MI
		case "PL":
			return core.PL
		case "VS":
			return core.VS
		case "VC":
			return core.VC
		case "HI":
			return core.HI
		case "LS":
			return core.LS
		case "GE":
			return core.GE
		case "LT":
			return core.LT
		case "GT":
			return core.GT
		case "LE":
			return core.LE
		case "AL":
			return core.AL
		case "NV":
			return core.NV
		default:
			return core.AL // Default to always for unknown strings
		}
	}

	// Convert numeric condition to integer for evaluation
	var conditionCode int
	switch v := cond.(type) {
	case int:
		conditionCode = v
	case uint32:
		conditionCode = int(v)
	case uint64:
		conditionCode = int(v)
	default:
		// Try to convert to int using string conversion as last resort
		// This handles disassemble.condition type (uint32 based)
		condStr := fmt.Sprintf("%v", cond)
		if condInt, err := strconv.Atoi(condStr); err == nil {
			conditionCode = condInt
		} else {
			return core.AL // Default to always
		}
	}

	switch conditionCode {
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
	default:
		return core.AL
	}
}

// parseConditionFromString provides fallback string-based condition parsing
func (e *ConditionalExecutor) parseConditionFromString(instrStr string) core.ConditionCode {
	instrLower := strings.ToLower(instrStr)

	// Check for both dotted (.eq) and non-dotted (eq) condition suffixes
	if strings.Contains(instrLower, ".eq") || strings.HasSuffix(instrLower, " eq") {
		return core.EQ
	} else if strings.Contains(instrLower, ".ne") || strings.HasSuffix(instrLower, " ne") {
		return core.NE
	} else if strings.Contains(instrLower, ".cs") || strings.Contains(instrLower, ".hs") || strings.HasSuffix(instrLower, " cs") || strings.HasSuffix(instrLower, " hs") {
		return core.CS
	} else if strings.Contains(instrLower, ".cc") || strings.Contains(instrLower, ".lo") || strings.HasSuffix(instrLower, " cc") || strings.HasSuffix(instrLower, " lo") {
		return core.CC
	} else if strings.Contains(instrLower, ".mi") || strings.HasSuffix(instrLower, " mi") {
		return core.MI
	} else if strings.Contains(instrLower, ".pl") || strings.HasSuffix(instrLower, " pl") {
		return core.PL
	} else if strings.Contains(instrLower, ".vs") || strings.HasSuffix(instrLower, " vs") {
		return core.VS
	} else if strings.Contains(instrLower, ".vc") || strings.HasSuffix(instrLower, " vc") {
		return core.VC
	} else if strings.Contains(instrLower, ".hi") || strings.HasSuffix(instrLower, " hi") {
		return core.HI
	} else if strings.Contains(instrLower, ".ls") || strings.HasSuffix(instrLower, " ls") {
		return core.LS
	} else if strings.Contains(instrLower, ".ge") || strings.HasSuffix(instrLower, " ge") {
		return core.GE
	} else if strings.Contains(instrLower, ".lt") || strings.HasSuffix(instrLower, " lt") {
		return core.LT
	} else if strings.Contains(instrLower, ".gt") || strings.HasSuffix(instrLower, " gt") {
		return core.GT
	} else if strings.Contains(instrLower, ".le") || strings.HasSuffix(instrLower, " le") {
		return core.LE
	}

	return core.AL // Default to always
}

// evaluateCondition evaluates a condition code against current flags
func (e *ConditionalExecutor) evaluateCondition(state core.State, cond core.ConditionCode) bool {
	n := state.GetN()
	z := state.GetZ()
	c := state.GetC()
	v := state.GetV()

	var result bool
	switch cond {
	case core.EQ:
		result = z
	case core.NE:
		result = !z
	case core.CS:
		result = c
	case core.CC:
		result = !c
	case core.MI:
		result = n
	case core.PL:
		result = !n
	case core.VS:
		result = v
	case core.VC:
		result = !v
	case core.HI:
		result = c && !z
	case core.LS:
		result = !c || z
	case core.GE:
		result = n == v
	case core.LT:
		result = n != v
	case core.GT:
		result = !z && (n == v)
	case core.LE:
		result = z || (n != v)
	case core.AL:
		result = true
	case core.NV:
		result = false
	default:
		result = false
	}

	return result
}

// checkSubtractionOverflow checks for signed overflow in subtraction (a - b)
func (e *ConditionalExecutor) checkSubtractionOverflow(a, b, result uint64) bool {
	// Overflow occurs when:
	// - Subtracting a positive from negative gives positive (negative - positive = positive)
	// - Subtracting a negative from positive gives negative (positive - negative = negative)
	aSign := (a & (1 << 63)) != 0
	bSign := (b & (1 << 63)) != 0
	rSign := (result & (1 << 63)) != 0

	return (aSign != bSign) && (aSign != rSign)
}

// RegisterConditionalInstructions registers all conditional instructions
func RegisterConditionalInstructions(registry *Registry) {
	registry.Register("CSEL", NewConditionalExecutor("CSEL", "Conditional select"))
	registry.Register("CSINC", NewConditionalExecutor("CSINC", "Conditional select increment"))
	registry.Register("CSINV", NewConditionalExecutor("CSINV", "Conditional select invert"))
	registry.Register("CSNEG", NewConditionalExecutor("CSNEG", "Conditional select negate"))
	registry.Register("CSET", NewConditionalExecutor("CSET", "Conditional set"))
	registry.Register("CSETM", NewConditionalExecutor("CSETM", "Conditional set mask"))
	registry.Register("CINC", NewConditionalExecutor("CINC", "Conditional increment"))
	registry.Register("CINV", NewConditionalExecutor("CINV", "Conditional invert"))
	registry.Register("CNEG", NewConditionalExecutor("CNEG", "Conditional negate"))
	// CCMP is registered in compare instructions as it's fundamentally a compare operation
}
