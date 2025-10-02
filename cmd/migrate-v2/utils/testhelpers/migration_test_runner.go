package testhelpers

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

// MigrationTestCase represents a test case for migration testing
type MigrationTestCase struct {
	Name        string
	Input       string
	Expected    string
	Description string // Optional description of what's being tested
}

// MigrationTestRunner provides utilities for testing migrations
type MigrationTestRunner struct {
	migration    internal.ResourceMigration
	resourceType string
}

// NewMigrationTestRunner creates a new test runner for a migration
func NewMigrationTestRunner(migration internal.ResourceMigration, resourceType string) *MigrationTestRunner {
	return &MigrationTestRunner{
		migration:    migration,
		resourceType: resourceType,
	}
}

// RunConfigTests runs configuration migration tests
func (r *MigrationTestRunner) RunConfigTests(t *testing.T, tests []MigrationTestCase) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Parse the input config
			file, diags := hclwrite.ParseConfig([]byte(tt.Input), "test.tf", hcl.InitialPos)
			require.False(t, diags.HasErrors(), "Failed to parse config: %s", diags.Error())

			// Find the resource block
			resourceBlock := r.findResourceBlock(file, r.resourceType)
			require.NotNil(t, resourceBlock, "Resource block not found for type: %s", r.resourceType)

			// Create migration context
			ctx := NewTestContext()

			// Apply migration
			err := r.migration.MigrateConfig(resourceBlock, ctx)
			require.NoError(t, err)

			// Format and compare output
			output := string(hclwrite.Format(file.Bytes()))

			// Normalize whitespace for comparison
			expected := NormalizeHCL(tt.Expected)
			actual := NormalizeHCL(output)

			assert.Equal(t, expected, actual, "Migration output doesn't match expected")
		})
	}
}

// RunStateTests runs state migration tests
func (r *MigrationTestRunner) RunStateTests(t *testing.T, tests []StateTestCase) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ctx := NewTestContext()

			// Make a copy of the input state to avoid mutation
			stateCopy := CopyState(tt.Input)

			err := r.migration.MigrateState(stateCopy, ctx)
			require.NoError(t, err)

			assert.Equal(t, tt.Expected, stateCopy, "State migration output doesn't match expected")
		})
	}
}

// findResourceBlock finds a resource block by type
func (r *MigrationTestRunner) findResourceBlock(file *hclwrite.File, resourceType string) *hclwrite.Block {
	for _, block := range file.Body().Blocks() {
		if block.Type() == "resource" && len(block.Labels()) >= 2 {
			if block.Labels()[0] == resourceType {
				return block
			}
		}
	}
	return nil
}

// StateTestCase represents a test case for state migration
type StateTestCase struct {
	Name     string
	Input    map[string]interface{}
	Expected map[string]interface{}
}

// NewTestContext creates a new migration context for testing
func NewTestContext() *internal.MigrationContext {
	return &internal.MigrationContext{
		Diagnostics: []internal.Diagnostic{},
		Metrics:     &internal.MigrationMetrics{},
		Options: internal.MigrationOptions{
			Verbose: false,
			DryRun:  false,
		},
	}
}

// CopyState creates a deep copy of state map
func CopyState(state map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range state {
		switch val := v.(type) {
		case map[string]interface{}:
			result[k] = CopyState(val)
		case []interface{}:
			copied := make([]interface{}, len(val))
			copy(copied, val)
			result[k] = copied
		default:
			result[k] = v
		}
	}
	return result
}

// NormalizeHCL normalizes HCL for comparison by removing excessive whitespace
func NormalizeHCL(hclStr string) string {
	// Parse and re-format to ensure consistent formatting
	file, diags := hclwrite.ParseConfig([]byte(hclStr), "test.tf", hcl.InitialPos)
	if diags.HasErrors() {
		// If parsing fails, return original
		return hclStr
	}

	formatted := string(hclwrite.Format(file.Bytes()))

	// Clean up excessive blank lines (more than one consecutive blank line)
	// This is a workaround for hclwrite not properly cleaning up whitespace after block removal
	lines := strings.Split(formatted, "\n")
	var cleaned []string
	blankCount := 0

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			blankCount++
			// Allow at most one blank line
			if blankCount <= 1 {
				cleaned = append(cleaned, line)
			}
		} else {
			blankCount = 0
			cleaned = append(cleaned, line)
		}
	}

	return strings.Join(cleaned, "\n")
}

// AssertNoDiagnosticErrors checks that no error diagnostics were generated
func AssertNoDiagnosticErrors(t *testing.T, ctx *internal.MigrationContext) {
	for _, diag := range ctx.Diagnostics {
		if diag.Severity == internal.DiagnosticSeverityError {
			t.Errorf("Unexpected error diagnostic: %s - %s", diag.Summary, diag.Detail)
		}
	}
}

// AssertDiagnosticWarning checks that a specific warning was generated
func AssertDiagnosticWarning(t *testing.T, ctx *internal.MigrationContext, summary string) {
	for _, diag := range ctx.Diagnostics {
		if diag.Severity == internal.DiagnosticSeverityWarning && strings.Contains(diag.Summary, summary) {
			return // Found it
		}
	}
	t.Errorf("Expected warning diagnostic with summary containing: %s", summary)
}
