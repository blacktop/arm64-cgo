package emulate

import "testing"

// Encodings from llvm-mc (-mattr=+pauth-lr). Immediate forms encode the label as a negative,
// word-scaled offset in imm16 bits 20:5; register forms carry the register in the low bits.
const (
	pacibsppc     = 0xdac1a7fe
	paciasppc     = 0xdac1a3fe
	pacnbibsppc   = 0xdac187fe
	pacm          = 0xd50324ff
	pacia171615   = 0xdac18bfe
	autia171615   = 0xdac1bbfe
	retabsppcBase = 0x5520001f
	retaasppcBase = 0x5500001f
	autibsppcBase = 0xf3a0001f
	autiasppcBase = 0xf380001f
	autiasppcrX3  = 0xdac1907e
	autiasppcrXZR = 0xdac193fe
	retabsppcrX3  = 0xd65f0fe3
)

const (
	prologAddr = 0x1000
	epilogAddr = 0x1040
	stackPtr   = 0x7fff_0000
	returnAddr = 0x0000_4444_8888
)

func labelForm(base uint32, from, label uint64) uint32 {
	return base | uint32((from-label)/4)<<5
}

func newPAuthLREngine(t *testing.T) *Engine {
	t.Helper()
	engine := NewEngine()
	engine.SetSP(stackPtr)
	engine.SetRegister(30, returnAddr)
	return engine
}

func exec(t *testing.T, engine *Engine, pc uint64, instr uint32) {
	t.Helper()
	if err := engine.ExecuteInstruction(pc, instr); err != nil {
		t.Fatalf("ExecuteInstruction(%#x, %#x) failed: %v", pc, instr, err)
	}
}

// The iOS 27 SDK ARM64_STACK_PROLOG_ENH / ARM64_STACK_EPILOG_ENH pair: pacibsppc in the
// prologue, retabsppc <label of the prologue> in the epilogue.
func TestPAuthLR_PrologEpilogRoundTrip(t *testing.T) {
	engine := newPAuthLREngine(t)

	exec(t, engine, prologAddr, pacibsppc)
	if lr := engine.GetRegister(30); lr == returnAddr || lr&0x00FFFFFFFFFFFFFF != returnAddr {
		t.Fatalf("LR = %#x after pacibsppc, want signed copy of %#x", lr, returnAddr)
	}
	if engine.GetPC() != prologAddr+4 {
		t.Fatalf("PC = %#x after pacibsppc, want %#x", engine.GetPC(), prologAddr+4)
	}

	signedLR := engine.GetRegister(30)
	exec(t, engine, epilogAddr, labelForm(retabsppcBase, epilogAddr, prologAddr))
	if engine.GetPC() != returnAddr {
		t.Errorf("PC = %#x after retabsppc, want %#x", engine.GetPC(), returnAddr)
	}
	if engine.GetRegister(30) != signedLR {
		t.Errorf("LR = %#x after retabsppc, want unchanged signed value %#x", engine.GetRegister(30), signedLR)
	}
}

// A mismatched authenticated return branches to the poisoned target but still leaves LR alone.
func TestPAuthLR_ReturnMismatchDoesNotWriteLR(t *testing.T) {
	engine := newPAuthLREngine(t)
	exec(t, engine, prologAddr, pacibsppc)
	signedLR := engine.GetRegister(30)
	exec(t, engine, epilogAddr, labelForm(retaasppcBase, epilogAddr, prologAddr))
	if engine.GetPC() != 0 {
		t.Errorf("PC = %#x after key-mismatched retaasppc, want 0", engine.GetPC())
	}
	if engine.GetRegister(30) != signedLR {
		t.Errorf("LR = %#x, want unchanged signed value %#x", engine.GetRegister(30), signedLR)
	}
}

func TestPAuthLR_MismatchPoisonsLR(t *testing.T) {
	tests := []struct {
		name  string
		setup func(e *Engine)
		instr uint32
	}{
		{"wrong label", func(*Engine) {}, labelForm(autibsppcBase, epilogAddr, prologAddr+8)},
		{"wrong key", func(*Engine) {}, labelForm(autiasppcBase, epilogAddr, prologAddr)},
		{"wrong SP", func(e *Engine) { e.SetSP(stackPtr + 0x10) }, labelForm(autibsppcBase, epilogAddr, prologAddr)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := newPAuthLREngine(t)
			exec(t, engine, prologAddr, pacibsppc)
			tt.setup(engine)
			exec(t, engine, epilogAddr, tt.instr)
			if lr := engine.GetRegister(30); lr != 0 {
				t.Errorf("LR = %#x, want 0 (poisoned)", lr)
			}
		})
	}
}

