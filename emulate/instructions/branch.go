package instructions

import (
	"fmt"
	"strings"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// BranchExecutor handles branch and jump instructions
type BranchExecutor struct {
	*BaseExecutor
}

// NewBranchExecutor creates a new branch instruction executor
func NewBranchExecutor(mnemonic, description string) *BranchExecutor {
	return &BranchExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// ValidateInstruction performs basic validation for branch instructions
func (e *BranchExecutor) ValidateInstruction(instr *disassemble.Instruction) error {
	if instr == nil {
		return core.NewEmulationError(core.ErrInvalidInstruction, 0, e.mnemonic, "nil instruction")
	}
	// For branch instructions, we don't validate against the operation string
	// since conditional branches may share operation constants
	return nil
}

// Execute executes branch instructions
func (e *BranchExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	switch e.mnemonic {
	case "B":
		return e.executeB(state, instr)
	case "BL":
		return e.executeBL(state, instr)
	case "BLR":
		return e.executeBLR(state, instr)
	case "BR":
		return e.executeBR(state, instr)
	case "RET":
		return e.executeRET(state, instr)
	case "CBZ":
		return e.executeCBZ(state, instr)
	case "CBNZ":
		return e.executeCBNZ(state, instr)
	case "TBZ":
		return e.executeTBZ(state, instr)
	case "TBNZ":
		return e.executeTBNZ(state, instr)
	// Conditional branches
	case "B.EQ", "B.NE", "B.CS", "B.CC", "B.MI", "B.PL", "B.VS", "B.VC",
		"B.HI", "B.LS", "B.GE", "B.LT", "B.GT", "B.LE", "B.AL", "B.NV":
		return e.executeConditionalBranch(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("branch instruction %s not implemented", e.mnemonic))
	}
}

// B - Unconditional branch
func (e *BranchExecutor) executeB(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "B requires target address")
	}

	// Calculate target address
	target, err := e.calculateBranchTarget(state, instr.Operands[0])
	if err != nil {
		return err
	}

	state.SetPC(target)
	return nil
}

// BL - Branch with link
func (e *BranchExecutor) executeBL(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BL requires target address")
	}

	// Save return address in link register (X30)
	returnAddr := state.GetPC() + 4
	state.SetX(30, returnAddr)

	// Calculate target address
	target, err := e.calculateBranchTarget(state, instr.Operands[0])
	if err != nil {
		return err
	}

	state.SetPC(target)
	return nil
}

// BLR - Branch with link to register
func (e *BranchExecutor) executeBLR(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BLR requires target register")
	}

	// Save return address in link register (X30)
	returnAddr := state.GetPC() + 4
	state.SetX(30, returnAddr)

	// Get target address from register
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BLR requires register operand")
	}

	targetReg := core.MapRegister(instr.Operands[0].Registers[0])
	if targetReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid target register")
	}

	target := state.GetX(targetReg)
	state.SetPC(target)
	return nil
}

// BR - Branch to register
func (e *BranchExecutor) executeBR(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BR requires target register")
	}

	// Get target address from register
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "BR requires register operand")
	}

	targetReg := core.MapRegister(instr.Operands[0].Registers[0])
	if targetReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid target register")
	}

	target := state.GetX(targetReg)
	state.SetPC(target)
	return nil
}

// RET - Return from subroutine
func (e *BranchExecutor) executeRET(state core.State, instr *disassemble.Instruction) error {
	var targetReg int
	if len(instr.Operands) > 0 && len(instr.Operands[0].Registers) > 0 {
		// RET with specific register
		targetReg = core.MapRegister(instr.Operands[0].Registers[0])
		if targetReg == -1 {
			return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), "invalid return register")
		}
	} else {
		// Default RET uses X30 (link register)
		targetReg = 30
	}

	target := state.GetX(targetReg)
	state.SetPC(target)
	return nil
}

// CBZ - Compare and branch if zero
func (e *BranchExecutor) executeCBZ(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CBZ requires register and target")
	}

	// Get register value
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CBZ requires register operand")
	}

	testReg := core.MapRegister(instr.Operands[0].Registers[0])
	if testReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid test register")
	}

	value := state.GetX(testReg)

	// Check if this is a 32-bit operation (W register)
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	if uint32(instr.Operands[0].Registers[0]) >= 1 && uint32(instr.Operands[0].Registers[0]) <= 31 {
		value = uint64(state.GetW(testReg))
	}

	// Branch if zero
	if value == 0 {
		target, err := e.calculateBranchTarget(state, instr.Operands[1])
		if err != nil {
			return err
		}
		state.SetPC(target)
	} else {
		// Continue to next instruction
		state.SetPC(state.GetPC() + 4)
	}

	return nil
}

