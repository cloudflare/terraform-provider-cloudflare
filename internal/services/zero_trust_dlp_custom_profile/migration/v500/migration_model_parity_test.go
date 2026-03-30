// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_profile"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_profile/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZeroTrustDLPCustomProfileMigrationModelSchemaParity verifies that TargetZeroTrustDLPCustomProfileModel
// (used in UpgradeFromVersion0, UpgradeFromV0 → resp.State.Set, MoveState → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestZeroTrustDLPCustomProfileMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetZeroTrustDLPCustomProfileModel)(nil)
	schema := zero_trust_dlp_custom_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
