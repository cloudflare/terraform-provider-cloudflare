// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultEnvelope struct {
	Result FilterModel `json:"result"`
}

type FilterModel struct {
	ZoneID      types.String `tfsdk:"zone_id" path:"zone_id,required"`
	FilterID    types.String `tfsdk:"filter_id" path:"filter_id,optional"`
	Expression  types.String `tfsdk:"expression" json:"expression,required"`
	Description types.String `tfsdk:"description" json:"description,computed"`
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,computed"`
	Ref         types.String `tfsdk:"ref" json:"ref,computed"`
}

func (m FilterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FilterModel) MarshalJSONForUpdate(state FilterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
