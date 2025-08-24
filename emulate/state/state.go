package state

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/blacktop/arm64-cgo/emulate/core"
)

// ARM64State implements the core.State interface
type ARM64State struct {
	// General purpose registers X0-X30
	x [31]uint64
	// SIMD/FP registers Q0-Q31 (128-bit each, stored as [2]uint64)
	q [32][2]uint64
	// Stack pointer
	sp uint64
	// Program counter
	pc uint64
	// Processor state (condition flags: NZCV in bits 31-28)
	pstate uint64

	// Register validity tracking
	// Bitmask tracking which X registers hold valid values (bits 0-30)
	valid uint32
	// Bitmask tracking whether a register is 64-bit (X) or 32-bit (W) (bits 0-30)
	wide uint32
	// Bitmask tracking which Q registers hold valid values (bits 0-31)
	qvalid uint32

	// Host memory tracking
	// 32x 2-bit fields tracking memory validity
	host uint64

	// Emulation control and safety
	instructionCount uint64
	maxInstructions  uint64

	// Error tracking
	lastError string
	errorPC   uint64

	// Simple stack memory simulation
	stack     []byte
	stackBase uint64
	// Memory map for simple memory simulation (address -> data)
	memory map[uint64][]byte
	// Tracks explicitly mapped memory blocks (aligned to 8-byte)
	mapped map[uint64]bool
	// Custom handlers
	memReadHandler core.MemoryReadHandler
	stringHandler  core.StringPoolHandler
	ptrResolver    core.PointerResolver
}

// NewState creates a new ARM64 processor state with default values
func NewState() *ARM64State {
	return &ARM64State{
		x:                [31]uint64{},
		q:                [32][2]uint64{},
		sp:               0x7ffff0000000, // Default stack pointer
		pc:               0,
		pstate:           0,
		valid:            0, // All registers start as invalid
		wide:             0,
		qvalid:           0, // All Q registers start as invalid
		host:             0, // All host memory invalid
		instructionCount: 0,
		maxInstructions:  10000, // Default safety limit
		lastError:        "",
		errorPC:          0,
		stack:            make([]byte, 0x100000), // 1MB stack
		stackBase:        0x7ffff0000000,
		memory:           make(map[uint64][]byte),
		mapped:           make(map[uint64]bool),
		memReadHandler:   nil,
		stringHandler:    nil,
		ptrResolver:      nil,
	}
}

// NewStateWithConfig creates a new state with custom configuration
func NewStateWithConfig(config core.StateConfig) *ARM64State {
	stackSize := config.StackSize
	if stackSize == 0 {
		stackSize = 0x100000 // Default 1MB
	}

	stackBase := config.StackBase
	if stackBase == 0 {
		stackBase = 0x7ffff0000000 // Default stack base
	}

	state := &ARM64State{
		x:                [31]uint64{},
		q:                [32][2]uint64{},
		sp:               config.InitialSP,
		pc:               config.InitialPC,
		pstate:           0,
		valid:            0,
		wide:             0,
		qvalid:           0,
		host:             0,
		instructionCount: 0,
		maxInstructions:  10000, // Default safety limit
		lastError:        "",
		errorPC:          0,
		stack:            make([]byte, stackSize),
		stackBase:        stackBase,
		memory:           make(map[uint64][]byte),
		mapped:           make(map[uint64]bool),
		memReadHandler:   config.MemoryHandler,
		stringHandler:    config.StringHandler,
		ptrResolver:      config.PointerHandler,
	}

	if state.sp == 0 {
		state.sp = stackBase
	}

	return state
}

// Register access methods
func (s *ARM64State) GetX(reg int) uint64 {
	if reg < 0 || reg >= 31 {
		return 0
	}
	return s.x[reg]
}

