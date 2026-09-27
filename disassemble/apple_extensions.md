# Instructions Beyond the Generated Decoder

The instructions below are not produced by the generated decoder (`decode0.c`), either because
they are Apple proprietary or because they postdate its spec data. They are matched by hand
before the generated decoder runs (`decode.c`).

## Apple Proprietary Instructions (`decode_apple.c`)

Apple instructions live in the `0x0020xxxx` range and fault unless the kernel enables them.

| Instruction | Encoding | Syntax |
| --- | --- | --- |
| AMX | `0x00201000 \| op << 5 \| Xn`, op 0..22 (`0x00201000`–`0x002012df`) | `ldx x0`, `fma64 x0`, ... |
| AMX SET/CLR | `0x00201220` / `0x00201221` (op 17) | `set` / `clr` |
| WKDMC | `0x00200800 \| rD << 5 \| rS` | `wkdmc xS, xD` |
| WKDMD | `0x00200c00 \| rD << 5 \| rS` | `wkdmd xS, xD` |
| GEXIT | `0x00201400` | `gexit` |
| GENTER | `0x00201420 \| imm5` | `genter #imm` |

AMX op numbers: 0 ldx, 1 ldy, 2 stx, 3 sty, 4 ldz, 5 stz, 6 ldzi, 7 stzi, 8 extrx, 9 extry,
10 fma64, 11 fms64, 12 fma32, 13 fms32, 14 mac16, 15 fma16, 16 fms16, 17 set/clr, 18 vecint,
19 vecfp, 20 matint, 21 matfp, 22 genlut.

Not decoded yet: MUL53 (`0x00200000`–`0x002007ff`), `at_as1elx` (`0x00201440 | Xa`) and
SDSB (`0x00201460`–`0x00201463`). `ENC_MUL53HI`, `ENC_MUL53LO` and `ENC_SDSB` exist but
nothing produces them.

## Arm Future Architecture Technologies (`decode_poe2.c`)

ArmvFAT.0 instructions used by macOS 27 (the Mac18,5 kernelcache and TXM call into
guarded code with `tenter` instead of `genter`). Encodings follow binutils.

| Instruction | Feature | Encoding | Syntax |
| --- | --- | --- | --- |
| TENTER | FEAT_TEV | `0xd4e00000 \| NB << 17 \| imm7 << 5` | `tenter #imm{, nb}` |
| TEXIT | FEAT_TEV | `0xd6ff03e0 \| NB << 10` | `texit{ nb}` |
| TCHANGEF (register) | FEAT_S1POE2 | `0xd5800000 \| NB << 17 \| Xn << 5 \| Xd` | `tchangef xd, xn{, nb}` |
| TCHANGEF (immediate) | FEAT_S1POE2 | `0xd5900000 \| NB << 17 \| imm7 << 5 \| Xd` | `tchangef xd, #imm{, nb}` |
| TCHANGEB (register) | FEAT_S1POE2 | `0xd5840000 \| NB << 17 \| Xn << 5 \| Xd` | `tchangeb xd, xn{, nb}` |
| TCHANGEB (immediate) | FEAT_S1POE2 | `0xd5940000 \| NB << 17 \| imm7 << 5 \| Xd` | `tchangeb xd, #imm{, nb}` |

## References

- [Asahi Linux Apple Instructions Documentation](https://github.com/AsahiLinux/docs/blob/main/docs/hw/cpu/apple-instructions.md)
- [Dougall Johnson's AMX notes](https://gist.github.com/dougallj/7a75a3be1ec69ca550e7c36dc75e0d6f)
- [Arm feature names: Future Architecture Technologies](https://support.arm.com/documentation/109697/2026_06/Feature-descriptions/Future-Architecture-Technologies)
- binutils "aarch64: Add support for POE2 instructions" and "aarch64: Add support for TEV instructions"
