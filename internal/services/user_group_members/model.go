// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_group_members

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserGroupMembersResultEnvelope struct {
	Result *[]*UserGroupMembersMembersModel `json:"result"`
}

type UserGroupMembersModel struct {
	ID          types.String                     `tfsdk:"id" json:"-,computed"`
	UserGroupID types.String                     `tfsdk:"user_group_id" path:"user_group_id,required"`
	AccountID   types.String                     `tfsdk:"account_id" path:"account_id,required"`
	FuzzyEmail  types.String                     `tfsdk:"fuzzy_email" query:"fuzzyEmail,optional"`
	Direction   types.String                     `tfsdk:"direction" query:"direction,computed_optional"`
	Page        types.Float64                    `tfsdk:"page" query:"page,computed_optional"`
	PerPage     types.Float64                    `tfsdk:"per_page" query:"per_page,computed_optional"`
	Members     *[]*UserGroupMembersMembersModel `tfsdk:"members" json:"members,required"`
}

func (m UserGroupMembersModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Members)
}

func (m UserGroupMembersModel) MarshalJSONForUpdate(state UserGroupMembersModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Members, state.Members)
}

type UserGroupMembersMembersModel struct {
	ID types.String `tfsdk:"id" json:"id,required"`
}
