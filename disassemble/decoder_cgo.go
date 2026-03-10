package disassemble

/*
#cgo CFLAGS: -I${SRCDIR}

#include <string.h>
#include "decode.h"

static inline void zero_instruction(Instruction *instr)
{
	memset(instr, 0, sizeof(*instr));
}
*/
import "C"

// Decoder reuses cgo-side decode buffers across calls.
//
// Note: C.Instruction is still heap-allocated because of cgo pointer
// rules. The performance win comes from reusing these buffers instead
// of allocating fresh ones on every decode.
//
// Decoder is not safe for concurrent use by multiple goroutines.
// The zero value is ready to use.
type Decoder struct {
	cInstr C.Instruction
	batch  []C.Instruction
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
