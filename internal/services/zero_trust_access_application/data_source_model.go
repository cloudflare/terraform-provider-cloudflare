// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_application

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessApplicationResultDataSourceEnvelope struct {
	Result ZeroTrustAccessApplicationDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessApplicationDataSourceModel struct {
	ID                       types.String                                                                          `tfsdk:"id" json:"-,computed"`
	AppID                    types.String                                                                          `tfsdk:"app_id" path:"app_id,optional"`
	AccountID                types.String                                                                          `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                   types.String                                                                          `tfsdk:"zone_id" path:"zone_id,optional"`
	AllowAuthenticateViaWARP types.Bool                                                                            `tfsdk:"allow_authenticate_via_warp" json:"allow_authenticate_via_warp,computed"`
	AppLauncherLogoURL       types.String                                                                          `tfsdk:"app_launcher_logo_url" json:"app_launcher_logo_url,computed"`
	AppLauncherVisible       types.Bool                                                                            `tfsdk:"app_launcher_visible" json:"app_launcher_visible,computed"`
	AUD                      types.String                                                                          `tfsdk:"aud" json:"aud,computed"`
	AutoRedirectToIdentity   types.Bool                                                                            `tfsdk:"auto_redirect_to_identity" json:"auto_redirect_to_identity,computed"`
	BgColor                  types.String                                                                          `tfsdk:"bg_color" json:"bg_color,computed"`
	CreatedAt                timetypes.RFC3339                                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomDenyMessage        types.String                                                                          `tfsdk:"custom_deny_message" json:"custom_deny_message,computed"`
	CustomDenyURL            types.String                                                                          `tfsdk:"custom_deny_url" json:"custom_deny_url,computed"`
	CustomNonIdentityDenyURL types.String                                                                          `tfsdk:"custom_non_identity_deny_url" json:"custom_non_identity_deny_url,computed"`
	Domain                   types.String                                                                          `tfsdk:"domain" json:"domain,computed"`
	EnableBindingCookie      types.Bool                                                                            `tfsdk:"enable_binding_cookie" json:"enable_binding_cookie,computed"`
	HeaderBgColor            types.String                                                                          `tfsdk:"header_bg_color" json:"header_bg_color,computed"`
	HTTPOnlyCookieAttribute  types.Bool                                                                            `tfsdk:"http_only_cookie_attribute" json:"http_only_cookie_attribute,computed"`
	LogoURL                  types.String                                                                          `tfsdk:"logo_url" json:"logo_url,computed"`
	Name                     types.String                                                                          `tfsdk:"name" json:"name,computed"`
	OptionsPreflightBypass   types.Bool                                                                            `tfsdk:"options_preflight_bypass" json:"options_preflight_bypass,computed"`
	PathCookieAttribute      types.Bool                                                                            `tfsdk:"path_cookie_attribute" json:"path_cookie_attribute,computed"`
	SameSiteCookieAttribute  types.String                                                                          `tfsdk:"same_site_cookie_attribute" json:"same_site_cookie_attribute,computed"`
	ServiceAuth401Redirect   types.Bool                                                                            `tfsdk:"service_auth_401_redirect" json:"service_auth_401_redirect,computed"`
	SessionDuration          types.String                                                                          `tfsdk:"session_duration" json:"session_duration,computed"`
	SkipAppLauncherLoginPage types.Bool                                                                            `tfsdk:"skip_app_launcher_login_page" json:"skip_app_launcher_login_page,computed"`
	SkipInterstitial         types.Bool                                                                            `tfsdk:"skip_interstitial" json:"skip_interstitial,computed"`
	Type                     types.String                                                                          `tfsdk:"type" json:"type,computed"`
	UpdatedAt                timetypes.RFC3339                                                                     `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AllowedIdPs              customfield.List[types.String]                                                        `tfsdk:"allowed_idps" json:"allowed_idps,computed"`
	CustomPages              customfield.List[types.String]                                                        `tfsdk:"custom_pages" json:"custom_pages,computed"`
	SelfHostedDomains        customfield.List[types.String]                                                        `tfsdk:"self_hosted_domains" json:"self_hosted_domains,computed"`
	Tags                     customfield.List[types.String]                                                        `tfsdk:"tags" json:"tags,computed"`
	CORSHeaders              customfield.NestedObject[ZeroTrustAccessApplicationCORSHeadersDataSourceModel]        `tfsdk:"cors_headers" json:"cors_headers,computed"`
	Destinations             customfield.NestedObjectList[ZeroTrustAccessApplicationDestinationsDataSourceModel]   `tfsdk:"destinations" json:"destinations,computed"`
	FooterLinks              customfield.NestedObjectList[ZeroTrustAccessApplicationFooterLinksDataSourceModel]    `tfsdk:"footer_links" json:"footer_links,computed"`
	LandingPageDesign        customfield.NestedObject[ZeroTrustAccessApplicationLandingPageDesignDataSourceModel]  `tfsdk:"landing_page_design" json:"landing_page_design,computed"`
	Policies                 customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesDataSourceModel]       `tfsdk:"policies" json:"policies,computed"`
	SaaSApp                  customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppDataSourceModel]            `tfsdk:"saas_app" json:"saas_app,computed"`
	SCIMConfig               customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigDataSourceModel]         `tfsdk:"scim_config" json:"scim_config,computed"`
	TargetCriteria           customfield.NestedObjectList[ZeroTrustAccessApplicationTargetCriteriaDataSourceModel] `tfsdk:"target_criteria" json:"target_criteria,computed"`
	Filter                   *ZeroTrustAccessApplicationFindOneByDataSourceModel                                   `tfsdk:"filter"`
}

func (m *ZeroTrustAccessApplicationDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessApplicationGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessApplicationDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationListParams{}

	if !m.Filter.AUD.IsNull() {
		params.AUD = cloudflare.F(m.Filter.AUD.ValueString())
	}
	if !m.Filter.Domain.IsNull() {
		params.Domain = cloudflare.F(m.Filter.Domain.ValueString())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}
	if !m.Filter.Search.IsNull() {
		params.Search = cloudflare.F(m.Filter.Search.ValueString())
	}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessApplicationCORSHeadersDataSourceModel struct {
	AllowAllHeaders  types.Bool                     `tfsdk:"allow_all_headers" json:"allow_all_headers,computed"`
	AllowAllMethods  types.Bool                     `tfsdk:"allow_all_methods" json:"allow_all_methods,computed"`
	AllowAllOrigins  types.Bool                     `tfsdk:"allow_all_origins" json:"allow_all_origins,computed"`
	AllowCredentials types.Bool                     `tfsdk:"allow_credentials" json:"allow_credentials,computed"`
	AllowedHeaders   customfield.List[types.String] `tfsdk:"allowed_headers" json:"allowed_headers,computed"`
	AllowedMethods   customfield.List[types.String] `tfsdk:"allowed_methods" json:"allowed_methods,computed"`
	AllowedOrigins   customfield.List[types.String] `tfsdk:"allowed_origins" json:"allowed_origins,computed"`
	MaxAge           types.Float64                  `tfsdk:"max_age" json:"max_age,computed"`
}

type ZeroTrustAccessApplicationDestinationsDataSourceModel struct {
	Type       types.String `tfsdk:"type" json:"type,computed"`
	URI        types.String `tfsdk:"uri" json:"uri,computed"`
	CIDR       types.String `tfsdk:"cidr" json:"cidr,computed"`
	Hostname   types.String `tfsdk:"hostname" json:"hostname,computed"`
	L4Protocol types.String `tfsdk:"l4_protocol" json:"l4_protocol,computed"`
	PortRange  types.String `tfsdk:"port_range" json:"port_range,computed"`
	VnetID     types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type ZeroTrustAccessApplicationFooterLinksDataSourceModel struct {
	Name types.String `tfsdk:"name" json:"name,computed"`
	URL  types.String `tfsdk:"url" json:"url,computed"`
}

type ZeroTrustAccessApplicationLandingPageDesignDataSourceModel struct {
	ButtonColor     types.String `tfsdk:"button_color" json:"button_color,computed"`
	ButtonTextColor types.String `tfsdk:"button_text_color" json:"button_text_color,computed"`
	ImageURL        types.String `tfsdk:"image_url" json:"image_url,computed"`
	Message         types.String `tfsdk:"message" json:"message,computed"`
	Title           types.String `tfsdk:"title" json:"title,computed"`
}

type ZeroTrustAccessApplicationPoliciesDataSourceModel struct {
	ID                           types.String                                                                                  `tfsdk:"id" json:"id,computed"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesApprovalGroupsDataSourceModel] `tfsdk:"approval_groups" json:"approval_groups,computed"`
	ApprovalRequired             types.Bool                                                                                    `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Decision                     types.String                                                                                  `tfsdk:"decision" json:"decision,computed"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesExcludeDataSourceModel]        `tfsdk:"exclude" json:"exclude,computed"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesIncludeDataSourceModel]        `tfsdk:"include" json:"include,computed"`
	IsolationRequired            types.Bool                                                                                    `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                                                                  `tfsdk:"name" json:"name,computed"`
	PurposeJustificationPrompt   types.String                                                                                  `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
	PurposeJustificationRequired types.Bool                                                                                    `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessApplicationPoliciesRequireDataSourceModel]        `tfsdk:"require" json:"require,computed"`
	SessionDuration              types.String                                                                                  `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                                                             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Precedence                   types.Int64                                                                                   `tfsdk:"precedence" json:"precedence,computed"`
	ConnectionRules              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesConnectionRulesDataSourceModel]    `tfsdk:"connection_rules" json:"connection_rules,computed"`
}

type ZeroTrustAccessApplicationPoliciesApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64                  `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed"`
	EmailListUUID   types.String                   `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesExcludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesIncludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessApplicationPoliciesRequireServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesDataSourceModel struct {
	SSH customfield.NestedObject[ZeroTrustAccessApplicationPoliciesConnectionRulesSSHDataSourceModel] `tfsdk:"ssh" json:"ssh,computed"`
}

type ZeroTrustAccessApplicationPoliciesConnectionRulesSSHDataSourceModel struct {
	Usernames       customfield.List[types.String] `tfsdk:"usernames" json:"usernames,computed"`
	AllowEmailAlias types.Bool                     `tfsdk:"allow_email_alias" json:"allow_email_alias,computed"`
}

type ZeroTrustAccessApplicationSaaSAppDataSourceModel struct {
	AuthType                      types.String                                                                                       `tfsdk:"auth_type" json:"auth_type,computed"`
	ConsumerServiceURL            types.String                                                                                       `tfsdk:"consumer_service_url" json:"consumer_service_url,computed"`
	CreatedAt                     timetypes.RFC3339                                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CustomAttributes              customfield.NestedObjectList[ZeroTrustAccessApplicationSaaSAppCustomAttributesDataSourceModel]     `tfsdk:"custom_attributes" json:"custom_attributes,computed"`
	DefaultRelayState             types.String                                                                                       `tfsdk:"default_relay_state" json:"default_relay_state,computed"`
	IdPEntityID                   types.String                                                                                       `tfsdk:"idp_entity_id" json:"idp_entity_id,computed"`
	NameIDFormat                  types.String                                                                                       `tfsdk:"name_id_format" json:"name_id_format,computed"`
	NameIDTransformJsonata        types.String                                                                                       `tfsdk:"name_id_transform_jsonata" json:"name_id_transform_jsonata,computed"`
	PublicKey                     types.String                                                                                       `tfsdk:"public_key" json:"public_key,computed"`
	SAMLAttributeTransformJsonata types.String                                                                                       `tfsdk:"saml_attribute_transform_jsonata" json:"saml_attribute_transform_jsonata,computed"`
	SPEntityID                    types.String                                                                                       `tfsdk:"sp_entity_id" json:"sp_entity_id,computed"`
	SSOEndpoint                   types.String                                                                                       `tfsdk:"sso_endpoint" json:"sso_endpoint,computed"`
	UpdatedAt                     timetypes.RFC3339                                                                                  `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	AccessTokenLifetime           types.String                                                                                       `tfsdk:"access_token_lifetime" json:"access_token_lifetime,computed"`
	AllowPKCEWithoutClientSecret  types.Bool                                                                                         `tfsdk:"allow_pkce_without_client_secret" json:"allow_pkce_without_client_secret,computed"`
	AppLauncherURL                types.String                                                                                       `tfsdk:"app_launcher_url" json:"app_launcher_url,computed"`
	ClientID                      types.String                                                                                       `tfsdk:"client_id" json:"client_id,computed"`
	ClientSecret                  types.String                                                                                       `tfsdk:"client_secret" json:"client_secret,computed"`
	CustomClaims                  customfield.NestedObjectList[ZeroTrustAccessApplicationSaaSAppCustomClaimsDataSourceModel]         `tfsdk:"custom_claims" json:"custom_claims,computed"`
	GrantTypes                    customfield.List[types.String]                                                                     `tfsdk:"grant_types" json:"grant_types,computed"`
	GroupFilterRegex              types.String                                                                                       `tfsdk:"group_filter_regex" json:"group_filter_regex,computed"`
	HybridAndImplicitOptions      customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsDataSourceModel] `tfsdk:"hybrid_and_implicit_options" json:"hybrid_and_implicit_options,computed"`
	RedirectURIs                  customfield.List[types.String]                                                                     `tfsdk:"redirect_uris" json:"redirect_uris,computed"`
	RefreshTokenOptions           customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsDataSourceModel]      `tfsdk:"refresh_token_options" json:"refresh_token_options,computed"`
	Scopes                        customfield.List[types.String]                                                                     `tfsdk:"scopes" json:"scopes,computed"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesDataSourceModel struct {
	FriendlyName types.String                                                                                     `tfsdk:"friendly_name" json:"friendly_name,computed"`
	Name         types.String                                                                                     `tfsdk:"name" json:"name,computed"`
	NameFormat   types.String                                                                                     `tfsdk:"name_format" json:"name_format,computed"`
	Required     types.Bool                                                                                       `tfsdk:"required" json:"required,computed"`
	Source       customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

type ZeroTrustAccessApplicationSaaSAppCustomAttributesSourceDataSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsDataSourceModel struct {
	Name     types.String                                                                                 `tfsdk:"name" json:"name,computed"`
	Required types.Bool                                                                                   `tfsdk:"required" json:"required,computed"`
	Scope    types.String                                                                                 `tfsdk:"scope" json:"scope,computed"`
	Source   customfield.NestedObject[ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceDataSourceModel] `tfsdk:"source" json:"source,computed"`
}

type ZeroTrustAccessApplicationSaaSAppCustomClaimsSourceDataSourceModel struct {
	Name      types.String                  `tfsdk:"name" json:"name,computed"`
	NameByIdP customfield.Map[types.String] `tfsdk:"name_by_idp" json:"name_by_idp,computed"`
}

type ZeroTrustAccessApplicationSaaSAppHybridAndImplicitOptionsDataSourceModel struct {
	ReturnAccessTokenFromAuthorizationEndpoint types.Bool `tfsdk:"return_access_token_from_authorization_endpoint" json:"return_access_token_from_authorization_endpoint,computed"`
	ReturnIDTokenFromAuthorizationEndpoint     types.Bool `tfsdk:"return_id_token_from_authorization_endpoint" json:"return_id_token_from_authorization_endpoint,computed"`
}

type ZeroTrustAccessApplicationSaaSAppRefreshTokenOptionsDataSourceModel struct {
	Lifetime types.String `tfsdk:"lifetime" json:"lifetime,computed"`
}

type ZeroTrustAccessApplicationSCIMConfigDataSourceModel struct {
	IdPUID             types.String                                                                                `tfsdk:"idp_uid" json:"idp_uid,computed"`
	RemoteURI          types.String                                                                                `tfsdk:"remote_uri" json:"remote_uri,computed"`
	Authentication     customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigAuthenticationDataSourceModel] `tfsdk:"authentication" json:"authentication,computed"`
	DeactivateOnDelete types.Bool                                                                                  `tfsdk:"deactivate_on_delete" json:"deactivate_on_delete,computed"`
	Enabled            types.Bool                                                                                  `tfsdk:"enabled" json:"enabled,computed"`
	Mappings           customfield.NestedObjectList[ZeroTrustAccessApplicationSCIMConfigMappingsDataSourceModel]   `tfsdk:"mappings" json:"mappings,computed"`
}

type ZeroTrustAccessApplicationSCIMConfigAuthenticationDataSourceModel struct {
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

type ZeroTrustAccessApplicationSCIMConfigMappingsDataSourceModel struct {
	Schema           types.String                                                                                    `tfsdk:"schema" json:"schema,computed"`
	Enabled          types.Bool                                                                                      `tfsdk:"enabled" json:"enabled,computed"`
	Filter           types.String                                                                                    `tfsdk:"filter" json:"filter,computed"`
	Operations       customfield.NestedObject[ZeroTrustAccessApplicationSCIMConfigMappingsOperationsDataSourceModel] `tfsdk:"operations" json:"operations,computed"`
	Strictness       types.String                                                                                    `tfsdk:"strictness" json:"strictness,computed"`
	TransformJsonata types.String                                                                                    `tfsdk:"transform_jsonata" json:"transform_jsonata,computed"`
}

type ZeroTrustAccessApplicationSCIMConfigMappingsOperationsDataSourceModel struct {
	Create types.Bool `tfsdk:"create" json:"create,computed"`
	Delete types.Bool `tfsdk:"delete" json:"delete,computed"`
	Update types.Bool `tfsdk:"update" json:"update,computed"`
}

type ZeroTrustAccessApplicationTargetCriteriaDataSourceModel struct {
	Port             types.Int64                                     `tfsdk:"port" json:"port,computed"`
	Protocol         types.String                                    `tfsdk:"protocol" json:"protocol,computed"`
	TargetAttributes customfield.Map[customfield.List[types.String]] `tfsdk:"target_attributes" json:"target_attributes,computed"`
}

type ZeroTrustAccessApplicationFindOneByDataSourceModel struct {
	AUD    types.String `tfsdk:"aud" query:"aud,optional"`
	Domain types.String `tfsdk:"domain" query:"domain,optional"`
	Name   types.String `tfsdk:"name" query:"name,optional"`
	Search types.String `tfsdk:"search" query:"search,optional"`
}
