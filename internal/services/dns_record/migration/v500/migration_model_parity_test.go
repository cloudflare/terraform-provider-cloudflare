// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestDNSRecordMigrationModelSchemaParity verifies that TargetDNSRecordModel
// (used in UpgradeFromV0, UpgradeFromLegacyV3 → resp.State.Set, MoveState → resp.TargetState.Set)
// stays in sync with the live ResourceSchema. Adding a field to the live schema without updating
// TargetDNSRecordModel will cause this test to fail.
func TestDNSRecordMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetDNSRecordModel)(nil)
	schema := dns_record.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
