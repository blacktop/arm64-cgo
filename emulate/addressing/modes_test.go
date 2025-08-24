package addressing

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// mockState implements core.State for testing
type mockState struct {
	registers [31]uint64
	sp        uint64
	pc        uint64
	flags     uint64
}

func newMockState() *mockState {
	return &mockState{
		registers: [31]uint64{},
		sp:        0x7ffff0000000,
		pc:        0x100000000,
		flags:     0,
	}
}

func (m *mockState) GetX(reg int) uint64 {
	if reg >= 0 && reg < 31 {
		return m.registers[reg]
	}
	return 0
}

func (m *mockState) SetX(reg int, val uint64) {
	if reg >= 0 && reg < 31 {
		m.registers[reg] = val
	}
}

func (m *mockState) GetW(reg int) uint32 {
	return uint32(m.GetX(reg))
}

func (m *mockState) SetW(reg int, val uint32) {
	m.SetX(reg, uint64(val))
}

func (m *mockState) GetPC() uint64     { return m.pc }
func (m *mockState) SetPC(addr uint64) { m.pc = addr }
func (m *mockState) GetSP() uint64     { return m.sp }
func (m *mockState) SetSP(addr uint64) { m.sp = addr }

func (m *mockState) GetN() bool { return (m.flags & (1 << 31)) != 0 }
func (m *mockState) SetN(val bool) {
	if val {
		m.flags |= (1 << 31)
	} else {
		m.flags &= ^uint64(1 << 31)
	}
}

func (m *mockState) GetZ() bool { return (m.flags & (1 << 30)) != 0 }
func (m *mockState) SetZ(val bool) {
	if val {
		m.flags |= (1 << 30)
	} else {
		m.flags &= ^uint64(1 << 30)
	}
}

func (m *mockState) GetC() bool { return (m.flags & (1 << 29)) != 0 }
func (m *mockState) SetC(val bool) {
	if val {
		m.flags |= (1 << 29)
	} else {
		m.flags &= ^uint64(1 << 29)
	}
}

func (m *mockState) GetV() bool { return (m.flags & (1 << 28)) != 0 }
func (m *mockState) SetV(val bool) {
	if val {
		m.flags |= (1 << 28)
	} else {
		m.flags &= ^uint64(1 << 28)
	}
}

func (m *mockState) UpdateFlags(result uint64, isNegative bool) {
	m.SetZ(result == 0)
	m.SetN(isNegative)
}

// Stub implementations for memory operations
func (m *mockState) ReadMemory(addr uint64, size int) ([]byte, error) {
	return make([]byte, size), nil
}
func (m *mockState) WriteMemory(addr uint64, data []byte)                         {}
func (m *mockState) ReadUint64(addr uint64) (uint64, error)                       { return 0, nil }
func (m *mockState) ReadUint32(addr uint64) (uint32, error)                       { return 0, nil }
func (m *mockState) WriteUint64(addr uint64, value uint64)                        {}
func (m *mockState) WriteUint32(addr uint64, value uint32)                        {}
func (m *mockState) ReadString(addr uint64) (string, error)                       { return "", nil }
func (m *mockState) ReadStringWithLength(addr uint64, length int) (string, error) { return "", nil }
func (m *mockState) ResolvePointer(addr uint64) (uint64, error)                   { return addr, nil }
func (m *mockState) FollowPointerChain(addr uint64, maxDepth int) (uint64, error) { return addr, nil }
func (m *mockState) ReadStringAtPointer(ptrAddr uint64) (string, error)           { return "", nil }
func (m *mockState) SetMemoryReadHandler(handler core.MemoryReadHandler)          {}
func (m *mockState) SetStringPoolHandler(handler core.StringPoolHandler)          {}
func (m *mockState) SetPointerResolver(resolver core.PointerResolver)             {}
func (m *mockState) Clone() core.State                                            { return &mockState{} }
func (m *mockState) Reset()                                                       {}

