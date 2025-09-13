package emulate

import (
	"errors"
	"sync/atomic"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// BoundsGuard provides opt-in execution bounds based on an address range
// while allowing calls to escape the range. It halts only when execution
// has returned to the original call frame (depth 0) and the PC has moved
// beyond the configured end address (or outside the range, if strict).
//
// Typical usage:
//
//	guard, _ := NewBoundsGuard(start, end, WithStrictRange(true))
//	cfg := DefaultEngineConfig()
//	cfg.ShouldHaltPreHandler = guard.Pre
//	cfg.ShouldHaltPostHandler = guard.Post
//	eng := NewEngineWithConfig(cfg)
//
// Start is inclusive; End is exclusive.
type BoundsGuard struct {
	start    uint64
	end      uint64
	strict   bool  // if true, halt when depth==0 and PC is outside [start,end)
	maxDepth int64 // clamp depth to this maximum (saturation)

	depth int64 // call depth relative to the starting frame (atomic)
}

// BoundOption configures a BoundsGuard.
type BoundOption func(*BoundsGuard)

// WithStrictRange toggles strict range checking at depth 0.
// When true (default), we halt if PC < Start or PC >= End at depth 0.
// When false, we only halt if PC >= End at depth 0.
func WithStrictRange(strict bool) BoundOption { return func(b *BoundsGuard) { b.strict = strict } }

// WithMaxDepth sets the saturation limit for depth tracking (default 1<<20).
func WithMaxDepth(max int64) BoundOption { return func(b *BoundsGuard) { b.maxDepth = max } }

// NewBoundsGuard constructs a guard with sensible defaults.
func NewBoundsGuard(start, end uint64, opts ...BoundOption) (*BoundsGuard, error) {
	if start >= end {
		return nil, errors.New("invalid range: start must be < end")
	}
	g := &BoundsGuard{start: start, end: end, strict: true, maxDepth: 1 << 20}
	for _, opt := range opts {
		opt(g)
	}
	if g.maxDepth <= 0 {
		g.maxDepth = 1 << 20
	}
	return g, nil
}

// Start returns the inclusive start address.
func (g *BoundsGuard) Start() uint64 { return g.start }

// End returns the exclusive end address.
func (g *BoundsGuard) End() uint64 { return g.end }

// Strict reports whether strict range checking is enabled.
func (g *BoundsGuard) Strict() bool { return g.strict }

// MaxDepth returns the saturation limit for depth tracking.
func (g *BoundsGuard) MaxDepth() int64 { return g.maxDepth }

// Reset clears the tracked depth back to 0.
func (g *BoundsGuard) Reset() { atomic.StoreInt64(&g.depth, 0) }

// Pre is intended for Engine.ShouldHaltPreHandler.
// It halts when we are back at depth 0 and PC has reached the end (or left the range if Strict).
func (g *BoundsGuard) Pre(state core.State, _ core.InstructionInfo) bool {
	pc := state.GetPC()
	if atomic.LoadInt64(&g.depth) == 0 {
		if g.strict {
			if pc < g.start || pc >= g.end {
				return true
			}
		} else if pc >= g.end {
			return true
		}
	}
	return false
}

// Post is intended for Engine.ShouldHaltPostHandler.
// It updates call depth by observing executed branch mnemonics.
func (g *BoundsGuard) Post(state core.State, info core.InstructionInfo) bool {
	// Track calls/returns via operation codes (no string allocations)
	switch info.Instruction.Operation {
	case disassemble.ARM64_BL, disassemble.ARM64_BLR,
		disassemble.ARM64_BLRAA, disassemble.ARM64_BLRAAZ,
		disassemble.ARM64_BLRAB, disassemble.ARM64_BLRABZ:
		// depth++ with saturation at MaxDepth
		for {
			d := atomic.LoadInt64(&g.depth)
			if d >= g.maxDepth {
				break
			}
			if atomic.CompareAndSwapInt64(&g.depth, d, d+1) {
				break
			}
		}

	case disassemble.ARM64_RET, disassemble.ARM64_RETAA, disassemble.ARM64_RETAB,
		disassemble.ARM64_ERET, disassemble.ARM64_ERETAA, disassemble.ARM64_ERETAB:
		// depth-- with floor at 0
		for {
			d := atomic.LoadInt64(&g.depth)
			if d == 0 {
				break
			}
			if atomic.CompareAndSwapInt64(&g.depth, d, d-1) {
				break
			}
		}

	case disassemble.ARM64_BR:
		// Treat BR X30 (LR) as an indirect return
		if len(info.Instruction.Operands) > 0 && len(info.Instruction.Operands[0].Registers) > 0 {
			reg := info.Instruction.Operands[0].Registers[0]
			mapped := core.MapRegister(reg)
			// Handle both mapped and direct enum checks for robustness.
			if mapped == 30 || reg == disassemble.REG_X30 {
				for {
					d := atomic.LoadInt64(&g.depth)
					if d == 0 {
						break
					}
					if atomic.CompareAndSwapInt64(&g.depth, d, d-1) {
						break
					}
				}
			}
		}
	}

	// Optional immediate halt after this instruction if we’re out of range at depth 0
	// (covers tail calls like B/BR out of range without waiting for next fetch)
	if atomic.LoadInt64(&g.depth) == 0 {
		pc := state.GetPC()
		if g.strict {
			if pc < g.start || pc >= g.end {
				// Re-check depth to avoid a rare race if depth changed concurrently
				if atomic.LoadInt64(&g.depth) == 0 {
					return true
				}
			}
		} else if pc >= g.end {
			if atomic.LoadInt64(&g.depth) == 0 {
				return true
			}
		}
	}

	return false
}
