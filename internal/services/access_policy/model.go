// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPolicyResultEnvelope struct {
	Result AccessPolicyModel `json:"result,computed"`
}

type AccessPolicyModel struct {
	ID                           types.String                        `tfsdk:"id" json:"id,computed"`
	AppID                        types.String                        `tfsdk:"app_id" path:"app_id"`
	AccountID                    types.String                        `tfsdk:"account_id" path:"account_id"`
	ZoneID                       types.String                        `tfsdk:"zone_id" path:"zone_id"`
	Decision                     types.String                        `tfsdk:"decision" json:"decision"`
	Include                      *[]*AccessPolicyIncludeModel        `tfsdk:"include" json:"include"`
	Name                         types.String                        `tfsdk:"name" json:"name"`
	ApprovalGroups               *[]*AccessPolicyApprovalGroupsModel `tfsdk:"approval_groups" json:"approval_groups"`
	ApprovalRequired             types.Bool                          `tfsdk:"approval_required" json:"approval_required"`
	Exclude                      *[]*AccessPolicyExcludeModel        `tfsdk:"exclude" json:"exclude"`
	IsolationRequired            types.Bool                          `tfsdk:"isolation_required" json:"isolation_required"`
	Precedence                   types.Int64                         `tfsdk:"precedence" json:"precedence"`
	PurposeJustificationPrompt   types.String                        `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	PurposeJustificationRequired types.Bool                          `tfsdk:"purpose_justification_required" json:"purpose_justification_required"`
	Require                      *[]*AccessPolicyRequireModel        `tfsdk:"require" json:"require"`
	SessionDuration              types.String                        `tfsdk:"session_duration" json:"session_duration"`
	CreatedAt                    timetypes.RFC3339                   `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt                    timetypes.RFC3339                   `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessPolicyIncludeModel struct {
	Email                *AccessPolicyIncludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyIncludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyIncludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                        `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyIncludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyIncludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                        `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyIncludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyIncludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyIncludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyIncludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyIncludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyIncludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyIncludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyIncludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyIncludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyIncludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyIncludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyIncludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyIncludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyIncludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyIncludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessPolicyApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type AccessPolicyExcludeModel struct {
	Email                *AccessPolicyExcludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyExcludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyExcludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                        `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyExcludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyExcludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                        `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyExcludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyExcludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyExcludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyExcludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyExcludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyExcludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyExcludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyExcludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyExcludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyExcludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyExcludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyExcludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyExcludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyExcludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyExcludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type AccessPolicyRequireModel struct {
	Email                *AccessPolicyRequireEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *AccessPolicyRequireEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *AccessPolicyRequireEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                        `tfsdk:"everyone" json:"everyone"`
	IP                   *AccessPolicyRequireIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *AccessPolicyRequireIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                        `tfsdk:"certificate" json:"certificate"`
	Group                *AccessPolicyRequireGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *AccessPolicyRequireAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *AccessPolicyRequireGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *AccessPolicyRequireGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *AccessPolicyRequireOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *AccessPolicyRequireSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *AccessPolicyRequireServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *AccessPolicyRequireExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *AccessPolicyRequireGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *AccessPolicyRequireAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *AccessPolicyRequireDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type AccessPolicyRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type AccessPolicyRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type AccessPolicyRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccessPolicyRequireAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type AccessPolicyRequireGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type AccessPolicyRequireGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type AccessPolicyRequireSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type AccessPolicyRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type AccessPolicyRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type AccessPolicyRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type AccessPolicyRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type AccessPolicyRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}
