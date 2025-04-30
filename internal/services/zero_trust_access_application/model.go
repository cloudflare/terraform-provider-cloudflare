// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationResultEnvelope struct {
	Result ZeroTrustAccessApplicationModel `json:"result"`
}

type ZeroTrustAccessApplicationModel struct {
	ID                          types.String                                                               `tfsdk:"id" json:"id,computed"`
	AccountID                   types.String                                                               `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                      types.String                                                               `tfsdk:"zone_id" path:"zone_id,optional"`
	AllowAuthenticateViaWARP    types.Bool                                                                 `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,optional"`
	AllowIframe                 types.Bool                                                                 `tfsdk:"allow_iframe" json:"allow_iframe,optional"`
	AppLauncherLogoURL          types.String                                                               `tfsdk:"app_launcher_logo_url" json:"app_launcher_logo_url,optional"`
	BgColor                     types.String                                                               `tfsdk:"bg_color" json:"bg_color,optional"`
	CustomDenyMessage           types.String                                                               `tfsdk:"custom_deny_message" json:"custom_deny_message,optional"`
	CustomDenyURL               types.String                                                               `tfsdk:"custom_deny_url" json:"custom_deny_url,optional"`
	CustomNonIdentityDenyURL    types.String                                                               `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url,optional"`
	Domain                      types.String                                                               `tfsdk:"domain" json:"domain,optional"`
	HeaderBgColor               types.String                                                               `tfsdk:"header_bg_color" json:"header_bg_color,optional"`
	LogoURL                     types.String                                                               `tfsdk:"logo_url" json:"logo_url,optional"`
	Name                        types.String                                                               `tfsdk:"name" json:"name,optional"`
	OptionsPreflightBypass      types.Bool                                                                 `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass,optional"`
	ReadServiceTokensFromHeader types.String                                                               `tfsdk:"read_service_tokens_from_header" json:"read_service_tokens_from_header,optional"`
	SameSiteCookieAttribute     types.String                                                               `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute,optional"`
	ServiceAuth401Redirect      types.Bool                                                                 `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect,optional"`
	SkipInterstitial            types.Bool                                                                 `tfsdk:"skip_interstitial" json:"skip_interstitial,optional"`
	Type                        types.String                                                               `tfsdk:"type" json:"type,optional"`
	AllowedIdPs                 *[]types.String                                                            `tfsdk:"allowed_idps" json:"allowed_idps,optional"`
	CustomPages                 *[]types.String                                                            `tfsdk:"custom_pages" json:"custom_pages,optional"`
	SelfHostedDomains           *[]types.String                                                            `tfsdk:"self_hosted_domains" json:"self_hosted_domains,optional"`
	Tags                        *[]types.String                                                            `tfsdk:"tags" json:"tags,optional"`
	CORSHeaders                 *ZeroTrustAccessApplicationCORSHeadersModel                                `tfsdk:"cors_headers" json:"cors_headers,optional"`
	Destinations                *[]*ZeroTrustAccessApplicationDestinationsModel                            `tfsdk:"destinations" json:"destinations,optional"`
	FooterLinks                 *[]*ZeroTrustAccessApplicationFooterLinksModel                             `tfsdk:"footer_links" json:"footer_links,optional"`
	Policies                    *[]*ZeroTrustAccessApplicationPoliciesModel                                `tfsdk:"policies" json:"policies,optional"`
	SCIMConfig                  *ZeroTrustAccessApplicationSCIMConfigModel                                 `tfsdk:"scim_config" json:"scim_config,optional"`
	TargetCriteria              *[]*ZeroTrustAccessApplicationTargetCriteriaModel                          `tfsdk:"target_criteria" json:"target_criteria,optional"`
	AppLauncherVisible          types.Bool                                                                 `tfsdk:"app_launcher_visible" json:"app_launcher_visible,computed_optional"`
	AutoRedirectToIdentity      types.Bool                                                                 `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed_optional"`
	EnableBindingCookie         types.Bool                                                                 `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed_optional"`
	HTTPOnlyCookieAttribute     types.Bool                                                                 `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute,computed_optional"`
	PathCookieAttribute         types.Bool                                                                 `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute,computed_optional"`
	SessionDuration             types.String                                                               `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	SkipAppLauncherLoginPage    types.Bool                                                                 `tfsdk:"skip_app_launcher_login_page" json:"skip_app_launcher_login_page,computed_optional"`
	LandingPageDesign           customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignModel] `tfsdk:"landing_page_design" json:"landing_page_design,computed_optional"`
	SaaSApp                     customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppModel]           `tfsdk:"saas_app" json:"saas_app,computed_optional"`
	AUD                         types.String                                                               `tfsdk:"aud" json:"aud,computed"`
	CreatedAt                   timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt                   timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessApplicationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessApplicationModel) MarshalJSONForUpdate(state ZeroTrustAccessApplicationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessApplicationCORSHeadersModel struct {
	AllowAllHeaders  types.Bool      `tfsdk:"allow_all_headers" json:"allow_all_headers,optional"`
	AllowAllMethods  types.Bool      `tfsdk:"allow_all_methods" json:"allow_all_methods,optional"`
	AllowAllOrigins  types.Bool      `tfsdk:"allow_all_origins" json:"allow_all_origins,optional"`
	AllowCredentials types.Bool      `tfsdk:"allow_credentials" json:"allow_credentials,optional"`
	AllowedHeaders   *[]types.String `tfsdk:"allowed_headers" json:"allowed_headers,optional"`
	AllowedMethods   *[]types.String `tfsdk:"allowed_methods" json:"allowed_methods,optional"`
	AllowedOrigins   *[]types.String `tfsdk:"allowed_origins" json:"allowed_origins,optional"`
	MaxAge           types.Float64   `tfsdk:"max_age" json:"max_age,optional"`
}

type ZeroTrustAccessApplicationDestinationsModel struct {
	Type       types.String `tfsdk:"type" json:"type,optional"`
	URI        types.String `tfsdk:"uri" json:"uri,optional"`
	CIDR       types.String `tfsdk:"cidr" json:"cidr,optional"`
	Hostname   types.String `tfsdk:"hostname" json:"hostname,optional"`
	L4Protocol types.String `tfsdk:"l4_protocol" json:"l4_protocol,optional"`
	PortRange  types.String `tfsdk:"port_range" json:"port_range,optional"`
	VnetID     types.String `tfsdk:"vnet_id" json:"vnet_id,optional"`
}

type ZeroTrustAccessApplicationFooterLinksModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
	URL  types.String `tfsdk:"url" json:"url,required"`
}

type ZeroTrustAccessApplicationPoliciesModel struct {
	ID              types.String                                            `tfsdk:"id" json:"id,optional"`
	Precedence      types.Int64                                             `tfsdk:"precedence" json:"precedence,optional"`
	Decision        types.String                                            `tfsdk:"decision" json:"decision,optional"`
	Include         *[]*ZeroTrustAccessApplicationPoliciesIncludeModel      `tfsdk:"include" json:"include,optional"`
	Name            types.String                                            `tfsdk:"name" json:"name,optional"`
	ConnectionRules *ZeroTrustAccessApplicationPoliciesConnectionRulesModel `tfsdk:"connection_rules" json:"connection_rules,optional"`
	Exclude         *[]*ZeroTrustAccessApplicationPoliciesExcludeModel      `tfsdk:"exclude" json:"exclude,optional"`
	Require         *[]*ZeroTrustAccessApplicationPoliciesRequireModel      `tfsdk:"require" json:"require,optional"`
}

type ZeroTrustAccessApplicationPoliciesIncludeModel struct {
	Group                *ZeroTrustAccessApplicationPoliciesIncludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessApplicationPoliciesIncludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessApplicationPoliciesIncludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessApplicationPoliciesIncludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesIncludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessApplicationPoliciesIncludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessApplicationPoliciesIncludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessApplicationPoliciesIncludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessApplicationPoliciesIncludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessApplicationPoliciesIncludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessApplicationPoliciesIncludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesIncludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessApplicationPoliciesIncludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessApplicationPoliciesIncludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessApplicationPoliciesIncludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessApplicationPoliciesIncludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessApplicationPoliciesIncludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessApplicationPoliciesIncludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessApplicationPoliciesIncludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesModel struct {
	SSH *ZeroTrustAccessApplicationPoliciesConnectionRulesSSHModel `tfsdk:"ssh" json:"ssh,optional"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesSSHModel struct {
	Usernames       *[]types.String `tfsdk:"usernames" json:"usernames,required"`
	AllowEmailAlias types.Bool      `tfsdk:"allow_email_alias" json:"allow_email_alias,optional"`
}

type ZeroTrustAccessApplicationPoliciesExcludeModel struct {
	Group                *ZeroTrustAccessApplicationPoliciesExcludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessApplicationPoliciesExcludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessApplicationPoliciesExcludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessApplicationPoliciesExcludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesExcludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessApplicationPoliciesExcludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessApplicationPoliciesExcludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessApplicationPoliciesExcludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessApplicationPoliciesExcludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessApplicationPoliciesExcludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessApplicationPoliciesExcludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesExcludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessApplicationPoliciesExcludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessApplicationPoliciesExcludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessApplicationPoliciesExcludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessApplicationPoliciesExcludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessApplicationPoliciesExcludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessApplicationPoliciesExcludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessApplicationPoliciesExcludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireModel struct {
	Group                *ZeroTrustAccessApplicationPoliciesRequireGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessApplicationPoliciesRequireAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessApplicationPoliciesRequireAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessApplicationPoliciesRequireAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesRequireCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessApplicationPoliciesRequireCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessApplicationPoliciesRequireGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessApplicationPoliciesRequireDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessApplicationPoliciesRequireEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessApplicationPoliciesRequireEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessApplicationPoliciesRequireEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesRequireEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessApplicationPoliciesRequireGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessApplicationPoliciesRequireLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessApplicationPoliciesRequireIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessApplicationPoliciesRequireIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessApplicationPoliciesRequireOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessApplicationPoliciesRequireSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessApplicationPoliciesRequireServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessApplicationPoliciesRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessApplicationPoliciesRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessApplicationSCIMConfigModel struct {
	IdPUID             types.String                                             `tfsdk:"idp_uid" json:"idp_uid,required"`
	RemoteURI          types.String                                             `tfsdk:"remote_uri" json:"remote_uri,required"`
	Authentication     *ZeroTrustAccessApplicationSCIMConfigAuthenticationModel `tfsdk:"authentication" json:"authentication,optional"`
	DeactivateOnDelete types.Bool                                               `tfsdk:"deactivate_on_delete" json:"deactivate_on_delete,optional"`
	Enabled            types.Bool                                               `tfsdk:"enabled" json:"enabled,optional"`
	Mappings           *[]*ZeroTrustAccessApplicationSCIMConfigMappingsModel    `tfsdk:"mappings" json:"mappings,optional"`
}

type ZeroTrustAccessApplicationSCIMConfigAuthenticationModel struct {
	Password         types.String    `tfsdk:"password" json:"password,optional"`
	Scheme           types.String    `tfsdk:"scheme" json:"scheme,required"`
	User             types.String    `tfsdk:"user" json:"user,optional"`
	Token            types.String    `tfsdk:"token" json:"token,optional"`
	AuthorizationURL types.String    `tfsdk:"authorization_url" json:"authorization_url,optional"`
	ClientID         types.String    `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret     types.String    `tfsdk:"client_secret" json:"client_secret,optional"`
	TokenURL         types.String    `tfsdk:"token_url" json:"token_url,optional"`
	Scopes           *[]types.String `tfsdk:"scopes" json:"scopes,optional"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsModel struct {
	Schema           types.String                                                 `tfsdk:"schema" json:"schema,required"`
	Enabled          types.Bool                                                   `tfsdk:"enabled" json:"enabled,optional"`
	Filter           types.String                                                 `tfsdk:"filter" json:"filter,optional"`
	Operations       *ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel `tfsdk:"operations" json:"operations,optional"`
	Strictness       types.String                                                 `tfsdk:"strictness" json:"strictness,optional"`
	TransformJsonata types.String                                                 `tfsdk:"transform_jsonata" json:"transform_jsonata,optional"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel struct {
	Create types.Bool `tfsdk:"create" json:"create,optional"`
	Delete types.Bool `tfsdk:"delete" json:"delete,optional"`
	Update types.Bool `tfsdk:"update" json:"update,optional"`
}

type ZeroTrustAccessApplicationTargetCriteriaModel struct {
	Port             types.Int64                 `tfsdk:"port" json:"port,required"`
	Protocol         types.String                `tfsdk:"protocol" json:"protocol,required"`
	TargetAttributes *map[string]*[]types.String `tfsdk:"target_attributes" json:"target_attributes,required"`
}

type ZeroTrustAccessApplicationLandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color" json:"button_color,optional"`
	ButtonTextColor types.String `tfsdk:"button_text_color" json:"button_text_color,optional"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,optional"`
	Message         types.String `tfsdk:"message" json:"message,optional"`
	Title           types.String `tfsdk:"title" json:"title,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppModel struct {
	AuthType                      types.String                                                    `tfsdk:"auth_type" json:"auth_type,computed_optional"`
	ConsumerServiceURL            types.String                                                    `tfsdk:"consumer_service_url" json:"consumer_service_url,optional"`
	CreatedAt                     timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomAttributes              *[]*ZeroTrustAccessApplicationSaaSAppCustomAttributesModel      `tfsdk:"custom_attributes" json:"custom_attributes,optional"`
	DefaultRelayState             types.String                                                    `tfsdk:"default_relay_state" json:"default_relay_state,optional"`
	IdPEntityID                   types.String                                                    `tfsdk:"idp_entity_id" json:"idp_entity_id,computed_optional"`
	NameIDFormat                  types.String                                                    `tfsdk:"name_id_format" json:"name_id_format,optional"`
	NameIDTransformJsonata        types.String                                                    `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata,optional"`
	PublicKey                     types.String                                                    `tfsdk:"public_key" json:"public_key,computed_optional"`
	SAMLAttributeTransformJsonata types.String                                                    `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata,optional"`
	SPEntityID                    types.String                                                    `tfsdk:"sp_entity_id" json:"sp_entity_id,optional"`
	SSOEndpoint                   types.String                                                    `tfsdk:"sso_endpoint" json:"sso_endpoint,computed_optional"`
	UpdatedAt                     timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AccessTokenLifetime           types.String                                                    `tfsdk:"access_token_lifetime" json:"access_token_lifetime,optional"`
	AllowPKCEWithoutClientSecret  types.Bool                                                      `tfsdk:"allow_pkce_without_client_secret" json:"allow_pkce_without_client_secret,optional"`
	AppLauncherURL                types.String                                                    `tfsdk:"app_launcher_url" json:"app_launcher_url,optional"`
	ClientID                      types.String                                                    `tfsdk:"client_id" json:"client_id,computed_optional"`
	ClientSecret                  types.String                                                    `tfsdk:"client_secret" json:"client_secret,computed_optional"`
	CustomClaims                  *[]*ZeroTrustAccessApplicationSaaSAppCustomClaimsModel          `tfsdk:"custom_claims" json:"custom_claims,optional"`
	GrantTypes                    *[]types.String                                                 `tfsdk:"grant_types" json:"grant_types,optional"`
	GroupFilterRegex              types.String                                                    `tfsdk:"group_filter_regex" json:"group_filter_regex,optional"`
	HybridAndImplicitOptions      *ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel `tfsdk:"hybrid_and_implicit_options" json:"hybrid_and_implicit_options,optional"`
	RedirectURIs                  *[]types.String                                                 `tfsdk:"redirect_uris" json:"redirect_uris,optional"`
	RefreshTokenOptions           *ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel      `tfsdk:"refresh_token_options" json:"refresh_token_options,optional"`
	Scopes                        *[]types.String                                                 `tfsdk:"scopes" json:"scopes,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesModel struct {
	FriendlyName types.String                                                  `tfsdk:"friendly_name" json:"friendly_name,optional"`
	Name         types.String                                                  `tfsdk:"name" json:"name,optional"`
	NameFormat   types.String                                                  `tfsdk:"name_format" json:"name_format,optional"`
	Required     types.Bool                                                    `tfsdk:"required" json:"required,optional"`
	Source       *ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel `tfsdk:"source" json:"source,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel struct {
	Name      types.String                                                              `tfsdk:"name" json:"name,optional"`
	NameByIdP *[]*ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPModel `tfsdk:"name_by_idp" json:"name_by_idp,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceNameByIdPModel struct {
	IdPID      types.String `tfsdk:"idp_id" json:"idp_id,optional"`
	SourceName types.String `tfsdk:"source_name" json:"source_name,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsModel struct {
	Name     types.String                                              `tfsdk:"name" json:"name,optional"`
	Required types.Bool                                                `tfsdk:"required" json:"required,optional"`
	Scope    types.String                                              `tfsdk:"scope" json:"scope,optional"`
	Source   *ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel `tfsdk:"source" json:"source,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel struct {
	Name      types.String             `tfsdk:"name" json:"name,optional"`
	NameByIdP *map[string]types.String `tfsdk:"name_by_idp" json:"name_by_idp,optional"`
}

type ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint" json:"return_access_token_from_authorization_endpoint,optional"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint" json:"return_id_token_from_authorization_endpoint,optional"`
}

type ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel struct {
	Lifetime types.String `tfsdk:"lifetime" json:"lifetime,optional"`
}
