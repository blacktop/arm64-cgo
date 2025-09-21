# Apple ARM64 Extensions

This document describes Apple-specific ARM64 instructions and system registers that need to be added to the disassembler.

## Instructions

Based on Asahi Linux documentation and Capstone implementation:

### AMX (Apple Matrix Extension) Instructions

The AMX instructions use implementation-specific encodings in the `0x0020xxxx` range. These need to be enabled via system registers before use.

- **AMX configuration/state management** (exact encodings TBD)
- These instructions trap by default and require kernel support

### Guarded Execution Mode

- **GENTER** - Enter guarded execution mode
  - Encoding: Implementation-specific (likely HINT-space)
  - Faults by default without kernel support

- **GEXIT** - Exit guarded execution mode
  - Encoding: Implementation-specific (likely HINT-space)
  - Faults by default without kernel support

### Memory Compression

- **WKDMC** - Compress memory page
  - Encoding: Implementation-specific

- **WKDMD** - Decompress memory page
  - Encoding: Implementation-specific

### Other Apple Extensions

- **MUL53** - 53-bit multiply
  - Encoding: Implementation-specific

- **AT AS1ELx** - Address translation variants
  - Encoding: System instruction space

- **SDSB** variants (osh, nsh, ish, sy) - Speculative data synchronization barrier
  - Encoding: Barrier instruction space

## System Registers (Already Present)

The following Apple AMX system registers are already defined in `sysregs_gen.h`:

- `AMX_STATE_T_EL1` (S3_4_c15_c1_3)
- `AMX_CONFIG_EL1` (S3_4_c15_c1_4)
- `AMX_STATE_EL1` (S3_4_c15_c3_0)
- `AMX_STATUS_EL1` (S3_4_c15_c3_6)
- `AMX_CONFIG_EL12` (S3_4_c15_c4_6)
- `AMX_CONFIG_EL2` (S3_4_c15_c4_7)
- `AMXIDR_EL1` (S3_6_c15_c2_7)

## Implementation Notes

1. Many of these instructions use implementation-specific encodings that may be in:
   - HINT space (for NOP-compatible instructions)
   - System instruction space (SYS/SYSL encodings)
   - Implementation-specific regions (0x0020xxxx)

2. These instructions typically:
   - Fault by default (need kernel enablement)
   - Require specific CPU features/models
   - May be gated by system registers

3. To add support:
   - Need exact opcode encodings for each instruction
   - May need to handle as special cases in decode.c
   - Should add feature detection flags

## References

- [Asahi Linux Apple Instructions Documentation](https://github.com/AsahiLinux/docs/blob/main/docs/hw/cpu/apple-instructions.md)
- [Capstone PR #2692](https://github.com/capstone-engine/capstone/pull/2692)
- [Dougall's ARM64 Documentation](https://dougallj.github.io/applecpu/firestorm-int.html)

## TODO

To fully implement Apple instruction support:

1. Obtain exact opcode encodings for each instruction
2. Add instruction enum entries to operations.h
3. Update decode logic to recognize these encodings
4. Add formatting support in format.c
5. Test on Apple Silicon hardware