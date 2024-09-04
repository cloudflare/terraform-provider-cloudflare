// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultEnvelope struct {
	Result ZeroTrustAccessGroupModel `json:"result"`
}

type ZeroTrustAccessGroupModel struct {
	ID        types.String                                                   `tfsdk:"id" json:"id,computed"`
	AccountID types.String                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	Name      types.String                                                   `tfsdk:"name" json:"name,required"`
	Include   *[]*ZeroTrustAccessGroupIncludeModel                           `tfsdk:"include" json:"include,required"`
	IsDefault types.Bool                                                     `tfsdk:"is_default" json:"is_default,optional"`
	Exclude   customfield.NestedObjectList[ZeroTrustAccessGroupExcludeModel] `tfsdk:"exclude" json:"exclude,computed_optional"`
	Require   customfield.NestedObjectList[ZeroTrustAccessGroupRequireModel] `tfsdk:"require" json:"require,computed_optional"`
	CreatedAt timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessGroupIncludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupIncludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                         `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupIncludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupIncludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                         `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupIncludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupIncludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupIncludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupIncludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupIncludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupIncludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupIncludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupIncludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupIncludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupIncludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupIncludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id,required"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Name         types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupIncludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupIncludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupIncludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
}

type ZeroTrustAccessGroupIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessGroupIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessGroupExcludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                         `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupExcludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupExcludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                         `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupExcludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupExcludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupExcludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupExcludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupExcludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupExcludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupExcludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupExcludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupExcludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupExcludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupExcludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id,required"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Name         types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupExcludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupExcludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupExcludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
}

type ZeroTrustAccessGroupExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessGroupExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessGroupRequireModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessGroupRequireEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupRequireEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupRequireEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                         `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupRequireIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupRequireIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                         `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessGroupRequireGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupRequireAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupRequireGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupRequireGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupRequireOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupRequireSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupRequireServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupRequireExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupRequireGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupRequireAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupRequireDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id,required"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Name         types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupRequireGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupRequireOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,required"`
	Email        types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupRequireSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
}

type ZeroTrustAccessGroupRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessGroupRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}