// CBNZ - Compare and branch if not zero
func (e *BranchExecutor) executeCBNZ(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CBNZ requires register and target")
	}

	// Get register value
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "CBNZ requires register operand")
	}

	testReg := core.MapRegister(instr.Operands[0].Registers[0])
	if testReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid test register")
	}

	value := state.GetX(testReg)

	// Check if this is a 32-bit operation (W register)
	// W registers are 1-31, X registers are 34-64 in the disassemble package
	if uint32(instr.Operands[0].Registers[0]) >= 1 && uint32(instr.Operands[0].Registers[0]) <= 31 {
		value = uint64(state.GetW(testReg))
	}

	// Branch if not zero
	if value != 0 {
		target, err := e.calculateBranchTarget(state, instr.Operands[1])
		if err != nil {
			return err
		}
		state.SetPC(target)
	} else {
		// Continue to next instruction
		state.SetPC(state.GetPC() + 4)
	}

	return nil
}

// TBZ - Test bit and branch if zero
func (e *BranchExecutor) executeTBZ(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TBZ requires register, bit position, and target")
	}

	// Get register value
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TBZ requires register operand")
	}

	testReg := core.MapRegister(instr.Operands[0].Registers[0])
	if testReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid test register")
	}

	value := state.GetX(testReg)
	bitPos := uint64(instr.Operands[1].Immediate)

	// Validate bit position
	if bitPos >= 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid bit position")
	}

	// Test bit and branch if zero
	if (value & (1 << bitPos)) == 0 {
		target, err := e.calculateBranchTarget(state, instr.Operands[2])
		if err != nil {
			return err
		}
		state.SetPC(target)
	} else {
		// Continue to next instruction
		state.SetPC(state.GetPC() + 4)
	}

	return nil
}

// TBNZ - Test bit and branch if not zero
func (e *BranchExecutor) executeTBNZ(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TBNZ requires register, bit position, and target")
	}

	// Get register value
	if len(instr.Operands[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "TBNZ requires register operand")
	}

	testReg := core.MapRegister(instr.Operands[0].Registers[0])
	if testReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid test register")
	}

	value := state.GetX(testReg)
	bitPos := uint64(instr.Operands[1].Immediate)

	// Validate bit position
	if bitPos >= 64 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid bit position")
	}

	// Test bit and branch if not zero
	if (value & (1 << bitPos)) != 0 {
		target, err := e.calculateBranchTarget(state, instr.Operands[2])
		if err != nil {
			return err
		}
		state.SetPC(target)
	} else {
		// Continue to next instruction
		state.SetPC(state.GetPC() + 4)
	}

	return nil
}

// executeConditionalBranch handles conditional branch instructions
func (e *BranchExecutor) executeConditionalBranch(state core.State, instr *disassemble.Instruction) error {
	if len(instr.Operands) < 1 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "conditional branch requires target address")
	}

	// Extract condition code from mnemonic (e.g., "B.EQ" -> EQ)
	condition := e.extractConditionCode(e.mnemonic)

	// Evaluate condition
	if e.evaluateCondition(state, condition) {
		target, err := e.calculateBranchTarget(state, instr.Operands[0])
		if err != nil {
			return err
		}
		state.SetPC(target)
	} else {
		// Continue to next instruction
		state.SetPC(state.GetPC() + 4)
	}

	return nil
}

// Helper methods

// calculateBranchTarget calculates the target address for a branch
func (e *BranchExecutor) calculateBranchTarget(state core.State, operand disassemble.Operand) (uint64, error) {
	// Check if this is a register operand first
	if len(operand.Registers) > 0 {
		// Register indirect branch
		regIdx := core.MapRegister(operand.Registers[0])
		if regIdx == -1 {
			return 0, core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				"BRANCH", "invalid target register")
		}
		return state.GetX(regIdx), nil
	} else {
		// Immediate operand is already an absolute target provided by the disassembler
		return operand.Immediate, nil
	}
}

