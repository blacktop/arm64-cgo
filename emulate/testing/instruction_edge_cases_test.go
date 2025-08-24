//go:build darwin

package testing

import (
	"testing"
)

// Edge cases and boundary conditions that could break emulator implementations
// These tests are critical for ensuring robustness

// TestArithmeticEdgeCases tests edge cases for arithmetic instructions
func TestArithmeticEdgeCases(t *testing.T) {
	testCases := []InstructionTestCase{
		// Overflow/underflow edge cases
		{
			Name:         "ADD_max_positive_overflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0x91}},    // ADD X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFFFFFFFFFF}, // Max positive 64-bit
			Description:  "ADD causing positive overflow",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000}, // Should wrap to negative
		},
		{
			Name:         "ADD_max_unsigned_overflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0x91}},    // ADD X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // Max unsigned 64-bit
			Description:  "ADD causing unsigned overflow",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // Should wrap to 0
		},
		{
			Name:         "SUB_min_negative_underflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xd1}},    // SUB X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000}, // Min negative 64-bit
			Description:  "SUB causing negative underflow",
			ExpectedRegs: map[int]uint64{0: 0x7FFFFFFFFFFFFFFF}, // Should wrap to positive
		},
		{
			Name:         "SUB_zero_underflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xd1}}, // SUB X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0},               // Zero
			Description:  "SUB from zero causing underflow",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // Should wrap to max unsigned
		},

		// 32-bit vs 64-bit boundary cases
		{
			Name:         "ADD_W_high_bits_cleared",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0x11}}, // ADD W0, W0, #1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "ADD 32-bit clears high bits",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // High 32 bits cleared
		},
		{
			Name:         "SUB_W_high_bits_cleared",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0x51}}, // SUB W0, W0, #1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0x123456780000000A},
			Description:  "SUB 32-bit clears high bits",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000009}, // High 32 bits cleared
		},

		// Zero operand cases
		{
			Name:         "ADD_with_XZR",
			Instructions: [][]byte{{0x1f, 0x00, 0x00, 0x8b}}, // ADD XZR, X0, X0 (result discarded)
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "ADD with XZR destination (result discarded)",
			ExpectedRegs: map[int]uint64{0: 42}, // X0 unchanged
		},
		{
			Name:         "ADD_XZR_operand",
			Instructions: [][]byte{{0x00, 0x00, 0x1f, 0x8b}}, // ADD X0, X0, XZR
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "ADD with XZR as operand (should read as 0)",
			ExpectedRegs: map[int]uint64{0: 42}, // X0 + 0 = X0
		},
		{
			Name:         "MUL_by_zero",
			Instructions: [][]byte{{0x00, 0x7c, 0x1f, 0x9b}}, // MUL X0, X0, XZR
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "MUL by zero (XZR)",
			ExpectedRegs: map[int]uint64{0: 0},
		},

		// Large immediate values
		{
			Name:         "ADD_max_immediate",
			Instructions: [][]byte{{0x00, 0xfc, 0x3f, 0x91}}, // ADD X0, X0, #4095 (max 12-bit immediate)
			InitialRegs:  map[int]uint64{0: 1},
			Description:  "ADD with maximum immediate value",
			ExpectedRegs: map[int]uint64{0: 4096},
		},
		{
			Name:         "SUB_max_immediate",
			Instructions: [][]byte{{0x00, 0xfc, 0x3f, 0xd1}}, // SUB X0, X0, #4095
			InitialRegs:  map[int]uint64{0: 5000},
			Description:  "SUB with maximum immediate value",
			ExpectedRegs: map[int]uint64{0: 905},
		},

		// Division edge cases
		{
			Name:         "UDIV_max_dividend_min_divisor",
			Instructions: [][]byte{{0x00, 0x08, 0xc1, 0x9a}}, // UDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 1},
			Description:  "UDIV maximum dividend by minimum divisor",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
		},
		{
			Name:         "SDIV_most_negative_by_minus_one",
			Instructions: [][]byte{{0x00, 0x0c, 0xc1, 0x9a}},                           // SDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000, 1: 0xFFFFFFFFFFFFFFFF}, // -1
			Description:  "SDIV most negative by -1 (overflow case)",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000}, // Should saturate
		},

		// Flag setting edge cases for ADDS/SUBS
		{
			Name:         "ADDS_carry_and_overflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xb1}}, // ADDS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFFFFFFFFFF},
			Description:  "ADDS causing both carry and signed overflow",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000},
			CheckFlags:   true, // Should set V (overflow) but not C (carry)
		},
		{
			Name:         "ADDS_unsigned_carry_no_overflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xb1}}, // ADDS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "ADDS causing unsigned carry but no signed overflow",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
			CheckFlags:   true, // Should set C (carry) and Z (zero) but not V
		},
		{
			Name:         "SUBS_borrow_and_overflow",
			Instructions: [][]byte{{0x00, 0x04, 0x00, 0xf1}}, // SUBS X0, X0, #1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000},
			Description:  "SUBS causing both borrow and signed overflow",
			ExpectedRegs: map[int]uint64{0: 0x7FFFFFFFFFFFFFFF},
			CheckFlags:   true, // Should set V (overflow) and C (no borrow)
		},

		// Multiply overflow cases
		{
			Name:         "MUL_overflow_64bit",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x9b}}, // MUL X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0xFFFFFFFFFFFFFFFF},
			Description:  "MUL causing 64-bit overflow",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000001}, // Low 64 bits of result
		},
		{
			Name:         "MADD_overflow_and_wraparound",
			Instructions: [][]byte{{0x00, 0x04, 0x02, 0x9b}}, // MADD X0, X0, X2, X1
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0xFFFFFFFFFFFFFFFF, 2: 0xFFFFFFFFFFFFFFFF},
			Description:  "MADD with maximum values causing wraparound",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000}, // Should wrap to 0
		},

		// Shift amount edge cases (though shifts are in other instructions)
		{
			Name:         "ADD_register_max_shift",
			Instructions: [][]byte{{0x00, 0xfc, 0x01, 0x8b}}, // ADD X0, X0, X1, LSL #63
			InitialRegs:  map[int]uint64{0: 0, 1: 1},
			Description:  "ADD with maximum shift amount",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000}, // 1 << 63
		},
	}

	runInstructionTests(t, "ArithmeticEdgeCases", testCases)
}

