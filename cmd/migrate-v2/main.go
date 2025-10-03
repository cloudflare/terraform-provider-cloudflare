package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources"
)

var (
	configDir      string
	stateFile      string
	sourceVersion  string
	targetVersion  string
	resourceTarget string
	dryRun         bool
	preview        bool
	backup         bool
	parallel       bool
	workers        int
	verbose        bool
	reportFile     string
)

func init() {
	flag.StringVar(&configDir, "config", "", "Directory containing Terraform configuration files (required)")
	flag.StringVar(&stateFile, "state", "", "Path to Terraform state file (optional)")
	flag.StringVar(&sourceVersion, "from", "v4", "Source provider version")
	flag.StringVar(&targetVersion, "to", "v5", "Target provider version")
	flag.StringVar(&resourceTarget, "target", "", "Target specific resource type")
	flag.BoolVar(&dryRun, "dry-run", false, "Preview changes without applying them")
	flag.BoolVar(&preview, "preview", false, "Show detailed preview of changes")
	flag.BoolVar(&backup, "backup", true, "Create backups before modifying files")
	flag.BoolVar(&parallel, "parallel", false, "Process files in parallel")
	flag.IntVar(&workers, "workers", 4, "Number of parallel workers")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.StringVar(&reportFile, "report", "", "Write migration report to file")
}

func main() {
	// Custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Terraform Provider Cloudflare Migration Tool (v4 → v5)\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Preview changes without applying\n")
		fmt.Fprintf(os.Stderr, "  %s -config ./terraform -dry-run -preview\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Migrate configuration files\n")
		fmt.Fprintf(os.Stderr, "  %s -config ./terraform\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Migrate with state file\n")
		fmt.Fprintf(os.Stderr, "  %s -config ./terraform -state terraform.tfstate\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Migrate specific resource type\n")
		fmt.Fprintf(os.Stderr, "  %s -config ./terraform -target cloudflare_zero_trust_access_application\n\n", os.Args[0])
	}

	flag.Parse()

	// Show help if no arguments provided or if -h/-help was used
	if len(os.Args) == 1 || configDir == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Initialize migration registry
	registry := internal.NewDefaultRegistry()

	// Automatically register all resource migrations
	// This uses the init() functions in each resource package
	if err := resources.RegisterAll(registry); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register migrations: %v\n", err)
		os.Exit(1)
	}

	// Create migration context
	ctx := &internal.MigrationContext{
		SourceVersion: sourceVersion,
		TargetVersion: targetVersion,
		DryRun:        dryRun,
		Verbose:       verbose,
		WorkingDir:    configDir,
		StateFile:     stateFile,
		Diagnostics:   []internal.Diagnostic{},
		Metrics:       &internal.MigrationMetrics{},
		Options: internal.MigrationOptions{
			DryRun:       dryRun,
			Verbose:      verbose,
			Parallel:     parallel,
			Workers:      workers,
			Preview:      preview,
			Backup:       backup,
			AutoRollback: false, // Could be added as a flag later
			WorkingDir:   configDir,
		},
		Progress: internal.NewProgressTracker(),
		Output:   os.Stdout,
	}

	// Configure progress tracker
	ctx.Progress.SetVerbose(verbose)

	// Create orchestrator
	orchestrator := internal.NewOrchestrator(registry, ctx.Options)

	// Process configuration files
	if err := orchestrator.MigrateDirectory(configDir, ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing configuration: %v\n", err)
		os.Exit(1)
	}

	// Generate and display report
	report := generateReport(ctx)
	displayReport(report)

	// Write report to file if requested
	if reportFile != "" {
		if err := writeReportToFile(report, reportFile); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing report: %v\n", err)
		}
	}

	// Exit with appropriate code
	if ctx.Metrics.FailedResources > 0 {
		os.Exit(1)
	}
}

// MigrationReport represents the migration report
type MigrationReport struct {
	Timestamp         time.Time
	SourceVersion     string
	TargetVersion     string
	TotalResources    int
	MigratedResources int
	FailedResources   int
	Warnings          int
	Diagnostics       []internal.Diagnostic
	DryRun            bool
}

func generateReport(ctx *internal.MigrationContext) MigrationReport {
	report := MigrationReport{
		Timestamp:         time.Now(),
		SourceVersion:     sourceVersion,
		TargetVersion:     targetVersion,
		TotalResources:    ctx.Metrics.TotalResources,
		MigratedResources: ctx.Metrics.MigratedResources,
		FailedResources:   ctx.Metrics.FailedResources,
		Warnings:          ctx.Metrics.WarningCount,
		Diagnostics:       ctx.Diagnostics,
		DryRun:            dryRun,
	}

	return report
}

func displayReport(report MigrationReport) {
	fmt.Println("\n=== Migration Report ===")
	fmt.Printf("Timestamp: %s\n", report.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("Migration: %s -> %s\n", report.SourceVersion, report.TargetVersion)

	if report.DryRun {
		fmt.Println("Mode: DRY RUN (no changes applied)")
	} else {
		fmt.Println("Mode: APPLIED")
	}

	fmt.Println("\nStatistics:")
	fmt.Printf("  Total resources found: %d\n", report.TotalResources)
	fmt.Printf("  Successfully migrated: %d\n", report.MigratedResources)
	fmt.Printf("  Failed migrations: %d\n", report.FailedResources)
	fmt.Printf("  Warnings: %d\n", report.Warnings)

	if len(report.Diagnostics) > 0 {
		fmt.Println("\nDiagnostics:")
		for _, diag := range report.Diagnostics {
			severity := "INFO"
			switch diag.Severity {
			case internal.DiagnosticSeverityError:
				severity = "ERROR"
			case internal.DiagnosticSeverityWarning:
				severity = "WARNING"
			}
			fmt.Printf("  [%s] %s: %s\n", severity, diag.Summary, diag.Detail)
		}
	}

	if report.FailedResources > 0 {
		fmt.Println("\n⚠️  Some migrations failed. Please review the errors above.")
	} else if report.MigratedResources > 0 {
		fmt.Println("\n✅ Migration completed successfully!")
	} else {
		fmt.Println("\nℹ️  No resources required migration.")
	}
}

func writeReportToFile(report MigrationReport, filename string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write report file: %w", err)
	}

	fmt.Printf("Report written to: %s\n", filename)
	return nil
}
