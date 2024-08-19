// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultDataSourceEnvelope struct {
	Result ZeroTrustAccessGroupDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessGroupResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessGroupDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessGroupDataSourceModel struct {
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	GroupID   types.String                                     `tfsdk:"group_id" path:"group_id"`
	ZoneID    types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339                                `tfsdk:"updated_at" json:"updated_at,computed"`
	ID        types.String                                     `tfsdk:"id" json:"id"`
	Name      types.String                                     `tfsdk:"name" json:"name"`
	Exclude   *[]*ZeroTrustAccessGroupExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude"`
	Include   *[]*ZeroTrustAccessGroupIncludeDataSourceModel   `tfsdk:"include" json:"include"`
	IsDefault *[]*ZeroTrustAccessGroupIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default"`
	Require   *[]*ZeroTrustAccessGroupRequireDataSourceModel   `tfsdk:"require" json:"require"`
	Filter    *ZeroTrustAccessGroupFindOneByDataSourceModel    `tfsdk:"filter"`
}

type ZeroTrustAccessGroupExcludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIncludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIsDefaultDataSourceModel struct {
	Email                *ZeroTrustAccessGroupIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                            `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                            `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                            `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupRequireDataSourceModel struct {
	Email                *ZeroTrustAccessGroupRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
