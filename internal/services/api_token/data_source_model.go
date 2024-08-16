// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokenResultDataSourceEnvelope struct {
	Result APITokenDataSourceModel `json:"result,computed"`
}

type APITokenResultListDataSourceEnvelope struct {
	Result *[]*APITokenDataSourceModel `json:"result,computed"`
}

type APITokenDataSourceModel struct {
	TokenID    types.String                        `tfsdk:"token_id" path:"token_id"`
	ExpiresOn  timetypes.RFC3339                   `tfsdk:"expires_on" json:"expires_on"`
	ID         types.String                        `tfsdk:"id" json:"id"`
	IssuedOn   timetypes.RFC3339                   `tfsdk:"issued_on" json:"issued_on"`
	LastUsedOn timetypes.RFC3339                   `tfsdk:"last_used_on" json:"last_used_on"`
	ModifiedOn timetypes.RFC3339                   `tfsdk:"modified_on" json:"modified_on"`
	Name       types.String                        `tfsdk:"name" json:"name"`
	NotBefore  timetypes.RFC3339                   `tfsdk:"not_before" json:"not_before"`
	Status     types.String                        `tfsdk:"status" json:"status"`
	Condition  *APITokenConditionDataSourceModel   `tfsdk:"condition" json:"condition"`
	Policies   *[]*APITokenPoliciesDataSourceModel `tfsdk:"policies" json:"policies"`
	Filter     *APITokenFindOneByDataSourceModel   `tfsdk:"filter"`
}

type APITokenConditionDataSourceModel struct {
	RequestIP *APITokenConditionRequestIPDataSourceModel `tfsdk:"request_ip" json:"request.ip"`
}

type APITokenConditionRequestIPDataSourceModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

type APITokenPoliciesDataSourceModel struct {
	ID               types.String                                        `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                        `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups *[]*APITokenPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        jsontypes.Normalized                                `tfsdk:"resources" json:"resources,computed"`
}

type APITokenPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String         `tfsdk:"id" json:"id,computed"`
	Meta jsontypes.Normalized `tfsdk:"meta" json:"meta"`
	Name types.String         `tfsdk:"name" json:"name,computed"`
}

type APITokenFindOneByDataSourceModel struct {
	Direction types.String `tfsdk:"direction" query:"direction"`
}
