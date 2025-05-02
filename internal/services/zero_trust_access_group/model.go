// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessGroupResultEnvelope struct {
	Result ZeroTrustAccessGroupModel `json:"result"`
}

type ZeroTrustAccessGroupModel struct {
	ID        types.String                         `tfsdk:"id" json:"id,computed"`
	AccountID types.String                         `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                         `tfsdk:"zone_id" path:"zone_id,optional"`
	Name      types.String                         `tfsdk:"name" json:"name,required"`
	Include   *[]*ZeroTrustAccessGroupIncludeModel `tfsdk:"include" json:"include,required"`
	IsDefault types.Bool                           `tfsdk:"is_default" json:"is_default,optional"`
	Exclude   *[]*ZeroTrustAccessGroupExcludeModel `tfsdk:"exclude" json:"exclude,optional"`
	Require   *[]*ZeroTrustAccessGroupRequireModel `tfsdk:"require" json:"require,optional"`
	CreatedAt timetypes.RFC3339                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt timetypes.RFC3339                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
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
	LoginMethod          *ZeroTrustAccessGroupIncludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
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

type ZeroTrustAccessGroupIncludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
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
	Group                *ZeroTrustAccessGroupExcludeGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessGroupExcludeAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessGroupExcludeAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessGroupExcludeAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessGroupExcludeAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessGroupExcludeCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessGroupExcludeCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessGroupExcludeGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessGroupExcludeDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessGroupExcludeEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessGroupExcludeEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessGroupExcludeEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessGroupExcludeEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupExcludeExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessGroupExcludeGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessGroupExcludeGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessGroupExcludeLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessGroupExcludeIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessGroupExcludeIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessGroupExcludeOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessGroupExcludeSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessGroupExcludeServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
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

type ZeroTrustAccessGroupExcludeLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
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
	Group                *ZeroTrustAccessGroupRequireGroupModel                `tfsdk:"group" json:"group,optional"`
	AnyValidServiceToken *ZeroTrustAccessGroupRequireAnyValidServiceTokenModel `tfsdk:"any_valid_service_token" json:"any_valid_service_token,optional"`
	AuthContext          *ZeroTrustAccessGroupRequireAuthContextModel          `tfsdk:"auth_context" json:"auth_context,optional"`
	AuthMethod           *ZeroTrustAccessGroupRequireAuthMethodModel           `tfsdk:"auth_method" json:"auth_method,optional"`
	AzureAD              *ZeroTrustAccessGroupRequireAzureADModel              `tfsdk:"azure_ad" json:"azureAD,optional"`
	Certificate          *ZeroTrustAccessGroupRequireCertificateModel          `tfsdk:"certificate" json:"certificate,optional"`
	CommonName           *ZeroTrustAccessGroupRequireCommonNameModel           `tfsdk:"common_name" json:"common_name,optional"`
	Geo                  *ZeroTrustAccessGroupRequireGeoModel                  `tfsdk:"geo" json:"geo,optional"`
	DevicePosture        *ZeroTrustAccessGroupRequireDevicePostureModel        `tfsdk:"device_posture" json:"device_posture,optional"`
	EmailDomain          *ZeroTrustAccessGroupRequireEmailDomainModel          `tfsdk:"email_domain" json:"email_domain,optional"`
	EmailList            *ZeroTrustAccessGroupRequireEmailListModel            `tfsdk:"email_list" json:"email_list,optional"`
	Email                *ZeroTrustAccessGroupRequireEmailModel                `tfsdk:"email" json:"email,optional"`
	Everyone             *ZeroTrustAccessGroupRequireEveryoneModel             `tfsdk:"everyone" json:"everyone,optional"`
	ExternalEvaluation   *ZeroTrustAccessGroupRequireExternalEvaluationModel   `tfsdk:"external_evaluation" json:"external_evaluation,optional"`
	GitHubOrganization   *ZeroTrustAccessGroupRequireGitHubOrganizationModel   `tfsdk:"github_organization" json:"github-organization,optional"`
	GSuite               *ZeroTrustAccessGroupRequireGSuiteModel               `tfsdk:"gsuite" json:"gsuite,optional"`
	LoginMethod          *ZeroTrustAccessGroupRequireLoginMethodModel          `tfsdk:"login_method" json:"login_method,optional"`
	IPList               *ZeroTrustAccessGroupRequireIPListModel               `tfsdk:"ip_list" json:"ip_list,optional"`
	IP                   *ZeroTrustAccessGroupRequireIPModel                   `tfsdk:"ip" json:"ip,optional"`
	Okta                 *ZeroTrustAccessGroupRequireOktaModel                 `tfsdk:"okta" json:"okta,optional"`
	SAML                 *ZeroTrustAccessGroupRequireSAMLModel                 `tfsdk:"saml" json:"saml,optional"`
	ServiceToken         *ZeroTrustAccessGroupRequireServiceTokenModel         `tfsdk:"service_token" json:"service_token,optional"`
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

type ZeroTrustAccessGroupRequireLoginMethodModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
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
