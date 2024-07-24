// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPoliciesResultListDataSourceEnvelope struct {
	Result *[]*AccessPoliciesItemsDataSourceModel `json:"result,computed"`
}

type AccessPoliciesDataSourceModel struct {
	AppID     types.String                           `tfsdk:"app_id" path:"app_id"`
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                           `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Items     *[]*AccessPoliciesItemsDataSourceModel `tfsdk:"items"`
}

type AccessPoliciesItemsDataSourceModel struct {
	ID                           types.String                                         `tfsdk:"id" json:"id"`
	ApprovalGroups               *[]*AccessPoliciesItemsApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups"`
	ApprovalRequired             types.Bool                                           `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed"`
	Decision                     types.String                                         `tfsdk:"decision" json:"decision"`
	Exclude                      *[]*AccessPoliciesItemsExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*AccessPoliciesItemsIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	IsolationRequired            types.Bool                                           `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                         `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                         `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                                           `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      *[]*AccessPoliciesItemsRequireDataSourceModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                                         `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                    `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessPoliciesItemsApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type AccessPoliciesItemsExcludeDataSourceModel struct {
	Email                *AccessPoliciesItemsExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesItemsExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesItemsExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesItemsExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesItemsExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesItemsExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesItemsExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesItemsExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesItemsExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesItemsExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesItemsExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesItemsExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesItemsExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesItemsExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesItemsExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesItemsExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesItemsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesItemsIncludeDataSourceModel struct {
	Email                *AccessPoliciesItemsIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesItemsIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesItemsIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesItemsIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesItemsIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesItemsIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesItemsIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesItemsIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesItemsIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesItemsIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesItemsIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesItemsIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesItemsIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesItemsIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesItemsIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesItemsIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesItemsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesItemsRequireDataSourceModel struct {
	Email                *AccessPoliciesItemsRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPoliciesItemsRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPoliciesItemsRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                         `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPoliciesItemsRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPoliciesItemsRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                         `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPoliciesItemsRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPoliciesItemsRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPoliciesItemsRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPoliciesItemsRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPoliciesItemsRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPoliciesItemsRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPoliciesItemsRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                         `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPoliciesItemsRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPoliciesItemsRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPoliciesItemsRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPoliciesItemsRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPoliciesItemsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
