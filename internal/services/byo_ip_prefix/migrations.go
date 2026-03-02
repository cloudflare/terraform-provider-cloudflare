package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ByoIPPrefixResource)(nil)

func (r *ByoIPPrefixResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceCloudflareByoIPPrefixSchema()
	return map[int64]resource.StateUpgrader{
		// Handle state from the v4 SDKv2 provider (schema_version=0).
		// Transforms field renames (prefix_id→id), drops removed fields (advertisement),
		// and initializes new v5 fields (asn, cidr, and computed fields) to null.
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},
	}
}
