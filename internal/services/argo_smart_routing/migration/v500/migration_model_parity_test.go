// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_smart_routing"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_smart_routing/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestArgoSmartRoutingMigrationModelSchemaParity verifies that TargetArgoSmartRoutingModel
// (used in UpgradeArgoToSmartRouting → resp.State.Set, MoveArgoToSmartRouting → resp.TargetState.Set)
// stays in sync with the live ResourceSchema. Adding a field to the live schema without updating
// TargetArgoSmartRoutingModel will cause this test to fail.
func TestArgoSmartRoutingMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetArgoSmartRoutingModel)(nil)
	schema := argo_smart_routing.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