func (s *ARM64State) SetX(reg int, val uint64) {
	if reg >= 0 && reg <= 30 {
		s.x[reg] = val
		s.setValid(reg, true)
		s.setWide(reg, true)
	} else if reg == 31 {
		// Register 31 can be SP or XZR depending on context
		// For now, allow writing to register 31 as SP
		s.sp = val
	}
}

func (s *ARM64State) GetW(reg int) uint32 {
	if reg < 0 || reg >= 31 {
		return 0
	}
	return uint32(s.x[reg])
}

func (s *ARM64State) SetW(reg int, val uint32) {
	if reg >= 0 && reg <= 30 {
		s.x[reg] = uint64(val)
		s.setValid(reg, true)
		s.setWide(reg, false)
	}
}

// Special register access
func (s *ARM64State) GetPC() uint64 {
	return s.pc
}

func (s *ARM64State) SetPC(addr uint64) {
	s.pc = addr
}

func (s *ARM64State) GetSP() uint64 {
	return s.sp
}

func (s *ARM64State) SetSP(addr uint64) {
	s.sp = addr
}

// Flag access methods
func (s *ARM64State) GetN() bool {
	return (s.pstate & (1 << 31)) != 0
}

func (s *ARM64State) SetN(val bool) {
	if val {
		s.pstate |= (1 << 31)
	} else {
		s.pstate &^= (1 << 31)
	}
}

func (s *ARM64State) GetZ() bool {
	return (s.pstate & (1 << 30)) != 0
}

func (s *ARM64State) SetZ(val bool) {
	if val {
		s.pstate |= (1 << 30)
	} else {
		s.pstate &^= (1 << 30)
	}
}

func (s *ARM64State) GetC() bool {
	return (s.pstate & (1 << 29)) != 0
}

func (s *ARM64State) SetC(val bool) {
	if val {
		s.pstate |= (1 << 29)
	} else {
		s.pstate &^= (1 << 29)
	}
}

func (s *ARM64State) GetV() bool {
	return (s.pstate & (1 << 28)) != 0
}

func (s *ARM64State) SetV(val bool) {
	if val {
		s.pstate |= (1 << 28)
	} else {
		s.pstate &^= (1 << 28)
	}
}

func (s *ARM64State) UpdateFlags(result uint64, isNegative bool) {
	s.SetN(isNegative)
	s.SetZ(result == 0)
}

// Register validity tracking methods

// SetRegisterValid marks a register as containing valid data
func (s *ARM64State) SetRegisterValid(reg uint8, valid bool) {
	if reg >= 31 {
		return
	}
	if valid {
		s.valid |= (1 << reg)
	} else {
		s.valid &= ^(uint32(1) << reg)
	}
}

// IsRegisterValid checks if a register contains valid data
func (s *ARM64State) IsRegisterValid(reg uint8) bool {
	if reg >= 31 {
		return false
	}
	return (s.valid & (1 << reg)) != 0
}

// SetRegisterWide marks a register as 64-bit (vs 32-bit)
func (s *ARM64State) SetRegisterWide(reg uint8, wide bool) {
	if reg >= 31 {
		return
	}
	if wide {
		s.wide |= (1 << reg)
	} else {
		s.wide &= ^(uint32(1) << reg)
	}
}

// IsRegisterWide checks if register is 64-bit
func (s *ARM64State) IsRegisterWide(reg uint8) bool {
	if reg >= 31 {
		return false
	}
	return (s.wide & (1 << reg)) != 0
}

// ClearRegisterValidity clears all validity bits
func (s *ARM64State) ClearRegisterValidity() {
	s.valid = 0
	s.wide = 0
	s.qvalid = 0
}

// GetValidityMask returns current validity bitmasks for debugging
func (s *ARM64State) GetValidityMask() (valid, wide, qvalid uint32) {
	return s.valid, s.wide, s.qvalid
}

// Legacy methods for backward compatibility (use int parameter like before)
func (s *ARM64State) setValid(reg int, valid bool) {
	if reg < 0 || reg >= 31 {
		return
	}
	s.SetRegisterValid(uint8(reg), valid)
}

