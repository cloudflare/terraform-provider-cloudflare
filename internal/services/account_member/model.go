// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultEnvelope struct {
	Result AccountMemberModel `json:"result,computed"`
}

type AccountMemberModel struct {
	ID        types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID types.String                   `tfsdk:"account_id" path:"account_id"`
	Roles     *[]types.String                `tfsdk:"roles" json:"roles"`
	Policies  *[]*AccountMemberPoliciesModel `tfsdk:"policies" json:"policies"`
	Email     types.String                   `tfsdk:"email" json:"email"`
	Status    types.String                   `tfsdk:"status" json:"status,computed"`
}

type AccountMemberPoliciesModel struct {
	ID               types.String                                   `tfsdk:"id" json:"id,computed"`
	Access           types.String                                   `tfsdk:"access" json:"access"`
	PermissionGroups *[]*AccountMemberPoliciesPermissionGroupsModel `tfsdk:"permission_groups" json:"permission_groups"`
	ResourceGroups   *[]*AccountMemberPoliciesResourceGroupsModel   `tfsdk:"resource_groups" json:"resource_groups"`
}

type AccountMemberPoliciesPermissionGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}

type AccountMemberPoliciesResourceGroupsModel struct {
	ID types.String `tfsdk:"id" json:"id"`
}
