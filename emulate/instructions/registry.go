package instructions

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

// Registry implements the core.InstructionRegistry interface
type Registry struct {
	mu        sync.RWMutex
	executors map[string]core.InstructionExecutor
	aliases   map[string]string // Maps aliases to canonical names
}

// NewRegistry creates a new instruction registry
func NewRegistry() *Registry {
	return &Registry{
		executors: make(map[string]core.InstructionExecutor),
		aliases:   make(map[string]string),
	}
}

// Register registers an instruction executor for a given mnemonic
func (r *Registry) Register(mnemonic string, executor core.InstructionExecutor) error {
	if mnemonic == "" {
		return fmt.Errorf("empty mnemonic")
	}
	if executor == nil {
		return fmt.Errorf("nil executor")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	canonical := strings.ToUpper(mnemonic)
	r.executors[canonical] = executor
	return nil
}

// RegisterWithAliases registers an instruction executor with multiple aliases
func (r *Registry) RegisterWithAliases(primary string, executor core.InstructionExecutor, aliases ...string) error {
	if err := r.Register(primary, executor); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	canonical := strings.ToUpper(primary)
	for _, alias := range aliases {
		if alias != "" {
			r.aliases[strings.ToUpper(alias)] = canonical
		}
	}

	return nil
}

// Get retrieves an instruction executor for a given mnemonic
func (r *Registry) Get(mnemonic string) (core.InstructionExecutor, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	canonical := strings.ToUpper(mnemonic)

	// Check for direct match first
	if executor, exists := r.executors[canonical]; exists {
		return executor, true
	}

	// Check for alias
	if primary, exists := r.aliases[canonical]; exists {
		if executor, exists := r.executors[primary]; exists {
			return executor, true
		}
	}

	return nil, false
}

// List returns all registered instruction mnemonics
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	mnemonics := make([]string, 0, len(r.executors))
	for mnemonic := range r.executors {
		mnemonics = append(mnemonics, mnemonic)
	}

	sort.Strings(mnemonics)
	return mnemonics
}

// Clear removes all registered executors
func (r *Registry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.executors = make(map[string]core.InstructionExecutor)
	r.aliases = make(map[string]string)
}

// Count returns the number of registered executors
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.executors)
}

// HasInstruction checks if an instruction is supported
func (r *Registry) HasInstruction(mnemonic string) bool {
	_, exists := r.Get(mnemonic)
	return exists
}

// BaseExecutor provides common functionality for instruction executors
type BaseExecutor struct {
	mnemonic    string
	description string
}

// NewBaseExecutor creates a new base executor
func NewBaseExecutor(mnemonic, description string) *BaseExecutor {
	return &BaseExecutor{
		mnemonic:    strings.ToUpper(mnemonic),
		description: description,
	}
}

// Supports checks if this executor supports the given mnemonic
func (e *BaseExecutor) Supports(mnemonic string) bool {
	return strings.ToUpper(mnemonic) == e.mnemonic
}

// GetMnemonic returns the mnemonic this executor handles
func (e *BaseExecutor) GetMnemonic() string {
	return e.mnemonic
}

// GetDescription returns the description of this instruction
func (e *BaseExecutor) GetDescription() string {
	return e.description
}

// ValidateInstruction performs basic validation on an instruction
func (e *BaseExecutor) ValidateInstruction(instr *disassemble.Instruction) error {
	if instr == nil {
		return core.NewEmulationError(core.ErrInvalidInstruction, 0, e.mnemonic, "nil instruction")
	}

	if !e.Supports(fmt.Sprintf("%v", instr.Operation)) {
		return core.NewEmulationError(core.ErrInvalidInstruction, 0, fmt.Sprintf("%v", instr.Operation),
			fmt.Sprintf("executor %s does not support %s", e.mnemonic, fmt.Sprintf("%v", instr.Operation)))
	}

	return nil
}

// ExecutorFunc is a function type that implements InstructionExecutor
type ExecutorFunc struct {
	mnemonic string
	fn       func(state core.State, instr *disassemble.Instruction) error
}

// NewExecutorFunc creates a new function-based executor
func NewExecutorFunc(mnemonic string, fn func(state core.State, instr *disassemble.Instruction) error) *ExecutorFunc {
	return &ExecutorFunc{
		mnemonic: strings.ToUpper(mnemonic),
		fn:       fn,
	}
}

// Execute executes the instruction using the function
func (e *ExecutorFunc) Execute(state core.State, instr *disassemble.Instruction) error {
	return e.fn(state, instr)
}

// Supports checks if this executor supports the given mnemonic
func (e *ExecutorFunc) Supports(mnemonic string) bool {
	return strings.ToUpper(mnemonic) == e.mnemonic
}

