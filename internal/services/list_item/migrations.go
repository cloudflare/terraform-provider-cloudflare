// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package list_item

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*ListItemResource)(nil)

// UpgradeState registers state upgraders for schema version changes.
//
// Schema version history:
//   - v4 SDKv2: schema_version=0
//   - v4.52.5 Plugin Framework: schema_version=1 (hostname/redirect as ListNestedBlock)
//   - v5 production (v5.0-v5.18): schema_version=1 (hostname/redirect as SingleNestedAttribute)
//   - v5 current: schema_version=500
//
// IMPORTANT: Both v4.52.5 and v5 production state are at schema_version=1, but they have
// incompatible formats:
//   - v4.52.5: hostname/redirect stored as ARRAY (ListNestedBlock)
//   - v5: hostname/redirect stored as OBJECT (SingleNestedAttribute)
//
// Upgrade paths:
//
//	0: v4 SDKv2 -> full transform
//	1: AMBIGUOUS - v4.52.5 format (array) or v5 production format (object), detected at runtime
func (r *ListItemResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	sourceSchema := v500.SourceListItemSchema()

	return map[int64]resource.StateUpgrader{
		// Handle upgrades from v4 SDKv2 (schema_version=0) to v500
		0: {
			PriorSchema:   &sourceSchema,
			StateUpgrader: v500.UpgradeFromV0,
		},

		// Handle schema_version=1 -- ambiguous between v4.52.5 and v5 production.
		//
		// PriorSchema is nil so the framework skips pre-decoding req.State entirely.
		// Both paths are handled via req.RawState.JSON in UpgradeFromV1Ambiguous:
		// - v4.52.5 format (hostname/redirect is array): full transform via raw state
		// - v5 format (hostname/redirect is object): no-op re-decode with target schema
		1: {
			PriorSchema:   nil,
			StateUpgrader: v500.UpgradeFromV1Ambiguous,
		},
	}
}
