// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestLogpushJobMigrationModelSchemaParity verifies that TargetLogpushJobModel
// (used in UpgradeFromVersion0, UpgradeFromV4 → resp.State.Set) stays in sync
// with the live ResourceSchema. Adding a field to the live schema without updating
// TargetLogpushJobModel will cause this test to fail.
func TestLogpushJobMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetLogpushJobModel)(nil)
	schema := logpush_job.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
