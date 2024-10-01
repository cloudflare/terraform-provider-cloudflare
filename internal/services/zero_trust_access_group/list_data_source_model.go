// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessGroupsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessGroupsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                             `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessGroupsDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessGroupListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessGroupsResultDataSourceModel struct {
	ID        types.String                                                                `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339                                                           `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Exclude   customfield.NestedObjectList[ZeroTrustAccessGroupsExcludeDataSourceModel]   `tfsdk:"exclude" json:"exclude,computed"`
	Include   customfield.NestedObjectList[ZeroTrustAccessGroupsIncludeDataSourceModel]   `tfsdk:"include" json:"include,computed"`
	IsDefault customfield.NestedObjectList[ZeroTrustAccessGroupsIsDefaultDataSourceModel] `tfsdk:"is_default" json:"is_default,computed"`
	Name      types.String                                                                `tfsdk:"name" json:"name,computed"`
	Require   customfield.NestedObjectList[ZeroTrustAccessGroupsRequireDataSourceModel]   `tfsdk:"require" json:"require,computed"`
	UpdatedAt timetypes.RFC3339                                                           `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessGroupsExcludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupsExcludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupsExcludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupsExcludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupsExcludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupsExcludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupsExcludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupsExcludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupsExcludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupsExcludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupsExcludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupsExcludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupsExcludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupsExcludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupsExcludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupsExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupsExcludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupsExcludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupsExcludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupsExcludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessGroupsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsIncludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupsIncludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupsIncludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupsIncludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupsIncludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupsIncludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupsIncludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupsIncludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupsIncludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupsIncludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupsIncludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupsIncludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupsIncludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupsIncludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupsIncludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupsIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupsIncludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupsIncludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupsIncludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupsIncludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessGroupsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsIsDefaultDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupsIsDefaultDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupsIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupsIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIsDefaultSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupsIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsRequireDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupsRequireEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupsRequireEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupsRequireEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupsRequireEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupsRequireIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupsRequireIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupsRequireCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupsRequireGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupsRequireAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupsRequireGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupsRequireGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupsRequireOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupsRequireSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupsRequireServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupsRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupsRequireExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupsRequireGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupsRequireAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupsRequireDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessGroupsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
