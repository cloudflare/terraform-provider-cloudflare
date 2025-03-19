// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomPagesResultEnvelope struct {
	Result CustomPagesModel `json:"result"`
}

type CustomPagesModel struct {
	Identifier types.String `tfsdk:"identifier" path:"identifier,required"`
	AccountID  types.String `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID     types.String `tfsdk:"zone_id" path:"zone_id,optional"`
	State      types.String `tfsdk:"state" json:"state,required"`
	URL        types.String `tfsdk:"url" json:"url,required"`
}

func (m CustomPagesModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CustomPagesModel) MarshalJSONForUpdate(state CustomPagesModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
