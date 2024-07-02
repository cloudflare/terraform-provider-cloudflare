// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultEnvelope struct {
	Result APITokenModel `json:"result,computed"`
}

type APITokenResultDataSourceEnvelope struct {
	Result APITokenDataSourceModel `json:"result,computed"`
}

type APITokensResultDataSourceEnvelope struct {
	Result APITokensDataSourceModel `json:"result,computed"`
}

type APITokenModel struct {
	TokenID   types.String              `tfsdk:"token_id" path:"token_id"`
	Name      types.String              `tfsdk:"name" json:"name"`
	Policies  *[]*APITokenPoliciesModel `tfsdk:"policies" json:"policies"`
	Condition *APITokenConditionModel   `tfsdk:"condition" json:"condition"`
	ExpiresOn types.String              `tfsdk:"expires_on" json:"expires_on"`
	NotBefore types.String              `tfsdk:"not_before" json:"not_before"`
	Status    types.String              `tfsdk:"status" json:"status"`
	Value     types.String              `tfsdk:"value" json:"value,computed"`
	ID        types.String              `tfsdk:"id" json:"id,computed"`
}

type APITokenPoliciesModel struct {
	ID               types.String                              `tfsdk:"id" json:"id,computed"`
	Effect           types.String                              `tfsdk:"effect" json:"effect"`
	PermissionGroups *[]*APITokenPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups"`
	Resources        types.String                              `tfsdk:"resources" json:"resources"`
}

type APITokenPoliciesPermissionGroupsModel struct {
	ID   types.String `tfsdk:"id" json:"id,computed"`
	Meta types.String `tfsdk:"meta" json:"meta"`
	Name types.String `tfsdk:"name" json:"name,computed"`
}

type APITokenConditionModel struct {
	RequestIP *APITokenConditionRequestIPModel `tfsdk:"request_ip" json:"request_ip"`
}

type APITokenConditionRequestIPModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

type APITokenDataSourceModel struct {
}

type APITokensDataSourceModel struct {
}