// extractConditionCode extracts condition code from instruction string
func (e *BranchExecutor) extractConditionCode(instrStr string) core.ConditionCode {
	parts := strings.Split(instrStr, ".")
	if len(parts) < 2 {
		return core.AL // Default to always
	}

	switch strings.ToUpper(parts[1]) {
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
		return core.AL
	}
}

// evaluateCondition evaluates a condition code against current flags
func (e *BranchExecutor) evaluateCondition(state core.State, cond core.ConditionCode) bool {
	n := state.GetN()
	z := state.GetZ()
	c := state.GetC()
	v := state.GetV()

	switch cond {
	case core.EQ:
		return z // Equal (Z set)
	case core.NE:
		return !z // Not equal (Z clear)
	case core.CS:
		return c // Carry set (C set)
	case core.CC:
		return !c // Carry clear (C clear)
	case core.MI:
		return n // Minus/negative (N set)
	case core.PL:
		return !n // Plus/positive or zero (N clear)
	case core.VS:
		return v // Overflow set (V set)
	case core.VC:
		return !v // Overflow clear (V clear)
	case core.HI:
		return c && !z // Higher (unsigned): C set and Z clear
	case core.LS:
		return !c || z // Lower or same (unsigned): C clear or Z set
	case core.GE:
		return n == v // Greater or equal (signed): N equals V
	case core.LT:
		return n != v // Less than (signed): N not equal to V
	case core.GT:
		return !z && (n == v) // Greater than (signed): Z clear and N equals V
	case core.LE:
		return z || (n != v) // Less or equal (signed): Z set or N not equal to V
	case core.AL:
		return true // Always
	case core.NV:
		return false // Never (deprecated, should not be used)
	default:
		return false
	}
}

// RegisterBranchInstructions registers all branch instructions
func RegisterBranchInstructions(registry *Registry) {
	// Unconditional branches
	registry.Register("B", NewBranchExecutor("B", "Branch"))
	registry.Register("BL", NewBranchExecutor("BL", "Branch with link"))
	registry.Register("BLR", NewBranchExecutor("BLR", "Branch with link to register"))
	registry.Register("BR", NewBranchExecutor("BR", "Branch to register"))
	registry.Register("RET", NewBranchExecutor("RET", "Return from subroutine"))

	// Compare and branch
	registry.Register("CBZ", NewBranchExecutor("CBZ", "Compare and branch if zero"))
	registry.Register("CBNZ", NewBranchExecutor("CBNZ", "Compare and branch if not zero"))

	// Test bit and branch
	registry.Register("TBZ", NewBranchExecutor("TBZ", "Test bit and branch if zero"))
	registry.Register("TBNZ", NewBranchExecutor("TBNZ", "Test bit and branch if not zero"))

	// Conditional branches
	registry.Register("B.EQ", NewBranchExecutor("B.EQ", "Branch if equal"))
	registry.Register("B.NE", NewBranchExecutor("B.NE", "Branch if not equal"))
	registry.Register("B.CS", NewBranchExecutor("B.CS", "Branch if carry set"))
	registry.Register("B.CC", NewBranchExecutor("B.CC", "Branch if carry clear"))
	registry.Register("B.MI", NewBranchExecutor("B.MI", "Branch if minus"))
	registry.Register("B.PL", NewBranchExecutor("B.PL", "Branch if plus"))
	registry.Register("B.VS", NewBranchExecutor("B.VS", "Branch if overflow set"))
	registry.Register("B.VC", NewBranchExecutor("B.VC", "Branch if overflow clear"))
	registry.Register("B.HI", NewBranchExecutor("B.HI", "Branch if higher"))
	registry.Register("B.LS", NewBranchExecutor("B.LS", "Branch if lower or same"))
	registry.Register("B.GE", NewBranchExecutor("B.GE", "Branch if greater or equal"))
	registry.Register("B.LT", NewBranchExecutor("B.LT", "Branch if less than"))
	registry.Register("B.GT", NewBranchExecutor("B.GT", "Branch if greater than"))
	registry.Register("B.LE", NewBranchExecutor("B.LE", "Branch if less or equal"))
	registry.Register("B.AL", NewBranchExecutor("B.AL", "Branch always"))
	registry.Register("B.NV", NewBranchExecutor("B.NV", "Branch never"))

	// Aliases
	registry.RegisterWithAliases("B.CS", NewBranchExecutor("B.CS", "Branch if carry set"), "B.HS")
	registry.RegisterWithAliases("B.CC", NewBranchExecutor("B.CC", "Branch if carry clear"), "B.LO")
}

