/* Apple ARM64 Extensions Implementation Header
 * Based on:
 * - https://github.com/AsahiLinux/docs/blob/main/docs/hw/cpu/apple-instructions.md
 * - https://github.com/capstone-engine/capstone/pull/2692
 */

#pragma once

/* Apple-specific instruction opcodes
 * These are implementation-defined encodings that need special handling
 */

// AMX (Apple Matrix Extension) instructions
// These use implementation-specific encodings and trap by default
#define AMX_OPCODE_MASK     0xFFE0FC00
#define AMX_OPCODE_VALUE    0x00201000  // Base encoding for AMX ops

// Guarded execution mode instructions
// These likely use HINT space encodings
#define GENTER_ENCODING     0xD503201F  // HINT #0x00 (tentative)
#define GEXIT_ENCODING      0xD503203F  // HINT #0x01 (tentative)

// Memory compression instructions
// Implementation-specific encodings
#define WKDMC_ENCODING      0x00000000  // TBD - needs exact encoding
#define WKDMD_ENCODING      0x00000000  // TBD - needs exact encoding

// Special multiply
#define MUL53_ENCODING      0x00000000  // TBD - needs exact encoding

// Apple system registers (already in sysregs_gen.h but listed for reference)
// AMX_STATE_T_EL1  - S3_4_c15_c1_3
// AMX_CONFIG_EL1   - S3_4_c15_c1_4
// AMX_STATE_EL1    - S3_4_c15_c3_0
// AMX_STATUS_EL1   - S3_4_c15_c3_6
// AMX_CONFIG_EL12  - S3_4_c15_c4_6
// AMX_CONFIG_EL2   - S3_4_c15_c4_7
// AMXIDR_EL1       - S3_6_c15_c2_7

/* Helper function to detect Apple instructions */
static inline int is_apple_instruction(uint32_t opcode) {
    // Check for AMX instruction pattern
    if ((opcode & AMX_OPCODE_MASK) == AMX_OPCODE_VALUE) {
        return 1;
    }

    // Check for specific Apple instructions
    switch(opcode) {
        case GENTER_ENCODING:
        case GEXIT_ENCODING:
        // Add other known encodings when available
            return 1;
        default:
            return 0;
    }
}

/* Decode Apple-specific instructions
 * Returns 1 if decoded successfully, 0 otherwise
 */
static inline int decode_apple_instruction(uint32_t opcode, Instruction *instr) {
    // AMX instructions
    if ((opcode & AMX_OPCODE_MASK) == AMX_OPCODE_VALUE) {
        // TODO: Decode specific AMX instruction variant
        instr->operation = ARM64_HINT;  // Placeholder until we add AMX opcodes
        return 1;
    }

    // GENTER/GEXIT
    if (opcode == GENTER_ENCODING) {
        instr->operation = ARM64_HINT;  // Placeholder until we add GENTER opcode
        return 1;
    }
    if (opcode == GEXIT_ENCODING) {
        instr->operation = ARM64_HINT;  // Placeholder until we add GEXIT opcode
        return 1;
    }

    return 0;
}