func TestCalculateOffsetAddress(t *testing.T) {
	state := newMockState()
	state.SetX(0, 0x1000) // X0 = 0x1000

	tests := []struct {
		name     string
		operand  disassemble.Operand
		expected uint64
	}{
		{
			name: "positive immediate offset",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_OFFSET,
				Registers: []disassemble.Register{34}, // X0
				Immediate: 0x10,
				SignedImm: false,
			},
			expected: 0x1010,
		},
		{
			name: "negative signed immediate offset",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_OFFSET,
				Registers: []disassemble.Register{34}, // X0
				Immediate: 0x1FF,                      // -1 in 9-bit signed
				SignedImm: true,
			},
			expected: 0xFFF, // 0x1000 + (-1)
		},
		{
			name: "zero offset",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_OFFSET,
				Registers: []disassemble.Register{34}, // X0
				Immediate: 0,
				SignedImm: false,
			},
			expected: 0x1000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := calculateOffsetAddress(state, tt.operand)
			if err != nil {
				t.Fatalf("calculateOffsetAddress() error = %v", err)
			}
			if calc.Address != tt.expected {
				t.Errorf("calculateOffsetAddress() = 0x%x, want 0x%x", calc.Address, tt.expected)
			}
			if calc.RequiresWriteback {
				t.Errorf("calculateOffsetAddress() should not require writeback")
			}
		})
	}
}

func TestCalculatePreIndexAddress(t *testing.T) {
	state := newMockState()
	state.SetX(1, 0x2000) // X1 = 0x2000

	operand := disassemble.Operand{
		Class:     disassemble.MEM_PRE_IDX,
		Registers: []disassemble.Register{35}, // X1
		Immediate: 0x20,
		SignedImm: false,
	}

	calc, err := calculatePreIndexAddress(state, operand)
	if err != nil {
		t.Fatalf("calculatePreIndexAddress() error = %v", err)
	}

	expectedAddr := uint64(0x2020) // 0x2000 + 0x20
	if calc.Address != expectedAddr {
		t.Errorf("calculatePreIndexAddress() address = 0x%x, want 0x%x", calc.Address, expectedAddr)
	}
	if calc.WritebackValue != expectedAddr {
		t.Errorf("calculatePreIndexAddress() writeback = 0x%x, want 0x%x", calc.WritebackValue, expectedAddr)
	}
	if !calc.RequiresWriteback {
		t.Errorf("calculatePreIndexAddress() should require writeback")
	}
	if calc.BaseRegister != 1 {
		t.Errorf("calculatePreIndexAddress() base register = %d, want 1", calc.BaseRegister)
	}
}

func TestCalculatePostIndexAddress(t *testing.T) {
	state := newMockState()
	state.SetX(2, 0x3000) // X2 = 0x3000

	operand := disassemble.Operand{
		Class:     disassemble.MEM_POST_IDX,
		Registers: []disassemble.Register{36}, // X2
		Immediate: 0x30,
		SignedImm: false,
	}

	calc, err := calculatePostIndexAddress(state, operand)
	if err != nil {
		t.Fatalf("calculatePostIndexAddress() error = %v", err)
	}

	originalAddr := uint64(0x3000)
	writebackAddr := uint64(0x3030) // 0x3000 + 0x30

	if calc.Address != originalAddr {
		t.Errorf("calculatePostIndexAddress() address = 0x%x, want 0x%x", calc.Address, originalAddr)
	}
	if calc.WritebackValue != writebackAddr {
		t.Errorf("calculatePostIndexAddress() writeback = 0x%x, want 0x%x", calc.WritebackValue, writebackAddr)
	}
	if !calc.RequiresWriteback {
		t.Errorf("calculatePostIndexAddress() should require writeback")
	}
	if calc.BaseRegister != 2 {
		t.Errorf("calculatePostIndexAddress() base register = %d, want 2", calc.BaseRegister)
	}
}

