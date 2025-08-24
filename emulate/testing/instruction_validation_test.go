//go:build darwin

package testing

import (
	"fmt"
	"testing"
)

// Comprehensive per-instruction validation against Apple's Hypervisor Framework
// Every instruction we support should have a corresponding hypervisor test

// InstructionTestCase represents a single instruction test case
type InstructionTestCase struct {
	Name         string
	Instructions [][]byte
	InitialRegs  map[int]uint64
	InitialMem   map[uint64][]byte
	Description  string
	ExpectedRegs map[int]uint64 // Optional: for additional validation
	CheckFlags   bool           // Whether to validate flag states
}

// TestArithmeticInstructions tests all arithmetic instructions against hypervisor
func TestArithmeticInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// ADD - Immediate
		{
			Name:         "ADD_immediate",
			Instructions: [][]byte{{0x00, 0x7c, 0x00, 0x91}}, // ADD X0, X0, #31
			InitialRegs:  map[int]uint64{0: 100},
			Description:  "ADD with immediate value",
			ExpectedRegs: map[int]uint64{0: 131},
		},
		// ADD - Register
		{
			Name:         "ADD_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x8b}}, // ADD X0, X0, X1
			InitialRegs:  map[int]uint64{0: 50, 1: 25},
			Description:  "ADD with register operand",
			ExpectedRegs: map[int]uint64{0: 75},
		},
		// ADDS - Immediate (flag setting)
		{
			Name:         "ADDS_immediate",
			Instructions: [][]byte{{0x00, 0x7c, 0x00, 0xb1}}, // ADDS X0, X0, #31
			InitialRegs:  map[int]uint64{0: 100},
			Description:  "ADDS with immediate - sets flags",
			ExpectedRegs: map[int]uint64{0: 131},
			CheckFlags:   true,
		},
		// ADDS - Register (flag setting)
		{
			Name:         "ADDS_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xab}}, // ADDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 50, 1: 25},
			Description:  "ADDS with register - sets flags",
			ExpectedRegs: map[int]uint64{0: 75},
			CheckFlags:   true,
		},
		// SUB - Immediate
		{
			Name:         "SUB_immediate",
			Instructions: [][]byte{{0x00, 0x7c, 0x00, 0xd1}}, // SUB X0, X0, #31
			InitialRegs:  map[int]uint64{0: 100},
			Description:  "SUB with immediate value",
			ExpectedRegs: map[int]uint64{0: 69},
		},
		// SUB - Register
		{
			Name:         "SUB_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xcb}}, // SUB X0, X0, X1
			InitialRegs:  map[int]uint64{0: 100, 1: 25},
			Description:  "SUB with register operand",
			ExpectedRegs: map[int]uint64{0: 75},
		},
		// SUBS - Immediate (flag setting)
		{
			Name:         "SUBS_immediate",
			Instructions: [][]byte{{0x00, 0x7c, 0x00, 0xf1}}, // SUBS X0, X0, #31
			InitialRegs:  map[int]uint64{0: 100},
			Description:  "SUBS with immediate - sets flags",
			ExpectedRegs: map[int]uint64{0: 69},
			CheckFlags:   true,
		},
		// SUBS - Register (flag setting)
		{
			Name:         "SUBS_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xeb}}, // SUBS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 100, 1: 25},
			Description:  "SUBS with register - sets flags",
			ExpectedRegs: map[int]uint64{0: 75},
			CheckFlags:   true,
		},
		// MADD - Multiply-add
		{
			Name:         "MADD",
			Instructions: [][]byte{{0x00, 0x04, 0x02, 0x9b}}, // MADD X0, X0, X2, X1
			InitialRegs:  map[int]uint64{0: 5, 1: 10, 2: 3},
			Description:  "MADD - multiply-add operation",
			ExpectedRegs: map[int]uint64{0: 25}, // 10 + (5 * 3) = 25
		},
		// MUL
		{
			Name:         "MUL",
			Instructions: [][]byte{{0x00, 0x7c, 0x01, 0x9b}}, // MUL X0, X0, X1
			InitialRegs:  map[int]uint64{0: 7, 1: 6},
			Description:  "MUL - multiply operation",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		// UDIV
		{
			Name:         "UDIV",
			Instructions: [][]byte{{0x00, 0x08, 0xc1, 0x9a}}, // UDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 100, 1: 4},
			Description:  "UDIV - unsigned divide",
			ExpectedRegs: map[int]uint64{0: 25},
		},
		// SDIV
		{
			Name:         "SDIV",
			Instructions: [][]byte{{0x00, 0x0c, 0xc1, 0x9a}}, // SDIV X0, X0, X1
			InitialRegs:  map[int]uint64{0: 100, 1: 4},
			Description:  "SDIV - signed divide",
			ExpectedRegs: map[int]uint64{0: 25},
		},
	}

	runInstructionTests(t, "Arithmetic", testCases)
}

