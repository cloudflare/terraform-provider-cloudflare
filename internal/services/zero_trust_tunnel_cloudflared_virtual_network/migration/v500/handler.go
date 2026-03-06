package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// UpgradeFromV4 upgrades state from schema version 0 (v4 provider) to version 500.
// It ensures comment and is_default_network have their default values if absent.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var source SourceVirtualNetworkModel
	resp.Diagnostics.Append(req.State.Get(ctx, &source)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target, diags := Transform(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)
}

// UpgradeFromV5 is a no-op upgrader for state already at schema version 1.
// It is triggered when TF_MIG_TEST=1 advances the version from 1 to 500.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	resp.State.Raw = req.State.Raw
}
