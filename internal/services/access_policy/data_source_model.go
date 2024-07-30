// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyResultDataSourceEnvelope struct {
	Result AccessPolicyDataSourceModel `json:"result,computed"`
}

type AccessPolicyResultListDataSourceEnvelope struct {
	Result *[]*AccessPolicyDataSourceModel `json:"result,computed"`
}

type AccessPolicyDataSourceModel struct {
	AppID                        types.String                                  `tfsdk:"app_id" path:"app_id"`
	PolicyID                     types.String                                  `tfsdk:"policy_id" path:"policy_id"`
	AccountID                    types.String                                  `tfsdk:"account_id" path:"account_id"`
	ZoneID                       types.String                                  `tfsdk:"zone_id" path:"zone_id"`
	ID                           types.String                                  `tfsdk:"id" json:"id"`
	ApprovalGroups               *[]*AccessPolicyApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups"`
	ApprovalRequired             types.Bool                                    `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                             `tfsdk:"created_at" json:"created_at,computed"`
	Decision                     types.String                                  `tfsdk:"decision" json:"decision"`
	Exclude                      *[]*AccessPolicyExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*AccessPolicyIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	IsolationRequired            types.Bool                                    `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                  `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                  `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                                    `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      *[]*AccessPolicyRequireDataSourceModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                                  `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                             `tfsdk:"updated_at" json:"updated_at,computed"`
	Filter                       *AccessPolicyFindOneByDataSourceModel         `tfsdk:"filter"`
}

type AccessPolicyApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type AccessPolicyExcludeDataSourceModel struct {
	Email                *AccessPolicyExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPolicyExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPolicyExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPolicyExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPolicyExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPolicyExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPolicyExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPolicyExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPolicyExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPolicyExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPolicyIncludeDataSourceModel struct {
	Email                *AccessPolicyIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPolicyIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPolicyIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPolicyIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPolicyIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPolicyIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPolicyIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPolicyIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPolicyIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPolicyIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPolicyRequireDataSourceModel struct {
	Email                *AccessPolicyRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                  `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                  `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                  `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPolicyRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPolicyRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPolicyRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPolicyRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPolicyRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPolicyRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPolicyRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPolicyRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPolicyRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPolicyRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPolicyRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPolicyFindOneByDataSourceModel struct {
	AppID types.String `tfsdk:"app_id" path:"app_id"`
}
