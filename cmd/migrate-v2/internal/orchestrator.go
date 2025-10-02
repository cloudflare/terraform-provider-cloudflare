package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Orchestrator manages the migration process
type Orchestrator struct {
	registry      MigrationRegistry
	options       MigrationOptions
	backupManager *BackupManager
	previewGen    *PreviewGenerator
}

// NewOrchestrator creates a new migration orchestrator
func NewOrchestrator(registry MigrationRegistry, options MigrationOptions) *Orchestrator {
	workDir := "."
	if options.WorkingDir != "" {
		workDir = options.WorkingDir
	}

	return &Orchestrator{
		registry:      registry,
		options:       options,
		backupManager: NewBackupManager(workDir),
		previewGen:    NewPreviewGenerator(registry),
	}
}

// MigrateDirectory migrates all .tf files in a directory
func (o *Orchestrator) MigrateDirectory(dir string, ctx *MigrationContext) error {
	// Handle state file if provided
	if ctx.StateFile != "" {
		if err := o.MigrateStateFile(ctx.StateFile, ctx); err != nil {
			return fmt.Errorf("failed to migrate state file: %w", err)
		}
	}
	// Initialize progress tracking
	if ctx.Progress != nil {
		ctx.Progress.Start(4, "Migration")
		defer ctx.Progress.Complete()
	}

	// Step 1: Find files
	if ctx.Progress != nil {
		ctx.Progress.StartStep("Finding Terraform files")
	}

	// Find all .tf files
	tfFiles, err := filepath.Glob(filepath.Join(dir, "*.tf"))
	if err != nil {
		return fmt.Errorf("failed to find .tf files: %w", err)
	}

	if len(tfFiles) == 0 {
		return fmt.Errorf("no .tf files found in directory: %s", dir)
	}

	if ctx.Progress != nil {
		ctx.Progress.CompleteStep(fmt.Sprintf("Found %d files", len(tfFiles)))
	}

	// Step 2: Create backup (if not dry-run)
	var backup *Backup
	if !ctx.DryRun && o.options.Backup {
		if ctx.Progress != nil {
			ctx.Progress.StartStep("Creating backup")
		}

		backup, err = o.backupManager.CreateBackup(ctx)
		if err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}

		if ctx.Progress != nil {
			ctx.Progress.CompleteStep(fmt.Sprintf("Backup created: %s", backup.ID))
		}
		ctx.Info("Created backup: %s", backup.ID)
	}

	// Step 3: Generate preview (if requested or dry-run)
	if ctx.DryRun || o.options.Preview {
		if ctx.Progress != nil {
			ctx.Progress.StartStep("Generating preview")
		}

		preview, err := o.generatePreview(tfFiles, ctx)
		if err != nil {
			ctx.Warning("Failed to generate preview", err.Error())
		} else {
			fmt.Println(preview.RenderDiff())

			if ctx.DryRun {
				if ctx.Progress != nil {
					ctx.Progress.CompleteStep("Preview generated (dry-run mode)")
				}
				return nil // Don't proceed with actual migration in dry-run
			}
		}

		if ctx.Progress != nil {
			ctx.Progress.CompleteStep("Preview generated")
		}
	}

	// Step 4: Execute migration
	if ctx.Progress != nil {
		ctx.Progress.StartStep("Migrating files")
	}

	var migrationErr error
	if o.options.Parallel && !ctx.DryRun {
		migrationErr = o.migrateFilesParallel(tfFiles, ctx)
	} else {
		migrationErr = o.migrateFilesSequential(tfFiles, ctx)
	}

	// Handle migration failure with auto-rollback
	if migrationErr != nil && backup != nil && o.options.AutoRollback {
		ctx.Error("Migration failed, initiating rollback...")
		if err := o.backupManager.Rollback(backup.ID); err != nil {
			return fmt.Errorf("migration failed and rollback failed: %w", err)
		}
		ctx.Info("Successfully rolled back to pre-migration state")
		return migrationErr
	}

	if ctx.Progress != nil {
		if migrationErr != nil {
			ctx.Progress.FailStep(migrationErr)
		} else {
			ctx.Progress.CompleteStep("All files migrated successfully")
		}
	}

	return migrationErr
}

func (o *Orchestrator) migrateFilesSequential(files []string, ctx *MigrationContext) error {
	for _, file := range files {
		if err := o.MigrateFile(file, ctx); err != nil {
			ctx.AddError("File migration failed", err.Error(), file)
			ctx.Metrics.FailedResources++
		}
	}
	return nil
}

func (o *Orchestrator) migrateFilesParallel(files []string, ctx *MigrationContext) error {
	var wg sync.WaitGroup
	fileChan := make(chan string, len(files))
	errorChan := make(chan error, len(files))

	// Start workers
	for i := 0; i < o.options.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range fileChan {
				if err := o.MigrateFile(file, ctx); err != nil {
					errorChan <- fmt.Errorf("error processing %s: %w", file, err)
				}
			}
		}()
	}

	// Send files to process
	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)

	// Wait for workers to finish
	wg.Wait()
	close(errorChan)

	// Collect errors
	for err := range errorChan {
		ctx.AddError("File migration failed", err.Error(), "")
		ctx.Metrics.FailedResources++
	}

	return nil
}

