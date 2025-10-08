package core

// HookKind enumerates the supported hook categories for the emulator engine.
type HookKind int

const (
	// HookPreInstruction fires before the instruction executor runs, after decode.
	HookPreInstruction HookKind = iota
	// HookPostInstruction fires after the instruction executor completes.
	HookPostInstruction
	// HookUnimplementedInstruction fires when an instruction is unsupported or unimplemented.
	HookUnimplementedInstruction
)

// HookResult communicates actions a hook would like the engine to take.
type HookResult struct {
	// Halt stops execution after the current iteration.
	Halt bool
	// SkipInstruction requests that the current instruction be treated as a no-op.
	SkipInstruction bool
}

// PreInstructionHook observes an instruction prior to execution and may influence control-flow.
type PreInstructionHook func(state State, info InstructionInfo) HookResult

// PostInstructionHook observes an instruction after execution.
type PostInstructionHook func(state State, info InstructionInfo) HookResult

// UnimplementedInstructionHook observes instructions that lack emulator support.
type UnimplementedInstructionHook func(state State, info InstructionInfo, err error) HookResult
