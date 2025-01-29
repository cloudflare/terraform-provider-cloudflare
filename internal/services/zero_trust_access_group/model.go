// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultEnvelope struct {
	Result ZeroTrustAccessGroupModel `json:"result"`
}

type ZeroTrustAccessGroupModel struct {
	ID        types.String                                                   `tfsdk:"id" json:"id,computed"`
	AccountID types.String                                                   `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                   `tfsdk:"zone_id" path:"zone_id,optional"`
	Name      types.String                                                   `tfsdk:"name" json:"name,required"`
	Include   *[]*ZeroTrustAccessGroupIncludeModel                           `tfsdk:"include" json:"include,required"`
	IsDefault types.Bool                                                     `tfsdk:"is_default" json:"is_default,optional"`
	Exclude   customfield.NestedObjectList[ZeroTrustAccessGroupExcludeModel] `tfsdk:"exclude" json:"exclude,computed_optional"`
	Require   customfield.NestedObjectList[ZeroTrustAccessGroupRequireModel] `tfsdk:"require" json:"require,computed_optional"`
	CreatedAt timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustAccessGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustAccessGroupModel) MarshalJSONForUpdate(state ZeroTrustAccessGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustAccessGroupIncludeModel struct {
	Group                *ZeroTrustAccessGroupIncludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessGroupIncludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessGroupIncludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessGroupIncludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessGroupIncludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessGroupIncludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessGroupIncludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessGroupIncludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessGroupIncludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessGroupIncludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessGroupIncludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessGroupIncludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessGroupIncludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupIncludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessGroupIncludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessGroupIncludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	IPList               *ZeroTrustAccessGroupIncludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessGroupIncludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessGroupIncludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessGroupIncludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessGroupIncludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
}

type ZeroTrustAccessGroupIncludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessGroupIncludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupIncludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupIncludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupIncludeCertificateModel struct {
}

type ZeroTrustAccessGroupIncludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessGroupIncludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupIncludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessGroupIncludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupIncludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupIncludeEveryoneModel struct {
}

type ZeroTrustAccessGroupIncludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupIncludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessGroupIncludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupIncludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupIncludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupIncludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupIncludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupIncludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessGroupExcludeModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupExcludeGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupExcludeAuthContextModel]        `tfsdk:"auth_context" json:"auth_context,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupExcludeAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupExcludeAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	Certificate          *ZeroTrustAccessGroupExcludeCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupExcludeCommonNameModel]         `tfsdk:"common_name" json:"common_name,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupExcludeGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupExcludeDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupExcludeEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	Everyone             *ZeroTrustAccessGroupExcludeEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupExcludeExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupExcludeGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupExcludeGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupExcludeIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupExcludeIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupExcludeOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupExcludeSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupExcludeServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
}

type ZeroTrustAccessGroupExcludeGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessGroupExcludeAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupExcludeAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupExcludeAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupExcludeCertificateModel struct {
}

type ZeroTrustAccessGroupExcludeCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessGroupExcludeGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupExcludeDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessGroupExcludeEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupExcludeEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupExcludeEveryoneModel struct {
}

type ZeroTrustAccessGroupExcludeExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupExcludeGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessGroupExcludeGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupExcludeIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupExcludeIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupExcludeOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupExcludeSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupExcludeServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}

