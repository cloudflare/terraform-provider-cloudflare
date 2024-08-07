// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMembersResultListDataSourceEnvelope struct {
	Result *[]*AccountMembersResultDataSourceModel `json:"result,computed"`
}

type AccountMembersDataSourceModel struct {
	AccountID types.String                            `tfsdk:"account_id" path:"account_id"`
	Direction types.String                            `tfsdk:"direction" query:"direction"`
	Order     types.String                            `tfsdk:"order" query:"order"`
	Status    types.String                            `tfsdk:"status" query:"status"`
	MaxItems  types.Int64                             `tfsdk:"max_items"`
	Result    *[]*AccountMembersResultDataSourceModel `tfsdk:"result"`
}

type AccountMembersResultDataSourceModel struct {
	ID       types.String                                                `tfsdk:"id" json:"id,computed"`
	Policies *[]*AccountMembersPoliciesDataSourceModel                   `tfsdk:"policies" json:"policies"`
	Roles    *[]*AccountMembersRolesDataSourceModel                      `tfsdk:"roles" json:"roles"`
	Status   types.String                                                `tfsdk:"status" json:"status,computed"`
	User     customfield.NestedObject[AccountMembersUserDataSourceModel] `tfsdk:"user" json:"user,computed"`
}

type AccountMembersPoliciesDataSourceModel struct {
	ID               types.String                                              `tfsdk:"id" json:"id,computed"`
	Access           types.String                                              `tfsdk:"access" json:"access"`
	PermissionGroups *[]*AccountMembersPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups"`
	ResourceGroups   *[]*AccountMembersPoliciesResourceGroupsDataSourceModel   `tfsdk:"resource_groups" json:"resource_groups"`
}

type AccountMembersPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String         `tfsdk:"id" json:"id,computed"`
	Meta jsontypes.Normalized `tfsdk:"meta" json:"meta"`
	Name types.String         `tfsdk:"name" json:"name,computed"`
}

type AccountMembersPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                 `tfsdk:"id" json:"id,computed"`
	Scope *[]*AccountMembersPoliciesResourceGroupsScopeDataSourceModel `tfsdk:"scope" json:"scope,computed"`
	Meta  jsontypes.Normalized                                         `tfsdk:"meta" json:"meta"`
	Name  types.String                                                 `tfsdk:"name" json:"name,computed"`
}

type AccountMembersPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                        `tfsdk:"key" json:"key,computed"`
	Objects *[]*AccountMembersPoliciesResourceGroupsScopeObjectsDataSourceModel `tfsdk:"objects" json:"objects,computed"`
}

type AccountMembersPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type AccountMembersRolesDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	Description types.String    `tfsdk:"description" json:"description,computed"`
	Name        types.String    `tfsdk:"name" json:"name,computed"`
	Permissions *[]types.String `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountMembersUserDataSourceModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}
