package testing

import (
	"testing"

	"github.com/blacktop/arm64-cgo/emulate/state"
)

// TestCompareStates demonstrates and validates the CompareStates functionality
func TestCompareStates(t *testing.T) {
	t.Run("identical_states", func(t *testing.T) {
		// Create two identical states
		state1 := state.NewState()
		state1.SetX(0, 0x1234)
		state1.SetX(1, 0x5678)
		state1.SetSP(0x7000)
		state1.SetPC(0x10000000)

		state2 := state.NewState()
		state2.SetX(0, 0x1234)
		state2.SetX(1, 0x5678)
		state2.SetSP(0x7000)
		state2.SetPC(0x10000000)

		comparison := CompareStates(state1, state2)
		if comparison.HasDifferences() {
			t.Errorf("Expected no differences, but found: %s", comparison.String())
		}
	})

	t.Run("different_registers", func(t *testing.T) {
		// Create states with different register values
		state1 := state.NewState()
		state1.SetX(0, 0x1234)
		state1.SetX(1, 0x5678)

		state2 := state.NewState()
		state2.SetX(0, 0x4321)  // Different X0
		state2.SetX(1, 0x5678)  // Same X1

		comparison := CompareStates(state1, state2)
		if !comparison.HasDifferences() {
			t.Error("Expected differences, but found none")
		}

		// Check that we found exactly one difference (X0)
		if len(comparison.Differences) != 1 {
			t.Errorf("Expected 1 difference, got %d", len(comparison.Differences))
		}

		diff := comparison.Differences[0]
		if diff.Type != "register" || diff.Name != "X0" {
			t.Errorf("Expected register X0 difference, got %s %s", diff.Type, diff.Name)
		}
	})

	t.Run("different_pc_and_sp", func(t *testing.T) {
		// Create states with different PC and SP
		state1 := state.NewState()
		state1.SetPC(0x10000000)
		state1.SetSP(0x7000)

		state2 := state.NewState()
		state2.SetPC(0x20000000)  // Different PC
		state2.SetSP(0x8000)      // Different SP

		comparison := CompareStates(state1, state2)
		if !comparison.HasDifferences() {
			t.Error("Expected differences, but found none")
		}

		// Should find 2 differences: PC and SP
		if len(comparison.Differences) != 2 {
			t.Errorf("Expected 2 differences, got %d", len(comparison.Differences))
		}

		// Check that we found PC and SP differences
		foundPC, foundSP := false, false
		for _, diff := range comparison.Differences {
			if diff.Type == "register" && diff.Name == "PC" {
				foundPC = true
			}
			if diff.Type == "register" && diff.Name == "SP" {
				foundSP = true
			}
		}

		if !foundPC {
			t.Error("Expected PC difference not found")
		}
		if !foundSP {
			t.Error("Expected SP difference not found")
		}
	})

	t.Run("different_flags", func(t *testing.T) {
		// Create states with different flags
		state1 := state.NewState()
		state1.SetN(true)   // Set Negative flag
		state1.SetZ(false)  // Clear Zero flag

		state2 := state.NewState()
		state2.SetN(false)  // Clear Negative flag (different)
		state2.SetZ(false)  // Clear Zero flag (same)

		comparison := CompareStates(state1, state2)
		if !comparison.HasDifferences() {
			t.Error("Expected differences, but found none")
		}

		// Should find 1 difference: N flag
		flagDiffs := 0
		for _, diff := range comparison.Differences {
			if diff.Type == "flag" {
				flagDiffs++
				if diff.Name == "N" {
					if diff.OurValue != true || diff.HypervisorValue != false {
						t.Errorf("N flag difference values incorrect: our=%v, expected=%v", 
							diff.OurValue, diff.HypervisorValue)
					}
				}
			}
		}

		if flagDiffs != 1 {
			t.Errorf("Expected 1 flag difference, got %d", flagDiffs)
		}
	})

	t.Run("string_output", func(t *testing.T) {
		// Test the string formatting
		state1 := state.NewState()
		state1.SetX(0, 0x1234)
		state1.SetPC(0x10000000)

		state2 := state.NewState() 
		state2.SetX(0, 0x5678)
		state2.SetPC(0x20000000)

		comparison := CompareStates(state1, state2)
		output := comparison.String()

		// Should contain information about differences
		if !containsIgnoreCase(output, "found") {
			t.Error("String output should mention differences found")
		}
		if !containsIgnoreCase(output, "x0") {
			t.Error("String output should mention X0 register")
		}
		if !containsIgnoreCase(output, "pc") {
			t.Error("String output should mention PC register")
		}
	})
}