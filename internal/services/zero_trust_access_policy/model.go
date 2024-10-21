// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPolicyResultEnvelope struct {
	Result ZeroTrustAccessPolicyModel `json:"result"`
}

type ZeroTrustAccessPolicyModel struct {
	ID                           types.String                                                           `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String                                                           `tfsdk:"account_id" path:"account_id,required"`
	Decision                     types.String                                                           `tfsdk:"decision" json:"decision,required"`
	Name                         types.String                                                           `tfsdk:"name" json:"name,required"`
	Include                      *[]*ZeroTrustAccessPolicyIncludeModel                                  `tfsdk:"include" json:"include,required"`
	PurposeJustificationPrompt   types.String                                                           `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,optional"`
	ApprovalRequired             types.Bool                                                             `tfsdk:"approval_required" json:"approval_required,computed_optional"`
	IsolationRequired            types.Bool                                                             `tfsdk:"isolation_required" json:"isolation_required,computed_optional"`
	PurposeJustificationRequired types.Bool                                                             `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed_optional"`
	SessionDuration              types.String                                                           `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessPolicyApprovalGroupsModel] `tfsdk:"approval_groups" json:"approval_groups,computed_optional"`
	ConnectionRules              customfield.NestedObject[ZeroTrustAccessPolicyConnectionRulesModel]    `tfsdk:"connection_rules" json:"connection_rules,computed_optional"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPolicyExcludeModel]        `tfsdk:"exclude" json:"exclude,computed_optional"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPolicyRequireModel]        `tfsdk:"require" json:"require,computed_optional"`
	AppCount                     types.Int64                                                            `tfsdk:"app_count" json:"app_count,computed"`
	CreatedAt                    timetypes.RFC3339                                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Reusable                     types.Bool                                                             `tfsdk:"reusable" json:"reusable,computed"`
	UpdatedAt                    timetypes.RFC3339                                                      `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessPolicyModel) MarshalJSONForUpdate(state ZeroTrustAccessPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessPolicyIncludeModel struct {
	Email                *ZeroTrustAccessPolicyIncludeEmailModel                `tfsdk:"email" json:"email,optional"`
	EmailList            *ZeroTrustAccessPolicyIncludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	EmailDomain          *ZeroTrustAccessPolicyIncludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	Everyone             *ZeroTrustAccessPolicyIncludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	IP                   *ZeroTrustAccessPolicyIncludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	IPList               *ZeroTrustAccessPolicyIncludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	Certificate          *ZeroTrustAccessPolicyIncludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	Group                *ZeroTrustAccessPolicyIncludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AzureAD              *ZeroTrustAccessPolicyIncludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	GitHubOrganization   *ZeroTrustAccessPolicyIncludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessPolicyIncludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	Okta                 *ZeroTrustAccessPolicyIncludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessPolicyIncludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessPolicyIncludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	ExternalEvaluation   *ZeroTrustAccessPolicyIncludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	Geo                  *ZeroTrustAccessPolicyIncludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	AuthMethod           *ZeroTrustAccessPolicyIncludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	DevicePosture        *ZeroTrustAccessPolicyIncludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
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

type ZeroTrustAccessPolicyIncludeEveryoneModel struct {
}

type ZeroTrustAccessPolicyIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeCertificateModel struct {
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

type ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel struct {
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
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,required"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses,optional"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid,optional"`
}

type ZeroTrustAccessPolicyConnectionRulesModel struct {
	SSH customfield.NestedObject[ZeroTrustAccessPolicyConnectionRulesSSHModel] `tfsdk:"ssh" json:"ssh,computed_optional"`
}

type ZeroTrustAccessPolicyConnectionRulesSSHModel struct {
	Usernames *[]types.String `tfsdk:"usernames" json:"usernames,required"`
}

type ZeroTrustAccessPolicyExcludeModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             *ZeroTrustAccessPolicyExcludeEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          *ZeroTrustAccessPolicyExcludeCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyExcludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyExcludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyExcludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyExcludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyExcludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
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

type ZeroTrustAccessPolicyExcludeEveryoneModel struct {
}

type ZeroTrustAccessPolicyExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeCertificateModel struct {
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

type ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel struct {
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
	Everyone             *ZeroTrustAccessPolicyRequireEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyRequireIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyRequireIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          *ZeroTrustAccessPolicyRequireCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyRequireGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyRequireAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyRequireGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyRequireGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyRequireOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyRequireSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyRequireServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyRequireAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
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

type ZeroTrustAccessPolicyRequireEveryoneModel struct {
}

type ZeroTrustAccessPolicyRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessPolicyRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireCertificateModel struct {
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

type ZeroTrustAccessPolicyRequireAnyValidServiceTokenModel struct {
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
