// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_rule"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_rule/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZeroTrustDevicePostureRuleMigrationModelSchemaParity verifies that TargetDevicePostureRuleModel
// (used in UpgradeFromVersion0, UpgradeFromV0 → resp.State.Set, MoveState → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestZeroTrustDevicePostureRuleMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetDevicePostureRuleModel)(nil)
	schema := zero_trust_device_posture_rule.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
