// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/accounts"
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
	AccountID types.String                                                       `tfsdk:"account_id" path:"account_id,optional"`
	MemberID  types.String                                                       `tfsdk:"member_id" path:"member_id,optional"`
	ID        types.String                                                       `tfsdk:"id" json:"id,computed"`
	Status    types.String                                                       `tfsdk:"status" json:"status,computed"`
	Policies  customfield.NestedObjectList[AccountMemberPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Roles     customfield.NestedObjectList[AccountMemberRolesDataSourceModel]    `tfsdk:"roles" json:"roles,computed"`
	User      customfield.NestedObject[AccountMemberUserDataSourceModel]         `tfsdk:"user" json:"user,computed"`
	Filter    *AccountMemberFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *AccountMemberDataSourceModel) toReadParams(_ context.Context) (params accounts.MemberGetParams, diags diag.Diagnostics) {
	params = accounts.MemberGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *AccountMemberDataSourceModel) toListParams(_ context.Context) (params accounts.MemberListParams, diags diag.Diagnostics) {
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

type AccountMemberPoliciesDataSourceModel struct {
	ID               types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Access           types.String                                                                       `tfsdk:"access" json:"access,computed"`
	PermissionGroups customfield.NestedObjectList[AccountMemberPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	ResourceGroups   customfield.NestedObjectList[AccountMemberPoliciesResourceGroupsDataSourceModel]   `tfsdk:"resource_groups" json:"resource_groups,computed"`
}

type AccountMemberPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[AccountMemberPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                       `tfsdk:"name" json:"name,computed"`
}

type AccountMemberPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AccountMemberPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                                          `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[AccountMemberPoliciesResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  customfield.NestedObject[AccountMemberPoliciesResourceGroupsMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
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
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type AccountMemberRolesDataSourceModel struct {
	ID          types.String                   `tfsdk:"id" json:"id,computed"`
	Description types.String                   `tfsdk:"description" json:"description,computed"`
	Name        types.String                   `tfsdk:"name" json:"name,computed"`
	Permissions customfield.List[types.String] `tfsdk:"permissions" json:"permissions,computed"`
}

type AccountMemberUserDataSourceModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name,computed"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}

type AccountMemberFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Direction types.String `tfsdk:"direction" query:"direction,optional"`
	Order     types.String `tfsdk:"order" query:"order,optional"`
	Status    types.String `tfsdk:"status" query:"status,optional"`
}
