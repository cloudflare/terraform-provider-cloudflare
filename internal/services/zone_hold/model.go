// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZoneHoldResultEnvelope struct {
	Result ZoneHoldModel `json:"result"`
}

type ZoneHoldModel struct {
	ID                types.String `tfsdk:"id" json:"-,computed"`
	ZoneID            types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Hold              types.Bool   `tfsdk:"hold" json:"hold,computed"`
	HoldAfter         types.String `tfsdk:"hold_after" json:"hold_after,computed"`
	IncludeSubdomains types.String `tfsdk:"include_subdomains" json:"include_subdomains,computed"`
}

func (m ZoneHoldModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZoneHoldModel) MarshalJSONForUpdate(state ZoneHoldModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
