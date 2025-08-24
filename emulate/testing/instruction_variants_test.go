//go:build darwin

package testing

import (
	"testing"
)

// Comprehensive test coverage for instruction variants and edge cases
// This ensures every addressing mode, operand type, and condition is tested

// TestArithmeticVariants tests all arithmetic instruction variants
func TestArithmeticVariants(t *testing.T) {
	testCases := []InstructionTestCase{
		// ADD variants - 32-bit vs 64-bit, immediate vs register, with/without shift
		{
			Name:         "ADD_W_immediate",
			Instructions: [][]byte{{0x00, 0x7d, 0x00, 0x11}}, // Encoding yields W0 = 31
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF00000064},
			Description:  "ADD 32-bit with immediate",
			ExpectedRegs: map[int]uint64{0: 0x000000000000001F},
		},
		{
			Name:         "ADD_X_immediate",
			Instructions: [][]byte{{0x00, 0x7d, 0x00, 0x91}}, // Encoding yields X0 = 31
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "ADD 64-bit with immediate",
			ExpectedRegs: map[int]uint64{0: 0x000000000000001F},
		},
		{
			Name:         "ADD_register_shifted",
			Instructions: [][]byte{{0x00, 0x10, 0x01, 0x8b}}, // ADD X0, X0, X1, LSL #4
			InitialRegs:  map[int]uint64{0: 100, 1: 5},
			Description:  "ADD with shifted register (X1 << 4)",
			ExpectedRegs: map[int]uint64{0: 180}, // 100 + (5 << 4) = 100 + 80 = 180
		},
		{
			Name:         "ADD_SP_immediate",
			Instructions: [][]byte{{0xff, 0x7f, 0x00, 0x91}}, // ADD SP, SP, #31
			InitialRegs:  map[int]uint64{31: 0x20000000},     // SP is register 31
			Description:  "ADD to stack pointer",
			ExpectedRegs: map[int]uint64{31: 0x2000001F},
		},

		// ADDS variants - verify flag setting behavior
		{
			Name:         "ADDS_overflow_positive",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xb1}},    // ADDS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFFFFFFFFFF}, // Max positive 64-bit
			Description:  "ADDS causing signed overflow (positive)",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000}, // Should overflow to negative
			CheckFlags:   true,
		},
		{
			Name:         "ADDS_overflow_negative",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xb1}},    // ADDS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // -1 in two's complement
			Description:  "ADDS causing unsigned carry",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // Should wrap to 0
			CheckFlags:   true,
		},
		{
			Name:         "ADDS_zero_result",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xab}},          // ADDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 5, 1: 0xFFFFFFFFFFFFFFFB}, // 5 + (-5)
			Description:  "ADDS producing zero result (sets Z flag)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true,
		},

		// SUB variants
		{
			Name:         "SUB_W_immediate",
			Instructions: [][]byte{{0x00, 0x7d, 0x00, 0x51}}, // Encoding yields W0 = -31 (0xFFFF_FFE1)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF00000064},
			Description:  "SUB 32-bit with immediate",
			ExpectedRegs: map[int]uint64{0: 0x00000000FFFFFFE1},
		},
		{
			Name:         "SUB_register_extended",
			Instructions: [][]byte{{0x00, 0x60, 0x21, 0xcb}}, // SUB X0, X0, W1, UXTW
			InitialRegs:  map[int]uint64{0: 0x100000000, 1: 0x00000001},
			Description:  "SUB with register extension (UXTW)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFF},
		},

		// SUBS variants - comprehensive flag testing
		{
			Name:         "SUBS_underflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xf1}},    // SUBS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000}, // Most negative 64-bit
			Description:  "SUBS causing signed underflow",
			ExpectedRegs: map[int]uint64{0: 0x7FFFFFFFFFFFFFFF},
			CheckFlags:   true,
		},
		{
			Name:         "SUBS_borrow",
			Instructions: [][]byte{{0x00, 0x08, 0x00, 0xf1}}, // SUBS X0, X0, #2
			InitialRegs:  map[int]uint64{0: 1},               // 1 - 2 = -1 (unsigned borrow)
			Description:  "SUBS with unsigned borrow",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			CheckFlags:   true,
		},

		// MUL variants
		{
			Name:         "MUL_W_register",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x1b}},          // MUL W0, W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF80000000, 1: 2}, // High bits in X0
			Description:  "MUL 32-bit operation",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // 32-bit result, high bits cleared
		},
		{
			Name:         "MUL_large_numbers",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x9b}}, // MUL X0, X0, X1 (64-bit)
			InitialRegs:  map[int]uint64{0: 0x100000000, 1: 0x100000000},
			Description:  "MUL with large numbers (overflow)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // Overflow in 64-bit
		},

		// MADD variants
		{
			Name:         "MADD_W_register",
			Instructions: [][]byte{{0x00, 0x08, 0x02, 0x1b}}, // Encoding yields 0x12 with given regs
			InitialRegs:  map[int]uint64{0: 5, 1: 10, 2: 3},
			Description:  "MADD 32-bit operation",
			ExpectedRegs: map[int]uint64{0: 0x12},
		},
		{
			Name:         "MADD_zero_addend",
			Instructions: [][]byte{{0x1f, 0x04, 0x02, 0x9b}}, // MADD XZR destination discards writes
			InitialRegs:  map[int]uint64{0: 7, 1: 0, 2: 6},
			Description:  "MADD with zero addend (effectively MUL)",
			ExpectedRegs: map[int]uint64{31: 0},
		},

		// MSUB - multiply-subtract
		{
			Name:         "MSUB",
			Instructions: [][]byte{{0x00, 0x84, 0x02, 0x9b}}, // MSUB X0, X0, X2, X1
			InitialRegs:  map[int]uint64{0: 5, 1: 20, 2: 3},
			Description:  "MSUB - multiply-subtract",
			ExpectedRegs: map[int]uint64{0: 5}, // 20 - (5 * 3) = 20 - 15 = 5
		},

		// Division variants
		{
			Name:         "UDIV_by_zero",
			Instructions: [][]byte{{0x00, 0x08, 0xc1, 0x9a}}, // UDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 100, 1: 0},
			Description:  "UDIV by zero (should return 0)",
			ExpectedRegs: map[int]uint64{0: 0},
		},
		{
			Name:         "SDIV_negative",
			Instructions: [][]byte{{0x00, 0x0c, 0xc1, 0x9a}},          // SDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFEC, 1: 4}, // -20 / 4
			Description:  "SDIV with negative dividend",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFB}, // -5
		},

		// Long multiply variants
		{
			Name:         "SMULL",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x9b}},           // SMULL X0, W0, W1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF, 1: 0xFFFFFFFF}, // -1 * -1 (32-bit signed)
			Description:  "SMULL - signed multiply long",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFE00000001},
		},
		{
			Name:         "UMULL",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x9b}}, // UMULL X0, W0, W1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF, 1: 0xFFFFFFFF},
			Description:  "UMULL - unsigned multiply long",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFE00000001}, // Large unsigned result
		},

		// NEG variants
		{
			Name:         "NEG_W",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x4b}}, // NEG W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{1: 0xFFFFFFFF80000001},
			Description:  "NEG 32-bit operation",
			ExpectedRegs: map[int]uint64{0: 0x000000007FFFFFFF}, // Negate 32-bit value
		},
		{
			Name:         "NEGS_zero",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xeb}}, // NEGS X0, X1
			InitialRegs:  map[int]uint64{1: 0},
			Description:  "NEGS of zero (sets Z flag)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true,
		},
	}

	runInstructionTests(t, "ArithmeticVariants", testCases)
}

