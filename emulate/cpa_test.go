package emulate

import (
	"errors"
	"testing"

	"github.com/blacktop/arm64-cgo/emulate/core"
)

// FEAT_CPA integer forms execute as plain arithmetic (as if pointer checking is disabled).
func TestExecuteCPA(t *testing.T) {
	tests := []struct {
		name  string
		instr uint32
		setup func(e *Engine)
		check func(t *testing.T, e *Engine)
	}{
		{
			name:  "addpt x8, x27, x21, lsl #4",
			instr: 0x9a153368,
			setup: func(e *Engine) { e.SetRegister(27, 0x1000); e.SetRegister(21, 0x3) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 8, 0x1030) },
		},
		{
			name:  "addpt x8, sp, x5",
			instr: 0x9a0523e8,
			setup: func(e *Engine) { e.SetSP(0x7000); e.SetRegister(5, 0x10) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 8, 0x7010) },
		},
		{
			name:  "addpt x1, x2, xzr",
			instr: 0x9a1f2041,
			setup: func(e *Engine) { e.SetRegister(2, 0x1234) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 1, 0x1234) },
		},
		{
			name:  "subpt x17, x13, x3, lsl #7",
			instr: 0xda033db1,
			setup: func(e *Engine) { e.SetRegister(13, 0x1000); e.SetRegister(3, 0x2) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 17, 0x1000-0x100) },
		},
		{
			name:  "subpt sp, sp, x5",
			instr: 0xda0523ff,
			setup: func(e *Engine) { e.SetSP(0x7000); e.SetRegister(5, 0x10) },
			check: func(t *testing.T, e *Engine) {
				if got := e.GetSP(); got != 0x6ff0 {
					t.Errorf("SP = %#x, want %#x", got, 0x6ff0)
				}
			},
		},
		{
			name:  "maddpt x3, x2, x1, x4",
			instr: 0x9b611043,
			setup: func(e *Engine) { e.SetRegister(2, 6); e.SetRegister(1, 7); e.SetRegister(4, 100) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 3, 142) },
		},
		{
			name:  "msubpt x3, x2, x1, x4",
			instr: 0x9b619043,
			setup: func(e *Engine) { e.SetRegister(2, 6); e.SetRegister(1, 7); e.SetRegister(4, 100) },
			check: func(t *testing.T, e *Engine) { expectX(t, e, 3, 58) },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewEngine()
			tt.setup(engine)
			if err := engine.ExecuteInstruction(0x1000, tt.instr); err != nil {
				t.Fatalf("ExecuteInstruction failed: %v", err)
			}
			tt.check(t, engine)
			if engine.GetPC() != 0x1004 {
				t.Errorf("PC = %#x, want 0x1004", engine.GetPC())
			}
		})
	}
}

// SVE CPA forms share the ADDPT/SUBPT operation but must not execute as scalar arithmetic.
// They surface as ErrUnsupportedFeature so HookUnimplementedInstruction can intercept them.
func TestExecuteCPA_SVEFormsUnsupported(t *testing.T) {
	tests := []struct {
		name  string
		instr uint32
	}{
		{"addpt z3.d, z2.d, z1.d", 0x04e10843},
		{"subpt z3.d, p1/m, z3.d, z2.d", 0x04c50443},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := NewEngine()
			err := engine.ExecuteInstruction(0x1000, tt.instr)
			if !errors.Is(err, core.ErrUnsupportedFeature) {
				t.Fatalf("error = %v, want ErrUnsupportedFeature", err)
			}

			hooked := NewEngine()
			var seen []string
			hook := core.UnimplementedInstructionHook(func(_ core.State, info core.InstructionInfo, _ error) core.HookResult {
				seen = append(seen, info.Mnemonic)
				return core.HookResult{SkipInstruction: true}
			})
			if err := hooked.AddHook(core.HookUnimplementedInstruction, hook); err != nil {
				t.Fatalf("AddHook failed: %v", err)
			}
			if err := hooked.ExecuteInstruction(0x1000, tt.instr); err != nil {
				t.Fatalf("hooked ExecuteInstruction failed: %v", err)
			}
			if len(seen) != 1 {
				t.Fatalf("unimplemented hook fired %d times, want 1", len(seen))
			}
			if hooked.GetPC() != 0x1004 {
				t.Errorf("PC = %#x, want 0x1004 after skip", hooked.GetPC())
			}
		})
	}
}

// The scalar CPA mnemonics must be discoverable through the public support APIs.
func TestExecuteCPA_Advertised(t *testing.T) {
	engine := NewEngine()
	listed := make(map[string]bool)
	for _, mnemonic := range engine.ListSupportedInstructions() {
		listed[mnemonic] = true
	}
	for _, mnemonic := range []string{"ADDPT", "SUBPT", "MADDPT", "MSUBPT"} {
		if !listed[mnemonic] {
			t.Errorf("%s missing from ListSupportedInstructions", mnemonic)
		}
		executor, found := engine.registry.Get(mnemonic)
		if !found {
			t.Errorf("%s has no registered executor", mnemonic)
			continue
		}
		if !executor.Supports(mnemonic) {
			t.Errorf("executor for %s does not report Supports(%q)", mnemonic, mnemonic)
		}
	}
}

func expectX(t *testing.T, e *Engine, reg int, want uint64) {
	t.Helper()
	if got := e.GetRegister(reg); got != want {
		t.Errorf("X%d = %#x, want %#x", reg, got, want)
	}
}
