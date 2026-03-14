package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV1 handles state upgrades from earlier v5 versions (schema_version=1) to current v500.
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeFromV1(
	ctx context.Context,
	req resource.UpgradeStateRequest,
	resp *resource.UpgradeStateResponse,
) {
	tflog.Info(ctx, "Upgrading zero trust device managed networks state from schema_version=1")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}

// UpgradeFromLegacyV0 handles state upgrades for schema_version=0.
//
// For this resource we normalize published v5-shaped state that may have been
// persisted at schema_version=0 by older provider releases.
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero trust device managed networks state from schema_version=0")

	// Parse using v5-shaped model (config object), with schema version pinned to 0 by caller.
	var state TargetZeroTrustDeviceManagedNetworksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Backfill computed network_id from id when absent.
	if (state.NetworkID.IsNull() || state.NetworkID.IsUnknown()) && !state.ID.IsNull() && !state.ID.IsUnknown() {
		state.NetworkID = types.StringValue(state.ID.ValueString())
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)

	tflog.Info(ctx, "State upgrade from schema_version=0 completed successfully")
}
