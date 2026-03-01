// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeStateFrom0To500 handles the state upgrade from v4 (version 0) to v5 (version 500)
// This is the main entry point called by the Terraform Plugin Framework
func UpgradeStateFrom0To500(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Starting state upgrade from version 0 (v4) to version 500 (v5) for zone_dnssec")

	// Parse the old state using the source model and attribute types
	var source SourceCloudflareZoneDNSSECModel
	diags := req.State.Get(ctx, &source)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Debug(ctx, "Successfully parsed v4 state", map[string]interface{}{
		"zone_id": source.ZoneID.ValueString(),
	})

	// Transform the state from v4 to v5
	target, transformDiags := Transform(ctx, source)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to transform state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	// Set the new state
	diags = resp.State.Set(ctx, target)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to set v5 state", map[string]interface{}{
			"diagnostics": resp.Diagnostics,
		})
		return
	}

	tflog.Info(ctx, "Successfully completed state upgrade to version 500 for zone_dnssec", map[string]interface{}{
		"zone_id": target.ZoneID.ValueString(),
	})
}
