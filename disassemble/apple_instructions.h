/* Apple ARM64 Proprietary Instructions
 * Based on Capstone PR #2692: https://github.com/capstone-engine/capstone/pull/2692
 *
 * These encodings are from the test cases in the PR and are in little-endian format
 */

#pragma once

/* Apple AMX (Matrix Extension) Instructions
 * Base pattern: 0x00201xxx (little-endian: 0x1xxx2000)
 * These instructions operate on the AMX coprocessor
 */

// AMX Load/Store Instructions
#define AMX_LDX_ENCODING    0x00102000  // ldx - Load X register to AMX
#define AMX_LDY_ENCODING    0x20102000  // ldy - Load Y register to AMX
#define AMX_LDZ_ENCODING    0x80102000  // ldz - Load Z register to AMX
#define AMX_LDZI_ENCODING   0xC0102000  // ldzi - Load Z register immediate to AMX
#define AMX_STX_ENCODING    0x40102000  // stx - Store X register from AMX
#define AMX_STY_ENCODING    0x60102000  // sty - Store Y register from AMX
#define AMX_STZ_ENCODING    0xA0102000  // stz - Store Z register from AMX
#define AMX_STZI_ENCODING   0xE0102000  // stzi - Store Z register immediate from AMX

// AMX Extract Instructions
#define AMX_EXTRX_ENCODING  0x00112000  // extrx - Extract from X
#define AMX_EXTRY_ENCODING  0x20112000  // extry - Extract from Y

// AMX FMA Instructions (Fused Multiply-Add)
#define AMX_FMA64_ENCODING  0x40112000  // fma64 - 64-bit FMA
#define AMX_FMS64_ENCODING  0x60112000  // fms64 - 64-bit FMS
#define AMX_FMA32_ENCODING  0x80112000  // fma32 - 32-bit FMA
#define AMX_FMS32_ENCODING  0xA0112000  // fms32 - 32-bit FMS
#define AMX_FMA16_ENCODING  0xE0112000  // fma16 - 16-bit FMA
#define AMX_FMS16_ENCODING  0x00122000  // fms16 - 16-bit FMS

// AMX MAC Instruction
#define AMX_MAC16_ENCODING  0xC0112000  // mac16 - 16-bit MAC

// AMX Set/Clear
#define AMX_SET_ENCODING    0x20122000  // set - Set AMX state
#define AMX_CLR_ENCODING    0x21122000  // clr - Clear AMX state

// AMX Vector/Matrix Operations
#define AMX_VECINT_ENCODING 0x40122000  // vecint - Vector integer operation
#define AMX_VECFP_ENCODING  0x60122000  // vecfp - Vector FP operation
#define AMX_MATINT_ENCODING 0x80122000  // matint - Matrix integer operation
#define AMX_MATFP_ENCODING  0xA0122000  // matfp - Matrix FP operation
#define AMX_GENLUT_ENCODING 0xC0122000  // genlut - Generate lookup table

/* Guarded Execution Mode Instructions
 * Pattern: 0x002014xx
 */
#define GEXIT_ENCODING      0x00142000  // gexit - Exit guarded mode
#define GENTER_BASE         0x00142000  // genter - Enter guarded mode (base)
#define GENTER_IMM_MASK     0xFF         // Immediate mask for genter

/* Memory Compression Instructions (WKdm - Wilson-Kaplan direct-mapped)
 * Pattern: 0x002008xx
 */
#define WKDMC_ENCODING      0x22082000  // wkdmc - Compress memory page
#define WKDMD_ENCODING      0x62082000  // wkdmd - Decompress memory page

/* Special Multiply (53-bit)
 * These are SIMD instructions for 53-bit multiplication
 */
// Need exact encodings from MUL53HI/MUL53LO instructions

/* Speculation Data Synchronization Barrier
 * SDSB with different barrier options
 */
// Need exact encodings for SDSB variants

/* Helper macros for decoding Apple instructions */
#define IS_AMX_INSTRUCTION(opcode) \
    (((opcode) & 0xFFE0F000) == 0x00102000 || \
     ((opcode) & 0xFFE0F000) == 0x00112000 || \
     ((opcode) & 0xFFE0F000) == 0x00122000)

#define IS_GENTER_INSTRUCTION(opcode) \
    (((opcode) & 0xFFFFFF00) == 0x00142000 && ((opcode) & 0xFF) != 0)

#define IS_GEXIT_INSTRUCTION(opcode) \
    ((opcode) == 0x00142000)

#define IS_WKDM_INSTRUCTION(opcode) \
    ((opcode) == WKDMC_ENCODING || (opcode) == WKDMD_ENCODING)

#define GET_AMX_OPERATION(opcode) \
    (((opcode) >> 5) & 0x7)  // Bits [7:5] determine the AMX operation

#define GET_GENTER_IMMEDIATE(opcode) \
    ((opcode) & 0xFF)  // Low 8 bits are the immediate value