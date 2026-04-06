package byo_ip_prefix

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*ByoIPPrefixResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
// - v4 SDKv2: schema_version=0 with fields: id, account_id, prefix_id, description, advertisement
// - v5 production (v5.0-v5.15): schema_version=0 (Version was not yet set to 500)
// - v5 current: schema_version=500
//
// IMPORTANT: Both v4 and early v5 production state are at schema_version=0 but have
// incompatible formats:
// - v4: has "prefix_id" and "advertisement" fields
// - v5: has "asn", "cidr", and many computed fields; no "prefix_id" or "advertisement"
//
// Upgrade paths:
// 0: AMBIGUOUS - v4 SDKv2 format or v5 production format, detected at runtime
func (r *ByoIPPrefixResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Handle schema_version=0 -- ambiguous between v4 SDKv2 and v5 production.
		//
		// PriorSchema is nil so the framework skips pre-decoding req.State entirely.
		// Both paths are handled via req.RawState.JSON in UpgradeFromV0Ambiguous:
		// - v4 format (prefix_id present): full transform (rename fields, drop advertisement)
		// - v5 format (prefix_id absent): no-op re-decode with target schema
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV0Ambiguous,
		},
	}
}
