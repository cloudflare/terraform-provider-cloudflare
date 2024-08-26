// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPolicyResultEnvelope struct {
	Result ZeroTrustAccessPolicyModel `json:"result"`
}

type ZeroTrustAccessPolicyModel struct {
	ID                           types.String                                 `tfsdk:"id" json:"id,computed"`
	AppID                        types.String                                 `tfsdk:"app_id" path:"app_id"`
	AccountID                    types.String                                 `tfsdk:"account_id" path:"account_id"`
	ZoneID                       types.String                                 `tfsdk:"zone_id" path:"zone_id"`
	Decision                     types.String                                 `tfsdk:"decision" json:"decision"`
	Name                         types.String                                 `tfsdk:"name" json:"name"`
	Include                      *[]*ZeroTrustAccessPolicyIncludeModel        `tfsdk:"include" json:"include"`
	Precedence                   types.Int64                                  `tfsdk:"precedence" json:"precedence"`
	PurposeJustificationPrompt   types.String                                 `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	ApprovalGroups               *[]*ZeroTrustAccessPolicyApprovalGroupsModel `tfsdk:"approval_groups" json:"approval_groups"`
	Exclude                      *[]*ZeroTrustAccessPolicyExcludeModel        `tfsdk:"exclude" json:"exclude"`
	Require                      *[]*ZeroTrustAccessPolicyRequireModel        `tfsdk:"require" json:"require"`
	ApprovalRequired             types.Bool                                   `tfsdk:"approval_required" json:"approval_required"`
	IsolationRequired            types.Bool                                   `tfsdk:"isolation_required" json:"isolation_required"`
	PurposeJustificationRequired types.Bool                                   `tfsdk:"purpose_justification_required" json:"purpose_justification_required"`
	SessionDuration              types.String                                 `tfsdk:"session_duration" json:"session_duration"`
	CreatedAt                    timetypes.RFC3339                            `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt                    timetypes.RFC3339                            `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustAccessPolicyIncludeModel struct {
	Email                *ZeroTrustAccessPolicyIncludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyIncludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyIncludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                 `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyIncludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyIncludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                 `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyIncludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyIncludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyIncludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyIncludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyIncludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyIncludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyIncludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                 `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyIncludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyIncludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyIncludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyIncludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPolicyIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessPolicyIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessPolicyIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyIncludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessPolicyIncludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyIncludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyIncludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessPolicyIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessPolicyIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessPolicyIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessPolicyIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessPolicyIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type ZeroTrustAccessPolicyApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type ZeroTrustAccessPolicyExcludeModel struct {
	Email                *ZeroTrustAccessPolicyExcludeEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyExcludeEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyExcludeEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                 `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyExcludeIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyExcludeIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                 `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyExcludeGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyExcludeAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyExcludeGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyExcludeGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyExcludeOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyExcludeSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyExcludeServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                 `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyExcludeExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyExcludeGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyExcludeAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyExcludeDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPolicyExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessPolicyExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessPolicyExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyExcludeAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessPolicyExcludeGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyExcludeOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyExcludeSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessPolicyExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessPolicyExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessPolicyExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessPolicyExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessPolicyExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}

type ZeroTrustAccessPolicyRequireModel struct {
	Email                *ZeroTrustAccessPolicyRequireEmailModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyRequireEmailListModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyRequireEmailDomainModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                 `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyRequireIPModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyRequireIPListModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                 `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyRequireGroupModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyRequireAzureADModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyRequireGitHubOrganizationModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyRequireGSuiteModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyRequireOktaModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyRequireSAMLModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyRequireServiceTokenModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                 `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyRequireExternalEvaluationModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyRequireGeoModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyRequireAuthMethodModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyRequireDevicePostureModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessPolicyRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain"`
}

type ZeroTrustAccessPolicyRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip"`
}

type ZeroTrustAccessPolicyRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type ZeroTrustAccessPolicyRequireAzureADModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Name         types.String `tfsdk:"name" json:"name"`
}

type ZeroTrustAccessPolicyRequireGSuiteModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyRequireOktaModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id"`
	Email        types.String `tfsdk:"email" json:"email"`
}

type ZeroTrustAccessPolicyRequireSAMLModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value"`
}

type ZeroTrustAccessPolicyRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id"`
}

type ZeroTrustAccessPolicyRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url"`
}

type ZeroTrustAccessPolicyRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code"`
}

type ZeroTrustAccessPolicyRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method"`
}

type ZeroTrustAccessPolicyRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid"`
}
