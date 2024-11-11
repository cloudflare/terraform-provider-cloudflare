// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProviderResultEnvelope struct {
	Result ZeroTrustAccessIdentityProviderModel `json:"result"`
}

type ZeroTrustAccessIdentityProviderModel struct {
	ID         types.String                                                             `tfsdk:"id" json:"id,computed"`
	AccountID  types.String                                                             `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID     types.String                                                             `tfsdk:"zone_id" path:"zone_id,optional"`
	Name       types.String                                                             `tfsdk:"name" json:"name,required"`
	Type       types.String                                                             `tfsdk:"type" json:"type,required"`
	Config     *ZeroTrustAccessIdentityProviderConfigModel                              `tfsdk:"config" json:"config,required"`
	SCIMConfig customfield.NestedObject[ZeroTrustAccessIdentityProviderSCIMConfigModel] `tfsdk:"scim_config" json:"scim_config,computed_optional"`
}

func (m ZeroTrustAccessIdentityProviderModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessIdentityProviderModel) MarshalJSONForUpdate(state ZeroTrustAccessIdentityProviderModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessIdentityProviderConfigModel struct {
	Claims                   *[]types.String                                                `tfsdk:"claims" json:"claims,optional"`
	ClientID                 types.String                                                   `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret             types.String                                                   `tfsdk:"client_secret" json:"client_secret,optional"`
	ConditionalAccessEnabled types.Bool                                                     `tfsdk:"conditional_access_enabled" json:"conditional_access_enabled,optional"`
	DirectoryID              types.String                                                   `tfsdk:"directory_id" json:"directory_id,optional"`
	EmailClaimName           types.String                                                   `tfsdk:"email_claim_name" json:"email_claim_name,optional"`
	Prompt                   types.String                                                   `tfsdk:"prompt" json:"prompt,optional"`
	SupportGroups            types.Bool                                                     `tfsdk:"support_groups" json:"support_groups,optional"`
	CentrifyAccount          types.String                                                   `tfsdk:"centrify_account" json:"centrify_account,optional"`
	CentrifyAppID            types.String                                                   `tfsdk:"centrify_app_id" json:"centrify_app_id,optional"`
	AppsDomain               types.String                                                   `tfsdk:"apps_domain" json:"apps_domain,optional"`
	AuthURL                  types.String                                                   `tfsdk:"auth_url" json:"auth_url,optional"`
	CERTsURL                 types.String                                                   `tfsdk:"certs_url" json:"certs_url,optional"`
	Scopes                   *[]types.String                                                `tfsdk:"scopes" json:"scopes,optional"`
	TokenURL                 types.String                                                   `tfsdk:"token_url" json:"token_url,optional"`
	AuthorizationServerID    types.String                                                   `tfsdk:"authorization_server_id" json:"authorization_server_id,optional"`
	OktaAccount              types.String                                                   `tfsdk:"okta_account" json:"okta_account,optional"`
	OneloginAccount          types.String                                                   `tfsdk:"onelogin_account" json:"onelogin_account,optional"`
	PingEnvID                types.String                                                   `tfsdk:"ping_env_id" json:"ping_env_id,optional"`
	Attributes               *[]types.String                                                `tfsdk:"attributes" json:"attributes,optional"`
	EmailAttributeName       types.String                                                   `tfsdk:"email_attribute_name" json:"email_attribute_name,optional"`
	HeaderAttributes         *[]*ZeroTrustAccessIdentityProviderConfigHeaderAttributesModel `tfsdk:"header_attributes" json:"header_attributes,optional"`
	IdPPublicCERTs           *[]types.String                                                `tfsdk:"idp_public_certs" json:"idp_public_certs,optional"`
	IssuerURL                types.String                                                   `tfsdk:"issuer_url" json:"issuer_url,optional"`
	SignRequest              types.Bool                                                     `tfsdk:"sign_request" json:"sign_request,optional"`
	SSOTargetURL             types.String                                                   `tfsdk:"sso_target_url" json:"sso_target_url,optional"`
	RedirectURL              types.String                                                   `tfsdk:"redirect_url" json:"redirect_url,computed"`
}

type ZeroTrustAccessIdentityProviderConfigHeaderAttributesModel struct {
	AttributeName types.String `tfsdk:"attribute_name" json:"attribute_name,optional"`
	HeaderName    types.String `tfsdk:"header_name" json:"header_name,optional"`
}

type ZeroTrustAccessIdentityProviderSCIMConfigModel struct {
	Enabled                types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	GroupMemberDeprovision types.Bool   `tfsdk:"group_member_deprovision" json:"group_member_deprovision,optional"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision" json:"seat_deprovision,optional"`
	Secret                 types.String `tfsdk:"secret" json:"secret,computed"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision" json:"user_deprovision,optional"`
}
