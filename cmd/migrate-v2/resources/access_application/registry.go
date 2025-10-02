package access_application

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/access_application/migrations"
)

// RegisterMigrations registers all migrations for the access_application resource
func RegisterMigrations(registry internal.MigrationRegistry) error {
	// Register v4 to v5 migration
	if err := migrations.RegisterV4ToV5(registry); err != nil {
		return err
	}

	// Future: Register v5 to v6 migration when implemented
	// if err := migrations.RegisterV5ToV6(registry); err != nil {
	//     return err
	// }

	return nil
}
