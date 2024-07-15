// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessGroupResultEnvelope struct {
	Result AccessGroupModel `json:"result,computed"`
}

type AccessGroupModel struct {
	ID        types.String                `tfsdk:"id" json:"id,computed"`
	AccountID types.String                `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                `tfsdk:"zone_id" path:"zone_id"`
	Include   *[]*AccessGroupIncludeModel `tfsdk:"include" json:"include"`
	Name      types.String                `tfsdk:"name" json:"name"`
	Exclude   *[]*AccessGroupExcludeModel `tfsdk:"exclude" json:"exclude"`
	IsDefault types.Bool                  `tfsdk:"is_default" json:"is_default"`
	Require   *[]*AccessGroupRequireModel `tfsdk:"require" json:"require"`
	CreatedAt timetypes.RFC3339           `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339           `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessGroupIncludeModel struct {
	Email                *AccessGroupIncludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupIncludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupIncludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupIncludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupIncludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupIncludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupIncludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupIncludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupIncludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupIncludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupIncludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupIncludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupIncludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupIncludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupIncludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupIncludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupIncludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupIncludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupIncludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupIncludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupExcludeModel struct {
	Email                *AccessGroupExcludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupExcludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupExcludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupExcludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupExcludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupExcludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupExcludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupExcludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupExcludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupExcludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupExcludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupExcludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupExcludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupExcludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupExcludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupExcludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupExcludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupExcludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupExcludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupExcludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessGroupRequireModel struct {
	Email                *AccessGroupRequireEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessGroupRequireEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessGroupRequireEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                       `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessGroupRequireIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessGroupRequireIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                       `tfsdk:"certificate" json:"certificate"`
	Group                *AccessGroupRequireGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessGroupRequireAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessGroupRequireGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessGroupRequireGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessGroupRequireOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessGroupRequireSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessGroupRequireServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                       `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessGroupRequireExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessGroupRequireGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessGroupRequireAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessGroupRequireDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessGroupRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessGroupRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessGroupRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessGroupRequireAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessGroupRequireGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessGroupRequireGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessGroupRequireSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessGroupRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessGroupRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessGroupRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessGroupRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessGroupRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}
