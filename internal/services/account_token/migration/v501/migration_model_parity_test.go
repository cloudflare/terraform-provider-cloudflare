// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v501_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token"
	v501 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token/migration/v501"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestAccountTokenV501MigrationModelSchemaParity verifies that AccountTokenModelV501
// (used in UpgradeFromV500 → resp.State.Set) stays in sync
// with the live ResourceSchema. Adding a field to the live schema without updating
// AccountTokenModelV501 will cause this test to fail.
func TestAccountTokenV501MigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v501.AccountTokenModelV501)(nil)
	schema := account_token.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
