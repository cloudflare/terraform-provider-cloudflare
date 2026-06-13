// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_app

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlagshipAppResultEnvelope struct {
	Result FlagshipAppModel `json:"result"`
}

type FlagshipAppModel struct {
	ID        types.String `tfsdk:"id" json:"id,computed"`
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
	CreatedAt types.String `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt types.String `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy types.String `tfsdk:"updated_by" json:"updated_by,computed"`
}

func (m FlagshipAppModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FlagshipAppModel) MarshalJSONForUpdate(state FlagshipAppModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
