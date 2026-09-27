/* FEAT_S1POE2 (TCHANGEF/TCHANGEB) and FEAT_TEV (TENTER/TEXIT) instruction decoder
 * These ArmvFAT.0 instructions postdate the spec the generated decoder was
 * built from, so they are matched here by hand.
 * Encodings from binutils (opcodes/aarch64-tbl.h, "aarch64: Add support for
 * POE2 instructions" and "aarch64: Add support for TEV instructions").
 */

#include "decode.h"
#include "encodings_dec.h"
#include "pcode.h"

/* TCHANGE{F,B} <Xd>, <Xn>{, NB}: 1101 0101 1000 0BN0 0000 00nn nnnd dddd */
#define TCHANGE_REG_MASK  0xFFF9FC00
#define TCHANGE_REG_VALUE 0xD5800000

/* TCHANGE{F,B} <Xd>, #<imm7>{, NB}: 1101 0101 1001 0BN0 0000 iiii iiid dddd */
#define TCHANGE_IMM_MASK  0xFFF9F000
#define TCHANGE_IMM_VALUE 0xD5900000

/* TENTER #<imm7>{, NB}: 1101 0100 1110 00N0 0000 iiii iii0 0000 */
#define TENTER_MASK  0xFFFDF01F
#define TENTER_VALUE 0xD4E00000

/* TEXIT{ NB}: 1101 0110 1111 1111 0000 0N11 1110 0000 */
#define TEXIT_MASK   0xFFFFFBFF
#define TEXIT_VALUE  0xD6FF03E0

static int decode_tchange(uint32_t opcode, context* ctx, Instruction* instr)
{
	int backward = SLICE(opcode, 18, 18);

	if ((opcode & TCHANGE_REG_MASK) == TCHANGE_REG_VALUE) {
		instr->encoding = backward ? ENC_TCHANGEB_REG : ENC_TCHANGEF_REG;
		ctx->n = SLICE(opcode, 9, 5);
	} else if ((opcode & TCHANGE_IMM_MASK) == TCHANGE_IMM_VALUE) {
		instr->encoding = backward ? ENC_TCHANGEB_IMM : ENC_TCHANGEF_IMM;
		ctx->imm7 = SLICE(opcode, 11, 5);
	} else {
		return DECODE_STATUS_UNMATCHED;
	}

	ctx->d = SLICE(opcode, 4, 0);
	ctx->imm1 = SLICE(opcode, 17, 17);
	return DECODE_STATUS_OK;
}

int decode_poe2(uint32_t opcode, context* ctx, Instruction* instr)
{
	if ((opcode & TENTER_MASK) == TENTER_VALUE) {
		instr->encoding = ENC_TENTER_EX_EXCEPTION;
		ctx->imm7 = SLICE(opcode, 11, 5);
		ctx->imm1 = SLICE(opcode, 17, 17);
		return DECODE_STATUS_OK;
	}

	if ((opcode & TEXIT_MASK) == TEXIT_VALUE) {
		instr->encoding = ENC_TEXIT_64E_BRANCH_REG;
		ctx->imm1 = SLICE(opcode, 10, 10);
		return DECODE_STATUS_OK;
	}

	return decode_tchange(opcode, ctx, instr);
}