func TestCalculateRegisterAddress(t *testing.T) {
	state := newMockState()
	state.SetX(3, 0x4000) // X3 = 0x4000
	state.SetX(4, 0x100)  // X4 = 0x100

	tests := []struct {
		name     string
		operand  disassemble.Operand
		expected uint64
	}{
		{
			name: "base register only",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_REG,
				Registers: []disassemble.Register{37}, // X3
			},
			expected: 0x4000,
		},
		{
			name: "base + index register",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_REG,
				Registers: []disassemble.Register{37, 38}, // X3, X4
			},
			expected: 0x4100, // 0x4000 + 0x100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := calculateRegisterAddress(state, tt.operand)
			if err != nil {
				t.Fatalf("calculateRegisterAddress() error = %v", err)
			}
			if calc.Address != tt.expected {
				t.Errorf("calculateRegisterAddress() = 0x%x, want 0x%x", calc.Address, tt.expected)
			}
			if calc.RequiresWriteback {
				t.Errorf("calculateRegisterAddress() should not require writeback")
			}
		})
	}
}

func TestCalculateExtendedAddress(t *testing.T) {
	state := newMockState()
	state.SetX(5, 0x5000) // X5 = 0x5000
	state.SetX(6, 0x123)  // X6 = 0x123

	tests := []struct {
		name     string
		operand  disassemble.Operand
		expected uint64
	}{
		{
			name: "no extension",
			operand: disassemble.Operand{
				Class:          disassemble.MEM_EXTENDED,
				Registers:      []disassemble.Register{39, 40}, // X5, X6
				ShiftValueUsed: false,
			},
			expected: 0x5123, // 0x5000 + 0x123
		},
		{
			name: "with LSL extension",
			operand: disassemble.Operand{
				Class:          disassemble.MEM_EXTENDED,
				Registers:      []disassemble.Register{39, 40}, // X5, X6
				ShiftValueUsed: true,
				ShiftType:      ShiftTypeLSL,
				ShiftValue:     2,
			},
			expected: 0x548C, // 0x5000 + (0x123 << 2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := calculateExtendedAddress(state, tt.operand)
			if err != nil {
				t.Fatalf("calculateExtendedAddress() error = %v", err)
			}
			if calc.Address != tt.expected {
				t.Errorf("calculateExtendedAddress() = 0x%x, want 0x%x", calc.Address, tt.expected)
			}
			if calc.RequiresWriteback {
				t.Errorf("calculateExtendedAddress() should not require writeback")
			}
		})
	}
}

func TestApplyWriteback(t *testing.T) {
	state := newMockState()
	state.SetX(7, 0x6000) // X7 = 0x6000

	calc := &AddressCalculation{
		Address:           0x6020,
		WritebackValue:    0x6020,
		RequiresWriteback: true,
		BaseRegister:      7,
	}

	err := ApplyWriteback(state, calc)
	if err != nil {
		t.Fatalf("ApplyWriteback() error = %v", err)
	}

	if state.GetX(7) != 0x6020 {
		t.Errorf("ApplyWriteback() X7 = 0x%x, want 0x6020", state.GetX(7))
	}
}

func TestApplyWritebackSP(t *testing.T) {
	state := newMockState()
	state.SetSP(0x7000)

	calc := &AddressCalculation{
		Address:           0x7010,
		WritebackValue:    0x7010,
		RequiresWriteback: true,
		BaseRegister:      31, // SP
	}

	err := ApplyWriteback(state, calc)
	if err != nil {
		t.Fatalf("ApplyWriteback() error = %v", err)
	}

	if state.GetSP() != 0x7010 {
		t.Errorf("ApplyWriteback() SP = 0x%x, want 0x7010", state.GetSP())
	}
}

