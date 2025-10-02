package migrations

import (
	_ "embed"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
)

//go:embed v4_to_v5.yaml
var v4ToV5Config []byte

// AccessApplicationMigration handles migrations for zero_trust_access_application resource
// All transformation logic is defined in the embedded YAML configuration
type AccessApplicationMigration struct {
	*internal.Migration
}

// NewAccessApplicationMigration creates a new access application migration
func NewAccessApplicationMigration(config []byte) (*AccessApplicationMigration, error) {
	migration, err := internal.NewMigration(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create base migration: %w", err)
	}

	return &AccessApplicationMigration{
		Migration: migration,
	}, nil
}

// NewV4ToV5Migration creates a new v4 to v5 migration for access_application
func NewV4ToV5Migration() (*AccessApplicationMigration, error) {
	return NewAccessApplicationMigration(v4ToV5Config)
}

// RegisterV4ToV5 registers the v4 to v5 migration
func RegisterV4ToV5(registry internal.MigrationRegistry) error {
	migration, err := NewV4ToV5Migration()
	if err != nil {
		return err
	}
	return registry.Register(migration)
}
