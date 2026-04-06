// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestArgoTieredCachingMigrationModelSchemaParity verifies that TargetArgoTieredCachingModel
// (used in UpgradeArgoToTieredCaching → resp.State.Set, MoveArgoToTieredCaching → resp.TargetState.Set)
// stays in sync with the live ResourceSchema. Adding a field to the live schema without updating
// TargetArgoTieredCachingModel will cause this test to fail.
func TestArgoTieredCachingMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetArgoTieredCachingModel)(nil)
	schema := argo_tiered_caching.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