func TestApplyExtend(t *testing.T) {
	tests := []struct {
		name       string
		value      uint64
		extendType int
		amount     uint32
		expected   uint64
	}{
		{
			name:       "UXTW no shift",
			value:      0xFFFFFFFF12345678,
			extendType: ShiftTypeUXTW,
			amount:     0,
			expected:   0x12345678, // Zero-extend 32-bit
		},
		{
			name:       "UXTW with shift",
			value:      0x12345678,
			extendType: ShiftTypeUXTW,
			amount:     2,
			expected:   0x48D159E0, // Zero-extend then shift left 2
		},
		{
			name:       "SXTW positive",
			value:      0x12345678,
			extendType: ShiftTypeSXTW,
			amount:     0,
			expected:   0x12345678, // Sign-extend positive 32-bit
		},
		{
			name:       "SXTW negative",
			value:      0x80000000,
			extendType: ShiftTypeSXTW,
			amount:     0,
			expected:   0xFFFFFFFF80000000, // Sign-extend negative 32-bit
		},
		{
			name:       "UXTB",
			value:      0xFFFFFFFFFFFFFF80,
			extendType: ShiftTypeUXTB,
			amount:     0,
			expected:   0x80, // Zero-extend 8-bit
		},
		{
			name:       "SXTB positive",
			value:      0x7F,
			extendType: ShiftTypeSXTB,
			amount:     0,
			expected:   0x7F, // Sign-extend positive 8-bit
		},
		{
			name:       "SXTB negative",
			value:      0x80,
			extendType: ShiftTypeSXTB,
			amount:     0,
			expected:   0xFFFFFFFFFFFFFF80, // Sign-extend negative 8-bit
		},
		{
			name:       "UXTH",
			value:      0xFFFFFFFFFFFF8000,
			extendType: ShiftTypeUXTH,
			amount:     0,
			expected:   0x8000, // Zero-extend 16-bit
		},
		{
			name:       "SXTH positive",
			value:      0x7FFF,
			extendType: ShiftTypeSXTH,
			amount:     0,
			expected:   0x7FFF, // Sign-extend positive 16-bit
		},
		{
			name:       "SXTH negative",
			value:      0x8000,
			extendType: ShiftTypeSXTH,
			amount:     0,
			expected:   0xFFFFFFFFFFFF8000, // Sign-extend negative 16-bit
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyExtend(tt.value, tt.extendType, tt.amount)
			if result != tt.expected {
				t.Errorf("ApplyExtend() = 0x%x, want 0x%x", result, tt.expected)
			}
		})
	}
}

func TestApplyShift(t *testing.T) {
	tests := []struct {
		name      string
		value     uint64
		shiftType int
		amount    uint32
		expected  uint64
	}{
		{
			name:      "LSL",
			value:     0x12345678,
			shiftType: ShiftTypeLSL,
			amount:    4,
			expected:  0x123456780,
		},
		{
			name:      "LSR",
			value:     0x12345678,
			shiftType: ShiftTypeLSR,
			amount:    4,
			expected:  0x1234567,
		},
		{
			name:      "ASR positive",
			value:     0x12345678,
			shiftType: ShiftTypeASR,
			amount:    4,
			expected:  0x1234567,
		},
		{
			name:      "ASR negative",
			value:     0x8000000000000000,
			shiftType: ShiftTypeASR,
			amount:    4,
			expected:  0xF800000000000000,
		},
		{
			name:      "ROR",
			value:     0x12345678,
			shiftType: ShiftTypeROR,
			amount:    4,
			expected:  0x8000000001234567, // Rotate right 4 bits
		},
		{
			name:      "LSL overflow",
			value:     0x12345678,
			shiftType: ShiftTypeLSL,
			amount:    64,
			expected:  0, // Shift by 64 or more results in 0
		},
		{
			name:      "ASR large amount negative",
			value:     0x8000000000000000,
			shiftType: ShiftTypeASR,
			amount:    64,
			expected:  0xFFFFFFFFFFFFFFFF, // All 1s for negative
		},
		{
			name:      "ASR large amount positive",
			value:     0x7FFFFFFFFFFFFFFF,
			shiftType: ShiftTypeASR,
			amount:    64,
			expected:  0, // All 0s for positive
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyShift(tt.value, tt.shiftType, tt.amount)
			if result != tt.expected {
				t.Errorf("ApplyShift() = 0x%x, want 0x%x", result, tt.expected)
			}
		})
	}
}

