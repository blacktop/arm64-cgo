package instructions

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/addressing"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// MemoryExecutor handles memory load/store instructions
type MemoryExecutor struct {
	*BaseExecutor
}

// NewMemoryExecutor creates a new memory instruction executor
func NewMemoryExecutor(mnemonic, description string) *MemoryExecutor {
	return &MemoryExecutor{
		BaseExecutor: NewBaseExecutor(mnemonic, description),
	}
}

// Execute executes memory instructions
func (e *MemoryExecutor) Execute(state core.State, instr *disassemble.Instruction) error {
	if err := e.ValidateInstruction(instr); err != nil {
		return err
	}

	// Get the actual instruction mnemonic from the disassembler
	instrMnemonic := strings.ToUpper(instr.Operation.String())

	switch instrMnemonic {
	case "LDR":
		return e.executeLDR(state, instr)
	case "LDRB":
		return e.executeLDRB(state, instr)
	case "LDRH":
		return e.executeLDRH(state, instr)
	case "LDRSB":
		return e.executeLDRSB(state, instr)
	case "LDRSH":
		return e.executeLDRSH(state, instr)
	case "LDRSW":
		return e.executeLDRSW(state, instr)
	case "LDADDA":
		return e.executeLDADDA(state, instr)
	case "STR":
		return e.executeSTR(state, instr)
	case "STRB":
		return e.executeSTRB(state, instr)
	case "STRH":
		return e.executeSTRH(state, instr)
	case "LDP":
		return e.executeLDP(state, instr)
	case "STP":
		return e.executeSTP(state, instr)
	case "LDUR":
		return e.executeLDUR(state, instr)
	case "STUR":
		return e.executeSTUR(state, instr)
	default:
		return core.NewEmulationError(core.ErrUnsupportedFeature, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory instruction %s not implemented", instrMnemonic))
	}
}

// LDR - Load register (64-bit)
func (e *MemoryExecutor) executeLDR(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDR requires at least 2 operands")
	}

	// Check destination register
	if len(ops[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "no destination register")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	// Calculate address
	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}
	// Defer writeback until after loading to avoid overwriting destination
	shouldWriteback := calc.RequiresWriteback
	writebackValue := calc.WritebackValue
	writebackReg := calc.BaseRegister

	// Check if this is a 32-bit operation (W register)
	// W registers are 1-31, X registers are 34-64
	regID := uint32(ops[0].Registers[0])
	if regID >= 1 && regID <= 31 {
		// 32-bit load, zero-extend to 64-bit
		value, err := state.ReadUint32(calc.Address)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		state.SetW(dstReg, value)
	} else {
		// 64-bit load
		value, err := state.ReadUint64(calc.Address)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		state.SetX(dstReg, value)
	}

	// Apply writeback AFTER loading
	// ARM64 spec: When destination register equals base register in writeback operations,
	// the behavior is CONSTRAINED UNPREDICTABLE. We implement WBSUPPRESS (writeback suppressed)
	// to match common implementations like hypervisor framework
	if shouldWriteback && writebackReg != dstReg {
		state.SetX(writebackReg, writebackValue)
	}

	return nil
}

// LDRB - Load register byte (8-bit, zero-extended)
func (e *MemoryExecutor) executeLDRB(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDRB requires at least 2 operands")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Read 1 byte
	data, err := state.ReadMemory(calc.Address, 1)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
	}

	// Zero-extend to 64-bit
	value := uint64(data[0])
	state.SetX(dstReg, value)

	return nil
}

// LDRH - Load register halfword (16-bit, zero-extended)
func (e *MemoryExecutor) executeLDRH(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDRH requires at least 2 operands")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Read 2 bytes
	data, err := state.ReadMemory(calc.Address, 2)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
	}

	// Zero-extend to 64-bit
	value := uint64(binary.LittleEndian.Uint16(data))
	state.SetX(dstReg, value)

	return nil
}

