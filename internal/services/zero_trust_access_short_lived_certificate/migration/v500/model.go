package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy v4 SDKv2 Provider)
// ============================================================================

// SourceAccessCACertificateModel represents the legacy cloudflare_access_ca_certificate state from v4.x provider.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_access_ca_certificate
type SourceAccessCACertificateModel struct {
	ID            types.String `tfsdk:"id"`
	AccountID     types.String `tfsdk:"account_id"`
	ZoneID        types.String `tfsdk:"zone_id"`
	ApplicationID types.String `tfsdk:"application_id"` // Renamed to app_id in v5
	AUD           types.String `tfsdk:"aud"`
	PublicKey     types.String `tfsdk:"public_key"`
}

// ============================================================================
// Target Models (Current v5 Provider)
// ============================================================================

// TargetAccessShortLivedCertificateModel represents the current cloudflare_zero_trust_access_short_lived_certificate state.
// Schema version: 500
// Resource type: cloudflare_zero_trust_access_short_lived_certificate
type TargetAccessShortLivedCertificateModel struct {
	ID        types.String `tfsdk:"id"`
	AppID     types.String `tfsdk:"app_id"`
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`
	AUD       types.String `tfsdk:"aud"`
	PublicKey types.String `tfsdk:"public_key"`
}
