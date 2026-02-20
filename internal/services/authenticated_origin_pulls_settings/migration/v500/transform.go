package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts legacy v4 state to current v5 state.
// This function is used by both MoveState and UpgradeState.
//
// Transformation logic:
// - Copy: id, zone_id, enabled
// - Drop: hostname, authenticated_origin_pulls_certificate (removed in v5)
func Transform(ctx context.Context, source SourceCloudflareAuthenticatedOriginPullsModel) (TargetAuthenticatedOriginPullsSettingsModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	var target TargetAuthenticatedOriginPullsSettingsModel

	// Direct field mappings (no transformation needed)
	target.ID = source.ID
	target.ZoneID = source.ZoneID
	target.Enabled = source.Enabled

	// Fields removed in v5 (no action needed):
	// - hostname
	// - authenticated_origin_pulls_certificate

	return target, diags
}