// LDRSB - Load register signed byte (8-bit, sign-extended)
func (e *MemoryExecutor) executeLDRSB(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDRSB requires at least 2 operands")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Defer writeback until after loading to avoid overwriting destination
	shouldWriteback := calc.RequiresWriteback
	writebackValue := calc.WritebackValue
	writebackReg := calc.BaseRegister

	// Read 1 byte
	data, err := state.ReadMemory(calc.Address, 1)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
	}

	// Sign-extend to 64-bit
	value := uint64(int64(int8(data[0])))
	state.SetX(dstReg, value)

	// Apply writeback AFTER loading
	// ARM64 spec: When destination register equals base register in writeback operations,
	// the behavior is CONSTRAINED UNPREDICTABLE. We implement WBSUPPRESS (writeback suppressed)
	// to match common implementations like hypervisor framework
	if shouldWriteback && writebackReg != dstReg {
		state.SetX(writebackReg, writebackValue)
	}

	return nil
}

// LDRSH - Load register signed halfword (16-bit, sign-extended)
func (e *MemoryExecutor) executeLDRSH(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDRSH requires at least 2 operands")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Read 2 bytes
	data, err := state.ReadMemory(calc.Address, 2)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
	}

	// Sign-extend to 64-bit
	value := uint64(int64(int16(binary.LittleEndian.Uint16(data))))
	state.SetX(dstReg, value)

	return nil
}

// LDRSW - Load register signed word (32-bit, sign-extended)
func (e *MemoryExecutor) executeLDRSW(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDRSW requires at least 2 operands")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Defer writeback until after loading to avoid overwriting destination
	shouldWriteback := calc.RequiresWriteback
	writebackValue := calc.WritebackValue
	writebackReg := calc.BaseRegister

	// Read 4 bytes
	val32, err := state.ReadUint32(calc.Address)
	if err != nil {
		if errors.Is(err, core.ErrUnmappedMemory) {
			return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
		}
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", err))
	}

	// Sign-extend to 64-bit
	value := uint64(int64(int32(val32)))
	state.SetX(dstReg, value)

	// Apply writeback AFTER loading
	// ARM64 spec: When destination register equals base register in writeback operations,
	// the behavior is CONSTRAINED UNPREDICTABLE. We implement WBSUPPRESS (writeback suppressed)
	// to match common implementations like hypervisor framework
	if shouldWriteback && writebackReg != dstReg {
		state.SetX(writebackReg, writebackValue)
	}

	return nil
}

// STR - Store register
func (e *MemoryExecutor) executeSTR(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "STR requires at least 2 operands")
	}

	srcReg := core.MapRegister(ops[0].Registers[0])
	if srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Check if this is a 32-bit operation (W register)
	// W registers are 1-31, X registers are 34-64
	if uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31 {
		// 32-bit store
		value := state.GetW(srcReg)
		state.WriteUint32(calc.Address, value)
	} else {
		// 64-bit store
		value := state.GetX(srcReg)
		state.WriteUint64(calc.Address, value)
	}

	return nil
}

// LDADDA - Atomic load-add with acquire semantics
func (e *MemoryExecutor) executeLDADDA(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDADDA requires destination, source, and memory operands")
	}

	if len(ops[0].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "missing destination register")
	}
	if len(ops[1].Registers) == 0 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "missing source register")
	}

	dstReg := core.MapRegister(ops[0].Registers[0])
	srcReg := core.MapRegister(ops[1].Registers[0])
	if dstReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination register")
	}
	if srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	calc, err := addressing.CalculateAddress(state, ops[2])
	if err != nil {
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("address calculation failed: %v", err))
	}

	is32Bit := func(regID disassemble.Register) bool {
		id := uint32(regID)
		return (id >= 1 && id <= 32)
	}(ops[0].Registers[0])

	if is32Bit {
		oldVal, readErr := state.ReadUint32(calc.Address)
		if readErr != nil {
			if errors.Is(readErr, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", readErr))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", readErr))
		}

		addend := state.GetW(srcReg)
		newVal := oldVal + addend // uint32 arithmetic wraps automatically
		state.WriteUint32(calc.Address, newVal)
		state.SetW(dstReg, oldVal)
	} else {
		oldVal, readErr := state.ReadUint64(calc.Address)
		if readErr != nil {
			if errors.Is(readErr, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", readErr))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("memory read failed: %v", readErr))
		}

		addend := state.GetX(srcReg)
		newVal := oldVal + addend
		state.WriteUint64(calc.Address, newVal)
		state.SetX(dstReg, oldVal)
	}

	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("writeback failed: %v", err))
	}

	// Acquire semantics prevent later memory accesses from being reordered before this instruction.
	// In this single-threaded emulator model, we treat it as a full completion of the RMW cycle.
	return nil
}