func (s *ARM64State) IsValid(reg int) bool {
	if reg < 0 || reg >= 31 {
		return false
	}
	return s.IsRegisterValid(uint8(reg))
}

func (s *ARM64State) setWide(reg int, wide bool) {
	if reg < 0 || reg >= 31 {
		return
	}
	s.SetRegisterWide(uint8(reg), wide)
}

func (s *ARM64State) IsWide(reg int) bool {
	if reg < 0 || reg >= 31 {
		return false
	}
	return s.IsRegisterWide(uint8(reg))
}

// Q register access methods
// GetQ returns the value of a Q register as two 64-bit values (low, high)
func (s *ARM64State) GetQ(reg uint8) [2]uint64 {
	if reg >= 32 {
		return [2]uint64{0, 0}
	}
	return s.q[reg]
}

// SetQ sets the value of a Q register from two 64-bit values (low, high)
func (s *ARM64State) SetQ(reg uint8, low, high uint64) {
	if reg >= 32 {
		return
	}
	s.q[reg][0] = low
	s.q[reg][1] = high
	s.SetQRegisterValid(reg, true)
}

// SetQRegisterValid marks a Q register as containing valid data
func (s *ARM64State) SetQRegisterValid(reg uint8, valid bool) {
	if reg >= 32 {
		return
	}
	if valid {
		s.qvalid |= (1 << reg)
	} else {
		s.qvalid &= ^(uint32(1) << reg)
	}
}

// IsQRegisterValid checks if a Q register contains valid data
func (s *ARM64State) IsQRegisterValid(reg uint8) bool {
	if reg >= 32 {
		return false
	}
	return (s.qvalid & (1 << reg)) != 0
}

// Host memory validity tracking
// SetHostMemoryValid marks memory region as valid/invalid (2-bit level per index)
func (s *ARM64State) SetHostMemoryValid(index uint8, level uint8) {
	if index >= 32 || level > 3 { // Only 32 indices, 2 bits per index (0-3)
		return
	}
	// Clear the 2 bits for this index
	mask := uint64(3) << (index * 2)
	s.host &= ^mask
	// Set the new value
	s.host |= uint64(level) << (index * 2)
}

// IsHostMemoryValid checks if memory region is valid (returns 2-bit level)
func (s *ARM64State) IsHostMemoryValid(index uint8) uint8 {
	if index >= 32 {
		return 0
	}
	return uint8((s.host >> (index * 2)) & 3)
}

// ValidatePointer checks if a register contains a valid pointer within bounds
func (s *ARM64State) ValidatePointer(reg uint8, minAddr, maxAddr uint64) bool {
	if !s.IsRegisterValid(reg) || reg >= 31 {
		return false
	}
	addr := s.GetX(int(reg))
	return addr >= minAddr && addr <= maxAddr
}

