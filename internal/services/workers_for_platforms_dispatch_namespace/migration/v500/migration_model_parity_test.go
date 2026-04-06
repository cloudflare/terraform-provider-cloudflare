// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestWorkersForPlatformsDispatchNamespaceMigrationModelSchemaParity verifies that TargetWorkersForPlatformsDispatchNamespaceModel
// (used in UpgradeFromV0 → resp.State.Set, MoveFromWorkersForPlatformsNamespace → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestWorkersForPlatformsDispatchNamespaceMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetWorkersForPlatformsDispatchNamespaceModel)(nil)
	schema := workers_for_platforms_dispatch_namespace.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
