// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_organization

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessOrganizationResultEnvelope struct {
	Result AccessOrganizationModel `json:"result,computed"`
}

type AccessOrganizationModel struct {
	AccountID                      types.String                        `tfsdk:"account_id" path:"account_id"`
	ZoneID                         types.String                        `tfsdk:"zone_id" path:"zone_id"`
	AuthDomain                     types.String                        `tfsdk:"auth_domain" json:"auth_domain"`
	Name                           types.String                        `tfsdk:"name" json:"name"`
	AllowAuthenticateViaWARP       types.Bool                          `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp"`
	AutoRedirectToIdentity         types.Bool                          `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity"`
	IsUiReadOnly                   types.Bool                          `tfsdk:"is_ui_read_only" json:"is_ui_read_only"`
	LoginDesign                    *AccessOrganizationLoginDesignModel `tfsdk:"login_design" json:"login_design"`
	SessionDuration                types.String                        `tfsdk:"session_duration" json:"session_duration"`
	UiReadOnlyToggleReason         types.String                        `tfsdk:"ui_read_only_toggle_reason" json:"ui_read_only_toggle_reason"`
	UserSeatExpirationInactiveTime types.String                        `tfsdk:"user_seat_expiration_inactive_time" json:"user_seat_expiration_inactive_time"`
	WARPAuthSessionDuration        types.String                        `tfsdk:"warp_auth_session_duration" json:"warp_auth_session_duration"`
}

type AccessOrganizationLoginDesignModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path"`
	TextColor       types.String `tfsdk:"text_color" json:"text_color"`
}
