// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPolicyResultEnvelope struct {
	Result ZeroTrustAccessPolicyModel `json:"result"`
}

type ZeroTrustAccessPolicyModel struct {
	ID                           types.String                                                           `tfsdk:"id" json:"id,computed"`
	AppID                        types.String                                                           `tfsdk:"app_id" path:"app_id,required"`
	AccountID                    types.String                                                           `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID                       types.String                                                           `tfsdk:"zone_id" path:"zone_id,optional"`
	Decision                     types.String                                                           `tfsdk:"decision" json:"decision,required"`
	Name                         types.String                                                           `tfsdk:"name" json:"name,required"`
	Include                      *[]*ZeroTrustAccessPolicyIncludeModel                                  `tfsdk:"include" json:"include,required"`
	Precedence                   types.Int64                                                            `tfsdk:"precedence" json:"precedence,optional"`
	ApprovalRequired             types.Bool                                                             `tfsdk:"approval_required" json:"approval_required,computed_optional"`
	IsolationRequired            types.Bool                                                             `tfsdk:"isolation_required" json:"isolation_required,computed_optional"`
	PurposeJustificationPrompt   types.String                                                           `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed_optional"`
	PurposeJustificationRequired types.Bool                                                             `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed_optional"`
	SessionDuration              types.String                                                           `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessPolicyApprovalGroupsModel] `tfsdk:"approval_groups" json:"approval_groups,computed_optional"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPolicyExcludeModel]        `tfsdk:"exclude" json:"exclude,computed_optional"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPolicyRequireModel]        `tfsdk:"require" json:"require,computed_optional"`
	CreatedAt                    timetypes.RFC3339                                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt                    timetypes.RFC3339                                                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessPolicyIncludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyIncludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyIncludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyIncludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyIncludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyIncludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyIncludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyIncludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyIncludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyIncludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessPolicyIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyIncludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessPolicyIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessPolicyApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64                  `tfsdk:"approvals_needed" json:"approvals_needed,required"`
	EmailAddresses  customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed_optional"`
	EmailListUUID   types.String                   `tfsdk:"email_list_uuid" json:"email_list_uuid,computed_optional"`
}

type ZeroTrustAccessPolicyExcludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyExcludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyExcludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyExcludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyExcludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyExcludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyExcludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyExcludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyExcludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyExcludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessPolicyExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyExcludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessPolicyExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessPolicyRequireModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyRequireIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyRequireIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyRequireGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyRequireAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyRequireGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyRequireGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyRequireOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyRequireSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyRequireServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyRequireExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyRequireGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyRequireAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyRequireDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessPolicyRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessPolicyRequireSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessPolicyRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}