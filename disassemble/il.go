package disassemble

import "math"

type LowLevelILFunction struct {
}

func ones32(i uint32) uint32 {
	return math.MaxUint32 >> (32 - i)
}

func ones64(i uint64) uint64 {
	return math.MaxUint64 >> (64 - i)
}

// func (o *Operand) ExtractImmediate(int sizeof_imm) {

// 	if o.Class != IMM32 && o.Class != IMM64 {
// 		return il.Unimplemented()
// 	}

// 	imm := o.Immediate

// 	if o.ShiftValueUsed {
// 		switch o.ShiftType {
// 		case SHIFT_TYPE_LSL:
// 			imm = imm << uint64(o.ShiftValue)
// 			break
// 		case SHIFT_TYPE_LSR:
// 			imm = imm >> uint64(o.ShiftValue)
// 			break
// 		case SHIFT_TYPE_MSL:
// 			imm = (imm << o.ShiftValue) | uint64(ones32(o.ShiftValue))
// 			break
// 		case SHIFT_TYPE_ASR:
// 			fallthrough
// 		case SHIFT_TYPE_ROR:
// 			fallthrough
// 		case SHIFT_TYPE_UXTW:
// 			fallthrough
// 		case SHIFT_TYPE_SXTW:
// 			fallthrough
// 		case SHIFT_TYPE_SXTX:
// 			fallthrough
// 		case SHIFT_TYPE_UXTX:
// 			fallthrough
// 		case SHIFT_TYPE_SXTB:
// 			fallthrough
// 		case SHIFT_TYPE_SXTH:
// 			fallthrough
// 		case SHIFT_TYPE_UXTH:
// 			fallthrough
// 		case SHIFT_TYPE_UXTB:
// 			fallthrough
// 		case SHIFT_TYPE_END:
// 			fallthrough
// 		default:
// 			return il.Unimplemented()
// 		}
// 	}

// 	return imm & ONES(sizeof_imm*8)
// }

