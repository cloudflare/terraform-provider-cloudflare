// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
	"context"
	"os"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneDNSSECResource)(nil)

func (r *ZoneDNSSECResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
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
	sourceSchema := v500.SourceCloudflareZoneDNSSECSchema()

	return map[int64]resource.StateUpgrader{
		// Upgrade from v4 (SDKv2, version 0) to v5 (Plugin Framework, version 500)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeStateFrom0To500,
		},
	}
}