// TestLogicalInstructions tests all logical instructions against hypervisor
func TestLogicalInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// AND - Immediate
		{
			Name:         "AND_immediate",
			Instructions: [][]byte{{0x00, 0x00, 0x40, 0x92}}, // AND X0, X0, #0xFF
			InitialRegs:  map[int]uint64{0: 0x12345678ABCDEF00},
			Description:  "AND with immediate mask",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
		},
		// AND - Register
		{
			Name:         "AND_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0x8a}}, // AND X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xAAAAAAAAAAAAAAAA, 1: 0x5555555555555555},
			Description:  "AND with register operand",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
		},
		// ANDS - Immediate (flag setting)
		{
			Name:         "ANDS_immediate",
			Instructions: [][]byte{{0x00, 0x00, 0x40, 0xf2}}, // ANDS X0, X0, #0xFF
			InitialRegs:  map[int]uint64{0: 0x12345678ABCDEF00},
			Description:  "ANDS with immediate - sets flags",
			ExpectedRegs: map[int]uint64{0: 0x0000000000000000},
			CheckFlags:   true,
		},
		// ANDS - Register (flag setting)
		{
			Name:         "ANDS_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xea}}, // ANDS X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x8000000000000000, 1: 0x8000000000000000},
			Description:  "ANDS with register - sets flags",
			ExpectedRegs: map[int]uint64{0: 0x8000000000000000},
			CheckFlags:   true,
		},
		// ORR - Immediate
		{
			Name:         "ORR_immediate",
			Instructions: [][]byte{{0x00, 0x1c, 0x40, 0xb2}}, // ORR X0, X0, #0xFF
			InitialRegs:  map[int]uint64{0: 0x1234567800000000},
			Description:  "ORR with immediate value",
			ExpectedRegs: map[int]uint64{0: 0x12345678000000FF},
		},
		// ORR - Register
		{
			Name:         "ORR_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xaa}}, // ORR X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0x1234000000000000, 1: 0x0000567800000000},
			Description:  "ORR with register operand",
			ExpectedRegs: map[int]uint64{0: 0x1234567800000000},
		},
		// EOR - Exclusive OR
		{
			Name:         "EOR_register",
			Instructions: [][]byte{{0x00, 0x00, 0x01, 0xca}}, // EOR X0, X0, X1
			InitialRegs:  map[int]uint64{0: 0xAAAAAAAAAAAAAAAA, 1: 0x5555555555555555},
			Description:  "EOR with register operand",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
		},
		// TST - Test bits (ANDS with XZR destination)
		{
			Name:         "TST",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xea}}, // TST X0, X1 (ANDS XZR, X0, X1)
			InitialRegs:  map[int]uint64{0: 0x8000000000000000, 1: 0x8000000000000000},
			Description:  "TST - test bits operation",
			CheckFlags:   true,
		},
	}

	runInstructionTests(t, "Logical", testCases)
}

