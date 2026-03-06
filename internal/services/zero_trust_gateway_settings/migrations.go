// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_settings

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZeroTrustGatewaySettingsResource)(nil)
var _ resource.ResourceWithMoveState = (*ZeroTrustGatewaySettingsResource)(nil)

// MoveState registers state movers for the resource rename:
// cloudflare_teams_account → cloudflare_zero_trust_gateway_settings.
//
// This enables Terraform 1.8+ `moved` blocks to automatically trigger state migration
// when renaming resources from the old type to the new type.
func (r *ZeroTrustGatewaySettingsResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceV4ZeroTrustGatewaySettingsSchema()

	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveFromCloudflareTeamsAccount,
		},
	}
}

// UpgradeState registers state upgraders for schema version changes.
//
// Version history:
//   - 0: v4 SDKv2 state (full transformation needed)
//   - 1: Dormant production v5 state (GetSchemaVersion returns 1 normally)
//   - 500: Active migration version (GetSchemaVersion returns 500 when TF_MIG_TEST=1)
func (r *ZeroTrustGatewaySettingsResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	if os.Getenv("TF_MIG_TEST") == "" {
		return map[int64]resource.StateUpgrader{
			0: {
				PriorSchema: &targetSchema,
				StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
					resp.State.Raw = req.State.Raw
				},
			},
		}
	}

	// v4 source schema for parsing old SDKv2 state (schema_version=0)
	v4Schema := v500.SourceV4ZeroTrustGatewaySettingsSchema()

	// v5 schema at version=1 for no-op pass-through upgrader
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle state from v4 SDKv2 provider (schema_version=0)
		// Uses v4 PriorSchema to parse, then transforms flat structure to v5 nested
		0: {
			PriorSchema:   &v4Schema,
			StateUpgrader: v500.UpgradeFromV4,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// No-op version bump to 500; only triggered when TF_MIG_TEST=1
		1: {
			PriorSchema:   &v5SchemaVersion1,
			StateUpgrader: v500.UpgradeFromV5,
		},
	}
}
