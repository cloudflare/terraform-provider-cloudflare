// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/zero_trust"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPoliciesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[ZeroTrustAccessPoliciesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessPoliciesDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[ZeroTrustAccessPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessPoliciesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessPolicyListParams, diags diag.Diagnostics) {
  params = zero_trust.AccessPolicyListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  return
}

type ZeroTrustAccessPoliciesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AppCount types.Int64 `tfsdk:"app_count" json:"app_count,computed"`
ApprovalGroups customfield.NestedObjectList[ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel] `tfsdk:"approval_groups" json:"approval_groups,computed"`
ApprovalRequired types.Bool `tfsdk:"approval_required" json:"approval_required,computed"`
CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
Decision types.String `tfsdk:"decision" json:"decision,computed"`
Exclude customfield.NestedObjectList[ZeroTrustAccessPoliciesExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
Include customfield.NestedObjectList[ZeroTrustAccessPoliciesIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
IsolationRequired types.Bool `tfsdk:"isolation_required" json:"isolation_required,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
PurposeJustificationPrompt types.String `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
PurposeJustificationRequired types.Bool `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
Require customfield.NestedObjectList[ZeroTrustAccessPoliciesRequireDataSourceModel] `tfsdk:"require" json:"require,computed"`
Reusable types.Bool `tfsdk:"reusable" json:"reusable,computed"`
SessionDuration types.String `tfsdk:"session_duration" json:"session_duration,computed"`
UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel struct {
ApprovalsNeeded types.Float64 `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
EmailAddresses customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed"`
EmailListUUID types.String `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type ZeroTrustAccessPoliciesExcludeDataSourceModel struct {
Group customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGroupDataSourceModel] `tfsdk:"group" json:"group,computed"`
AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
AuthContext customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAuthContextDataSourceModel] `tfsdk:"auth_context" json:"auth_context,computed"`
AuthMethod customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel] `tfsdk:"auth_method" json:"auth_method,computed"`
AzureAD customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel] `tfsdk:"azure_ad" json:"azureAD,computed"`
Certificate customfield.NestedObject[ZeroTrustAccessPoliciesExcludeCertificateDataSourceModel] `tfsdk:"certificate" json:"certificate,computed"`
CommonName customfield.NestedObject[ZeroTrustAccessPoliciesExcludeCommonNameDataSourceModel] `tfsdk:"common_name" json:"common_name,computed"`
Geo customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGeoDataSourceModel] `tfsdk:"geo" json:"geo,computed"`
DevicePosture customfield.NestedObject[ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel] `tfsdk:"device_posture" json:"device_posture,computed"`
EmailDomain customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel] `tfsdk:"email_domain" json:"email_domain,computed"`
EmailList customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel] `tfsdk:"email_list" json:"email_list,computed"`
Email customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailDataSourceModel] `tfsdk:"email" json:"email,computed"`
Everyone customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEveryoneDataSourceModel] `tfsdk:"everyone" json:"everyone,computed"`
ExternalEvaluation customfield.NestedObject[ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
GitHubOrganization customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
GSuite customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel] `tfsdk:"gsuite" json:"gsuite,computed"`
LoginMethod customfield.NestedObject[ZeroTrustAccessPoliciesExcludeLoginMethodDataSourceModel] `tfsdk:"login_method" json:"login_method,computed"`
IPList customfield.NestedObject[ZeroTrustAccessPoliciesExcludeIPListDataSourceModel] `tfsdk:"ip_list" json:"ip_list,computed"`
IP customfield.NestedObject[ZeroTrustAccessPoliciesExcludeIPDataSourceModel] `tfsdk:"ip" json:"ip,computed"`
Okta customfield.NestedObject[ZeroTrustAccessPoliciesExcludeOktaDataSourceModel] `tfsdk:"okta" json:"okta,computed"`
SAML customfield.NestedObject[ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel] `tfsdk:"saml" json:"saml,computed"`
ServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel] `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPoliciesExcludeGroupDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPoliciesExcludeAuthContextDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AcID types.String `tfsdk:"ac_id" json:"ac_id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel struct {
AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessPoliciesExcludeCommonNameDataSourceModel struct {
CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPoliciesExcludeGeoDataSourceModel struct {
CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel struct {
IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel struct {
Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesExcludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel struct {
EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
KeysURL types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Team types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeLoginMethodDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPDataSourceModel struct {
IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesExcludeOktaDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel struct {
AttributeName types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel struct {
TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeDataSourceModel struct {
Group customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGroupDataSourceModel] `tfsdk:"group" json:"group,computed"`
AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
AuthContext customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAuthContextDataSourceModel] `tfsdk:"auth_context" json:"auth_context,computed"`
AuthMethod customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel] `tfsdk:"auth_method" json:"auth_method,computed"`
AzureAD customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel] `tfsdk:"azure_ad" json:"azureAD,computed"`
Certificate customfield.NestedObject[ZeroTrustAccessPoliciesIncludeCertificateDataSourceModel] `tfsdk:"certificate" json:"certificate,computed"`
CommonName customfield.NestedObject[ZeroTrustAccessPoliciesIncludeCommonNameDataSourceModel] `tfsdk:"common_name" json:"common_name,computed"`
Geo customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGeoDataSourceModel] `tfsdk:"geo" json:"geo,computed"`
DevicePosture customfield.NestedObject[ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel] `tfsdk:"device_posture" json:"device_posture,computed"`
EmailDomain customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel] `tfsdk:"email_domain" json:"email_domain,computed"`
EmailList customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel] `tfsdk:"email_list" json:"email_list,computed"`
Email customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailDataSourceModel] `tfsdk:"email" json:"email,computed"`
Everyone customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEveryoneDataSourceModel] `tfsdk:"everyone" json:"everyone,computed"`
ExternalEvaluation customfield.NestedObject[ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
GitHubOrganization customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
GSuite customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel] `tfsdk:"gsuite" json:"gsuite,computed"`
LoginMethod customfield.NestedObject[ZeroTrustAccessPoliciesIncludeLoginMethodDataSourceModel] `tfsdk:"login_method" json:"login_method,computed"`
IPList customfield.NestedObject[ZeroTrustAccessPoliciesIncludeIPListDataSourceModel] `tfsdk:"ip_list" json:"ip_list,computed"`
IP customfield.NestedObject[ZeroTrustAccessPoliciesIncludeIPDataSourceModel] `tfsdk:"ip" json:"ip,computed"`
Okta customfield.NestedObject[ZeroTrustAccessPoliciesIncludeOktaDataSourceModel] `tfsdk:"okta" json:"okta,computed"`
SAML customfield.NestedObject[ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel] `tfsdk:"saml" json:"saml,computed"`
ServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel] `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPoliciesIncludeGroupDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPoliciesIncludeAuthContextDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AcID types.String `tfsdk:"ac_id" json:"ac_id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel struct {
AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeCertificateDataSourceModel struct {
}

type ZeroTrustAccessPoliciesIncludeCommonNameDataSourceModel struct {
CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPoliciesIncludeGeoDataSourceModel struct {
CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel struct {
IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel struct {
Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesIncludeEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel struct {
EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
KeysURL types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Team types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeLoginMethodDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPDataSourceModel struct {
IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesIncludeOktaDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel struct {
AttributeName types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel struct {
TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesRequireDataSourceModel struct {
Group customfield.NestedObject[ZeroTrustAccessPoliciesRequireGroupDataSourceModel] `tfsdk:"group" json:"group,computed"`
AnyValidServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesRequireAnyValidServiceTokenDataSourceModel] `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
AuthContext customfield.NestedObject[ZeroTrustAccessPoliciesRequireAuthContextDataSourceModel] `tfsdk:"auth_context" json:"auth_context,computed"`
AuthMethod customfield.NestedObject[ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel] `tfsdk:"auth_method" json:"auth_method,computed"`
AzureAD customfield.NestedObject[ZeroTrustAccessPoliciesRequireAzureADDataSourceModel] `tfsdk:"azure_ad" json:"azureAD,computed"`
Certificate customfield.NestedObject[ZeroTrustAccessPoliciesRequireCertificateDataSourceModel] `tfsdk:"certificate" json:"certificate,computed"`
CommonName customfield.NestedObject[ZeroTrustAccessPoliciesRequireCommonNameDataSourceModel] `tfsdk:"common_name" json:"common_name,computed"`
Geo customfield.NestedObject[ZeroTrustAccessPoliciesRequireGeoDataSourceModel] `tfsdk:"geo" json:"geo,computed"`
DevicePosture customfield.NestedObject[ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel] `tfsdk:"device_posture" json:"device_posture,computed"`
EmailDomain customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel] `tfsdk:"email_domain" json:"email_domain,computed"`
EmailList customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailListDataSourceModel] `tfsdk:"email_list" json:"email_list,computed"`
Email customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailDataSourceModel] `tfsdk:"email" json:"email,computed"`
Everyone customfield.NestedObject[ZeroTrustAccessPoliciesRequireEveryoneDataSourceModel] `tfsdk:"everyone" json:"everyone,computed"`
ExternalEvaluation customfield.NestedObject[ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
GitHubOrganization customfield.NestedObject[ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
GSuite customfield.NestedObject[ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel] `tfsdk:"gsuite" json:"gsuite,computed"`
LoginMethod customfield.NestedObject[ZeroTrustAccessPoliciesRequireLoginMethodDataSourceModel] `tfsdk:"login_method" json:"login_method,computed"`
IPList customfield.NestedObject[ZeroTrustAccessPoliciesRequireIPListDataSourceModel] `tfsdk:"ip_list" json:"ip_list,computed"`
IP customfield.NestedObject[ZeroTrustAccessPoliciesRequireIPDataSourceModel] `tfsdk:"ip" json:"ip,computed"`
Okta customfield.NestedObject[ZeroTrustAccessPoliciesRequireOktaDataSourceModel] `tfsdk:"okta" json:"okta,computed"`
SAML customfield.NestedObject[ZeroTrustAccessPoliciesRequireSAMLDataSourceModel] `tfsdk:"saml" json:"saml,computed"`
ServiceToken customfield.NestedObject[ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel] `tfsdk:"service_token" json:"service_token,computed"`
}

type ZeroTrustAccessPoliciesRequireGroupDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireAnyValidServiceTokenDataSourceModel struct {
}

type ZeroTrustAccessPoliciesRequireAuthContextDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
AcID types.String `tfsdk:"ac_id" json:"ac_id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel struct {
AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesRequireAzureADDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireCertificateDataSourceModel struct {
}

type ZeroTrustAccessPoliciesRequireCommonNameDataSourceModel struct {
CommonName types.String `tfsdk:"common_name" json:"common_name,computed"`
}

type ZeroTrustAccessPoliciesRequireGeoDataSourceModel struct {
CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel struct {
IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel struct {
Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesRequireEveryoneDataSourceModel struct {
}

type ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel struct {
EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
KeysURL types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
Team types.String `tfsdk:"team" json:"team,computed"`
}

type ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel struct {
Email types.String `tfsdk:"email" json:"email,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireLoginMethodDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireIPListDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireIPDataSourceModel struct {
IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesRequireOktaDataSourceModel struct {
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesRequireSAMLDataSourceModel struct {
AttributeName types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel struct {
TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}
