package instructions

import (
	"testing"

	"github.com/blacktop/arm64-cgo/disassemble"
	"github.com/blacktop/arm64-cgo/emulate/core"
)

func TestDefaultRegistryIntegration(t *testing.T) {
	registry := DefaultRegistry()

	t.Run("Registry Creation", func(t *testing.T) {
		if registry == nil {
			t.Fatal("DefaultRegistry() returned nil")
		}

		count := registry.Count()
		if count == 0 {
			t.Fatal("Registry has no instructions registered")
		}

		t.Logf("Registry contains %d instructions", count)
	})

	t.Run("All Categories Registered", func(t *testing.T) {
		expectedCategories := []string{
			"Arithmetic", "Memory", "Branch", "Move",
			"Logical", "Compare", "Conditional", "System",
		}

		stats := registry.GetStats()

		for _, expectedCategory := range expectedCategories {
			if count, exists := stats.Categories[expectedCategory]; !exists || count == 0 {
				t.Errorf("Category %s is missing or has no instructions", expectedCategory)
			} else {
				t.Logf("Category %s has %d instructions", expectedCategory, count)
			}
		}
	})

	t.Run("Core Instructions Present", func(t *testing.T) {
		coreInstructions := []string{
			// Arithmetic
			"ADD", "SUB", "MUL", "UDIV", "SDIV",
			// Memory
			"LDR", "STR", "LDP", "STP",
			// Branch
			"B", "BL", "BR", "BLR", "RET",
			// Move
			"MOV", "MOVZ", "MOVK",
			// Logical
			"AND", "ORR", "EOR", "BIC",
			// Compare
			"CMP", "CMN",
			// Conditional
			"CSEL", "CSET",
			// System
			"NOP", "MRS", "MSR",
		}

		for _, instruction := range coreInstructions {
			if !registry.HasInstruction(instruction) {
				t.Errorf("Core instruction %s is not registered", instruction)
			}
		}
	})

	t.Run("Aliases Work Correctly", func(t *testing.T) {
		aliasTests := []struct {
			alias   string
			primary string
		}{
			{"B.HS", "B.CS"},
			{"B.LO", "B.CC"},
		}

		for _, test := range aliasTests {
			aliasExec, aliasExists := registry.Get(test.alias)
			primaryExec, primaryExists := registry.Get(test.primary)

			if !aliasExists {
				t.Errorf("Alias %s is not registered", test.alias)
				continue
			}

			if !primaryExists {
				t.Errorf("Primary instruction %s is not registered", test.primary)
				continue
			}

			// Both should return the same executor instance
			if aliasExec != primaryExec {
				t.Errorf("Alias %s does not resolve to the same executor as %s", test.alias, test.primary)
			}
		}
	})

	t.Run("No Duplicate Registrations", func(t *testing.T) {
		// Test that instructions like TST and CCMP are not registered multiple times
		// by checking that we can get them and they work consistently

		if tstExec, exists := registry.Get("TST"); !exists {
			t.Error("TST instruction is not registered")
		} else {
			// TST should be in logical category (it's AND with flags)
			if !tstExec.Supports("TST") {
				t.Error("TST executor does not support TST mnemonic")
			}
		}

		if ccmpExec, exists := registry.Get("CCMP"); !exists {
			t.Error("CCMP instruction is not registered")
		} else {
			if !ccmpExec.Supports("CCMP") {
				t.Error("CCMP executor does not support CCMP mnemonic")
			}
		}
	})

	t.Run("Statistics Are Accurate", func(t *testing.T) {
		stats := registry.GetStats()

		if stats.TotalInstructions != registry.Count() {
			t.Errorf("Stats total instructions (%d) does not match registry count (%d)",
				stats.TotalInstructions, registry.Count())
		}

		// Verify category counts add up
		totalFromCategories := 0
		for _, count := range stats.Categories {
			totalFromCategories += count
		}

		if totalFromCategories != stats.TotalInstructions {
			t.Errorf("Category counts (%d) do not add up to total instructions (%d)",
				totalFromCategories, stats.TotalInstructions)
		}

		// Check that we have reasonable coverage
		if stats.Coverage < 10 {
			t.Errorf("Coverage is too low: %.2f%%", stats.Coverage)
		}

		t.Logf("Registry statistics:")
		t.Logf("  Total instructions: %d", stats.TotalInstructions)
		t.Logf("  Total aliases: %d", stats.AliasCount)
		t.Logf("  Coverage: %.2f%%", stats.Coverage)

		for _, category := range stats.TopCategories {
			t.Logf("  %s: %d instructions (%.1f%%)",
				category.Name, category.Count, category.Percentage)
		}
	})

	t.Run("Instruction Info Complete", func(t *testing.T) {
		info := registry.GetInstructionInfo()

		if len(info) != registry.Count() {
			t.Errorf("Instruction info count (%d) does not match registry count (%d)",
				len(info), registry.Count())
		}

		// Check that all instructions have proper metadata
		for _, instrInfo := range info {
			if instrInfo.Mnemonic == "" {
				t.Error("Found instruction with empty mnemonic")
			}

			if instrInfo.Category == "" {
				t.Errorf("Instruction %s has empty category", instrInfo.Mnemonic)
			}

			if instrInfo.Category == "Other" {
				t.Logf("Warning: Instruction %s is categorized as 'Other'", instrInfo.Mnemonic)
			}

			if !instrInfo.Supported {
				t.Errorf("Instruction %s is marked as not supported", instrInfo.Mnemonic)
			}
		}
	})

	t.Run("Instructions By Category", func(t *testing.T) {
		byCategory := registry.GetInstructionsByCategory()

		// Verify all categories are present
		stats := registry.GetStats()
		for category := range stats.Categories {
			if instructions, exists := byCategory[category]; !exists {
				t.Errorf("Category %s is missing from GetInstructionsByCategory", category)
			} else if len(instructions) == 0 {
				t.Errorf("Category %s has no instructions in GetInstructionsByCategory", category)
			}
		}

		// Verify instructions are sorted within categories
		for category, instructions := range byCategory {
			for i := 1; i < len(instructions); i++ {
				if instructions[i-1] > instructions[i] {
					t.Errorf("Instructions in category %s are not sorted: %s > %s",
						category, instructions[i-1], instructions[i])
					break
				}
			}
		}
	})
}