// TestBitfieldInstructions tests all bitfield instructions against hypervisor
func TestBitfieldInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// BFM - Bitfield Move
		{
			Name:         "BFM",
			Instructions: [][]byte{{0x20, 0x1c, 0x40, 0xb3}}, // BFM X0, X1, #0, #7
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF, 1: 0x123456789ABCDEF0},
			Description:  "BFM - bitfield move operation",
		},
		// SBFM - Signed Bitfield Move
		{
			Name:         "SBFM",
			Instructions: [][]byte{{0x20, 0x1c, 0x44, 0x93}}, // SBFM X0, X1, #4, #7
			InitialRegs:  map[int]uint64{0: 0, 1: 0x80},
			Description:  "SBFM - signed bitfield move",
		},
		// UBFM - Unsigned Bitfield Move
		{
			Name:         "UBFM",
			Instructions: [][]byte{{0x20, 0x3c, 0x48, 0xd3}}, // UBFM X0, X1, #8, #15
			InitialRegs:  map[int]uint64{0: 0, 1: 0xFF00},
			Description:  "UBFM - unsigned bitfield move",
			ExpectedRegs: map[int]uint64{0: 0xFF},
		},
		// ASR - Arithmetic Shift Right (SBFM alias)
		{
			Name:         "ASR",
			Instructions: [][]byte{{0x20, 0xfc, 0x44, 0x93}}, // ASR X0, X1, #4
			InitialRegs:  map[int]uint64{0: 0, 1: 0x80},
			Description:  "ASR - arithmetic shift right",
		},
		// LSL - Logical Shift Left (UBFM alias)
		{
			Name:         "LSL",
			Instructions: [][]byte{{0x20, 0xec, 0x7c, 0xd3}}, // LSL X0, X1, #4
			InitialRegs:  map[int]uint64{0: 0, 1: 0xF},
			Description:  "LSL - logical shift left",
			ExpectedRegs: map[int]uint64{0: 0xF0},
		},
		// LSR - Logical Shift Right (UBFM alias)
		{
			Name:         "LSR",
			Instructions: [][]byte{{0x20, 0xfc, 0x44, 0xd3}}, // LSR X0, X1, #4
			InitialRegs:  map[int]uint64{0: 0, 1: 0xF0},
			Description:  "LSR - logical shift right",
			ExpectedRegs: map[int]uint64{0: 0xF},
		},
	}

	runInstructionTests(t, "Bitfield", testCases)
}

