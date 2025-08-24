# ARM64 Emulator Testing Framework

This package provides comprehensive testing for the ARM64 emulator using the macOS Hypervisor Framework as the authoritative reference implementation for instruction validation.

## Overview

The testing framework compares our ARM64 emulator implementation against the macOS Hypervisor Framework to ensure correctness and compatibility. This approach provides:

- **Authoritative Reference**: The Hypervisor Framework serves as the source of truth for ARM64 instruction execution
- **Comprehensive Validation**: Compares register states, flags, and memory after each instruction
- **Edge Case Discovery**: Helps identify corner cases our implementation might miss
- **Performance Testing**: Validates execution correctness across different scenarios

## Architecture

### File Organization

```
testing/
├── types.go              # Core type definitions and interfaces
├── test_helpers.go       # Shared utilities and common patterns
├── hypervisor_framework.go # macOS Hypervisor Framework integration (Darwin only)
├── framework_test.go     # Framework structure and build validation
├── integration_test.go   # Multi-instruction sequence validation
├── performance_test.go   # Performance benchmarks and comparisons
├── stress_test.go        # Resource limits and stress testing
├── error_handling_test.go # Error scenarios and recovery testing
└── README.md            # This documentation
```

### Core Components

#### Type System (`types.go`)
- `EmulatorState`: Complete processor state including registers, flags, and memory
- `ComparisonResult`: Results of comparing instruction execution between implementations
- `TestCase`: Individual test scenario with initial conditions and expected outcomes
- `TestSuite`: Collection of related test cases for organized validation
- `ValidationResult`: Comprehensive test suite execution results
- `StateDifference`: Detailed differences between expected and actual states

#### Shared Utilities (`test_helpers.go`)
- `CreateTestEngine()`: Standardized engine creation with configuration
- `CommonInstructions`: Frequently used instruction encodings
- `InstructionSequenceBuilder`: Type-safe instruction sequence construction
- `StateComparison`: Utilities for comparing engine states
- `PerformanceHelper`: Performance measurement utilities
- `MemoryHelper`: Memory usage tracking
- `ErrorHelper`: Robust error testing patterns

#### Platform Integration (`hypervisor_framework.go`)
- macOS-specific Hypervisor Framework integration (`//go:build darwin` constraint)
- External dependency: `hv` binary from https://github.com/blacktop/go-hypervisor
- Register state conversion and mapping
- Instruction execution comparison
- Memory state synchronization

## Test Categories

### 1. Framework Validation (`framework_test.go`)
Tests the testing framework itself without external dependencies:
- Type definition completeness and functionality
- Build system validation
- Instruction sequence structure validation
- Encoding format verification

### 2. Integration Testing (`integration_test.go`)
Multi-instruction sequence validation against Hypervisor Framework:
- Arithmetic dependency chains
- Memory operation sequences  
- Function call patterns
- Conditional execution flows
- Complex algorithmic patterns

### 3. Performance Testing (`performance_test.go`)
Performance benchmarks and comparisons:
- Single instruction execution benchmarks
- Sequence execution performance
- Memory efficiency validation
- Hypervisor Framework comparison (when available)

### 4. Stress Testing (`stress_test.go`)
Resource limits and high-load scenarios:
- Large instruction sequence execution
- Memory-intensive operations
- Resource exhaustion boundaries
- Long-running execution patterns
- Concurrent execution scenarios

### 5. Error Handling (`error_handling_test.go`)
Error scenarios and recovery mechanisms:
- Invalid instruction encoding handling
- Memory access violations
- Resource limit enforcement
- State corruption detection
- Recovery mechanism validation

## Build Requirements

### macOS (Full Testing)
The complete testing framework requires macOS with Hypervisor Framework support and the external `hv` utility:

#### Installation
```bash
# Install the go-hypervisor utility (required for Hypervisor Framework integration)
brew install blacktop/tap/go-hypervisor

# Verify installation
hv --version
```

#### Running Tests
```bash
# Run all tests including Hypervisor Framework validation
go test -v ./...

# Run performance benchmarks
go test -v -bench=. ./...

# Run stress tests (may take longer)
go test -v -timeout=30m ./...
```

### Other Platforms (Limited Testing)
On non-macOS platforms, framework validation and structure tests still work:

```bash
# Run framework validation tests only (non-Hypervisor tests)
go test -v -run TestFramework ./...
```

**Note**: Tests that require Hypervisor Framework integration use the `//go:build darwin` build constraint and will automatically be excluded on non-macOS platforms, including CI environments like GitHub Actions.

### Build Constraints
The testing framework uses Go build constraints to ensure platform compatibility:

- **`//go:build darwin`**: Files with this constraint only compile on macOS
  - `hypervisor_framework.go`
  - `integration_test.go` 
  - `performance_test.go` (macOS-specific benchmarks)
- **No constraints**: Files that work on all platforms
  - `framework_test.go`
  - `test_helpers.go`
  - `types.go`
  - `error_handling_test.go`
  - `stress_test.go`

This ensures that CI/CD pipelines on Linux (like GitHub Actions) can run core tests without requiring macOS-specific dependencies.

## Usage Examples

### Creating Test Cases

