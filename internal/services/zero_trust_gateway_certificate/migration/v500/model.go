package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareZeroTrustGatewayCertificateModel represents the legacy resource
// state from the v4.x SDKv2 provider (resource: cloudflare_zero_trust_gateway_certificate,
// implemented as resourceCloudflareTeamsCertificate in the SDKv2 provider).
// Schema version: 0 (SDKv2 default, no explicit SchemaVersion was set).
type SourceCloudflareZeroTrustGatewayCertificateModel struct {
	ID                 types.String `tfsdk:"id"`
	AccountID          types.String `tfsdk:"account_id"`
	Custom             types.Bool   `tfsdk:"custom"`
	GatewayManaged     types.Bool   `tfsdk:"gateway_managed"`
	ValidityPeriodDays types.Int64  `tfsdk:"validity_period_days"`
	Activate           types.Bool   `tfsdk:"activate"`
	InUse              types.Bool   `tfsdk:"in_use"`
	BindingStatus      types.String `tfsdk:"binding_status"`
	QsPackID           types.String `tfsdk:"qs_pack_id"`
	UploadedOn         types.String `tfsdk:"uploaded_on"`
	CreatedAt          types.String `tfsdk:"created_at"`
	ExpiresOn          types.String `tfsdk:"expires_on"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetZeroTrustGatewayCertificateModel represents the current resource state
// from the v5.x+ provider (resource: cloudflare_zero_trust_gateway_certificate).
// Schema version: 500.
//
// Note: This mirrors ZeroTrustGatewayCertificateModel in the parent package.
// It is duplicated here to keep the migration package self-contained.
type TargetZeroTrustGatewayCertificateModel struct {
	ID                 types.String      `tfsdk:"id"`
	AccountID          types.String      `tfsdk:"account_id"`
	ValidityPeriodDays types.Int64       `tfsdk:"validity_period_days"`
	Activate           types.Bool        `tfsdk:"activate"`
	BindingStatus      types.String      `tfsdk:"binding_status"`
	Certificate        types.String      `tfsdk:"certificate"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at"`
	ExpiresOn          timetypes.RFC3339 `tfsdk:"expires_on"`
	Fingerprint        types.String      `tfsdk:"fingerprint"`
	InUse              types.Bool        `tfsdk:"in_use"`
	IssuerOrg          types.String      `tfsdk:"issuer_org"`
	IssuerRaw          types.String      `tfsdk:"issuer_raw"`
	Type               types.String      `tfsdk:"type"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at"`
	UploadedOn         timetypes.RFC3339 `tfsdk:"uploaded_on"`
}
