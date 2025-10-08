package emulate

// ARM64 Instruction Detection and Decoding Helpers
// These functions provide convenience methods for quickly scanning ARM64 instruction patterns

// IsADRP checks if instruction is ADRP (Address of Page)
// ADRP encoding: 1001_0000_iiii_iiii_iiii_iiii_iiid_dddd (bits 31-24: 0x90)
func IsADRP(instr uint32) bool {
	return (instr & 0x9F000000) == 0x90000000
}

// IsBL checks if instruction is BL (Branch with Link)
// BL encoding: 1001_01ii_iiii_iiii_iiii_iiii_iiii_iiii (bits 31-26: 0x25)
func IsBL(instr uint32) bool {
	return (instr & 0xFC000000) == 0x94000000
}

// IsBLR checks if instruction is BLR (Branch with Link to Register)
// BLR encoding: 1101_0110_0011_1111_0000_00nn_nnn0_0000 (bits 31-10: 0xD63F0, bits 9-5: register)
func IsBLR(instr uint32) bool {
	return (instr & 0xFFFFFC1F) == 0xD63F0000
}

// IsSTR checks if instruction is STR (Store Register) with immediate offset
// STR (immediate) encoding: 1111_1001_00ii_iiii_iiii_iinn_nnnt_tttt (bits 31-22: 0x3E4)
func IsSTR(instr uint32) bool {
	// STR (immediate, post-index): 1111_1000_00ii_iiii_iiii_iinn_nnnt_tttt
	// STR (immediate, pre-index):  1111_1000_01ii_iiii_iiii_iinn_nnnt_tttt
	// STR (immediate, unsigned):   1111_1001_00ii_iiii_iiii_iinn_nnnt_tttt
	return ((instr & 0xFFC00000) == 0xF8000000) || // post/pre-index forms
		((instr & 0xFFC00000) == 0xF9000000) // unsigned immediate form
}

// IsLDR checks if instruction is LDR (Load Register) with immediate offset
// LDR (immediate) encoding: 1111_1001_01ii_iiii_iiii_iinn_nnnt_tttt (bits 31-22: 0x3E5)
func IsLDR(instr uint32) bool {
	// LDR (immediate, post-index): 1111_1000_01ii_iiii_iiii_iinn_nnnt_tttt
	// LDR (immediate, pre-index):  1111_1000_11ii_iiii_iiii_iinn_nnnt_tttt
	// LDR (immediate, unsigned):   1111_1001_01ii_iiii_iiii_iinn_nnnt_tttt
	return ((instr & 0xFFC00000) == 0xF8400000) || // post/pre-index forms
		((instr & 0xFFC00000) == 0xF9400000) // unsigned immediate form
}

// IsADD checks if instruction is ADD (Add)
// ADD (immediate) encoding: 1001_0001_xxii_iiii_iiii_iinn_nnnd_dddd (bits 31-24: 0x91)
// ADD (shifted register): 1000_1011_xx0m_mmmm_iiii_iinn_nnnd_dddd (bits 31-24: 0x8B)
func IsADD(instr uint32) bool {
	// ADD immediate: bits 31-24 == 0x91
	// ADD shifted register: bits 31-24 == 0x8B and bit 21 == 0
	return ((instr & 0xFF000000) == 0x91000000) ||
		((instr & 0xFF200000) == 0x8B000000)
}

// IsSTP checks if instruction is STP (Store Pair)
// STP encoding: 1010_1001_x0ii_iiii_iiii_iinn_nnnt_tttt (bit 22 = 0 for store)
func IsSTP(instr uint32) bool {
	// Check if it's a pair operation and bit 22 is 0 (store)
	return (instr&0xFE000000) == 0xA8000000 && (instr&0x00400000) == 0
}

// IsLDP checks if instruction is LDP (Load Pair)
// LDP encoding: 1010_1001_x1ii_iiii_iiii_iinn_nnnt_tttt (bit 22 = 1 for load)
func IsLDP(instr uint32) bool {
	// Check if it's a pair operation and bit 22 is 1 (load)
	return (instr&0xFE000000) == 0xA8000000 && (instr&0x00400000) != 0
}

