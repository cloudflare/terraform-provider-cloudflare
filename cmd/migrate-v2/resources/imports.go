package resources

// This file imports all resource packages to trigger their init() functions,
// which automatically register their migrations.
// When adding a new resource migration, simply add an import here.

import (
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/internal"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/access_application"
	"github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/argo"
	
	// Add new resource imports here as they are implemented:
	// "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/zone"
	// "github.com/cloudflare/terraform-provider-cloudflare/cmd/migrate-v2/resources/record"
)

// RegisterAll registers all resource migrations with the provided registry.
// When adding a new resource, add its RegisterMigrations call here.
func RegisterAll(registry internal.MigrationRegistry) error {
	// Register all resource migrations
	if err := access_application.RegisterMigrations(registry); err != nil {
		return err
	}
	
	if err := argo.RegisterMigrations(registry); err != nil {
		return err
	}
	
	// Add new resources here:
	// if err := zone.RegisterMigrations(registry); err != nil {
	//     return err
	// }
	
	return nil
}