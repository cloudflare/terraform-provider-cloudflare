// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultDataSourceEnvelope struct {
	Result AccountMemberDataSourceModel `json:"result,computed"`
}

type AccountMemberResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AccountMemberDataSourceModel] `json:"result,computed"`
}

type AccountMemberDataSourceModel struct {
	AccountID types.String                                               `tfsdk:"account_id" path:"account_id"`
	MemberID  types.String                                               `tfsdk:"member_id" path:"member_id"`
	ID        types.String                                               `tfsdk:"id" json:"id,computed"`
	Status    types.String                                               `tfsdk:"status" json:"status,computed"`
	User      customfield.NestedObject[AccountMemberUserDataSourceModel] `tfsdk:"user" json:"user,computed"`
	Policies  *[]*AccountMemberPoliciesDataSourceModel                   `tfsdk:"policies" json:"policies,computed_optional"`
	Roles     *[]*AccountMemberRolesDataSourceModel                      `tfsdk:"roles" json:"roles,computed_optional"`
	Filter    *AccountMemberFindOneByDataSourceModel                     `tfsdk:"filter"`
}

func (m *AccountMemberDataSourceModel) toReadParams() (params accounts.MemberGetParams, diags diag.Diagnostics) {
	params = accounts.MemberGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountMemberDataSourceModel) toListParams() (params accounts.MemberListParams, diags diag.Diagnostics) {
	params = accounts.MemberListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(accounts.MemberListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.Order.IsNull() {
		params.Order = cloudflare.F(accounts.MemberListParamsOrder(m.Filter.Order.ValueString()))
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(accounts.MemberListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type AccountMemberUserDataSourceModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name,computed_optional"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name,computed_optional"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}

type AccountMemberPoliciesDataSourceModel struct {
	ID               types.String                                             `tfsdk:"id" json:"id,computed"`
	Access           types.String                                             `tfsdk:"access" json:"access,computed_optional"`
	PermissionGroups *[]*AccountMemberPoliciesPermissionGroupsDataSourceModel `tfsdk:"permission_groups" json:"permission_groups,computed_optional"`
	ResourceGroups   *[]*AccountMemberPoliciesResourceGroupsDataSourceModel   `tfsdk:"resource_groups" json:"resource_groups,computed_optional"`
}

type AccountMemberPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                              `tfsdk:"id" json:"id,computed"`
	Meta *AccountMemberPoliciesPermissionGroupsMetaDataSourceModel `tfsdk:"meta" json:"meta,computed_optional"`
	Name types.String                                              `tfsdk:"name" json:"name,computed"`
}

type AccountMemberPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type AccountMemberPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                                          `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[AccountMemberPoliciesResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  *AccountMemberPoliciesResourceGroupsMetaDataSourceModel                               `tfsdk:"meta" json:"meta,computed_optional"`
	Name  types.String                                                                          `tfsdk:"name" json:"name,computed"`
}

type AccountMemberPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                                                 `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[AccountMemberPoliciesResourceGroupsScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type AccountMemberPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type AccountMemberPoliciesResourceGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed_optional"`
	Value types.String `tfsdk:"value" json:"value,computed_optional"`
}

type AccountMemberRolesDataSourceModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	Name        types.String `tfsdk:"name" json:"name,computed"`
	Permissions types.List   `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountMemberFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	Direction types.String `tfsdk:"direction" query:"direction"`
	Order     types.String `tfsdk:"order" query:"order"`
	Status    types.String `tfsdk:"status" query:"status"`
}