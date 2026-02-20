package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy v4 SDKv2 Provider)
// ============================================================================

// SourceAccessMutualTLSCertificateModel represents the legacy cloudflare_access_mutual_tls_certificate state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_mutual_tls_certificate
type SourceAccessMutualTLSCertificateModel struct {
	ID                  types.String `tfsdk:"id"`
	AccountID           types.String `tfsdk:"account_id"`
	ZoneID              types.String `tfsdk:"zone_id"`
	Certificate         types.String `tfsdk:"certificate"`
	Name                types.String `tfsdk:"name"`
	AssociatedHostnames types.Set    `tfsdk:"associated_hostnames"` // Set in v4, may be null if not configured
	ExpiresOn           types.String `tfsdk:"expires_on"`          // Plain string in v4
	Fingerprint         types.String `tfsdk:"fingerprint"`
}

// ============================================================================
// Target Models (Current v5 Provider)
// ============================================================================

// TargetAccessMTLSCertificateModel represents the current cloudflare_zero_trust_access_mtls_certificate state.
// Schema version: 500
// Resource type: cloudflare_zero_trust_access_mtls_certificate
type TargetAccessMTLSCertificateModel struct {
	ID                  types.String      `tfsdk:"id"`
	AccountID           types.String      `tfsdk:"account_id"`
	ZoneID              types.String      `tfsdk:"zone_id"`
	Certificate         types.String      `tfsdk:"certificate"`
	Name                types.String      `tfsdk:"name"`
	AssociatedHostnames *[]types.String   `tfsdk:"associated_hostnames"`
	ExpiresOn           timetypes.RFC3339 `tfsdk:"expires_on"`
	Fingerprint         types.String      `tfsdk:"fingerprint"`
}
