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

- **Main Interface**:
  - `disassemble.go` - Main Go API for disassembly

## CGO Integration

As of September 2025, this package uses direct CGO integration for operations and system registers rather than maintaining duplicate Go constants. This approach ensures:

1. **Single source of truth** - C headers define all values
2. **Automatic synchronization** - C updates are immediately available in Go
3. **No manual maintenance** - No need to update Go constants when C code changes
4. **Type safety** - Go's type system still enforced via wrapper types

## Regenerating CGO Wrappers

When the C headers are updated (e.g., new instructions added), regenerate the Go wrappers:

```bash
python3 scripts/generate_cgo_wrappers.py
```

This script:
- Parses `operations.h` to extract all ARM64 operation enums
- Parses `sysregs_gen.h` to extract system register enums
- Generates `operations_cgo.go` with all operations properly categorized
- Generates `sysregs_cgo.go` with commonly used system registers

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