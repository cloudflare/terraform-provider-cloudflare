package spectrum_application

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_application/migration/v500"
)

var _ resource.ResourceWithUpgradeState = (*SpectrumApplicationResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
//   - v4 (SDKv2): schema_version=0, dns/origin_dns/edge_ips/origin_port_range as JSON arrays
//   - early v5 (v5.0–v5.7): schema_version=0, same fields as JSON objects
//   - v5 production (v5.0–v5.18, with stepping stone): schema_version=1
//   - v5 current: schema_version=500
//
// Upgrade paths:
//
//  1. schema_version=0 → 500: AMBIGUOUS between v4 (list shape) and early v5
//     (object shape). PriorSchema MUST be nil — if we set it to the v4 source
//     schema, the Plugin Framework rejects early-v5 state pre-handler with
//     `AttributeName("dns"): invalid JSON, expected "[", got "{"` (see #7098).
//     The handler reads req.RawState.JSON itself and routes accordingly.
//
//  2. schema_version=1 → 500: No-op for v5 production state.
func (r *SpectrumApplicationResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	targetSchema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// schema_version=0 — ambiguous between v4 SDKv2 and early v5; format
		// detection happens in the handler against req.RawState.JSON.
		0: {
			StateUpgrader: v500.UpgradeFromV0Ambiguous,
		},

		// Handle state from v5 provider with version=1 (production dormant state)
		// No-op upgrade that just bumps the version to 500
		1: {
			PriorSchema:   &targetSchema,
			StateUpgrader: v500.UpgradeFromV1,
		},
	}
}
