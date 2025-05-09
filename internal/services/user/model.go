// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserResultEnvelope struct {
	Result UserModel `json:"result"`
}

type UserModel struct {
	Country   types.String `tfsdk:"country" json:"country,optional,no_refresh"`
	FirstName types.String `tfsdk:"first_name" json:"first_name,optional,no_refresh"`
	LastName  types.String `tfsdk:"last_name" json:"last_name,optional,no_refresh"`
	Telephone types.String `tfsdk:"telephone" json:"telephone,optional,no_refresh"`
	Zipcode   types.String `tfsdk:"zipcode" json:"zipcode,optional,no_refresh"`
}

func (m UserModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m UserModel) MarshalJSONForUpdate(state UserModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}
