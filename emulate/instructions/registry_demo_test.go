package instructions

import (
	"fmt"
	"testing"
)

// TestRegistryDemo demonstrates the complete instruction registry functionality
func TestRegistryDemo(t *testing.T) {
	registry := DefaultRegistry()

	fmt.Printf("\n=== ARM64 Instruction Registry Demo ===\n")

	// Show basic statistics
	stats := registry.GetStats()
	fmt.Printf("\nRegistry Statistics:\n")
	fmt.Printf("  Total Instructions: %d\n", stats.TotalInstructions)
	fmt.Printf("  Total Aliases: %d\n", stats.AliasCount)
	fmt.Printf("  Coverage: %.1f%%\n", stats.Coverage)

	// Show top categories
	fmt.Printf("\nTop Instruction Categories:\n")
	for i, category := range stats.TopCategories {
		if i >= 5 { // Show top 5
			break
		}
		fmt.Printf("  %d. %s: %d instructions (%.1f%%)\n",
			i+1, category.Name, category.Count, category.Percentage)
		fmt.Printf("     Examples: %v\n", category.Examples)
	}

	// Demonstrate alias functionality
	fmt.Printf("\nAlias Demonstration:\n")
	aliases := []struct{ alias, primary string }{
		{"B.HS", "B.CS"},
		{"B.LO", "B.CC"},
	}

	for _, alias := range aliases {
		aliasExec, aliasExists := registry.Get(alias.alias)
		primaryExec, primaryExists := registry.Get(alias.primary)

		if aliasExists && primaryExists && aliasExec == primaryExec {
			fmt.Printf("  ✓ %s → %s (same executor)\n", alias.alias, alias.primary)
		} else {
			fmt.Printf("  ✗ %s → %s (failed)\n", alias.alias, alias.primary)
		}
	}

	// Show instructions by category (sample)
	fmt.Printf("\nInstructions by Category (sample):\n")
	byCategory := registry.GetInstructionsByCategory()
	for category, instructions := range byCategory {
		if len(instructions) > 0 {
			fmt.Printf("  %s (%d): ", category, len(instructions))
			// Show first few instructions
			for i, instr := range instructions {
				if i >= 3 {
					fmt.Printf("...")
					break
				}
				if i > 0 {
					fmt.Printf(", ")
				}
				fmt.Printf("%s", instr)
			}
			fmt.Printf("\n")
		}
	}

	// Demonstrate instruction lookup
	fmt.Printf("\nInstruction Lookup Examples:\n")
	testInstructions := []string{"ADD", "LDR", "B.EQ", "MOV", "NONEXISTENT"}
	for _, instr := range testInstructions {
		if registry.HasInstruction(instr) {
			fmt.Printf("  ✓ %s: supported\n", instr)
		} else {
			fmt.Printf("  ✗ %s: not supported\n", instr)
		}
	}

	fmt.Printf("\n=== Demo Complete ===\n\n")
}
