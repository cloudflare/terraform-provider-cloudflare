// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultEnvelope struct {
	Result ZeroTrustAccessGroupModel `json:"result,computed"`
}

type ZeroTrustAccessGroupModel struct {
	ID        types.String                         `tfsdk:"id" json:"id,computed"`
	AccountID types.String                         `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                         `tfsdk:"zone_id" path:"zone_id"`
	Name      types.String                         `tfsdk:"name" json:"name"`
	Include   *[]*ZeroTrustAccessGroupIncludeModel `tfsdk:"include" json:"include"`
	IsDefault types.Bool                           `tfsdk:"is_default" json:"is_default"`
	Exclude   *[]*ZeroTrustAccessGroupExcludeModel `tfsdk:"exclude" json:"exclude"`
	Require   *[]*ZeroTrustAccessGroupRequireModel `tfsdk:"require" json:"require"`
	CreatedAt timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339                    `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustAccessGroupIncludeModel struct {
	Email                *ZeroTrustAccessGroupIncludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupIncludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupIncludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupIncludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupIncludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupIncludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupIncludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupIncludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupIncludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupIncludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupIncludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupIncludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupIncludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupIncludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupIncludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupIncludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessGroupIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessGroupIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupIncludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessGroupIncludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupIncludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupIncludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessGroupIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessGroupIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessGroupIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessGroupIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessGroupIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type ZeroTrustAccessGroupExcludeModel struct {
	Email                *ZeroTrustAccessGroupExcludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupExcludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupExcludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupExcludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupExcludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupExcludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupExcludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupExcludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupExcludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupExcludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupExcludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupExcludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupExcludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupExcludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupExcludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupExcludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessGroupExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessGroupExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupExcludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessGroupExcludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupExcludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupExcludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessGroupExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessGroupExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessGroupExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessGroupExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessGroupExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type ZeroTrustAccessGroupRequireModel struct {
	Email                *ZeroTrustAccessGroupRequireEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupRequireEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupRequireEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupRequireIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupRequireIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupRequireGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupRequireAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupRequireGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupRequireGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupRequireOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupRequireSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupRequireServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupRequireExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupRequireGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupRequireAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupRequireDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessGroupRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessGroupRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessGroupRequireAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessGroupRequireGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupRequireOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessGroupRequireSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessGroupRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessGroupRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessGroupRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessGroupRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessGroupRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}
