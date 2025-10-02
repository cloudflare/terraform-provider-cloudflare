package internal

import (
	"fmt"
	"io"
	"os"
)

// MigrationContext provides context for migrations
type MigrationContext struct {
	// Version information
	SourceVersion string
	TargetVersion string

	// Options
	DryRun      bool
	Verbose     bool
	AutoApprove bool

	// Paths
	WorkingDir string
	StateFile  string

	// Diagnostics collects warnings and errors during migration
	Diagnostics []Diagnostic

	// Metrics tracks migration statistics
	Metrics *MigrationMetrics

	// Options contains user-specified migration options
	Options MigrationOptions

	// Progress tracking
	Progress *ProgressTracker

	// Output writer (defaults to os.Stdout)
	Output io.Writer
}

// NewMigrationContext creates a new migration context with defaults
func NewMigrationContext() *MigrationContext {
	return &MigrationContext{
		Diagnostics: []Diagnostic{},
		Metrics:     &MigrationMetrics{},
		Progress:    NewProgressTracker(),
		Output:      os.Stdout,
	}
}

// WithDryRun sets dry-run mode
func (ctx *MigrationContext) WithDryRun(dryRun bool) *MigrationContext {
	ctx.DryRun = dryRun
	return ctx
}

// WithVerbose sets verbose mode
func (ctx *MigrationContext) WithVerbose(verbose bool) *MigrationContext {
	ctx.Verbose = verbose
	return ctx
}

// Log writes a log message if verbose mode is enabled
func (ctx *MigrationContext) Log(format string, args ...interface{}) {
	if ctx.Verbose {
		fmt.Fprintf(ctx.Output, "[DEBUG] "+format+"\n", args...)
	}
}

// Info writes an info message
func (ctx *MigrationContext) Info(format string, args ...interface{}) {
	fmt.Fprintf(ctx.Output, "[INFO] "+format+"\n", args...)
}

// AddDiagnostic adds a diagnostic to the context
func (ctx *MigrationContext) AddDiagnostic(severity DiagnosticSeverity, summary, detail, resource string) {
	ctx.Diagnostics = append(ctx.Diagnostics, Diagnostic{
		Severity: severity,
		Summary:  summary,
		Detail:   detail,
		Resource: resource,
	})
}

// AddInfo adds an info diagnostic
func (ctx *MigrationContext) AddInfo(summary, detail, resource string) {
	ctx.AddDiagnostic(DiagnosticSeverityInfo, summary, detail, resource)
}

// AddWarning adds a warning diagnostic
func (ctx *MigrationContext) AddWarning(summary, detail, resource string) {
	ctx.AddDiagnostic(DiagnosticSeverityWarning, summary, detail, resource)
	ctx.Metrics.WarningCount++
}

// AddError adds an error diagnostic
func (ctx *MigrationContext) AddError(summary, detail, resource string) {
	ctx.AddDiagnostic(DiagnosticSeverityError, summary, detail, resource)
}

// Warning writes a warning message with detail
func (ctx *MigrationContext) Warning(summary string, detail string) {
	ctx.AddWarning(summary, detail, "")
	if ctx.Progress != nil {
		ctx.Progress.Warning(fmt.Sprintf("%s: %s", summary, detail))
	} else {
		fmt.Fprintf(ctx.Output, "[WARNING] %s: %s\n", summary, detail)
	}
}

// Error writes an error message
func (ctx *MigrationContext) Error(message string) {
	if ctx.Progress != nil {
		ctx.Progress.Error(message)
	} else {
		fmt.Fprintf(ctx.Output, "[ERROR] %s\n", message)
	}
}