// Helper functions for branch instruction analysis

// IsBranchInstruction checks if a mnemonic is a branch instruction
func IsBranchInstruction(mnemonic string) bool {
	switch mnemonic {
	case "B", "BL", "BLR", "BR", "RET", "CBZ", "CBNZ", "TBZ", "TBNZ":
		return true
	default:
		return strings.HasPrefix(mnemonic, "B.")
	}
}

// IsUnconditionalBranch checks if a branch is unconditional
func IsUnconditionalBranch(mnemonic string) bool {
	switch mnemonic {
	case "B", "BL", "BLR", "BR", "RET":
		return true
	case "B.AL":
		return true
	default:
		return false
	}
}

// IsConditionalBranch checks if a branch is conditional
func IsConditionalBranch(mnemonic string) bool {
	if !IsBranchInstruction(mnemonic) {
		return false
	}
	return !IsUnconditionalBranch(mnemonic) && mnemonic != "CBZ" && mnemonic != "CBNZ" &&
		mnemonic != "TBZ" && mnemonic != "TBNZ"
}

// IsLinkBranch checks if a branch saves the return address
func IsLinkBranch(mnemonic string) bool {
	switch mnemonic {
	case "BL", "BLR":
		return true
	default:
		return false
	}
}

// IsRegisterBranch checks if a branch uses a register for the target
func IsRegisterBranch(mnemonic string) bool {
	switch mnemonic {
	case "BLR", "BR", "RET":
		return true
	default:
		return false
	}
}

// Enum-based helpers (prefer these in new code to avoid string matching costs)

// IsReturnOp returns true if the instruction is any return-like instruction.
func IsReturnOp(instr *disassemble.Instruction) bool {
	if instr == nil {
		return false
	}
	switch instr.Operation {
	case disassemble.ARM64_RET, disassemble.ARM64_RETAA, disassemble.ARM64_RETAB,
		disassemble.ARM64_ERET, disassemble.ARM64_ERETAA, disassemble.ARM64_ERETAB:
		return true
	default:
		return false
	}
}

// IsReturnMnemonic returns true if the mnemonic corresponds to a return-like instruction.
func IsReturnMnemonic(mnemonic string) bool {
	switch strings.ToUpper(mnemonic) {
	case "RET", "RETAA", "RETAB", "ERET", "ERETAA", "ERETAB":
		return true
	default:
		return false
	}
}

// IsConditionalBranchOp returns true for conditional B.<cc> operations.
func IsConditionalBranchOp(instr *disassemble.Instruction) bool {
	if instr == nil {
		return false
	}
	switch instr.Operation {
	case disassemble.ARM64_B_AL, disassemble.ARM64_B_CC, disassemble.ARM64_B_CS,
		disassemble.ARM64_B_EQ, disassemble.ARM64_B_GE, disassemble.ARM64_B_GT,
		disassemble.ARM64_B_HI, disassemble.ARM64_B_LE, disassemble.ARM64_B_LS,
		disassemble.ARM64_B_LT, disassemble.ARM64_B_MI, disassemble.ARM64_B_NE,
		disassemble.ARM64_B_NV, disassemble.ARM64_B_PL, disassemble.ARM64_B_VC,
		disassemble.ARM64_B_VS:
		return true
	default:
		return false
	}
}

// IsUnconditionalBranchOp returns true for non-return unconditional branches.
func IsUnconditionalBranchOp(instr *disassemble.Instruction) bool {
	if instr == nil {
		return false
	}
	switch instr.Operation {
	case disassemble.ARM64_B, disassemble.ARM64_BL, disassemble.ARM64_BLR, disassemble.ARM64_BR:
		return true
	default:
		return false
	}
}

// IsBranchOp returns true for any branch-like operation including returns.
func IsBranchOp(instr *disassemble.Instruction) bool {
	if IsUnconditionalBranchOp(instr) || IsConditionalBranchOp(instr) {
		return true
	}
	if instr == nil {
		return false
	}
	switch instr.Operation {
	case disassemble.ARM64_CBZ, disassemble.ARM64_CBNZ, disassemble.ARM64_TBZ, disassemble.ARM64_TBNZ:
		return true
	}
	return IsReturnOp(instr)
}
