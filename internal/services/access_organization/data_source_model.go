// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_organization

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessOrganizationResultDataSourceEnvelope struct {
	Result AccessOrganizationDataSourceModel `json:"result,computed"`
}

type AccessOrganizationDataSourceModel struct {
	AccountID                      types.String                                  `tfsdk:"account_id" path:"account_id"`
	ZoneID                         types.String                                  `tfsdk:"zone_id" path:"zone_id"`
	AllowAuthenticateViaWARP       types.Bool                                    `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp"`
	AuthDomain                     types.String                                  `tfsdk:"auth_domain" json:"auth_domain"`
	AutoRedirectToIdentity         types.Bool                                    `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity"`
	CreatedAt                      timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at"`
	CustomPages                    *AccessOrganizationCustomPagesDataSourceModel `tfsdk:"custom_pages" json:"custom_pages"`
	IsUIReadOnly                   types.Bool                                    `tfsdk:"is_ui_read_only" json:"is_ui_read_only"`
	LoginDesign                    *AccessOrganizationLoginDesignDataSourceModel `tfsdk:"login_design" json:"login_design"`
	Name                           types.String                                  `tfsdk:"name" json:"name"`
	SessionDuration                types.String                                  `tfsdk:"session_duration" json:"session_duration"`
	UIReadOnlyToggleReason         types.String                                  `tfsdk:"ui_read_only_toggle_reason" json:"ui_read_only_toggle_reason"`
	UpdatedAt                      timetypes.RFC3339                             `tfsdk:"updated_at" json:"updated_at"`
	UserSeatExpirationInactiveTime types.String                                  `tfsdk:"user_seat_expiration_inactive_time" json:"user_seat_expiration_inactive_time"`
	WARPAuthSessionDuration        types.String                                  `tfsdk:"warp_auth_session_duration" json:"warp_auth_session_duration"`
}

type AccessOrganizationCustomPagesDataSourceModel struct {
	Forbidden      types.String `tfsdk:"forbidden" json:"forbidden"`
	IdentityDenied types.String `tfsdk:"identity_denied" json:"identity_denied"`
}

type AccessOrganizationLoginDesignDataSourceModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path"`
	TextColor       types.String `tfsdk:"text_color" json:"text_color"`
}
