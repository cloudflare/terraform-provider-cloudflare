// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_identity_provider

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessIdentityProviderResultDataSourceEnvelope struct {
	Result ZeroTrustAccessIdentityProviderDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessIdentityProviderDataSourceModel struct {
	ID                   types.String                                                                               `tfsdk:"id" path:"identity_provider_id,computed"`
	IdentityProviderID   types.String                                                                               `tfsdk:"identity_provider_id" path:"identity_provider_id,optional"`
	AccountID            types.String                                                                               `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID               types.String                                                                               `tfsdk:"zone_id" path:"zone_id,optional"`
	Name                 types.String                                                                               `tfsdk:"name" json:"name,computed"`
	SAMLCertificateSetID types.String                                                                               `tfsdk:"saml_certificate_set_id" json:"saml_certificate_set_id,computed"`
	Type                 types.String                                                                               `tfsdk:"type" json:"type,computed"`
	Config               customfield.NestedObject[ZeroTrustAccessIdentityProviderConfigDataSourceModel]             `tfsdk:"config" json:"config,computed"`
	SAMLCertificateSet   customfield.NestedObject[ZeroTrustAccessIdentityProviderSAMLCertificateSetDataSourceModel] `tfsdk:"saml_certificate_set" json:"saml_certificate_set,computed"`
	SCIMConfig           customfield.NestedObject[ZeroTrustAccessIdentityProviderSCIMConfigDataSourceModel]         `tfsdk:"scim_config" json:"scim_config,computed"`
	Filter               *ZeroTrustAccessIdentityProviderFindOneByDataSourceModel                                   `tfsdk:"filter"`
}

func (m *ZeroTrustAccessIdentityProviderDataSourceModel) toReadParams(_ context.Context) (params zero_trust.IdentityProviderGetParams, diags diag.Diagnostics) {
	params = zero_trust.IdentityProviderGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessIdentityProviderDataSourceModel) toListParams(_ context.Context) (params zero_trust.IdentityProviderListParams, diags diag.Diagnostics) {
	params = zero_trust.IdentityProviderListParams{}

	if !m.Filter.SCIMEnabled.IsNull() {
		params.SCIMEnabled = cloudflare.F(m.Filter.SCIMEnabled.ValueString())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessIdentityProviderConfigDataSourceModel struct {
	Claims                   customfield.List[types.String]                                                                     `tfsdk:"claims" json:"claims,computed"`
	ClientID                 types.String                                                                                       `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret             types.String                                                                                       `tfsdk:"client_secret" json:"client_secret,computed"`
	ConditionalAccessEnabled types.Bool                                                                                         `tfsdk:"conditional_access_enabled" json:"conditional_access_enabled,computed"`
	DirectoryID              types.String                                                                                       `tfsdk:"directory_id" json:"directory_id,computed"`
	EmailClaimName           types.String                                                                                       `tfsdk:"email_claim_name" json:"email_claim_name,computed"`
	Prompt                   types.String                                                                                       `tfsdk:"prompt" json:"prompt,computed"`
	SupportGroups            types.Bool                                                                                         `tfsdk:"support_groups" json:"support_groups,computed"`
	CentrifyAccount          types.String                                                                                       `tfsdk:"centrify_account" json:"centrify_account,computed"`
	CentrifyAppID            types.String                                                                                       `tfsdk:"centrify_app_id" json:"centrify_app_id,computed"`
	AppsDomain               types.String                                                                                       `tfsdk:"apps_domain" json:"apps_domain,computed"`
	AuthURL                  types.String                                                                                       `tfsdk:"auth_url" json:"auth_url,computed"`
	CERTsURL                 types.String                                                                                       `tfsdk:"certs_url" json:"certs_url,computed"`
	PKCEEnabled              types.Bool                                                                                         `tfsdk:"pkce_enabled" json:"pkce_enabled,computed"`
	Scopes                   customfield.List[types.String]                                                                     `tfsdk:"scopes" json:"scopes,computed"`
	TokenURL                 types.String                                                                                       `tfsdk:"token_url" json:"token_url,computed"`
	AuthorizationServerID    types.String                                                                                       `tfsdk:"authorization_server_id" json:"authorization_server_id,computed"`
	OktaAccount              types.String                                                                                       `tfsdk:"okta_account" json:"okta_account,computed"`
	OneloginAccount          types.String                                                                                       `tfsdk:"onelogin_account" json:"onelogin_account,computed"`
	PingEnvID                types.String                                                                                       `tfsdk:"ping_env_id" json:"ping_env_id,computed"`
	Attributes               customfield.List[types.String]                                                                     `tfsdk:"attributes" json:"attributes,computed"`
	EmailAttributeName       types.String                                                                                       `tfsdk:"email_attribute_name" json:"email_attribute_name,computed"`
	EnableEncryption         types.Bool                                                                                         `tfsdk:"enable_encryption" json:"enable_encryption,computed"`
	HeaderAttributes         customfield.NestedObjectList[ZeroTrustAccessIdentityProviderConfigHeaderAttributesDataSourceModel] `tfsdk:"header_attributes" json:"header_attributes,computed"`
	IdPPublicCERTs           customfield.List[types.String]                                                                     `tfsdk:"idp_public_certs" json:"idp_public_certs,computed"`
	IssuerURL                types.String                                                                                       `tfsdk:"issuer_url" json:"issuer_url,computed"`
	SignRequest              types.Bool                                                                                         `tfsdk:"sign_request" json:"sign_request,computed"`
	SSOTargetURL             types.String                                                                                       `tfsdk:"sso_target_url" json:"sso_target_url,computed"`
	RedirectURL              types.String                                                                                       `tfsdk:"redirect_url" json:"redirect_url,computed"`
	RestrictToAccountMembers types.Bool                                                                                         `tfsdk:"restrict_to_account_members" json:"restrict_to_account_members,computed"`
}

type ZeroTrustAccessIdentityProviderConfigHeaderAttributesDataSourceModel struct {
	AttributeName types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	HeaderName    types.String `tfsdk:"header_name" json:"header_name,computed"`
}

type ZeroTrustAccessIdentityProviderSAMLCertificateSetDataSourceModel struct {
	CreatedAt           timetypes.RFC3339                                                                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UID                 types.String                                                                                                 `tfsdk:"uid" json:"uid,computed"`
	UpdatedAt           timetypes.RFC3339                                                                                            `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	CurrentCertificate  customfield.NestedObject[ZeroTrustAccessIdentityProviderSAMLCertificateSetCurrentCertificateDataSourceModel] `tfsdk:"current_certificate" json:"current_certificate,computed"`
	PreviousCertificate jsontypes.Normalized                                                                                         `tfsdk:"previous_certificate" json:"previous_certificate,computed"`
}

type ZeroTrustAccessIdentityProviderSAMLCertificateSetCurrentCertificateDataSourceModel struct {
	IsCurrent         types.Bool        `tfsdk:"is_current" json:"is_current,computed"`
	NotAfter          timetypes.RFC3339 `tfsdk:"not_after" json:"not_after,computed" format:"date-time"`
	PublicCertificate types.String      `tfsdk:"public_certificate" json:"public_certificate,computed"`
	UID               types.String      `tfsdk:"uid" json:"uid,computed"`
}

type ZeroTrustAccessIdentityProviderSCIMConfigDataSourceModel struct {
	Enabled                types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	IdentityUpdateBehavior types.String `tfsdk:"identity_update_behavior" json:"identity_update_behavior,computed"`
	SCIMBaseURL            types.String `tfsdk:"scim_base_url" json:"scim_base_url,computed"`
	SeatDeprovision        types.Bool   `tfsdk:"seat_deprovision" json:"seat_deprovision,computed"`
	Secret                 types.String `tfsdk:"secret" json:"secret,computed"`
	UserDeprovision        types.Bool   `tfsdk:"user_deprovision" json:"user_deprovision,computed"`
}

type ZeroTrustAccessIdentityProviderFindOneByDataSourceModel struct {
	SCIMEnabled types.String `tfsdk:"scim_enabled" query:"scim_enabled,optional"`
}