// TestMemoryEdgeCases tests edge cases for memory instructions
func TestMemoryEdgeCases(t *testing.T) {
	testCases := []InstructionTestCase{
		// Alignment edge cases
		{
			Name: "LDR_unaligned_address",
			Instructions: [][]byte{
				{0x01, 0x04, 0x00, 0xf9}, // STR X1, [X0, #8] - store to unaligned address
				{0x02, 0x04, 0x40, 0xf9}, // LDR X2, [X0, #8] - load from same unaligned address
			},
			InitialRegs: map[int]uint64{0: 0x10000001, 1: 0xFEDCBA9876543210}, // Unaligned address + test data
			Description: "LDR from unaligned address (store then load)",
		},
		{
			Name:         "STR_unaligned_address",
			Instructions: [][]byte{{0x01, 0x04, 0x00, 0xf9}},                   // STR X1, [X0, #8]
			InitialRegs:  map[int]uint64{0: 0x10000001, 1: 0x123456789ABCDEF0}, // Unaligned
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 32),
			},
			Description: "STR to unaligned address",
		},

		// Boundary addresses
		{
			Name: "LDR_at_memory_boundary",
			Instructions: [][]byte{
				{0x01, 0x00, 0x00, 0xf9}, // STR X1, [X0] - store to boundary address
				{0x02, 0x00, 0x40, 0xf9}, // LDR X2, [X0] - load from boundary address
			},
			InitialRegs: map[int]uint64{0: 0x10000FF8, 1: 0x9ABCDEF012345678}, // Boundary address + test data
			Description: "LDR at memory boundary (store then load)",
		},

		// Maximum offset cases
		{
			Name: "LDR_max_positive_offset",
			Instructions: [][]byte{
				{0x01, 0xfc, 0x3f, 0xf9}, // STR X1, [X0, #32760] - store at max offset
				{0x02, 0xfc, 0x7f, 0xf9}, // LDR X2, [X0, #32760] - load from max offset
			},
			InitialRegs: map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0},
			Description: "LDR with maximum positive offset (store then load)",
		},
		{
			Name: "LDUR_max_negative_offset",
			Instructions: [][]byte{
				{0x01, 0x00, 0x10, 0xf8}, // STUR X1, [X0, #-256] - store at negative offset
				{0x02, 0x00, 0x50, 0xf8}, // LDUR X2, [X0, #-256] - load from negative offset
			},
			InitialRegs: map[int]uint64{0: 0x10000100, 1: 0x9ABCDEF012345678},
			Description: "LDUR with maximum negative offset (store then load)",
		},

		// Pre/post-index edge cases
		{
			Name:         "LDR_pre_index_wrap",
			Instructions: [][]byte{{0x00, 0xfc, 0x5f, 0xf8}},    // LDR X0, [X0, #-1]!
			InitialRegs:  map[int]uint64{0: 0x0000000000000000}, // Will wrap around
			InitialMem: map[uint64][]byte{
				0xFFFFFFFFFFFFFFFF: {0x78, 0x56, 0x34, 0x12, 0xF0, 0xDE, 0xBC, 0x9A},
			},
			Description: "LDR pre-index causing address wrap",
		},

		// Load/store pair edge cases
		{
			Name: "LDP_overlapping_registers",
			Instructions: [][]byte{
				{0x01, 0x08, 0x00, 0xa9}, // STP X1, X2, [X0] - store pair
				{0x03, 0x00, 0x40, 0xa9}, // LDP X3, X0, [X0] - load pair with X0 as second dest
			},
			InitialRegs: map[int]uint64{0: 0x10000000, 1: 0x8877665544332211, 2: 0x00FFEEDDCCBBAA99},
			Description: "LDP with overlapping destination registers (store then load)",
		},
		{
			Name:         "STP_overlapping_registers",
			Instructions: [][]byte{{0x00, 0x00, 0x00, 0xa9}}, // STP X0, X0, [X0]
			InitialRegs:  map[int]uint64{0: 0x10000000},
			InitialMem: map[uint64][]byte{
				0x10000000: make([]byte, 32),
			},
			Description: "STP with same source register",
			// Should store X0 value to both memory locations
		},

		// Sign extension edge cases
		{
			Name: "LDRSB_positive_byte",
			Instructions: [][]byte{
				{0x01, 0x00, 0x00, 0x39}, // STRB W1, [X0] - store byte
				{0x02, 0x00, 0x80, 0x39}, // LDRSB X2, [X0] - load sign-extended byte
			},
			InitialRegs: map[int]uint64{0: 0x10000000, 1: 0x7F}, // Positive byte value
			Description: "LDRSB with positive byte (store then load)",
		},
		{
			Name: "LDRSW_boundary_value",
			Instructions: [][]byte{
				{0x01, 0x00, 0x00, 0xb9}, // STR W1, [X0] - store 32-bit value
				{0x02, 0x00, 0x80, 0xb9}, // LDRSW X2, [X0] - load sign-extended 32-bit
			},
			InitialRegs: map[int]uint64{0: 0x10000000, 1: 0x7FFFFFFF}, // Max positive 32-bit
			Description: "LDRSW with maximum positive 32-bit value (store then load)",
		},

		// Zero-length memory regions (edge case for our test framework)
		{
			Name: "LDR_from_zero_page",
			Instructions: [][]byte{
				{0x01, 0x00, 0x00, 0xf9}, // STR X1, [X0] - store value
				{0x02, 0x00, 0x40, 0xf9}, // LDR X2, [X0] - load it back
			},
			InitialRegs: map[int]uint64{0: 0x10000000, 1: 0xDEADBEEFCAFEBABE},
			Description: "LDR from memory page (store then load)",
		},
	}

	runInstructionTests(t, "MemoryEdgeCases", testCases)
}

