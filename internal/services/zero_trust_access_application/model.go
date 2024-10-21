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
	ID                       types.String                                                               `tfsdk:"id" json:"id,computed"`
	AccountID                types.String                                                               `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                   types.String                                                               `tfsdk:"zone_id" path:"zone_id,optional"`
	AllowAuthenticateViaWARP types.Bool                                                                 `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,optional"`
	AppLauncherLogoURL       types.String                                                               `tfsdk:"app_launcher_logo_url" json:"app_launcher_logo_url,optional"`
	BgColor                  types.String                                                               `tfsdk:"bg_color" json:"bg_color,optional"`
	CustomDenyMessage        types.String                                                               `tfsdk:"custom_deny_message" json:"custom_deny_message,optional"`
	CustomDenyURL            types.String                                                               `tfsdk:"custom_deny_url" json:"custom_deny_url,optional"`
	CustomNonIdentityDenyURL types.String                                                               `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url,optional"`
	Domain                   types.String                                                               `tfsdk:"domain" json:"domain,optional"`
	HeaderBgColor            types.String                                                               `tfsdk:"header_bg_color" json:"header_bg_color,optional"`
	LogoURL                  types.String                                                               `tfsdk:"logo_url" json:"logo_url,optional"`
	Name                     types.String                                                               `tfsdk:"name" json:"name,optional"`
	OptionsPreflightBypass   types.Bool                                                                 `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass,optional"`
	SameSiteCookieAttribute  types.String                                                               `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute,optional"`
	ServiceAuth401Redirect   types.Bool                                                                 `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect,optional"`
	SkipInterstitial         types.Bool                                                                 `tfsdk:"skip_interstitial" json:"skip_interstitial,optional"`
	Type                     types.String                                                               `tfsdk:"type" json:"type,optional"`
	AllowedIdPs              *[]types.String                                                            `tfsdk:"allowed_idps" json:"allowed_idps,optional"`
	CustomPages              *[]types.String                                                            `tfsdk:"custom_pages" json:"custom_pages,optional"`
	SelfHostedDomains        *[]types.String                                                            `tfsdk:"self_hosted_domains" json:"self_hosted_domains,optional"`
	Tags                     *[]types.String                                                            `tfsdk:"tags" json:"tags,optional"`
	AppLauncherVisible       types.Bool                                                                 `tfsdk:"app_launcher_visible" json:"app_launcher_visible,computed_optional"`
	AutoRedirectToIdentity   types.Bool                                                                 `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed_optional"`
	EnableBindingCookie      types.Bool                                                                 `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed_optional"`
	HTTPOnlyCookieAttribute  types.Bool                                                                 `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute,computed_optional"`
	PathCookieAttribute      types.Bool                                                                 `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute,computed_optional"`
	SessionDuration          types.String                                                               `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	SkipAppLauncherLoginPage types.Bool                                                                 `tfsdk:"skip_app_launcher_login_page" json:"skip_app_launcher_login_page,computed_optional"`
	CORSHeaders              customfield.NestedObject[ZeroTrustAccessApplicationCORSHeadersModel]       `tfsdk:"cors_headers" json:"cors_headers,computed_optional"`
	FooterLinks              customfield.NestedObjectList[ZeroTrustAccessApplicationFooterLinksModel]   `tfsdk:"footer_links" json:"footer_links,computed_optional"`
	LandingPageDesign        customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignModel] `tfsdk:"landing_page_design" json:"landing_page_design,computed_optional"`
	Policies                 customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesModel]      `tfsdk:"policies" json:"policies,computed_optional"`
	SaaSApp                  customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppModel]           `tfsdk:"saas_app" json:"saas_app,computed_optional"`
	SCIMConfig               customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigModel]        `tfsdk:"scim_config" json:"scim_config,computed_optional"`
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

type ZeroTrustAccessApplicationFooterLinksModel struct {
	Name types.String `tfsdk:"name" json:"name,required"`
	URL  types.String `tfsdk:"url" json:"url,required"`
}

type ZeroTrustAccessApplicationLandingPageDesignModel struct {
	ButtonColor     types.String `tfsdk:"button_color" json:"button_color,optional"`
	ButtonTextColor types.String `tfsdk:"button_text_color" json:"button_text_color,optional"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,optional"`
	Message         types.String `tfsdk:"message" json:"message,optional"`
	Title           types.String `tfsdk:"title" json:"title,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesModel struct {
	ID                           types.String                                                                        `tfsdk:"id" json:"id,optional"`
	Precedence                   types.Int64                                                                         `tfsdk:"precedence" json:"precedence,optional"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesApprovalGroupsModel] `tfsdk:"approval_groups" json:"approval_groups,computed_optional"`
	ApprovalRequired             types.Bool                                                                          `tfsdk:"approval_required" json:"approval_required,computed_optional"`
	ConnectionRules              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesConnectionRulesModel]    `tfsdk:"connection_rules" json:"connection_rules,computed_optional"`
	CreatedAt                    timetypes.RFC3339                                                                   `tfsdk:"created_at" json:"created_at,optional" format:"date-time"`
	Decision                     types.String                                                                        `tfsdk:"decision" json:"decision,optional"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesExcludeModel]        `tfsdk:"exclude" json:"exclude,computed_optional"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesIncludeModel]        `tfsdk:"include" json:"include,computed_optional"`
	IsolationRequired            types.Bool                                                                          `tfsdk:"isolation_required" json:"isolation_required,computed_optional"`
	Name                         types.String                                                                        `tfsdk:"name" json:"name,optional"`
	PurposeJustificationPrompt   types.String                                                                        `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,optional"`
	PurposeJustificationRequired types.Bool                                                                          `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed_optional"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesRequireModel]        `tfsdk:"require" json:"require,computed_optional"`
	SessionDuration              types.String                                                                        `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	UpdatedAt                    timetypes.RFC3339                                                                   `tfsdk:"updated_at" json:"updated_at,optional" format:"date-time"`
}

type ZeroTrustAccessApplicationPoliciesApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,required"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses,optional"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid,optional"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesModel struct {
	SSH customfield.NestedObject[ZeroTrustAccessApplicationPoliciesConnectionRulesSSHModel] `tfsdk:"ssh" json:"ssh,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesSSHModel struct {
	Usernames *[]types.String `tfsdk:"usernames" json:"usernames,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesExcludeEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesExcludeCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
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

type ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesIncludeEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesIncludeCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
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

type ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             *ZeroTrustAccessApplicationPoliciesRequireEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          *ZeroTrustAccessApplicationPoliciesRequireCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireEveryoneModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireCertificateModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
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

type ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessApplicationPoliciesRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessApplicationSaaSAppModel struct {
	AuthType                      types.String                                                                             `tfsdk:"auth_type" json:"auth_type,optional"`
	ConsumerServiceURL            types.String                                                                             `tfsdk:"consumer_service_url" json:"consumer_service_url,optional"`
	CreatedAt                     timetypes.RFC3339                                                                        `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomAttributes              customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomAttributesModel]         `tfsdk:"custom_attributes" json:"custom_attributes,computed_optional"`
	DefaultRelayState             types.String                                                                             `tfsdk:"default_relay_state" json:"default_relay_state,optional"`
	IdPEntityID                   types.String                                                                             `tfsdk:"idp_entity_id" json:"idp_entity_id,optional"`
	NameIDFormat                  types.String                                                                             `tfsdk:"name_id_format" json:"name_id_format,optional"`
	NameIDTransformJsonata        types.String                                                                             `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata,optional"`
	PublicKey                     types.String                                                                             `tfsdk:"public_key" json:"public_key,optional"`
	SAMLAttributeTransformJsonata types.String                                                                             `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata,optional"`
	SPEntityID                    types.String                                                                             `tfsdk:"sp_entity_id" json:"sp_entity_id,optional"`
	SSOEndpoint                   types.String                                                                             `tfsdk:"sso_endpoint" json:"sso_endpoint,optional"`
	UpdatedAt                     timetypes.RFC3339                                                                        `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AccessTokenLifetime           types.String                                                                             `tfsdk:"access_token_lifetime" json:"access_token_lifetime,optional"`
	AllowPKCEWithoutClientSecret  types.Bool                                                                               `tfsdk:"allow_pkce_without_client_secret" json:"allow_pkce_without_client_secret,optional"`
	AppLauncherURL                types.String                                                                             `tfsdk:"app_launcher_url" json:"app_launcher_url,optional"`
	ClientID                      types.String                                                                             `tfsdk:"client_id" json:"client_id,optional"`
	ClientSecret                  types.String                                                                             `tfsdk:"client_secret" json:"client_secret,optional"`
	CustomClaims                  customfield.NestedObjectList[ZeroTrustAccessApplicationSaaSAppCustomClaimsModel]         `tfsdk:"custom_claims" json:"custom_claims,computed_optional"`
	GrantTypes                    *[]types.String                                                                          `tfsdk:"grant_types" json:"grant_types,optional"`
	GroupFilterRegex              types.String                                                                             `tfsdk:"group_filter_regex" json:"group_filter_regex,optional"`
	HybridAndImplicitOptions      customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel] `tfsdk:"hybrid_and_implicit_options" json:"hybrid_and_implicit_options,computed_optional"`
	RedirectURIs                  *[]types.String                                                                          `tfsdk:"redirect_uris" json:"redirect_uris,optional"`
	RefreshTokenOptions           customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel]      `tfsdk:"refresh_token_options" json:"refresh_token_options,computed_optional"`
	Scopes                        *[]types.String                                                                          `tfsdk:"scopes" json:"scopes,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesModel struct {
	FriendlyName types.String                                                                           `tfsdk:"friendly_name" json:"friendly_name,optional"`
	Name         types.String                                                                           `tfsdk:"name" json:"name,optional"`
	NameFormat   types.String                                                                           `tfsdk:"name_format" json:"name_format,optional"`
	Required     types.Bool                                                                             `tfsdk:"required" json:"required,optional"`
	Source       customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel] `tfsdk:"source" json:"source,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceModel struct {
	Name      types.String            `tfsdk:"name" json:"name,optional"`
	NameByIdP map[string]types.String `tfsdk:"name_by_idp" json:"name_by_idp,optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsModel struct {
	Name     types.String                                                                       `tfsdk:"name" json:"name,optional"`
	Required types.Bool                                                                         `tfsdk:"required" json:"required,optional"`
	Scope    types.String                                                                       `tfsdk:"scope" json:"scope,optional"`
	Source   customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel] `tfsdk:"source" json:"source,computed_optional"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceModel struct {
	Name      types.String            `tfsdk:"name" json:"name,optional"`
	NameByIdP map[string]types.String `tfsdk:"name_by_idp" json:"name_by_idp,optional"`
}

type ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint" json:"return_access_token_from_authorization_endpoint,optional"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint" json:"return_id_token_from_authorization_endpoint,optional"`
}

type ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsModel struct {
	Lifetime types.String `tfsdk:"lifetime" json:"lifetime,optional"`
}

type ZeroTrustAccessApplicationSCIMConfigModel struct {
	IdPUID             types.String                                                                      `tfsdk:"idp_uid" json:"idp_uid,required"`
	RemoteURI          types.String                                                                      `tfsdk:"remote_uri" json:"remote_uri,required"`
	Authentication     customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigAuthenticationModel] `tfsdk:"authentication" json:"authentication,computed_optional"`
	DeactivateOnDelete types.Bool                                                                        `tfsdk:"deactivate_on_delete" json:"deactivate_on_delete,optional"`
	Enabled            types.Bool                                                                        `tfsdk:"enabled" json:"enabled,optional"`
	Mappings           customfield.NestedObjectList[ZeroTrustAccessApplicationSCIMConfigMappingsModel]   `tfsdk:"mappings" json:"mappings,computed_optional"`
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
	Schema           types.String                                                                          `tfsdk:"schema" json:"schema,required"`
	Enabled          types.Bool                                                                            `tfsdk:"enabled" json:"enabled,optional"`
	Filter           types.String                                                                          `tfsdk:"filter" json:"filter,optional"`
	Operations       customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel] `tfsdk:"operations" json:"operations,computed_optional"`
	TransformJsonata types.String                                                                          `tfsdk:"transform_jsonata" json:"transform_jsonata,optional"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsOperationsModel struct {
	Create types.Bool `tfsdk:"create" json:"create,optional"`
	Delete types.Bool `tfsdk:"delete" json:"delete,optional"`
	Update types.Bool `tfsdk:"update" json:"update,optional"`
}
