// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserGroupResultEnvelope struct {
	Result UserGroupModel `json:"result"`
}

type UserGroupModel struct {
	ID         types.String               `tfsdk:"id" json:"id,computed"`
	AccountID  types.String               `tfsdk:"account_id" path:"account_id,required"`
	Name       types.String               `tfsdk:"name" json:"name,required"`
	Policies   *[]*UserGroupPoliciesModel `tfsdk:"policies" json:"policies,optional"`
	CreatedOn  timetypes.RFC3339          `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn timetypes.RFC3339          `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m UserGroupModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserGroupModel) MarshalJSONForUpdate(state UserGroupModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type UserGroupPoliciesModel struct {
	Access           types.String                               `tfsdk:"access" json:"access,required"`
	PermissionGroups *[]*UserGroupPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	ResourceGroups   *[]*UserGroupPoliciesResourceGroupsModel   `tfsdk:"resource_groups" json:"resource_groups,required"`
}

type UserGroupPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type UserGroupPoliciesResourceGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}