// TestFlagBehaviorEdgeCases tests edge cases in flag setting behavior
func TestFlagBehaviorEdgeCases(t *testing.T) {
	testCases := []InstructionTestCase{
		// Flag combinations that are easy to get wrong
		{
			Name:         "ADDS_all_flags_clear",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xab}},                           // ADDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x4000000000000000, 1: 0x4000000000000000}, // 2^62 + 2^62
			Description:  "ADDS that should clear all flags",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000}, // Result is negative
			CheckFlags:   true,                                  // N=1, Z=0, C=0, V=1 (signed overflow)
		},
		{
			Name:         "ADDS_zero_plus_zero",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xab}}, // ADDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0, 1: 0},
			Description:  "ADDS zero plus zero (sets Z flag)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true, // N=0, Z=1, C=0, V=0
		},
		{
			Name:         "SUBS_equal_values",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xeb}}, // SUBS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0, 1: 0x123456789ABCDEF0},
			Description:  "SUBS equal values (sets Z and C flags)",
			ExpectedRegs: map[int]uint64{0: 0},
			CheckFlags:   true, // N=0, Z=1, C=1 (no borrow), V=0
		},
		{
			Name:         "SUBS_unsigned_underflow",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xeb}}, // SUBS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0, 1: 1},         // 0 - 1
			Description:  "SUBS causing unsigned underflow",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // -1 in unsigned
			CheckFlags:   true,                                  // N=1, Z=0, C=0 (borrow), V=0
		},

		// 32-bit flag behavior
		{
			Name:         "ADDS_W_overflow",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x2b}},  // ADDS W0, W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFF, 1: 1}, // Max positive 32-bit + 1
			Description:  "ADDS 32-bit overflow",
			ExpectedRegs: map[int]uint64{0: 0x0000000080000000}, // Negative in 32-bit
			CheckFlags:   true,                                  // N=1 (for 32-bit), Z=0, C=0, V=1
		},
		{
			Name:         "SUBS_W_underflow",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x6b}},  // SUBS W0, W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0x80000000, 1: 1}, // Min negative 32-bit - 1
			Description:  "SUBS 32-bit underflow",
			ExpectedRegs: map[int]uint64{0: 0x000000007FFFFFFF}, // Max positive 32-bit
			CheckFlags:   true,                                  // N=0 (for 32-bit), Z=0, C=1, V=1
		},

		// Logical instruction flag behavior
		{
			Name:         "ANDS_all_bits_set",
			Instructions: [][]byte{{0x00, 0x00, 0x00, 0xea}}, // ANDS X0, X0, X0
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "ANDS with all bits set (negative result)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			CheckFlags:   true, // N=1, Z=0, C=0, V=0
		},
		{
			Name:         "ANDS_W_sign_bit",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x6a}}, // ANDS W0, W0, W1 (32-bit)
			InitialRegs:  map[int]uint64{0: 0x80000000, 1: 0x80000000},
			Description:  "ANDS 32-bit with sign bit set",
			ExpectedRegs: map[int]uint64{0: 0x0000000080000000},
			CheckFlags:   true, // N=1 (for 32-bit), Z=0, C=0, V=0
		},

		// Compare instruction flag edge cases
		{
			Name:         "CMP_maximum_difference",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xeb}},                           // CMP X0, X1
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFFFFFFFFFF, 1: 0x8000000000000000}, // Max pos - max neg
			Description:  "CMP with maximum possible difference",
			CheckFlags:   true, // Should set multiple flags
		},
		{
			Name:         "CMN_negative_overflow",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xab}},                           // CMN X0, X1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000, 1: 0x8000000000000000}, // Min neg + min neg
			Description:  "CMN causing negative overflow",
			CheckFlags:   true, // Should set V flag (signed overflow)
		},

		// Conditional compare edge cases
		{
			Name: "CCMP_condition_false_nzcv",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (set condition to false)
				{0x2f, 0x08, 0x40, 0xfa}, // CCMP X1, #0, #15, EQ (all flags set via immediate)
			},
			InitialRegs: map[int]uint64{0: 2, 1: 5}, // Condition will be false
			Description: "CCMP with false condition using immediate NZCV=15",
			CheckFlags:  true, // Should have NZCV = 1111 (all flags set)
		},
		{
			Name: "CCMP_condition_true_compare",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x00, 0xf1}, // SUBS XZR, X0, #0 (sets Z flag)
				{0x20, 0x08, 0x40, 0xfa}, // CCMP X1, #0, #0, EQ (condition true, perform compare)
			},
			InitialRegs: map[int]uint64{0: 0, 1: 0}, // Both zero
			Description: "CCMP with true condition performing actual compare",
			CheckFlags:  true, // Should reflect result of X1 compared to 0
		},
	}

	runInstructionTests(t, "FlagEdgeCases", testCases)
}

