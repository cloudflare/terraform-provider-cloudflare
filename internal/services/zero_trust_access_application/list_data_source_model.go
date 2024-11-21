// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessApplicationsDataSourceModel struct {
	AccountID types.String                                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	AUD       types.String                                                                   `tfsdk:"aud" query:"aud,optional"`
	Domain    types.String                                                                   `tfsdk:"domain" query:"domain,optional"`
	Name      types.String                                                                   `tfsdk:"name" query:"name,optional"`
	Search    types.String                                                                   `tfsdk:"search" query:"search,optional"`
	MaxItems  types.Int64                                                                    `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessApplicationsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessApplicationsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationListParams{}

	if !m.AUD.IsNull() {
		params.AUD = cloudflare.F(m.AUD.ValueString())
	}
	if !m.Domain.IsNull() {
		params.Domain = cloudflare.F(m.Domain.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}
	if !m.Search.IsNull() {
		params.Search = cloudflare.F(m.Search.ValueString())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessApplicationsResultDataSourceModel struct {
	Domain                   types.String                                                                           `tfsdk:"domain" json:"domain,computed"`
	Type                     types.String                                                                           `tfsdk:"type" json:"type,computed"`
	ID                       types.String                                                                           `tfsdk:"id" json:"id,computed"`
	AllowAuthenticateViaWARP types.Bool                                                                             `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,computed"`
	AllowedIdPs              customfield.List[types.String]                                                         `tfsdk:"allowed_idps" json:"allowed_idps,computed"`
	AppLauncherVisible       types.Bool                                                                             `tfsdk:"app_launcher_visible" json:"app_launcher_visible,computed"`
	AUD                      types.String                                                                           `tfsdk:"aud" json:"aud,computed"`
	AutoRedirectToIdentity   types.Bool                                                                             `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed"`
	CORSHeaders              customfield.NestedObject[ZeroTrustAccessApplicationsCORSHeadersDataSourceModel]        `tfsdk:"cors_headers" json:"cors_headers,computed"`
	CreatedAt                timetypes.RFC3339                                                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomDenyMessage        types.String                                                                           `tfsdk:"custom_deny_message" json:"custom_deny_message,computed"`
	CustomDenyURL            types.String                                                                           `tfsdk:"custom_deny_url" json:"custom_deny_url,computed"`
	CustomNonIdentityDenyURL types.String                                                                           `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url,computed"`
	CustomPages              customfield.List[types.String]                                                         `tfsdk:"custom_pages" json:"custom_pages,computed"`
	Destinations             customfield.NestedObjectList[ZeroTrustAccessApplicationsDestinationsDataSourceModel]   `tfsdk:"destinations" json:"destinations,computed"`
	EnableBindingCookie      types.Bool                                                                             `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed"`
	HTTPOnlyCookieAttribute  types.Bool                                                                             `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute,computed"`
	LogoURL                  types.String                                                                           `tfsdk:"logo_url" json:"logo_url,computed"`
	Name                     types.String                                                                           `tfsdk:"name" json:"name,computed"`
	OptionsPreflightBypass   types.Bool                                                                             `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass,computed"`
	PathCookieAttribute      types.Bool                                                                             `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute,computed"`
	Policies                 customfield.NestedObjectList[ZeroTrustAccessApplicationsPoliciesDataSourceModel]       `tfsdk:"policies" json:"policies,computed"`
	SameSiteCookieAttribute  types.String                                                                           `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute,computed"`
	SCIMConfig               customfield.NestedObject[ZeroTrustAccessApplicationsSCIMConfigDataSourceModel]         `tfsdk:"scim_config" json:"scim_config,computed"`
	SelfHostedDomains        customfield.List[types.String]                                                         `tfsdk:"self_hosted_domains" json:"self_hosted_domains,computed"`
	ServiceAuth401Redirect   types.Bool                                                                             `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect,computed"`
	SessionDuration          types.String                                                                           `tfsdk:"session_duration" json:"session_duration,computed"`
	SkipInterstitial         types.Bool                                                                             `tfsdk:"skip_interstitial" json:"skip_interstitial,computed"`
	Tags                     customfield.List[types.String]                                                         `tfsdk:"tags" json:"tags,computed"`
	UpdatedAt                timetypes.RFC3339                                                                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	SaaSApp                  customfield.NestedObject[ZeroTrustAccessApplicationsSaaSAppDataSourceModel]            `tfsdk:"saas_app" json:"saas_app,computed"`
	AppLauncherLogoURL       types.String                                                                           `tfsdk:"app_launcher_logo_url" json:"app_launcher_logo_url,computed"`
	BgColor                  types.String                                                                           `tfsdk:"bg_color" json:"bg_color,computed"`
	FooterLinks              customfield.NestedObjectList[ZeroTrustAccessApplicationsFooterLinksDataSourceModel]    `tfsdk:"footer_links" json:"footer_links,computed"`
	HeaderBgColor            types.String                                                                           `tfsdk:"header_bg_color" json:"header_bg_color,computed"`
	LandingPageDesign        customfield.NestedObject[ZeroTrustAccessApplicationsLandingPageDesignDataSourceModel]  `tfsdk:"landing_page_design" json:"landing_page_design,computed"`
	SkipAppLauncherLoginPage types.Bool                                                                             `tfsdk:"skip_app_launcher_login_page" json:"skip_app_launcher_login_page,computed"`
	TargetCriteria           customfield.NestedObjectList[ZeroTrustAccessApplicationsTargetCriteriaDataSourceModel] `tfsdk:"target_criteria" json:"target_criteria,computed"`
}