func TestSignExtend(t *testing.T) {
	tests := []struct {
		name     string
		imm      uint64
		width    int
		expected int64
	}{
		{
			name:     "9-bit positive",
			imm:      0x100, // 256 in 9-bit (sign bit set, so negative)
			width:    9,
			expected: -256,
		},
		{
			name:     "9-bit negative",
			imm:      0x1FF, // -1 in 9-bit (all bits set)
			width:    9,
			expected: -1,
		},
		{
			name:     "9-bit max negative",
			imm:      0x100, // -256 in 9-bit (sign bit set, others clear)
			width:    9,
			expected: -256, // This is negative in 9-bit
		},
		{
			name:     "9-bit actual negative",
			imm:      0x180, // -128 in 9-bit
			width:    9,
			expected: -128,
		},
		{
			name:     "12-bit positive",
			imm:      0x7FF, // 2047 in 12-bit
			width:    12,
			expected: 2047,
		},
		{
			name:     "12-bit negative",
			imm:      0x800, // -2048 in 12-bit
			width:    12,
			expected: -2048,
		},
		{
			name:     "invalid width",
			imm:      0x123,
			width:    0,
			expected: 0x123, // Should return as-is
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := signExtend(tt.imm, tt.width)
			if result != tt.expected {
				t.Errorf("signExtend() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestMapRegister(t *testing.T) {
	tests := []struct {
		name     string
		regID    disassemble.Register
		expected int
	}{
		{
			name:     "X0",
			regID:    34,
			expected: 0,
		},
		{
			name:     "X30",
			regID:    64,
			expected: 30,
		},
		{
			name:     "W0",
			regID:    1,
			expected: 0,
		},
		{
			name:     "W30",
			regID:    31,
			expected: 30,
		},
		{
			name:     "invalid register",
			regID:    100,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapRegister(tt.regID)
			if result != tt.expected {
				t.Errorf("mapRegister() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCalculateAddressIntegration(t *testing.T) {
	state := newMockState()
	state.SetX(0, 0x1000)
	state.SetX(1, 0x2000)

	tests := []struct {
		name      string
		operand   disassemble.Operand
		expected  uint64
		writeback bool
	}{
		{
			name: "MEM_OFFSET",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_OFFSET,
				Registers: []disassemble.Register{34}, // X0
				Immediate: 0x10,
			},
			expected:  0x1010,
			writeback: false,
		},
		{
			name: "MEM_PRE_IDX",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_PRE_IDX,
				Registers: []disassemble.Register{35}, // X1
				Immediate: 0x20,
			},
			expected:  0x2020,
			writeback: true,
		},
		{
			name: "MEM_POST_IDX",
			operand: disassemble.Operand{
				Class:     disassemble.MEM_POST_IDX,
				Registers: []disassemble.Register{35}, // X1
				Immediate: 0x30,
			},
			expected:  0x2000, // Original address
			writeback: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := CalculateAddress(state, tt.operand)
			if err != nil {
				t.Fatalf("CalculateAddress() error = %v", err)
			}
			if calc.Address != tt.expected {
				t.Errorf("CalculateAddress() = 0x%x, want 0x%x", calc.Address, tt.expected)
			}
			if calc.RequiresWriteback != tt.writeback {
				t.Errorf("CalculateAddress() writeback = %v, want %v", calc.RequiresWriteback, tt.writeback)
			}
		})
	}
}