// TestBitfieldEdgeCases tests edge cases for bitfield instructions
func TestBitfieldEdgeCases(t *testing.T) {
	testCases := []InstructionTestCase{
		// Boundary bit positions
		{
			Name:         "BFM_bit_0_only",
			Instructions: [][]byte{{0x20, 0x00, 0x40, 0xb3}}, // BFM X0, X1, #0, #0 (single bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFE, 1: 0x0000000000000001},
			Description:  "BFM extracting only bit 0",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // Bit 0 set
		},
		{
			Name:         "BFM_bit_63_only",
			Instructions: [][]byte{{0x20, 0xfc, 0x7f, 0xb3}}, // BFM X0, X1, #63, #63 (MSB only)
			InitialRegs:  map[int]uint64{0: 0x7FFFFFFFFFFFFFFF, 1: 0x8000000000000000},
			Description:  "BFM extracting only bit 63 (MSB)",
			ExpectedRegs: map[int]uint64{0: 0x7FFFFFFFFFFFFFFF}, // MSB set in field, lower bits unchanged
		},

		// Full width operations
		{
			Name:         "UBFM_extract_full_width",
			Instructions: [][]byte{{0x20, 0xfc, 0x40, 0xd3}}, // UBFM X0, X1, #0, #63 (all bits)
			InitialRegs:  map[int]uint64{0: 0xDEADBEEFCAFEBABE, 1: 0x123456789ABCDEF0},
			Description:  "UBFM extracting full width (effectively MOV)",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},
		{
			Name:         "SBFM_sign_extend_full_width",
			Instructions: [][]byte{{0x20, 0xfc, 0x40, 0x93}}, // SBFM X0, X1, #0, #63
			InitialRegs:  map[int]uint64{0: 0, 1: 0x8000000000000000},
			Description:  "SBFM with full width (should preserve sign)",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000},
		},

		// Wrap-around cases (immr > imms)
		{
			Name:         "BFM_wrap_around",
			Instructions: [][]byte{{0x20, 0x1c, 0x48, 0xb3}}, // BFM X0, X1, #8, #7 (wrap case)
			InitialRegs:  map[int]uint64{0: 0x0000000000000000, 1: 0xFF00000000000000},
			Description:  "BFM with wrap-around (immr > imms)",
		},

		// 32-bit vs 64-bit edge cases
		{
			Name:         "BFM_32bit_high_bits",
			Instructions: [][]byte{{0x20, 0x1c, 0x04, 0x33}}, // BFM W0, W1, #4, #7 (32-bit)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x123456789ABCDEF0},
			Description:  "BFM 32-bit should clear high 32 bits",
		},
		{
			Name:         "UBFM_32bit_zero_extend",
			Instructions: [][]byte{{0x20, 0x1c, 0x00, 0x53}}, // UBFM W0, W1, #0, #7 (32-bit UXTB)
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x123456789ABCDEFF},
			Description:  "UBFM 32-bit zero extension",
			ExpectedRegs: map[int]uint64{0: 0x00000000000000FF}, // High bits cleared
		},

		// Shift by maximum amounts
		{
			Name:         "LSL_max_shift",
			Instructions: [][]byte{{0x20, 0x00, 0x41, 0xd3}}, // LSL X0, X1, #63
			InitialRegs:  map[int]uint64{0: 0, 1: 0x0000000000000001},
			Description:  "LSL by maximum shift amount (63)",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000},
		},
		{
			Name:         "LSR_max_shift",
			Instructions: [][]byte{{0x20, 0xfc, 0x7f, 0xd3}}, // LSR X0, X1, #63 (UBFM alias)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x8000000000000000},
			Description:  "LSR by maximum shift amount (63)",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000001},
		},
		{
			Name:         "ASR_negative_max_shift",
			Instructions: [][]byte{{0x20, 0xfc, 0x7f, 0x93}},          // ASR X0, X1, #63 (SBFM alias)
			InitialRegs:  map[int]uint64{0: 0, 1: 0x8000000000000000}, // Negative number
			Description:  "ASR of negative number by maximum shift",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF}, // All bits set (sign extension)
		},

		// Zero shift cases
		{
			Name:         "LSL_zero_shift",
			Instructions: [][]byte{{0x20, 0xfc, 0x40, 0xd3}}, // LSL X0, X1, #0 (effectively MOV)
			InitialRegs:  map[int]uint64{0: 0xDEADBEEFCAFEBABE, 1: 0x123456789ABCDEF0},
			Description:  "LSL by zero (should be identity)",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},

		// Pattern with all bits set/clear
		{
			Name:         "BFM_all_ones_source",
			Instructions: [][]byte{{0x20, 0x1c, 0x44, 0xb3}}, // BFM X0, X1, #4, #7
			InitialRegs:  map[int]uint64{0: 0x0000000000000000, 1: 0xFFFFFFFFFFFFFFFF},
			Description:  "BFM with all ones in source",
		},
		{
			Name:         "BFM_all_zeros_destination",
			Instructions: [][]byte{{0x20, 0x1c, 0x44, 0xb3}}, // BFM X0, X1, #4, #7
			InitialRegs:  map[int]uint64{0: 0x0000000000000000, 1: 0x123456789ABCDEF0},
			Description:  "BFM with all zeros in destination",
		},
	}

	runInstructionTests(t, "BitfieldEdgeCases", testCases)
}

