/* Apple ARM64 Proprietary Instruction Decoder
 * Handles Apple-specific instructions (AMX, GENTER/GEXIT, WKDM, etc.)
 */

#include "decode.h"
#include "encodings_dec.h"
#include "apple_instructions.h"

/* Decode Apple AMX instructions
 * Pattern: 0x002010xx through 0x002012xx
 */
int decode_apple_amx(uint32_t opcode, context* ctx, Instruction* instr)
{
    // AMX Load/Store operations (0x002010xx)
    if ((opcode & 0xFFFFFF00) == 0x00102000) {
        uint8_t op = opcode & 0xFF;

        // Extract register from bits [4:0]
        uint8_t reg = op & 0x1F;
        uint8_t variant = (op >> 5) & 0x7;

        switch (variant) {
            case 0: // LDX
                instr->encoding = ENC_AMX_LDX;
                break;
            case 1: // LDY
                instr->encoding = ENC_AMX_LDY;
                break;
            case 2: // STX
                instr->encoding = ENC_AMX_STX;
                break;
            case 3: // STY
                instr->encoding = ENC_AMX_STY;
                break;
            case 4: // LDZ
                instr->encoding = ENC_AMX_LDZ;
                break;
            case 5: // STZ
                instr->encoding = ENC_AMX_STZ;
                break;
            case 6: // LDZI
                instr->encoding = ENC_AMX_LDZI;
                break;
            case 7: // STZI
                instr->encoding = ENC_AMX_STZI;
                break;
        }

        // Set operand - X register
        ctx->n = reg;
        return DECODE_STATUS_OK;
    }

    // AMX Extract/FMA operations (0x002011xx)
    if ((opcode & 0xFFFFFF00) == 0x00112000) {
        uint8_t op = opcode & 0xFF;
        uint8_t reg = op & 0x1F;
        uint8_t variant = (op >> 5) & 0x7;

        switch (variant) {
            case 0: // EXTRX
                instr->encoding = ENC_AMX_EXTRX;
                break;
            case 1: // EXTRY
                instr->encoding = ENC_AMX_EXTRY;
                break;
            case 2: // FMA64
                instr->encoding = ENC_AMX_FMA64;
                break;
            case 3: // FMS64
                instr->encoding = ENC_AMX_FMS64;
                break;
            case 4: // FMA32
                instr->encoding = ENC_AMX_FMA32;
                break;
            case 5: // FMS32
                instr->encoding = ENC_AMX_FMS32;
                break;
            case 6: // MAC16
                instr->encoding = ENC_AMX_MAC16;
                break;
            case 7: // FMA16
                instr->encoding = ENC_AMX_FMA16;
                break;
        }

        ctx->n = reg;
        return DECODE_STATUS_OK;
    }

    // AMX Vector/Matrix operations (0x002012xx)
    if ((opcode & 0xFFFFFF00) == 0x00122000) {
        uint8_t op = opcode & 0xFF;
        uint8_t reg = op & 0x1F;
        uint8_t variant = (op >> 5) & 0x7;

        switch (variant) {
            case 0: // FMS16
                instr->encoding = ENC_AMX_FMS16;
                break;
            case 1: // SET/CLR
                if (op == 0x20) {
                    instr->encoding = ENC_AMX_SET;
                } else if (op == 0x21) {
                    instr->encoding = ENC_AMX_CLR;
                } else {
                    return DECODE_STATUS_UNMATCHED; // Unknown SET/CLR variant
                }
                break;
            case 2: // VECINT
                instr->encoding = ENC_AMX_VECINT;
                break;
            case 3: // VECFP
                instr->encoding = ENC_AMX_VECFP;
                break;
            case 4: // MATINT
                instr->encoding = ENC_AMX_MATINT;
                break;
            case 5: // MATFP
                instr->encoding = ENC_AMX_MATFP;
                break;
            case 6: // GENLUT
                instr->encoding = ENC_AMX_GENLUT;
                break;
            default: // variant 7 or any other unexpected value
                return DECODE_STATUS_UNMATCHED;
        }

        if (variant != 1) { // Not SET/CLR
            ctx->n = reg;
        }
        return DECODE_STATUS_OK;
    }

    return DECODE_STATUS_UNMATCHED;
}

/* Decode Apple guarded execution instructions */
int decode_apple_guarded(uint32_t opcode, context* ctx, Instruction* instr)
{
    // GEXIT: 0x00201400 (exact match)
    if (opcode == 0x00201400) {
        instr->encoding = ENC_GEXIT;
        return DECODE_STATUS_OK;
    }

    // GENTER: 0x002014xx (bits [4:0] contain the immediate, bit 5 is always set)
    // GENTER #0 is 0x00201420, GENTER #4 is 0x00201424
    if ((opcode & 0xFFFFFFE0) == 0x00201420) {
        instr->encoding = ENC_GENTER;
        ctx->imm = (opcode & 0x1F); // Immediate value in bits [4:0]
        return DECODE_STATUS_OK;
    }

    return DECODE_STATUS_UNMATCHED;
}

/* Decode Apple memory compression instructions */
int decode_apple_wkdm(uint32_t opcode, context* ctx, Instruction* instr)
{
    // WKDMC: 0x00200822
    if (opcode == 0x00200822) {
        instr->encoding = ENC_WKDMC;
        // Extract registers from encoding
        ctx->n = (opcode >> 5) & 0x1F;  // Rn
        ctx->t = opcode & 0x1F;         // Rt
        return DECODE_STATUS_OK;
    }

    // WKDMD: 0x00200862
    if (opcode == 0x00200862) {
        instr->encoding = ENC_WKDMD;
        // Extract registers from encoding
        ctx->n = (opcode >> 5) & 0x1F;  // Rn
        ctx->t = opcode & 0x1F;         // Rt
        return DECODE_STATUS_OK;
    }

    return DECODE_STATUS_UNMATCHED;
}

/* Main Apple instruction decoder entry point */
int decode_apple_instruction(uint32_t opcode, context* ctx, Instruction* instr)
{
    int rc;

    // Check for guarded execution first (0x002014xx pattern)
    rc = decode_apple_guarded(opcode, ctx, instr);
    if (rc == DECODE_STATUS_OK) return rc;

    // Check for AMX instructions (0x0020xxxx and 0x0010xxxx patterns)
    if ((opcode & 0xFFF00000) == 0x00200000 || (opcode & 0xFFF00000) == 0x00100000) {
        // Try AMX instructions
        rc = decode_apple_amx(opcode, ctx, instr);
        if (rc == DECODE_STATUS_OK) return rc;

        // Try WKdm compression
        rc = decode_apple_wkdm(opcode, ctx, instr);
        if (rc == DECODE_STATUS_OK) return rc;
    }

    // MUL53 instructions would be in SIMD space
    // SDSB would be in barrier space
    // These need more investigation for exact patterns

    return DECODE_STATUS_UNMATCHED;
}