// STRB - Store register byte
func (e *MemoryExecutor) executeSTRB(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "STRB requires at least 2 operands")
	}

	srcReg := core.MapRegister(ops[0].Registers[0])
	if srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Store only the low byte
	value := uint8(state.GetX(srcReg))
	state.WriteMemory(calc.Address, []byte{value})

	return nil
}

// STRH - Store register halfword
func (e *MemoryExecutor) executeSTRH(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 2 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "STRH requires at least 2 operands")
	}

	srcReg := core.MapRegister(ops[0].Registers[0])
	if srcReg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source register")
	}

	calc, err := addressing.CalculateAddress(state, ops[1])
	if err != nil {
		return err
	}

	// Apply writeback if required
	if err := addressing.ApplyWriteback(state, calc); err != nil {
		return err
	}

	// Store only the low 16 bits
	value := uint16(state.GetX(srcReg))
	data := make([]byte, 2)
	binary.LittleEndian.PutUint16(data, value)
	state.WriteMemory(calc.Address, data)

	return nil
}

// LDP - Load pair of registers
func (e *MemoryExecutor) executeLDP(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "LDP requires at least 3 operands")
	}

	dst1Reg := core.MapRegister(ops[0].Registers[0])
	dst2Reg := core.MapRegister(ops[1].Registers[0])
	if dst1Reg == -1 || dst2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid destination registers")
	}

	calc, err := addressing.CalculateAddress(state, ops[2])
	if err != nil {
		return err
	}

	// Defer writeback until after loading to avoid overwriting destinations
	shouldWriteback := calc.RequiresWriteback

	// Determine if this is 32-bit or 64-bit operation
	// W registers are 1-31, X registers are 34-64
	is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31

	if is32Bit {
		// Load two 32-bit values
		val1, err := state.ReadUint32(calc.Address)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("first read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("first read failed: %v", err))
		}
		val2, err := state.ReadUint32(calc.Address + 4)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("second read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("second read failed: %v", err))
		}
		state.SetW(dst1Reg, val1)
		state.SetW(dst2Reg, val2)
	} else {
		// Load two 64-bit values
		val1, err := state.ReadUint64(calc.Address)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("first read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("first read failed: %v", err))
		}
		val2, err := state.ReadUint64(calc.Address + 8)
		if err != nil {
			if errors.Is(err, core.ErrUnmappedMemory) {
				return core.NewEmulationError(core.ErrUnmappedMemory, state.GetPC(),
					fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("second read failed: %v", err))
			}
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("second read failed: %v", err))
		}
		state.SetX(dst1Reg, val1)
		state.SetX(dst2Reg, val2)
	}

	// Apply writeback AFTER loading (handles SP correctly)
	if shouldWriteback {
		if err := addressing.ApplyWriteback(state, calc); err != nil {
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("writeback failed: %v", err))
		}
	}

	return nil
}

// STP - Store pair of registers
func (e *MemoryExecutor) executeSTP(state core.State, instr *disassemble.Instruction) error {
	ops := instr.Operands
	if len(ops) < 3 {
		return core.NewEmulationError(core.ErrInvalidInstruction, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "STP requires at least 3 operands")
	}

	src1Reg := core.MapRegister(ops[0].Registers[0])
	src2Reg := core.MapRegister(ops[1].Registers[0])
	if src1Reg == -1 || src2Reg == -1 {
		return core.NewEmulationError(core.ErrInvalidRegister, state.GetPC(),
			fmt.Sprintf("%v", instr.Operation), "invalid source registers")
	}

	calc, err := addressing.CalculateAddress(state, ops[2])
	if err != nil {
		return err
	}

	// Defer writeback until after storing to avoid overwriting source registers
	shouldWriteback := calc.RequiresWriteback

	// Determine if this is 32-bit or 64-bit operation
	// W registers are 1-31, X registers are 34-64
	is32Bit := uint32(ops[0].Registers[0]) >= 1 && uint32(ops[0].Registers[0]) <= 31

	if is32Bit {
		// Store two 32-bit values
		val1 := state.GetW(src1Reg)
		val2 := state.GetW(src2Reg)
		state.WriteUint32(calc.Address, val1)
		state.WriteUint32(calc.Address+4, val2)
	} else {
		// Store two 64-bit values
		val1 := state.GetX(src1Reg)
		val2 := state.GetX(src2Reg)
		state.WriteUint64(calc.Address, val1)
		state.WriteUint64(calc.Address+8, val2)
	}

	// Apply writeback AFTER storing (handles SP correctly)
	if shouldWriteback {
		if err := addressing.ApplyWriteback(state, calc); err != nil {
			return core.NewEmulationError(core.ErrMemoryAccess, state.GetPC(),
				fmt.Sprintf("%v", instr.Operation), fmt.Sprintf("writeback failed: %v", err))
		}
	}

	return nil
}

