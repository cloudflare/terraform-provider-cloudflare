// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
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
	ApprovalRequired             types.Bool                                    `tfsdk:"approval_required" json:"approval_required"`
	CreatedAt                    types.String                                  `tfsdk:"created_at" json:"created_at"`
	Decision                     types.String                                  `tfsdk:"decision" json:"decision"`
	Exclude                      *[]*AccessPolicyExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*AccessPolicyIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	IsolationRequired            types.Bool                                    `tfsdk:"isolation_required" json:"isolation_required"`
	Name                         types.String                                  `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                  `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                                    `tfsdk:"purpose_justification_required" json:"purpose_justification_required"`
	Require                      *[]*AccessPolicyRequireDataSourceModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                                  `tfsdk:"session_duration" json:"session_duration"`
	UpdatedAt                    types.String                                  `tfsdk:"updated_at" json:"updated_at"`
	FindOneBy                    *AccessPolicyFindOneByDataSourceModel         `tfsdk:"find_one_by"`
}

type AccessPolicyApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64 `tfsdk:"approvals_needed" json:"approvals_needed"`
	EmailAddresses  types.String  `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String  `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type AccessPolicyExcludeDataSourceModel struct {
	Email                *AccessPolicyExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                          `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessPolicyIncludeDataSourceModel struct {
	Email                *AccessPolicyIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                          `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessPolicyRequireDataSourceModel struct {
	Email                *AccessPolicyRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             types.String                                          `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          types.String                                          `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken types.String                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessPolicyFindOneByDataSourceModel struct {
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