```go
// Using shared utilities
engine := CreateTestEngine(&EngineConfig{
    MaxInstructions: 1000,
    InitialRegs: map[int]uint64{0: 100, 1: 200},
})

// Using instruction builder
sequence := NewSequenceBuilder().
    Add(CommonInstructions.AddX0Imm1).
    Add(CommonInstructions.AddX1X0).
    Build()

// Executing with performance measurement
helper := NewTestHelper()
duration := helper.ExecuteInstructionSequence(t, engine, sequence, 0x10000000)
```

### State Comparison

```go
// Compare states with detailed diff
ourState := engine.GetState()
expectedState := referenceEngine.GetState()
comparison := CompareStates(ourState, expectedState)
if comparison.HasDifferences() {
    t.Errorf("State mismatch: %s", comparison.String())
}
```

### Error Testing

```go
errorHelper := NewErrorHelper()
err := engine.ExecuteInstruction(pc, invalidInstruction)
if !errorHelper.ContainsError(err, "failed to decode") {
    t.Error("Expected decode error not found")
}
```

## Development Guidelines

### Adding New Tests

1. **Framework Tests**: Add to `framework_test.go` for structural validation
2. **Integration Tests**: Add to `integration_test.go` for multi-instruction sequences  
3. **Performance Tests**: Add to `performance_test.go` with proper benchmarking
4. **Error Tests**: Add to `error_handling_test.go` with expected error patterns

### Test Naming Conventions

- Use descriptive names: `TestArithmeticChainExecution` not `TestTask8_3`
- Group related tests with subtests: `t.Run("subtest_name", func(t *testing.T) {...})`
- Follow Go conventions: `TestFunctionName` for test functions

### Best Practices

1. **Use Shared Utilities**: Leverage `test_helpers.go` to eliminate duplication
2. **Validate Assumptions**: Use framework tests to validate test structure
3. **Handle Platform Differences**: Use build tags for platform-specific functionality
4. **Document Complex Tests**: Add comments explaining test purpose and methodology
5. **Use Table-Driven Tests**: For systematic validation of multiple scenarios

### Performance Considerations

- Use `CreateBenchmarkEngine()` for performance tests to avoid allocation overhead
- Leverage `PerformanceHelper` and `MemoryHelper` for consistent measurements
- Consider test timeouts for long-running stress tests
- Use `testing.B` properly in benchmarks

## Test Validation Categories

### Instruction Sequences Validated

1. **Arithmetic Chains**: ADD, SUB, MUL sequences with register dependencies
2. **Memory Operations**: Load/store patterns with various addressing modes
3. **Function Calls**: Stack management, prologue/epilogue patterns
4. **Conditional Execution**: Branch instructions with flag dependencies
5. **Complex Algorithms**: Real-world computational patterns
6. **Edge Cases**: Boundary conditions, overflow scenarios, special values

### Platform Requirements

#### For Full Testing (macOS only)
- **macOS 10.15+**: Required for Hypervisor Framework integration
- **go-hypervisor**: External utility from https://github.com/blacktop/go-hypervisor
  - Install with: `brew install blacktop/tap/go-hypervisor` 
- **Go 1.19+**: Required for modern Go features used in tests
- **Xcode Command Line Tools**: Required for CGO compilation

#### For Basic Testing (All platforms)
- **Go 1.19+**: Required for framework validation and structure tests
- Tests with `//go:build darwin` will be automatically excluded on non-macOS platforms

## Troubleshooting

### Common Issues

1. **Build Failures on Non-macOS**: Expected - Hypervisor Framework tests are macOS-only and will be automatically excluded
2. **"hv binary not found in PATH"**: Install go-hypervisor with `brew install blacktop/tap/go-hypervisor`
3. **Test Timeouts**: Increase timeout for stress tests: `go test -timeout=30m`
4. **Memory Issues**: Large test suites may require increased memory limits
5. **Hypervisor Access**: Ensure proper permissions for Hypervisor Framework usage
6. **CI/CD Failures**: GitHub Actions and other Linux-based CI will automatically skip Hypervisor tests due to `//go:build darwin` constraints

### Debug Mode

Enable verbose testing for detailed output:
```bash
go test -v -run TestSpecificTest ./...
```

### Performance Analysis

Run with memory profiling:
```bash
go test -memprofile=mem.prof -bench=. ./...
go tool pprof mem.prof
```

## Contributing

When adding new tests:

1. Follow the established file organization
2. Use shared utilities from `test_helpers.go`
3. Add appropriate documentation
4. Ensure tests work on both macOS and other platforms (where applicable)
5. Update this README if adding new test categories

## Implementation Status

- ✅ Framework validation and structure tests
- ✅ Multi-instruction sequence integration tests  
- ✅ Performance benchmarking with Hypervisor Framework comparison
- ✅ Stress testing for resource limits and long-running execution
- ✅ Comprehensive error handling and recovery validation
- ✅ Shared utilities for consistent test patterns
- ✅ Platform-aware testing (macOS Hypervisor Framework integration)

This testing framework provides comprehensive validation of the ARM64 emulator implementation against the authoritative macOS Hypervisor Framework, ensuring correctness and performance across a wide range of instruction sequences and execution scenarios.