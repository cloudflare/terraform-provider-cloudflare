// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZeroTrustDeviceDefaultProfileMigrationModelSchemaParity verifies that TargetDefaultProfileModel
// (used in UpgradeFromV4 → resp.State.Set, MoveDeviceProfilesToDefaultProfile → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestZeroTrustDeviceDefaultProfileMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetDefaultProfileModel)(nil)
	schema := zero_trust_device_default_profile.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
