// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessPoliciesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessPoliciesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessPoliciesDataSourceModel struct {
	AppID     types.String                                                               `tfsdk:"app_id" path:"app_id,required"`
	AccountID types.String                                                               `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                               `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                                `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessPoliciesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessPoliciesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessPoliciesResultDataSourceModel struct {
	ID                           types.String                                                                       `tfsdk:"id" json:"id,computed"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel] `tfsdk:"approval_groups" json:"approval_groups,computed"`
	ApprovalRequired             types.Bool                                                                         `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                                                  `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Decision                     types.String                                                                       `tfsdk:"decision" json:"decision,computed"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPoliciesExcludeDataSourceModel]        `tfsdk:"exclude" json:"exclude,computed"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessPoliciesIncludeDataSourceModel]        `tfsdk:"include" json:"include,computed"`
	IsolationRequired            types.Bool                                                                         `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                                                       `tfsdk:"name" json:"name,computed"`
	PurposeJustificationPrompt   types.String                                                                       `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
	PurposeJustificationRequired types.Bool                                                                         `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPoliciesRequireDataSourceModel]        `tfsdk:"require" json:"require,computed"`
	SessionDuration              types.String                                                                       `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                                                  `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

type ZeroTrustAccessPoliciesApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64                  `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed"`
	EmailListUUID   types.String                   `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type ZeroTrustAccessPoliciesExcludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                      `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPoliciesExcludeIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPoliciesExcludeIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                      `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPoliciesExcludeOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                      `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPoliciesExcludeGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesExcludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesIncludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                      `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPoliciesIncludeIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPoliciesIncludeIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                      `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPoliciesIncludeOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                      `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPoliciesIncludeGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesIncludeSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPoliciesRequireDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                      `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPoliciesRequireIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPoliciesRequireIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                      `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPoliciesRequireGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPoliciesRequireAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPoliciesRequireOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPoliciesRequireSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                      `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPoliciesRequireGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPoliciesRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPoliciesRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPoliciesRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireOktaDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPoliciesRequireSAMLDataSourceModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPoliciesRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessPoliciesRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPoliciesRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPoliciesRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPoliciesRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
