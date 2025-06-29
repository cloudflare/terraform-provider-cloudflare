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
	ID                           types.String                                                    `tfsdk:"id" json:"id,computed"`
	AccountID                    types.String                                                    `tfsdk:"account_id" path:"account_id,required"`
	Decision                     types.String                                                    `tfsdk:"decision" json:"decision,required"`
	Name                         types.String                                                    `tfsdk:"name" json:"name,required"`
	ApprovalRequired             types.Bool                                                      `tfsdk:"approval_required" json:"approval_required,optional"`
	IsolationRequired            types.Bool                                                      `tfsdk:"isolation_required" json:"isolation_required,optional"`
	PurposeJustificationPrompt   types.String                                                    `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,optional"`
	PurposeJustificationRequired types.Bool                                                      `tfsdk:"purpose_justification_required" json:"purpose_justification_required,optional"`
	ApprovalGroups               *[]*ZeroTrustAccessPolicyApprovalGroupsModel                    `tfsdk:"approval_groups" json:"approval_groups,optional"`
	SessionDuration              types.String                                                    `tfsdk:"session_duration" json:"session_duration,computed_optional"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPolicyExcludeModel] `tfsdk:"exclude" json:"exclude,computed_optional"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessPolicyIncludeModel] `tfsdk:"include" json:"include,computed_optional"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPolicyRequireModel] `tfsdk:"require" json:"require,computed_optional"`
	CreatedAt                    timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt                    timetypes.RFC3339                                               `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessPolicyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessPolicyModel) MarshalJSONForUpdate(state ZeroTrustAccessPolicyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessPolicyApprovalGroupsModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,required"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses,optional"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid,optional"`
}

type ZeroTrustAccessPolicyExcludeModel struct {
	Group                *ZeroTrustAccessPolicyExcludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessPolicyExcludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessPolicyExcludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessPolicyExcludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessPolicyExcludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessPolicyExcludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessPolicyExcludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessPolicyExcludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessPolicyExcludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessPolicyExcludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessPolicyExcludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessPolicyExcludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessPolicyExcludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessPolicyExcludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessPolicyExcludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessPolicyExcludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessPolicyExcludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessPolicyExcludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessPolicyExcludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessPolicyExcludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessPolicyExcludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessPolicyExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessPolicyExcludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeCertificateModel struct {
}

type ZeroTrustAccessPolicyExcludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessPolicyExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessPolicyExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyExcludeEveryoneModel struct {
}

type ZeroTrustAccessPolicyExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessPolicyExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyExcludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
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

type ZeroTrustAccessPolicyIncludeModel struct {
	Group                *ZeroTrustAccessPolicyIncludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessPolicyIncludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessPolicyIncludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessPolicyIncludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessPolicyIncludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessPolicyIncludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessPolicyIncludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessPolicyIncludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessPolicyIncludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessPolicyIncludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessPolicyIncludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessPolicyIncludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessPolicyIncludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessPolicyIncludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessPolicyIncludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessPolicyIncludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessPolicyIncludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessPolicyIncludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessPolicyIncludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessPolicyIncludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessPolicyIncludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessPolicyIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessPolicyIncludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeCertificateModel struct {
}

type ZeroTrustAccessPolicyIncludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessPolicyIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessPolicyIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyIncludeEveryoneModel struct {
}

type ZeroTrustAccessPolicyIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessPolicyIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyIncludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
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

type ZeroTrustAccessPolicyRequireModel struct {
	Group                *ZeroTrustAccessPolicyRequireGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessPolicyRequireAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessPolicyRequireAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessPolicyRequireAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessPolicyRequireAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessPolicyRequireCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessPolicyRequireCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessPolicyRequireGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessPolicyRequireDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessPolicyRequireEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessPolicyRequireEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessPolicyRequireEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessPolicyRequireEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessPolicyRequireExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessPolicyRequireGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessPolicyRequireGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessPolicyRequireLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessPolicyRequireIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessPolicyRequireIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessPolicyRequireOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessPolicyRequireSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessPolicyRequireServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessPolicyRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessPolicyRequireAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessPolicyRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireCertificateModel struct {
}

type ZeroTrustAccessPolicyRequireCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessPolicyRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessPolicyRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessPolicyRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessPolicyRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessPolicyRequireEveryoneModel struct {
}

type ZeroTrustAccessPolicyRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessPolicyRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessPolicyRequireLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessPolicyRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
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
