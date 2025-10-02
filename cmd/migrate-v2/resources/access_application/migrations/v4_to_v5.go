package migrations

import (
	_ "embed"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/core"
)

//go:embed v4_to_v5.yaml
var migrationConfig []byte

// V4ToV5Migration handles the v4 to v5 migration for zero_trust_access_application
// All transformation logic is defined in the embedded YAML configuration
type V4ToV5Migration struct {
	*core.BaseMigration
}

// NewV4ToV5Migration creates a new v4 to v5 migration
func NewV4ToV5Migration() (*V4ToV5Migration, error) {
	base, err := core.NewBaseMigration(migrationConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create base migration: %w", err)
	}

	return &V4ToV5Migration{
		BaseMigration: base,
	}, nil
}

// RegisterV4ToV5 registers the v4 to v5 migration
func RegisterV4ToV5(registry core.MigrationRegistry) error {
	migration, err := NewV4ToV5Migration()
	if err != nil {
		return err
	}
	return registry.Register(migration)
}