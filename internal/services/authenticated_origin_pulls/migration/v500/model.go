package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareAuthenticatedOriginPullsModel represents the legacy resource state from v4.x provider.
// Schema version: 0 (default, not specified in v4)
// Resource type: cloudflare_authenticated_origin_pulls
//
// The v4 resource handled three modes in one resource:
// 1. Global AOP: zone_id + enabled
// 2. Per-Zone AOP: zone_id + authenticated_origin_pulls_certificate + enabled
// 3. Per-Hostname AOP: zone_id + hostname + authenticated_origin_pulls_certificate + enabled
//
// In v5, this resource only handles Per-Hostname AOP mode (mode 3).
// Modes 1 and 2 migrate to cloudflare_authenticated_origin_pulls_settings instead.
type SourceCloudflareAuthenticatedOriginPullsModel struct {
	ID                                  types.String `tfsdk:"id"`
	ZoneID                              types.String `tfsdk:"zone_id"`
	Hostname                            types.String `tfsdk:"hostname"`
	AuthenticatedOriginPullsCertificate types.String `tfsdk:"authenticated_origin_pulls_certificate"`
	Enabled                             types.Bool   `tfsdk:"enabled"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAuthenticatedOriginPullsModel represents the current resource state from v5.x+ provider.
// Schema version: 500 (after migration)
// Resource type: cloudflare_authenticated_origin_pulls
//
// Note: This should match the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetAuthenticatedOriginPullsModel struct {
	ID             types.String                                  `tfsdk:"id"`
	Hostname       types.String                                  `tfsdk:"hostname"`
	ZoneID         types.String                                  `tfsdk:"zone_id"`
	Config         *[]*TargetAuthenticatedOriginPullsConfigModel `tfsdk:"config"`
	CERTID         types.String                                  `tfsdk:"cert_id"`
	CERTStatus     types.String                                  `tfsdk:"cert_status"`
	CERTUpdatedAt  timetypes.RFC3339                             `tfsdk:"cert_updated_at"`
	CERTUploadedOn timetypes.RFC3339                             `tfsdk:"cert_uploaded_on"`
	Certificate    types.String                                  `tfsdk:"certificate"`
	CreatedAt      timetypes.RFC3339                             `tfsdk:"created_at"`
	Enabled        types.Bool                                    `tfsdk:"enabled"`
	ExpiresOn      timetypes.RFC3339                             `tfsdk:"expires_on"`
	Issuer         types.String                                  `tfsdk:"issuer"`
	PrivateKey     types.String                                  `tfsdk:"private_key"`
	SerialNumber   types.String                                  `tfsdk:"serial_number"`
	Signature      types.String                                  `tfsdk:"signature"`
	Status         types.String                                  `tfsdk:"status"`
	UpdatedAt      timetypes.RFC3339                             `tfsdk:"updated_at"`
}

// TargetAuthenticatedOriginPullsConfigModel represents the nested config structure in v5.
// The v5 API expects an array of config objects, but the resource enforces exactly one item.
type TargetAuthenticatedOriginPullsConfigModel struct {
	CERTID   types.String `tfsdk:"cert_id"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Hostname types.String `tfsdk:"hostname"`
}
