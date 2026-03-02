// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_dnssec

import (
	"context"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ZoneDNSSECResource)(nil)

func (r *ZoneDNSSECResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	sourceSchema := v500.SourceCloudflareZoneDNSSECSchema()

	return map[int64]resource.StateUpgrader{
		// Upgrade from v4 (SDKv2, version 0) to v5 (Plugin Framework, version 500)
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeStateFrom0To500,
		},
	}
}