// func (i *Instruction) GetLowLevelIL(Architecture* arch, uint64_t addr, LowLevelILFunction& il, Instruction& i, size_t addrSize) bool {
func (i *Instruction) GetLowLevelIL(il *LowLevelILFunction) bool {

	// int n_instrs_before = il.GetInstructionCount()

	//printf("%s() operation:%d encoding:%d\n", __func__, i.operation, i.encoding)

	// LowLevelILLabel trueLabel, falseLabel
	// switch (i.Operation){
	// case ARM64_ADD:
	// case ARM64_ADDS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Add(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])),
	// 					SETFLAGS)))
	// 	break
	// case ARM64_ADC:
	// case ARM64_ADCS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.AddCarry(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])),
	// 					il.Flag(IL_FLAG_C),
	// 					SETFLAGS)))
	// 	break
	// case ARM64_AND:
	// case ARM64_ANDS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.And(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])),
	// 					SETFLAGS)))
	// 	break
	// case ARM64_ADR:
	// case ARM64_ADRP:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0], il.ConstPointer(REGSZ_O(i.Operands[0]), IMM_O(i.Operands[1]))))
	// 	break
	// case ARM64_ASR:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0], il.ArithShiftRight(REGSZ_O(i.Operands[1]), ILREG_O(i.Operands[1]), ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[1])))))
	// 	break
	// case ARM64_AESD:
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))}, ARM64_INTRIN_AESD, {ILREG_O(i.Operands[0]), ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_AESE:
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))}, ARM64_INTRIN_AESE, {ILREG_O(i.Operands[0]), ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_BTI:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_HINT_BTI, {}))
	// 	break
	// case ARM64_B:
	// 	il.AddInstruction(DirectJump(arch, il, IMM_O(i.Operands[0]), addrSize))
	// 	break
	// case ARM64_B_NE:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_NE), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_EQ:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_E), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_CS:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_UGE), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_CC:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_ULT), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_MI:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_NEG), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_PL:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_POS), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_VS:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_O), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_VC:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_NO), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_HI:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_UGT), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_LS:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_ULE), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_GE:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_SGE), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_LT:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_SLT), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_GT:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_SGT), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_B_LE:
	// 	ConditionalJump(arch, il, il.FlagCondition(LLFC_SLE), addrSize, IMM_O(i.Operands[0]), addr + 4)
	// 	return false
	// case ARM64_BL:
	// 	il.AddInstruction(il.Call(il.ConstPointer(addrSize, IMM_O(i.Operands[0]))))
	// 	break
	// case ARM64_BLR:
	// case ARM64_BLRAA:
	// case ARM64_BLRAAZ:
	// case ARM64_BLRAB:
	// case ARM64_BLRABZ:
	// 	il.AddInstruction(il.Call(ILREG_O(i.Operands[0])))
	// 	break
	// case ARM64_BFC:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.And(REGSZ_O(i.Operands[0]),
	// 			il.Const(REGSZ_O(i.Operands[0]), ~(ONES(IMM_O(i.Operands[2])) << IMM_O(i.Operands[1]))),
	// 			ILREG_O(i.Operands[0]))))
	// 	break
	// case ARM64_BFI:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.Or(REGSZ_O(i.Operands[0]),
	// 			il.And(REGSZ_O(i.Operands[0]),
	// 				il.Const(REGSZ_O(i.Operands[0]), ~(ONES(IMM_O(i.Operands[3])) << IMM_O(i.Operands[2]))),
	// 				ILREG_O(i.Operands[0])),
	// 			il.ShiftLeft(REGSZ_O(i.Operands[0]),
	// 				il.And(REGSZ_O(i.Operands[0]),
	// 					il.Const(REGSZ_O(i.Operands[0]), ONES(IMM_O(i.Operands[3]))),
	// 					ILREG_O(i.Operands[1])),
	// 				il.Const(0, IMM_O(i.Operands[2]))))))
	// 	break
	// case ARM64_BFXIL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.Or(REGSZ_O(i.Operands[0]),
	// 			il.And(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				il.Const(REGSZ_O(i.Operands[0]), ~ONES(IMM_O(i.Operands[3])))),
	// 			il.LogicalShiftRight(REGSZ_O(i.Operands[0]),
	// 				il.And(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					il.Const(REGSZ_O(i.Operands[0]), ONES(IMM_O(i.Operands[3])) << IMM_O(i.Operands[2]))),
	// 				il.Const(0, IMM_O(i.Operands[2]))))))
	// 	break
	// case ARM64_BR:
	// case ARM64_BRAA:
	// case ARM64_BRAAZ:
	// case ARM64_BRAB:
	// case ARM64_BRABZ:
	// 	il.AddInstruction(il.Jump(ILREG_O(i.Operands[0])))
	// 	return false
	// case ARM64_BIC:
	// case ARM64_BICS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.And(REGSZ_O(i.Operands[1]),
	// 					ILREG_O(i.Operands[1]),
	// 					il.Not(REGSZ_O(i.Operands[1]),
	// 						ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[1]))), SETFLAGS)
	// 						))
	// 	break
	// case ARM64_CAS: // these compare-and-swaps can be 32 or 64 bit
	// case ARM64_CASA:
	// case ARM64_CASAL:
	// case ARM64_CASL:
	// 	GenIfElse(il,
	// 		il.CompareEqual(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[0]), il.Load(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[2]))),
	// 		il.Store(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[2]), ILREG_O(i.Operands[1])),
	// 		0
	// 	)
	// 	break
	// case ARM64_CASAH: // these compare-and-swaps are 16 bit
	// case ARM64_CASALH:
	// case ARM64_CASH:
	// case ARM64_CASLH:
	// 	GenIfElse(il,
	// 		il.CompareEqual(REGSZ_O(i.Operands[0]), ExtractRegister(il, i.Operands[0], 0, 2, false, 2), il.Load(2, ILREG_O(i.Operands[2]))),
	// 		il.Store(2, ILREG_O(i.Operands[2]), ExtractRegister(il, i.Operands[1], 0, 2, false, 2)),
	// 		0
	// 	)
	// 	break
	// case ARM64_CASAB: // these compare-and-swaps are 8 bit
	// case ARM64_CASALB:
	// case ARM64_CASB:
	// case ARM64_CASLB:
	// 	GenIfElse(il,
	// 		il.CompareEqual(REGSZ_O(i.Operands[0]), ExtractRegister(il, i.Operands[0], 0, 1, false, 1), il.Load(1, ILREG_O(i.Operands[2]))),
	// 		il.Store(1, ILREG_O(i.Operands[2]), ExtractRegister(il, i.Operands[1], 0, 1, false, 1)),
	// 		0
	// 	)
	// 	break
	// case ARM64_CBNZ:
	// 	ConditionalJump(arch, il,
	// 			il.CompareNotEqual(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				il.Const(REGSZ_O(i.Operands[0]), 0)),
	// 			addrSize, IMM_O(i.Operands[1]), addr + 4)
	// 	return false
	// case ARM64_CBZ:
	// 	ConditionalJump(arch, il,
	// 			il.CompareEqual(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				il.Const(REGSZ_O(i.Operands[0]), 0)),
	// 			addrSize, IMM_O(i.Operands[1]), addr + 4)
	// 	return false
	// case ARM64_CMN:
	// 	il.AddInstruction(il.Add(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 	break
	// case ARM64_CCMN:
	// 	{
	// 		LowLevelILLabel trueCode, falseCode, done

	// 		il.AddInstruction(il.If(GetCondition(il, i.Operands[3].cond), trueCode, falseCode))

	// 		il.MarkLabel(trueCode)
	// 		il.AddInstruction(il.Add(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[0]),
	// 					ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(falseCode)
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_N, il.Const(0, (IMM_O(i.Operands[2]) >> 3) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_Z, il.Const(0, (IMM_O(i.Operands[2]) >> 2) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_C, il.Const(0, (IMM_O(i.Operands[2]) >> 1) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_V, il.Const(0, (IMM_O(i.Operands[2]) >> 0) & 1)))

	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(done)
	// 	}
	// 	break
	// case ARM64_CMP:
	// 	il.AddInstruction(il.Sub(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 	break
	// case ARM64_CCMP:
	// 	{
	// 		LowLevelILLabel trueCode, falseCode, done

	// 		il.AddInstruction(il.If(GetCondition(il, i.Operands[3].cond), trueCode, falseCode))

	// 		il.MarkLabel(trueCode)
	// 		il.AddInstruction(il.Sub(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[0]),
	// 					ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(falseCode)
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_N, il.Const(0, (IMM_O(i.Operands[2]) >> 3) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_Z, il.Const(0, (IMM_O(i.Operands[2]) >> 2) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_C, il.Const(0, (IMM_O(i.Operands[2]) >> 1) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_V, il.Const(0, (IMM_O(i.Operands[2]) >> 0) & 1)))

	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(done)
	// 	}
	// 	break
	// case ARM64_CLREX:
	// 		il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_CLREX, {}))
	// 	break
	// case ARM64_CSEL:
	// case ARM64_FCSEL:
	// 	GenIfElse(il, GetCondition(il, i.Operands[3].cond),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[2])))
	// 	break
	// case ARM64_CSINC:
	// 	GenIfElse(il, GetCondition(il, i.Operands[3].cond),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])),
	// 		ILSETREG_O(i.Operands[0], ILADDREG_O(i.Operands[2], il.Const(REGSZ_O(i.Operands[0]), 1))))
	// 	break
	// case ARM64_CSINV:
	// 	GenIfElse(il, GetCondition(il, i.Operands[3].cond),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])),
	// 		ILSETREG_O(i.Operands[0], il.Not(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_CSNEG:
	// 	GenIfElse(il, GetCondition(il, i.Operands[3].cond),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])),
	// 		ILSETREG_O(i.Operands[0], il.Neg(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_CSET:
	// 	GenIfElse(il, GetCondition(il, i.Operands[1].cond),
	// 		ILSETREG_O(i.Operands[0], il.Const(REGSZ_O(i.Operands[0]), 1)),
	// 		ILSETREG_O(i.Operands[0], il.Const(REGSZ_O(i.Operands[0]), 0)))
	// 	break
	// case ARM64_CSETM:
	// 	GenIfElse(il, GetCondition(il, i.Operands[1].cond),
	// 		ILSETREG_O(i.Operands[0], il.Const(REGSZ_O(i.Operands[0]), -1)),
	// 		ILSETREG_O(i.Operands[0], il.Const(REGSZ_O(i.Operands[0]), 0)))
	// 	break
	// case ARM64_CINC:
	// 	GenIfElse(il, GetCondition(il, i.Operands[2].cond),
	// 		ILSETREG_O(i.Operands[0], ILADDREG_O(i.Operands[1], il.Const(REGSZ_O(i.Operands[0]), 1))),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])))
	// 	break
	// case ARM64_CINV:
	// 	GenIfElse(il, GetCondition(il, i.Operands[2].cond),
	// 		ILSETREG_O(i.Operands[0], il.Not(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]))),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])))
	// 	break
	// case ARM64_CNEG:
	// 	GenIfElse(il, GetCondition(il, i.Operands[2].cond),
	// 		ILSETREG_O(i.Operands[0], il.Neg(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]))),
	// 		ILSETREG_O(i.Operands[0], ILREG_O(i.Operands[1])))
	// 	break
	// case ARM64_CLZ:
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))}, ARM64_INTRIN_CLZ, {ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_DC:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_DC, {ILREG_O(i.Operands[1])})) /* i.Operands[0] is <dc_op> */
	// 	break
	// case ARM64_DMB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_DMB, {}))
	// 	break
	// case ARM64_DSB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_DSB, {}))
	// 	break
	// case ARM64_EON:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Xor(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					il.Not(REGSZ_O(i.Operands[0]), ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0]))))))
	// 	break
	// case ARM64_EOR:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Xor(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])))))
	// 	break
	// case ARM64_ESB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_ESB, {}))
	// 	break
	// case ARM64_EXTR:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.LogicalShiftRight(REGSZ_O(i.Operands[0]) * 2,
	// 					il.Or(REGSZ_O(i.Operands[0]) * 2,
	// 						il.ShiftLeft(REGSZ_O(i.Operands[0]) * 2, ILREG_O(i.Operands[1]), il.Const(1, REGSZ_O(i.Operands[0]) * 8)),
	// 						ILREG_O(i.Operands[2])),
	// 					il.Const(1, IMM_O(i.Operands[3])))))
	// 	break
	// case ARM64_FADD:
	// 	switch(i.encoding)
	// 	{
	// 		case ENC_FADD_H_FLOATDP2:
	// 		case ENC_FADD_S_FLOATDP2:
	// 		case ENC_FADD_D_FLOATDP2:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatAdd(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 			break
	// 		case ENC_FADD_ASIMDSAME_ONLY:
	// 		case ENC_FADD_ASIMDSAMEFP16_ONLY:
	// 		{
	// 			Register srcs[16], dsts[16]
	// 			int dst_n = unpack_vector(i.Operands[0], dsts)
	// 			int src_n = unpack_vector(i.Operands[1], srcs)
	// 			if((dst_n != src_n) || dst_n==0)
	// 				ABORT_LIFT

	// 			int rsize = get_register_size(dsts[0])
	// 			for(int i=0 i<dst_n ++i)
	// 				il.AddInstruction(il.FloatAdd(rsize, ILREG(dsts[i]), ILREG(srcs[i])))
	// 		}
	// 			break
	// 		default:
	// 			il.AddInstruction(il.Unimplemented())
	// 	}
	// 	break
	// case ARM64_FCCMP:
	// case ARM64_FCCMPE:
	// 	{
	// 		LowLevelILLabel trueCode, falseCode, done

	// 		il.AddInstruction(il.If(GetCondition(il, i.Operands[3].cond), trueCode, falseCode))

	// 		il.MarkLabel(trueCode)
	// 		il.AddInstruction(il.FloatSub(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[0]),
	// 					ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(falseCode)
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_N, il.Const(0, (IMM_O(i.Operands[2]) >> 3) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_Z, il.Const(0, (IMM_O(i.Operands[2]) >> 2) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_C, il.Const(0, (IMM_O(i.Operands[2]) >> 1) & 1)))
	// 		il.AddInstruction(il.SetFlag(IL_FLAG_V, il.Const(0, (IMM_O(i.Operands[2]) >> 0) & 1)))

	// 		il.AddInstruction(il.Goto(done))

	// 		il.MarkLabel(done)
	// 	}
	// 	break
	// case ARM64_FCMP:
	// case ARM64_FCMPE:
	// 	il.AddInstruction(il.FloatSub(REGSZ_O(i.Operands[0]),
	// 				ILREG_O(i.Operands[0]),
	// 				ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 	break
	// case ARM64_FSUB:
	// 	switch(i.encoding) {
	// 		case ENC_FSUB_H_FLOATDP2:
	// 		case ENC_FSUB_S_FLOATDP2:
	// 		case ENC_FSUB_D_FLOATDP2:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatSub(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 			break
	// 		case ENC_FSUB_ASIMDSAME_ONLY:
	// 		case ENC_FSUB_ASIMDSAMEFP16_ONLY:
	// 		{
	// 			Register srcs[16], dsts[16]
	// 			int dst_n = unpack_vector(i.Operands[0], dsts)
	// 			int src_n = unpack_vector(i.Operands[1], srcs)
	// 			if((dst_n != src_n) || dst_n==0)
	// 				ABORT_LIFT

	// 			int rsize = get_register_size(dsts[0])
	// 			for(int i=0 i<dst_n ++i)
	// 				il.AddInstruction(il.FloatSub(rsize, ILREG(dsts[i]), ILREG(srcs[i])))
	// 		}
	// 			break
	// 		default:
	// 			il.AddInstruction(il.Unimplemented())
	// 	}
	// 	break
	// case ARM64_FCVT:
	// {
	// 	int float_sz = 0
	// 	switch(i.encoding) {
	// 		/* non-SVE is straight register-to-register */
	// 		case ENC_FCVT_HS_FLOATDP1: // convert to half (2-byte)
	// 		case ENC_FCVT_HD_FLOATDP1:
	// 			float_sz = 2
	// 		case ENC_FCVT_SH_FLOATDP1: // convert to single (4-byte)
	// 		case ENC_FCVT_SD_FLOATDP1:
	// 			if(!float_sz) float_sz = 4
	// 		case ENC_FCVT_DH_FLOATDP1: // convert to double (8-byte)
	// 		case ENC_FCVT_DS_FLOATDP1:
	// 			if(!float_sz) float_sz = 8
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0], GetFloat(il, i.Operands[1], float_sz)))
	// 			break
	// 		/* future: support SVE versions with predicated execution and z register file */
	// 		default:
	// 			ABORT_LIFT
	// 	}
	// 	break
	// }
	// case ARM64_FDIV:
	// 	switch(i.encoding) {
	// 		case ENC_FDIV_H_FLOATDP2:
	// 		case ENC_FDIV_S_FLOATDP2:
	// 		case ENC_FDIV_D_FLOATDP2:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatDiv(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 			break
	// 		default:
	// 			il.AddInstruction(il.Unimplemented())
	// 	}
	// 	break
	// case ARM64_FMOV:
	// 	switch(i.encoding) {
	// 		case ENC_FMOV_64VX_FLOAT2INT:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatToInt(REGSZ_O(i.Operands[0]), ILREG(vector_reg_minimize(i.Operands[1])))))
	// 			break
	// 		case ENC_FMOV_V64I_FLOAT2INT:
	// 		{
	// 			Register minreg = vector_reg_minimize(i.Operands[0])
	// 			il.AddInstruction(il.SetRegister(get_register_size(minreg), minreg,
	// 				il.FloatToInt(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]))))
	// 			break
	// 		}
	// 		case ENC_FMOV_32H_FLOAT2INT:
	// 		case ENC_FMOV_32S_FLOAT2INT:
	// 		case ENC_FMOV_64H_FLOAT2INT:
	// 		case ENC_FMOV_64D_FLOAT2INT:
	// 		case ENC_FMOV_D64_FLOAT2INT:
	// 		case ENC_FMOV_H32_FLOAT2INT:
	// 		case ENC_FMOV_H64_FLOAT2INT:
	// 		case ENC_FMOV_S32_FLOAT2INT:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatToInt(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]))))
	// 			break
	// 		case ENC_FMOV_H_FLOATIMM:
	// 		case ENC_FMOV_S_FLOATIMM:
	// 		case ENC_FMOV_D_FLOATIMM:
	// 		{
	// 			int float_sz = 2
	// 			if(i.encoding==ENC_FMOV_S_FLOATIMM) float_sz = 4
	// 			if(i.encoding==ENC_FMOV_D_FLOATIMM) float_sz = 8
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				GetFloat(il, i.Operands[1], float_sz)))
	// 			break
	// 		}
	// 		case ENC_FMOV_H_FLOATDP1:
	// 		case ENC_FMOV_S_FLOATDP1:
	// 		case ENC_FMOV_D_FLOATDP1:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ILREG_O(i.Operands[1])))
	// 			break
	// 		case ENC_FMOV_ASIMDIMM_D2_D:
	// 		case ENC_FMOV_ASIMDIMM_H_H:
	// 		case ENC_FMOV_ASIMDIMM_S_S:
	// 		{
	// 			int float_sz = 2
	// 			if(i.encoding==ENC_FMOV_ASIMDIMM_S_S) float_sz = 4
	// 			if(i.encoding==ENC_FMOV_ASIMDIMM_D2_D) float_sz = 8

	// 			Register regs[16]
	// 			int dst_n = unpack_vector(i.Operands[0], regs)
	// 			for(int i=0 i<dst_n ++i)
	// 				il.AddInstruction(ILSETREG(regs[i], GetFloat(il, i.Operands[1], float_sz)))
	// 			break
	// 		}
	// 		default:
	// 			il.AddInstruction(il.Unimplemented())
	// 	}
	// 	break
	// case ARM64_FMUL:
	// 	switch(i.encoding) {
	// 		case ENC_FMUL_H_FLOATDP2:
	// 		case ENC_FMUL_S_FLOATDP2:
	// 		case ENC_FMUL_D_FLOATDP2:
	// 			il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.FloatMult(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 			break
	// 		default:
	// 			il.AddInstruction(il.Unimplemented())
	// 	}
	// 	break
	// case ARM64_ERET:
	// case ARM64_ERETAA:
	// case ARM64_ERETAB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_ERET, {}))
	// 	il.AddInstruction(il.Trap(0))
	// 	return false
	// case ARM64_ISB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_ISB, {}))
	// 	break
	// case ARM64_LDAR:
	// case ARM64_LDAXR:
	// 	LoadStoreOperand(il, true, i.Operands[0], i.Operands[1], 0)
	// 	break
	// case ARM64_LDARB:
	// case ARM64_LDAXRB:
	// 	LoadStoreOperandSize(il, true, false, 1, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDARH:
	// case ARM64_LDAXRH:
	// 	LoadStoreOperandSize(il, true, false, 2, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDP:
	// case ARM64_LDNP:
	// 	LoadStoreOperandPair(il, true, i.Operands[0], i.Operands[1], i.Operands[2])
	// 	break
	// case ARM64_LDR:
	// case ARM64_LDUR:
	// case ARM64_LDRAA:
	// case ARM64_LDRAB:
	// 	LoadStoreOperand(il, true, i.Operands[0], i.Operands[1], 0)
	// 	break
	// case ARM64_LDRB:
	// case ARM64_LDURB:
	// 	LoadStoreOperandSize(il, true, false, 1, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDRH:
	// case ARM64_LDURH:
	// 	LoadStoreOperandSize(il, true, false, 2, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDRSB:
	// case ARM64_LDURSB:
	// 	LoadStoreOperandSize(il, true, true, 1, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDRSH:
	// case ARM64_LDURSH:
	// 	LoadStoreOperandSize(il, true, true, 2, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDRSW:
	// case ARM64_LDURSW:
	// 	LoadStoreOperandSize(il, true, true, 4, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LD1:
	// 	LoadStoreVector(il, true, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_LDADD:
	// case ARM64_LDADDA:
	// case ARM64_LDADDL:
	// case ARM64_LDADDAL:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 0)
	// 	il.AddInstruction(il.Store(REGSZ_O(i.Operands[2]), ILREG_O(i.Operands[2]),
	// 		il.Add(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[0]), ILREG_O(i.Operands[1]))))
	// 	break
	// case ARM64_LDADDB:
	// case ARM64_LDADDAB:
	// case ARM64_LDADDLB:
	// case ARM64_LDADDALB:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 1)
	// 	il.AddInstruction(il.Store(REGSZ_O(i.Operands[2]), ILREG_O(i.Operands[2]),
	// 		il.Add(1, il.LowPart(1,ILREG_O(i.Operands[0])), il.LowPart(1,ILREG_O(i.Operands[1])))))
	// 	break
	// case ARM64_LDADDH:
	// case ARM64_LDADDAH:
	// case ARM64_LDADDLH:
	// case ARM64_LDADDALH:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 2)
	// 	il.AddInstruction(il.Store(REGSZ_O(i.Operands[2]), ILREG_O(i.Operands[2]),
	// 		il.Add(2, il.LowPart(2,ILREG_O(i.Operands[0])), il.LowPart(2,ILREG_O(i.Operands[1])))))
	// 	break
	// case ARM64_LSL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.ShiftLeft(REGSZ_O(i.Operands[1]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[1])))))
	// 	break
	// case ARM64_LSR:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.LogicalShiftRight(REGSZ_O(i.Operands[1]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[1])))))
	// 	break
	// case ARM64_MOV:
	// {
	// 	Register regs[16]
	// 	int n = unpack_vector(i.Operands[0], regs)

	// 	if (n == 1)
	// 		il.AddInstruction(ILSETREG(regs[0], ReadILOperand(il, i.Operands[1], get_register_size(regs[0]))))
	// 	else
	// 		ABORT_LIFT

	// 	break
	// }
	// case ARM64_MOVI:
	// {
	// 	Register regs[16]
	// 	int n = unpack_vector(i.Operands[0], regs)
	// 	for(int i=0 i<n ++i)
	// 		il.AddInstruction(ILSETREG(regs[i], ILCONST_O(get_register_size(regs[i]), i.Operands[1])))
	// 	break
	// }
	// case ARM64_MVN:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 					il.Not(REGSZ_O(i.Operands[0]), ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])))))
	// 	break
	// case ARM64_MOVK:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Or(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[0]), il.Const(REGSZ_O(i.Operands[0]), IMM_O(i.Operands[1]) << i.Operands[1].shiftValue))))
	// 	break
	// case ARM64_MOVZ:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Const(REGSZ_O(i.Operands[0]), IMM_O(i.Operands[1]) << i.Operands[1].shiftValue)))
	// 	break
	// case ARM64_MUL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Mult(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_MADD:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ILADDREG_O(i.Operands[3], il.Mult(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_MRS:
	// 	{
	// 		ExprId reg = ILREG_O(i.Operands[1])
	// 		const char *name = get_system_register_name((SystemReg)i.Operands[1].sysreg)

	// 		if (strlen(name) == 0) {
	// 			LogWarn("Unknown system register %d @ 0x%" PRIx64 ": S%d_%d_c%d_c%d_%d, using catch-all system register instead\n",
	// 					i.Operands[1].sysreg, addr, i.Operands[1].implspec[0], i.Operands[1].implspec[1], i.Operands[1].implspec[2],
	// 					i.Operands[1].implspec[3], i.Operands[1].implspec[4])
	// 			reg = il.Register(8, SYSREG_UNKNOWN)
	// 		}

	// 		il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))},
	// 					ARM64_INTRIN_MRS,
	// 					{reg}))
	// 		break
	// 	}
	// case ARM64_MSUB:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[3]), il.Mult(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_MNEG:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]),
	// 					il.Const(8, 0),
	// 					il.Mult(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_MSR:
	// 	{
	// 		uint32_t dst = i.Operands[0].sysreg
	// 		const char *name = get_system_register_name((SystemReg)dst)

	// 		if (strlen(name) == 0) {
	// 			LogWarn("Unknown system register %d @ 0x%" PRIx64 ": S%d_%d_c%d_c%d_%d, using catch-all system register instead\n",
	// 					dst, addr, i.Operands[0].implspec[0], i.Operands[0].implspec[1], i.Operands[0].implspec[2],
	// 					i.Operands[0].implspec[3], i.Operands[0].implspec[4])
	// 			dst = SYSREG_UNKNOWN
	// 		}

	// 		switch (i.Operands[1].operandClass) {
	// 			case IMM32:
	// 				il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(dst)},
	// 							ARM64_INTRIN_MSR,
	// 							{il.Const(4, IMM_O(i.Operands[1]))}))
	// 				break
	// 			case REG:
	// 				il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(dst)},
	// 							ARM64_INTRIN_MSR,
	// 							{ILREG_O(i.Operands[1])}))
	// 				break
	// 			default:
	// 				LogError("unknown MSR operand class: %x\n", i.Operands[1].operandClass)
	// 				break
	// 		}
	// 		break
	// 	}
	// case ARM64_NEG:
	// case ARM64_NEGS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Neg(REGSZ_O(i.Operands[0]), ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS)))
	// 	break
	// case ARM64_NOP:
	// 	il.AddInstruction(il.Nop())
	// 	break

	// case ARM64_AUTDA:
	// case ARM64_AUTDB:
	// case ARM64_AUTIA:
	// case ARM64_AUTIB:
	// case ARM64_PACDA:
	// case ARM64_PACDB:
	// case ARM64_PACIA:
	// case ARM64_PACIB:
	// 	// <Xd> is address, <Xn> is modifier
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))},
	// 				operation_to_intrinsic(i.operation),
	// 				{ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_PACGA:
	// 	// <Xd> is address, <Xn>, <Xm> are modifiers, keys
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))},
	// 				operation_to_intrinsic(i.operation),
	// 				{ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])}))
	// 	break
	// case ARM64_AUTIB1716:
	// case ARM64_PACIA1716:
	// case ARM64_PACIB1716:
	// 	// x17 is address, x16 is modifier
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_X17)},
	// 				operation_to_intrinsic(i.operation),
	// 				{il.Register(8, REG_X16)}))
	// 	break
	// case ARM64_AUTDZA:
	// case ARM64_AUTDZB:
	// case ARM64_AUTIZA:
	// case ARM64_AUTIZB:
	// case ARM64_PACDZA:
	// case ARM64_PACDZB:
	// case ARM64_PACIZA:
	// case ARM64_PACIZB:
	// case ARM64_XPACI:
	// case ARM64_XPACD:
	// 	// <Xd> is address, modifier is omitted or 0
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))},
	// 				operation_to_intrinsic(i.operation),
	// 				{}))
	// 	break
	// case ARM64_AUTIBZ:
	// case ARM64_PACIAZ:
	// case ARM64_PACIBZ:
	// case ARM64_XPACLRI:
	// 	// x30 is address, modifier is omitted or 0
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_X30)},
	// 				operation_to_intrinsic(i.operation),
	// 				{}))
	// 	break
	// case ARM64_AUTIBSP:
	// case ARM64_PACIASP:
	// case ARM64_PACIBSP:
	// 	// x30 is address, sp is modifier
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_X30)},
	// 				operation_to_intrinsic(i.operation),
	// 				{il.Register(8, REG_SP)}))
	// 	break
	// case ARM64_PRFUM:
	// case ARM64_PRFM:
	// 	// TODO use the PRFM types when we have a better option than defining 18 different intrinsics to account for:
	// 	// - 3 types {PLD, PLI, PST}
	// 	// - 3 targets {L1, L2, L3}
	// 	// - 2 policies {KEEP, STM}
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_PRFM, {ReadILOperand(il, i.Operands[1], 8)}))
	// 	break
	// case ARM64_ORN:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Or(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					il.Not(REGSZ_O(i.Operands[0]), ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0]))))))
	// 	break
	// case ARM64_ORR:
	// case ARM64_ORRS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Or(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])), SETFLAGS)))
	// 	break
	// case ARM64_PSB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_PSBCSYNC, {}))
	// 	break
	// case ARM64_RET:
	// case ARM64_RETAA:
	// case ARM64_RETAB:
	// {
	// 	ExprId reg = (i.Operands[0].operandClass == REG) ? ILREG_O(i.Operands[0]) : il.Register(8, REG_X30)
	// 	il.AddInstruction(il.Return(reg))
	// 	break
	// }
	// case ARM64_REVB: // SVE only
	// case ARM64_REVH:
	// case ARM64_REVW:
	// 	il.AddInstruction(il.Unimplemented())
	// 	break
	// case ARM64_REV16:
	// case ARM64_REV32:
	// case ARM64_REV64:
	// case ARM64_REV:
	// 	if(IS_SVE_O(i.Operands[0])) {
	// 		il.AddInstruction(il.Unimplemented())
	// 		break
	// 	}
	// 	// if LLIL_BSWAP ever gets added, replace
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))}, ARM64_INTRIN_REV, {ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_RBIT:
	// 	il.AddInstruction(il.Intrinsic({RegisterOrFlag::Register(REG_O(i.Operands[0]))}, ARM64_INTRIN_RBIT, {ILREG_O(i.Operands[1])}))
	// 	break
	// case ARM64_ROR:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.RotateRight(REGSZ_O(i.Operands[1]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[1])))))
	// 	break
	// case ARM64_SBC:
	// case ARM64_SBCS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.SubBorrow(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[1]),
	// 					ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])),
	// 					il.Not(0, il.Flag(IL_FLAG_C)),
	// 					SETFLAGS)))
	// 	break
	// case ARM64_SBFIZ:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 					il.ArithShiftRight(REGSZ_O(i.Operands[0]),
	// 						il.ShiftLeft(REGSZ_O(i.Operands[0]),
	// 							ExtractBits(il, i.Operands[1], IMM_O(i.Operands[3]), 0),
	// 							il.Const(1, (REGSZ_O(i.Operands[0])*8)-IMM_O(i.Operands[3]))),
	// 						il.Const(1, (REGSZ_O(i.Operands[0])*8)-IMM_O(i.Operands[2])-IMM_O(i.Operands[3])))))
	// 	break
	// case ARM64_SBFX:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 					il.ArithShiftRight(REGSZ_O(i.Operands[0]),
	// 						il.ShiftLeft(REGSZ_O(i.Operands[0]),
	// 							ExtractBits(il, i.Operands[1], IMM_O(i.Operands[3]), IMM_O(i.Operands[2])),
	// 							il.Const(1, (REGSZ_O(i.Operands[0])*8)-IMM_O(i.Operands[3])-IMM_O(i.Operands[2]))),
	// 						il.Const(1, (REGSZ_O(i.Operands[0])*8)-IMM_O(i.Operands[3])))))
	// 	break
	// case ARM64_SDIV:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.DivSigned(REGSZ_O(i.Operands[1]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_SEV:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_SEV, {}))
	// 	break
	// case ARM64_SEVL:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_SEVL, {}))
	// 	break
	// case ARM64_SHL:
	// {
	// 	Register srcs[16], dsts[16]
	// 	int dst_n = unpack_vector(i.Operands[0], dsts)
	// 	int src_n = unpack_vector(i.Operands[1], srcs)

	// 	if((dst_n != src_n) || dst_n==0)
	// 		ABORT_LIFT

	// 	int rsize = get_register_size(dsts[0])
	// 	for(int i=0 i<dst_n ++i) {
	// 		il.AddInstruction(il.SetRegister(rsize, dsts[i],
	// 			il.ShiftLeft(rsize, il.Register(rsize, srcs[i]), il.Const(0, IMM_O(i.Operands[2])))))
	// 	}

	// 	break
	// }
	// case ARM64_ST1:
	// 	LoadStoreVector(il, false, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_STP:
	// case ARM64_STNP:
	// 	LoadStoreOperandPair(il, false, i.Operands[0], i.Operands[1], i.Operands[2])
	// 	break
	// case ARM64_STR:
	// case ARM64_STLR:
	// case ARM64_STUR:
	// 	LoadStoreOperand(il, false, i.Operands[0], i.Operands[1], 0)
	// 	break
	// case ARM64_STRB:
	// case ARM64_STLRB:
	// case ARM64_STURB:
	// 	LoadStoreOperandSize(il, false, false, 1, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_STRH:
	// case ARM64_STLRH:
	// case ARM64_STURH:
	// 	LoadStoreOperandSize(il, false, false, 2, i.Operands[0], i.Operands[1])
	// 	break
	// case ARM64_SUB:
	// case ARM64_SUBS:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ReadILOperand(il, i.Operands[2], REGSZ_O(i.Operands[0])),
	// 				SETFLAGS)))
	// 	break
	// case ARM64_SVC:
	// case ARM64_HVC:
	// case ARM64_SMC:
	// {
	// 	/* b31,b30==xx of fake register mark transition to ELxx */
	// 	uint32_t el_mark = 0
	// 	if(i.operation == ARM64_SVC) el_mark = 0x40000000
	// 	else if(i.operation == ARM64_HVC) el_mark = 0x80000000
	// 	else if(i.operation == ARM64_SMC) el_mark = 0xC0000000
	// 	/* b15..b0 of fake register still holds syscall number */
	// 	il.AddInstruction(il.SetRegister(4, FAKEREG_SYSCALL_INFO, il.Const(4, el_mark | IMM_O(i.Operands[0]))))
	// 	il.AddInstruction(il.SystemCall())
	// 	break
	// }
	// case ARM64_SWP: /* word (4) or doubleword (8) */
	// case ARM64_SWPA:
	// case ARM64_SWPL:
	// case ARM64_SWPAL:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 0)
	// 	LoadStoreOperand(il, false, i.Operands[0], i.Operands[2], 0)
	// 	break
	// case ARM64_SWPB: /* byte (1) */
	// case ARM64_SWPAB:
	// case ARM64_SWPLB:
	// case ARM64_SWPALB:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 1)
	// 	il.AddInstruction(il.Store(1, ILREG_O(i.Operands[2]), il.LowPart(1, ILREG_O(i.Operands[0]))))
	// 	break
	// case ARM64_SWPH: /* half-word (2) */
	// case ARM64_SWPAH:
	// case ARM64_SWPLH:
	// case ARM64_SWPALH:
	// 	LoadStoreOperand(il, true, i.Operands[1], i.Operands[2], 2)
	// 	il.AddInstruction(il.Store(2, ILREG_O(i.Operands[2]), il.LowPart(2, ILREG_O(i.Operands[0]))))
	// 	break
	// case ARM64_SXTB:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ExtractRegister(il, i.Operands[1], 0, 1, true, REGSZ_O(i.Operands[0]))))
	// 	break
	// case ARM64_SXTH:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ExtractRegister(il, i.Operands[1], 0, 2, true, REGSZ_O(i.Operands[0]))))
	// 	break
	// case ARM64_SXTW:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ExtractRegister(il, i.Operands[1], 0, 4, true, REGSZ_O(i.Operands[0]))))
	// 	break
	// case ARM64_TBNZ:
	// 	ConditionalJump(arch, il,
	// 		il.CompareNotEqual(REGSZ_O(i.Operands[0]),
	// 			ExtractBit(il, i.Operands[0], IMM_O(i.Operands[1])),
	// 			il.Const(REGSZ_O(i.Operands[0]), 0)),
	// 		addrSize, IMM_O(i.Operands[2]), addr + 4)
	// 	return false
	// case ARM64_TBZ:
	// 	ConditionalJump(arch, il,
	// 		il.CompareEqual(REGSZ_O(i.Operands[0]),
	// 			ExtractBit(il, i.Operands[0], IMM_O(i.Operands[1])),
	// 			il.Const(REGSZ_O(i.Operands[0]), 0)),
	// 		addrSize, IMM_O(i.Operands[2]), addr + 4)
	// 	return false
	// case ARM64_TST:
	// 	il.AddInstruction(il.And(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[0]),
	// 					ReadILOperand(il, i.Operands[1], REGSZ_O(i.Operands[0])), SETFLAGS))
	// 	break
	// case ARM64_UMADDL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Add(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[3]),
	// 					il.MultDoublePrecUnsigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_UMULL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.MultDoublePrecUnsigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_UMSUBL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[3]),
	// 					il.MultDoublePrecUnsigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_UMNEGL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]),
	// 					il.Const(8, 0),
	// 					il.MultDoublePrecUnsigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_UXTL:
	// case ARM64_UXTL2:
	// {
	// 	Register srcs[16], dsts[16]
	// 	int dst_n = unpack_vector(i.Operands[0], dsts)
	// 	int src_n = unpack_vector(i.Operands[1], srcs)

	// 	if(src_n==0 || dst_n==0) ABORT_LIFT
	// 	if(i.operation==ARM64_UXTL && (src_n != dst_n)) ABORT_LIFT
	// 	if(i.operation==ARM64_UXTL2 && (src_n != 2*dst_n)) ABORT_LIFT

	// 	for(int i=0 i<dst_n ++i) {
	// 		if(i.operation==ARM64_UXTL)
	// 			il.AddInstruction(ILSETREG(dsts[i], ILREG(srcs[i])))
	// 		else
	// 			il.AddInstruction(ILSETREG(dsts[i], ILREG(srcs[i + src_n/2])))
	// 	}

	// 	break
	// }
	// case ARM64_SMADDL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Add(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[3]),
	// 					il.MultDoublePrecSigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_USHR:
	// {
	// 	Register srcs[16], dsts[16]
	// 	int dst_n = unpack_vector(i.Operands[0], dsts)
	// 	int src_n = unpack_vector(i.Operands[1], srcs)

	// 	if((dst_n != src_n) || dst_n==0)
	// 		ABORT_LIFT

	// 	int rsize = get_register_size(dsts[0])
	// 	for(int i=0 i<dst_n ++i) {
	// 		il.AddInstruction(il.SetRegister(rsize, dsts[i],
	// 			il.LogicalShiftRight(rsize, il.Register(rsize, srcs[i]), il.Const(0, IMM_O(i.Operands[2])))))
	// 	}

	// 	break
	// }
	// case ARM64_SMULL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.MultDoublePrecSigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_SMSUBL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]),
	// 					ILREG_O(i.Operands[3]),
	// 					il.MultDoublePrecSigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_SMNEGL:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.Sub(REGSZ_O(i.Operands[0]),
	// 					il.Const(8, 0),
	// 					il.MultDoublePrecSigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])))))
	// 	break
	// case ARM64_UMULH:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.LogicalShiftRight(16,
	// 			il.MultDoublePrecUnsigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])),
	// 			il.Const(1, 8))))
	// 	break
	// case ARM64_SMULH:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.LogicalShiftRight(16,
	// 			il.MultDoublePrecSigned(REGSZ_O(i.Operands[0]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2])),
	// 			il.Const(1, 8))))
	// 	break
	// case ARM64_UDIV:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				il.DivUnsigned(REGSZ_O(i.Operands[1]), ILREG_O(i.Operands[1]), ILREG_O(i.Operands[2]))))
	// 	break
	// case ARM64_UBFIZ:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 		il.ZeroExtend(REGSZ_O(i.Operands[0]), il.ShiftLeft(REGSZ_O(i.Operands[1]),
	// 			il.And(REGSZ_O(i.Operands[1]), ILREG_O(i.Operands[1]), il.Const(REGSZ_O(i.Operands[1]), (1LL << IMM_O(i.Operands[3])) - 1)),
	// 				il.Const(1, IMM_O(i.Operands[2]))))))
	// 	break
	// case ARM64_UBFX:
	// {
	// 	// ubfx <dst>, <src>, <src_lsb>, <src_len>
	// 	int src_lsb = IMM_O(i.Operands[2])
	// 	int src_len = IMM_O(i.Operands[3])
	// 	if(src_lsb==0 && (src_len==8 || src_len==16 || src_len==32 || src_len==64)) {
	// 		il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 			il.LowPart(src_len/8, ILREG_O(i.Operands[1]))))
	// 	}
	// 	else {
	// 		il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 			il.ZeroExtend(REGSZ_O(i.Operands[0]), il.And(REGSZ_O(i.Operands[1]),
	// 				il.LogicalShiftRight(REGSZ_O(i.Operands[1]), ILREG_O(i.Operands[1]), il.Const(1, IMM_O(i.Operands[2]))),
	// 					il.Const(REGSZ_O(i.Operands[1]), (1LL << IMM_O(i.Operands[3])) - 1)))))
	// 	}
	// 	break
	// }
	// case ARM64_UXTB:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ExtractRegister(il, i.Operands[1], 0, 1, false, REGSZ_O(i.Operands[0]))))
	// 	break
	// case ARM64_UXTH:
	// 	il.AddInstruction(ILSETREG_O(i.Operands[0],
	// 				ExtractRegister(il, i.Operands[1], 0, 2, false, REGSZ_O(i.Operands[0]))))
	// 	break
	// case ARM64_WFE:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_WFE, {}))
	// 	break
	// case ARM64_WFI:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_WFI, {}))
	// 	break
	// case ARM64_BRK:
	// 	il.AddInstruction(il.Trap(IMM_O(i.Operands[0]))) // FIXME Breakpoint may need a parameter (IMM_O(i.Operands[0])))
	// 	return false
	// case ARM64_DUP:
	// 	{
	// 		if(i.encoding != ENC_DUP_ASIMDINS_DR_R) ABORT_LIFT
	// 		Register regs[16]
	// 		int regs_n = unpack_vector(i.Operands[0], regs)
	// 		if(regs_n <= 0) ABORT_LIFT
	// 		int lane_sz = REGSZ(regs[0])
	// 		for(int i=0 i<regs_n ++i)
	// 			il.AddInstruction(ILSETREG(regs[i], ExtractRegister(il, i.Operands[1], 0, lane_sz, 0, lane_sz)))
	// 	}
	// 	break
	// case ARM64_DGH:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_HINT_DGH, {}))
	// 	break
	// case ARM64_TSB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_HINT_TSB, {}))
	// 	break
	// case ARM64_CSDB:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_HINT_CSDB, {}))
	// 	break
	// case ARM64_HINT:
	// 	if ((IMM_O(i.Operands[0]) & ~0b110) == 0b100000)
	// 		il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_HINT_BTI, {}))
	// 	else
	// 		LogWarn("unknown hint operand: 0x%" PRIx64 "\n", IMM_O(i.Operands[0]))
	// 	break
	// case ARM64_HLT:
	// 	il.AddInstruction(il.Trap(IMM_O(i.Operands[0])))
	// 	return false
	// case ARM64_YIELD:
	// 	il.AddInstruction(il.Intrinsic({}, ARM64_INTRIN_YIELD, {}))
	// 	break
	// default:
	// 	break
	// }

	// if(il.GetInstructionCount() > n_instrs_before)
	// 	return true

	// NeonGetLowLevelILForInstruction(arch, addr, il, i, addrSize)
	// if(il.GetInstructionCount() > n_instrs_before)
	// 	return true

	// il.AddInstruction(il.Unimplemented())
	return true
}
