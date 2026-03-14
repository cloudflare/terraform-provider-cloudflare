package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from cloudflare_byo_ip_prefix schema_version=0 (v4 SDKv2)
// to schema_version=1 (v5 Plugin Framework).
//
// This is triggered when the v5 provider encounters state with schema_version=0, which occurs
// when users migrate from the v4 provider using tf-migrate.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_byo_ip_prefix state from schema_version=0 (v4) to schema_version=1 (v5)")

	// Parse the v4 state using the source schema
	var sourceState SourceCloudflareByoIPPrefixModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to v5 state
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "cloudflare_byo_ip_prefix state upgrade from schema_version=0 completed successfully")
}
