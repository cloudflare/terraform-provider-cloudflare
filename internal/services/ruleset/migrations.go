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
// - v4 Plugin Framework: schema_version=1 (after internal V0->V1 ratelimit rename)
// - v5 production (v5.0-v5.18): schema_version=1 (GetSchemaVersion(1, 500) returned 1)
// - v5 current: schema_version=500
//
// IMPORTANT: Both v4 and v5 production state are at schema_version=1, but they have
// incompatible formats:
// - v4: action_parameters stored as ARRAY (ListNestedBlock)
// - v5: action_parameters stored as OBJECT (SingleNestedAttribute)
//
// Upgrade paths:
// 0: v5 fresh resources (schema_version=0) -> no-op
// 1: AMBIGUOUS - v4 format (array) or v5 production format (object), detected at runtime
func (r *RulesetResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	v5Schema := ResourceSchema(ctx)

	return map[int64]resource.StateUpgrader{
		// Handle fresh v5 resources (version 0 -> 500): no-op.
		0: {
			PriorSchema:   &v5Schema,
			StateUpgrader: v500.UpgradeFromV5,
		},

		// Handle schema_version=1 -- ambiguous between v4 and v5 production.
		//
		// PriorSchema is nil so the framework skips pre-decoding req.State entirely.
		// Both paths are handled via req.RawState.JSON in UpgradeFromV1Ambiguous:
		// - v4 format (action_parameters is array): full transform via raw state
		// - v5 format (action_parameters is object): no-op re-decode with target schema
		1: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV1Ambiguous,
		},
	}
}
