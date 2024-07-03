// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_policy

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessPoliciesResultListDataSourceEnvelope struct {
	Result *[]*AccessPoliciesItemsDataSourceModel `json:"result,computed"`
}

type AccessPoliciesDataSourceModel struct {
	AppID     types.String                           `tfsdk:"app_id" path:"app_id"`
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                           `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Items     *[]*AccessPoliciesItemsDataSourceModel `tfsdk:"items"`
}

type AccessPoliciesItemsDataSourceModel struct {
	ID                           types.String                                         `tfsdk:"id" json:"id,computed"`
	ApprovalGroups               *[]*AccessPoliciesItemsApprovalGroupsDataSourceModel `tfsdk:"approval_groups" json:"approval_groups,computed"`
	ApprovalRequired             types.Bool                                           `tfsdk:"approval_required" json:"approval_required,computed"`
	CreatedAt                    types.String                                         `tfsdk:"created_at" json:"created_at,computed"`
	Decision                     types.String                                         `tfsdk:"decision" json:"decision,computed"`
	Exclude                      *[]*AccessPoliciesItemsExcludeDataSourceModel        `tfsdk:"exclude" json:"exclude,computed"`
	Include                      *[]*AccessPoliciesItemsIncludeDataSourceModel        `tfsdk:"include" json:"include,computed"`
	IsolationRequired            types.Bool                                           `tfsdk:"isolation_required" json:"isolation_required,computed"`
	Name                         types.String                                         `tfsdk:"name" json:"name,computed"`
	PurposeJustificationPrompt   types.String                                         `tfsdk:"purpose_justification_prompt" json:"purpose_justification_prompt,computed"`
	PurposeJustificationRequired types.Bool                                           `tfsdk:"purpose_justification_required" json:"purpose_justification_required,computed"`
	Require                      *[]*AccessPoliciesItemsRequireDataSourceModel        `tfsdk:"require" json:"require,computed"`
	SessionDuration              types.String                                         `tfsdk:"session_duration" json:"session_duration,computed"`
	UpdatedAt                    types.String                                         `tfsdk:"updated_at" json:"updated_at,computed"`
}

type AccessPoliciesItemsApprovalGroupsDataSourceModel struct {
	ApprovalsNeeded types.Float64   `tfsdk:"approvals_needed" json:"approvals_needed,computed"`
	EmailAddresses  *[]types.String `tfsdk:"email_addresses" json:"email_addresses,computed"`
	EmailListUUID   types.String    `tfsdk:"email_list_uuid" json:"email_list_uuid,computed"`
}

type AccessPoliciesItemsExcludeDataSourceModel struct {
	Everyone             types.String `tfsdk:"everyone" json:"everyone,computed"`
	Certificate          types.String `tfsdk:"certificate" json:"certificate,computed"`
	AnyValidServiceToken types.String `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
}

type AccessPoliciesItemsExcludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsExcludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsExcludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsExcludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsExcludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsExcludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsExcludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsExcludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsExcludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsExcludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsExcludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsExcludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesItemsIncludeDataSourceModel struct {
	Everyone             types.String `tfsdk:"everyone" json:"everyone,computed"`
	Certificate          types.String `tfsdk:"certificate" json:"certificate,computed"`
	AnyValidServiceToken types.String `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
}

type AccessPoliciesItemsIncludeEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsIncludeIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsIncludeIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsIncludeAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsIncludeGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsIncludeGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsIncludeSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsIncludeServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsIncludeExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsIncludeGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsIncludeAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsIncludeDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}

type AccessPoliciesItemsRequireDataSourceModel struct {
	Everyone             types.String `tfsdk:"everyone" json:"everyone,computed"`
	Certificate          types.String `tfsdk:"certificate" json:"certificate,computed"`
	AnyValidServiceToken types.String `tfsdk:"any_valid_service_token" json:"any_valid_service_token,computed"`
}

type AccessPoliciesItemsRequireEmailDataSourceModel struct {
	Email types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireEmailListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireEmailDomainDataSourceModel struct {
	Domain types.String `tfsdk:"domain" json:"domain,computed"`
}

type AccessPoliciesItemsRequireIPDataSourceModel struct {
	IP types.String `tfsdk:"ip" json:"ip,computed"`
}

type AccessPoliciesItemsRequireIPListDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireGroupDataSourceModel struct {
	ID types.String `tfsdk:"id" json:"id,computed"`
}

type AccessPoliciesItemsRequireAzureADDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
}

type AccessPoliciesItemsRequireGitHubOrganizationDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Name         types.String `tfsdk:"name" json:"name,computed"`
}

type AccessPoliciesItemsRequireGSuiteDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireOktaDataSourceModel struct {
	ConnectionID types.String `tfsdk:"connection_id" json:"connection_id,computed"`
	Email        types.String `tfsdk:"email" json:"email,computed"`
}

type AccessPoliciesItemsRequireSAMLDataSourceModel struct {
	AttributeName  types.String `tfsdk:"attribute_name" json:"attribute_name,computed"`
	AttributeValue types.String `tfsdk:"attribute_value" json:"attribute_value,computed"`
}

type AccessPoliciesItemsRequireServiceTokenDataSourceModel struct {
	TokenID types.String `tfsdk:"token_id" json:"token_id,computed"`
}

type AccessPoliciesItemsRequireExternalEvaluationDataSourceModel struct {
	EvaluateURL types.String `tfsdk:"evaluate_url" json:"evaluate_url,computed"`
	KeysURL     types.String `tfsdk:"keys_url" json:"keys_url,computed"`
}

type AccessPoliciesItemsRequireGeoDataSourceModel struct {
	CountryCode types.String `tfsdk:"country_code" json:"country_code,computed"`
}

type AccessPoliciesItemsRequireAuthMethodDataSourceModel struct {
	AuthMethod types.String `tfsdk:"auth_method" json:"auth_method,computed"`
}

type AccessPoliciesItemsRequireDevicePostureDataSourceModel struct {
	IntegrationUID types.String `tfsdk:"integration_uid" json:"integration_uid,computed"`
}
