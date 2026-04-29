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
// We detect the format at runtime:
// - If load_shedding is an array (or missing), it's v4 format → transform
// - If load_shedding is an object, it's v5 format → no-op (copy state through)
//
// Key transformations (v4 only):
// - load_shedding: Array[0] → NestedObject (SDK v2 TypeList MaxItems:1)
// - origin_steering: Array[0] → NestedObject (SDK v2 TypeList MaxItems:1)
// - origins.header: Complex nested structure transformation
// - check_regions: Set → List
// - origins: Set → List
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading load_balancer_pool state from schema_version=0")

	// Try to parse with v4 schema
	// If it succeeds, it's v4 format and needs transformation
	// If it fails (e.g., "expected '[', got '{'"), it's v5 format and needs no-op
	var sourceState SourceCloudflareLoadBalancerPoolModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)

	if resp.Diagnostics.HasError() {
		// Parsing failed - likely v5 format (Plugin Framework) where load_shedding/origin_steering are objects
		tflog.Info(ctx, "Failed to parse as v4 format, assuming v5 format - performing no-op upgrade")
		tflog.Debug(ctx, "Parse error details", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})

		// Clear diagnostics and do no-op
		resp.Diagnostics = nil
		resp.State.Raw = req.State.Raw
		return
	}

	// Successfully parsed as v4 - perform transformation
	tflog.Info(ctx, "Successfully parsed as v4 format (SDKv2) - performing transformation")

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
