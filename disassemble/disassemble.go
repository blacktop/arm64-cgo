package disassemble

/*
#cgo CFLAGS: -I${SRCDIR}

#include "decode.h"
#include "format.h"

void disassemble(uint64_t addr, uint8_t *data, int len, char *result, bool verbose)
{
	Instruction instr;
	memset(&instr, 0, sizeof(instr));

	aarch64_decompose(*(uint32_t *)data, &instr, addr);

	aarch64_disassemble(&instr, result, 1024);
}

*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// GetOpCodeByteString returns the opcodes as a string of hex bytes
func GetOpCodeByteString(opcode uint32) string {
	op := new(bytes.Buffer)
	err := binary.Write(op, binary.LittleEndian, opcode)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("% x", op.Bytes())
}

// Disassemble disassembles an instruction
func Disassemble(addr uint64, instruction uint32, verbose bool) (string, error) {

	var results []byte
	cResults := (*C.char)(C.CBytes(results))

	instrBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(instrBytes, instruction)

	C.disassemble(
		C.uint64_t(addr),                   // uint64_t addr
		(*C.uint8_t)(C.CBytes(instrBytes)), // uint8_t *data
		C.int(4),                           // int len
		cResults,                           // char *result
		C.bool(verbose))                    // bool verbose

	return C.GoString(cResults), nil
}
