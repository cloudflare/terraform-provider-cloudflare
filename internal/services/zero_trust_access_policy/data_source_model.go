// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPolicyResultDataSourceEnvelope struct {
	Result ZeroTrustAccessPolicyDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessPolicyResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessPolicyDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessPolicyDataSourceModel struct {
	AccountID                    types.String                                                                     `tfsdk:"account_id" path:"account_id,optional"`
	PolicyID                     types.String                                                                     `tfsdk:"policy_id" path:"policy_id,optional"`
	AppCount                     types.Int64                                                                      `tfsdk:"app_count" json:"app_count,computed"`
	ApprovalRequired             types.Bool                                                                       `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Decision                     types.String                                                                     `tfsdk:"decision" json:"decision,computed"`
	ID                           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	IsolationRequired            types.Bool                                                                       `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                                                     `tfsdk:"name" json:"name,computed"`
	PurposeJustificationPrompt   types.String                                                                     `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
	PurposeJustificationRequired types.Bool                                                                       `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Reusable                     types.Bool                                                                       `tfsdk:"reusable" json:"reusable,computed"`
	SessionDuration              types.String                                                                     `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                                                `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessPolicyApprovalGroupsDataSourceModel] `tfsdk:"approval_groups" json:"approval_groups,computed"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPolicyExcludeDataSourceModel]        `tfsdk:"exclude" json:"exclude,computed"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessPolicyIncludeDataSourceModel]        `tfsdk:"include" json:"include,computed"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPolicyRequireDataSourceModel]        `tfsdk:"require" json:"require,computed"`
	Filter                       *ZeroTrustAccessPolicyFindOneByDataSourceModel                                   `tfsdk:"filter"`
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessPolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessPolicyGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessPolicyListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type ZeroTrustAccessPolicyApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64                  `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed"`
	EmailListUUID   types.String                   `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type ZeroTrustAccessPolicyExcludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessPolicyExcludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPolicyExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessPolicyExcludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyExcludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessPolicyExcludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessPolicyExcludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyExcludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessPolicyExcludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyExcludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPolicyExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPolicyExcludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessPolicyExcludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPolicyExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPolicyIncludeDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessPolicyIncludeGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPolicyIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessPolicyIncludeAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyIncludeAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessPolicyIncludeCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessPolicyIncludeCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyIncludeGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessPolicyIncludeEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyIncludeServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPolicyIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPolicyIncludeAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessPolicyIncludeCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPolicyIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPolicyRequireDataSourceModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessPolicyRequireGroupDataSourceModel]                `tfsdk:"group" json:"group,computed"`
	AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPolicyRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessPolicyRequireAuthContextDataSourceModel]          `tfsdk:"auth_context" json:"auth_context,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel]           `tfsdk:"auth_method" json:"auth_method,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyRequireAzureADDataSourceModel]              `tfsdk:"azure_ad" json:"azureAD,computed"`
	Certificate          customfield.NestedObject[ZeroTrustAccessPolicyRequireCertificateDataSourceModel]          `tfsdk:"certificate" json:"certificate,computed"`
	CommonName           customfield.NestedObject[ZeroTrustAccessPolicyRequireCommonNameDataSourceModel]           `tfsdk:"common_name" json:"common_name,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyRequireGeoDataSourceModel]                  `tfsdk:"geo" json:"geo,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel]        `tfsdk:"device_posture" json:"device_posture,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel]          `tfsdk:"email_domain" json:"email_domain,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailListDataSourceModel]            `tfsdk:"email_list" json:"email_list,computed"`
	Email                customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailDataSourceModel]                `tfsdk:"email" json:"email,computed"`
	Everyone             customfield.NestedObject[ZeroTrustAccessPolicyRequireEveryoneDataSourceModel]             `tfsdk:"everyone" json:"everyone,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel]   `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel]   `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyRequireGSuiteDataSourceModel]               `tfsdk:"gsuite" json:"gsuite,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyRequireIPListDataSourceModel]               `tfsdk:"ip_list" json:"ip_list,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyRequireIPDataSourceModel]                   `tfsdk:"ip" json:"ip,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyRequireOktaDataSourceModel]                 `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyRequireSAMLDataSourceModel]                 `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyRequireServiceTokenDataSourceModel]         `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPolicyRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPolicyRequireAuthContextDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessPolicyRequireCommonNameDataSourceModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPolicyRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
	Team               types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPolicyRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPolicyFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
}
