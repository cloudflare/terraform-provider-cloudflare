// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultDataSourceEnvelope struct {
	Result ZeroTrustAccessGroupDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessGroupDataSourceModel struct {
	ID        types.String                                                               `tfsdk:"id" json:"-,computed"`
	GroupID   types.String                                                               `tfsdk:"group_id" path:"group_id,optional"`
	AccountID types.String                                                               `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                               `tfsdk:"zone_id" path:"zone_id,optional"`
	CreatedAt timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name      types.String                                                               `tfsdk:"name" json:"name,computed"`
	UpdatedAt timetypes.RFC3339                                                          `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Exclude   customfield.NestedObjectList[ZeroTrustAccessGroupExcludeDataSourceModel]   `tfsdk:"exclude" json:"exclude,computed"`
	Include   customfield.NestedObjectList[ZeroTrustAccessGroupIncludeDataSourceModel]   `tfsdk:"include" json:"include,computed"`
	IsDefault customfield.NestedObjectList[ZeroTrustAccessGroupIsDefaultDataSourceModel] `tfsdk:"is_default" json:"is_default,computed"`
	Require   customfield.NestedObjectList[ZeroTrustAccessGroupRequireDataSourceModel]   `tfsdk:"require" json:"require,computed"`
	Filter    *ZeroTrustAccessGroupFindOneByDataSourceModel                              `tfsdk:"filter"`
}

func (m *ZeroTrustAccessGroupDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessGroupGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupGetParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessGroupDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessGroupListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupListParams{}

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

type ZeroTrustAccessGroupExcludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupExcludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupExcludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupExcludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupExcludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupExcludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupExcludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupExcludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupExcludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	LoginMethod          customfield.NestedObject[ZeroTrustAccessGroupExcludeLoginMethodDataSourceModel]          `tfsdk:"login_method" json:"login_method,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupExcludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupExcludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupExcludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupExcludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessGroupExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupExcludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupExcludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessGroupExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessGroupExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupExcludeLoginMethodDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIncludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupIncludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupIncludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupIncludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupIncludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupIncludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupIncludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupIncludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupIncludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	LoginMethod          customfield.NestedObject[ZeroTrustAccessGroupIncludeLoginMethodDataSourceModel]          `tfsdk:"login_method" json:"login_method,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupIncludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupIncludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupIncludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupIncludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessGroupIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupIncludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupIncludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessGroupIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessGroupIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIncludeLoginMethodDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupIsDefaultGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupIsDefaultAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupIsDefaultAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupIsDefaultCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupIsDefaultCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupIsDefaultGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupIsDefaultEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupIsDefaultEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	LoginMethod          customfield.NestedObject[ZeroTrustAccessGroupIsDefaultLoginMethodDataSourceModel]          `tfsdk:"login_method" json:"login_method,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupIsDefaultIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupIsDefaultIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupIsDefaultOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessGroupIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupIsDefaultAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupIsDefaultCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessGroupIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultLoginMethodDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIsDefaultOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupRequireDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupRequireGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessGroupRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupRequireAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupRequireAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupRequireAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessGroupRequireCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupRequireCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupRequireGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupRequireDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupRequireEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupRequireEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupRequireEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessGroupRequireEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupRequireGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	LoginMethod          customfield.NestedObject[ZeroTrustAccessGroupRequireLoginMethodDataSourceModel]          `tfsdk:"login_method" json:"login_method,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupRequireIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupRequireIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupRequireOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupRequireSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupRequireServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessGroupRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessGroupRequireAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessGroupRequireCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessGroupRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessGroupRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupRequireLoginMethodDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessGroupRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupFindOneByDataSourceModel struct {
	Name   types.String `tfsdk:"name" query:"name,optional"`
	Search types.String `tfsdk:"search" query:"search,optional"`
}
