package disassemble

// Op is a fixed-size operand that avoids heap allocations.
// All slice fields from Operand are replaced with fixed arrays.
type Op struct {
	Class          operandClass
	ArrSpec        arrangementSpec
	Registers      [MAX_REGISTERS]Register
	NumRegisters   uint8
	Condition      condition
	ImplSpec       [5]byte
	HasImplSpec    bool
	SysReg         SystemReg
	LaneUsed       bool
	Lane           uint32
	Immediate      uint64
	ShiftType      shiftType
	ShiftValueUsed bool
	ShiftValue     uint32
	Extend         shiftType
	SignedImm      bool
	PredQual       byte
	MulVl          bool
	Tile           uint16
	Slice          sliceType
	Name           [MAX_NAME]byte
	NameLen        uint8
}

// NameString returns the name as a Go string.
func (o *Op) NameString() string {
	return string(o.Name[:o.NameLen])
}

// RegSlice returns the registers as a slice (allocates).
func (o *Op) RegSlice() []Register {
	return append([]Register(nil), o.Registers[:o.NumRegisters]...)
}

// GetImmediate applies shift to the immediate value.
func (o *Op) GetImmediate() uint64 {
	imm := o.Immediate
	if o.ShiftValueUsed {
		switch o.ShiftType {
		case SHIFT_TYPE_LSL:
			imm = imm << uint64(o.ShiftValue)
		case SHIFT_TYPE_LSR:
			imm = imm >> uint64(o.ShiftValue)
		case SHIFT_TYPE_MSL:
			imm = imm<<uint64(o.ShiftValue) |
				^uint64(1)>>(64-o.ShiftValue)
		}
	}
	return imm
}

// ToOperand converts to the heap-allocated Operand type.
func (o *Op) ToOperand() Operand {
	op := Operand{
		Class:          o.Class,
		ArrSpec:        o.ArrSpec,
		Condition:      o.Condition,
		SysReg:         o.SysReg,
		LaneUsed:       o.LaneUsed,
		Lane:           o.Lane,
		Immediate:      o.Immediate,
		ShiftType:      o.ShiftType,
		ShiftValueUsed: o.ShiftValueUsed,
		ShiftValue:     o.ShiftValue,
		Extend:         o.Extend,
		SignedImm:      o.SignedImm,
		PredQual:       o.PredQual,
		MulVl:          o.MulVl,
		Tile:           o.Tile,
		Slice:          o.Slice,
		Name:           o.NameString(),
	}
	if o.NumRegisters > 0 {
		op.Registers = make([]Register, o.NumRegisters)
		copy(op.Registers, o.Registers[:o.NumRegisters])
	}
	if o.HasImplSpec {
		op.ImplSpec = make([]byte, 5)
		copy(op.ImplSpec, o.ImplSpec[:])
	}
	return op
}

// Inst is a fixed-size instruction that can live on the stack.
type Inst struct {
	Address   uint64
	Raw       uint32
	Encoding  encoding
	Operation Operation
	SetFlags  FlagEffect
	Operands  [MAX_OPERANDS]Op
	NumOps    uint8
}

// Disassemble lazily formats the instruction as a string.
// Only callers that need the string pay the cost.
func (inst *Inst) Disassemble() (string, error) {
	var buf [1024]byte
	return Disassemble(inst.Address, inst.Raw, &buf)
}

// ToInstruction converts to the heap-allocated Instruction type.
func (inst *Inst) ToInstruction() *Instruction {
	i := &Instruction{
		Address:   inst.Address,
		Raw:       inst.Raw,
		Encoding:  inst.Encoding,
		Operation: inst.Operation,
		SetFlags:  inst.SetFlags,
	}
	if inst.NumOps > 0 {
		i.Operands = make([]Operand, inst.NumOps)
		for idx := uint8(0); idx < inst.NumOps; idx++ {
			i.Operands[idx] = inst.Operands[idx].ToOperand()
		}
	}
	return i
}