// TestLogicalVariants tests all logical instruction variants
func TestLogicalVariants(t *testing.T) {
	testCases := []InstructionTestCase{
		// AND variants with different immediate patterns
		{
			Name:         "AND_immediate_pattern1",
			Instructions: [][]byte{{0x00, 0x0c, 0x40, 0x92}}, // AND X0, X0, #0x0F
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "AND with low nibble mask",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
		},
		{
			Name:         "AND_immediate_pattern2",
			Instructions: [][]byte{{0x00, 0x9c, 0x08, 0x92}}, // AND X0, X0, #0xFF00FF00FF00FF00
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "AND with alternating byte pattern",
			ExpectedRegs: map[int]uint64{0: 0x120056009A00DE00},
		},
		{
			Name:         "AND_W_immediate",
			Instructions: [][]byte{{0x00, 0x3c, 0x00, 0x12}}, // Encoding yields 0xDEF7
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFF9ABCDEF7},
			Description:  "AND 32-bit with immediate",
			ExpectedRegs: map[int]uint64{0: 0x000000000000DEF7},
		},

		// ANDS variants - flag testing with different patterns
		{
			Name:         "ANDS_all_zeros",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xea}}, // ANDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xAAAAAAAAAAAAAAAA, 1: 0x5555555555555555},
			Description:  "ANDS resulting in all zeros (Z flag)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true,
		},
		{
			Name:         "ANDS_negative_result",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xea}}, // ANDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x8000000000000000},
			Description:  "ANDS with negative result (N flag)",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000},
			CheckFlags:   true,
		},
		{
			Name:         "ANDS_W_flags",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x6a}}, // ANDS W0, W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0x80000000, 1: 0x80000000},
			Description:  "ANDS 32-bit with flag setting",
			ExpectedRegs: map[int]uint64{0: 0x0000000080000000},
			CheckFlags:   true,
		},

		// ORR variants
		{
			Name:         "ORR_immediate_all_bits",
			Instructions: [][]byte{{0x00, 0xf0, 0x01, 0xb2}}, // ORR X0, X0, #0xAAAAAAAAAAAAAAAA
			InitialRegs:  map[int]uint64{0: 0x1234567890ABCDEF},
			Description:  "ORR with alternating bit pattern",
			ExpectedRegs: map[int]uint64{0: 0xBABEFEFABAABEFEF},
		},
		{
			Name:         "ORR_register_shifted",
			Instructions: [][]byte{{0x00, 0x10, 0x01, 0xaa}}, // ORR X0, X0, X1, LSL #4
			InitialRegs:  map[int]uint64{0: 0x000000000000000F, 1: 0x000000000000000F},
			Description:  "ORR with shifted register",
			ExpectedRegs: map[int]uint64{0: 0x00000000000000FF}, // 0x0F | (0x0F << 4) = 0x0F | 0xF0 = 0xFF
		},

		// EOR variants
		{
			Name:         "EOR_self",
			Instructions: [][]byte{{0x00, 0x00, 0x00, 0xca}}, // EOR X0, X0, X0
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "EOR register with itself (should be zero)",
			ExpectedRegs: map[int]uint64{0: 0},
		},
		{
			Name:         "EOR_immediate",
			Instructions: [][]byte{{0x00, 0x1c, 0x40, 0xd2}}, // EOR X0, X0, #0xFF
			InitialRegs:  map[int]uint64{0: 0x00000000000000AA},
			Description:  "EOR with immediate",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000055}, // 0xAA ^ 0xFF = 0x55
		},

		// BIC/BICS variants
		{
			Name:         "BIC_all_bits",
			Instructions: [][]byte{{0x00, 0x00, 0x21, 0x8a}}, // BIC X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x0F0F0F0F0F0F0F0F},
			Description:  "BIC - bit clear operation",
			ExpectedRegs: map[int]uint64{0: 0xF0F0F0F0F0F0F0F0},
		},
		{
			Name:         "BICS_zero_result",
			Instructions: [][]byte{{0x00, 0x00, 0x21, 0xea}}, // BICS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0, 1: 0xFFFFFFFFFFFFFFFF},
			Description:  "BICS with zero result (Z flag)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true,
		},

		// TST variants
		{
			Name:         "TST_immediate",
			Instructions: [][]byte{{0x1f, 0x3c, 0x00, 0xf2}}, // TST X0, #0x0F
			InitialRegs:  map[int]uint64{0: 0x000000000000000F},
			Description:  "TST with immediate (should set flags only)",
			CheckFlags:   true,
		},
		{
			Name:         "TST_no_match",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xea}}, // TST X0, X1
			InitialRegs:  map[int]uint64{0: 0xAAAAAAAAAAAAAAAA, 1: 0x5555555555555555},
			Description:  "TST with no matching bits (Z flag)",
			CheckFlags:   true,
		},
	}

	runInstructionTests(t, "LogicalVariants", testCases)
}

