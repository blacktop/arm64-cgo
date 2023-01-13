# arm64-cgo

[![Go](https://github.com/blacktop/arm64-cgo/actions/workflows/go.yml/badge.svg)](https://github.com/blacktop/arm64-cgo/actions/workflows/go.yml)
![GitHub all releases](https://img.shields.io/github/downloads/blacktop/arm64-cgo/total)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/blacktop/arm64-cgo)](https://github.com/blacktop/arm64-cgo/releases/latest)
[![PkgGoDev](https://pkg.go.dev/badge/blacktop/arm64-cgo)](https://pkg.go.dev/github.com/blacktop/arm64-cgo/disassemble)
![GitHub](https://img.shields.io/github/license/blacktop/arm64-cgo?color=blue)

> Golang bindings for the Binary Ninja [Arm64 Disassembler](https://github.com/Vector35/arch-arm64) with support for **ARMv9A (2021-06 spec)**

## Getting Started

```bash
go get github.com/blacktop/arm64-cgo
```

```go
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/blacktop/arm64-cgo/disassemble"
)

func main() {

	var startAddr uint64
	var instrValue uint32
	var results [1024]byte

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(data)

	for {
		err = binary.Read(r, binary.LittleEndian, &instrValue)

		if err == io.EOF {
			break
		}

		instruction, err := disassemble.Disassemble(startAddr, instrValue, &results)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#08x:  %s\t%s\n",
			uint64(startAddr),
			disassemble.GetOpCodeByteString(instrValue),
			instruction)

		startAddr += uint64(binary.Size(uint32(0)))
	}
}
```

## MachO disassembler CLI `disass`

### Install

```bash
go install github.com/blacktop/arm64-cgo/cmd/disass@latest
```

With [Homebrew](https://brew.sh)

```bash
brew install blacktop/tap/disass
```

> OR download from [Releases](https://github.com/blacktop/arm64-cgo/releases/latest)

### Usage

```bash
disass hello-mte --vaddr 0x100007f1c
```

```s
_main:
0x100007ef0:  7f 23 03 d5	pacibsp
0x100007ef4:  ff 43 01 d1	sub	sp, sp, #0x50
0x100007ef8:  fd 7b 04 a9	stp	x29, x30, [sp, #0x40]
0x100007efc:  fd 03 01 91	add	x29, sp, #0x40
0x100007f00:  08 00 00 b0	adrp	x8, 0x100008000
0x100007f04:  08 09 40 f9	ldr	x8, [x8, #0x10]
0x100007f08:  08 01 40 f9	ldr	x8, [x8]
0x100007f0c:  a8 83 1f f8	stur	x8, [x29, #-0x8]
0x100007f10:  08 00 80 d2	mov	x8, #0
0x100007f14:  e8 13 c8 9a	irg	x8, sp, x8
0x100007f18:  09 01 82 91	addg	x9, x8, #0x20, #0
👉100007f1c:  29 09 20 d9	stg	x9, [x9]
0x100007f20:  08 05 81 91	addg	x8, x8, #0x10, #0x1
0x100007f24:  08 09 20 d9	stg	x8, [x8]
0x100007f28:  0a 00 80 52	mov	w10, #0
0x100007f2c:  2a 01 00 b9	str	w10, [x9]
0x100007f30:  e8 07 00 f9	str	x8, [sp, #0x8]
0x100007f34:  ca ff ff 97	bl	_test
0x100007f38:  e8 07 40 f9	ldr	x8, [sp, #0x8]
0x100007f3c:  00 01 00 39	strb	w0, [x8]
0x100007f40:  0a 01 40 39	ldrb	w10, [x8]
0x100007f44:  e9 03 0a aa	mov	x9, x10
0x100007f48:  eb 03 00 91	mov	x11, sp
0x100007f4c:  69 01 00 f9	str	x9, [x11]
0x100007f50:  00 00 00 90	adrp	x0, 0x100007000
0x100007f54:  00 d0 3e 91	add	x0, x0, #0xfb4 ; "%c\n"
0x100007f58:  13 00 00 94	bl	0x100007fa4
0x100007f5c:  ff 1b a0 d9	st2g	sp, [sp, #0x10]
0x100007f60:  08 00 00 b0	adrp	x8, 0x100008000
0x100007f64:  08 09 40 f9	ldr	x8, [x8, #0x10]
0x100007f68:  08 01 40 f9	ldr	x8, [x8]
0x100007f6c:  a9 83 5f f8	ldur	x9, [x29, #-0x8]
0x100007f70:  08 01 09 eb	subs	x8, x8, x9
0x100007f74:  c1 00 00 54	b.ne	0x100007f8c
0x100007f78:  01 00 00 14	b	0x100007f7c
0x100007f7c:  00 00 80 52	mov	w0, #0
0x100007f80:  fd 7b 44 a9	ldp	x29, x30, [sp, #0x40]
0x100007f84:  ff 43 01 91	add	sp, sp, #0x50
0x100007f88:  ff 0f 5f d6	retab
0x100007f8c:  02 00 00 94	bl	0x100007f94
0x100007f90:  20 00 20 d4	brk	#0x1
```

### Make it pretty 💄

```bash
brew install bat
```

```bash
disass hello-mte --vaddr 0x100007f1c \
	| bat -l s --tabs 0 -p --theme Nord --wrap=never --pager "less -S"
```

### Output as JSON

```bash
disass hello-mte --symbol _main --json | jq .
```

```json
[
   {
      "address": 4294999644,
      "raw": 3573752703,
      "encoding": "PACIBSP_HI_hints",
      "operation": "pacibsp",
      "disassembly": "pacibsp"
   },
   {
      "address": 4294999648,
      "raw": 3506619391,
      "encoding": "SUB_64_addsub_imm",
      "operation": "sub",
      "operands": [
         {
            "class": "REG",
            "registers": [
               "sp"
            ]
         },
         {
            "class": "REG",
            "registers": [
               "sp"
            ]
         },
         {
            "class": "IMM64",
            "immediate": 176
         }
      ],
      "disassembly": "sub\tsp, sp, #0xb0"
   },
   {
      "address": 4294999652,
      "raw": 2836036605,
      "encoding": "STP_64_ldstpair_off",
      "operation": "stp",
      "operands": [
         {
            "class": "REG",
            "registers": [
               "x29"
            ]
         },
         {
            "class": "REG",
            "registers": [
               "x30"
            ]
         },
         {
            "class": "MEM_OFFSET",
            "registers": [
               "sp"
            ],
            "immediate": 160,
            "signed_imm": true
         }
      ],
      "disassembly": "stp\tx29, x30, [sp, #0xa0]"
   },
   <SNIP>
  {
      "address": 4294999788,
      "raw": 3558866976,
      "encoding": "BRK_EX_exception",
      "operation": "brk",
      "operands": [
         {
            "class": "IMM32",
            "immediate": 1
         }
      ],
      "disassembly": "brk\t#0x1"
   }
]
```

## License

MIT Copyright (c) 2021-2023 blacktop

Apache License, Version 2.0 Copyright 2020-2021 Vector 35 Inc.