func TestRegistryThreadSafety(t *testing.T) {
	registry := NewRegistry()

	// Test concurrent registration and access
	done := make(chan bool, 2)

	// Goroutine 1: Register instructions
	go func() {
		defer func() { done <- true }()

		for i := 0; i < 100; i++ {
			executor := NewExecutorFunc("TEST", func(state core.State, instr *disassemble.Instruction) error {
				return nil
			})
			registry.Register("TEST", executor)
		}
	}()

	// Goroutine 2: Access registry
	go func() {
		defer func() { done <- true }()

		for i := 0; i < 100; i++ {
			registry.Get("TEST")
			registry.List()
			registry.Count()
			registry.HasInstruction("TEST")
		}
	}()

	// Wait for both goroutines
	<-done
	<-done

	// Verify final state
	if !registry.HasInstruction("TEST") {
		t.Error("TEST instruction was not properly registered")
	}
}

func TestRegistryErrorHandling(t *testing.T) {
	registry := NewRegistry()

	t.Run("Empty Mnemonic", func(t *testing.T) {
		executor := NewExecutorFunc("", func(state core.State, instr *disassemble.Instruction) error {
			return nil
		})

		err := registry.Register("", executor)
		if err == nil {
			t.Error("Expected error for empty mnemonic")
		}
	})

	t.Run("Nil Executor", func(t *testing.T) {
		err := registry.Register("TEST", nil)
		if err == nil {
			t.Error("Expected error for nil executor")
		}
	})

	t.Run("Nonexistent Instruction", func(t *testing.T) {
		_, exists := registry.Get("NONEXISTENT")
		if exists {
			t.Error("Should not find nonexistent instruction")
		}

		if registry.HasInstruction("NONEXISTENT") {
			t.Error("Should not have nonexistent instruction")
		}
	})
}

func BenchmarkRegistryOperations(b *testing.B) {
	registry := DefaultRegistry()

	b.Run("Get", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			registry.Get("ADD")
		}
	})

	b.Run("HasInstruction", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			registry.HasInstruction("ADD")
		}
	})

	b.Run("List", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			registry.List()
		}
	})

	b.Run("GetStats", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			registry.GetStats()
		}
	})

	b.Run("GetInstructionInfo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			registry.GetInstructionInfo()
		}
	})
}
