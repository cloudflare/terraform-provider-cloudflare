// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterResultEnvelope struct {
	Result *[]*FilterBodyModel `json:"result"`
}

type FilterModel struct {
	ID          types.String        `tfsdk:"id" json:"id,computed"`
	ZoneID      types.String        `tfsdk:"zone_id" path:"zone_id,required"`
	Body        *[]*FilterBodyModel `tfsdk:"body" json:"body,required,no_refresh"`
	Description types.String        `tfsdk:"description" json:"description,optional"`
	Expression  types.String        `tfsdk:"expression" json:"expression,optional"`
	Paused      types.Bool          `tfsdk:"paused" json:"paused,optional"`
	Ref         types.String        `tfsdk:"ref" json:"ref,optional"`
}

func (m FilterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m FilterModel) MarshalJSONForUpdate(state FilterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m.Body, state.Body)
}

type FilterBodyModel struct {
	ID          types.String `tfsdk:"id" json:"id,computed"`
	Description types.String `tfsdk:"description" json:"description,optional"`
	Expression  types.String `tfsdk:"expression" json:"expression,optional"`
	Paused      types.Bool   `tfsdk:"paused" json:"paused,optional"`
	Ref         types.String `tfsdk:"ref" json:"ref,optional"`
}