// IsADR checks if instruction is ADR (Address)
// ADR encoding: 0001_0000_iiii_iiii_iiii_iiii_iiid_dddd (bits 31-24: 0x10)
func IsADR(instr uint32) bool {
	return (instr & 0x9F000000) == 0x10000000
}

// IsSTPFramePointer checks if STP stores x29,x30 (frame pointer setup)
// Frame pointer operations typically store FP (x29) and LR (x30) together
func IsSTPFramePointer(instr uint32) bool {
	if !IsSTP(instr) {
		return false
	}
	// Extract registers: bits 4-0 (Rt) and bits 14-10 (Rt2)
	rt := instr & 0x1F
	rt2 := (instr >> 10) & 0x1F
	// Check if storing x29 (FP) and x30 (LR)
	return (rt == 29 && rt2 == 30) || (rt == 30 && rt2 == 29)
}

// IsLDPFramePointer checks if LDP loads x29,x30 (frame pointer restore)
func IsLDPFramePointer(instr uint32) bool {
	if !IsLDP(instr) {
		return false
	}
	// Extract registers: bits 4-0 (Rt) and bits 14-10 (Rt2)
	rt := instr & 0x1F
	rt2 := (instr >> 10) & 0x1F
	// Check if loading x29 (FP) and x30 (LR)
	return (rt == 29 && rt2 == 30) || (rt == 30 && rt2 == 29)
}

// IsLinearInst checks if instruction doesn't change control flow
// This matches iometa's is_linear_inst() logic exactly - critical for function boundary detection
// Linear instructions include arithmetic, loads, stores, but exclude branches (except BL in stack frames)
func IsLinearInst(instr uint32) bool {
	// Branch instructions (control flow changes) - these are NOT linear
	if IsBL(instr) || IsBLR(instr) {
		return false
	}

	// B (unconditional branch): 0001_01ii_iiii_iiii_iiii_iiii_iiii_iiii
	if (instr & 0xFC000000) == 0x14000000 {
		return false
	}

	// B.cond (conditional branch): 0101_0100_iiii_iiii_iiii_iiii_iii0_cccc
	if (instr & 0xFF000010) == 0x54000000 {
		return false
	}

	// CBZ/CBNZ (compare and branch): xx11_0100_iiii_iiii_iiii_iiii_iiit_tttt
	if (instr & 0x7E000000) == 0x34000000 {
		return false
	}

	// TBZ/TBNZ (test bit and branch): x011_0110_iiii_iiii_iiii_iiii_iiit_tttt
	if (instr & 0x7E000000) == 0x36000000 {
		return false
	}

	// BR (branch to register): 1101_0110_0001_1111_0000_00nn_nnn0_0000
	if (instr & 0xFFFFFC1F) == 0xD61F0000 {
		return false
	}

	// RET (return): 1101_0110_0010_1111_0000_00nn_nnn0_0000
	if (instr & 0xFFFFFC1F) == 0xD65F0000 {
		return false
	}

	// All other instructions are linear (arithmetic, loads, stores, etc.)
	return true
}

// Instruction decoding helper functions

// ExtractBLTarget extracts the target address from BL instruction
// BL immediate: 26-bit signed immediate, shifted left by 2
func ExtractBLTarget(instr uint32, pc uint64) uint64 {
	if !IsBL(instr) {
		return 0
	}
	// Extract 26-bit immediate and sign extend
	imm26 := instr & 0x03FFFFFF
	// Sign extend from 26 bits to 32 bits
	var offset int32
	if (imm26 & 0x02000000) != 0 {
		// Negative: set upper bits to 1
		offset = int32(imm26 | 0xFC000000)
	} else {
		// Positive: upper bits already 0
		offset = int32(imm26)
	}
	// Shift left by 2 (multiply by 4) and add to PC
	return pc + uint64(offset<<2)
}