// Memory operations
func (s *ARM64State) ReadMemory(addr uint64, size int) ([]byte, error) {
	// First try custom memory handler if available
	if s.memReadHandler != nil {
		data, err := s.memReadHandler(addr, size)
		if err == nil {
			return data, nil
		}
		// If custom handler fails, fall through to local memory
	}

	// Check if it's a stack access (any part of the read touches the stack)
	if addr >= s.stackBase && addr < s.stackBase+uint64(len(s.stack)) {
		offset := addr - s.stackBase
		if offset+uint64(size) > uint64(len(s.stack)) {
			return nil, fmt.Errorf("%w: stack read out of bounds at 0x%x", core.ErrMemoryAccess, addr)
		}
		return s.stack[offset : offset+uint64(size)], nil
	}

	// Also check if the read would start beyond the stack but the address is close to stack range
	if addr >= s.stackBase+uint64(len(s.stack)) && addr < s.stackBase+uint64(len(s.stack))+8 {
		return nil, fmt.Errorf("%w: read beyond stack bounds at 0x%x", core.ErrMemoryAccess, addr)
	}

	// Check if the entire requested range is within mapped regions
	startAddr := addr & ^uint64(7)
	endAddr := (addr + uint64(size) - 1) & ^uint64(7)

	// If this is a very high address (likely invalid), return error immediately
	if addr >= 0xFFFFFFFF00000000 {
		return nil, fmt.Errorf("%w: invalid memory address 0x%x", core.ErrMemoryAccess, addr)
	}

	// Check if the range touches any mapped regions
	hasAnyMapped := false
	for a := startAddr; a <= endAddr; a += 8 {
		if s.mapped[a] {
			hasAnyMapped = true
			break
		}
	}

	// If no part of the range is mapped, return error for out-of-bounds access
	if !hasAnyMapped {
		return nil, fmt.Errorf("%w: read from unmapped memory at 0x%x", core.ErrMemoryAccess, addr)
	}

	// Check local memory map - handle multi-block reads
	result := make([]byte, size)
	bytesRead := 0

	for bytesRead < size {
		currentAddr := addr + uint64(bytesRead)
		alignedAddr := currentAddr & ^uint64(7)
		offset := currentAddr - alignedAddr

		if data, exists := s.memory[alignedAddr]; exists {
			// Check if offset is within the data block
			if int(offset) >= len(data) {
				// Offset is beyond this block, advance to next byte and continue
				result[bytesRead] = 0
				bytesRead++
				continue
			}

			// How many bytes can we read from this block?
			availableInBlock := len(data) - int(offset)
			toRead := min(size-bytesRead, availableInBlock)

			if toRead > 0 {
				copy(result[bytesRead:bytesRead+toRead], data[offset:offset+uint64(toRead)])
				bytesRead += toRead
			} else {
				// No more data in this block, fill remaining with zeros
				break
			}
		} else {
			// No data in this block - check if it's mapped
			if s.mapped[alignedAddr] {
				// Mapped but no data - fill with zeros (valid for mapped regions)
				result[bytesRead] = 0
				bytesRead++
			} else {
				// Unmapped - this should not happen since we checked above
				break
			}
		}
	}

	// Fill any remaining bytes with zeros for mapped regions
	for bytesRead < size {
		result[bytesRead] = 0
		bytesRead++
	}

	return result, nil
}

func (s *ARM64State) WriteMemory(addr uint64, data []byte) {
	// Check if it's a stack access
	if addr >= s.stackBase && addr < s.stackBase+uint64(len(s.stack)) {
		offset := addr - s.stackBase
		if offset+uint64(len(data)) <= uint64(len(s.stack)) {
			copy(s.stack[offset:], data)
			return
		}
	}

	// Handle multi-block writes for memory map
	bytesWritten := 0

	for bytesWritten < len(data) {
		currentAddr := addr + uint64(bytesWritten)
		alignedAddr := currentAddr & ^uint64(7)
		offset := currentAddr - alignedAddr

		// Only allow writes to explicitly mapped blocks (or those already existing)
		if _, exists := s.memory[alignedAddr]; !exists {
			if !s.mapped[alignedAddr] {
				// Unmapped: ignore write to match hv behavior
				bytesWritten++
				continue
			}
			// If mapped but no storage yet, allocate
			s.memory[alignedAddr] = make([]byte, 8)
		}

		// How many bytes can we write to this block?
		availableInBlock := 8 - int(offset)
		toWrite := min(len(data)-bytesWritten, availableInBlock)

		// Extend the block if needed
		needed := int(offset) + toWrite
		if needed > len(s.memory[alignedAddr]) {
			newData := make([]byte, needed)
			copy(newData, s.memory[alignedAddr])
			s.memory[alignedAddr] = newData
		}

		// Copy data to this block
		copy(s.memory[alignedAddr][offset:offset+uint64(toWrite)], data[bytesWritten:bytesWritten+toWrite])
		bytesWritten += toWrite
	}
}

