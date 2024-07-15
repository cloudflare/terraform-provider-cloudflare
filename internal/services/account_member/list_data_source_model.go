// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMembersResultListDataSourceEnvelope struct {
	Result *[]*AccountMembersItemsDataSourceModel `json:"result,computed"`
}

type AccountMembersDataSourceModel struct {
	AccountID types.String                           `tfsdk:"account_id" path:"account_id"`
	Direction types.String                           `tfsdk:"direction" query:"direction"`
	Order     types.String                           `tfsdk:"order" query:"order"`
	Page      types.Float64                          `tfsdk:"page" query:"page"`
	PerPage   types.Float64                          `tfsdk:"per_page" query:"per_page"`
	Status    types.String                           `tfsdk:"status" query:"status"`
	MaxItems  types.Int64                            `tfsdk:"max_items"`
	Items     *[]*AccountMembersItemsDataSourceModel `tfsdk:"items"`
}

type AccountMembersItemsDataSourceModel struct {
	ID       types.String                                   `tfsdk:"id" json:"id,computed"`
	Policies *[]*AccountMembersItemsPoliciesDataSourceModel `tfsdk:"policies" json:"policies"`
	Roles    *[]*AccountMembersItemsRolesDataSourceModel    `tfsdk:"roles" json:"roles"`
	Status   types.String                                   `tfsdk:"status" json:"status,computed"`
}

type AccountMembersItemsPoliciesDataSourceModel struct {
	ID               types.String                                                   `tfsdk:"id" json:"id,computed"`
	Access           types.String                                                   `tfsdk:"access" json:"access"`
	PermissionGroups *[]*AccountMembersItemsPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups"`
	ResourceGroups   *[]*AccountMembersItemsPoliciesResourceGroupsDataSourceModel   `tfsdk:"resource_groups" json:"resource_groups"`
}

type AccountMembersItemsPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String         `tfsdk:"id" json:"id,computed"`
	Meta jsontypes.Normalized `tfsdk:"meta" json:"meta"`
	Name types.String         `tfsdk:"name" json:"name,computed"`
}

type AccountMembersItemsPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                      `tfsdk:"id" json:"id,computed"`
	Scope *[]*AccountMembersItemsPoliciesResourceGroupsScopeDataSourceModel `tfsdk:"scope" json:"scope,computed"`
	Meta  jsontypes.Normalized                                              `tfsdk:"meta" json:"meta"`
	Name  types.String                                                      `tfsdk:"name" json:"name,computed"`
}

type AccountMembersItemsPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                             `tfsdk:"key" json:"key,computed"`
	Objects *[]*AccountMembersItemsPoliciesResourceGroupsScopeObjectsDataSourceModel `tfsdk:"objects" json:"objects,computed"`
}

type AccountMembersItemsPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type AccountMembersItemsRolesDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	Description types.String    `tfsdk:"description" json:"description,computed"`
	Name        types.String    `tfsdk:"name" json:"name,computed"`
	Permissions *[]types.String `tfsdk:"permissions" json:"permissions,computed"`
}
