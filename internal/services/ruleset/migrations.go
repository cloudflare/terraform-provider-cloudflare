// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*RulesetResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
//   - v4 Plugin Framework: schema_version=0 (may contain rule version/time metadata)
//     or 1 (after retiring that metadata)
//   - v5 production (v5.0-v5.18): schema_version=1 (GetSchemaVersion(1, 500) returned 1)
//   - v5 current: schema_version=500
//
// IMPORTANT: Both v4 and v5 state can be at schema_version=0 or 1, but they have
// incompatible formats:
// - v4: action_parameters stored as ARRAY (ListNestedBlock)
// - v5: action_parameters stored as OBJECT (SingleNestedAttribute)
//
// Upgrade paths:
// 0, 1: AMBIGUOUS - v4 format (array) or v5 format (object), detected at runtime
func (r *RulesetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		// Neither version identifies the format. Skip framework pre-decoding so
		// v4 list blocks are not rejected by the v5 object schema (#7135).
		0: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromAmbiguous,
		},
		1: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromAmbiguous,
		},
	}
}
