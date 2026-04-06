// File generated to ensure migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration Target struct,
// this test will fail, catching the drift in CI before it causes a runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_hostname_certificate"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_hostname_certificate/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestAuthenticatedOriginPullsHostnameCertificateMigrationModelSchemaParity verifies that V5Model
// (used in UpgradeFromV0 → resp.State.Set, MoveState → resp.TargetState.Set) stays in sync
// with the live ResourceSchema. Adding a field to the live schema without updating
// V5Model will cause this test to fail.
func TestAuthenticatedOriginPullsHostnameCertificateMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.V5Model)(nil)
	schema := authenticated_origin_pulls_hostname_certificate.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
