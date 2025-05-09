// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultEnvelope struct {
	Result AccountMemberModel `json:"result"`
}

type AccountMemberModel struct {
	ID        types.String                                             `tfsdk:"id" json:"id,computed"`
	AccountID types.String                                             `tfsdk:"account_id" path:"account_id,required"`
	Email     types.String                                             `tfsdk:"email" json:"email,required,no_refresh"`
	Status    types.String                                             `tfsdk:"status" json:"status,computed_optional"`
	Roles     *[]types.String                                          `tfsdk:"roles" json:"roles,optional,no_refresh"`
	Policies  customfield.NestedObjectList[AccountMemberPoliciesModel] `tfsdk:"policies" json:"policies,computed_optional"`
	User      customfield.NestedObject[AccountMemberUserModel]         `tfsdk:"user" json:"user,computed"`
}

func (m AccountMemberModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccountMemberModel) MarshalJSONForUpdate(state AccountMemberModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccountMemberPoliciesModel struct {
	ID               types.String                                   `tfsdk:"id" json:"id,computed"`
	Access           types.String                                   `tfsdk:"access" json:"access,required"`
	PermissionGroups *[]*AccountMemberPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups,required"`
	ResourceGroups   *[]*AccountMemberPoliciesResourceGroupsModel   `tfsdk:"resource_groups" json:"resource_groups,required"`
}

type AccountMemberPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type AccountMemberPoliciesResourceGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}

type AccountMemberUserModel struct {
	Email                          types.String `tfsdk:"email" json:"email,computed"`
	ID                             types.String `tfsdk:"id" json:"id,computed"`
	FirstName                      types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName                       types.String `tfsdk:"last_name" json:"last_name,computed"`
	TwoFactorAuthenticationEnabled types.Bool   `tfsdk:"two_factor_authentication_enabled" json:"two_factor_authentication_enabled,computed"`
}
