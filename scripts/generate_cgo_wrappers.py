#!/usr/bin/env python3
"""
Generate CGO wrapper files for ARM64 operations and system registers.
Parses C header files and generates Go code with proper CGO bindings.
"""

import re
import sys
from pathlib import Path


def parse_operations_header(header_path):
    """Parse operations.h to extract all ARM64 operation enums."""
    operations = []

    with open(header_path, 'r') as f:
        content = f.read()

    # Find the enum Operation block
    enum_pattern = r'enum\s+Operation\s*\{([^}]+)\}'
    match = re.search(enum_pattern, content, re.DOTALL)

    if not match:
        raise ValueError("Could not find 'enum Operation' in header file")

    enum_content = match.group(1)

    # Extract each enum value
    for line in enum_content.split('\n'):
        line = line.strip()
        if line.startswith('ARM64_'):
            # Handle lines like "ARM64_ADD," or "ARM64_ERROR"
            op_name = line.split(',')[0].split('=')[0].strip()
            operations.append(op_name)

    return operations


def parse_sysregs_header(header_path):
    """Parse sysregs_gen.h to extract system register enums."""
    sysregs = []

    with open(header_path, 'r') as f:
        content = f.read()

    # Find the enum SystemReg block
    enum_pattern = r'enum\s+SystemReg\s*\{([^}]+)\}'
    match = re.search(enum_pattern, content, re.DOTALL)

    if not match:
        raise ValueError("Could not find 'enum SystemReg' in header file")

    enum_content = match.group(1)

    # Extract each enum value - looking for commonly used ones
    for line in enum_content.split('\n'):
        line = line.strip()
        # Get SYSREG_NONE and a few common REG_ entries for examples
        if line.startswith('SYSREG_NONE') or (line.startswith('REG_') and len(sysregs) < 10):
            reg_name = line.split('=')[0].strip().rstrip(',')
            sysregs.append(reg_name)

    return sysregs


def categorize_operations(operations):
    """Categorize operations by type for better organization."""
    categories = {
        'Error': [],
        'Arithmetic': [],
        'Logical': [],
        'Bitfield': [],
        'Move': [],
        'Memory': [],
        'Branch': [],
        'Conditional Branch': [],
        'Compare and Branch': [],
        'Conditional Select': [],
        'System': [],
        'Synchronization': [],
        'PAC/Auth': [],
        'Address': [],
        'SIMD/FP': [],
        'SVE': [],
        'Crypto': [],
        'Other': []
    }

    for op in operations:
        name = op.replace('ARM64_', '')

        # Categorize based on instruction patterns
        if op == 'ARM64_ERROR':
            categories['Error'].append(op)
        elif name.startswith('B_'):
            categories['Conditional Branch'].append(op)
        elif name in ['ADD', 'ADDS', 'SUB', 'SUBS', 'MUL', 'MADD', 'MSUB', 'SDIV', 'UDIV',
                      'NEG', 'NEGS', 'ADC', 'ADCS', 'SBC', 'SBCS', 'MNEG', 'SMULL', 'UMULL',
                      'SMLAL', 'UMLAL', 'SMULH', 'UMULH']:
            categories['Arithmetic'].append(op)
        elif name in ['AND', 'ANDS', 'ORR', 'ORN', 'EOR', 'EON', 'BIC', 'BICS', 'TST', 'MVN']:
            categories['Logical'].append(op)
        elif name in ['BFM', 'SBFM', 'UBFM', 'BFI', 'BFXIL', 'SBFX', 'UBFX', 'SBFIZ', 'UBFIZ',
                      'LSL', 'LSR', 'ASR', 'ROR', 'EXTR']:
            categories['Bitfield'].append(op)
        elif name.startswith('MOV') or name in ['MOVN', 'MOVK']:
            categories['Move'].append(op)
        elif any(name.startswith(prefix) for prefix in ['LD', 'ST', 'PRFM', 'LDAR', 'STLR',
                                                         'LDXR', 'STXR', 'LDAPR', 'LDADD']):
            categories['Memory'].append(op)
        elif name in ['B', 'BL', 'BR', 'BLR', 'RET', 'ERET', 'DRPS', 'BTI'] or \
             name.startswith('RET') or name.startswith('ERET') or name.startswith('BLR') or \
             name.startswith('BRA'):
            categories['Branch'].append(op)
        elif name in ['CBZ', 'CBNZ', 'TBZ', 'TBNZ']:
            categories['Compare and Branch'].append(op)
        elif name.startswith('CS') or name.startswith('CINC') or name.startswith('CINV') or \
             name.startswith('CNEG') or name in ['CSEL', 'CSET', 'CSETM']:
            categories['Conditional Select'].append(op)
        elif name in ['MRS', 'MSR', 'NOP', 'HINT', 'HLT', 'BRK', 'SVC', 'HVC', 'SMC',
                      'SYS', 'SYSL', 'IC', 'DC', 'AT', 'TLBI', 'SB']:
            categories['System'].append(op)
        elif name in ['ISB', 'DSB', 'DMB', 'YIELD', 'WFE', 'WFI', 'SEV', 'SEVL', 'CLREX']:
            categories['Synchronization'].append(op)
        elif any(name.startswith(prefix) for prefix in ['PAC', 'AUT', 'XPAC']) or \
             any(name.endswith(suffix) for suffix in ['AA', 'AB', 'AZ', 'BZ']):
            categories['PAC/Auth'].append(op)
        elif name in ['ADR', 'ADRP']:
            categories['Address'].append(op)
        elif any(name.startswith(prefix) for prefix in ['FADD', 'FSUB', 'FMUL', 'FDIV', 'FABS',
                                                         'FNEG', 'FSQRT', 'FCMP', 'FCVT', 'FMOV',
                                                         'SCVTF', 'UCVTF', 'FCVTZ', 'FRINT']):
            categories['SIMD/FP'].append(op)
        elif name.startswith('Z') or any(name.startswith(prefix) for prefix in ['ADDVL', 'ADDPL',
                                                                                 'RDVL', 'RDPL']):
            categories['SVE'].append(op)
        elif any(name.startswith(prefix) for prefix in ['AES', 'SHA', 'SM3', 'SM4']):
            categories['Crypto'].append(op)
        else:
            categories['Other'].append(op)

    # Remove empty categories
    return {k: v for k, v in categories.items() if v}


