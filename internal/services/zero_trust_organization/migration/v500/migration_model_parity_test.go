// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZeroTrustOrganizationMigrationModelSchemaParity verifies that TargetZeroTrustOrganizationModel
// (used in MoveState → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestZeroTrustOrganizationMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetZeroTrustOrganizationModel)(nil)
	schema := zero_trust_organization.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}

// TestZeroTrustOrganizationMigrationSourceModelSchemaParity verifies that
// SourceCloudflareAccessOrganizationModel (used in MoveState and
// UpgradeFromLegacyV0 to decode v4 state) declares every attribute that
// SourceCloudflareAccessOrganizationSchema defines. A missing field causes the
// Plugin Framework to return a Value Conversion Error when reading v4 state,
// blocking terraform plan after migration.
//
// Regression test for ESCALATION-11029: warp_auth_non_browser_401 was added to
// the source schema but omitted from the source model struct.
func TestZeroTrustOrganizationMigrationSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.SourceCloudflareAccessOrganizationModel)(nil)
	schema := v500.SourceCloudflareAccessOrganizationSchema()
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
