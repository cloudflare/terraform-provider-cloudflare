package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles v4 state at schema_version=0.
// Merges items + items_with_description into unified items set.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_list state from v4 (schema_version=0)")

	var v4State SourceTeamsListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "V4→V5 zero_trust_list state upgrade completed")
}

// UpgradeFromV5V1 handles v5 state at version=1 — no-op.
func UpgradeFromV5V1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_list state from version=1 to current (no-op)")
	resp.State.Raw = req.State.Raw
}