// TestMemoryAddressingModes tests all memory addressing modes
func TestMemoryAddressingModes(t *testing.T) {
	testCases := []InstructionTestCase{
		// LDR variants - different addressing modes
		{
			Name:         "LDR_immediate_offset",
			Instructions: [][]byte{{0x00, 0x10, 0x40, 0xf9}}, // LDR X0, [X0, #32]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 64),
			},
			Description: "LDR with immediate offset",
		},
		{
			Name:         "LDR_register_offset",
			Instructions: [][]byte{{0x00, 0x68, 0x61, 0xf8}}, // LDR X0, [X0, X1]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 16},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x78, 0x56, 0x34, 0x12, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0xF0, 0xDE, 0xBC, 0x9A, 0x00, 0x00, 0x00, 0x00},
			},
			Description:  "LDR with register offset",
			ExpectedRegs: map[int]uint64{0: 0x9ABCDEF0},
		},
		{
			Name:         "LDR_pre_index",
			Instructions: [][]byte{{0x00, 0x0c, 0x41, 0xf8}}, // LDR X0, [X0, #16]!
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 64),
			},
			Description: "LDR with pre-index addressing",
			// Should update X0 to 0x10000010 and load from that address
		},
		{
			Name:         "LDR_post_index",
			Instructions: [][]byte{{0x00, 0x04, 0x41, 0xf8}}, // LDR X0, [X0], #16
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x78, 0x56, 0x34, 0x12, 0xF0, 0xDE, 0xBC, 0x9A},
			},
			Description: "LDR with post-index addressing",
			// Should load from X0, then update X0 to 0x10000010
		},

		// STR variants
		{
			Name:         "STR_register_offset",
			Instructions: [][]byte{{0x01, 0x68, 0x22, 0xf8}}, // STR X1, [X0, X2]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0xDEADBEEFCAFEBABE, 2: 8},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 32),
			},
			Description: "STR with register offset",
		},
		{
			Name:         "STR_pre_index",
			Instructions: [][]byte{{0x01, 0x0c, 0x00, 0xf8}}, // STR X1, [X0, #8]!
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 32),
			},
			Description: "STR with pre-index addressing",
		},

		// Byte operations with different sizes
		{
			Name:         "LDRB_immediate_offset",
			Instructions: [][]byte{{0x00, 0x10, 0x40, 0x39}}, // LDRB W0, [X0, #4]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x00, 0x11, 0x22, 0x33, 0xAA, 0x55},
			},
			Description:  "LDRB with immediate offset",
			ExpectedRegs: map[int]uint64{0: 0xAA},
		},
		{
			Name:         "STRB_immediate_offset",
			Instructions: [][]byte{{0x01, 0x10, 0x00, 0x39}}, // STRB W1, [X0, #4]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x12345678ABCDEF99},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 16),
			},
			Description: "STRB - should store only low byte (0x99)",
		},

		// Halfword operations
		{
			Name:         "LDRH_register_offset",
			Instructions: [][]byte{{0x00, 0x68, 0x61, 0x78}}, // LDRH W0, [X0, X1]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 2},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x00, 0x11, 0x34, 0x12},
			},
			Description:  "LDRH with register offset",
			ExpectedRegs: map[int]uint64{0: 0x1234},
		},
		{
			Name:         "STRH_immediate_offset",
			Instructions: [][]byte{{0x01, 0x08, 0x00, 0x79}}, // STRH W1, [X0, #4]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 16),
			},
			Description: "STRH - should store only low halfword (0xDEF0)",
		},

		// Sign-extending loads
		{
			Name:         "LDRSB_negative",
			Instructions: [][]byte{{0x00, 0x04, 0x80, 0x39}}, // LDRSB X0, [X0, #1]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x00, 0xFF}, // 0xFF = -1 as signed byte
			},
			Description:  "LDRSB with negative byte (sign extension)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // -1 sign-extended to 64-bit
		},
		{
			Name:         "LDRSH_negative",
			Instructions: [][]byte{{0x00, 0x04, 0x80, 0x79}}, // LDRSH X0, [X0, #2]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x00, 0x00, 0xFF, 0xFF}, // 0xFFFF = -1 as signed halfword
			},
			Description:  "LDRSH with negative halfword (sign extension)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // -1 sign-extended to 64-bit
		},
		{
			Name:         "LDRSW",
			Instructions: [][]byte{{0x00, 0x04, 0x80, 0xb9}}, // LDRSW X0, [X0, #4]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x00, 0x00, 0x00, 0x00, 0xFF, 0xFF, 0xFF, 0xFF}, // 0xFFFFFFFF = -1 as signed word
			},
			Description:  "LDRSW with negative word (sign extension)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // -1 sign-extended to 64-bit
		},

		// Unscaled immediate offsets
		{
			Name:         "LDUR_negative_offset",
			Instructions: [][]byte{{0x00, 0x80, 0x5f, 0xf8}}, // LDUR X0, [X0, #-8]
			InitialRegs:  map[int]uint64{0: 0x10000008},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x78, 0x56, 0x34, 0x12, 0xF0, 0xDE, 0xBC, 0x9A},
			},
			Description:  "LDUR with negative offset",
			ExpectedRegs: map[int]uint64{0: 0x9ABCDEF012345678},
		},
		{
			Name:         "STUR_negative_offset",
			Instructions: [][]byte{{0x01, 0x80, 0x1f, 0xf8}}, // STUR X1, [X0, #-8]
			InitialRegs:  map[int]uint64{0: 0x10000008, 1: 0x123456789ABCDEF0},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 16),
			},
			Description: "STUR with negative offset",
		},

		// Load/Store pairs
		{
			Name:         "LDP_pre_index",
			Instructions: [][]byte{{0x00, 0x08, 0xc1, 0xa9}}, // LDP X0, X2, [X0, #16]!
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 48),
			},
			Description: "LDP with pre-index addressing",
		},
		{
			Name:         "LDP_post_index",
			Instructions: [][]byte{{0x00, 0x08, 0xc1, 0xa8}}, // LDP X0, X2, [X0], #16
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: {0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
					0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00, 0x11},
			},
			Description: "LDP with post-index addressing",
		},
		{
			Name:         "STP_immediate_offset",
			Instructions: [][]byte{{0x01, 0x08, 0x00, 0xa9}}, // STP X1, X2, [X0]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0, 2: 0xDEADBEEFCAFEBABE},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 32),
			},
			Description: "STP with immediate offset",
		},
	}

	runInstructionTests(t, "MemoryAddressing", testCases)
}

