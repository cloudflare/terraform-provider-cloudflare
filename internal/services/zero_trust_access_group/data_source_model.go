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

type ZeroTrustAccessGroupResultDataSourceEnvelope struct {
	Result ZeroTrustAccessGroupDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessGroupResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessGroupDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessGroupDataSourceModel struct {
	AccountID types.String                                     `tfsdk:"account_id" path:"account_id"`
	GroupID   types.String                                     `tfsdk:"group_id" path:"group_id"`
	ZoneID    types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt timetypes.RFC3339                                `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt timetypes.RFC3339                                `tfsdk:"updated_at" json:"updated_at,computed"`
	ID        types.String                                     `tfsdk:"id" json:"id,computed_optional"`
	Name      types.String                                     `tfsdk:"name" json:"name,computed_optional"`
	Exclude   *[]*ZeroTrustAccessGroupExcludeDataSourceModel   `tfsdk:"exclude" json:"exclude,computed_optional"`
	Include   *[]*ZeroTrustAccessGroupIncludeDataSourceModel   `tfsdk:"include" json:"include,computed_optional"`
	IsDefault *[]*ZeroTrustAccessGroupIsDefaultDataSourceModel `tfsdk:"is_default" json:"is_default,computed_optional"`
	Require   *[]*ZeroTrustAccessGroupRequireDataSourceModel   `tfsdk:"require" json:"require,computed_optional"`
	Filter    *ZeroTrustAccessGroupFindOneByDataSourceModel    `tfsdk:"filter"`
}

func (m *ZeroTrustAccessGroupDataSourceModel) toReadParams() (params zero_trust.AccessGroupGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessGroupDataSourceModel) toListParams() (params zero_trust.AccessGroupListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessGroupListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessGroupExcludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupExcludeEmailDataSourceModel              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            *ZeroTrustAccessGroupExcludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          *ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   *ZeroTrustAccessGroupExcludeIPDataSourceModel                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               *ZeroTrustAccessGroupExcludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                *ZeroTrustAccessGroupExcludeGroupDataSourceModel              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              *ZeroTrustAccessGroupExcludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   *ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               *ZeroTrustAccessGroupExcludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 *ZeroTrustAccessGroupExcludeOktaDataSourceModel               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 *ZeroTrustAccessGroupExcludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         *ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  *ZeroTrustAccessGroupExcludeGeoDataSourceModel                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           *ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        *ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIncludeDataSourceModel struct {
	Email                *ZeroTrustAccessGroupIncludeEmailDataSourceModel              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            *ZeroTrustAccessGroupIncludeEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          *ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   *ZeroTrustAccessGroupIncludeIPDataSourceModel                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               *ZeroTrustAccessGroupIncludeIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                *ZeroTrustAccessGroupIncludeGroupDataSourceModel              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              *ZeroTrustAccessGroupIncludeAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   *ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               *ZeroTrustAccessGroupIncludeGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 *ZeroTrustAccessGroupIncludeOktaDataSourceModel               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 *ZeroTrustAccessGroupIncludeSAMLDataSourceModel               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         *ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  *ZeroTrustAccessGroupIncludeGeoDataSourceModel                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           *ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        *ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupIsDefaultDataSourceModel struct {
	Email                *ZeroTrustAccessGroupIsDefaultEmailDataSourceModel              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            *ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          *ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                            `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   *ZeroTrustAccessGroupIsDefaultIPDataSourceModel                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               *ZeroTrustAccessGroupIsDefaultIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                            `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                *ZeroTrustAccessGroupIsDefaultGroupDataSourceModel              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              *ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   *ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               *ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 *ZeroTrustAccessGroupIsDefaultOktaDataSourceModel               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 *ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         *ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                            `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  *ZeroTrustAccessGroupIsDefaultGeoDataSourceModel                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           *ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        *ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupIsDefaultEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupIsDefaultIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupIsDefaultAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupIsDefaultGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupIsDefaultSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupIsDefaultServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupIsDefaultExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupIsDefaultGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupIsDefaultAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupIsDefaultDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupRequireDataSourceModel struct {
	Email                *ZeroTrustAccessGroupRequireEmailDataSourceModel              `tfsdk:"email" json:"email,computed_optional"`
	EmailList            *ZeroTrustAccessGroupRequireEmailListDataSourceModel          `tfsdk:"email_list" json:"email_list,computed_optional"`
	EmailDomain          *ZeroTrustAccessGroupRequireEmailDomainDataSourceModel        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	Everyone             jsontypes.Normalized                                          `tfsdk:"everyone" json:"everyone,computed_optional"`
	IP                   *ZeroTrustAccessGroupRequireIPDataSourceModel                 `tfsdk:"ip" json:"ip,computed_optional"`
	IPList               *ZeroTrustAccessGroupRequireIPListDataSourceModel             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	Certificate          jsontypes.Normalized                                          `tfsdk:"certificate" json:"certificate,computed_optional"`
	Group                *ZeroTrustAccessGroupRequireGroupDataSourceModel              `tfsdk:"group" json:"group,computed_optional"`
	AzureAD              *ZeroTrustAccessGroupRequireAzureADDataSourceModel            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	GitHubOrganization   *ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               *ZeroTrustAccessGroupRequireGSuiteDataSourceModel             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	Okta                 *ZeroTrustAccessGroupRequireOktaDataSourceModel               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 *ZeroTrustAccessGroupRequireSAMLDataSourceModel               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         *ZeroTrustAccessGroupRequireServiceTokenDataSourceModel       `tfsdk:"service_token" json:"service_token,computed_optional"`
	AnyValidServiceToken jsontypes.Normalized                                          `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed_optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	Geo                  *ZeroTrustAccessGroupRequireGeoDataSourceModel                `tfsdk:"geo" json:"geo,computed_optional"`
	AuthMethod           *ZeroTrustAccessGroupRequireAuthMethodDataSourceModel         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	DevicePosture        *ZeroTrustAccessGroupRequireDevicePostureDataSourceModel      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
}

type ZeroTrustAccessGroupRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type ZeroTrustAccessGroupRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type ZeroTrustAccessGroupRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type ZeroTrustAccessGroupRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type ZeroTrustAccessGroupRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type ZeroTrustAccessGroupRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type ZeroTrustAccessGroupRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type ZeroTrustAccessGroupRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type ZeroTrustAccessGroupRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type ZeroTrustAccessGroupRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type ZeroTrustAccessGroupRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type ZeroTrustAccessGroupFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
