package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// This performs the transformation from v4 → v5 format:
// - cache_type="smart" → value="on"
// - cache_type="off" → value="off"
// - cache_type="generic" → value="off" (argo_tiered_caching with value="on" is created from config)
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading tiered_cache state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 model
	var v4State SourceTieredCacheModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, transformDiags := Transform(ctx, v4State)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading tiered_cache state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// Transform converts source (legacy v4) state to target (current v5) state.
// This function handles the cache_type to value transformation.
func Transform(ctx context.Context, source SourceTieredCacheModel) (*TargetTieredCacheModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Transform cache_type to value
	var newValue string
	cacheType := source.CacheType.ValueString()
	switch cacheType {
	case "smart":
		newValue = "on"
	case "off":
		newValue = "off"
	case "generic":
		// generic means argo tiered caching is ON and tiered caching is OFF
		// tf-migrate creates both cloudflare_tiered_cache (value=off) and cloudflare_argo_tiered_caching (value=on)
		// The argo_tiered_caching resource will be created on apply from the config (not from state upgrade)
		newValue = "off"
	default:
		// For unknown values (e.g., variable references), preserve as-is
		// The value will be validated by the schema validator
		newValue = cacheType
	}

	// Create the target state
	target := &TargetTieredCacheModel{
		ID:         source.ID,
		ZoneID:     source.ZoneID,
		Value:      types.StringValue(newValue),
		Editable:   source.Editable,
		ModifiedOn: source.ModifiedOn,
	}

	return target, diags
}
