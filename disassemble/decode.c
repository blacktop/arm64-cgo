#include "decode.h"
#include "feature_flags.h"

int decode_spec(context* ctx, Instruction* dec);        // from decode0.cpp
int decode_scratchpad(context* ctx, Instruction* dec);  // from decode_scratchpad.c
int decode_apple_instruction(uint32_t opcode, context* ctx, Instruction* instr); // from decode_apple.c

int aarch64_decompose(uint32_t instructionValue, Instruction* instr, uint64_t address)
{
	context ctx = {0};
	ctx.halted = 1;  // enable disassembly of exception instructions like DCPS1
	ctx.insword = instructionValue;
	ctx.address = address;
	ARCH_FEATURES_ENABLE_ALL(ctx.decode_features);
	ARCH_FEATURES_ENABLE_ALL(ctx.pcode_features);
	ctx.EDSCR_HDE = 1;

	/* Try Apple proprietary instructions first (if enabled) */
	int rc = decode_apple_instruction(instructionValue, &ctx, instr);
	if (rc == DECODE_STATUS_OK) {
		/* Apple instruction decoded successfully */
		return decode_scratchpad(&ctx, instr);
	}

	/* have the spec-generated code populate all the pcode variables */
	rc = decode_spec(&ctx, instr);

	if (rc != DECODE_STATUS_OK)
	{
		/* exceptional cases where we accept a non-OK decode status */
		if (rc == DECODE_STATUS_END_OF_INSTRUCTION && instr->encoding == ENC_HINT_HM_HINTS)
		{
			while (0)
				;
		}
		/* no exception! fail! */
		else
			return rc;
	}

	/* if UDF encoding, return undefined */
	// if(instr->encoding == ENC_UDF_ONLY_PERM_UNDEF)
	//	return DECODE_STATUS_UNDEFINED;

	/* convert the pcode variables to list of operands, etc. */
	return decode_scratchpad(&ctx, instr);
}