// MapRegion explicitly maps a memory region so subsequent reads/writes are permitted.
// The region is tracked as 8-byte aligned blocks like the memory backing store.
func (s *ARM64State) MapRegion(addr uint64, size int) {
	if size <= 0 {
		return
	}
	start := addr & ^uint64(7)
	end := addr + uint64(size)
	for a := start; a < end; a += 8 {
		s.mapped[a] = true
		if _, exists := s.memory[a]; !exists {
			s.memory[a] = make([]byte, 8)
		}
	}
}

func (s *ARM64State) ReadUint64(addr uint64) (uint64, error) {
	data, err := s.ReadMemory(addr, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(data), nil
}

func (s *ARM64State) ReadUint32(addr uint64) (uint32, error) {
	data, err := s.ReadMemory(addr, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(data), nil
}

func (s *ARM64State) WriteUint64(addr uint64, value uint64) {
	if addr%8 != 0 {
		// In a real scenario, this would be an alignment fault.
		// For now, we'll allow it but should be revisited.
	}

	// Auto-map the region if needed
	s.MapRegion(addr, 8)

	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	s.WriteMemory(addr, data)
}

func (s *ARM64State) WriteUint32(addr uint64, value uint32) {
	if addr%4 != 0 {
		// In a real scenario, this would be an alignment fault.
	}

	// Auto-map the region if needed
	s.MapRegion(addr, 4)

	data := make([]byte, 4)
	binary.LittleEndian.PutUint32(data, value)
	s.WriteMemory(addr, data)
}

// String operations
func (s *ARM64State) ReadString(addr uint64) (string, error) {
	// Try custom string handler first
	if s.stringHandler != nil {
		return s.stringHandler(addr)
	}

	// Read string byte by byte until null terminator
	var result strings.Builder
	for i := 0; i < 1024; i++ { // Limit to prevent infinite loops
		data, err := s.ReadMemory(addr+uint64(i), 1)
		if err != nil {
			return "", fmt.Errorf("failed to read string at 0x%x: %w", addr, err)
		}

		if data[0] == 0 {
			break
		}

		result.WriteByte(data[0])
	}

	return result.String(), nil
}

func (s *ARM64State) ReadStringWithLength(addr uint64, length int) (string, error) {
	if length <= 0 {
		return "", nil
	}

	data, err := s.ReadMemory(addr, length)
	if err != nil {
		return "", fmt.Errorf("failed to read string at 0x%x: %w", addr, err)
	}

	// Find null terminator
	nullPos := bytes.IndexByte(data, 0)
	if nullPos >= 0 {
		data = data[:nullPos]
	}

	return string(data), nil
}

// Pointer operations
func (s *ARM64State) ResolvePointer(addr uint64) (uint64, error) {
	if s.ptrResolver != nil {
		return s.ptrResolver(addr)
	}

	// Default implementation: read 8 bytes as pointer
	return s.ReadUint64(addr)
}

func (s *ARM64State) FollowPointerChain(addr uint64, maxDepth int) (uint64, error) {
	current := addr
	for depth := 0; depth < maxDepth; depth++ {
		next, err := s.ResolvePointer(current)
		if err != nil {
			// If we can't resolve further, return the current address
			if depth == 0 {
				return 0, fmt.Errorf("failed to resolve initial pointer at addr=0x%x: %w", current, err)
			}
			return current, nil
		}

		// If we're pointing to the same address, we've reached the end
		if next == current {
			return next, nil
		}

		// Try to read from next address to see if it contains another pointer
		// If it fails, then next is the final value
		_, err = s.ResolvePointer(next)
		if err != nil {
			// Can't resolve further, so next is the final address
			return next, nil
		}

		current = next
	}

	return current, nil
}

func (s *ARM64State) ReadStringAtPointer(ptrAddr uint64) (string, error) {
	strAddr, err := s.ResolvePointer(ptrAddr)
	if err != nil {
		return "", fmt.Errorf("failed to resolve string pointer at 0x%x: %w", ptrAddr, err)
	}

	return s.ReadString(strAddr)
}

// Handler management
func (s *ARM64State) SetMemoryReadHandler(handler core.MemoryReadHandler) {
	s.memReadHandler = handler
}

func (s *ARM64State) SetStringPoolHandler(handler core.StringPoolHandler) {
	s.stringHandler = handler
}

func (s *ARM64State) SetPointerResolver(resolver core.PointerResolver) {
	s.ptrResolver = resolver
}

// State management
func (s *ARM64State) Clone() core.State {
	clone := &ARM64State{
		x:                s.x, // Arrays are copied by value
		q:                s.q, // Arrays are copied by value
		sp:               s.sp,
		pc:               s.pc,
		pstate:           s.pstate,
		valid:            s.valid,
		wide:             s.wide,
		qvalid:           s.qvalid,
		host:             s.host,
		instructionCount: s.instructionCount,
		maxInstructions:  s.maxInstructions,
		lastError:        s.lastError,
		errorPC:          s.errorPC,
		stack:            make([]byte, len(s.stack)),
		stackBase:        s.stackBase,
		memory:           make(map[uint64][]byte),
		mapped:           make(map[uint64]bool),
		memReadHandler:   s.memReadHandler,
		stringHandler:    s.stringHandler,
		ptrResolver:      s.ptrResolver,
	}

	// Deep copy stack and memory
	copy(clone.stack, s.stack)
	for addr, data := range s.memory {
		clonedData := make([]byte, len(data))
		copy(clonedData, data)
		clone.memory[addr] = clonedData
	}

	// Deep copy mapped regions
	for addr, mapped := range s.mapped {
		clone.mapped[addr] = mapped
	}

	return clone
}

func (s *ARM64State) Reset() {
	s.x = [31]uint64{}
	s.q = [32][2]uint64{}
	s.sp = s.stackBase
	s.pc = 0
	s.pstate = 0
	s.valid = 0
	s.wide = 0
	s.qvalid = 0
	s.host = 0
	s.instructionCount = 0
	s.lastError = ""
	s.errorPC = 0

	// Clear stack
	for i := range s.stack {
		s.stack[i] = 0
	}

	// Clear memory map and mapped regions
	s.memory = make(map[uint64][]byte)
	s.mapped = make(map[uint64]bool)
}

// Utility methods for state initialization
func (s *ARM64State) SetupFunctionCall(returnAddr uint64) {
	s.SetX(30, returnAddr) // LR (Link Register)
}

func (s *ARM64State) SetupFunctionCallWithArgs(returnAddr uint64, args ...uint64) {
	s.SetX(30, returnAddr)
	for i, arg := range args {
		if i < 8 { // X0-X7 are argument registers
			s.SetX(i, arg)
		}
	}
}

func (s *ARM64State) SetStackPointer(addr uint64) {
	s.sp = addr
	s.stackBase = addr
}

func (s *ARM64State) LoadRegisters(values []uint64) {
	for i, val := range values {
		if i < 31 {
			s.SetX(i, val)
		}
	}
}

func (s *ARM64State) ClearRegisters() {
	s.x = [31]uint64{}
	s.valid = 0
	s.wide = 0
}

func (s *ARM64State) SetupTestCase(startAddr, returnAddr uint64) {
	s.pc = startAddr
	s.SetX(30, returnAddr)
}

// Context setup helpers
func (s *ARM64State) SetupKernelContext() {
	// Initialize for kernel-mode execution
	s.sp = 0xfffffff000000000 // Kernel stack
	s.stackBase = 0xfffffff000000000
	s.pstate = 0 // Kernel mode, interrupts enabled
}

func (s *ARM64State) SetupUserContext() {
	// Initialize for user-mode execution
	s.sp = 0x7ffff0000000 // User stack
	s.stackBase = 0x7ffff0000000
	s.pstate = 0 // User mode
}

// Emulation safety and error handling methods

// IncrementInstructionCount safely tracks emulation progress
func (s *ARM64State) IncrementInstructionCount() error {
	s.instructionCount++
	if s.instructionCount > s.maxInstructions {
		return fmt.Errorf("emulation exceeded maximum instructions (%d)", s.maxInstructions)
	}
	return nil
}

// SetEmulationError records emulation errors for debugging
func (s *ARM64State) SetEmulationError(pc uint64, err string) {
	s.errorPC = pc
	s.lastError = err
}

// HasEmulationError checks if emulation encountered errors
func (s *ARM64State) HasEmulationError() bool {
	return s.lastError != ""
}

// GetEmulationStatus returns current emulation status
func (s *ARM64State) GetEmulationStatus() (instructionCount uint64, hasError bool, errorMsg string) {
	return s.instructionCount, s.HasEmulationError(), s.lastError
}

// SetMaxInstructions sets the maximum instruction limit for emulation safety
func (s *ARM64State) SetMaxInstructions(max uint64) {
	s.maxInstructions = max
}

// GetInstructionCount returns the current instruction count
func (s *ARM64State) GetInstructionCount() uint64 {
	return s.instructionCount
}

// ResetInstructionCount resets the instruction counter
func (s *ARM64State) ResetInstructionCount() {
	s.instructionCount = 0
}

// ClearEmulationError clears any recorded emulation errors
func (s *ARM64State) ClearEmulationError() {
	s.lastError = ""
	s.errorPC = 0
}

// Debugging and validation methods

// DumpRegisterState returns human-readable register state with validity info
func (s *ARM64State) DumpRegisterState() string {
	var buf strings.Builder

	fmt.Fprintf(&buf, "PC: 0x%016x  SP: 0x%016x\n", s.pc, s.sp)
	fmt.Fprintf(&buf, "PSTATE: 0x%016x (N:%t Z:%t C:%t V:%t)\n",
		s.pstate, s.GetN(), s.GetZ(), s.GetC(), s.GetV())
	fmt.Fprintf(&buf, "Valid: 0x%08x  Wide: 0x%08x  QValid: 0x%08x\n",
		s.valid, s.wide, s.qvalid)
	fmt.Fprintf(&buf, "Host: 0x%016x  InstrCount: %d  MaxInstr: %d\n",
		s.host, s.instructionCount, s.maxInstructions)

	if s.HasEmulationError() {
		fmt.Fprintf(&buf, "ERROR at PC 0x%016x: %s\n", s.errorPC, s.lastError)
	}

	// Show X registers with validity indicators
	for i := range 31 {
		valid := s.IsRegisterValid(uint8(i))
		wide := s.IsRegisterWide(uint8(i))
		validChar := ' '
		if valid {
			if wide {
				validChar = 'X'
			} else {
				validChar = 'W'
			}
		}
		if i%2 == 0 {
			fmt.Fprintf(&buf, "\n")
		}
		fmt.Fprintf(&buf, "X%-2d:%c 0x%016x  ", i, validChar, s.x[i])
	}
	fmt.Fprintf(&buf, "\n")

	return buf.String()
}

// ValidateInternalState checks state consistency
func (s *ARM64State) ValidateInternalState() error {
	// Check register validity consistency
	if s.valid != 0 {
		for i := uint8(0); i < 31; i++ {
			if s.IsRegisterValid(i) && s.GetX(int(i)) == 0 {
				// This might be valid (register can contain 0), but worth noting
				continue
			}
		}
	}

	// Check Q register validity consistency
	if s.qvalid != 0 {
		for i := uint8(0); i < 32; i++ {
			if s.IsQRegisterValid(i) {
				q := s.GetQ(i)
				if q[0] == 0 && q[1] == 0 {
					// This might be valid (Q register can contain 0)
					continue
				}
			}
		}
	}

	// Check instruction count vs max
	if s.instructionCount > s.maxInstructions {
		return fmt.Errorf("instruction count (%d) exceeds maximum (%d)",
			s.instructionCount, s.maxInstructions)
	}

	return nil
}

// CompareState compares two states for testing (returns differences)
func (s *ARM64State) CompareState(other *ARM64State) []string {
	if other == nil {
		return []string{"other state is nil"}
	}

	var diffs []string

	// Compare basic state
	if s.pc != other.pc {
		diffs = append(diffs, fmt.Sprintf("PC: 0x%x vs 0x%x", s.pc, other.pc))
	}
	if s.sp != other.sp {
		diffs = append(diffs, fmt.Sprintf("SP: 0x%x vs 0x%x", s.sp, other.sp))
	}
	if s.pstate != other.pstate {
		diffs = append(diffs, fmt.Sprintf("PSTATE: 0x%x vs 0x%x", s.pstate, other.pstate))
	}

	// Compare validity masks
	if s.valid != other.valid {
		diffs = append(diffs, fmt.Sprintf("Valid: 0x%x vs 0x%x", s.valid, other.valid))
	}
	if s.wide != other.wide {
		diffs = append(diffs, fmt.Sprintf("Wide: 0x%x vs 0x%x", s.wide, other.wide))
	}
	if s.qvalid != other.qvalid {
		diffs = append(diffs, fmt.Sprintf("QValid: 0x%x vs 0x%x", s.qvalid, other.qvalid))
	}

	// Compare registers
	for i := 0; i < 31; i++ {
		if s.x[i] != other.x[i] {
			diffs = append(diffs, fmt.Sprintf("X%d: 0x%x vs 0x%x", i, s.x[i], other.x[i]))
		}
	}

	// Compare Q registers
	for i := 0; i < 32; i++ {
		if s.q[i] != other.q[i] {
			diffs = append(diffs, fmt.Sprintf("Q%d: [0x%x,0x%x] vs [0x%x,0x%x]",
				i, s.q[i][0], s.q[i][1], other.q[i][0], other.q[i][1]))
		}
	}

	// Compare emulation state
	if s.instructionCount != other.instructionCount {
		diffs = append(diffs, fmt.Sprintf("InstrCount: %d vs %d", s.instructionCount, other.instructionCount))
	}
	if s.lastError != other.lastError {
		diffs = append(diffs, fmt.Sprintf("LastError: %q vs %q", s.lastError, other.lastError))
	}

	return diffs
}

// Dump methods for debugging
func (s *ARM64State) DumpRegisters(w io.Writer) {
	fmt.Fprintf(w, "PC: 0x%016x  SP: 0x%016x\n", s.pc, s.sp)
	fmt.Fprintf(w, "PSTATE: 0x%016x (N:%t Z:%t C:%t V:%t)\n",
		s.pstate, s.GetN(), s.GetZ(), s.GetC(), s.GetV())

	for i := range 31 {
		if i%4 == 0 {
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "X%-2d: 0x%016x  ", i, s.x[i])
	}
	fmt.Fprintf(w, "\n")
}

func (s *ARM64State) DumpStack(w io.Writer, size int) {
	if size <= 0 {
		size = 256
	}

	fmt.Fprintf(w, "Stack (top %d bytes from SP=0x%x):\n", size, s.sp)

	if s.sp >= s.stackBase && s.sp < s.stackBase+uint64(len(s.stack)) {
		offset := s.sp - s.stackBase
		remaining := uint64(len(s.stack)) - offset
		if uint64(size) > remaining {
			size = int(remaining)
		}

		data := s.stack[offset : offset+uint64(size)]
		for i := 0; i < len(data); i += 16 {
			end := min(i+16, len(data))

			fmt.Fprintf(w, "0x%016x: ", s.sp+uint64(i))
			for j := i; j < end; j++ {
				fmt.Fprintf(w, "%02x ", data[j])
			}
			fmt.Fprintf(w, "\n")
		}
	}
}
