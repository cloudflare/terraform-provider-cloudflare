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
// This handles v4 state where config is a list block (array in JSON).
// We parse using the source model, transform to target model, and set the upgraded state.
func UpgradeFromLegacyV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero trust device managed networks state from schema_version=0")

	// Parse using v4-shaped source model (config as list block / array)
	var sourceState SourceCloudflareDeviceManagedNetworksModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5 target model (config as SingleNestedAttribute / object)
	targetState := TargetZeroTrustDeviceManagedNetworksModel{
		ID:        sourceState.ID,
		AccountID: sourceState.AccountID,
		Name:      sourceState.Name,
		Type:      sourceState.Type,
	}

	// Backfill computed network_id from id
	if !sourceState.ID.IsNull() && !sourceState.ID.IsUnknown() {
		targetState.NetworkID = types.StringValue(sourceState.ID.ValueString())
	}

	// Transform config from list (array) to object
	if len(sourceState.Config) > 0 {
		targetState.Config = &TargetConfigModel{
			TLSSockaddr: sourceState.Config[0].TLSSockaddr,
			Sha256:      sourceState.Config[0].Sha256,
		}
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, &targetState)...)

	tflog.Info(ctx, "State upgrade from schema_version=0 completed successfully")
}