// ExtractADRPTarget extracts the target address from ADRP instruction
// ADRP: 21-bit immediate (page address), with lo:hi bit arrangement
func ExtractADRPTarget(instr uint32, pc uint64) uint64 {
	if !IsADRP(instr) {
		return 0
	}
	// Extract immhi (bits 23-5) and immlo (bits 30-29)
	immhi := (instr >> 5) & 0x7FFFF // 19 bits
	immlo := (instr >> 29) & 0x3    // 2 bits

	// Combine: immlo:immhi (21 bits total)
	imm21 := (immlo << 19) | immhi

	// Sign extend from 21 bits
	var offset int32
	if (imm21 & 0x100000) != 0 {
		// Negative
		offset = int32(imm21 | 0xFFE00000)
	} else {
		// Positive
		offset = int32(imm21)
	}

	// ADRP: Add to page-aligned PC (PC & ~0xFFF), shift by 12
	pagePC := pc & ^uint64(0xFFF)
	return pagePC + uint64(offset<<12)
}

// ExtractADRTarget extracts the target address from ADR instruction
// ADR: 21-bit immediate (byte address), with lo:hi bit arrangement
func ExtractADRTarget(instr uint32, pc uint64) uint64 {
	if !IsADR(instr) {
		return 0
	}
	// Extract immhi (bits 23-5) and immlo (bits 30-29)
	immhi := (instr >> 5) & 0x7FFFF // 19 bits
	immlo := (instr >> 29) & 0x3    // 2 bits

	// Combine: immlo:immhi (21 bits total)
	imm21 := (immlo << 19) | immhi

	// Sign extend from 21 bits
	var offset int32
	if (imm21 & 0x100000) != 0 {
		// Negative
		offset = int32(imm21 | 0xFFE00000)
	} else {
		// Positive
		offset = int32(imm21)
	}

	// ADR: Add to current PC
	return pc + uint64(offset)
}

// ExtractRegister extracts register number from instruction operand
// operandIndex: 0 for destination, 1 for first source, 2 for second source, etc.
func ExtractRegister(instr uint32, operandIndex int) uint8 {
	switch operandIndex {
	case 0: // Destination register (Rd) - bits 4-0
		return uint8(instr & 0x1F)
	case 1: // First source register (Rn) - bits 9-5
		return uint8((instr >> 5) & 0x1F)
	case 2: // Second source register (Rm) - bits 20-16
		return uint8((instr >> 16) & 0x1F)
	case 3: // Third register operand (Ra) - bits 14-10 (for some instructions like MADD)
		return uint8((instr >> 10) & 0x1F)
	default:
		return 0xFF // Invalid
	}
}

// ExtractImmediate extracts immediate value from instruction
// This handles common immediate encodings - specific instructions may need custom extraction
func ExtractImmediate(instr uint32) int64 {
	// For ADD/SUB immediate: 12-bit immediate in bits 21-10, with optional shift
	if IsADD(instr) && ((instr & 0xFF000000) == 0x91000000) {
		imm12 := (instr >> 10) & 0xFFF
		sh := (instr >> 22) & 0x1 // shift flag
		if sh != 0 {
			return int64(imm12 << 12) // shift by 12 if sh=1
		}
		return int64(imm12)
	}

	// For LDR/STR immediate: various immediate encodings depending on addressing mode
	if IsLDR(instr) || IsSTR(instr) {
		// Unsigned immediate (bits 21-10): imm12 * 8 for 64-bit loads/stores
		if ((instr & 0xFFC00000) == 0xF9000000) || ((instr & 0xFFC00000) == 0xF9400000) {
			imm12 := (instr >> 10) & 0xFFF
			return int64(imm12 * 8) // scale by 8 for 64-bit
		}
		// Signed immediate (bits 20-12): 9-bit signed for pre/post-index
		if ((instr & 0xFFC00000) == 0xF8000000) || ((instr & 0xFFC00000) == 0xF8400000) {
			imm9 := (instr >> 12) & 0x1FF
			if (imm9 & 0x100) != 0 {
				return int64(int32(imm9 | 0xFFFFFE00)) // sign extend
			}
			return int64(imm9)
		}
	}

	return 0 // Unknown or no immediate
}