// TestMemoryInstructions tests all memory instructions against hypervisor
func TestMemoryInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// LDR - Load register (unsigned offset) - Combined STR+LDR test
		{
			Name: "LDR_unsigned_offset",
			Instructions: [][]byte{
				{0x01, 0x08, 0x00, 0xf9}, // STR X1, [X0, #16] - store test data
				{0x00, 0x08, 0x40, 0xf9}, // LDR X0, [X0, #16] - load it back
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x9ABCDEF012345678},
			Description:  "LDR with unsigned offset (via STR+LDR)",
			ExpectedRegs: map[int]uint64{0: 0x9ABCDEF012345678},
		},
		// STR - Store register
		{
			Name:         "STR",
			Instructions: [][]byte{{0x01, 0x00, 0x00, 0xf9}}, // STR X1, [X0]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0xDEADBEEFCAFEBABE},
			Description:  "STR - store register",
		},
		// LDRB - Load byte - Combined STRB+LDRB test
		{
			Name: "LDRB",
			Instructions: [][]byte{
				{0x01, 0x04, 0x00, 0x39}, // STRB W1, [X0, #1] - store test byte
				{0x00, 0x04, 0x40, 0x39}, // LDRB W0, [X0, #1] - load it back
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0xFF},
			Description:  "LDRB - load byte (via STRB+LDRB)",
			ExpectedRegs: map[int]uint64{0: 0xFF},
		},
		// LDRH - Load halfword - Combined STRH+LDRH test
		{
			Name: "LDRH",
			Instructions: [][]byte{
				{0x01, 0x04, 0x00, 0x79}, // STRH W1, [X0, #2] - store test halfword
				{0x00, 0x04, 0x40, 0x79}, // LDRH W0, [X0, #2] - load it back
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x1234},
			Description:  "LDRH - load halfword (via STRH+LDRH)",
			ExpectedRegs: map[int]uint64{0: 0x1234},
		},
		// LDRSB - Load signed byte - Combined STRB+LDRSB test
		{
			Name: "LDRSB",
			Instructions: [][]byte{
				{0x01, 0x04, 0x00, 0x39}, // STRB W1, [X0, #1] - store 0x80 (negative byte)
				{0x00, 0x04, 0x80, 0x39}, // LDRSB X0, [X0, #1] - load with sign extension
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x80},
			Description:  "LDRSB - load signed byte (via STRB+LDRSB)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFFF80},
		},
		// LDRSH - Load signed halfword - Combined STRH+LDRSH test
		{
			Name: "LDRSH",
			Instructions: [][]byte{
				{0x01, 0x04, 0x00, 0x79}, // STRH W1, [X0, #2] - store 0x8000 (negative halfword)
				{0x00, 0x04, 0x80, 0x79}, // LDRSH X0, [X0, #2] - load with sign extension
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x8000},
			Description:  "LDRSH - load signed halfword (via STRH+LDRSH)",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFF8000},
		},
		// LDUR - Load unscaled - Combined STUR+LDUR test
		{
			Name: "LDUR",
			Instructions: [][]byte{
				{0x01, 0x80, 0x00, 0xf8}, // STUR X1, [X0, #8] - store test data
				{0x00, 0x80, 0x40, 0xf8}, // LDUR X0, [X0, #8] - load it back
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x9ABCDEF012345678},
			Description:  "LDUR - load unscaled offset (via STUR+LDUR)",
			ExpectedRegs: map[int]uint64{0: 0x9ABCDEF012345678},
		},
		// STUR - Store unscaled
		{
			Name:         "STUR",
			Instructions: [][]byte{{0x01, 0x80, 0x00, 0xf8}}, // STUR X1, [X0, #8]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0},
			Description:  "STUR - store unscaled offset",
		},
		// LDP - Load pair - Combined STP+LDP test
		{
			Name: "LDP",
			Instructions: [][]byte{
				{0x01, 0x08, 0x01, 0xa9}, // STP X1, X2, [X0, #16] - store test data pair
				{0x00, 0x08, 0x41, 0xa9}, // LDP X0, X2, [X0, #16] - load them back (X0 gets X1's value, X2 gets X2's value)
			},
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x12345678, 2: 0x9ABCDEF0},
			Description:  "LDP - load pair (via STP+LDP)",
			ExpectedRegs: map[int]uint64{0: 0x12345678, 2: 0x9ABCDEF0},
		},
		// STP - Store pair
		{
			Name:         "STP",
			Instructions: [][]byte{{0x01, 0x08, 0x00, 0xa9}}, // STP X1, X2, [X0]
			InitialRegs:  map[int]uint64{0: 0x10000000, 1: 0x123456789ABCDEF0, 2: 0xDEADBEEFCAFEBABE},
			Description:  "STP - store pair",
		},
	}

	runInstructionTests(t, "Memory", testCases)
}

// TestMoveInstructions tests all move instructions against hypervisor
func TestMoveInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// MOV - Move register
		{
			Name:         "MOV_register",
			Instructions: [][]byte{{0xe0, 0x03, 0x01, 0xaa}}, // MOV X0, X1
			InitialRegs:  map[int]uint64{1: 0x123456789ABCDEF0},
			Description:  "MOV register to register",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABCDEF0},
		},
		// MOVZ - Move wide with zero
		{
			Name:         "MOVZ",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0xd2}}, // MOVZ X0, #0x2345
			InitialRegs:  map[int]uint64{0: 0xFFFFFFFFFFFFFFFF},
			Description:  "MOVZ - move wide with zero",
			ExpectedRegs: map[int]uint64{0: 0x2345},
		},
		// MOVN - Move wide with NOT
		{
			Name:         "MOVN",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0x92}}, // MOVN X0, #0x2345
			InitialRegs:  map[int]uint64{0: 0},
			Description:  "MOVN - move wide with NOT",
			ExpectedRegs: map[int]uint64{0: 0xFFFFFFFFFFFFDCBA},
		},
		// MOVK - Move wide with keep
		{
			Name:         "MOVK",
			Instructions: [][]byte{{0xa0, 0x68, 0x84, 0xf2}}, // MOVK X0, #0x2345
			InitialRegs:  map[int]uint64{0: 0x123456789ABCDEF0},
			Description:  "MOVK - move wide with keep",
			ExpectedRegs: map[int]uint64{0: 0x123456789ABC2345},
		},
	}

	runInstructionTests(t, "Move", testCases)
}