type ZeroTrustAccessGroupRequireModel struct {
	Group                customfield.NestedObject[ZeroTrustAccessGroupRequireGroupModel]              `tfsdk:"group" json:"group,computed_optional"`
	AnyValidServiceToken *ZeroTrustAccessGroupRequireAnyValidServiceTokenModel                        `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          customfield.NestedObject[ZeroTrustAccessGroupRequireAuthContextModel]        `tfsdk:"auth_context" json:"auth_context,computed_optional"`
	AuthMethod           customfield.NestedObject[ZeroTrustAccessGroupRequireAuthMethodModel]         `tfsdk:"auth_method" json:"auth_method,computed_optional"`
	AzureAD              customfield.NestedObject[ZeroTrustAccessGroupRequireAzureADModel]            `tfsdk:"azure_ad" json:"azureAD,computed_optional"`
	Certificate          *ZeroTrustAccessGroupRequireCertificateModel                                 `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           customfield.NestedObject[ZeroTrustAccessGroupRequireCommonNameModel]         `tfsdk:"common_name" json:"common_name,computed_optional"`
	Geo                  customfield.NestedObject[ZeroTrustAccessGroupRequireGeoModel]                `tfsdk:"geo" json:"geo,computed_optional"`
	DevicePosture        customfield.NestedObject[ZeroTrustAccessGroupRequireDevicePostureModel]      `tfsdk:"device_posture" json:"device_posture,computed_optional"`
	EmailDomain          customfield.NestedObject[ZeroTrustAccessGroupRequireEmailDomainModel]        `tfsdk:"email_domain" json:"email_domain,computed_optional"`
	EmailList            customfield.NestedObject[ZeroTrustAccessGroupRequireEmailListModel]          `tfsdk:"email_list" json:"email_list,computed_optional"`
	Email                customfield.NestedObject[ZeroTrustAccessGroupRequireEmailModel]              `tfsdk:"email" json:"email,computed_optional"`
	Everyone             *ZeroTrustAccessGroupRequireEveryoneModel                                    `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   customfield.NestedObject[ZeroTrustAccessGroupRequireExternalEvaluationModel] `tfsdk:"external_evaluation" json:"external_evaluation,computed_optional"`
	GitHubOrganization   customfield.NestedObject[ZeroTrustAccessGroupRequireGitHubOrganizationModel] `tfsdk:"github_organization" json:"github-organization,computed_optional"`
	GSuite               customfield.NestedObject[ZeroTrustAccessGroupRequireGSuiteModel]             `tfsdk:"gsuite" json:"gsuite,computed_optional"`
	IPList               customfield.NestedObject[ZeroTrustAccessGroupRequireIPListModel]             `tfsdk:"ip_list" json:"ip_list,computed_optional"`
	IP                   customfield.NestedObject[ZeroTrustAccessGroupRequireIPModel]                 `tfsdk:"ip" json:"ip,computed_optional"`
	Okta                 customfield.NestedObject[ZeroTrustAccessGroupRequireOktaModel]               `tfsdk:"okta" json:"okta,computed_optional"`
	SAML                 customfield.NestedObject[ZeroTrustAccessGroupRequireSAMLModel]               `tfsdk:"saml" json:"saml,computed_optional"`
	ServiceToken         customfield.NestedObject[ZeroTrustAccessGroupRequireServiceTokenModel]       `tfsdk:"service_token" json:"service_token,computed_optional"`
}

type ZeroTrustAccessGroupRequireGroupModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireAnyValidServiceTokenModel struct {
}

type ZeroTrustAccessGroupRequireAuthContextModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	AcID               types.String `tfsdk:"ac_id" json:"ac_id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupRequireAuthMethodModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,required"`
}

type ZeroTrustAccessGroupRequireAzureADModel struct {
	ID                 types.String `tfsdk:"id" json:"id,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupRequireCertificateModel struct {
}

type ZeroTrustAccessGroupRequireCommonNameModel struct {
	CommonName types.String `tfsdk:"common_name" json:"common_name,required"`
}

type ZeroTrustAccessGroupRequireGeoModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,required"`
}

type ZeroTrustAccessGroupRequireDevicePostureModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,required"`
}

type ZeroTrustAccessGroupRequireEmailDomainModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,required"`
}

type ZeroTrustAccessGroupRequireEmailListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireEmailModel struct {
	Email types.String `tfsdk:"email" json:"email,required"`
}

type ZeroTrustAccessGroupRequireEveryoneModel struct {
}

type ZeroTrustAccessGroupRequireExternalEvaluationModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,required"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,required"`
}

type ZeroTrustAccessGroupRequireGitHubOrganizationModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
	Team               types.String `tfsdk:"team" json:"team,optional"`
}

type ZeroTrustAccessGroupRequireGSuiteModel struct {
	Email              types.String `tfsdk:"email" json:"email,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupRequireIPListModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type ZeroTrustAccessGroupRequireIPModel struct {
	IP types.String `tfsdk:"ip" json:"ip,required"`
}

type ZeroTrustAccessGroupRequireOktaModel struct {
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
	Name               types.String `tfsdk:"name" json:"name,required"`
}

type ZeroTrustAccessGroupRequireSAMLModel struct {
	AttributeName      types.String `tfsdk:"attribute_name" json:"attribute_name,required"`
	AttributeValue     types.String `tfsdk:"attribute_value" json:"attribute_value,required"`
	IdentityProviderID types.String `tfsdk:"identity_provider_id" json:"identity_provider_id,required"`
}

type ZeroTrustAccessGroupRequireServiceTokenModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,required"`
}
