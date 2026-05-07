// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to schema_version=500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
//
// Why this exists: Terraform requires explicit upgraders to be defined for version tracking,
// even when the schema is identical. This ensures the schema_version is updated in the statefile.
func UpgradeFromV0(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading load_balancer_pool state from schema_version=0 (no-op for v5 same-version states)")
	// No-op upgrade: schema is compatible, just copy raw state through
	// We use the raw state value directly to avoid issues with custom field type serialization
	resp.State.Raw = req.State.Raw
}

// UpgradeFromLegacyV0 handles state upgrades from schema_version=0.
//
// IMPORTANT: schema_version=0 is used by BOTH:
// 1. v4 (SDKv2) provider - needs transformation (load_shedding/origin_steering as arrays)
// 2. Early v5 (5.0.0-5.15.x) releases - already in correct format (as objects)
//
// Detection strategy:
// We try to unmarshal the raw state using the TARGET (current v5) schema type first.
// - If it succeeds → state is already v5 format → no-op (copy through)
// - If it fails (e.g., load_shedding is array not object) → v4 format → transform
//
// This approach avoids relying on req.State (decoded with v4 PriorSchema) for v5 state,
// which would strip fields that only exist in v5 (e.g. disabled_at, networks).
//
// Key transformations (v4 only):
// - load_shedding: Array[0] → NestedObject (SDK v2 TypeList MaxItems:1)
// - origin_steering: Array[0] → NestedObject (SDK v2 TypeList MaxItems:1)
// - origins.header: Complex nested structure transformation
// - check_regions: Set → List
// - origins: Set → List
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_pool state from schema_version=0")

	// First, try to unmarshal raw state with the TARGET schema type.
	// This succeeds for v5 state (load_shedding/origin_steering are objects)
	// and fails for v4 state (they are arrays).
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	val, err := req.RawState.Unmarshal(targetType)
	if err == nil {
		// v5 format detected - state is already compatible with current schema.
		// Fields absent from the JSON (e.g. newly added fields) become null.
		tflog.Info(ctx, "State is v5 format - performing no-op upgrade via RawState")
		resp.State.Raw = val
		return
	}

	// Target schema unmarshal failed - this is v4 format (arrays for load_shedding etc.)
	tflog.Info(ctx, "State is v4 format (SDKv2) - performing transformation",
		map[string]interface{}{"unmarshal_err": err.Error()})

	// Parse with v4 PriorSchema (req.State was decoded using v4 source schema)
	var sourceState SourceCloudflareLoadBalancerPoolModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Debug(ctx, "Parsed v4 source state successfully", map[string]interface{}{
		"id":         sourceState.ID.ValueString(),
		"account_id": sourceState.AccountID.ValueString(),
		"name":       sourceState.Name.ValueString(),
	})

	// Transform to target
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set upgraded state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Info(ctx, "State upgrade from v4 load_balancer_pool completed successfully")
}
