// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_token

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APITokensResultListDataSourceEnvelope struct {
	Result *[]*APITokensResultDataSourceModel `json:"result,computed"`
}

type APITokensDataSourceModel struct {
	Direction types.String                       `tfsdk:"direction" query:"direction"`
	MaxItems  types.Int64                        `tfsdk:"max_items"`
	Result    *[]*APITokensResultDataSourceModel `tfsdk:"result"`
}

type APITokensResultDataSourceModel struct {
	ID         types.String                         `tfsdk:"id" json:"id,computed"`
	Condition  *APITokensConditionDataSourceModel   `tfsdk:"condition" json:"condition"`
	ExpiresOn  timetypes.RFC3339                    `tfsdk:"expires_on" json:"expires_on"`
	IssuedOn   timetypes.RFC3339                    `tfsdk:"issued_on" json:"issued_on,computed"`
	LastUsedOn timetypes.RFC3339                    `tfsdk:"last_used_on" json:"last_used_on,computed"`
	ModifiedOn timetypes.RFC3339                    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name       types.String                         `tfsdk:"name" json:"name"`
	NotBefore  timetypes.RFC3339                    `tfsdk:"not_before" json:"not_before"`
	Policies   *[]*APITokensPoliciesDataSourceModel `tfsdk:"policies" json:"policies"`
	Status     types.String                         `tfsdk:"status" json:"status"`
}

type APITokensConditionDataSourceModel struct {
	RequestIP *APITokensConditionRequestIPDataSourceModel `tfsdk:"request_ip" json:"request.ip"`
}

type APITokensConditionRequestIPDataSourceModel struct {
	In    *[]types.String `tfsdk:"in" json:"in"`
	NotIn *[]types.String `tfsdk:"not_in" json:"not_in"`
}

type APITokensPoliciesDataSourceModel struct {
	ID               types.String                                         `tfsdk:"id" json:"id,computed"`
	Effect           types.String                                         `tfsdk:"effect" json:"effect,computed"`
	PermissionGroups *[]*APITokensPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups,computed"`
	Resources        jsontypes.Normalized                                 `tfsdk:"resources" json:"resources,computed"`
}

type APITokensPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String         `tfsdk:"id" json:"id,computed"`
	Meta jsontypes.Normalized `tfsdk:"meta" json:"meta"`
	Name types.String         `tfsdk:"name" json:"name,computed"`
}
