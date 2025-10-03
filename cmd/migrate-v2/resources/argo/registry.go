package argo

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/argo/migrations"
)

// RegisterMigrations registers all migrations for the argo resource
func RegisterMigrations(registry internal.MigrationRegistry) error {
	// Register v4 to v5 migration
	if err := migrations.RegisterV4ToV5(registry); err != nil {
		return err
	}

	return nil
}