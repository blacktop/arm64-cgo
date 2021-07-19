package disassemble

//go:generate stringer -type=arrangementSpec,operandClass,condition,shiftType,failureCode,Group  -output disassemble_string.go

/*
#cgo CFLAGS: -I${SRCDIR}

#include "decode.h"
#include "format.h"

void disassemble(uint64_t addr, uint8_t *data, int len, char *result)
{
	Instruction instr;
	memset(&instr, 0, sizeof(instr));

	aarch64_decompose(*(uint32_t *)data, &instr, addr);

	aarch64_disassemble(&instr, result, 1024);
}

int aarch64_decompose(uint32_t instructionValue, Instruction *instr, uint64_t address);
int aarch64_disassemble(Instruction *instruction, char *buf, size_t buf_sz);
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	MAX_REGISTERS = 5
	MAX_NAME      = 16
	MAX_OPERANDS  = 5
)

type arrangementSpec uint32

/* these are used in lookup tables elsewhere, modify with caution */
const (
	ARRSPEC_NONE arrangementSpec = 0

	ARRSPEC_FULL = 1 /* 128-bit v-reg unsplit, eg: REG_V0_Q0 */

	/* 128 bit v-reg considered as... */
	ARRSPEC_2DOUBLES = 2 /* (.2d) two 64-bit double-precision: REG_V0_D1, REG_V0_D0 */
	ARRSPEC_4SINGLES = 3 /* (.4s) four 32-bit single-precision: REG_V0_S3, REG_V0_S2, REG_V0_S1, REG_V0_S0 */
	ARRSPEC_8HALVES  = 4 /* (.8h) eight 16-bit half-precision: REG_V0_H7, REG_V0_H6, (..., REG_V0_H0 */
	ARRSPEC_16BYTES  = 5 /* (.16b) sixteen 8-bit values: REG_V0_B15, REG_V0_B14, (..., REG_V0_B01 */

	/* low 64-bit of v-reg considered as... */
	ARRSPEC_1DOUBLE  = 6 /* (.d) one 64-bit double-precision: REG_V0_D0 */
	ARRSPEC_2SINGLES = 7 /* (.2s) two 32-bit single-precision: REG_V0_S1, REG_V0_S0 */
	ARRSPEC_4HALVES  = 8 /* (.4h) four 16-bit half-precision: REG_V0_H3, REG_V0_H2, REG_V0_H1, REG_V0_H0 */
	ARRSPEC_8BYTES   = 9 /* (.8b) eight 8-bit values: REG_V0_B7, REG_V0_B6, (..., REG_V0_B0 */

	/* low 32-bit of v-reg considered as... */
	ARRSPEC_1SINGLE = 10 /* (.s) one 32-bit single-precision: REG_V0_S0 */
	ARRSPEC_2HALVES = 11 /* (.2h) two 16-bit half-precision: REG_V0_H1, REG_V0_H0 */
	ARRSPEC_4BYTES  = 12 /* (.4b) four 8-bit values: REG_V0_B3, REG_V0_B2, REG_V0_B1, REG_V0_B0 */

	/* low 16-bit of v-reg considered as... */
	ARRSPEC_1HALF = 13 /* (.h) one 16-bit half-precision: REG_V0_H0 */

	/* low 8-bit of v-reg considered as... */
	ARRSPEC_1BYTE = 14 /* (.b) one 8-bit byte: REG_V0_B0 */
)

type operandClass uint32

const (
	NONE operandClass = 0
	IMM32
	IMM64
	FIMM32
	STR_IMM
	REG
	MULTI_REG
	SYS_REG
	MEM_REG
	MEM_PRE_IDX
	MEM_POST_IDX
	MEM_OFFSET
	MEM_EXTENDED
	LABEL
	CONDITION
	NAME
	IMPLEMENTATION_SPECIFIC
)

type condition uint32

const (
	COND_EQ condition = iota
	COND_NE
	COND_CS
	COND_CC
	COND_MI
	COND_PL
	COND_VS
	COND_VC
	COND_HI
	COND_LS
	COND_GE
	COND_LT
	COND_GT
	COND_LE
	COND_AL
	COND_NV
	END_CONDITION
)

type shiftType uint32

const (
	SHIFT_TYPE_NONE shiftType = iota
	SHIFT_TYPE_LSL
	SHIFT_TYPE_LSR
	SHIFT_TYPE_ASR
	SHIFT_TYPE_ROR
	SHIFT_TYPE_UXTW
	SHIFT_TYPE_SXTW
	SHIFT_TYPE_SXTX
	SHIFT_TYPE_UXTX
	SHIFT_TYPE_SXTB
	SHIFT_TYPE_SXTH
	SHIFT_TYPE_UXTH
	SHIFT_TYPE_UXTB
	SHIFT_TYPE_MSL
	SHIFT_TYPE_END
)

type failureCode uint32

const (
	DISASM_SUCCESS failureCode = iota
	INVALID_ARGUMENTS
	FAILED_TO_DISASSEMBLE_OPERAND
	FAILED_TO_DISASSEMBLE_OPERATION
	FAILED_TO_DISASSEMBLE_REGISTER
	FAILED_TO_DECODE_INSTRUCTION
	OUTPUT_BUFFER_TOO_SMALL
	OPERAND_IS_NOT_REGISTER
	NOT_MEMORY_OPERAND
)

