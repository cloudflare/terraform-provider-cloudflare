// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessGroupsResultListDataSourceEnvelope struct {
	Result *[]*AccessGroupsItemsDataSourceModel `json:"result,computed"`
}

type AccessGroupsDataSourceModel struct {
	AccountID types.String                         `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                         `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                          `tfsdk:"max_items"`
	Items     *[]*AccessGroupsItemsDataSourceModel `tfsdk:"items"`
}

type AccessGroupsItemsDataSourceModel struct {
	ID        types.String                                  `tfsdk:"id" json:"id"`
	CreatedAt timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at,computed"`
	Exclude   *[]*AccessGroupsItemsExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude"`
	Include   *[]*AccessGroupsItemsIncludeDataSourceModel   `tfsdk:"include" json:"include"`
	IsDefault *[]*AccessGroupsItemsIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default"`
	Name      types.String                                  `tfsdk:"name" json:"name"`
	Require   *[]*AccessGroupsItemsRequireDataSourceModel   `tfsdk:"require" json:"require"`
	UpdatedAt timetypes.RFC3339                             `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessGroupsItemsExcludeDataSourceModel struct {
	Email                *AccessGroupsItemsExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsItemsExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsItemsExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsItemsExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsItemsExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsItemsExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsItemsExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsItemsExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsItemsExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsItemsExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsItemsExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsItemsExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsItemsExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsItemsExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsItemsExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsItemsExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsItemsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsItemsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsItemsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsItemsExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsItemsExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsItemsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsItemsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsItemsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsItemsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsItemsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsItemsIncludeDataSourceModel struct {
	Email                *AccessGroupsItemsIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsItemsIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsItemsIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsItemsIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsItemsIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsItemsIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsItemsIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsItemsIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsItemsIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsItemsIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsItemsIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsItemsIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsItemsIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsItemsIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsItemsIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsItemsIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsItemsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsItemsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsItemsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsItemsIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsItemsIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsItemsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsItemsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsItemsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsItemsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsItemsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsItemsIsDefaultDataSourceModel struct {
	Email                *AccessGroupsItemsIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsItemsIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsItemsIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsItemsIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsItemsIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsItemsIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsItemsIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsItemsIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsItemsIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsItemsIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsItemsIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsItemsIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsItemsIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsItemsIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsItemsIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsItemsIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsItemsIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsItemsIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsItemsIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsItemsIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsItemsIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsItemsIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsItemsIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsItemsIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsItemsIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsItemsIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessGroupsItemsRequireDataSourceModel struct {
	Email                *AccessGroupsItemsRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupsItemsRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupsItemsRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupsItemsRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupsItemsRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupsItemsRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupsItemsRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupsItemsRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupsItemsRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupsItemsRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupsItemsRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupsItemsRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupsItemsRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupsItemsRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupsItemsRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupsItemsRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupsItemsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessGroupsItemsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessGroupsItemsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessGroupsItemsRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessGroupsItemsRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessGroupsItemsRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessGroupsItemsRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessGroupsItemsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessGroupsItemsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessGroupsItemsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessGroupsItemsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessGroupsItemsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
