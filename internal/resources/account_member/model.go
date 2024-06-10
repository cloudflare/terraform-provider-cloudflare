// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_member

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccountMemberResultEnvelope struct {
	Result AccountMemberModel `json:"result,computed"`
}

type AccountMemberModel struct {
	AccountID types.String    `tfsdk:"account_id" path:"account_id"`
	MemberID  types.String    `tfsdk:"member_id" path:"member_id"`
	Email     types.String    `tfsdk:"email" json:"email"`
	Roles     *[]types.String `tfsdk:"roles" json:"roles"`
	Status    types.String    `tfsdk:"status" json:"status"`
}
