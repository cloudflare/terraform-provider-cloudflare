# Testing Guide

This guide explains how to test migrations and ensure they work correctly.

## Unit Testing

### Testing Transformations

Each transformation should have unit tests:

```go
func TestAttributeRenamer(t *testing.T) {
    // Create test HCL
    hcl := `resource "cloudflare_access_application" "example" {
        domain = "example.com"
        enabled = true
    }`
    
    // Parse HCL
    file, _ := hclwrite.ParseConfig([]byte(hcl), "test.tf", hcl.InitialPos)
    block := file.Body().Blocks()[0]
    
    // Apply transformation
    renamer := AttributeRenamer(map[string]string{
        "domain": "hostname",
    })
    err := renamer(block, &TransformContext{})
    
    // Verify results
    assert.NoError(t, err)
    assert.NotNil(t, block.Body().GetAttribute("hostname"))
    assert.Nil(t, block.Body().GetAttribute("domain"))
}
```

### Testing Migrations

Test complete migrations with real configurations:

```go
func TestAccessApplication_V4toV5(t *testing.T) {
    // Load test configuration
    input := loadTestData(t, "testdata/v4_config.tf")
    expected := loadTestData(t, "testdata/v5_expected.tf")
    
    // Create migration
    migration := NewV4toV5Migration()
    
    // Run migration
    ctx := NewMigrationContext()
    err := migration.MigrateConfig(input, ctx)
    
    // Compare results
    assert.NoError(t, err)
    assert.Equal(t, expected, input.BuildTokens(nil).Bytes())
}
```

## Integration Testing

### Test Files Structure

```
resources/access_application/migrations/
├── v4_to_v5.yaml
├── v4_to_v5_test.go
└── testdata/
    ├── v4_basic.tf
    ├── v4_complex.tf
    ├── v5_basic_expected.tf
    └── v5_complex_expected.tf
```

### Running Integration Tests

```bash
# Run all tests
make test

# Run specific resource tests
go test ./resources/access_application/migrations -v

# Run with coverage
go test -cover ./...

# Run specific test case
go test -run TestAccessApplication_V4toV5 ./resources/access_application/migrations
```

## Test Patterns

### Table-Driven Tests

Use table-driven tests for multiple scenarios:

```go
func TestMigration(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "basic configuration",
            input:    "testdata/basic.tf",
            expected: "testdata/basic_expected.tf",
        },
        {
            name:     "with nested blocks",
            input:    "testdata/nested.tf",
            expected: "testdata/nested_expected.tf",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test logic here
        })
    }
}
```

### Edge Cases to Test

Always test these scenarios:

1. **Empty configurations** - Resources with no attributes
2. **Missing attributes** - Optional fields not present
3. **Complex nesting** - Deeply nested structures
4. **Special characters** - Values with quotes, newlines, etc.
5. **References** - Variable and resource references
6. **Comments** - Preserve inline and block comments
7. **Formatting** - Maintain proper HCL formatting

## State Migration Testing

Test state file transformations:

```go
func TestStateMigration(t *testing.T) {
    // Load test state
    state := map[string]interface{}{
        "schema_version": 1,
        "attributes": map[string]interface{}{
            "domain": "example.com",
        },
    }
    
    // Apply migration
    migration := NewV4toV5Migration()
    err := migration.MigrateState(state, ctx)
    
    // Verify state changes
    assert.NoError(t, err)
    assert.Equal(t, 2, state["schema_version"])
    assert.Equal(t, "example.com", state["attributes"].(map[string]interface{})["hostname"])
}
```

## Testing YAML Configurations

Validate YAML configurations load correctly:

```go
func TestYAMLConfiguration(t *testing.T) {
    yamlContent := `
version: "1.0"
resource_type: cloudflare_access_application
migration:
  from: "4.*"
  to: "5.*"
attribute_renames:
  domain: hostname
`
    
    migration, err := NewMigration([]byte(yamlContent))
    assert.NoError(t, err)
    assert.NotNil(t, migration)
}
```

## Test Helpers

### Common Test Utilities

```go
// Load test file
func loadTestData(t *testing.T, path string) *hclwrite.File {
    content, err := os.ReadFile(path)
    require.NoError(t, err)
    
    file, diags := hclwrite.ParseConfig(content, path, hcl.InitialPos)
    require.False(t, diags.HasErrors())
    
    return file
}

// Compare HCL output
func assertHCLEqual(t *testing.T, expected, actual *hclwrite.File) {
    expectedBytes := expected.BuildTokens(nil).Bytes()
    actualBytes := actual.BuildTokens(nil).Bytes()
    
    if !bytes.Equal(expectedBytes, actualBytes) {
        t.Errorf("HCL mismatch:\nExpected:\n%s\nActual:\n%s",
            string(expectedBytes), string(actualBytes))
    }
}
```

## Continuous Integration

### GitHub Actions Example

```yaml
name: Tests
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run tests
        run: make test
      
      - name: Generate coverage
        run: go test -coverprofile=coverage.out ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v3
        with:
          file: ./coverage.out
```

## Best Practices

1. **Test both success and failure cases** - Ensure errors are handled properly
2. **Use real-world examples** - Base tests on actual customer configurations
3. **Keep tests fast** - Use small, focused test cases
4. **Test incrementally** - Build up from simple to complex scenarios
5. **Document test purpose** - Clear test names and comments
6. **Avoid test interdependence** - Each test should be independent
7. **Use assertions wisely** - Clear error messages on failure

## Debugging Tests

### Verbose Output

```bash
# Run with verbose output
go test -v ./...

# Debug specific test
go test -v -run TestName ./package
```

### Using Delve Debugger

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug test
dlv test ./resources/access_application/migrations -- -test.run TestName
```

### Test Output Inspection

```go
// Add debug output in tests
t.Logf("Before transformation: %s", block.BuildTokens(nil).Bytes())
// Apply transformation
t.Logf("After transformation: %s", block.BuildTokens(nil).Bytes())
```