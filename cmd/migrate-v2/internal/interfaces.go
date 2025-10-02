package internal

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ResourceMigration defines the interface for resource-specific migrations
type ResourceMigration interface {
	// ResourceType returns the resource type this migration handles (e.g., "cloudflare_zero_trust_access_application")
	ResourceType() string

	// SourceVersion returns the source version (e.g., "v4")
	SourceVersion() string

	// TargetVersion returns the target version (e.g., "v5")
	TargetVersion() string

	// MigrateConfig transforms a resource's configuration
	MigrateConfig(block *hclwrite.Block, ctx *MigrationContext) error

	// MigrateState transforms a resource's state
	MigrateState(state map[string]interface{}, ctx *MigrationContext) error

	// Validate checks if the migration can be applied to the given resource
	Validate(block *hclwrite.Block) error
}

// Diagnostic represents a warning or error during migration
type Diagnostic struct {
	Severity DiagnosticSeverity
	Summary  string
	Detail   string
	Resource string
	Line     int
	Column   int
}

// DiagnosticSeverity represents the severity of a diagnostic
type DiagnosticSeverity int

const (
	DiagnosticSeverityError DiagnosticSeverity = iota
	DiagnosticSeverityWarning
	DiagnosticSeverityInfo
)

// MigrationMetrics tracks statistics about the migration
type MigrationMetrics struct {
	TotalResources       int
	MigratedResources    int
	FailedResources      int
	WarningCount         int
	ManualMigrationCount int
}

// MigrationOptions contains user-specified options for migration
type MigrationOptions struct {
	DryRun       bool
	Verbose      bool
	Parallel     bool
	Workers      int
	Preview      bool
	Backup       bool
	AutoRollback bool
	WorkingDir   string
}

// MigrationRegistry maintains a registry of available migrations
type MigrationRegistry interface {
	// Register adds a migration to the registry
	Register(migration ResourceMigration) error

	// Get retrieves migrations for a specific resource type and version transition
	Get(resourceType, sourceVersion, targetVersion string) (ResourceMigration, error)

	// GetPath finds a migration path from source to target version
	GetPath(resourceType, sourceVersion, targetVersion string) ([]ResourceMigration, error)

	// ListAvailable lists all registered migrations
	ListAvailable() []MigrationInfo
}

// MigrationInfo provides information about a registered migration
type MigrationInfo struct {
	ResourceType  string
	SourceVersion string
	TargetVersion string
	Description   string
}
