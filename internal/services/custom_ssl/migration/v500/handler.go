package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from the v4 SDKv2 provider (schema_version=1) to v5 (version=500).
//
// The v4 cloudflare_custom_ssl resource stored certificate fields in a nested
// custom_ssl_options TypeList MaxItems:1 block. This handler reads the v4 state
// and transforms it to the flat v5 structure.
//
// Key transformations:
//   - custom_ssl_options[0].certificate   → certificate (flat)
//   - custom_ssl_options[0].private_key   → private_key (flat)
//   - custom_ssl_options[0].bundle_method → bundle_method (flat)
//   - custom_ssl_options[0].type          → type (flat)
//   - custom_ssl_options[0].geo_restrictions (string) → geo_restrictions.label (nested object)
//   - custom_ssl_priority                 → dropped (write-only, not in v5)
//   - priority (Int64)                    → priority (Float64)
//   - uploaded_on/modified_on/expires_on (String) → timetypes.RFC3339
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading custom_ssl state from v4 SDKv2 provider (schema_version=1)")

	// Parse v4 state using the v4 source schema and model.
	var v4State SourceCloudflareCustomSSLModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5.
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write the transformed state.
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)

	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade that only bumps the schema version number.
// It is triggered when TF_MIG_TEST=1 causes GetSchemaVersion to return 500
// while existing state files have version=1.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading custom_ssl state from version=1 to version=500 (no-op)")

	// Copy raw state directly — no transformation needed.
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