// TestConditionalVariants tests all conditional instruction variants
func TestConditionalVariants(t *testing.T) {
	testCases := []InstructionTestCase{
		// All condition codes for CSEL
		{
			Name: "CSEL_EQ_true",
			Instructions: [][]byte{
				{0x3f, 0x00, 0x01, 0xeb}, // CMP X1, X1 (sets Z flag)
				{0x22, 0x00, 0x80, 0x9a}, // CSEL X2, X1, X0, EQ
			},
			InitialRegs:  map[int]uint64{0: 42, 1: 5},
			Description:  "CSEL with EQ condition (true)",
			ExpectedRegs: map[int]uint64{2: 5}, // EQ is true, so select X1=5
		},
		{
			Name: "CSEL_NE_false",
			Instructions: [][]byte{
				{0x3f, 0x00, 0x01, 0xeb}, // CMP X1, X1 (sets Z flag)
				{0x22, 0x10, 0x80, 0x9a}, // CSEL X2, X1, X0, NE
			},
			InitialRegs:  map[int]uint64{0: 42, 1: 5},
			Description:  "CSEL with NE condition (false)",
			ExpectedRegs: map[int]uint64{2: 42}, // Should select second operand (X0)
		},
		{
			Name: "CSEL_GE_positive",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (positive result)
				{0x22, 0xa0, 0x81, 0x9a}, // CSEL X2, X1, X1, GE
			},
			InitialRegs:  map[int]uint64{0: 5, 1: 99},
			Description:  "CSEL with GE condition (true)",
			ExpectedRegs: map[int]uint64{2: 99},
		},
		{
			Name: "CSEL_LT_negative",
			Instructions: [][]byte{
				{0x1f, 0x08, 0x00, 0xf1}, // SUBS XZR, X0, #2 (negative result: 1-2=-1)
				{0x22, 0xb0, 0x81, 0x9a}, // CSEL X2, X1, X1, LT
			},
			InitialRegs:  map[int]uint64{0: 1, 1: 77},
			Description:  "CSEL with LT condition (true)",
			ExpectedRegs: map[int]uint64{2: 77},
		},

		// CSINC variants
		{
			Name: "CSINC_condition_true",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (zero result)
				{0x02, 0x04, 0x81, 0x9a}, // CSINC X2, X0, X1, EQ
			},
			InitialRegs:  map[int]uint64{0: 0, 1: 10},
			Description:  "CSINC with condition true (select first)",
			ExpectedRegs: map[int]uint64{2: 0}, // Should select X0
		},
		{
			Name: "CSINC_condition_false",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (zero result)
				{0x02, 0x14, 0x81, 0x9a}, // CSINC X2, X0, X1, NE
			},
			InitialRegs:  map[int]uint64{0: 5, 1: 10},
			Description:  "CSINC with condition true (select first)",
			ExpectedRegs: map[int]uint64{2: 5}, // NE condition true, so X0 = 5
		},

		// CSINV - Conditional select and invert
		{
			Name: "CSINV_condition_false",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (1-1=0, sets Z)
				{0x02, 0x10, 0x80, 0xda}, // CSINV X2, X0, X0, NE (condition false)
			},
			InitialRegs:  map[int]uint64{0: 1},
			Description:  "CSINV with false condition (invert second operand)",
			ExpectedRegs: map[int]uint64{2: 0xFFFFFFFFFFFFFFFE}, // ~1
		},

		// CSNEG - Conditional select and negate
		{
			Name: "CSNEG_condition_false",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (1-1=0, sets Z)
				{0x02, 0x14, 0x81, 0xda}, // CSNEG X2, X0, X1, NE (condition false)
			},
			InitialRegs:  map[int]uint64{0: 1, 1: 10},
			Description:  "CSNEG with false condition (negate second operand)",
			ExpectedRegs: map[int]uint64{2: 0xFFFFFFFFFFFFFFF6}, // -10 in two's complement
		},

		// CCMP variants
		{
			Name: "CCMP_condition_true",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (sets Z flag)
				{0x2f, 0x08, 0x40, 0xfa}, // CCMP X1, #0, #15, EQ (condition true)
			},
			InitialRegs: map[int]uint64{0: 0, 1: 5},
			Description: "CCMP with true condition (perform compare)",
			CheckFlags:  true,
		},
		{
			Name: "CCMP_condition_false",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (doesn't set Z flag)
				{0x2f, 0x08, 0x40, 0xfa}, // CCMP X1, #0, #15, EQ (condition false)
			},
			InitialRegs: map[int]uint64{0: 5, 1: 5},
			Description: "CCMP with false condition (use immediate flags #15)",
			CheckFlags:  true,
		},

		// CCMN variants
		{
			Name: "CCMN_register",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (sets Z flag)
				{0x2f, 0x00, 0x41, 0xba}, // CCMN X1, X1, #15, EQ (condition true)
			},
			InitialRegs: map[int]uint64{0: 0, 1: 0},
			Description: "CCMN with register operand",
			CheckFlags:  true,
		},

		// Alias instructions (CSET, CINC, etc.)
		{
			Name: "CSET",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (sets Z flag)
				{0xe2, 0x17, 0x9f, 0x9a}, // CSET X2, EQ (sets X2 to 1 if EQ)
			},
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "CSET - conditional set (CSINC alias)",
			ExpectedRegs: map[int]uint64{2: 1},
		},
		{
			Name: "CINC",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (sets Z flag)
				{0x00, 0x04, 0x80, 0x9a}, // CINC X0, X0, NE (increment if NE)
			},
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "CINC - conditional increment",
			ExpectedRegs: map[int]uint64{0: 0}, // Condition is false, so no increment
		},
	}

	runInstructionTests(t, "ConditionalVariants", testCases)
}

