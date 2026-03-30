// This file is intentionally NOT code-generated.
// It ensures the migration target model stays in sync with the live resource schema.
// If a field is added to the live model/schema without being added to the migration
// Target struct, this test will fail, catching the drift in CI before it causes a
// runtime panic.

package v500_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_config"
	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_config/migration/v500"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

// TestZeroTrustTunnelCloudflaredConfigMigrationModelSchemaParity verifies that TargetV5TunnelConfigModel
// (used in UpgradeFromV0, UpgradeFromV4, UpgradeFromV5 → resp.State.Set, MoveStateFromTunnelConfig → resp.TargetState.Set) stays in sync with the live ResourceSchema.
func TestZeroTrustTunnelCloudflaredConfigMigrationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*v500.TargetV5TunnelConfigModel)(nil)
	schema := zero_trust_tunnel_cloudflared_config.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateMigrationModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
