// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_application

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessApplicationResultEnvelope struct {
	Result AccessApplicationModel `json:"result"`
}

type AccessApplicationModel struct {
	ID                       types.String `tfsdk:"id" json:"id,computed"`
	AccountID                types.String `tfsdk:"account_id" path:"account_id,required"`
	ZoneID                   types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Name                     types.String `tfsdk:"name" json:"name,required"`
	Domain                   types.String `tfsdk:"domain" json:"domain,computed"`
	Type                     types.String `tfsdk:"type" json:"type,computed"`
	SessionDuration          types.String `tfsdk:"session_duration" json:"session_duration,computed"`
	AutoRedirectToIdentity   types.Bool   `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed"`
	EnableBindingCookie      types.Bool   `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed"`
	AllowedIDPs              types.List   `tfsdk:"allowed_idps" json:"allowed_idps,computed"`
	CustomDenyMessage        types.String `tfsdk:"custom_deny_message" json:"custom_deny_message,computed"`
	CustomDenyURL            types.String `tfsdk:"custom_deny_url" json:"custom_deny_url,computed"`
	CustomNonIdentityDenyURL types.String `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url,computed"`
	LogoURL                  types.String `tfsdk:"logo_url" json:"logo_url,computed"`
	CreatedAt                types.String `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt                types.String `tfsdk:"updated_at" json:"updated_at,computed"`
}