// TestConditionalInstructions tests conditional instructions against hypervisor
func TestConditionalInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// CSEL - Conditional select
		{
			Name: "CSEL_condition_true",
			Instructions: [][]byte{
				{0x3f, 0x04, 0x00, 0xf1}, // SUBS XZR, X1, #1 (set flags: 5-1=4, positive)
				{0x22, 0x10, 0x81, 0x9a}, // CSEL X2, X1, X1, NE (should select X1)
			},
			InitialRegs:  map[int]uint64{1: 5},
			Description:  "CSEL with true condition",
			ExpectedRegs: map[int]uint64{2: 5},
		},
		// CSINC - Conditional select and increment
		{
			Name: "CSINC_condition_true",
			Instructions: [][]byte{
				{0x1f, 0x00, 0x01, 0xeb}, // CMP X0, X1 (0 vs 0, sets Z flag)
				{0x22, 0x00, 0x81, 0x9a}, // CSINC X2, X0, X1, EQ (true, so X0)
			},
			InitialRegs:  map[int]uint64{0: 0, 1: 0},
			Description:  "CSINC with true condition",
			ExpectedRegs: map[int]uint64{2: 0}, // EQ condition true, so X0 = 0
		},
		// CMP - Compare (SUBS with XZR destination)
		{
			Name:         "CMP",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xeb}}, // CMP X0, X1
			InitialRegs:  map[int]uint64{0: 10, 1: 5},
			Description:  "CMP - compare operation",
			CheckFlags:   true,
		},
		// CMN - Compare negative (ADDS with XZR destination)
		{
			Name:         "CMN",
			Instructions: [][]byte{{0x1f, 0x00, 0x01, 0xab}}, // CMN X0, X1
			InitialRegs:  map[int]uint64{0: 5, 1: 10},
			Description:  "CMN - compare negative",
			CheckFlags:   true,
		},
		// CCMP - Conditional compare
		{
			Name: "CCMP",
			Instructions: [][]byte{
				{0x1f, 0x04, 0x00, 0xf1}, // SUBS XZR, X0, #1 (set condition)
				{0x2f, 0x08, 0x41, 0xfa}, // CCMP X1, #1, #15, EQ
			},
			InitialRegs: map[int]uint64{0: 1, 1: 1},
			Description: "CCMP - conditional compare",
			CheckFlags:  true,
		},
	}

	runInstructionTests(t, "Conditional", testCases)
}

// TestBranchInstructions tests branch instructions (limited - no actual branching)
func TestBranchInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// ADR - Address generation
		{
			Name:         "ADR",
			Instructions: [][]byte{{0x00, 0x00, 0x00, 0x10}}, // ADR X0, #0
			InitialRegs:  map[int]uint64{},
			Description:  "ADR - address generation",
			ExpectedRegs: map[int]uint64{0: 0x10000000}, // Should be current PC
		},
		// ADRP - Address generation (page)
		{
			Name:         "ADRP",
			Instructions: [][]byte{{0x00, 0x00, 0x00, 0x90}}, // ADRP X0, #0
			InitialRegs:  map[int]uint64{},
			Description:  "ADRP - address generation page",
			ExpectedRegs: map[int]uint64{0: 0x10000000}, // Page-aligned address
		},
		// Note: We can't test actual branch instructions (B, BL, etc.) easily
		// since they would change control flow and our test framework expects
		// linear execution. The integration tests cover branching behavior.
	}

	runInstructionTests(t, "Branch", testCases)
}

