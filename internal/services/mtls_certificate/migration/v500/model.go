package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareMTLSCertificateModel represents the source cloudflare_mtls_certificate state structure.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
// Used by UpgradeFromV4 to parse legacy state.
//
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_mtls_certificate.go
type SourceCloudflareMTLSCertificateModel struct {
	ID           types.String `tfsdk:"id"`
	AccountID    types.String `tfsdk:"account_id"`
	CA           types.Bool   `tfsdk:"ca"`
	Certificates types.String `tfsdk:"certificates"`
	Name         types.String `tfsdk:"name"`
	PrivateKey   types.String `tfsdk:"private_key"`
	// Computed fields from v4 - stored as plain strings (TypeString)
	Issuer       types.String `tfsdk:"issuer"`
	Signature    types.String `tfsdk:"signature"`
	SerialNumber types.String `tfsdk:"serial_number"`
	UploadedOn   types.String `tfsdk:"uploaded_on"`
	ExpiresOn    types.String `tfsdk:"expires_on"`
	// Note: updated_at does NOT exist in v4 - it is a new v5 computed field
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetMTLSCertificateModel represents the target cloudflare_mtls_certificate state structure (v500).
// This matches the structure in the parent package's model.go file (MTLSCertificateModel).
type TargetMTLSCertificateModel struct {
	ID           types.String `tfsdk:"id"`
	AccountID    types.String `tfsdk:"account_id"`
	CA           types.Bool   `tfsdk:"ca"`
	Certificates types.String `tfsdk:"certificates"`
	Name         types.String `tfsdk:"name"`
	PrivateKey   types.String `tfsdk:"private_key"`
	// Computed fields - upgraded to timetypes.RFC3339 in v5
	ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on"`
	Issuer       types.String      `tfsdk:"issuer"`
	SerialNumber types.String      `tfsdk:"serial_number"`
	Signature    types.String      `tfsdk:"signature"`
	// New computed field in v5 - not present in v4, set to null during migration
	UpdatedAt  timetypes.RFC3339 `tfsdk:"updated_at"`
	UploadedOn timetypes.RFC3339 `tfsdk:"uploaded_on"`
}
