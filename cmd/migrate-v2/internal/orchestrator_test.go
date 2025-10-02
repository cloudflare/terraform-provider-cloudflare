package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrchestrator_MigrateFile(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test HCL file
	testFile := filepath.Join(tempDir, "test.tf")
	testContent := `
resource "test_resource" "example" {
  old_attr = "value"
  keep_attr = "keep"
}

resource "other_resource" "example" {
  attr = "value"
}
`
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create mock registry with migration
	registry := NewDefaultRegistry()
	migration := &mockMigration{
		resourceType:  "test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}
	require.NoError(t, registry.Register(migration))

	// Create orchestrator
	options := MigrationOptions{
		DryRun:  true, // Don't actually write files
		Verbose: false,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate file
	err = orchestrator.MigrateFile(testFile, ctx)
	assert.NoError(t, err)

	// Verify metrics
	assert.Equal(t, 1, ctx.Metrics.TotalResources) // Only test_resource should be counted
	assert.Equal(t, 1, ctx.Metrics.MigratedResources)
	assert.Equal(t, 0, ctx.Metrics.FailedResources)
}

func TestOrchestrator_MigrateFile_NoMigrationAvailable(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test HCL file with resource that has no migration
	testFile := filepath.Join(tempDir, "test.tf")
	testContent := `
resource "unmigrated_resource" "example" {
  attr = "value"
}
`
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create empty registry
	registry := NewDefaultRegistry()

	// Create orchestrator
	options := MigrationOptions{
		DryRun: true,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate file - should not fail even if no migration available
	err = orchestrator.MigrateFile(testFile, ctx)
	assert.NoError(t, err)

	// Verify metrics - no resources should be migrated
	assert.Equal(t, 0, ctx.Metrics.TotalResources)
	assert.Equal(t, 0, ctx.Metrics.MigratedResources)
	assert.Equal(t, 0, ctx.Metrics.FailedResources)
}

func TestOrchestrator_MigrateFile_MigrationError(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test HCL file
	testFile := filepath.Join(tempDir, "test.tf")
	testContent := `
resource "failing_resource" "example" {
  attr = "value"
}
`
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create mock registry with failing migration
	registry := NewDefaultRegistry()
	migration := &mockMigration{
		resourceType:  "failing_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
		migrateError:  fmt.Errorf("migration failed"),
	}
	require.NoError(t, registry.Register(migration))

	// Create orchestrator
	options := MigrationOptions{
		DryRun: true,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate file
	err = orchestrator.MigrateFile(testFile, ctx)
	assert.NoError(t, err) // File migration shouldn't fail, but resource should be marked as failed

	// Verify metrics
	assert.Equal(t, 1, ctx.Metrics.TotalResources)
	assert.Equal(t, 0, ctx.Metrics.MigratedResources)
	assert.Equal(t, 1, ctx.Metrics.FailedResources)

	// Verify error diagnostic was added
	assert.NotEmpty(t, ctx.Diagnostics)
	hasError := false
	for _, diag := range ctx.Diagnostics {
		if diag.Severity == DiagnosticSeverityError {
			hasError = true
			break
		}
	}
	assert.True(t, hasError)
}

func TestOrchestrator_MigrateDirectory(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create multiple test files
	files := []struct {
		name    string
		content string
	}{
		{
			name: "file1.tf",
			content: `
resource "test_resource" "one" {
  attr = "value1"
}
`,
		},
		{
			name: "file2.tf",
			content: `
resource "test_resource" "two" {
  attr = "value2"
}
`,
		},
		{
			name: "file3.tf",
			content: `
resource "other_resource" "three" {
  attr = "value3"
}
`,
		},
	}

	for _, f := range files {
		err = ioutil.WriteFile(filepath.Join(tempDir, f.name), []byte(f.content), 0644)
		require.NoError(t, err)
	}

	// Create mock registry
	registry := NewDefaultRegistry()
	migration := &mockMigration{
		resourceType:  "test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}
	require.NoError(t, registry.Register(migration))

	// Create orchestrator
	options := MigrationOptions{
		DryRun:   true,
		Parallel: false,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate directory
	err = orchestrator.MigrateDirectory(tempDir, ctx)
	assert.NoError(t, err)

	// Verify metrics - should have processed 2 test_resource instances
	assert.Equal(t, 2, ctx.Metrics.TotalResources)
	assert.Equal(t, 2, ctx.Metrics.MigratedResources)
	assert.Equal(t, 0, ctx.Metrics.FailedResources)
}

func TestOrchestrator_MigrateDirectory_NoTfFiles(t *testing.T) {
	// Create a temporary directory with no .tf files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a non-.tf file
	err = ioutil.WriteFile(filepath.Join(tempDir, "readme.md"), []byte("# README"), 0644)
	require.NoError(t, err)

	// Create orchestrator
	registry := NewDefaultRegistry()
	options := MigrationOptions{
		DryRun: true,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate directory - should fail with no .tf files
	err = orchestrator.MigrateDirectory(tempDir, ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no .tf files found")
}

func TestOrchestrator_MigrateDirectory_Parallel(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create multiple test files
	for i := 0; i < 10; i++ {
		filename := fmt.Sprintf("file%d.tf", i)
		content := fmt.Sprintf(`
resource "test_resource" "resource_%d" {
  attr = "value%d"
}
`, i, i)
		err = ioutil.WriteFile(filepath.Join(tempDir, filename), []byte(content), 0644)
		require.NoError(t, err)
	}

	// Create mock registry
	registry := NewDefaultRegistry()
	migration := &mockMigration{
		resourceType:  "test_resource",
		sourceVersion: "v4",
		targetVersion: "v5",
	}
	require.NoError(t, registry.Register(migration))

	// Create orchestrator with parallel processing
	options := MigrationOptions{
		DryRun:   true,
		Parallel: true,
		Workers:  4,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate directory in parallel
	err = orchestrator.MigrateDirectory(tempDir, ctx)
	assert.NoError(t, err)

	// Verify all resources were processed
	assert.Equal(t, 10, ctx.Metrics.TotalResources)
	assert.Equal(t, 10, ctx.Metrics.MigratedResources)
	assert.Equal(t, 0, ctx.Metrics.FailedResources)
}

func TestOrchestrator_MigrateFile_InvalidHCL(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test file with invalid HCL
	testFile := filepath.Join(tempDir, "invalid.tf")
	testContent := `
resource "test_resource" "example" {
  attr = "value"
  # Missing closing brace
`
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create orchestrator
	registry := NewDefaultRegistry()
	options := MigrationOptions{
		DryRun: true,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate file - should fail due to parse error
	err = orchestrator.MigrateFile(testFile, ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse HCL")
}

func TestOrchestrator_MigrateFile_WriteFile(t *testing.T) {
	// Create a temporary directory for test files
	tempDir, err := ioutil.TempDir("", "orchestrator_test")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create test HCL file
	testFile := filepath.Join(tempDir, "test.tf")
	testContent := `resource "test_resource" "example" {
  attr = "value"
}`
	err = ioutil.WriteFile(testFile, []byte(testContent), 0644)
	require.NoError(t, err)

	// Create mock registry with migration that modifies the block
	registry := NewDefaultRegistry()
	migration := &mockMigrationWithTransform{
		mockMigration: mockMigration{
			resourceType:  "test_resource",
			sourceVersion: "v4",
			targetVersion: "v5",
		},
		transform: func(block *hclwrite.Block) {
			// Simply set a raw attribute for testing
			block.Body().SetAttributeRaw("new_attr", hclwrite.Tokens{
				&hclwrite.Token{Type: hclsyntax.TokenIdent, Bytes: []byte(`"new_value"`)},
			})
		},
	}
	require.NoError(t, registry.Register(migration))

	// Create orchestrator WITHOUT dry-run
	options := MigrationOptions{
		DryRun:  false, // Actually write the file
		Verbose: false,
	}
	orchestrator := NewOrchestrator(registry, options)

	// Create context
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Options:     options,
	}

	// Migrate file
	err = orchestrator.MigrateFile(testFile, ctx)
	assert.NoError(t, err)

	// Read the modified file and verify it was written
	content, err := ioutil.ReadFile(testFile)
	require.NoError(t, err)
	assert.Contains(t, string(content), "new_attr")
}

// Helper types for testing

type mockMigrationWithTransform struct {
	mockMigration
	transform func(*hclwrite.Block)
}

func (m *mockMigrationWithTransform) MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error {
	if m.transform != nil {
		m.transform(block)
	}
	return m.migrateError
}