// TestSystemInstructions tests system instructions against hypervisor
func TestSystemInstructions(t *testing.T) {
	testCases := []InstructionTestCase{
		// NOP - No operation
		{
			Name:         "NOP",
			Instructions: [][]byte{{0x1f, 0x20, 0x03, 0xd5}}, // NOP
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "NOP - no operation",
			ExpectedRegs: map[int]uint64{0: 42}, // Should be unchanged
		},
		// HINT instructions (variations of NOP)
		{
			Name:         "YIELD",
			Instructions: [][]byte{{0x3f, 0x20, 0x03, 0xd5}}, // YIELD
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "YIELD - yield hint",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "WFE",
			Instructions: [][]byte{{0x5f, 0x20, 0x03, 0xd5}}, // WFE
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "WFE - wait for event",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "WFI",
			Instructions: [][]byte{{0x7f, 0x20, 0x03, 0xd5}}, // WFI
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "WFI - wait for interrupt",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "SEV",
			Instructions: [][]byte{{0x9f, 0x20, 0x03, 0xd5}}, // SEV
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "SEV - send event",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "SEVL",
			Instructions: [][]byte{{0xbf, 0x20, 0x03, 0xd5}}, // SEVL
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "SEVL - send event local",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		// Memory barriers
		{
			Name:         "DSB",
			Instructions: [][]byte{{0x9f, 0x3f, 0x03, 0xd5}}, // DSB SY
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "DSB - data synchronization barrier",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "DMB",
			Instructions: [][]byte{{0xbf, 0x3f, 0x03, 0xd5}}, // DMB SY
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "DMB - data memory barrier",
			ExpectedRegs: map[int]uint64{0: 42},
		},
		{
			Name:         "ISB",
			Instructions: [][]byte{{0xdf, 0x3f, 0x03, 0xd5}}, // ISB
			InitialRegs:  map[int]uint64{0: 42},
			Description:  "ISB - instruction synchronization barrier",
			ExpectedRegs: map[int]uint64{0: 42},
		},
	}

	runInstructionTests(t, "System", testCases)
}

// runInstructionTests executes a set of instruction test cases
func runInstructionTests(t *testing.T, category string, testCases []InstructionTestCase) {
	framework, err := NewHypervisorFramework()
	if err != nil {
		t.Skipf("Hypervisor framework not available for %s tests: %v", category, err)
		return
	}
	defer framework.Close()

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", category, tc.Name), func(t *testing.T) {
			// Reset framework to clean state before each test
			if err := framework.Reset(); err != nil {
				t.Skipf("Failed to reset framework: %v", err)
				return
			}

			// Set up initial state
			if err := framework.SetRegisterState(tc.InitialRegs); err != nil {
				t.Skipf("Failed to set register state: %v", err)
				return
			}

			if len(tc.InitialMem) > 0 {
				if err := framework.SetMemoryState(tc.InitialMem); err != nil {
					t.Skipf("Failed to set memory state: %v", err)
					return
				}
			}

			// Execute instruction and compare with Hypervisor
			// Use a different PC address to avoid conflicts with test memory
			result, err := framework.ExecuteSequenceAndCompare(tc.Instructions, 0x40000000)
			if err != nil {
				t.Errorf("Failed to execute instruction %s: %v", tc.Name, err)
				return
			}

			// Count critical differences (excluding PC which is expected to differ)
			criticalDiffs := 0
			for _, diff := range result.Differences {
				if diff.Name != "PC" {
					criticalDiffs++
					t.Logf("Difference in %s: %s %s: our=0x%x, hypervisor=0x%x",
						tc.Name, diff.Type, diff.Name, diff.OurValue, diff.HypervisorValue)
				}
			}

			// Fail test if there are critical differences
			if criticalDiffs > 0 {
				t.Errorf("Found %d critical differences in instruction %s", criticalDiffs, tc.Name)
			} else {
				t.Logf("✅ %s validated successfully: %s", tc.Name, tc.Description)
			}

			// Additional validation if expected registers are specified
			if len(tc.ExpectedRegs) > 0 {
				for reg, expectedVal := range tc.ExpectedRegs {
					if actualVal, exists := result.AfterState.Registers[reg]; exists {
						if actualVal != expectedVal {
							t.Errorf("Register X%d: expected 0x%x, got 0x%x", reg, expectedVal, actualVal)
						}
					}
				}
			}

			// Flag validation for instructions that should set flags
			if tc.CheckFlags {
				// Note: Flag validation would be more complex and would require
				// checking the specific flag bits in the CPSR register.
				// For now, we rely on the hypervisor comparison.
				t.Logf("Flag behavior validated via hypervisor comparison for %s", tc.Name)
			}
		})
	}
}