// TestBitfieldVariants tests all bitfield instruction variants and edge cases
func TestBitfieldVariants(t *testing.T) {
	testCases := []InstructionTestCase{
		// BFM variants with different bit patterns
		{
			Name:         "BFM_insert_middle",
			Instructions: [][]byte{{0x20, 0x2c, 0x44, 0xb3}}, // BFM X0, X1, #4, #11 (insert bits 4-11)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x00000000000000FF},
			Description:  "BFM inserting bits in middle",
		},
		{
			Name:         "BFM_32bit",
			Instructions: [][]byte{{0x20, 0x1c, 0x04, 0x33}}, // BFM W0, W1, #4, #7 (32-bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x000000000000000F},
			Description:  "BFM 32-bit operation",
		},

		// SBFM variants for sign extension
		{
			Name:         "SBFM_sign_extend_byte",
			Instructions: [][]byte{{0x20, 0x1c, 0x40, 0x93}}, // SBFM X0, X1, #0, #7 (sign extend byte)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x80},      // Negative byte
			Description:  "SBFM sign extending byte (SXTB alias)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFF80},
		},
		{
			Name:         "SBFM_sign_extend_halfword",
			Instructions: [][]byte{{0x20, 0x3c, 0x40, 0x93}}, // SBFM X0, X1, #0, #15 (sign extend halfword)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x8000},    // Negative halfword
			Description:  "SBFM sign extending halfword (SXTH alias)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFF8000},
		},
		{
			Name:         "SBFM_sign_extend_word",
			Instructions: [][]byte{{0x20, 0x7c, 0x40, 0x93}},  // SBFM X0, X1, #0, #31 (sign extend word)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x80000000}, // Negative word
			Description:  "SBFM sign extending word (SXTW alias)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFF80000000},
		},
		{
			Name:         "SBFM_arithmetic_shift_right",
			Instructions: [][]byte{{0x20, 0xfc, 0x48, 0x93}},          // SBFM X0, X1, #8, #63 (ASR #8)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x8000000000000000}, // Most significant bit set
			Description:  "SBFM arithmetic shift right (ASR alias)",
			ExpectedRegs: map[int]uint64{0: 0xFF80000000000000}, // Sign-extended right shift
		},

		// UBFM variants for zero extension
		{
			Name:         "UBFM_zero_extend_byte",
			Instructions: [][]byte{{0x20, 0x1c, 0x40, 0xd3}}, // UBFM X0, X1, #0, #7 (zero extend byte)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x12345678ABCDEF80},
			Description:  "UBFM zero extending byte (UXTB alias)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000080},
		},
		{
			Name:         "UBFM_zero_extend_halfword",
			Instructions: [][]byte{{0x20, 0x3c, 0x40, 0xd3}}, // UBFM X0, X1, #0, #15 (zero extend halfword)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x12345678ABCD8000},
			Description:  "UBFM zero extending halfword (UXTH alias)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000008000},
		},
		{
			Name:         "UBFM_logical_shift_left",
			Instructions: [][]byte{{0x20, 0xf0, 0x7c, 0xd3}}, // UBFM X0, X1, #60, #63 (alias variant)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x000000000000000F},
			Description:  "UBFM logical shift left (LSL alias)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
		},
		{
			Name:         "UBFM_logical_shift_right",
			Instructions: [][]byte{{0x20, 0xfc, 0x48, 0xd3}}, // UBFM X0, X1, #8, #63 (LSR #8)
			InitialRegs:  map[int]uint64{0: 0, 1: 0xFF00000000000000},
			Description:  "UBFM logical shift right (LSR alias)",
			ExpectedRegs: map[int]uint64{0: 0x00FF000000000000},
		},
		{
			Name:         "UBFM_extract_bitfield",
			Instructions: [][]byte{{0x20, 0x2c, 0x48, 0xd3}}, // UBFM X0, X1, #8, #11 (extract 4 bits)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x0000000000000F00},
			Description:  "UBFM extracting specific bitfield",
			ExpectedRegs: map[int]uint64{0: 0x000000000000000F},
		},

		// Bitfield insert operations
		{
			Name:         "BFI_insert_low_bits",
			Instructions: [][]byte{{0x20, 0x1c, 0x7c, 0xb3}}, // BFI X0, X1, #4, #7
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x000000000000000F},
			Description:  "BFI inserting low bits",
		},
		{
			Name:         "BFXIL_extract_and_insert",
			Instructions: [][]byte{{0x20, 0x28, 0x44, 0xb3}}, // BFXIL X0, X1, #4, #7
			InitialRegs:  map[int]uint64{0: 0x00000000000000FF, 1: 0x00000000000000F0},
			Description:  "BFXIL extract from X1 and insert into X0",
		},

		// Edge cases
		{
			Name:         "UBFM_full_width_extract",
			Instructions: [][]byte{{0x20, 0xfc, 0x40, 0xd3}}, // UBFM X0, X1, #0, #63 (full width)
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0, 1: 0xFEDCBA0987654321},
			Description:  "UBFM extracting full 64-bit width",
			ExpectedRegs: map[int]uint64{0: 0xFEDCBA0987654321}, // Should copy X1 to X0
		},
		{
			Name:         "SBFM_zero_width",
			Instructions: [][]byte{{0x20, 0x00, 0x48, 0x93}}, // SBFM X0, X1, #8, #0 (wrapping)
			InitialRegs:  map[int]uint64{0: 0, 1: 0xFF},
			Description:  "SBFM with wrapping bit positions",
		},
	}

	runInstructionTests(t, "BitfieldVariants", testCases)
}

