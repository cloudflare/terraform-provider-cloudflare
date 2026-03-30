// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZoneDNSSECMigrationModelSchemaParity verifies that TargetZoneDNSSECModel
// (used in UpgradeStateFrom0To500 → resp.State.Set) stays in sync with the live ResourceSchema.
func TestZoneDNSSECMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetZoneDNSSECModel)(nil)
	schema := zone_dnssec.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