// DefaultRegistry creates and returns a registry with all standard ARM64 instructions
func DefaultRegistry() *Registry {
	registry := NewRegistry()

	// Register all instruction families
	RegisterArithmeticInstructions(registry)
	RegisterMemoryInstructions(registry)
	RegisterBranchInstructions(registry)
	RegisterMoveInstructions(registry)
	RegisterLogicalInstructions(registry)
	RegisterCompareInstructions(registry)
	RegisterConditionalInstructions(registry)
	RegisterSystemInstructions(registry)

	// Register additional aliases for common instruction variants
	registerCommonAliases(registry)

	return registry
}

// registerCommonAliases registers common instruction aliases
func registerCommonAliases(registry *Registry) {
	// Get executors for alias registration
	if subExec, exists := registry.Get("SUB"); exists {
		// NEG is typically an alias for SUB Xd, XZR, Xm
		registry.RegisterWithAliases("SUB", subExec, "NEG")
	}

	if movExec, exists := registry.Get("MOV"); exists {
		// MV is a common shorthand for MOV
		registry.RegisterWithAliases("MOV", movExec, "MV")
	}

	if lslExec, exists := registry.Get("LSL"); exists {
		// SHL is an alias for LSL (shift left)
		registry.RegisterWithAliases("LSL", lslExec, "SHL")
	}

	if lsrExec, exists := registry.Get("LSR"); exists {
		// SHR is an alias for LSR (logical shift right)
		registry.RegisterWithAliases("LSR", lsrExec, "SHR")
	}

	// Branch condition aliases are already handled in RegisterBranchInstructions
	// Additional common aliases can be added here as needed
}

