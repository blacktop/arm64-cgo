package disassemble

//go:generate stringer -type=arrangementSpec,operandClass,condition,shiftType,returnCode,Group  -output disassemble_string.go

/*
#cgo CFLAGS: -I${SRCDIR}

#include <stdlib.h>

#include "decode.h"
#include "format.h"

int disassemble(uint64_t addr, uint32_t instrValue, int len, char *result)
{
	Instruction instr;
	memset(&instr, 0, sizeof(instr));

	int rc = aarch64_decompose(instrValue, &instr, addr);
	if(rc != DECODE_STATUS_OK)
		return rc;

	return aarch64_disassemble(&instr, result, 1024);
}

int aarch64_decompose(uint32_t instructionValue, Instruction *instr, uint64_t address);
int aarch64_disassemble(Instruction *instruction, char *buf, size_t buf_sz);
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"unsafe"
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
	NONE operandClass = iota
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

type returnCode int

// decode return values
const (
	DECODE_STATUS_ERROR_OPERANDS     returnCode = -9
	DECODE_STATUS_ASSERT_FAILED      returnCode = -8 // failed an assert
	DECODE_STATUS_UNREACHABLE        returnCode = -7 // ran into pcode Unreachable()
	DECODE_STATUS_LOST               returnCode = -6 // descended past checks, ie: "SEE encoding_up_higher"
	DECODE_STATUS_END_OF_INSTRUCTION returnCode = -5 // spec decode EndOfInstruction(), instruction executes as NOP
	DECODE_STATUS_UNDEFINED          returnCode = -4 // spec says this encoding is undefined, often due to a disallowed field or a missing feature, eg: "if !HaveBF16Ext() then UNDEFINED;"
	DECODE_STATUS_UNALLOCATED        returnCode = -3 // spec says this space is unallocated, eg: UNALLOCATED_10_branch_reg
	DECODE_STATUS_UNMATCHED          returnCode = -2 // decoding logic fell through the spec's checks
	DECODE_STATUS_RESERVED           returnCode = -1 // spec says this space is reserved, eg: RESERVED_36_asisdsame
)

// FailureCode - these get returned by the disassemble_instruction() function
const (
	SUCCESS_OK returnCode = iota
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

// Operand is an arm64 instruction operand object
type Operand struct {
	Class     operandClass    `json:"class,omitempty"`
	ArrSpec   arrangementSpec `json:"arr_spec,omitempty"`
	Registers []register      `json:"registers,omitempty"`

	Condition condition `json:"condition,omitempty"` // for class CONDITION

	ImplSpec []byte `json:"impl_spec,omitempty"` // for class IMPLEMENTATION_SPECIFIC

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

// MarshalJSON is the operand's custom JSON marshaler
func (o *Operand) MarshalJSON() ([]byte, error) {
	var regs []string
	for _, r := range o.Registers {
		regs = append(regs, r.String())
	}
	return json.Marshal(struct {
		Class          string   `json:"class,omitempty"`
		ArrSpec        string   `json:"arr_spec,omitempty"`
		Registers      []string `json:"registers,omitempty"`
		Condition      string   `json:"condition,omitempty"` // for class CONDITION
		ImplSpec       []byte   `json:"impl_spec,omitempty"` // for class IMPLEMENTATION_SPECIFIC
		SysReg         string   `json:"sys_reg,omitempty"`   // for class SYS_REG
		LaneUsed       bool     `json:"lane_used,omitempty"`
		Lane           uint32   `json:"lane,omitempty"`
		Immediate      int64    `json:"immediate,omitempty"` // TODO: is it dangerous casting this as a int64 without checking SignedImm first
		ShiftType      string   `json:"shift_type,omitempty"`
		ShiftValueUsed bool     `json:"shift_value_used,omitempty"`
		ShiftValue     uint32   `json:"shift_value,omitempty"`
		Extend         string   `json:"extend,omitempty"`
		SignedImm      bool     `json:"signed_imm,omitempty"`
		PredQual       byte     `json:"pred_qual,omitempty"` // predicate register qualifier ('z' or 'm')
		MulVl          bool     `json:"mul_vl,omitempty"`    // whether MEM_OFFSET has the offset "mul vl"
		Name           string   `json:"name,omitempty"`      // or class NAME
	}{
		Class:          o.Class.String(),
		ArrSpec:        o.ArrSpec.String(),
		Registers:      regs,
		Condition:      o.Condition.String(),
		ImplSpec:       o.ImplSpec,
		SysReg:         o.SysReg.String(),
		LaneUsed:       o.LaneUsed,
		Lane:           o.Lane,
		Immediate:      int64(o.Immediate),
		ShiftType:      o.ShiftType.String(),
		ShiftValueUsed: o.ShiftValueUsed,
		ShiftValue:     o.ShiftValue,
		Extend:         o.Extend.String(),
		SignedImm:      o.SignedImm,
		PredQual:       o.PredQual,
		MulVl:          o.MulVl,
		Name:           o.Name,
	})
}

// Instruction is an arm64 instruction object
type Instruction struct {
	Raw         uint32    `json:"raw,omitempty"`
	Encoding    encoding  `json:"encoding,omitempty"`
	Operation   operation `json:"operation,omitempty"`
	Operands    []Operand `json:"operands,omitempty"`
	SetFlags    bool      `json:"set_flags,omitempty"`
	Disassembly string    `json:"disassembly,omitempty"`
}

// OpCodes returns the instruction's opcodes as a string of hex bytes
func (i *Instruction) OpCodes() string {
	return GetOpCodeByteString(i.Raw)
}

func (i *Instruction) String() string {
	return i.Disassembly
}

// MarshalJSON is the instruction's custom JSON marshaler
func (i *Instruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Raw         uint32    `json:"raw,omitempty"`
		Encoding    string    `json:"encoding,omitempty"`
		Operation   string    `json:"operation,omitempty"`
		Operands    []Operand `json:"operands,omitempty"`
		SetFlags    bool      `json:"set_flags,omitempty"`
		Disassembly string    `json:"disassembly,omitempty"`
	}{
		Raw:         i.Raw,
		Encoding:    i.Encoding.String(),
		Operation:   i.Operation.String(),
		Operands:    i.Operands,
		SetFlags:    i.SetFlags,
		Disassembly: i.Disassembly,
	})
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
func Disassemble(addr uint64, instructionValue uint32, results *[1024]byte) (string, error) {

	if rc := C.disassemble(
		C.uint64_t(addr),                   // uint64_t addr
		C.uint32_t(instructionValue),       // uint32_t instrValue
		C.int(4),                           // int len
		(*C.char)(unsafe.Pointer(results)), // char *result
	); returnCode(rc) != SUCCESS_OK {
		return "", fmt.Errorf("failed to disassemble instruction %#x: %s", instructionValue, returnCode(rc).String())
	}

	return C.GoString((*C.char)(unsafe.Pointer(results))), nil
}

// Decompose decomposes an instruction
func Decompose(addr uint64, instructionValue uint32, results *[1024]byte) (*Instruction, error) {

	var err error
	var instruction C.Instruction

	if rc := C.aarch64_decompose(
		C.uint32_t(instructionValue), // uint32_t instructionValue
		&instruction,                 // Instruction *instr,
		C.uint64_t(addr),             // uint64_t address
	); returnCode(rc) != SUCCESS_OK {
		return nil, fmt.Errorf("failed to decompose instruction %#x: %s", instructionValue, returnCode(rc).String())
	}

	i := goInstruction(&instruction)

	i.Disassembly, err = Disassemble(addr, instructionValue, results)
	if err != nil {
		return nil, err
	}

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
			if !allZero(op.implspec) {
				for k, ispec := range op.implspec {
					i.Operands[idx].ImplSpec[k] = byte(ispec)
				}
			}
		}
	}

	return i
}

func allZero(s [5]C.uchar) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
