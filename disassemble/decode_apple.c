/* Apple ARM64 Proprietary Instruction Decoder
 * Handles Apple-specific instructions (AMX, GENTER/GEXIT, WKDM) in the 0x0020xxxx range.
 * See apple_extensions.md for the encodings.
 */

#include "decode.h"
#include "encodings_dec.h"
#include "pcode.h"

#define AMX_OP_SET_CLR 17

static const enum ENCODING amx_encodings[] = {
	ENC_AMX_LDX, ENC_AMX_LDY, ENC_AMX_STX, ENC_AMX_STY,
	ENC_AMX_LDZ, ENC_AMX_STZ, ENC_AMX_LDZI, ENC_AMX_STZI,
	ENC_AMX_EXTRX, ENC_AMX_EXTRY, ENC_AMX_FMA64, ENC_AMX_FMS64,
	ENC_AMX_FMA32, ENC_AMX_FMS32, ENC_AMX_MAC16, ENC_AMX_FMA16,
	ENC_AMX_FMS16, ENC_UNKNOWN /* SET/CLR */, ENC_AMX_VECINT, ENC_AMX_VECFP,
	ENC_AMX_MATINT, ENC_AMX_MATFP, ENC_AMX_GENLUT,
};

/* AMX: 0x00201000 | op << 5 | Xn, op 0..22 */
static int decode_apple_amx(uint32_t opcode, context* ctx, Instruction* instr)
{
	if ((opcode & 0xFFFFFC00) != 0x00201000)
		return DECODE_STATUS_UNMATCHED;

	uint32_t op = SLICE(opcode, 9, 5);
	uint32_t reg = SLICE(opcode, 4, 0);
	if (op >= sizeof(amx_encodings) / sizeof(amx_encodings[0]))
		return DECODE_STATUS_UNMATCHED;

	if (op == AMX_OP_SET_CLR) {
		if (reg > 1)
			return DECODE_STATUS_UNMATCHED;
		instr->encoding = reg ? ENC_AMX_CLR : ENC_AMX_SET;
		return DECODE_STATUS_OK;
	}

	instr->encoding = amx_encodings[op];
	ctx->n = reg;
	return DECODE_STATUS_OK;
}

/* GEXIT: 0x00201400, GENTER: 0x00201420 | imm5 */
static int decode_apple_guarded(uint32_t opcode, context* ctx, Instruction* instr)
{
	if (opcode == 0x00201400) {
		instr->encoding = ENC_GEXIT;
		return DECODE_STATUS_OK;
	}

	if ((opcode & 0xFFFFFFE0) == 0x00201420) {
		instr->encoding = ENC_GENTER;
		ctx->imm = SLICE(opcode, 4, 0);
		return DECODE_STATUS_OK;
	}

	return DECODE_STATUS_UNMATCHED;
}

/* WKDMC: 0x00200800 | rD << 5 | rS, WKDMD: 0x00200c00 | rD << 5 | rS
 * rS is the source address and receives the status, rD is the destination.
 */
static int decode_apple_wkdm(uint32_t opcode, context* ctx, Instruction* instr)
{
	switch (opcode & 0xFFFFFC00) {
		case 0x00200800:
			instr->encoding = ENC_WKDMC;
			break;
		case 0x00200C00:
			instr->encoding = ENC_WKDMD;
			break;
		default:
			return DECODE_STATUS_UNMATCHED;
	}

	ctx->d = SLICE(opcode, 9, 5);
	ctx->s = SLICE(opcode, 4, 0);
	return DECODE_STATUS_OK;
}

int decode_apple_instruction(uint32_t opcode, context* ctx, Instruction* instr)
{
	if ((opcode & 0xFFFF0000) != 0x00200000)
		return DECODE_STATUS_UNMATCHED;

	if (decode_apple_guarded(opcode, ctx, instr) == DECODE_STATUS_OK)
		return DECODE_STATUS_OK;
	if (decode_apple_amx(opcode, ctx, instr) == DECODE_STATUS_OK)
		return DECODE_STATUS_OK;
	return decode_apple_wkdm(opcode, ctx, instr);
}