func TestPAuthLR_RegisterModifierForms(t *testing.T) {
	t.Run("autiasppcr x3 matches paciasppc address", func(t *testing.T) {
		engine := newPAuthLREngine(t)
		exec(t, engine, prologAddr, paciasppc)
		engine.SetRegister(3, prologAddr)
		exec(t, engine, epilogAddr, autiasppcrX3)
		if engine.GetRegister(30) != returnAddr {
			t.Errorf("LR = %#x, want %#x", engine.GetRegister(30), returnAddr)
		}
	})
	t.Run("autiasppcr xzr uses zero modifier", func(t *testing.T) {
		engine := newPAuthLREngine(t)
		exec(t, engine, prologAddr, paciasppc)
		exec(t, engine, epilogAddr, autiasppcrXZR)
		if lr := engine.GetRegister(30); lr != 0 {
			t.Errorf("LR = %#x, want 0 (poisoned: PC modifier was %#x, not zero)", lr, prologAddr)
		}
	})
	t.Run("retabsppcr x3 returns through LR", func(t *testing.T) {
		engine := newPAuthLREngine(t)
		exec(t, engine, prologAddr, pacnbibsppc)
		engine.SetRegister(3, prologAddr)
		signedLR := engine.GetRegister(30)
		exec(t, engine, epilogAddr, retabsppcrX3)
		if engine.GetPC() != returnAddr {
			t.Errorf("PC = %#x, want %#x", engine.GetPC(), returnAddr)
		}
		if engine.GetRegister(30) != signedLR {
			t.Errorf("LR = %#x, want unchanged signed value %#x", engine.GetRegister(30), signedLR)
		}
	})
}

func TestPAuthLR_X17Forms(t *testing.T) {
	engine := NewEngine()
	engine.SetRegister(17, 0x0000_1234_5678)
	engine.SetRegister(16, 0xaaaa)
	engine.SetRegister(15, 0x5555)
	exec(t, engine, prologAddr, pacia171615)
	if x17 := engine.GetRegister(17); x17 == 0x0000_1234_5678 {
		t.Fatalf("X17 = %#x after pacia171615, want a signed value", x17)
	}
	exec(t, engine, prologAddr+4, autia171615)
	if x17 := engine.GetRegister(17); x17 != 0x0000_1234_5678 {
		t.Errorf("X17 = %#x after autia171615, want %#x", x17, 0x0000_1234_5678)
	}
}

// Kernel addresses have bit 55 set; signing must not lose that and authentication must restore
// bits 63:56 from it, so a kernel return address survives the prologue/epilogue round trip.
func TestPAuthLR_HighHalfAddresses(t *testing.T) {
	const kernelLR uint64 = 0xffff_fe00_1234_5678
	t.Run("retabsppc branches to the full kernel address", func(t *testing.T) {
		engine := newPAuthLREngine(t)
		engine.SetRegister(30, kernelLR)
		exec(t, engine, prologAddr, pacibsppc)
		if lr := engine.GetRegister(30); lr&pointerHighHalfBit == 0 {
			t.Fatalf("LR = %#x after pacibsppc, bit 55 must survive signing", lr)
		}
		exec(t, engine, epilogAddr, labelForm(retabsppcBase, epilogAddr, prologAddr))
		if engine.GetPC() != kernelLR {
			t.Errorf("PC = %#x, want %#x", engine.GetPC(), kernelLR)
		}
	})
	t.Run("autibsppc restores the full kernel address", func(t *testing.T) {
		engine := newPAuthLREngine(t)
		engine.SetRegister(30, kernelLR)
		exec(t, engine, prologAddr, pacibsppc)
		exec(t, engine, epilogAddr, labelForm(autibsppcBase, epilogAddr, prologAddr))
		if lr := engine.GetRegister(30); lr != kernelLR {
			t.Errorf("LR = %#x, want %#x", lr, kernelLR)
		}
	})
	t.Run("autia171615 restores a kernel X17", func(t *testing.T) {
		engine := NewEngine()
		engine.SetRegister(17, kernelLR)
		exec(t, engine, prologAddr, pacia171615)
		exec(t, engine, prologAddr+4, autia171615)
		if x17 := engine.GetRegister(17); x17 != kernelLR {
			t.Errorf("X17 = %#x, want %#x", x17, kernelLR)
		}
	})
}

// A signed pointer whose payload was modified must not authenticate even with its tag intact.
func TestPAuthLR_TamperedPayloadRejected(t *testing.T) {
	engine := newPAuthLREngine(t)
	exec(t, engine, prologAddr, pacibsppc)
	engine.SetRegister(30, engine.GetRegister(30)^0x10)
	exec(t, engine, epilogAddr, labelForm(retabsppcBase, epilogAddr, prologAddr))
	if engine.GetPC() != 0 {
		t.Errorf("PC = %#x after return through tampered LR, want 0 (poisoned)", engine.GetPC())
	}
}

const pointerHighHalfBit = 1 << 55

func TestPAuthLR_PACMIsNop(t *testing.T) {
	engine := newPAuthLREngine(t)
	exec(t, engine, prologAddr, pacm)
	if engine.GetRegister(30) != returnAddr || engine.GetPC() != prologAddr+4 {
		t.Errorf("pacm changed state: LR=%#x PC=%#x", engine.GetRegister(30), engine.GetPC())
	}
}
