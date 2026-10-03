// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_block_sender"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_block_sender/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestEmailSecurityBlockSenderMigrationModelSchemaParity verifies that TargetModel
// (used in the v500 upgrader → resp.State.Set) stays in sync with the live
// ResourceSchema. Adding a field to the live schema without updating TargetModel
// will cause this test to fail.
func TestEmailSecurityBlockSenderMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetModel)(nil)
	schema := email_security_block_sender.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
