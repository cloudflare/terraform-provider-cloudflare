// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_policy

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
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
	AccountID                    types.String                                           `tfsdk:"account_id" path:"account_id"`
	AppID                        types.String                                           `tfsdk:"app_id" path:"app_id"`
	PolicyID                     types.String                                           `tfsdk:"policy_id" path:"policy_id"`
	ZoneID                       types.String                                           `tfsdk:"zone_id" path:"zone_id"`
	ApprovalRequired             types.Bool                                             `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    timetypes.RFC3339                                      `tfsdk:"created_at" json:"created_at,computed"`
	IsolationRequired            types.Bool                                             `tfsdk:"isolation_required" json:"isolation_required,computed"`
	PurposeJustificationRequired types.Bool                                             `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	SessionDuration              types.String                                           `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    timetypes.RFC3339                                      `tfsdk:"updated_at" json:"updated_at,computed"`
	Decision                     types.String                                           `tfsdk:"decision" json:"decision"`
	ID                           types.String                                           `tfsdk:"id" json:"id"`
	Name                         types.String                                           `tfsdk:"name" json:"name"`
	PurposeJustificationPrompt   types.String                                           `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt"`
	ApprovalGroups               *[]*ZeroTrustAccessPolicyApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups"`
	Exclude                      *[]*ZeroTrustAccessPolicyExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude"`
	Include                      *[]*ZeroTrustAccessPolicyIncludeDataSourceModel        `tfsdk:"include" json:"include"`
	Require                      *[]*ZeroTrustAccessPolicyRequireDataSourceModel        `tfsdk:"require" json:"require"`
	Filter                       *ZeroTrustAccessPolicyFindOneByDataSourceModel         `tfsdk:"filter"`
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toReadParams() (params zero_trust.AccessApplicationPolicyGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessPolicyDataSourceModel) toListParams() (params zero_trust.AccessApplicationPolicyListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessApplicationPolicyListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessPolicyApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid"`
}

type ZeroTrustAccessPolicyExcludeDataSourceModel struct {
	Email                *ZeroTrustAccessPolicyExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
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
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPolicyExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
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
	Email                *ZeroTrustAccessPolicyIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
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
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPolicyIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
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
	Email                *ZeroTrustAccessPolicyRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessPolicyRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessPolicyRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessPolicyRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessPolicyRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessPolicyRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessPolicyRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessPolicyRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessPolicyRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessPolicyRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessPolicyRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessPolicyRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessPolicyRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessPolicyRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessPolicyRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
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
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessPolicyRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessPolicyRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessPolicyRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
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
	AppID     types.String `tfsdk:"app_id" path:"app_id"`
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
