package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrationContext_AddDiagnostic(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Add different severity diagnostics
	ctx.AddDiagnostic(DiagnosticSeverityInfo, "Info message", "Info detail", "test_resource")
	ctx.AddDiagnostic(DiagnosticSeverityWarning, "Warning message", "Warning detail", "test_resource")
	ctx.AddDiagnostic(DiagnosticSeverityError, "Error message", "Error detail", "test_resource")

	// Verify diagnostics were added
	assert.Len(t, ctx.Diagnostics, 3)

	assert.Equal(t, DiagnosticSeverityInfo, ctx.Diagnostics[0].Severity)
	assert.Equal(t, "Info message", ctx.Diagnostics[0].Summary)
	assert.Equal(t, "Info detail", ctx.Diagnostics[0].Detail)
	assert.Equal(t, "test_resource", ctx.Diagnostics[0].Resource)

	assert.Equal(t, DiagnosticSeverityWarning, ctx.Diagnostics[1].Severity)
	assert.Equal(t, "Warning message", ctx.Diagnostics[1].Summary)

	assert.Equal(t, DiagnosticSeverityError, ctx.Diagnostics[2].Severity)
	assert.Equal(t, "Error message", ctx.Diagnostics[2].Summary)
}

func TestMigrationContext_AddInfo(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	ctx.AddInfo("Info summary", "Info details", "resource_a")

	assert.Len(t, ctx.Diagnostics, 1)
	assert.Equal(t, DiagnosticSeverityInfo, ctx.Diagnostics[0].Severity)
	assert.Equal(t, "Info summary", ctx.Diagnostics[0].Summary)
	assert.Equal(t, "Info details", ctx.Diagnostics[0].Detail)
	assert.Equal(t, "resource_a", ctx.Diagnostics[0].Resource)

	// Info should not increment warning count
	assert.Equal(t, 0, ctx.Metrics.WarningCount)
}

func TestMigrationContext_AddWarning(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	ctx.AddWarning("Warning 1", "Detail 1", "resource_a")
	ctx.AddWarning("Warning 2", "Detail 2", "resource_b")

	assert.Len(t, ctx.Diagnostics, 2)
	assert.Equal(t, DiagnosticSeverityWarning, ctx.Diagnostics[0].Severity)
	assert.Equal(t, DiagnosticSeverityWarning, ctx.Diagnostics[1].Severity)

	// Warning count should be incremented
	assert.Equal(t, 2, ctx.Metrics.WarningCount)
}

func TestMigrationContext_AddError(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	ctx.AddError("Error summary", "Error details", "failing_resource")

	assert.Len(t, ctx.Diagnostics, 1)
	assert.Equal(t, DiagnosticSeverityError, ctx.Diagnostics[0].Severity)
	assert.Equal(t, "Error summary", ctx.Diagnostics[0].Summary)
	assert.Equal(t, "Error details", ctx.Diagnostics[0].Detail)
	assert.Equal(t, "failing_resource", ctx.Diagnostics[0].Resource)
}

func TestMigrationContext_MixedDiagnostics(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Add a mix of diagnostics
	ctx.AddInfo("Starting migration", "Processing resource", "resource_a")
	ctx.AddWarning("Deprecated attribute", "Attribute 'old' is deprecated", "resource_a")
	ctx.AddError("Migration failed", "Unable to convert attribute", "resource_a")
	ctx.AddInfo("Completed migration", "Resource processed", "resource_a")
	ctx.AddWarning("Manual review needed", "Complex transformation", "resource_b")

	assert.Len(t, ctx.Diagnostics, 5)
	assert.Equal(t, 2, ctx.Metrics.WarningCount)

	// Count diagnostics by severity
	var infoCount, warningCount, errorCount int
	for _, diag := range ctx.Diagnostics {
		switch diag.Severity {
		case DiagnosticSeverityInfo:
			infoCount++
		case DiagnosticSeverityWarning:
			warningCount++
		case DiagnosticSeverityError:
			errorCount++
		}
	}

	assert.Equal(t, 2, infoCount)
	assert.Equal(t, 2, warningCount)
	assert.Equal(t, 1, errorCount)
}

func TestMigrationContext_EmptyDiagnostics(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Empty context should have no diagnostics
	assert.Empty(t, ctx.Diagnostics)
	assert.Equal(t, 0, ctx.Metrics.WarningCount)
	assert.Equal(t, 0, ctx.Metrics.TotalResources)
	assert.Equal(t, 0, ctx.Metrics.MigratedResources)
	assert.Equal(t, 0, ctx.Metrics.FailedResources)
	assert.Equal(t, 0, ctx.Metrics.ManualMigrationCount)
}

func TestMigrationContext_MetricsTracking(t *testing.T) {
	ctx := &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
	}

	// Simulate migration process
	ctx.Metrics.TotalResources = 10
	ctx.Metrics.MigratedResources = 7
	ctx.Metrics.FailedResources = 2
	ctx.Metrics.ManualMigrationCount = 1

	// Add corresponding diagnostics
	for i := 0; i < 7; i++ {
		ctx.AddInfo("Migrated", "Successfully migrated", "resource")
	}

	for i := 0; i < 2; i++ {
		ctx.AddError("Failed", "Migration failed", "resource")
	}

	ctx.AddWarning("Manual", "Manual migration needed", "resource")

	// Verify metrics
	assert.Equal(t, 10, ctx.Metrics.TotalResources)
	assert.Equal(t, 7, ctx.Metrics.MigratedResources)
	assert.Equal(t, 2, ctx.Metrics.FailedResources)
	assert.Equal(t, 1, ctx.Metrics.ManualMigrationCount)
	assert.Equal(t, 1, ctx.Metrics.WarningCount)
}

func TestMigrationContext_Options(t *testing.T) {
	tests := []struct {
		name    string
		options MigrationOptions
	}{
		{
			name: "dry run mode",
			options: MigrationOptions{
				DryRun:   true,
				Verbose:  false,
				Parallel: false,
				Workers:  1,
			},
		},
		{
			name: "verbose mode",
			options: MigrationOptions{
				DryRun:   false,
				Verbose:  true,
				Parallel: false,
				Workers:  1,
			},
		},
		{
			name: "parallel mode",
			options: MigrationOptions{
				DryRun:   false,
				Verbose:  false,
				Parallel: true,
				Workers:  4,
			},
		},
		{
			name: "all options enabled",
			options: MigrationOptions{
				DryRun:   true,
				Verbose:  true,
				Parallel: true,
				Workers:  8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &MigrationContext{
				Diagnostics: []Diagnostic{},
				Metrics:     &MigrationMetrics{},
				Options:     tt.options,
			}

			assert.Equal(t, tt.options.DryRun, ctx.Options.DryRun)
			assert.Equal(t, tt.options.Verbose, ctx.Options.Verbose)
			assert.Equal(t, tt.options.Parallel, ctx.Options.Parallel)
			assert.Equal(t, tt.options.Workers, ctx.Options.Workers)
		})
	}
}