// TestSystemInstructionEdgeCases tests edge cases for system instructions
func TestSystemInstructionEdgeCases(t *testing.T) {
	testCases := []InstructionTestCase{
		// Multiple NOPs
		{
			Name: "multiple_NOPs",
			Instructions: [][]byte{
				{0x1f, 0x20, 0x03, 0xd5}, // NOP
				{0x1f, 0x20, 0x03, 0xd5}, // NOP
				{0x1f, 0x20, 0x03, 0xd5}, // NOP
			},
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "Multiple NOPs should have no effect",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},

		// HINT instruction variants
		{
			Name:         "HINT_reserved_value",
			Instructions: [][]byte{{0xff, 0x20, 0x03, 0xd5}}, // HINT #31 (reserved)
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "HINT with reserved value (should be NOP)",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},

		// Memory barriers with different scopes
		{
			Name:         "DSB_inner_shareable",
			Instructions: [][]byte{{0xbf, 0x33, 0x03, 0xd5}}, // DSB ISH
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "DSB inner shareable",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "DMB_outer_shareable",
			Instructions: [][]byte{{0xbf, 0x32, 0x03, 0xd5}}, // DMB OSH
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "DMB outer shareable",
			ExpectedRegs: map[int]uint64{0: 42},
		},

		// ISB variants
		{
			Name:         "ISB_with_immediate",
			Instructions: [][]byte{{0xff, 0x3f, 0x03, 0xd5}}, // ISB #15
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "ISB with immediate value",
			ExpectedRegs: map[int]uint64{0: 42},
		},
	}

	runInstructionTests(t, "SystemEdgeCases", testCases)
}
