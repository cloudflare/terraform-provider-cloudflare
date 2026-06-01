package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SourceCloudflareAccessOrganizationModel represents the source state structure from v4.
// This corresponds to schema_version=0 from the legacy (SDKv2) cloudflare provider.
//
// Handles BOTH v4 resource names (they share identical schemas):
// - cloudflare_access_organization
// - cloudflare_zero_trust_access_organization
//
// Used by both MoveState (Terraform 1.8+) and UpgradeFromLegacyV0 (Terraform < 1.8)
// to parse legacy state.
type SourceCloudflareAccessOrganizationModel struct {
	ID                             types.String             `tfsdk:"id"`
	AccountID                      types.String             `tfsdk:"account_id"`
	ZoneID                         types.String             `tfsdk:"zone_id"`
	AuthDomain                     types.String             `tfsdk:"auth_domain"`
	Name                           types.String             `tfsdk:"name"`
	IsUIReadOnly                   types.Bool               `tfsdk:"is_ui_read_only"`
	UIReadOnlyToggleReason         types.String             `tfsdk:"ui_read_only_toggle_reason"`
	UserSeatExpirationInactiveTime types.String             `tfsdk:"user_seat_expiration_inactive_time"`
	AutoRedirectToIdentity         types.Bool               `tfsdk:"auto_redirect_to_identity"`
	SessionDuration                types.String             `tfsdk:"session_duration"`
	AllowAuthenticateViaWARP       types.Bool               `tfsdk:"allow_authenticate_via_warp"`
	WARPAuthSessionDuration        types.String             `tfsdk:"warp_auth_session_duration"`
	LoginDesign                    []SourceLoginDesignModel `tfsdk:"login_design"`
	CustomPages                    []SourceCustomPagesModel `tfsdk:"custom_pages"`
}

// SourceLoginDesignModel represents the login_design nested structure in v4.
// In SDK v2, this was stored as a list with MaxItems: 1, so state contains an array.
type SourceLoginDesignModel struct {
	BackgroundColor types.String `tfsdk:"background_color"`
	TextColor       types.String `tfsdk:"text_color"`
	LogoPath        types.String `tfsdk:"logo_path"`
	HeaderText      types.String `tfsdk:"header_text"`
	FooterText      types.String `tfsdk:"footer_text"`
}

// SourceCustomPagesModel represents the custom_pages nested structure in v4.
// In SDK v2, this was stored as a list with MaxItems: 1, so state contains an array.
type SourceCustomPagesModel struct {
	IdentityDenied types.String `tfsdk:"identity_denied"`
	Forbidden      types.String `tfsdk:"forbidden"`
}

// TargetZeroTrustOrganizationModel represents the target state structure for v5.
// This corresponds to schema_version=500 in the current provider.
// Must match zero_trust_organization.ZeroTrustOrganizationModel structure for core fields.
type TargetZeroTrustOrganizationModel struct {
	AccountID                              types.String            `tfsdk:"account_id"`
	ZoneID                                 types.String            `tfsdk:"zone_id"`
	AuthDomain                             types.String            `tfsdk:"auth_domain"`
	DenyUnmatchedRequests                  types.Bool              `tfsdk:"deny_unmatched_requests"`
	Name                                   types.String            `tfsdk:"name"`
	SessionDuration                        types.String            `tfsdk:"session_duration"`
	UserSeatExpirationInactiveTime         types.String            `tfsdk:"user_seat_expiration_inactive_time"`
	WARPAuthSessionDuration                types.String            `tfsdk:"warp_auth_session_duration"`
	DenyUnmatchedRequestsExemptedZoneNames *[]types.String         `tfsdk:"deny_unmatched_requests_exempted_zone_names"`
	CustomPages                            *TargetCustomPagesModel `tfsdk:"custom_pages"`
	LoginDesign                            *TargetLoginDesignModel `tfsdk:"login_design"`
	MfaConfig                              *TargetMfaConfigModel   `tfsdk:"mfa_config"`
	AllowAuthenticateViaWARP               types.Bool              `tfsdk:"allow_authenticate_via_warp"`
	AutoRedirectToIdentity                 types.Bool              `tfsdk:"auto_redirect_to_identity"`
	IsUIReadOnly                           types.Bool              `tfsdk:"is_ui_read_only"`
	MfaConfigurationAllowed                types.Bool              `tfsdk:"mfa_configuration_allowed"`
	MfaRequiredForAllApps                  types.Bool              `tfsdk:"mfa_required_for_all_apps"`
	UIReadOnlyToggleReason                 types.String            `tfsdk:"ui_read_only_toggle_reason"`
}

// TargetCustomPagesModel represents the custom_pages nested structure in v5.
// In Framework, this is a SingleNestedAttribute (object), not a list.
type TargetCustomPagesModel struct {
	Forbidden      types.String `tfsdk:"forbidden"`
	IdentityDenied types.String `tfsdk:"identity_denied"`
}

// TargetLoginDesignModel represents the login_design nested structure in v5.
// In Framework, this is a SingleNestedAttribute (object), not a list.
type TargetLoginDesignModel struct {
	BackgroundColor types.String `tfsdk:"background_color"`
	FooterText      types.String `tfsdk:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text"`
	LogoPath        types.String `tfsdk:"logo_path"`
	TextColor       types.String `tfsdk:"text_color"`
}

// TargetMfaConfigModel represents the mfa_config nested structure in v5.
// This is a new field in v5, not present in v4.
type TargetMfaConfigModel struct {
	AllowedAuthenticators *[]types.String `tfsdk:"allowed_authenticators"`
	SessionDuration       types.String    `tfsdk:"session_duration"`
}
