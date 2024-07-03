// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessGroupResultDataSourceEnvelope struct {
	Result AccessGroupDataSourceModel `json:"result,computed"`
}

type AccessGroupResultListDataSourceEnvelope struct {
	Result *[]*AccessGroupDataSourceModel `json:"result,computed"`
}

type AccessGroupDataSourceModel struct {
	GroupID   types.String                            `tfsdk:"group_id" path:"group_id"`
	AccountID types.String                            `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                            `tfsdk:"zone_id" path:"zone_id"`
	ID        types.String                            `tfsdk:"id" json:"id"`
	CreatedAt types.String                            `tfsdk:"created_at" json:"created_at"`
	Exclude   *[]*AccessGroupExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude"`
	Include   *[]*AccessGroupIncludeDataSourceModel   `tfsdk:"include" json:"include"`
	IsDefault *[]*AccessGroupIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default"`
	Name      types.String                            `tfsdk:"name" json:"name"`
	Require   *[]*AccessGroupRequireDataSourceModel   `tfsdk:"require" json:"require"`
	UpdatedAt types.String                            `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy *AccessGroupFindOneByDataSourceModel    `tfsdk:"find_one_by"`
}

type AccessGroupExcludeDataSourceModel struct {
	Email                *AccessGroupExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupIncludeDataSourceModel struct {
	Email                *AccessGroupIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupIsDefaultDataSourceModel struct {
	Email                *AccessGroupIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                           `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupRequireDataSourceModel struct {
	Email                *AccessGroupRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