def generate_operations_go(operations, output_path):
    """Generate the operations_cgo.go file."""

    categories = categorize_operations(operations)

    content = '''package disassemble

/*
#include "operations.h"
*/
import "C"

// Operation wraps the C enum Operation
type Operation C.enum_Operation

// Operation constants - generated from C enum Operation
const (
'''

    for category, ops in categories.items():
        if ops:
            content += f'\t// {category}\n'
            for op in sorted(ops):
                content += f'\t{op} = Operation(C.{op})\n'
            content += '\n'

    content += ''')

// String returns the string representation of an operation
func (o Operation) String() string {
	// Use the C formatting function
	cStr := C.operation_to_str(C.enum_Operation(o))
	if cStr != nil {
		return C.GoString(cStr)
	}
	return "UNKNOWN"
}

// IsValid checks if the operation is valid
func (o Operation) IsValid() bool {
	return o != Operation(C.ARM64_ERROR)
}
'''

    with open(output_path, 'w') as f:
        f.write(content)

    print(f"Generated {output_path}")
    print(f"  Total operations: {len(operations)}")
    print(f"  Categories: {len(categories)}")


def generate_sysregs_go(sysregs, output_path):
    """Generate the sysregs_cgo.go file."""

    content = '''package disassemble

/*
#include "sysregs_gen.h"
#include "sysregs_fmt_gen.h"
*/
import "C"

// SystemReg wraps the C enum SystemReg
type SystemReg C.enum_SystemReg

// SystemReg constants - commonly used ones
const (
'''

    for reg in sysregs:
        content += f'\t{reg} = SystemReg(C.{reg})\n'

    content += ''')

// String returns the string representation of a system register
func (s SystemReg) String() string {
	// Use the C formatting function
	cStr := C.get_system_register_name(C.enum_SystemReg(s))
	if cStr != nil {
		return C.GoString(cStr)
	}
	return "UNKNOWN"
}

// IsValid checks if the system register is valid
func (s SystemReg) IsValid() bool {
	return s != SystemReg(C.SYSREG_NONE)
}
'''

    with open(output_path, 'w') as f:
        f.write(content)

    print(f"Generated {output_path}")
    print(f"  Sample sysregs: {len(sysregs)}")


def main():
    # Set up paths
    script_dir = Path(__file__).parent
    project_root = script_dir.parent
    disassemble_dir = project_root / 'disassemble'

    # Parse headers
    operations_header = disassemble_dir / 'operations.h'
    sysregs_header = disassemble_dir / 'sysregs_gen.h'

    if not operations_header.exists():
        print(f"Error: {operations_header} not found", file=sys.stderr)
        sys.exit(1)

    if not sysregs_header.exists():
        print(f"Error: {sysregs_header} not found", file=sys.stderr)
        sys.exit(1)

    # Parse operations
    print(f"Parsing {operations_header}...")
    operations = parse_operations_header(operations_header)

    # Parse system registers
    print(f"Parsing {sysregs_header}...")
    sysregs = parse_sysregs_header(sysregs_header)

    # Generate Go files
    operations_go = disassemble_dir / 'operations_cgo.go'
    sysregs_go = disassemble_dir / 'sysregs_cgo.go'

    generate_operations_go(operations, operations_go)
    generate_sysregs_go(sysregs, sysregs_go)

    print("\nDone! Generated CGO wrapper files.")
    print("\nTo regenerate these files in the future, run:")
    print(f"  python3 {Path(__file__).relative_to(project_root)}")


if __name__ == '__main__':
    main()