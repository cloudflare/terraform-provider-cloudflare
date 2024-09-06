// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_organization

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustOrganizationResultDataSourceEnvelope struct {
	Result ZeroTrustOrganizationDataSourceModel `json:"result,computed"`
}

type ZeroTrustOrganizationDataSourceModel struct {
	AccountID                      types.String                                     `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                         types.String                                     `tfsdk:"zone_id" path:"zone_id,optional"`
	AllowAuthenticateViaWARP       types.Bool                                       `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,optional"`
	AuthDomain                     types.String                                     `tfsdk:"auth_domain" json:"auth_domain,optional"`
	CreatedAt                      timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,optional" format:"date-time"`
	IsUIReadOnly                   types.Bool                                       `tfsdk:"is_ui_read_only" json:"is_ui_read_only,optional"`
	Name                           types.String                                     `tfsdk:"name" json:"name,optional"`
	SessionDuration                types.String                                     `tfsdk:"session_duration" json:"session_duration,optional"`
	UIReadOnlyToggleReason         types.String                                     `tfsdk:"ui_read_only_toggle_reason" json:"ui_read_only_toggle_reason,optional"`
	UpdatedAt                      timetypes.RFC3339                                `tfsdk:"updated_at" json:"updated_at,optional" format:"date-time"`
	UserSeatExpirationInactiveTime types.String                                     `tfsdk:"user_seat_expiration_inactive_time" json:"user_seat_expiration_inactive_time,optional"`
	WARPAuthSessionDuration        types.String                                     `tfsdk:"warp_auth_session_duration" json:"warp_auth_session_duration,optional"`
	CustomPages                    *ZeroTrustOrganizationCustomPagesDataSourceModel `tfsdk:"custom_pages" json:"custom_pages,optional"`
	LoginDesign                    *ZeroTrustOrganizationLoginDesignDataSourceModel `tfsdk:"login_design" json:"login_design,optional"`
	AutoRedirectToIdentity         types.Bool                                       `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed_optional"`
}

func (m *ZeroTrustOrganizationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.OrganizationListParams, diags diag.Diagnostics) {
	params = zero_trust.OrganizationListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustOrganizationCustomPagesDataSourceModel struct {
	Forbidden      types.String `tfsdk:"forbidden" json:"forbidden,computed"`
	IdentityDenied types.String `tfsdk:"identity_denied" json:"identity_denied,computed"`
}

type ZeroTrustOrganizationLoginDesignDataSourceModel struct {
	BackgroundColor types.String `tfsdk:"background_color" json:"background_color,computed"`
	FooterText      types.String `tfsdk:"footer_text" json:"footer_text,computed"`
	HeaderText      types.String `tfsdk:"header_text" json:"header_text,computed"`
	LogoPath        types.String `tfsdk:"logo_path" json:"logo_path,computed"`
	TextColor       types.String `tfsdk:"text_color" json:"text_color,computed"`
}
