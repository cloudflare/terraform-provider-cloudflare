// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestByoIPPrefixMigrationModelSchemaParity verifies that TargetByoIPPrefixModel
// (used in UpgradeFromV0Ambiguous, UpgradeFromV0 → resp.State.Set) stays in sync
// with the live ResourceSchema. Adding a field to the live schema without updating
// TargetByoIPPrefixModel will cause this test to fail.
func TestByoIPPrefixMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetByoIPPrefixModel)(nil)
	schema := byo_ip_prefix.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