type ZeroTrustAccessApplicationsCORSHeadersDataSourceModel struct {
	AllowAllHeaders  types.Bool                     `tfsdk:"allow_all_headers" json:"allow_all_headers,computed"`
	AllowAllMethods  types.Bool                     `tfsdk:"allow_all_methods" json:"allow_all_methods,computed"`
	AllowAllOrigins  types.Bool                     `tfsdk:"allow_all_origins" json:"allow_all_origins,computed"`
	AllowCredentials types.Bool                     `tfsdk:"allow_credentials" json:"allow_credentials,computed"`
	AllowedHeaders   customfield.List[types.String] `tfsdk:"allowed_headers" json:"allowed_headers,computed"`
	AllowedMethods   customfield.List[types.String] `tfsdk:"allowed_methods" json:"allowed_methods,computed"`
	AllowedOrigins   customfield.List[types.String] `tfsdk:"allowed_origins" json:"allowed_origins,computed"`
	MaxAge           types.Float64                  `tfsdk:"max_age" json:"max_age,computed"`
}

type ZeroTrustAccessApplicationsDestinationsDataSourceModel struct {
	Type types.String `tfsdk:"type" json:"type,computed"`
	URI  types.String `tfsdk:"uri" json:"uri,computed"`
}

type ZeroTrustAccessApplicationsPoliciesDataSourceModel struct {
	ID        types.String                                                                            `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339                                                                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Decision  types.String                                                                            `tfsdk:"decision" json:"decision,computed"`
	Exclude   customfield.NestedObjectList[ZeroTrustAccessApplicationsPoliciesExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
	Include   customfield.NestedObjectList[ZeroTrustAccessApplicationsPoliciesIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
	Name      types.String                                                                            `tfsdk:"name" json:"name,computed"`
	Require   customfield.NestedObjectList[ZeroTrustAccessApplicationsPoliciesRequireDataSourceModel] `tfsdk:"require" json:"require,computed"`
	UpdatedAt timetypes.RFC3339                                                                       `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesExcludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesExcludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesExcludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesIncludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesIncludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesIncludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationsPoliciesRequireServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesRequireAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesRequireCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationsPoliciesRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationsPoliciesRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationsSCIMConfigDataSourceModel struct {
	IdPUID             types.String                                                                                 `tfsdk:"idp_uid" json:"idp_uid,computed"`
	RemoteURI          types.String                                                                                 `tfsdk:"remote_uri" json:"remote_uri,computed"`
	Authentication     customfield.NestedObject[ZeroTrustAccessApplicationsSCIMConfigAuthenticationDataSourceModel] `tfsdk:"authentication" json:"authentication,computed"`
	DeactivateOnDelete types.Bool                                                                                   `tfsdk:"deactivate_on_delete" json:"deactivate_on_delete,computed"`
	Enabled            types.Bool                                                                                   `tfsdk:"enabled" json:"enabled,computed"`
	Mappings           customfield.NestedObjectList[ZeroTrustAccessApplicationsSCIMConfigMappingsDataSourceModel]   `tfsdk:"mappings" json:"mappings,computed"`
}

type ZeroTrustAccessApplicationsSCIMConfigAuthenticationDataSourceModel struct {
	Password         types.String                   `tfsdk:"password" json:"password,computed"`
	Scheme           types.String                   `tfsdk:"scheme" json:"scheme,computed"`
	User             types.String                   `tfsdk:"user" json:"user,computed"`
	Token            types.String                   `tfsdk:"token" json:"token,computed"`
	AuthorizationURL types.String                   `tfsdk:"authorization_url" json:"authorization_url,computed"`
	ClientID         types.String                   `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret     types.String                   `tfsdk:"client_secret" json:"client_secret,computed"`
	TokenURL         types.String                   `tfsdk:"token_url" json:"token_url,computed"`
	Scopes           customfield.List[types.String] `tfsdk:"scopes" json:"scopes,computed"`
}

type ZeroTrustAccessApplicationsSCIMConfigMappingsDataSourceModel struct {
	Schema           types.String                                                                                     `tfsdk:"schema" json:"schema,computed"`
	Enabled          types.Bool                                                                                       `tfsdk:"enabled" json:"enabled,computed"`
	Filter           types.String                                                                                     `tfsdk:"filter" json:"filter,computed"`
	Operations       customfield.NestedObject[ZeroTrustAccessApplicationsSCIMConfigMappingsOperationsDataSourceModel] `tfsdk:"operations" json:"operations,computed"`
	Strictness       types.String                                                                                     `tfsdk:"strictness" json:"strictness,computed"`
	TransformJsonata types.String                                                                                     `tfsdk:"transform_jsonata" json:"transform_jsonata,computed"`
}

type ZeroTrustAccessApplicationsSCIMConfigMappingsOperationsDataSourceModel struct {
	Create types.Bool `tfsdk:"create" json:"create,computed"`
	Delete types.Bool `tfsdk:"delete" json:"delete,computed"`
	Update types.Bool `tfsdk:"update" json:"update,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppDataSourceModel struct {
	AuthType                      types.String                                                                                        `tfsdk:"auth_type" json:"auth_type,computed"`
	ConsumerServiceURL            types.String                                                                                        `tfsdk:"consumer_service_url" json:"consumer_service_url,computed"`
	CreatedAt                     timetypes.RFC3339                                                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomAttributes              customfield.NestedObjectList[ZeroTrustAccessApplicationsSaaSAppCustomAttributesDataSourceModel]     `tfsdk:"custom_attributes" json:"custom_attributes,computed"`
	DefaultRelayState             types.String                                                                                        `tfsdk:"default_relay_state" json:"default_relay_state,computed"`
	IdPEntityID                   types.String                                                                                        `tfsdk:"idp_entity_id" json:"idp_entity_id,computed"`
	NameIDFormat                  types.String                                                                                        `tfsdk:"name_id_format" json:"name_id_format,computed"`
	NameIDTransformJsonata        types.String                                                                                        `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata,computed"`
	PublicKey                     types.String                                                                                        `tfsdk:"public_key" json:"public_key,computed"`
	SAMLAttributeTransformJsonata types.String                                                                                        `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata,computed"`
	SPEntityID                    types.String                                                                                        `tfsdk:"sp_entity_id" json:"sp_entity_id,computed"`
	SSOEndpoint                   types.String                                                                                        `tfsdk:"sso_endpoint" json:"sso_endpoint,computed"`
	UpdatedAt                     timetypes.RFC3339                                                                                   `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AccessTokenLifetime           types.String                                                                                        `tfsdk:"access_token_lifetime" json:"access_token_lifetime,computed"`
	AllowPKCEWithoutClientSecret  types.Bool                                                                                          `tfsdk:"allow_pkce_without_client_secret" json:"allow_pkce_without_client_secret,computed"`
	AppLauncherURL                types.String                                                                                        `tfsdk:"app_launcher_url" json:"app_launcher_url,computed"`
	ClientID                      types.String                                                                                        `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret                  types.String                                                                                        `tfsdk:"client_secret" json:"client_secret,computed"`
	CustomClaims                  customfield.NestedObjectList[ZeroTrustAccessApplicationsSaaSAppCustomClaimsDataSourceModel]         `tfsdk:"custom_claims" json:"custom_claims,computed"`
	GrantTypes                    customfield.List[types.String]                                                                      `tfsdk:"grant_types" json:"grant_types,computed"`
	GroupFilterRegex              types.String                                                                                        `tfsdk:"group_filter_regex" json:"group_filter_regex,computed"`
	HybridAndImplicitOptions      customfield.NestedObject[ZeroTrustAccessApplicationsSaaSAppHybridAndImplicitOptionsDataSourceModel] `tfsdk:"hybrid_and_implicit_options" json:"hybrid_and_implicit_options,computed"`
	RedirectURIs                  customfield.List[types.String]                                                                      `tfsdk:"redirect_uris" json:"redirect_uris,computed"`
	RefreshTokenOptions           customfield.NestedObject[ZeroTrustAccessApplicationsSaaSAppRefreshTokenOptionsDataSourceModel]      `tfsdk:"refresh_token_options" json:"refresh_token_options,computed"`
	Scopes                        customfield.List[types.String]                                                                      `tfsdk:"scopes" json:"scopes,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppCustomAttributesDataSourceModel struct {
	FriendlyName types.String                                                                                      `tfsdk:"friendly_name" json:"friendly_name,computed"`
	Name         types.String                                                                                      `tfsdk:"name" json:"name,computed"`
	NameFormat   types.String                                                                                      `tfsdk:"name_format" json:"name_format,computed"`
	Required     types.Bool                                                                                        `tfsdk:"required" json:"required,computed"`
	Source       customfield.NestedObject[ZeroTrustAccessApplicationsSaaSAppCustomAttributesSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppCustomAttributesSourceDataSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppCustomClaimsDataSourceModel struct {
	Name     types.String                                                                                  `tfsdk:"name" json:"name,computed"`
	Required types.Bool                                                                                    `tfsdk:"required" json:"required,computed"`
	Scope    types.String                                                                                  `tfsdk:"scope" json:"scope,computed"`
	Source   customfield.NestedObject[ZeroTrustAccessApplicationsSaaSAppCustomClaimsSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppCustomClaimsSourceDataSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppHybridAndImplicitOptionsDataSourceModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint" json:"return_access_token_from_authorization_endpoint,computed"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint" json:"return_id_token_from_authorization_endpoint,computed"`
}

type ZeroTrustAccessApplicationsSaaSAppRefreshTokenOptionsDataSourceModel struct {
	Lifetime types.String `tfsdk:"lifetime" json:"lifetime,computed"`
}

type ZeroTrustAccessApplicationsFooterLinksDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	URL  types.String `tfsdk:"url" json:"url,computed"`
}

type ZeroTrustAccessApplicationsLandingPageDesignDataSourceModel struct {
	ButtonColor     types.String `tfsdk:"button_color" json:"button_color,computed"`
	ButtonTextColor types.String `tfsdk:"button_text_color" json:"button_text_color,computed"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,computed"`
	Message         types.String `tfsdk:"message" json:"message,computed"`
	Title           types.String `tfsdk:"title" json:"title,computed"`
}

type ZeroTrustAccessApplicationsTargetCriteriaDataSourceModel struct {
	Port             types.Int64                                     `tfsdk:"port" json:"port,computed"`
	Protocol         types.String                                    `tfsdk:"protocol" json:"protocol,computed"`
	TargetAttributes customfield.Map[customfield.List[types.String]] `tfsdk:"target_attributes" json:"target_attributes,computed"`
}
