package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareAuthenticatedOriginPullsModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (actual v4 schema version)
// Resource type: cloudflare_authenticated_origin_pulls
//
// This model includes ALL v4 fields, including ones that will be removed in v5:
// - hostname: Per-hostname AOP configuration (removed in v5)
// - authenticated_origin_pulls_certificate: Certificate ID for AOP (removed in v5)
type SourceCloudflareAuthenticatedOriginPullsModel struct {
	ID                                   types.String `tfsdk:"id"`
	ZoneID                              types.String `tfsdk:"zone_id"`
	Hostname                            types.String `tfsdk:"hostname"`
	AuthenticatedOriginPullsCertificate types.String `tfsdk:"authenticated_origin_pulls_certificate"`
	Enabled                             types.Bool   `tfsdk:"enabled"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAuthenticatedOriginPullsSettingsModel represents the current resource state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_authenticated_origin_pulls_settings
//
// Note: This matches the model in the parent package's model.go file.
// The v5 model is simplified - it only supports zone-wide AOP settings.
type TargetAuthenticatedOriginPullsSettingsModel struct {
	ID      types.String `tfsdk:"id"`
	ZoneID  types.String `tfsdk:"zone_id"`
	Enabled types.Bool   `tfsdk:"enabled"`
}