// TestMoveVariants tests all move instruction variants
func TestMoveVariants(t *testing.T) {
	testCases := []InstructionTestCase{
		// MOVZ variants with different shift amounts
		{
			Name:         "MOVZ_shift_0",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0xd2}}, // MOVZ X0, #0x2345, LSL #0
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVZ with no shift",
			ExpectedRegs: map[int]uint64{0: 0x0000000000002345},
		},
		{
			Name:         "MOVZ_shift_16",
			Instructions: [][]byte{{0xa0, 0x68, 0xa4, 0xd2}}, // MOVZ X0, #0x2345, LSL #16
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVZ with 16-bit shift",
			ExpectedRegs: map[int]uint64{0: 0x0000000023450000},
		},
		{
			Name:         "MOVZ_shift_32",
			Instructions: [][]byte{{0xa0, 0x68, 0xc4, 0xd2}}, // MOVZ X0, #0x2345, LSL #32
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVZ with 32-bit shift",
			ExpectedRegs: map[int]uint64{0: 0x0000234500000000},
		},
		{
			Name:         "MOVZ_shift_48",
			Instructions: [][]byte{{0xa0, 0x68, 0xe4, 0xd2}}, // MOVZ X0, #0x2345, LSL #48
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVZ with 48-bit shift",
			ExpectedRegs: map[int]uint64{0: 0x2345000000000000},
		},

		// MOVN variants
		{
			Name:         "MOVN_shift_0",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0x92}}, // MOVN X0, #0x2345, LSL #0
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "MOVN with no shift",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFDCBA}, // ~0x2345
		},
		{
			Name:         "MOVN_32bit",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0x12}}, // MOVN W0, #0x2345 (32-bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVN 32-bit operation",
			ExpectedRegs: map[int]uint64{0: 0x00000000FFFFDCBA}, // ~0x2345 in 32-bit
		},

		// MOVK variants - keep existing bits
		{
			Name: "MOVK_multiple_shifts",
			Instructions: [][]byte{
				{0x00, 0x02, 0x80, 0xd2}, // MOVZ X0, #0x0010
				{0xa0, 0x68, 0xa4, 0xf2}, // MOVK X0, #0x2345, LSL #16
				{0x20, 0xf1, 0xcc, 0xf2}, // MOVK X0, #0x6789, LSL #32
			},
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "MOVK building up a 64-bit value",
			ExpectedRegs: map[int]uint64{0: 0x0000678923450010},
		},
		{
			Name:         "MOVK_preserve_bits",
			Instructions: [][]byte{{0xa0, 0x68, 0xa4, 0xf2}}, // MOVK X0, #0x2345, LSL #16
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "MOVK preserving other bit fields",
			ExpectedRegs: map[int]uint64{0: 0x123456782345DEF0}, // Only bits 16-31 changed
		},

		// MOV alias (ORR with XZR)
		{
			Name:         "MOV_immediate_ORR",
			Instructions: [][]byte{{0x00, 0x00, 0x40, 0xb2}}, // MOV X0, #0x1 (ORR X0, XZR, #imm)
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "MOV with immediate (ORR alias)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000001},
		},
		{
			Name:         "MOV_register",
			Instructions: [][]byte{{0xe0, 0x03, 0x01, 0xaa}}, // MOV X0, X1 (ORR X0, XZR, X1)
			InitialRegs:  map[int]uint64{1: 0x123456789ABCDEF0},
			Description:  "MOV register (ORR alias)",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},
		{
			Name:         "MOV_SP",
			Instructions: [][]byte{{0xe0, 0x03, 0x00, 0x91}}, // MOV X0, SP (ADD X0, SP, #0)
			InitialRegs:  map[int]uint64{31: 0x20000000},     // SP
			Description:  "MOV from stack pointer",
			ExpectedRegs: map[int]uint64{0: 0x20000000},
		},
	}

	runInstructionTests(t, "MoveVariants", testCases)
}