// InstructionInfo provides metadata about registered instructions
type InstructionInfo struct {
	Mnemonic    string   `json:"mnemonic"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Aliases     []string `json:"aliases"`
	Supported   bool     `json:"supported"`
}

// GetInstructionInfo returns information about all registered instructions
func (r *Registry) GetInstructionInfo() []InstructionInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	info := make([]InstructionInfo, 0, len(r.executors))

	for mnemonic, executor := range r.executors {
		instrInfo := InstructionInfo{
			Mnemonic:  mnemonic,
			Category:  categorizeInstruction(mnemonic),
			Supported: true,
		}

		// Try to get description from various executor types
		instrInfo.Description = getExecutorDescription(executor)

		// Find aliases for this instruction
		aliases := make([]string, 0)
		for alias, canonical := range r.aliases {
			if canonical == mnemonic {
				aliases = append(aliases, alias)
			}
		}
		sort.Strings(aliases)
		instrInfo.Aliases = aliases

		info = append(info, instrInfo)
	}

	sort.Slice(info, func(i, j int) bool {
		if info[i].Category != info[j].Category {
			return info[i].Category < info[j].Category
		}
		return info[i].Mnemonic < info[j].Mnemonic
	})

	return info
}

// getExecutorDescription attempts to extract description from various executor types
func getExecutorDescription(executor core.InstructionExecutor) string {
	// Try different executor types that might have descriptions
	switch e := executor.(type) {
	case interface{ GetDescription() string }:
		return e.GetDescription()
	default:
		return "ARM64 instruction"
	}
}

// GetInstructionsByCategory returns instructions grouped by category
func (r *Registry) GetInstructionsByCategory() map[string][]string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make(map[string][]string)

	for mnemonic := range r.executors {
		category := categorizeInstruction(mnemonic)
		categories[category] = append(categories[category], mnemonic)
	}

	// Sort instructions within each category
	for category := range categories {
		sort.Strings(categories[category])
	}

	return categories
}

// Stats provides comprehensive statistics about the registry
type Stats struct {
	TotalInstructions int            `json:"total_instructions"`
	Categories        map[string]int `json:"categories"`
	AliasCount        int            `json:"alias_count"`
	TopCategories     []CategoryStat `json:"top_categories"`
	Coverage          float64        `json:"coverage_percentage"`
}

// CategoryStat provides statistics for a specific instruction category
type CategoryStat struct {
	Name       string   `json:"name"`
	Count      int      `json:"count"`
	Percentage float64  `json:"percentage"`
	Examples   []string `json:"examples"`
}

// GetStats returns comprehensive statistics about the registry
func (r *Registry) GetStats() Stats {
	r.mu.RLock()
	defer r.mu.RUnlock()

	totalInstructions := len(r.executors)
	categories := make(map[string]int)
	categoryExamples := make(map[string][]string)

	// Count categories and collect examples
	for mnemonic := range r.executors {
		category := categorizeInstruction(mnemonic)
		categories[category]++

		// Collect up to 3 examples per category
		if len(categoryExamples[category]) < 3 {
			categoryExamples[category] = append(categoryExamples[category], mnemonic)
		}
	}

	// Create top categories list
	topCategories := make([]CategoryStat, 0, len(categories))
	for category, count := range categories {
		percentage := float64(count) / float64(totalInstructions) * 100
		topCategories = append(topCategories, CategoryStat{
			Name:       category,
			Count:      count,
			Percentage: percentage,
			Examples:   categoryExamples[category],
		})
	}

	// Sort by count (descending)
	sort.Slice(topCategories, func(i, j int) bool {
		return topCategories[i].Count > topCategories[j].Count
	})

	// Estimate coverage (rough estimate based on common ARM64 instructions)
	estimatedTotalARM64Instructions := 300.0 // Rough estimate
	coverage := float64(totalInstructions) / estimatedTotalARM64Instructions * 100
	if coverage > 100 {
		coverage = 100
	}

	return Stats{
		TotalInstructions: totalInstructions,
		Categories:        categories,
		AliasCount:        len(r.aliases),
		TopCategories:     topCategories,
		Coverage:          coverage,
	}
}

// categorizeInstruction provides comprehensive categorization based on mnemonic
func categorizeInstruction(mnemonic string) string {
	// Arithmetic operations
	if isArithmeticInstruction(mnemonic) {
		return "Arithmetic"
	}

	// Memory operations
	if isMemoryInstruction(mnemonic) {
		return "Memory"
	}

	// Branch operations
	if isBranchInstruction(mnemonic) {
		return "Branch"
	}

	// Move operations
	if isMoveInstruction(mnemonic) {
		return "Move"
	}

	// Logical operations
	if isLogicalInstruction(mnemonic) {
		return "Logical"
	}

	// Compare operations
	if isCompareInstruction(mnemonic) {
		return "Compare"
	}

	// Conditional operations
	if isConditionalInstruction(mnemonic) {
		return "Conditional"
	}

	// System operations
	if isSystemInstruction(mnemonic) {
		return "System"
	}

	// Cryptographic operations
	if isCryptoInstruction(mnemonic) {
		return "Crypto"
	}

	return "Other"
}

// Helper functions for instruction categorization
func isArithmeticInstruction(mnemonic string) bool {
	arithmeticPrefixes := []string{"ADD", "SUB", "MUL", "DIV", "MADD", "MSUB", "NEG", "SMULL", "UMULL", "UDIV", "SDIV"}
	for _, prefix := range arithmeticPrefixes {
		if strings.HasPrefix(mnemonic, prefix) {
			return true
		}
	}
	return false
}

func isMemoryInstruction(mnemonic string) bool {
	memoryPrefixes := []string{"LDR", "STR", "LDP", "STP", "LDUR", "STUR"}
	for _, prefix := range memoryPrefixes {
		if strings.HasPrefix(mnemonic, prefix) {
			return true
		}
	}
	return false
}

func isBranchInstruction(mnemonic string) bool {
	// Handle conditional branches first
	if strings.HasPrefix(mnemonic, "B.") {
		return true
	}

	branchInstructions := []string{"B", "BL", "BR", "BLR", "RET", "CBZ", "CBNZ", "TBZ", "TBNZ"}
	for _, instr := range branchInstructions {
		if mnemonic == instr {
			return true
		}
	}
	return false
}

func isMoveInstruction(mnemonic string) bool {
	movePrefixes := []string{"MOV", "ADR", "SXT", "UXT"}
	for _, prefix := range movePrefixes {
		if strings.HasPrefix(mnemonic, prefix) {
			return true
		}
	}
	return false
}

func isLogicalInstruction(mnemonic string) bool {
	logicalInstructions := []string{"AND", "ANDS", "ORR", "EOR", "BIC", "BICS", "TST",
		"BFM", "BFI", "BFXIL", "SBFM", "SBFIZ", "SBFX", "UBFM", "UBFIZ", "UBFX",
		"ASR", "LSL", "LSR"}
	for _, instr := range logicalInstructions {
		if strings.HasPrefix(mnemonic, instr) {
			return true
		}
	}
	return false
}

func isCompareInstruction(mnemonic string) bool {
	compareInstructions := []string{"CMP", "CMN", "CCMP", "CCMN"}
	for _, instr := range compareInstructions {
		if strings.HasPrefix(mnemonic, instr) {
			return true
		}
	}
	return false
}

func isConditionalInstruction(mnemonic string) bool {
	conditionalPrefixes := []string{"CSEL", "CSINC", "CSINV", "CSNEG", "CSET", "CINC", "CINV", "CNEG"}
	for _, prefix := range conditionalPrefixes {
		if strings.HasPrefix(mnemonic, prefix) {
			return true
		}
	}
	return false
}

func isSystemInstruction(mnemonic string) bool {
	systemInstructions := []string{"NOP", "MRS", "MSR", "SYS", "SYSL", "ISB", "DSB", "DMB",
		"HINT", "YIELD", "WFE", "WFI", "HLT", "SEV", "SEVL"}
	for _, instr := range systemInstructions {
		if strings.HasPrefix(mnemonic, instr) {
			return true
		}
	}
	return false
}

func isCryptoInstruction(mnemonic string) bool {
	cryptoPrefixes := []string{"AES", "SHA", "PMULL"}
	for _, prefix := range cryptoPrefixes {
		if strings.Contains(mnemonic, prefix) {
			return true
		}
	}
	return false
}