// Utility functions for instruction analysis

// GetInstructionLength returns the length of an ARM64 instruction (always 4 bytes)
func GetInstructionLength(instr uint32) int {
	return 4
}

// IsValidInstruction performs basic validation of an ARM64 instruction
func IsValidInstruction(instr uint32) bool {
	// All 32-bit values are potentially valid ARM64 instructions
	// More sophisticated validation would require full decoding
	return true
}

// InstructionInfo provides detailed information about an instruction
type InstructionAnalysis struct {
	Instruction      uint32
	PC               uint64
	IsLinear         bool
	IsBranch         bool
	IsLoad           bool
	IsStore          bool
	IsPairOp         bool
	UsesFramePointer bool
	TargetAddress    uint64 // For branches
	Mnemonic         string
}

// AnalyzeInstruction provides comprehensive analysis of an ARM64 instruction
func AnalyzeInstruction(instr uint32, pc uint64) *InstructionAnalysis {
	analysis := &InstructionAnalysis{
		Instruction: instr,
		PC:          pc,
	}

	// Classify instruction type
	analysis.IsLinear = IsLinearInst(instr)
	analysis.IsBranch = IsBL(instr) || IsBLR(instr) ||
		((instr & 0xFC000000) == 0x14000000) || // B
		((instr & 0xFF000010) == 0x54000000) || // B.cond
		((instr & 0x7E000000) == 0x34000000) || // CBZ/CBNZ
		((instr & 0x7E000000) == 0x36000000) // TBZ/TBNZ

	analysis.IsLoad = IsLDR(instr) || IsLDP(instr)
	analysis.IsStore = IsSTR(instr) || IsSTP(instr)
	analysis.IsPairOp = IsLDP(instr) || IsSTP(instr)
	analysis.UsesFramePointer = IsLDPFramePointer(instr) || IsSTPFramePointer(instr)

	// Extract target address for branches
	if IsBL(instr) {
		analysis.TargetAddress = ExtractBLTarget(instr, pc)
		analysis.Mnemonic = "BL"
	} else if IsBLR(instr) {
		analysis.Mnemonic = "BLR"
	} else if IsADRP(instr) {
		analysis.TargetAddress = ExtractADRPTarget(instr, pc)
		analysis.Mnemonic = "ADRP"
	} else if IsADR(instr) {
		analysis.TargetAddress = ExtractADRTarget(instr, pc)
		analysis.Mnemonic = "ADR"
	}

	return analysis
}

// ScanInstructionPattern scans a byte buffer for specific ARM64 instruction patterns
// This is useful for quickly finding patterns in large binaries
func ScanInstructionPattern(data []byte, pattern func(uint32) bool) []uint64 {
	if len(data)%4 != 0 {
		return nil
	}

	var matches []uint64
	for i := 0; i < len(data); i += 4 {
		instr := uint32(data[i]) | uint32(data[i+1])<<8 |
			uint32(data[i+2])<<16 | uint32(data[i+3])<<24
		if pattern(instr) {
			matches = append(matches, uint64(i))
		}
	}
	return matches
}

// ScanBLInstructions finds all BL instructions in a byte buffer
func ScanBLInstructions(data []byte) []uint64 {
	return ScanInstructionPattern(data, IsBL)
}

// ScanADRPInstructions finds all ADRP instructions in a byte buffer
func ScanADRPInstructions(data []byte) []uint64 {
	return ScanInstructionPattern(data, IsADRP)
}

// ScanFramePointerOps finds all frame pointer operations (STP/LDP with x29,x30)
func ScanFramePointerOps(data []byte) []uint64 {
	return ScanInstructionPattern(data, func(instr uint32) bool {
		return IsSTPFramePointer(instr) || IsLDPFramePointer(instr)
	})
}