type Group uint32

const (
	GROUP_UNALLOCATED Group = iota
	GROUP_DATA_PROCESSING_IMM
	GROUP_BRANCH_EXCEPTION_SYSTEM
	GROUP_LOAD_STORE
	GROUP_DATA_PROCESSING_REG
	GROUP_DATA_PROCESSING_SIMD
	GROUP_DATA_PROCESSING_SIMD2
	END_GROUP
)

type Operand struct {
	Class     operandClass    `json:"class,omitempty"`
	ArrSpec   arrangementSpec `json:"arr_spec,omitempty"`
	Registers []register      `json:"registers,omitempty"`

	Condition condition `json:"condition,omitempty"` // for class CONDITION

	ImplSpec [MAX_REGISTERS]byte `json:"impl_spec,omitempty"` // for class IMPLEMENTATION_SPECIFIC

	SysReg systemReg `json:"sys_reg,omitempty"` // for class SYS_REG

	LaneUsed       bool      `json:"lane_used,omitempty"`
	Lane           uint32    `json:"lane,omitempty"`
	Immediate      uint64    `json:"immediate,omitempty"`
	ShiftType      shiftType `json:"shift_type,omitempty"`
	ShiftValueUsed bool      `json:"shift_value_used,omitempty"`
	ShiftValue     uint32    `json:"shift_value,omitempty"`
	Extend         shiftType `json:"extend,omitempty"`
	SignedImm      bool      `json:"signed_imm,omitempty"`
	PredQual       byte      `json:"pred_qual,omitempty"` // predicate register qualifier ('z' or 'm')
	MulVl          bool      `json:"mul_vl,omitempty"`    // whether MEM_OFFSET has the offset "mul vl"

	Name string `json:"name,omitempty"` // or class NAME
}

type Instruction struct {
	Raw         uint32    `json:"raw,omitempty"`
	Encoding    encoding  `json:"encoding,omitempty"`
	Operation   operation `json:"operation,omitempty"`
	Operands    []Operand `json:"operands,omitempty"`
	SetFlags    bool      `json:"set_flags,omitempty"`
	Disassembly string    `json:"disassembly,omitempty"`
}

func (i *Instruction) OpCodes() string {
	return GetOpCodeByteString(i.Raw)
}

func (i *Instruction) String() string {
	return i.Disassembly
}

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
func Disassemble(addr uint64, instruction uint32) (string, error) {

	var results []byte
	cResults := (*C.char)(C.CBytes(results))

	instrBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(instrBytes, instruction)

	C.disassemble(
		C.uint64_t(addr),                   // uint64_t addr
		(*C.uint8_t)(C.CBytes(instrBytes)), // uint8_t *data
		C.int(4),                           // int len
		cResults)                           // char *result

	return C.GoString(cResults), nil
}

// Decompose decomposes an instruction
func Decompose(addr uint64, instructionValue uint32) (*Instruction, error) {

	var instr []byte
	instruction := (*C.Instruction)(C.CBytes(instr))

	var results []byte
	cResults := (*C.char)(C.CBytes(results))

	instrBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(instrBytes, instructionValue)

	C.aarch64_decompose(*(*C.uint32_t)(C.CBytes(instrBytes)), instruction, C.uint64_t(addr))

	i := goInstruction(instruction)

	if ret := C.aarch64_disassemble(instruction, cResults, C.size_t(1024)); ret != 0 {
		return nil, fmt.Errorf("failed to disassemble instruction %#x: %s", instructionValue, failureCode(ret))
	}

	i.Disassembly = C.GoString(cResults)

	return i, nil
}

// goInstruction converts the cgo version into Go vesrion
func goInstruction(instr *C.Instruction) *Instruction {

	i := &Instruction{
		Raw:       uint32(instr.insword),
		Encoding:  encoding(instr.encoding),
		SetFlags:  bool(instr.setflags),
		Operation: operation(instr.operation),
	}

	for idx, op := range instr.operands {
		if op.operandClass != C.OperandClass(NONE) {
			i.Operands = append(i.Operands, Operand{
				Class:          operandClass(op.operandClass),
				ArrSpec:        arrangementSpec(op.arrSpec),
				Condition:      condition(op.cond),
				SysReg:         systemReg(op.sysreg),
				LaneUsed:       bool(op.laneUsed),
				Lane:           uint32(op.lane),
				Immediate:      uint64(op.immediate),
				ShiftType:      shiftType(op.shiftType),
				ShiftValueUsed: bool(op.shiftValueUsed),
				Extend:         shiftType(op.extend),
				SignedImm:      bool(op.signedImm),
				PredQual:       byte(op.pred_qual),
				MulVl:          bool(op.mul_vl),
				Name:           C.GoString(&op.name[0]),
			})
			for _, reg := range op.reg {
				if reg != C.Register(REG_NONE) {
					i.Operands[idx].Registers = append(i.Operands[idx].Registers, register(reg))
				}
			}
			for k, ispec := range op.implspec {
				i.Operands[idx].ImplSpec[k] = byte(ispec)
			}
		}
	}

	return i
}