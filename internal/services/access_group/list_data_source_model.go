// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessGroupsResultListDataSourceEnvelope struct {
	Result *[]*AccessGroupsResultDataSourceModel `json:"result,computed"`
}

type AccessGroupsDataSourceModel struct {
	AccountID types.String                          `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                          `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                           `tfsdk:"max_items"`
	Result    *[]*AccessGroupsResultDataSourceModel `tfsdk:"result"`
}

type AccessGroupsResultDataSourceModel struct {
	ID        types.String                             `tfsdk:"id" json:"id"`
	CreatedAt timetypes.RFC3339                        `tfsdk:"created_at" json:"created_at,computed"`
	Exclude   *[]*AccessGroupsExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude"`
	Include   *[]*AccessGroupsIncludeDataSourceModel   `tfsdk:"include" json:"include"`
	IsDefault *[]*AccessGroupsIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default"`
	Name      types.String                             `tfsdk:"name" json:"name"`
	Require   *[]*AccessGroupsRequireDataSourceModel   `tfsdk:"require" json:"require"`
	UpdatedAt timetypes.RFC3339                        `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessGroupsExcludeDataSourceModel struct {
	Email                *AccessGroupsExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsIncludeDataSourceModel struct {
	Email                *AccessGroupsIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsIsDefaultDataSourceModel struct {
	Email                *AccessGroupsIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                    `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                    `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsRequireDataSourceModel struct {
	Email                *AccessGroupsRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
