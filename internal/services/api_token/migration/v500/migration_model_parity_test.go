// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestAPITokenMigrationModelSchemaParity verifies that TargetAPITokenModel
// (used in UpgradeFromV4 → resp.State.Set) stays in sync
// with the live ResourceSchema. Adding a field to the live schema without updating
// TargetAPITokenModel will cause this test to fail.
func TestAPITokenMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetAPITokenModel)(nil)
	schema := api_token.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
