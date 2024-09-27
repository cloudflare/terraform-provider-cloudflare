// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
	AppID                        types.String                                                                     `tfsdk:"app_id" path:"app_id,optional"`
	PolicyID                     types.String                                                                     `tfsdk:"policy_id" path:"policy_id,optional"`
	ZoneID                       types.String                                                                     `tfsdk:"zone_id" path:"zone_id,optional"`
	ApprovalRequired             types.Bool                                                                       `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Decision                     types.String                                                                     `tfsdk:"decision" json:"decision,computed"`
	ID                           types.String                                                                     `tfsdk:"id" json:"id,computed"`
	IsolationRequired            types.Bool                                                                       `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                                                     `tfsdk:"name" json:"name,computed"`
	PurposeJustificationPrompt   types.String                                                                     `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
	PurposeJustificationRequired types.Bool                                                                       `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	SessionDuration              types.String                                                                     `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                                                `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	ApprovalGroups               customfield.NestedObjectList[ZeroTrustAccessPolicyApprovalGroupsDataSourceModel] `tfsdk:"approval_groups" json:"approval_groups,computed"`
	Exclude                      customfield.NestedObjectList[ZeroTrustAccessPolicyExcludeDataSourceModel]        `tfsdk:"exclude" json:"exclude,computed"`
	Include                      customfield.NestedObjectList[ZeroTrustAccessPolicyIncludeDataSourceModel]        `tfsdk:"include" json:"include,computed"`
	Require                      customfield.NestedObjectList[ZeroTrustAccessPolicyRequireDataSourceModel]        `tfsdk:"require" json:"require,computed"`
	Filter                       *ZeroTrustAccessPolicyFindOneByDataSourceModel                                   `tfsdk:"filter"`
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toReadParams(_ context.Context) (params zero_trust.AccessApplicationPolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessApplicationPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessPolicyApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64                  `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  customfield.List[types.String] `tfsdk:"email_addresses" json:"email_addresses,computed"`
	EmailListUUID   types.String                   `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type ZeroTrustAccessPolicyExcludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                    `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyExcludeIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                    `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyExcludeGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyExcludeAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyExcludeSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyExcludeServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyExcludeGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyExcludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
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

type ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyIncludeDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                    `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyIncludeIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                    `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyIncludeGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyIncludeAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyIncludeSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyIncludeServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyIncludeGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyIncludeAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
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

type ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyRequireDataSourceModel struct {
	Email                customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailDataSourceModel]              `tfsdk:"email" json:"email,computed"`
	EmailList            customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailListDataSourceModel]          `tfsdk:"email_list" json:"email_list,computed"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel]        `tfsdk:"email_domain" json:"email_domain,computed"`
	Everyone             jsontypes.Normalized                                                                    `tfsdk:"everyone" json:"everyone,computed"`
	IP                   customfield.NestedObject[ZeroTrustAccessPolicyRequireIPDataSourceModel]                 `tfsdk:"ip" json:"ip,computed"`
	IPList               customfield.NestedObject[ZeroTrustAccessPolicyRequireIPListDataSourceModel]             `tfsdk:"ip_list" json:"ip_list,computed"`
	Certificate          jsontypes.Normalized                                                                    `tfsdk:"certificate" json:"certificate,computed"`
	Group                customfield.NestedObject[ZeroTrustAccessPolicyRequireGroupDataSourceModel]              `tfsdk:"group" json:"group,computed"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessPolicyRequireAzureADDataSourceModel]            `tfsdk:"azure_ad" json:"azureAD,computed"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel] `tfsdk:"github_organization" json:"github-organization,computed"`
	GSuite               customfield.NestedObject[ZeroTrustAccessPolicyRequireGSuiteDataSourceModel]             `tfsdk:"gsuite" json:"gsuite,computed"`
	Okta                 customfield.NestedObject[ZeroTrustAccessPolicyRequireOktaDataSourceModel]               `tfsdk:"okta" json:"okta,computed"`
	SAML                 customfield.NestedObject[ZeroTrustAccessPolicyRequireSAMLDataSourceModel]               `tfsdk:"saml" json:"saml,computed"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessPolicyRequireServiceTokenDataSourceModel]       `tfsdk:"service_token" json:"service_token,computed"`
	AnyValidServiceToken jsontypes.Normalized                                                                    `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed"`
	Geo                  customfield.NestedObject[ZeroTrustAccessPolicyRequireGeoDataSourceModel]                `tfsdk:"geo" json:"geo,computed"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel]         `tfsdk:"auth_method" json:"auth_method,computed"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel]      `tfsdk:"device_posture" json:"device_posture,computed"`
}

type ZeroTrustAccessPolicyRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessPolicyRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessPolicyRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessPolicyRequireAzureADDataSourceModel struct {
	ID                 types.String `tfsdk:"id" json:"id,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
	Name               types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyRequireGSuiteDataSourceModel struct {
	Email              types.String `tfsdk:"email" json:"email,computed"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,computed"`
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

type ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessPolicyRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessPolicyFindOneByDataSourceModel struct {
	AppID     types.String `tfsdk:"app_id" path:"app_id,required"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id,optional"`
}
