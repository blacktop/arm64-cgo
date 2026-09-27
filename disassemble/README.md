# Disassemble Package

This package provides Go bindings for the Binary Ninja ARM64 disassembler with ARMv9A support.

## Architecture

The package uses CGO to interface with the C disassembler code, providing type-safe Go wrappers around C enums and functions.

### Key Files

- **C Code**: Core disassembler implementation from Binary Ninja
  - `decode*.c/h` - Instruction decoding logic
  - `format.c/h` - Instruction formatting
  - `operations.h` - Operation enum definitions (1689 operations!)
  - `sysregs_gen.h` - System register definitions
  - `sysregs_fmt_gen.h` - System register formatting

- **Go Wrappers** (Auto-generated):
  - `operations_cgo.go` - Wraps C enum Operation
  - `sysregs_cgo.go` - Wraps C enum SystemReg
  - `encodings_cgo.go` - Wraps C enum ENCODING

- **Main Interface**:
  - `disassemble.go` - Main Go API for disassembly

## CGO Integration

As of September 2025, this package uses direct CGO integration for operations and system registers rather than maintaining duplicate Go constants. This approach ensures:

1. **Single source of truth** - C headers define all values
2. **Automatic synchronization** - C updates are immediately available in Go
3. **No manual maintenance** - No need to update Go constants when C code changes
4. **Type safety** - Go's type system still enforced via wrapper types

## Breaking Change: Encoding Constants

`ENC_*` constants and `Encoding.String()` are now generated from the C `enum ENCODING`
(`encodings_cgo.go`). The previous hand-maintained `encoding.go` had drifted from the C enum:
nearly every constant held the wrong value, so comparisons against `Instruction.Encoding`,
`Encoding.String()`, the JSON `enc` field and `Blocks()` gave wrong results.

- Every retained encoding constant except `ENC_UNKNOWN` changed value to match the decoder.
- 959 names that are not in the C enum were removed, with no compatibility aliases:
  - 878 RESERVED/UNALLOCATED placeholders.
  - 34 predicated SVE encodings that are now split by predication, e.g. `ENC_ABS_Z_P_Z_` is now
    `ENC_ABS_Z_P_Z_M` (merging) and `ENC_ABS_Z_P_Z_Z` (zeroing).
  - 47 encodings that were renamed or merged, mostly aliases that now include the alias's own
    operand form, e.g. `ENC_MOV_DUP_Z_I_` is now `ENC_MOV_Z_I__DUP_Z_I_` and
    `ENC_MOV_MOVA_Z_P_RZA_B` is now `ENC_MOV_Z_P_RZA_B_MOVA_Z_P_RZA_B`. `ENC_DUP_P_P_PI_` is now
    `ENC_PSEL_P_PPI_`, and `ENC_FCMLA_ASIMDELEM_C_H`/`_C_S` merged into `ENC_FCMLA_ADVSIMD_ELT`.
    Search `encodings_cgo.go` for the base encoding name to find the new constant.
- `Encoding.String()` and JSON `enc` return the names from `enc_to_str()` in `encodings_fmt.c`
  (e.g. `B_only_branch_imm`), which differ in case and spelling from the old strings.

## Regenerating CGO Wrappers

When the C headers are updated (e.g., new instructions added), regenerate the Go wrappers:

```bash
python3 scripts/generate_cgo_wrappers.py
```

This script:
- Parses `operations.h` to extract all ARM64 operation enums
- Parses `sysregs_gen.h` to extract system register enums
- Parses `encodings_dec.h` to extract encoding enums and `encodings_fmt.c` for their names
- Generates `operations_cgo.go` with all operations properly categorized
- Generates `sysregs_cgo.go` with commonly used system registers
- Generates `encodings_cgo.go` with all encodings
- Runs `gofmt` on the generated files

CI reruns the script and fails if the committed wrappers differ, so edit the C headers or the
script, never the generated `*_cgo.go` files.

## Usage

```go
import "github.com/blacktop/arm64-cgo/disassemble"

// Disassemble a single instruction
instruction := uint32(0xb1007c00)  // adds x0, x0, #31
result, err := disassemble.Disassemble(0x1000, instruction, &buffer)

// Check operation type
if instr.Operation == disassemble.ARM64_ADDS {
    // Handle ADDS instruction
}

// Get string representation
fmt.Println(instr.Operation.String())  // prints "adds"
```