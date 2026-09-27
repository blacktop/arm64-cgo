package disassemble

/*
#cgo CFLAGS: -I${SRCDIR}

#include "decode.h"
*/
import "C"

import "sync"

// singleDecoders backs the one-shot DecomposeInto and Decompose wrappers,
// so each call reuses a C.Instruction instead of allocating one. Only
// cInstr is used, so pooled Decoders never hold batch buffers.
var singleDecoders = sync.Pool{New: func() any { return new(Decoder) }}

// Decoder reuses cgo-side decode buffers across calls so that the
// per-decode heap allocation is paid once, not on every call.
//
// Decoder is not safe for concurrent use by multiple goroutines.
// The zero value is ready to use.
type Decoder struct {
	cInstr C.Instruction
	batch  []C.Instruction
	status []DecodeStatus
}

// NewDecoder creates a reusable decoder. If batchCap is greater than 0,
// the batch buffer is preallocated to that capacity.
func NewDecoder(batchCap int) *Decoder {
	decoder := &Decoder{}
	if batchCap > 0 {
		decoder.batch = make([]C.Instruction, 0, batchCap)
	}
	return decoder
}

func (d *Decoder) ensureBatch(n int) {
	if cap(d.batch) < n {
		d.batch = make([]C.Instruction, n)
		return
	}
	d.batch = d.batch[:n]
}