// LDUR - Load register (unscaled offset)
func (e *MemoryExecutor) executeLDUR(state core.State, instr *disassemble.Instruction) error {
	// LDUR is similar to LDR but with unscaled immediate offset
	// For emulation purposes, we can treat it the same as LDR
	return e.executeLDR(state, instr)
}

// STUR - Store register (unscaled offset)
func (e *MemoryExecutor) executeSTUR(state core.State, instr *disassemble.Instruction) error {
	// STUR is similar to STR but with unscaled immediate offset
	// For emulation purposes, we can treat it the same as STR
	return e.executeSTR(state, instr)
}

// RegisterMemoryInstructions registers all memory instructions
func RegisterMemoryInstructions(registry *Registry) {
	// Load instructions
	registry.Register("LDR", NewMemoryExecutor("LDR", "Load register"))
	registry.Register("LDRB", NewMemoryExecutor("LDRB", "Load register byte (zero-extended)"))
	registry.Register("LDRH", NewMemoryExecutor("LDRH", "Load register halfword (zero-extended)"))
	registry.Register("LDRSB", NewMemoryExecutor("LDRSB", "Load register signed byte"))
	registry.Register("LDRSH", NewMemoryExecutor("LDRSH", "Load register signed halfword"))
	registry.Register("LDRSW", NewMemoryExecutor("LDRSW", "Load register signed word"))
	registry.Register("LDUR", NewMemoryExecutor("LDUR", "Load register (unscaled offset)"))
	// Atomic load/store operations
	registry.Register("LDADDA", NewMemoryExecutor("LDADDA", "Atomic load-add with acquire semantics"))
	// Store instructions
	registry.Register("STR", NewMemoryExecutor("STR", "Store register"))
	registry.Register("STRB", NewMemoryExecutor("STRB", "Store register byte"))
	registry.Register("STRH", NewMemoryExecutor("STRH", "Store register halfword"))
	registry.Register("STUR", NewMemoryExecutor("STUR", "Store register (unscaled offset)"))

	// Load/store pair instructions
	registry.Register("LDP", NewMemoryExecutor("LDP", "Load pair of registers"))
	registry.Register("STP", NewMemoryExecutor("STP", "Store pair of registers"))
}

// Additional helper functions for memory operations

// IsLoadInstruction checks if a mnemonic is a load instruction
func IsLoadInstruction(mnemonic string) bool {
	switch mnemonic {
	case "LDR", "LDRB", "LDRH", "LDRSB", "LDRSH", "LDRSW", "LDUR", "LDP", "LDADDA":
		return true
	default:
		return false
	}
}

// IsStoreInstruction checks if a mnemonic is a store instruction
func IsStoreInstruction(mnemonic string) bool {
	switch mnemonic {
	case "STR", "STRB", "STRH", "STUR", "STP", "LDADDA":
		return true
	default:
		return false
	}
}

// IsPairInstruction checks if a mnemonic is a pair instruction
func IsPairInstruction(mnemonic string) bool {
	switch mnemonic {
	case "LDP", "STP":
		return true
	default:
		return false
	}
}

// GetMemoryAccessSize returns the size in bytes of a memory access
func GetMemoryAccessSize(mnemonic string, is32Bit bool) int {
	switch mnemonic {
	case "LDRB", "STRB", "LDRSB":
		return 1
	case "LDRH", "STRH", "LDRSH":
		return 2
	case "LDRSW":
		return 4
	case "LDR", "STR", "LDUR", "STUR":
		if is32Bit {
			return 4
		}
		return 8
	case "LDADDA":
		if is32Bit {
			return 4
		}
		return 8
	case "LDP", "STP":
		if is32Bit {
			return 8 // Two 32-bit values
		}
		return 16 // Two 64-bit values
	default:
		return 0
	}
}
