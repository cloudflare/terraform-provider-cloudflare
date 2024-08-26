// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessGroupsResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessGroupsDataSourceModel struct {
	AccountID types.String                                                             `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                                             `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                                              `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessGroupsDataSourceModel) toListParams() (params zero_trust.AccessGroupListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessGroupsResultDataSourceModel struct {
	ID        types.String                                      `tfsdk:"id" json:"id"`
	CreatedAt timetypes.RFC3339                                 `tfsdk:"created_at" json:"created_at,computed"`
	Exclude   *[]*ZeroTrustAccessGroupsExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude"`
	Include   *[]*ZeroTrustAccessGroupsIncludeDataSourceModel   `tfsdk:"include" json:"include"`
	IsDefault *[]*ZeroTrustAccessGroupsIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default"`
	Name      types.String                                      `tfsdk:"name" json:"name"`
	Require   *[]*ZeroTrustAccessGroupsRequireDataSourceModel   `tfsdk:"require" json:"require"`
	UpdatedAt timetypes.RFC3339                                 `tfsdk:"updated_at" json:"updated_at,computed"`
}

type ZeroTrustAccessGroupsExcludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupsExcludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupsExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupsExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupsExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupsExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupsExcludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupsExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupsExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupsExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupsExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupsExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupsExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupsExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupsExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupsExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupsExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupsExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsIncludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupsIncludeEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupsIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupsIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupsIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupsIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupsIncludeGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupsIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupsIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupsIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupsIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupsIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupsIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupsIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupsIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupsIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupsIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupsIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsIsDefaultDataSourceModel struct {
	Email                *ZeroTrustAccessGroupsIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupsIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupsIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                             `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupsIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupsIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                             `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupsIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupsIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupsIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupsIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupsIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupsIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupsIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                             `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupsIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupsIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupsIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupsIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupsIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupsIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupsRequireDataSourceModel struct {
	Email                *ZeroTrustAccessGroupsRequireEmailDataSourceModel              `tfsdk:"email" json:"email"`
	EmailList            *ZeroTrustAccessGroupsRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list"`
	EmailDomain          *ZeroTrustAccessGroupsRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain"`
	Everyone             jsontypes.Normalized                                           `tfsdk:"everyone" json:"everyone"`
	IP                   *ZeroTrustAccessGroupsRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip"`
	IPList               *ZeroTrustAccessGroupsRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list"`
	Certificate          jsontypes.Normalized                                           `tfsdk:"certificate" json:"certificate"`
	Group                *ZeroTrustAccessGroupsRequireGroupDataSourceModel              `tfsdk:"group" json:"group"`
	AzureAD              *ZeroTrustAccessGroupsRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD"`
	GitHubOrganization   *ZeroTrustAccessGroupsRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization"`
	GSuite               *ZeroTrustAccessGroupsRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite"`
	Okta                 *ZeroTrustAccessGroupsRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta"`
	SAML                 *ZeroTrustAccessGroupsRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml"`
	ServiceToken         *ZeroTrustAccessGroupsRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token"`
	AnyValidServiceToken jsontypes.Normalized                                           `tfsdk:"any_valid_service_token" json:"any_valid_service_token"`
	ExternalEvaluation   *ZeroTrustAccessGroupsRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation"`
	Geo                  *ZeroTrustAccessGroupsRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo"`
	AuthMethod           *ZeroTrustAccessGroupsRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method"`
	DevicePosture        *ZeroTrustAccessGroupsRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture"`
}

type ZeroTrustAccessGroupsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupsRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupsRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupsRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupsRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
