// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustOrganizationResultEnvelope struct {
	Result ZeroTrustOrganizationModel `json:"result"`
}

type ZeroTrustOrganizationModel struct {
	ID                             types.String                                                    `tfsdk:"id" json:"-,computed"`
	Name                           types.String                                                    `tfsdk:"name" json:"name,required"`
	AccountID                      types.String                                                    `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                         types.String                                                    `tfsdk:"zone_id" path:"zone_id,optional"`
	AuthDomain                     types.String                                                    `tfsdk:"auth_domain" json:"auth_domain,required"`
	AllowAuthenticateViaWARP       types.Bool                                                      `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,computed_optional"`
	AutoRedirectToIdentity         types.Bool                                                      `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed_optional"`
	IsUIReadOnly                   types.Bool                                                      `tfsdk:"is_ui_read_only" json:"is_ui_read_only,computed_optional"`
	SessionDuration                types.String                                                    `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	UIReadOnlyToggleReason         types.String                                                    `tfsdk:"ui_read_only_toggle_reason" json:"ui_read_only_toggle_reason,computed_optional"`
	UserSeatExpirationInactiveTime types.String                                                    `tfsdk:"user_seat_expiration_inactive_time" json:"user_seat_expiration_inactive_time,computed_optional"`
	WARPAuthSessionDuration        types.String                                                    `tfsdk:"warp_auth_session_duration" json:"warp_auth_session_duration,computed_optional"`
	CustomPages                    customfield.NestedObject[ZeroTrustOrganizationCustomPagesModel] `tfsdk:"custom_pages" json:"custom_pages,computed_optional"`
	LoginDesign                    customfield.NestedObject[ZeroTrustOrganizationLoginDesignModel] `tfsdk:"login_design" json:"login_design,computed_optional"`
	CreatedAt                      timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt                      timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustOrganizationCustomPagesModel struct {
	Forbidden      types.String `tfsdk:"forbidden" json:"forbidden,computed_optional"`
	IdentityDenied types.String `tfsdk:"identity_denied" json:"identity_denied,computed_optional"`
}

type ZeroTrustOrganizationLoginDesignModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color,computed_optional"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text,computed_optional"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text,computed_optional"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path,computed_optional"`
	TextColor       types.String `tfsdk:"text_color" json:"text_color,computed_optional"`
}
