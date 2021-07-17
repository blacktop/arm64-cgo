package main

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
	"io"

	"github.com/blacktop/go-macho"
)

func getOpCodeByteString(opcode uint32) string {
	op := new(bytes.Buffer)
	err := binary.Write(op, binary.LittleEndian, opcode)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("% x", op.Bytes())
}

func main() {
	m, err := macho.Open("../../Proteas/hello-mte/hello-mte")
	if err != nil {
		panic(err)
	}

	symAddr, err := m.FindSymbolAddress("_test")
	if err != nil {
		panic(err)
	}

	fn, err := m.GetFunctionForVMAddr(symAddr)
	if err != nil {
		panic(err)
	}

	data, err := m.GetFunctionData(fn)
	if err != nil {
		panic(err)
	}

	var instrValue uint32
	r := bytes.NewReader(data)

	fmt.Println("_test:")

	for {
		addr, _ := r.Seek(0, io.SeekCurrent)

		err = binary.Read(r, binary.LittleEndian, &instrValue)

		if err == io.EOF {
			break
		}

		symAddr += uint64(addr)

		instruction, err := Disassemble(symAddr, instrValue, true)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#08x:  %s\t%s\n", uint64(symAddr), getOpCodeByteString(instrValue), instruction)
	}
}

func Disassemble(addr uint64, instruction uint32, verbose bool) (string, error) {

	instrBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(instrBytes, instruction)

	var results []byte
	resultsPointer := C.CBytes(results)
	cResults := (*C.char)(resultsPointer)

	dataPointer := C.CBytes(instrBytes)
	cData := (*C.uint8_t)(dataPointer)

	C.disassemble(C.uint64_t(addr), cData, C.int(4), cResults, C.bool(verbose))

	return C.GoString(cResults), nil
}
