// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPoliciesResultListDataSourceEnvelope struct {
	Result *[]*AccessPoliciesResultDataSourceModel `json:"result,computed"`
}

type AccessPoliciesDataSourceModel struct {
	AppID     types.String                            `tfsdk:"app_id" path:"app_id"`
	AccountID types.String                            `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                            `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Result    *[]*AccessPoliciesResultDataSourceModel `tfsdk:"result"`
}

type AccessPoliciesResultDataSourceModel struct {
	ID                           types.String                                    `tfsdk:"id" json:"id"`
	ApprovalGroups               *[]*AccessPoliciesApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups"`
	ApprovalRequired             types.Bool                                      `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                               `tfsdk:"created_at" json:"created_at,computed"`
	Decision                     types.String                                    `tfsdk:"decision" json:"decision"`
	Exclude                      *[]*AccessPoliciesExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*AccessPoliciesIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	IsolationRequired            types.Bool                                      `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                    `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                    `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                                      `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      *[]*AccessPoliciesRequireDataSourceModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                                    `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                               `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessPoliciesApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type AccessPoliciesExcludeDataSourceModel struct {
	Email                *AccessPoliciesExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                    `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                    `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesIncludeDataSourceModel struct {
	Email                *AccessPoliciesIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                    `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                    `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesRequireDataSourceModel struct {
	Email                *AccessPoliciesRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                    `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                    `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
