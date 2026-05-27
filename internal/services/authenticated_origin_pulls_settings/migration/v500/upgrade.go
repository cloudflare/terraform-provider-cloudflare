package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromLegacyV0 handles state upgrade from legacy v4 (schema version 0) to v5 (schema version 500).
//
// This upgrader is triggered when:
// - User has v4 state (schema_version: 0)
// - User runs `terraform state mv` to rename from cloudflare_authenticated_origin_pulls to cloudflare_authenticated_origin_pulls_settings
// - Terraform detects schema version mismatch and calls this upgrader
//
// Note: For Terraform 1.8+, MoveState will be triggered instead via `moved` blocks.
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading authenticated_origin_pulls_settings state from legacy v0 to v500")

	// Parse the source state (legacy v4 format with schema version 0)
	var sourceState SourceCloudflareAuthenticatedOriginPullsModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target (current v5 format)
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from v0 to v500 completed successfully")
}

// UpgradeFromV1 handles state upgrade from v5 (schema version 1) to v5 (schema version 500).
//
// This is a no-op upgrader that exists solely to support the schema version rollout mechanism.
// The v1 and v500 schemas are identical, so no transformation is needed — just a version bump.
//
// IMPORTANT: This upgrader is registered with PriorSchema=nil in migrations.go, so
// req.State is NOT populated by the framework. We must use req.RawState directly.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Debug(ctx, "No-op state upgrade from v1 to v500 (schema versions are compatible)")

	// req.State is nil because PriorSchema is nil — use RawState instead.
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, err := req.RawState.Unmarshal(targetType)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal state during v1→v500 upgrade for authenticated_origin_pulls_settings",
			"The raw state could not be read with the current schema. This may indicate state corruption. Error: "+err.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
}