// MigrateStateFile migrates a Terraform state file
func (o *Orchestrator) MigrateStateFile(filename string, ctx *MigrationContext) error {
	if ctx.Options.Verbose {
		ctx.Log("Processing state file: %s", filename)
	}

	// Read state file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read state file: %w", err)
	}

	// Create backup if requested
	if o.options.Backup && !o.options.DryRun {
		backupName := fmt.Sprintf("%s.backup.%s", filename, time.Now().Format("20060102-150405"))
		if err := ioutil.WriteFile(backupName, content, 0644); err != nil {
			return fmt.Errorf("failed to create state backup: %w", err)
		}
	}

	// Parse JSON state
	var state map[string]interface{}
	if err := json.Unmarshal(content, &state); err != nil {
		return fmt.Errorf("failed to parse state JSON: %w", err)
	}

	modified := false

	// Process each resource in the state
	if resources, ok := state["resources"].([]interface{}); ok {
		for _, res := range resources {
			resource := res.(map[string]interface{})
			resourceType := resource["type"].(string)

			// Look up migration for this resource type
			migration, err := o.registry.Get(resourceType, ctx.SourceVersion, ctx.TargetVersion)
			if err != nil || migration == nil {
				continue
			}

			if ctx.Options.Verbose {
				ctx.Log("Migrating state for resource type: %s", resourceType)
			}

			// Process each instance
			if instances, ok := resource["instances"].([]interface{}); ok {
				for _, inst := range instances {
					instance := inst.(map[string]interface{})
					if attrs, ok := instance["attributes"].(map[string]interface{}); ok {
						// Run state migration
						if err := migration.MigrateState(attrs, ctx); err != nil {
							ctx.AddError("State migration failed", err.Error(), resourceType)
							continue
						}

						// Handle schema version update
						if schemaVersion, exists := attrs["schema_version"]; exists {
							instance["schema_version"] = schemaVersion
							delete(attrs, "schema_version")
						}

						modified = true
						ctx.Metrics.MigratedResources++
					}
				}
			}
		}
	}

	// Write back the migrated state if modified and not dry run
	if modified && !o.options.DryRun {
		output, err := json.MarshalIndent(state, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal state: %w", err)
		}

		if err := ioutil.WriteFile(filename, output, 0644); err != nil {
			return fmt.Errorf("failed to write migrated state: %w", err)
		}

		if ctx.Options.Verbose {
			ctx.Log("State file migrated successfully")
		}
	}

	return nil
}

// MigrateFile migrates a single configuration file
func (o *Orchestrator) MigrateFile(filename string, ctx *MigrationContext) error {
	if ctx.Options.Verbose {
		ctx.Log("Processing file: %s", filename)
	}

	// Read file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create backup if requested
	if !o.options.DryRun && ctx.Options.Verbose {
		backupName := fmt.Sprintf("%s.backup.%s", filename, time.Now().Format("20060102-150405"))
		if err := ioutil.WriteFile(backupName, content, 0644); err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}
	}

	// Parse HCL
	file, diags := hclwrite.ParseConfig(content, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	modified := false

	// Process each resource block
	for _, block := range file.Body().Blocks() {
		if block.Type() != "resource" {
			continue
		}

		labels := block.Labels()
		if len(labels) < 2 {
			continue
		}

		resourceType := labels[0]

		if ctx.Options.Verbose {
			ctx.Log("Found resource: %s", resourceType)
		}

		// Get migration for this resource (v4 -> v5 hardcoded for now)
		migration, err := o.registry.Get(resourceType, "v4", "v5")
		if err != nil {
			if ctx.Options.Verbose {
				ctx.Log("No migration available for %s: %v", resourceType, err)
			}
			continue // No migration available
		}

		ctx.Metrics.TotalResources++

		// Apply migration
		if err := migration.MigrateConfig(block, ctx); err != nil {
			ctx.AddError("Migration failed", err.Error(), resourceType)
			ctx.Metrics.FailedResources++
		} else {
			ctx.Metrics.MigratedResources++
			modified = true
		}
	}

	// Write modified file if not dry-run
	if modified && !o.options.DryRun {
		output := hclwrite.Format(file.Bytes())
		if err := ioutil.WriteFile(filename, output, 0644); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}
	}

	return nil
}

// generatePreview generates a migration preview for the given files
func (o *Orchestrator) generatePreview(files []string, ctx *MigrationContext) (*MigrationPreview, error) {
	var allBlocks []*hclwrite.Block

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			ctx.AddWarning("Failed to read file for preview", err.Error(), file)
			continue
		}

		hclFile, diags := hclwrite.ParseConfig(content, file, hcl.InitialPos)
		if diags.HasErrors() {
			ctx.AddWarning("Failed to parse file for preview", diags.Error(), file)
			continue
		}

		allBlocks = append(allBlocks, hclFile.Body().Blocks()...)
	}

	if len(allBlocks) == 0 {
		return nil, fmt.Errorf("no valid blocks found for preview")
	}

	return o.previewGen.GeneratePreview(allBlocks, ctx.SourceVersion, ctx.TargetVersion)
}
