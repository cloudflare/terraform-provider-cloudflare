package zero_trust_access_mtls_hostname_settings

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_hostname_settings/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessMTLSHostnameSettingsResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
func (r *ZeroTrustAccessMTLSHostnameSettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		// Production mode: preserve existing upgraders only
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// Test mode (TF_MIG_TEST=1): full StateUpgrader migration
	v4Schema := v500.SourceMTLSHostnameSettingsSchema()

	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV0,
		},
		// Handle state from v5 Plugin Framework provider (version=1)
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
