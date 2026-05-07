// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/iam"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserGroupResultDataSourceEnvelope struct {
	Result UserGroupDataSourceModel `json:"result,computed"`
}

type UserGroupDataSourceModel struct {
	ID          types.String                                                   `tfsdk:"id" path:"user_group_id,computed"`
	UserGroupID types.String                                                   `tfsdk:"user_group_id" path:"user_group_id,optional"`
	AccountID   types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn   timetypes.RFC3339                                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339                                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                   `tfsdk:"name" json:"name,computed"`
	Policies    customfield.NestedObjectList[UserGroupPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
	Filter      *UserGroupFindOneByDataSourceModel                             `tfsdk:"filter"`
}

func (m *UserGroupDataSourceModel) toReadParams(_ context.Context) (params iam.UserGroupGetParams, diags diag.Diagnostics) {
	params = iam.UserGroupGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *UserGroupDataSourceModel) toListParams(_ context.Context) (params iam.UserGroupListParams, diags diag.Diagnostics) {
	params = iam.UserGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Direction.IsNull() {
		params.Direction = cloudflare.F(iam.UserGroupListParamsDirection(m.Filter.Direction.ValueString()))
	}
	if !m.Filter.FuzzyName.IsNull() {
		params.FuzzyName = cloudflare.F(m.Filter.FuzzyName.ValueString())
	}
	if !m.Filter.Name.IsNull() {
		params.Name = cloudflare.F(m.Filter.Name.ValueString())
	}

	return
}

type UserGroupPoliciesDataSourceModel struct {
	ID               types.String                                                                   `tfsdk:"id" json:"id,computed,force_encode,encode_state_for_unknown"`
	Access           types.String                                                                   `tfsdk:"access" json:"access,computed"`
	PermissionGroups customfield.NestedObjectList[UserGroupPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	ResourceGroups   customfield.NestedObjectList[UserGroupPoliciesResourceGroupsDataSourceModel]   `tfsdk:"resource_groups" json:"resource_groups,computed"`
}

type UserGroupPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[UserGroupPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                   `tfsdk:"name" json:"name,computed"`
}

type UserGroupPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type UserGroupPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                                      `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[UserGroupPoliciesResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  customfield.NestedObject[UserGroupPoliciesResourceGroupsMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
	Name  types.String                                                                      `tfsdk:"name" json:"name,computed"`
}

type UserGroupPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                                             `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[UserGroupPoliciesResourceGroupsScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type UserGroupPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type UserGroupPoliciesResourceGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type UserGroupFindOneByDataSourceModel struct {
	ID        types.String `tfsdk:"id" query:"id,optional"`
	Direction types.String `tfsdk:"direction" query:"direction,computed_optional"`
	FuzzyName types.String `tfsdk:"fuzzy_name" query:"fuzzyName,optional"`
	Name      types.String `tfsdk:"name" query:"name,optional"`
}
