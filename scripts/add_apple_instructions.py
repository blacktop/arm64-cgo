#!/usr/bin/env python3
"""
Add Apple-specific ARM64 instructions to operations.h
"""

import sys
from pathlib import Path

# Apple instructions to add (based on Capstone PR #2692)
APPLE_INSTRUCTIONS = [
    # AMX (Apple Matrix Extension) - Load/Store
    "ARM64_AMX_LDX",     # Load X register to AMX
    "ARM64_AMX_LDY",     # Load Y register to AMX
    "ARM64_AMX_LDZ",     # Load Z register to AMX
    "ARM64_AMX_LDZI",    # Load Z register immediate to AMX
    "ARM64_AMX_STX",     # Store X register from AMX
    "ARM64_AMX_STY",     # Store Y register from AMX
    "ARM64_AMX_STZ",     # Store Z register from AMX
    "ARM64_AMX_STZI",    # Store Z register immediate from AMX

    # AMX Extract
    "ARM64_AMX_EXTRX",   # Extract from X
    "ARM64_AMX_EXTRY",   # Extract from Y

    # AMX FMA/FMS Operations
    "ARM64_AMX_FMA16",   # 16-bit FMA
    "ARM64_AMX_FMA32",   # 32-bit FMA
    "ARM64_AMX_FMA64",   # 64-bit FMA
    "ARM64_AMX_FMS16",   # 16-bit FMS
    "ARM64_AMX_FMS32",   # 32-bit FMS
    "ARM64_AMX_FMS64",   # 64-bit FMS
    "ARM64_AMX_MAC16",   # 16-bit MAC

    # AMX Vector/Matrix Operations
    "ARM64_AMX_VECINT",  # Vector integer operation
    "ARM64_AMX_VECFP",   # Vector FP operation
    "ARM64_AMX_MATINT",  # Matrix integer operation
    "ARM64_AMX_MATFP",   # Matrix FP operation
    "ARM64_AMX_GENLUT",  # Generate lookup table

    # AMX Control
    "ARM64_AMX_SET",     # Set AMX state
    "ARM64_AMX_CLR",     # Clear AMX state

    # Guarded execution
    "ARM64_GENTER",      # Enter guarded mode
    "ARM64_GEXIT",       # Exit guarded mode

    # Memory compression (WKdm - Wilson-Kaplan direct-mapped)
    "ARM64_WKDMC",       # Compress memory page
    "ARM64_WKDMD",       # Decompress memory page

    # Special multiply (53-bit for JavaScript double precision)
    "ARM64_MUL53HI",     # 53-bit multiply high
    "ARM64_MUL53LO",     # 53-bit multiply low

    # Speculation barrier
    "ARM64_SDSB",        # Speculative data synchronization barrier
]

def add_apple_instructions_to_header(header_path):
    """Add Apple instructions to operations.h"""

    with open(header_path, 'r') as f:
        lines = f.readlines()

    # Find where to insert (before ARM64_END or at end of enum)
    insert_idx = -1
    in_enum = False

    for i, line in enumerate(lines):
        if 'enum Operation {' in line:
            in_enum = True
        elif in_enum:
            if 'ARM64_END' in line or '};' in line:
                insert_idx = i
                break

    if insert_idx == -1:
        print("Could not find insertion point in operations.h")
        return False

    # Create the lines to insert
    insert_lines = [
        "\t// Apple-specific instructions\n",
    ]
    for instr in APPLE_INSTRUCTIONS:
        insert_lines.append(f"\t{instr},\n")
    insert_lines.append("\n")

    # Insert the new instructions
    lines[insert_idx:insert_idx] = insert_lines

    # Write back
    with open(header_path, 'w') as f:
        f.writelines(lines)

    print(f"Added {len(APPLE_INSTRUCTIONS)} Apple instructions to {header_path}")
    return True


def update_operations_count(header_path):
    """Update the operations count in format strings"""

    # After adding instructions, regenerate the CGO wrappers
    import subprocess
    script_dir = Path(__file__).parent
    gen_script = script_dir / "generate_cgo_wrappers.py"

    if gen_script.exists():
        subprocess.run([sys.executable, str(gen_script)])
        print("Regenerated CGO wrappers")


def main():
    script_dir = Path(__file__).parent
    project_root = script_dir.parent
    operations_h = project_root / 'disassemble' / 'operations.h'

    if not operations_h.exists():
        print(f"Error: {operations_h} not found", file=sys.stderr)
        sys.exit(1)

    # Backup the original
    import shutil
    backup = operations_h.with_suffix('.h.bak')
    shutil.copy2(operations_h, backup)
    print(f"Created backup: {backup}")

    # Add Apple instructions
    if add_apple_instructions_to_header(operations_h):
        print("\n✅ Successfully added Apple instructions to operations.h")
        print("\nNext steps:")
        print("1. Update decode.c to handle these new instruction encodings")
        print("2. Update format.c to format these instructions properly")
        print("3. Run: python3 scripts/generate_cgo_wrappers.py")
        print("4. Test on Apple Silicon hardware")
    else:
        print("\n❌ Failed to add Apple instructions")
        print(f"Restoring from backup: {backup}")
        shutil.copy2(backup, operations_h)
        sys.exit(1)


if __name__ == '__main__':
    main()