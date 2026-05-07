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

type UserGroupsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[UserGroupsResultDataSourceModel] `json:"result,computed"`
}

type UserGroupsDataSourceModel struct {
	AccountID types.String                                                  `tfsdk:"account_id" path:"account_id,required"`
	FuzzyName types.String                                                  `tfsdk:"fuzzy_name" query:"fuzzyName,optional"`
	ID        types.String                                                  `tfsdk:"id" query:"id,optional"`
	Name      types.String                                                  `tfsdk:"name" query:"name,optional"`
	Direction types.String                                                  `tfsdk:"direction" query:"direction,computed_optional"`
	MaxItems  types.Int64                                                   `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[UserGroupsResultDataSourceModel] `tfsdk:"result"`
}

func (m *UserGroupsDataSourceModel) toListParams(_ context.Context) (params iam.UserGroupListParams, diags diag.Diagnostics) {
	params = iam.UserGroupListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Direction.IsNull() {
		params.Direction = cloudflare.F(iam.UserGroupListParamsDirection(m.Direction.ValueString()))
	}
	if !m.FuzzyName.IsNull() {
		params.FuzzyName = cloudflare.F(m.FuzzyName.ValueString())
	}
	if !m.Name.IsNull() {
		params.Name = cloudflare.F(m.Name.ValueString())
	}

	return
}

type UserGroupsResultDataSourceModel struct {
	ID         types.String                                                    `tfsdk:"id" json:"id,computed"`
	CreatedOn  timetypes.RFC3339                                               `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339                                               `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name       types.String                                                    `tfsdk:"name" json:"name,computed"`
	Policies   customfield.NestedObjectList[UserGroupsPoliciesDataSourceModel] `tfsdk:"policies" json:"policies,computed"`
}

type UserGroupsPoliciesDataSourceModel struct {
	ID               types.String                                                                    `tfsdk:"id" json:"id,computed,force_encode,encode_state_for_unknown"`
	Access           types.String                                                                    `tfsdk:"access" json:"access,computed"`
	PermissionGroups customfield.NestedObjectList[UserGroupsPoliciesPermissionGroupsDataSourceModel] `tfsdk:"permission_groups" json:"permission_groups,computed"`
	ResourceGroups   customfield.NestedObjectList[UserGroupsPoliciesResourceGroupsDataSourceModel]   `tfsdk:"resource_groups" json:"resource_groups,computed"`
}

type UserGroupsPoliciesPermissionGroupsDataSourceModel struct {
	ID   types.String                                                                    `tfsdk:"id" json:"id,computed"`
	Meta customfield.NestedObject[UserGroupsPoliciesPermissionGroupsMetaDataSourceModel] `tfsdk:"meta" json:"meta,computed"`
	Name types.String                                                                    `tfsdk:"name" json:"name,computed"`
}

type UserGroupsPoliciesPermissionGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type UserGroupsPoliciesResourceGroupsDataSourceModel struct {
	ID    types.String                                                                       `tfsdk:"id" json:"id,computed"`
	Scope customfield.NestedObjectList[UserGroupsPoliciesResourceGroupsScopeDataSourceModel] `tfsdk:"scope" json:"scope,computed"`
	Meta  customfield.NestedObject[UserGroupsPoliciesResourceGroupsMetaDataSourceModel]      `tfsdk:"meta" json:"meta,computed"`
	Name  types.String                                                                       `tfsdk:"name" json:"name,computed"`
}

type UserGroupsPoliciesResourceGroupsScopeDataSourceModel struct {
	Key     types.String                                                                              `tfsdk:"key" json:"key,computed"`
	Objects customfield.NestedObjectList[UserGroupsPoliciesResourceGroupsScopeObjectsDataSourceModel] `tfsdk:"objects" json:"objects,computed"`
}

type UserGroupsPoliciesResourceGroupsScopeObjectsDataSourceModel struct {
	Key types.String `tfsdk:"key" json:"key,computed"`
}

type UserGroupsPoliciesResourceGroupsMetaDataSourceModel struct {
	Key   types.String `tfsdk:"key" json:"key,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}
