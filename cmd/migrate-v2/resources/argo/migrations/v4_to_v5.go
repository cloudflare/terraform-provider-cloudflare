package migrations

import (
	_ "embed"
	"fmt"

	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

// ArgoMigration extends the base migration to handle resource splitting warnings
type ArgoMigration struct {
	*internal.Migration
}

// ResourceType returns the resource type
func (m *ArgoMigration) ResourceType() string {
	return m.Migration.ResourceType()
}

// SourceVersion returns the source version
func (m *ArgoMigration) SourceVersion() string {
	return m.Migration.SourceVersion()
}

// TargetVersion returns the target version
func (m *ArgoMigration) TargetVersion() string {
	return m.Migration.TargetVersion()
}

// MigrateConfig delegates to base migration
func (m *ArgoMigration) MigrateConfig(block *hclwrite.Block, ctx *internal.MigrationContext) error {
	// Since we handle resource splitting in TransformFile,
	// we just delegate to the base migration for any remaining config transformations
	return m.Migration.MigrateConfig(block, ctx)
}

// MigrateState delegates to base migration
func (m *ArgoMigration) MigrateState(state map[string]interface{}, ctx *internal.MigrationContext) error {
	return m.Migration.MigrateState(state, ctx)
}

// RequiresFileTransformation returns true since argo needs file-level processing
func (m *ArgoMigration) RequiresFileTransformation() bool {
	return true
}

// RegisterV4ToV5 registers the v4 to v5 migration for argo
func RegisterV4ToV5(registry internal.MigrationRegistry) error {
	baseMigration, err := internal.NewMigration(v4ToV5Config)
	if err != nil {
		return fmt.Errorf("failed to create argo migration: %w", err)
	}

	// Wrap with ArgoMigration to add custom behavior
	migration := &ArgoMigration{
		Migration: baseMigration,
	}

	return registry.Register(migration)
}
