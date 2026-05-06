// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_group/migration/v500"
)

var _ resource.ResourceWithMoveState = (*ZeroTrustAccessGroupResource)(nil)
var _ resource.ResourceWithUpgradeState = (*ZeroTrustAccessGroupResource)(nil)

// MoveState handles moves from cloudflare_access_group (v0) to cloudflare_zero_trust_access_group (v500).
// This is triggered when users use the `moved` block (Terraform 1.8+):
//
//	moved {
//	    from = cloudflare_access_group.example
//	    to   = cloudflare_zero_trust_access_group.example
//	}
func (r *ZeroTrustAccessGroupResource) MoveState(ctx context.Context) []resource.StateMover {
	sourceSchema := v500.SourceV4ZeroTrustAccessGroupSchema()
	return []resource.StateMover{
		{
			SourceSchema: &sourceSchema,
			StateMover:   v500.MoveState,
		},
	}
}

// UpgradeState handles schema version upgrades for cloudflare_zero_trust_access_group.
//
// This handles two upgrade paths:
// 1. v4 state (schema_version=0) → v5 (version=500): Full transformation
// 2. v5 state (version=1) → v5 (version=500): No-op upgrade
func (r *ZeroTrustAccessGroupResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	// v5 schema for version=1 upgrader (override version to match production state)
	v5SchemaVersion1 := ResourceSchema(ctx)
	v5SchemaVersion1.Version = 1

	return map[int64]resource.StateUpgrader{
		// Handle BOTH version=0 formats:
		// - v4 SDKv2 state (boolean selectors like any_valid_service_token) -> full transform
		// - early v5 published state (object selectors) -> no-op version bump
		//
		// PriorSchema is nil to avoid the framework trying to decode v4 state
		// using the v5 schema, which fails for boolean-to-object type changes
		// (e.g., any_valid_service_token: bool in v4 → SingleNestedAttribute in v5).
		0: {
			PriorSchema:   nil,
			StateUpgrader: upgradeFromV0,
		},

		// Handle state from v5 Plugin Framework provider (version=1)
		// No-op: schema is compatible, just bumps version to 500
		1: {
			PriorSchema: &v5SchemaVersion1,
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				resp.State.Raw = req.State.Raw
			},
		},
	}
}

// upgradeFromV0 dispatches version=0 state to the correct handler based on format.
// Version=0 state can be either:
// - v4 SDKv2 format: boolean selectors (any_valid_service_token = false), list fields
// - early v5 format: object selectors (any_valid_service_token = {}), published before version bump
func upgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil for schema version 0 migration")
		return
	}

	// Try decoding as early v5 (version=0) first.
	// If this succeeds, state is already v5-shaped and can be passed through.
	v5SchemaVersion0 := ResourceSchema(ctx)
	v5SchemaVersion0.Version = 0
	v5Type := v5SchemaVersion0.Type().TerraformType(ctx)
	v5RawValue, err := req.RawState.Unmarshal(v5Type)
	if err == nil {
		resp.State.Raw = v5RawValue
		return
	}

	// Not decodable as v5 version=0; treat as v4 SDKv2 and transform.
	v500.UpgradeFromV0(ctx, req, resp)
}
