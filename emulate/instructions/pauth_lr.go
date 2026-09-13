package instructions

import (
	"fmt"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// PAuthLRExecutor executes FEAT_PAuth_LR (PC-relative pointer authentication) instructions
// using the shared pseudo-PAC scheme in pac.go. The PC-relative forms fold the address of the
// signing instruction into the modifier, so an AUT/RET that names the matching label (or a
// register holding that address) strips the pseudo-PAC and any mismatch poisons the pointer.
type PAuthLRExecutor struct {
	*BaseExecutor
}

// NewPAuthLRExecutor creates a new FEAT_PAuth_LR executor for one mnemonic.
func NewPAuthLRExecutor(mnemonic, description string) *PAuthLRExecutor {
	return &PAuthLRExecutor{BaseExecutor: NewBaseExecutor(mnemonic, description)}
}

const linkRegister = 30

// Execute dispatches the FEAT_PAuth_LR instruction.
func (e *PAuthLRExecutor) Execute(state core.State, inst *disassemble.Inst) error {
	if err := e.ValidateInstruction(inst); err != nil {
		return err
	}

	switch inst.Operation {
	case disassemble.ARM64_PACM:
		return nil
	case disassemble.ARM64_PACIASPPC, disassemble.ARM64_PACNBIASPPC:
		signRegister(state, linkRegister, combineModifiers(state.GetSP(), state.GetPC()), pseudoPACKeyA)
		return nil
	case disassemble.ARM64_PACIBSPPC, disassemble.ARM64_PACNBIBSPPC:
		signRegister(state, linkRegister, combineModifiers(state.GetSP(), state.GetPC()), pseudoPACKeyB)
		return nil
	case disassemble.ARM64_PACIA171615:
		signRegister(state, 17, combineModifiers(state.GetX(16), state.GetX(15)), pseudoPACKeyA)
		return nil
	case disassemble.ARM64_PACIB171615:
		signRegister(state, 17, combineModifiers(state.GetX(16), state.GetX(15)), pseudoPACKeyB)
		return nil
	case disassemble.ARM64_AUTIA171615:
		authenticateRegister(state, 17, combineModifiers(state.GetX(16), state.GetX(15)), pseudoPACKeyA)
		return nil
	case disassemble.ARM64_AUTIB171615:
		authenticateRegister(state, 17, combineModifiers(state.GetX(16), state.GetX(15)), pseudoPACKeyB)
		return nil
	case disassemble.ARM64_AUTIASPPC, disassemble.ARM64_AUTIASPPCR:
		return e.authenticateLR(state, inst, pseudoPACKeyA, false)
	case disassemble.ARM64_AUTIBSPPC, disassemble.ARM64_AUTIBSPPCR:
		return e.authenticateLR(state, inst, pseudoPACKeyB, false)
	case disassemble.ARM64_RETAASPPC, disassemble.ARM64_RETAASPPCR:
		return e.authenticateLR(state, inst, pseudoPACKeyA, true)
	case disassemble.ARM64_RETABSPPC, disassemble.ARM64_RETABSPPCR:
		return e.authenticateLR(state, inst, pseudoPACKeyB, true)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			inst.Operation.String(), fmt.Sprintf("PAuth_LR instruction %s not implemented", e.mnemonic))
	}
}

// authenticateLR authenticates X30 against SP combined with the PC modifier named by the
// instruction's only operand (a label or a register). The AUT forms write the result back to
// X30. The RET forms branch to the authenticated value and leave X30 unchanged, as the Arm
// pseudocode authenticates into a temporary target.
func (e *PAuthLRExecutor) authenticateLR(state core.State, inst *disassemble.Inst, key uint8, ret bool) error {
	pcModifier, err := pcModifierOperand(state, inst)
	if err != nil {
		return err
	}
	modifier := combineModifiers(state.GetSP(), pcModifier)
	if ret {
		state.SetPC(authenticatePointer(state.GetX(linkRegister), modifier, key))
		return nil
	}
	authenticateRegister(state, linkRegister, modifier, key)
	return nil
}

// pcModifierOperand reads the PC modifier: a label address for the immediate forms, or the
// value of Xn (XZR reads as zero) for the register forms.
func pcModifierOperand(state core.State, inst *disassemble.Inst) (uint64, error) {
	if inst.NumOps < 1 {
		return 0, core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			inst.Operation.String(), "PAuth_LR authenticate requires a modifier operand")
	}
	op := inst.Operands[0]
	switch op.Class {
	case disassemble.LABEL:
		return op.Immediate, nil
	case disassemble.REG:
		if op.NumRegisters == 0 {
			break
		}
		if op.Registers[0] == disassemble.REG_XZR {
			return 0, nil
		}
		reg := core.MapRegister(op.Registers[0])
		if reg == -1 {
			return 0, core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
				inst.Operation.String(), fmt.Sprintf("invalid modifier register: %d", op.Registers[0]))
		}
		return state.GetX(reg), nil
	}
	return 0, core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
		inst.Operation.String(), "PAuth_LR modifier must be a label or X register")
}

func signRegister(state core.State, reg int, modifier uint64, key uint8) {
	state.SetX(reg, signPointer(state.GetX(reg), modifier, key))
}

func authenticateRegister(state core.State, reg int, modifier uint64, key uint8) {
	state.SetX(reg, authenticatePointer(state.GetX(reg), modifier, key))
}

// RegisterPAuthLRInstructions registers the FEAT_PAuth_LR instructions.
func RegisterPAuthLRInstructions(registry *Registry) {
	instructions := []struct{ mnemonic, description string }{
		{"PACM", "Pointer authentication modifier hint"},
		{"PACIASPPC", "Sign LR using SP and PC (A-key)"},
		{"PACIBSPPC", "Sign LR using SP and PC (B-key)"},
		{"PACNBIASPPC", "Sign LR using SP and PC, not a branch target (A-key)"},
		{"PACNBIBSPPC", "Sign LR using SP and PC, not a branch target (B-key)"},
		{"PACIA171615", "Sign X17 using X16 and X15 (A-key)"},
		{"PACIB171615", "Sign X17 using X16 and X15 (B-key)"},
		{"AUTIA171615", "Authenticate X17 using X16 and X15 (A-key)"},
		{"AUTIB171615", "Authenticate X17 using X16 and X15 (B-key)"},
		{"AUTIASPPC", "Authenticate LR using SP and a PC-relative label (A-key)"},
		{"AUTIBSPPC", "Authenticate LR using SP and a PC-relative label (B-key)"},
		{"AUTIASPPCR", "Authenticate LR using SP and a register modifier (A-key)"},
		{"AUTIBSPPCR", "Authenticate LR using SP and a register modifier (B-key)"},
		{"RETAASPPC", "Authenticate LR using SP and a PC-relative label, then return (A-key)"},
		{"RETABSPPC", "Authenticate LR using SP and a PC-relative label, then return (B-key)"},
		{"RETAASPPCR", "Authenticate LR using SP and a register modifier, then return (A-key)"},
		{"RETABSPPCR", "Authenticate LR using SP and a register modifier, then return (B-key)"},
	}
	for _, instr := range instructions {
		registry.Register(instr.mnemonic, NewPAuthLRExecutor(instr.mnemonic, instr.description))
	}
}
